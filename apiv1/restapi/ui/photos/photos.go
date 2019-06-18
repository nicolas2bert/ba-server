package photos

import (
	"sync"

	"github.com/go-openapi/runtime/middleware"
	"github.com/nicolas2bert/ba-server/apiv1/auth"
	apiContext "github.com/nicolas2bert/ba-server/apiv1/restapi/context"
	"github.com/nicolas2bert/ba-server/apiv1/restapi/intern/users"
	"github.com/nicolas2bert/ba-server/gen/models"
	"github.com/nicolas2bert/ba-server/gen/restapi/operations/ui"
)

func removeNilFromItems(sl models.Photos) models.Photos {
	newSlice := models.Photos{}
	for _, s := range sl {
		if s != nil {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

type logkey string

func GetPhotosHandler(params ui.GetPhotosParams, principal *auth.PrincipalBA) middleware.Responder {
	userID := params.ID

	ctx := params.HTTPRequest.Context()
	l := apiContext.GetLog(ctx, "GetPhotosHandler")

	u, err := users.GetUserID(userID)
	if err != nil {
		l.Error("get user")
		return err
	}

	photos, err := getFlickrPhotos(*u)
	if err != nil {
		l.Error("get photos")
		return err
	}

	items := make(models.Photos, len(photos))

	wg := &sync.WaitGroup{}
	for i, photo := range photos {
		wg.Add(1)
		go func(items models.Photos, i int, photo photoScheme) {
			photoInfo := getFlickrInfoPhoto(photo.ID, *u, wg)
			if photoInfo == nil {
				l.Error("get photo infos")
			}
			items[i] = photoInfo
		}(items, i, photo)
	}

	wg.Wait()

	items = removeNilFromItems(items)
	l.Info("ok")
	return ui.NewGetPhotosOK().WithPayload(items)
}
