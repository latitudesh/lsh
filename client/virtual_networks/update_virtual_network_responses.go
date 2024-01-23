// Code generated by go-swagger; DO NOT EDIT.

package virtual_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/models"
)

// UpdateVirtualNetworkReader is a Reader for the UpdateVirtualNetwork structure.
type UpdateVirtualNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVirtualNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVirtualNetworkOK()
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
	default:
		return nil, runtime.NewAPIError("[PATCH /virtual_networks/{virtual_network_id}] update-virtual-network", response, response.Code())
	}
}

// NewUpdateVirtualNetworkOK creates a UpdateVirtualNetworkOK with default headers values
func NewUpdateVirtualNetworkOK() *UpdateVirtualNetworkOK {
	return &UpdateVirtualNetworkOK{}
}

/*
UpdateVirtualNetworkOK describes a response with status code 200, with default header values.

Updated
*/
type UpdateVirtualNetworkOK struct {
	Payload *models.VirtualNetwork
}

// IsSuccess returns true when this update virtual network o k response has a 2xx status code
func (o *UpdateVirtualNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update virtual network o k response has a 3xx status code
func (o *UpdateVirtualNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update virtual network o k response has a 4xx status code
func (o *UpdateVirtualNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update virtual network o k response has a 5xx status code
func (o *UpdateVirtualNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update virtual network o k response a status code equal to that given
func (o *UpdateVirtualNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update virtual network o k response
func (o *UpdateVirtualNetworkOK) Code() int {
	return 200
}

func (o *UpdateVirtualNetworkOK) Error() string {
	return fmt.Sprintf("[PATCH /virtual_networks/{virtual_network_id}][%d] updateVirtualNetworkOK  %+v", 200, o.Payload)
}

func (o *UpdateVirtualNetworkOK) String() string {
	return fmt.Sprintf("[PATCH /virtual_networks/{virtual_network_id}][%d] updateVirtualNetworkOK  %+v", 200, o.Payload)
}

func (o *UpdateVirtualNetworkOK) GetPayload() *models.VirtualNetwork {
	return o.Payload
}

func (o *UpdateVirtualNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVirtualNetworkForbidden creates a UpdateVirtualNetworkForbidden with default headers values
func NewUpdateVirtualNetworkForbidden() *UpdateVirtualNetworkForbidden {
	return &UpdateVirtualNetworkForbidden{}
}

/*
UpdateVirtualNetworkForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateVirtualNetworkForbidden struct {
	Payload *models.VirtualNetwork
}

// IsSuccess returns true when this update virtual network forbidden response has a 2xx status code
func (o *UpdateVirtualNetworkForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update virtual network forbidden response has a 3xx status code
func (o *UpdateVirtualNetworkForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update virtual network forbidden response has a 4xx status code
func (o *UpdateVirtualNetworkForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update virtual network forbidden response has a 5xx status code
func (o *UpdateVirtualNetworkForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update virtual network forbidden response a status code equal to that given
func (o *UpdateVirtualNetworkForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update virtual network forbidden response
func (o *UpdateVirtualNetworkForbidden) Code() int {
	return 403
}

func (o *UpdateVirtualNetworkForbidden) Error() string {
	return fmt.Sprintf("[PATCH /virtual_networks/{virtual_network_id}][%d] updateVirtualNetworkForbidden  %+v", 403, o.Payload)
}

func (o *UpdateVirtualNetworkForbidden) String() string {
	return fmt.Sprintf("[PATCH /virtual_networks/{virtual_network_id}][%d] updateVirtualNetworkForbidden  %+v", 403, o.Payload)
}

func (o *UpdateVirtualNetworkForbidden) GetPayload() *models.VirtualNetwork {
	return o.Payload
}

func (o *UpdateVirtualNetworkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateVirtualNetworkBody update virtual network body
swagger:model UpdateVirtualNetworkBody
*/
type UpdateVirtualNetworkBody struct {

	// data
	// Required: true
	Data *UpdateVirtualNetworkParamsBodyData `json:"data"`
}

// Validate validates this update virtual network body
func (o *UpdateVirtualNetworkBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateVirtualNetworkBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update virtual network body based on the context it is used
func (o *UpdateVirtualNetworkBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateVirtualNetworkBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVirtualNetworkBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVirtualNetworkBody) UnmarshalBinary(b []byte) error {
	var res UpdateVirtualNetworkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateVirtualNetworkParamsBodyData update virtual network params body data
swagger:model UpdateVirtualNetworkParamsBodyData
*/
type UpdateVirtualNetworkParamsBodyData struct {

	// attributes
	Attributes *UpdateVirtualNetworkParamsBodyDataAttributes `json:"attributes,omitempty"`

	// type
	// Required: true
	// Enum: [virtual_network]
	Type *string `json:"type"`
}

// Validate validates this update virtual network params body data
func (o *UpdateVirtualNetworkParamsBodyData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateVirtualNetworkParamsBodyData) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(o.Attributes) { // not required
		return nil
	}

	if o.Attributes != nil {
		if err := o.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data" + "." + "attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

var updateVirtualNetworkParamsBodyDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual_network"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateVirtualNetworkParamsBodyDataTypeTypePropEnum = append(updateVirtualNetworkParamsBodyDataTypeTypePropEnum, v)
	}
}

const (

	// UpdateVirtualNetworkParamsBodyDataTypeVirtualNetwork captures enum value "virtual_network"
	UpdateVirtualNetworkParamsBodyDataTypeVirtualNetwork string = "virtual_network"
)

// prop value enum
func (o *UpdateVirtualNetworkParamsBodyData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateVirtualNetworkParamsBodyDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateVirtualNetworkParamsBodyData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"data"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update virtual network params body data based on the context it is used
func (o *UpdateVirtualNetworkParamsBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateVirtualNetworkParamsBodyData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if o.Attributes != nil {

		if swag.IsZero(o.Attributes) { // not required
			return nil
		}

		if err := o.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "data" + "." + "attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "data" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVirtualNetworkParamsBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVirtualNetworkParamsBodyData) UnmarshalBinary(b []byte) error {
	var res UpdateVirtualNetworkParamsBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateVirtualNetworkParamsBodyDataAttributes update virtual network params body data attributes
swagger:model UpdateVirtualNetworkParamsBodyDataAttributes
*/
type UpdateVirtualNetworkParamsBodyDataAttributes struct {

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this update virtual network params body data attributes
func (o *UpdateVirtualNetworkParamsBodyDataAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update virtual network params body data attributes based on context it is used
func (o *UpdateVirtualNetworkParamsBodyDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVirtualNetworkParamsBodyDataAttributes) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVirtualNetworkParamsBodyDataAttributes) UnmarshalBinary(b []byte) error {
	var res UpdateVirtualNetworkParamsBodyDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
