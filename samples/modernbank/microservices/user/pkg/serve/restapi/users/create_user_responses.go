// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	model "github.com/tetrateio/training/samples/modernbank/microservices/user/pkg/model"
)

// CreateUserCreatedCode is the HTTP code returned for type CreateUserCreated
const CreateUserCreatedCode int = 201

/*CreateUserCreated Created

swagger:response createUserCreated
*/
type CreateUserCreated struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`

	/*
	  In: Body
	*/
	Payload *model.User `json:"body,omitempty"`
}

// NewCreateUserCreated creates CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {

	return &CreateUserCreated{}
}

// WithVersion adds the version to the create user created response
func (o *CreateUserCreated) WithVersion(version string) *CreateUserCreated {
	o.Version = version
	return o
}

// SetVersion sets the version to the create user created response
func (o *CreateUserCreated) SetVersion(version string) {
	o.Version = version
}

// WithPayload adds the payload to the create user created response
func (o *CreateUserCreated) WithPayload(payload *model.User) *CreateUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create user created response
func (o *CreateUserCreated) SetPayload(payload *model.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUserConflictCode is the HTTP code returned for type CreateUserConflict
const CreateUserConflictCode int = 409

/*CreateUserConflict User alreadys exists

swagger:response createUserConflict
*/
type CreateUserConflict struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`
}

// NewCreateUserConflict creates CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {

	return &CreateUserConflict{}
}

// WithVersion adds the version to the create user conflict response
func (o *CreateUserConflict) WithVersion(version string) *CreateUserConflict {
	o.Version = version
	return o
}

// SetVersion sets the version to the create user conflict response
func (o *CreateUserConflict) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *CreateUserConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// CreateUserInternalServerErrorCode is the HTTP code returned for type CreateUserInternalServerError
const CreateUserInternalServerErrorCode int = 500

/*CreateUserInternalServerError Internal server error

swagger:response createUserInternalServerError
*/
type CreateUserInternalServerError struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`
}

// NewCreateUserInternalServerError creates CreateUserInternalServerError with default headers values
func NewCreateUserInternalServerError() *CreateUserInternalServerError {

	return &CreateUserInternalServerError{}
}

// WithVersion adds the version to the create user internal server error response
func (o *CreateUserInternalServerError) WithVersion(version string) *CreateUserInternalServerError {
	o.Version = version
	return o
}

// SetVersion sets the version to the create user internal server error response
func (o *CreateUserInternalServerError) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *CreateUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
