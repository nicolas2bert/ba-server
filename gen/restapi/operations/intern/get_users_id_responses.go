// Code generated by go-swagger; DO NOT EDIT.

package intern

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/nicolas2bert/ba-server/gen/models"
)

// GetUsersIDOKCode is the HTTP code returned for type GetUsersIDOK
const GetUsersIDOKCode int = 200

/*GetUsersIDOK OK

swagger:response getUsersIdOK
*/
type GetUsersIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUsersIDOK creates GetUsersIDOK with default headers values
func NewGetUsersIDOK() *GetUsersIDOK {

	return &GetUsersIDOK{}
}

// WithPayload adds the payload to the get users Id o k response
func (o *GetUsersIDOK) WithPayload(payload *models.User) *GetUsersIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users Id o k response
func (o *GetUsersIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
