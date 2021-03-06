// Code generated by go-swagger; DO NOT EDIT.

package intern

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	"github.com/nicolas2bert/ba-server/apiv1/auth"
)

// SaveUserHandlerFunc turns a function with the right signature into a save user handler
type SaveUserHandlerFunc func(SaveUserParams, *auth.PrincipalBA) middleware.Responder

// Handle executing the request and returning a response
func (fn SaveUserHandlerFunc) Handle(params SaveUserParams, principal *auth.PrincipalBA) middleware.Responder {
	return fn(params, principal)
}

// SaveUserHandler interface for that can handle valid save user params
type SaveUserHandler interface {
	Handle(SaveUserParams, *auth.PrincipalBA) middleware.Responder
}

// NewSaveUser creates a new http.Handler for the save user operation
func NewSaveUser(ctx *middleware.Context, handler SaveUserHandler) *SaveUser {
	return &SaveUser{Context: ctx, Handler: handler}
}

/*SaveUser swagger:route POST /users intern saveUser

Create user from internal ba webserver

*/
type SaveUser struct {
	Context *middleware.Context
	Handler SaveUserHandler
}

func (o *SaveUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSaveUserParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *auth.PrincipalBA
	if uprinc != nil {
		principal = uprinc.(*auth.PrincipalBA) // this is really a auth.PrincipalBA, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SaveUserBadRequestBody save user bad request body
// swagger:model SaveUserBadRequestBody
type SaveUserBadRequestBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this save user bad request body
func (o *SaveUserBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SaveUserBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SaveUserBadRequestBody) UnmarshalBinary(b []byte) error {
	var res SaveUserBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
