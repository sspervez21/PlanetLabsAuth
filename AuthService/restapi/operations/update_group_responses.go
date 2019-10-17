// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "PlanetLabs/AuthService/models"
)

// UpdateGroupOKCode is the HTTP code returned for type UpdateGroupOK
const UpdateGroupOKCode int = 200

/*UpdateGroupOK OK

swagger:response updateGroupOK
*/
type UpdateGroupOK struct {
}

// NewUpdateGroupOK creates UpdateGroupOK with default headers values
func NewUpdateGroupOK() *UpdateGroupOK {

	return &UpdateGroupOK{}
}

// WriteResponse to the client
func (o *UpdateGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateGroupNotFoundCode is the HTTP code returned for type UpdateGroupNotFound
const UpdateGroupNotFoundCode int = 404

/*UpdateGroupNotFound Group does not exists

swagger:response updateGroupNotFound
*/
type UpdateGroupNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.BadRequest `json:"body,omitempty"`
}

// NewUpdateGroupNotFound creates UpdateGroupNotFound with default headers values
func NewUpdateGroupNotFound() *UpdateGroupNotFound {

	return &UpdateGroupNotFound{}
}

// WithPayload adds the payload to the update group not found response
func (o *UpdateGroupNotFound) WithPayload(payload *models.BadRequest) *UpdateGroupNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update group not found response
func (o *UpdateGroupNotFound) SetPayload(payload *models.BadRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGroupNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}