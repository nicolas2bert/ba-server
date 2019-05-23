package users

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/nicolas2bert/ba-server/apiv1/auth"
	"github.com/nicolas2bert/ba-server/gen/models"
	"github.com/nicolas2bert/ba-server/gen/restapi/operations/intern"
)

var usersInMemory = map[string]*models.User{}

func GetUserID(id string) (*models.User, error) {
	return usersInMemory[id], nil
}

func GetUsersIDHandler(params intern.GetUsersIDParams, principal *auth.PrincipalBA) middleware.Responder {
	return intern.NewGetUsersIDOK().WithPayload(usersInMemory[params.ID])
}

func SaveUserHandler(params intern.SaveUserParams, principal *auth.PrincipalBA) middleware.Responder {
	fmt.Printf("params.User!!!!: %v", params.User)
	usersInMemory[*params.User.ID] = params.User
	return intern.NewSaveUserOK()
}
