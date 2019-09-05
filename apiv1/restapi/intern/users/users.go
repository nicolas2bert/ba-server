package users

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/nicolas2bert/ba-server/apiv1/auth"
	apiContext "github.com/nicolas2bert/ba-server/apiv1/restapi/context"
	"github.com/nicolas2bert/ba-server/gen/models"
	"github.com/nicolas2bert/ba-server/gen/restapi/operations/intern"
	"github.com/nicolas2bert/ba-server/gen/restapi/operations/ui"
)

var usersInMemory = map[string]*models.User{}

func GetUserID(id string) (*models.User, middleware.Responder) {
	if usersInMemory[id] == nil {
		return nil, ui.NewGetPhotosNotFound()
	}
	return usersInMemory[id], nil
}

func GetUsersIDHandler(params intern.GetUsersIDParams, principal *auth.PrincipalBA) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	l := apiContext.GetLog(ctx, "GetUsersIDHandler")
	user, err := GetUserID(params.ID)
	if err != nil {
		return err
	}
	l.Info("ok")
	return intern.NewGetUsersIDOK().WithPayload(user)
}

func SaveUserHandler(params intern.SaveUserParams, principal *auth.PrincipalBA) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	l := apiContext.GetLog(ctx, "SaveUserHandler")

	usersInMemory[*params.User.ID] = params.User

	l.Info("ok")
	return intern.NewSaveUserOK()
}
