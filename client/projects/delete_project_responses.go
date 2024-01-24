// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/latitudesh/lsh/internal/api"
)

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := api.NewUnauthorized()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := api.NewForbidden()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := api.NewNotFound()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := api.NewNotAcceptable()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /projects/{id_or_slug}] delete-project", response, response.Code())
	}
}

// NewDeleteProjectNoContent creates a DeleteProjectNoContent with default headers values
func NewDeleteProjectNoContent() *DeleteProjectNoContent {
	return &DeleteProjectNoContent{}
}

/*
DeleteProjectNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteProjectNoContent struct {
}

// IsSuccess returns true when this delete project no content response has a 2xx status code
func (o *DeleteProjectNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project no content response has a 3xx status code
func (o *DeleteProjectNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project no content response has a 4xx status code
func (o *DeleteProjectNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project no content response has a 5xx status code
func (o *DeleteProjectNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project no content response a status code equal to that given
func (o *DeleteProjectNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete project no content response
func (o *DeleteProjectNoContent) Code() int {
	return 204
}

func (o *DeleteProjectNoContent) Error() string {
	return fmt.Sprintf("[DELETE /projects/{id_or_slug}][%d] deleteProjectNoContent ", 204)
}

func (o *DeleteProjectNoContent) String() string {
	return fmt.Sprintf("[DELETE /projects/{id_or_slug}][%d] deleteProjectNoContent ", 204)
}

func (o *DeleteProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
