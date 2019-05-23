package photos

import (
	"fmt"
	"sync"

	"github.com/go-openapi/runtime/middleware"
	"github.com/nicolas2bert/ba-server/apiv1/auth"
	"github.com/nicolas2bert/ba-server/apiv1/restapi/intern/users"
	"github.com/nicolas2bert/ba-server/gen/restapi/operations/ui"
)

func removeNilFromItems(sl []*ui.GetPhotosOKBodyItems0) []*ui.GetPhotosOKBodyItems0 {
	newSlice := []*ui.GetPhotosOKBodyItems0{}
	for _, s := range sl {
		if s != nil {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

func GetPhotosHandler(params ui.GetPhotosParams, principal *auth.PrincipalBA) middleware.Responder {
	fmt.Printf("\n GetPhotosHandler => principal!!: %v \n", principal)
	fmt.Printf("\n flickr_consumer_secret!!!: %v \n", flickr_consumer_secret)
	fmt.Printf("\n flickr_consumer_key!!!: %v \n", flickr_consumer_key)
	userID := params.ID

	u, err := users.GetUserID(userID)
	if err != nil {
		return ui.NewGetPhotosBadRequest()
	}

	photos, resp := getFlickrPhotos(*u)

	items := make([]*ui.GetPhotosOKBodyItems0, len(photos))

	wg := &sync.WaitGroup{}
	if resp != nil {
		return resp
	}
	for i, photo := range photos {
		wg.Add(1)
		go func(items []*ui.GetPhotosOKBodyItems0, i int, photo photoScheme) {
			photoInfo, _ := getFlickrInfoPhoto(photo.ID, *u, wg)
			items[i] = photoInfo
		}(items, i, photo)
	}

	wg.Wait()

	items = removeNilFromItems(items)
	return ui.NewGetPhotosOK().WithPayload(items)
}
