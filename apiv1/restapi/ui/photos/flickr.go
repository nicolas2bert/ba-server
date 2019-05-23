package photos

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/go-openapi/runtime/middleware"
	flickrclient "github.com/nicolas2bert/ba-server/flickrclient"
	"github.com/nicolas2bert/ba-server/gen/models"
	"github.com/nicolas2bert/ba-server/gen/restapi/operations/ui"
)

type photoScheme struct {
	ID          string
	Secret      string
	Title       string
	Farm        int
	Server      string
	Description struct {
		Content string `json:"_content"`
	}
	Dates struct{ Posted int64 }
}

type photosScheme struct {
	Stat   string
	Photos struct{ Photo []photoScheme }
}

type photosInfoScheme struct {
	Stat  string
	Photo photoScheme
}

var flickr_consumer_secret = os.Getenv("FLICKR_SECRET")

var flickr_consumer_key = os.Getenv("FLICKR_KEY")

var fc = flickrclient.Client{
	ConsumerSecret: flickr_consumer_secret,
	ConsumerKey:    flickr_consumer_key,
}

func getFlickrInfoPhoto(id string, user models.User, wg *sync.WaitGroup) (*ui.GetPhotosOKBodyItems0, error) {

	fc.Token = *user.FlickrToken
	fc.SecretToken = *user.FlickrSecretToken

	args := map[string]string{
		"photo_id": id,
	}
	bodyInfo, err := fc.Request("flickr.photos.getInfo", args)
	if err != nil {
		fmt.Printf("error: GET getInfo")
		return nil, err
	}
	defer wg.Done()
	var info photosInfoScheme
	json.Unmarshal(bodyInfo, &info)
	item := ui.GetPhotosOKBodyItems0{
		ID:          id,
		Description: info.Photo.Description.Content,
		URL:         "https://farm" + strconv.Itoa(info.Photo.Farm) + ".staticflickr.com/" + info.Photo.Server + "/" + info.Photo.ID + "_" + info.Photo.Secret + ".jpg",
	}
	return &item, nil
}

func getFlickrPhotos(user models.User) ([]photoScheme, middleware.Responder) {

	fc.Token = *user.FlickrToken
	fc.SecretToken = *user.FlickrSecretToken

	args := map[string]string{
		"user_id": *user.ID,
	}
	body, err := fc.Request("flickr.people.getPhotos", args)
	if err != nil {
		fmt.Printf("\n err !!! %v \n", err)
		return nil, ui.NewGetPhotosNotFound()
	}

	var v photosScheme
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Printf("\n err !!! %v \n", err)
		return nil, ui.NewGetPhotosNotFound()
	}

	if v.Stat != "ok" {
		return nil, ui.NewGetPhotosNotFound()
	}

	if len(v.Photos.Photo) == 0 {
		return nil, ui.NewGetPhotosNotFound()
	}

	return v.Photos.Photo, nil
}
