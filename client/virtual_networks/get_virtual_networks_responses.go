// Code generated by go-swagger; DO NOT EDIT.

package virtual_networks

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

// GetVirtualNetworksReader is a Reader for the GetVirtualNetworks structure.
type GetVirtualNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVirtualNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVirtualNetworksOK()
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
	default:
		return nil, runtime.NewAPIError("[GET /virtual_networks] get-virtual-networks", response, response.Code())
	}
}

// NewGetVirtualNetworksOK creates a GetVirtualNetworksOK with default headers values
func NewGetVirtualNetworksOK() *GetVirtualNetworksOK {
	return &GetVirtualNetworksOK{}
}

/*
GetVirtualNetworksOK describes a response with status code 200, with default header values.

List virtual networks
*/
type GetVirtualNetworksOK struct {
	Payload *models.VirtualNetworks
}

// IsSuccess returns true when this get virtual networks o k response has a 2xx status code
func (o *GetVirtualNetworksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get virtual networks o k response has a 3xx status code
func (o *GetVirtualNetworksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get virtual networks o k response has a 4xx status code
func (o *GetVirtualNetworksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get virtual networks o k response has a 5xx status code
func (o *GetVirtualNetworksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get virtual networks o k response a status code equal to that given
func (o *GetVirtualNetworksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get virtual networks o k response
func (o *GetVirtualNetworksOK) Code() int {
	return 200
}

func (o *GetVirtualNetworksOK) Error() string {
	return fmt.Sprintf("[GET /virtual_networks][%d] getVirtualNetworksOK  %+v", 200, o.Payload)
}

func (o *GetVirtualNetworksOK) String() string {
	return fmt.Sprintf("[GET /virtual_networks][%d] getVirtualNetworksOK  %+v", 200, o.Payload)
}

func (o *GetVirtualNetworksOK) GetPayload() *models.VirtualNetworks {
	return o.Payload
}

func (o *GetVirtualNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualNetworks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
