// Code generated by go-swagger; DO NOT EDIT.

package virtual_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/models"
)

// GetVirtualNetworkReader is a Reader for the GetVirtualNetwork structure.
type GetVirtualNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVirtualNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVirtualNetworkOK()
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
		return nil, runtime.NewAPIError("[GET /virtual_networks/{id}] get-virtual-network", response, response.Code())
	}
}

// NewGetVirtualNetworkOK creates a GetVirtualNetworkOK with default headers values
func NewGetVirtualNetworkOK() *GetVirtualNetworkOK {
	return &GetVirtualNetworkOK{}
}

/*
GetVirtualNetworkOK describes a response with status code 200, with default header values.

List private networks
*/
type GetVirtualNetworkOK struct {
	Payload *GetVirtualNetworkOKBody
}

// IsSuccess returns true when this get virtual network o k response has a 2xx status code
func (o *GetVirtualNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get virtual network o k response has a 3xx status code
func (o *GetVirtualNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get virtual network o k response has a 4xx status code
func (o *GetVirtualNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get virtual network o k response has a 5xx status code
func (o *GetVirtualNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get virtual network o k response a status code equal to that given
func (o *GetVirtualNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get virtual network o k response
func (o *GetVirtualNetworkOK) Code() int {
	return 200
}

func (o *GetVirtualNetworkOK) Error() string {
	return fmt.Sprintf("[GET /virtual_networks/{id}][%d] getVirtualNetworkOK  %+v", 200, o.Payload)
}

func (o *GetVirtualNetworkOK) String() string {
	return fmt.Sprintf("[GET /virtual_networks/{id}][%d] getVirtualNetworkOK  %+v", 200, o.Payload)
}

func (o *GetVirtualNetworkOK) GetPayload() *GetVirtualNetworkOKBody {
	return o.Payload
}

func (o *GetVirtualNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetVirtualNetworkOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetVirtualNetworkOKBody get virtual network o k body
swagger:model GetVirtualNetworkOKBody
*/
type GetVirtualNetworkOKBody struct {

	// data
	Data *models.VirtualNetwork `json:"data,omitempty"`
}

// Validate validates this get virtual network o k body
func (o *GetVirtualNetworkOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetVirtualNetworkOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getVirtualNetworkOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getVirtualNetworkOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get virtual network o k body based on the context it is used
func (o *GetVirtualNetworkOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetVirtualNetworkOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getVirtualNetworkOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getVirtualNetworkOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetVirtualNetworkOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVirtualNetworkOKBody) UnmarshalBinary(b []byte) error {
	var res GetVirtualNetworkOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
