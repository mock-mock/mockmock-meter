// Code generated by go-swagger; DO NOT EDIT.

package mock_mock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/mock-mock/mockmock-meter/backend/api/models"
)

// PostMockOKCode is the HTTP code returned for type PostMockOK
const PostMockOKCode int = 200

/*PostMockOK OK

swagger:response postMockOK
*/
type PostMockOK struct {

	/*
	  In: Body
	*/
	Payload *models.SLACKResponse `json:"body,omitempty"`
}

// NewPostMockOK creates PostMockOK with default headers values
func NewPostMockOK() *PostMockOK {

	return &PostMockOK{}
}

// WithPayload adds the payload to the post mock o k response
func (o *PostMockOK) WithPayload(payload *models.SLACKResponse) *PostMockOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post mock o k response
func (o *PostMockOK) SetPayload(payload *models.SLACKResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMockOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMockMethodNotAllowedCode is the HTTP code returned for type PostMockMethodNotAllowed
const PostMockMethodNotAllowedCode int = 405

/*PostMockMethodNotAllowed Invalid input

swagger:response postMockMethodNotAllowed
*/
type PostMockMethodNotAllowed struct {
}

// NewPostMockMethodNotAllowed creates PostMockMethodNotAllowed with default headers values
func NewPostMockMethodNotAllowed() *PostMockMethodNotAllowed {

	return &PostMockMethodNotAllowed{}
}

// WriteResponse to the client
func (o *PostMockMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
