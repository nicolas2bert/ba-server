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

// GetUsersIDHandlerFunc turns a function with the right signature into a get users Id handler
type GetUsersIDHandlerFunc func(GetUsersIDParams, *auth.PrincipalBA) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersIDHandlerFunc) Handle(params GetUsersIDParams, principal *auth.PrincipalBA) middleware.Responder {
	return fn(params, principal)
}

// GetUsersIDHandler interface for that can handle valid get users Id params
type GetUsersIDHandler interface {
	Handle(GetUsersIDParams, *auth.PrincipalBA) middleware.Responder
}

// NewGetUsersID creates a new http.Handler for the get users Id operation
func NewGetUsersID(ctx *middleware.Context, handler GetUsersIDHandler) *GetUsersID {
	return &GetUsersID{Context: ctx, Handler: handler}
}

/*GetUsersID swagger:route GET /users/{id} intern getUsersId

Get user from internal ba webserver

*/
type GetUsersID struct {
	Context *middleware.Context
	Handler GetUsersIDHandler
}

func (o *GetUsersID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersIDParams()

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

// GetUsersIDBadRequestBody get users ID bad request body
// swagger:model GetUsersIDBadRequestBody
type GetUsersIDBadRequestBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get users ID bad request body
func (o *GetUsersIDBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetUsersIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
