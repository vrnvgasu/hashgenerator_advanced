// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"service2/models"
)

// GetCheckOKCode is the HTTP code returned for type GetCheckOK
const GetCheckOKCode int = 200

/*
GetCheckOK Success

swagger:response getCheckOK
*/
type GetCheckOK struct {

	/*
	  In: Body
	*/
	Payload models.ArrayOfHash `json:"body,omitempty"`
}

// NewGetCheckOK creates GetCheckOK with default headers values
func NewGetCheckOK() *GetCheckOK {

	return &GetCheckOK{}
}

// WithPayload adds the payload to the get check o k response
func (o *GetCheckOK) WithPayload(payload models.ArrayOfHash) *GetCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get check o k response
func (o *GetCheckOK) SetPayload(payload models.ArrayOfHash) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.ArrayOfHash{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCheckNoContentCode is the HTTP code returned for type GetCheckNoContent
const GetCheckNoContentCode int = 204

/*
GetCheckNoContent No Content

swagger:response getCheckNoContent
*/
type GetCheckNoContent struct {
}

// NewGetCheckNoContent creates GetCheckNoContent with default headers values
func NewGetCheckNoContent() *GetCheckNoContent {

	return &GetCheckNoContent{}
}

// WriteResponse to the client
func (o *GetCheckNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// GetCheckBadRequestCode is the HTTP code returned for type GetCheckBadRequest
const GetCheckBadRequestCode int = 400

/*
GetCheckBadRequest Bad request

swagger:response getCheckBadRequest
*/
type GetCheckBadRequest struct {
}

// NewGetCheckBadRequest creates GetCheckBadRequest with default headers values
func NewGetCheckBadRequest() *GetCheckBadRequest {

	return &GetCheckBadRequest{}
}

// WriteResponse to the client
func (o *GetCheckBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetCheckInternalServerErrorCode is the HTTP code returned for type GetCheckInternalServerError
const GetCheckInternalServerErrorCode int = 500

/*
GetCheckInternalServerError Internal Server Error

swagger:response getCheckInternalServerError
*/
type GetCheckInternalServerError struct {
}

// NewGetCheckInternalServerError creates GetCheckInternalServerError with default headers values
func NewGetCheckInternalServerError() *GetCheckInternalServerError {

	return &GetCheckInternalServerError{}
}

// WriteResponse to the client
func (o *GetCheckInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
