package users

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/nicolas2bert/ba-server/apiv1/auth"
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
	return intern.NewGetUsersIDOK().WithPayload(usersInMemory[params.ID])
}

func SaveUserHandler(params intern.SaveUserParams, principal *auth.PrincipalBA) middleware.Responder {
	usersInMemory[*params.User.ID] = params.User
	return intern.NewSaveUserOK()
}
