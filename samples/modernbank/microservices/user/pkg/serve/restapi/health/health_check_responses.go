// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// HealthCheckOKCode is the HTTP code returned for type HealthCheckOK
const HealthCheckOKCode int = 200

/*HealthCheckOK OK

swagger:response healthCheckOK
*/
type HealthCheckOK struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`
}

// NewHealthCheckOK creates HealthCheckOK with default headers values
func NewHealthCheckOK() *HealthCheckOK {

	return &HealthCheckOK{}
}

// WithVersion adds the version to the health check o k response
func (o *HealthCheckOK) WithVersion(version string) *HealthCheckOK {
	o.Version = version
	return o
}

// SetVersion sets the version to the health check o k response
func (o *HealthCheckOK) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *HealthCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
