package users

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/nicolas2bert/ba-server/apiv1/auth"
	apiContext "github.com/nicolas2bert/ba-server/apiv1/restapi/context"
	"github.com/nicolas2bert/ba-server/db"
	"github.com/nicolas2bert/ba-server/gen/models"
	"github.com/nicolas2bert/ba-server/gen/restapi/operations/intern"
)

var usersInMemory = map[string]*models.User{}

func GetUserID(ctx context.Context, id string) (*models.User, middleware.Responder) {
	user, err := db.GetUser(ctx, id)
	if err != nil {
		return nil, intern.NewGetUsersIDInternalServerError()
	}
	return user, nil
}

func GetUsersIDHandler(params intern.GetUsersIDParams, principal *auth.PrincipalBA) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	l := apiContext.GetLog(ctx, "GetUsersIDHandler")
	user, err := GetUserID(ctx, params.ID)
	if err != nil {
		l.Info("fail to get user from user.ID")
		return err
	}
	l.Info("ok")
	return intern.NewGetUsersIDOK().WithPayload(user)
}

func SaveUserHandler(params intern.SaveUserParams, principal *auth.PrincipalBA) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	l := apiContext.GetLog(ctx, "SaveUserHandler")

	err := db.CreateUser(ctx, params.User)
	if err != nil {
		l.WithError(err).Info("fail to create user")
		return intern.NewSaveUserInternalServerError()
	}

	l.Info("ok")
	return intern.NewSaveUserOK()
}
