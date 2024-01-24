// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/models"
)

// GetServerReader is a Reader for the GetServer structure.
type GetServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := apierrors.NewUnauthorized()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := apierrors.NewNotFound()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /servers/{server_id}] get-server", response, response.Code())
	}
}

// NewGetServerOK creates a GetServerOK with default headers values
func NewGetServerOK() *GetServerOK {
	return &GetServerOK{}
}

/*
GetServerOK describes a response with status code 200, with default header values.

Success
*/
type GetServerOK struct {
	Payload *models.Server
}

// IsSuccess returns true when this get server o k response has a 2xx status code
func (o *GetServerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get server o k response has a 3xx status code
func (o *GetServerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get server o k response has a 4xx status code
func (o *GetServerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get server o k response has a 5xx status code
func (o *GetServerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get server o k response a status code equal to that given
func (o *GetServerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get server o k response
func (o *GetServerOK) Code() int {
	return 200
}

func (o *GetServerOK) Error() string {
	return fmt.Sprintf("[GET /servers/{server_id}][%d] getServerOK  %+v", 200, o.Payload)
}

func (o *GetServerOK) String() string {
	return fmt.Sprintf("[GET /servers/{server_id}][%d] getServerOK  %+v", 200, o.Payload)
}

func (o *GetServerOK) GetPayload() *models.Server {
	return o.Payload
}

func (o *GetServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
