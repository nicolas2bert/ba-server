package flickrclient

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	SecretToken    string
}

func getSigningBaseString(basicURL string) string {
	u, _ := url.Parse(basicURL)
	uCopy := *u
	uCopy.RawQuery = ""
	endpointURL := uCopy.String()
	endpointURL = endpointURL[:len(endpointURL)]
	requestURL := url.QueryEscape(endpointURL)
	args := u.Query()
	flickrEncoded := strings.Replace(args.Encode(), "+", "%20", -1)
	query := url.QueryEscape(flickrEncoded)
	signature := fmt.Sprintf("%s&%s&%s", "GET", requestURL, query)
	// GET&https%3A%2F%2Fapi.flickr.com%2Fservices%2Frest&format%3Djson%26method%3Dflickr.people.getPhotos%26nojsoncallback%3D1%26oauth_consumer_key%3D<consumer_key>%26oauth_nonce%3D12345678%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1553741660%26oauth_token%3D72157707323250944-a493dcfeb7be7d29%26oauth_version%3D1.0%26user_id%3D147032531%2540N08
	return signature
}

func (c *Client) getOauthSignedURL(basicURL string) string {
	fmt.Printf("\n c.ConsumerKey!!!: %v \n", c.ConsumerKey)
	fmt.Printf("\n c.ConsumerSecret!!!: %v \n", c.ConsumerSecret)
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	oAuthURL := basicURL + "&nojsoncallback=1" +
		"&oauth_nonce=12345678" +
		"&format=json" +
		"&oauth_consumer_key=" + c.ConsumerKey +
		"&oauth_timestamp=" + timestamp +
		"&oauth_signature_method=HMAC-SHA1" +
		"&oauth_version=1.0" +
		"&oauth_token=" + c.Token

	key := fmt.Sprintf("%s&%s", url.QueryEscape(c.ConsumerSecret), url.QueryEscape(c.SecretToken))
	baseString := getSigningBaseString(oAuthURL)

	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(baseString))
	oauthSignature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return oAuthURL + "&oauth_signature=" + oauthSignature
}

func (c *Client) Request(method string, args map[string]string) ([]byte, error) {
	if c.Token == "" || c.SecretToken == "" {
		fmt.Printf("error: missing token")
		return nil, fmt.Errorf("missing token")
	}
	v := url.Values{}
	for key, value := range args {
		v.Add(key, value)
	}

	basicURL := "https://api.flickr.com/services/rest/?" +
		"method=" + method + "&" + v.Encode()

	signedURL := c.getOauthSignedURL(basicURL)

	fmt.Printf("\n signedURL!!!: %v \n", signedURL)

	res, err := http.Get(signedURL)
	if err != nil {
		fmt.Printf("\n GET err!!!: %v \n", err)
		return nil, fmt.Errorf("GET")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("\n read err!!!: %v \n", err)
		return nil, fmt.Errorf("read")
	}

	return body, nil
}
