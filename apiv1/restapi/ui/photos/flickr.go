package photos

import (
	"encoding/json"
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
	Farm        int
	Server      string
	Description struct {
		Content string `json:"_content"`
	}
	Dates struct{ Posted int64 }
}

type photoWithTitleScheme struct {
	ID     string
	Secret string
	Title  struct {
		Content string `json:"_content"`
	}
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
	Photo photoWithTitleScheme
}

var flickr_consumer_secret = os.Getenv("FLICKR_SECRET")

var flickr_consumer_key = os.Getenv("FLICKR_KEY")

var fc = flickrclient.Client{
	ConsumerSecret: flickr_consumer_secret,
	ConsumerKey:    flickr_consumer_key,
}

func getFlickrInfoPhoto(id string, user models.User, wg *sync.WaitGroup) *models.PhotosItems0 {
	defer wg.Done()
	fc.Token = *user.FlickrToken
	fc.SecretToken = *user.FlickrSecretToken

	args := map[string]string{
		"photo_id": id,
	}
	bodyInfo, err := fc.Request("flickr.photos.getInfo", args)
	if err != nil {
		return nil
	}
	var info photosInfoScheme
	json.Unmarshal(bodyInfo, &info)
	item := models.PhotosItems0{
		ID:          id,
		Title:       info.Photo.Title.Content,
		Description: info.Photo.Description.Content,
		URL:         "https://farm" + strconv.Itoa(info.Photo.Farm) + ".staticflickr.com/" + info.Photo.Server + "/" + info.Photo.ID + "_" + info.Photo.Secret + "_c.jpg",
	}
	return &item
}

func getFlickrPhotos(user models.User) ([]photoScheme, middleware.Responder) {

	fc.Token = *user.FlickrToken
	fc.SecretToken = *user.FlickrSecretToken

	args := map[string]string{
		// "user_id": *user.ID,
		"user_id": "147032531@N08",
	}
	body, err := fc.Request("flickr.people.getPhotos", args)
	if err != nil {
		return nil, ui.NewGetPhotosInternalServerError()
	}

	var v photosScheme
	err = json.Unmarshal(body, &v)
	if err != nil {
		return nil, ui.NewGetPhotosInternalServerError()
	}

	if v.Stat != "ok" {
		return nil, ui.NewGetPhotosInternalServerError()
	}

	if len(v.Photos.Photo) == 0 {
		return nil, ui.NewGetPhotosInternalServerError()
	}

	return v.Photos.Photo, nil
}
