// Code generated by go-swagger; DO NOT EDIT.

package ui

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/nicolas2bert/ba-server/gen/models"
)

// GetPhotosOKCode is the HTTP code returned for type GetPhotosOK
const GetPhotosOKCode int = 200

/*GetPhotosOK list of photos

swagger:response getPhotosOK
*/
type GetPhotosOK struct {

	/*
	  In: Body
	*/
	Payload models.Photos `json:"body,omitempty"`
}

// NewGetPhotosOK creates GetPhotosOK with default headers values
func NewGetPhotosOK() *GetPhotosOK {

	return &GetPhotosOK{}
}

// WithPayload adds the payload to the get photos o k response
func (o *GetPhotosOK) WithPayload(payload models.Photos) *GetPhotosOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get photos o k response
func (o *GetPhotosOK) SetPayload(payload models.Photos) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPhotosOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Photos, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetPhotosBadRequestCode is the HTTP code returned for type GetPhotosBadRequest
const GetPhotosBadRequestCode int = 400

/*GetPhotosBadRequest Bad request

swagger:response getPhotosBadRequest
*/
type GetPhotosBadRequest struct {

	/*
	  In: Body
	*/
	Payload *GetPhotosBadRequestBody `json:"body,omitempty"`
}

// NewGetPhotosBadRequest creates GetPhotosBadRequest with default headers values
func NewGetPhotosBadRequest() *GetPhotosBadRequest {

	return &GetPhotosBadRequest{}
}

// WithPayload adds the payload to the get photos bad request response
func (o *GetPhotosBadRequest) WithPayload(payload *GetPhotosBadRequestBody) *GetPhotosBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get photos bad request response
func (o *GetPhotosBadRequest) SetPayload(payload *GetPhotosBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPhotosBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPhotosNotFoundCode is the HTTP code returned for type GetPhotosNotFound
const GetPhotosNotFoundCode int = 404

/*GetPhotosNotFound Not Found

swagger:response getPhotosNotFound
*/
type GetPhotosNotFound struct {
}

// NewGetPhotosNotFound creates GetPhotosNotFound with default headers values
func NewGetPhotosNotFound() *GetPhotosNotFound {

	return &GetPhotosNotFound{}
}

// WriteResponse to the client
func (o *GetPhotosNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetPhotosInternalServerErrorCode is the HTTP code returned for type GetPhotosInternalServerError
const GetPhotosInternalServerErrorCode int = 500

/*GetPhotosInternalServerError Server Error

swagger:response getPhotosInternalServerError
*/
type GetPhotosInternalServerError struct {
}

// NewGetPhotosInternalServerError creates GetPhotosInternalServerError with default headers values
func NewGetPhotosInternalServerError() *GetPhotosInternalServerError {

	return &GetPhotosInternalServerError{}
}

// WriteResponse to the client
func (o *GetPhotosInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
