// Code generated by go-swagger; DO NOT EDIT.

package api_keys

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

	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/models"
)

// UpdateAPIKeyReader is a Reader for the UpdateAPIKey structure.
type UpdateAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := api.NewBadRequest()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := api.NewUnauthorized()
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
		return nil, runtime.NewAPIError("[PUT /auth/api_keys/{id}] update-api-key", response, response.Code())
	}
}

// NewUpdateAPIKeyOK creates a UpdateAPIKeyOK with default headers values
func NewUpdateAPIKeyOK() *UpdateAPIKeyOK {
	return &UpdateAPIKeyOK{}
}

/*
UpdateAPIKeyOK describes a response with status code 200, with default header values.

API Key Updated
*/
type UpdateAPIKeyOK struct {
	Payload *UpdateAPIKeyOKBody
}

// IsSuccess returns true when this update Api key o k response has a 2xx status code
func (o *UpdateAPIKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update Api key o k response has a 3xx status code
func (o *UpdateAPIKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Api key o k response has a 4xx status code
func (o *UpdateAPIKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Api key o k response has a 5xx status code
func (o *UpdateAPIKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update Api key o k response a status code equal to that given
func (o *UpdateAPIKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update Api key o k response
func (o *UpdateAPIKeyOK) Code() int {
	return 200
}

func (o *UpdateAPIKeyOK) Error() string {
	return fmt.Sprintf("[PUT /auth/api_keys/{id}][%d] updateApiKeyOK  %+v", 200, o.Payload)
}

func (o *UpdateAPIKeyOK) String() string {
	return fmt.Sprintf("[PUT /auth/api_keys/{id}][%d] updateApiKeyOK  %+v", 200, o.Payload)
}

func (o *UpdateAPIKeyOK) GetPayload() *UpdateAPIKeyOKBody {
	return o.Payload
}

func (o *UpdateAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateAPIKeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAPIKeyBadRequest creates a UpdateAPIKeyBadRequest with default headers values
func NewUpdateAPIKeyBadRequest() *UpdateAPIKeyBadRequest {
	return &UpdateAPIKeyBadRequest{}
}

/*
UpdateAPIKeyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateAPIKeyBadRequest struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this update Api key bad request response has a 2xx status code
func (o *UpdateAPIKeyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Api key bad request response has a 3xx status code
func (o *UpdateAPIKeyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Api key bad request response has a 4xx status code
func (o *UpdateAPIKeyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Api key bad request response has a 5xx status code
func (o *UpdateAPIKeyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update Api key bad request response a status code equal to that given
func (o *UpdateAPIKeyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update Api key bad request response
func (o *UpdateAPIKeyBadRequest) Code() int {
	return 400
}

func (o *UpdateAPIKeyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /auth/api_keys/{id}][%d] updateApiKeyBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAPIKeyBadRequest) String() string {
	return fmt.Sprintf("[PUT /auth/api_keys/{id}][%d] updateApiKeyBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAPIKeyBadRequest) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *UpdateAPIKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAPIKeyNotFound creates a UpdateAPIKeyNotFound with default headers values
func NewUpdateAPIKeyNotFound() *UpdateAPIKeyNotFound {
	return &UpdateAPIKeyNotFound{}
}

/*
UpdateAPIKeyNotFound describes a response with status code 404, with default header values.

API Key Not Found
*/
type UpdateAPIKeyNotFound struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this update Api key not found response has a 2xx status code
func (o *UpdateAPIKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update Api key not found response has a 3xx status code
func (o *UpdateAPIKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Api key not found response has a 4xx status code
func (o *UpdateAPIKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update Api key not found response has a 5xx status code
func (o *UpdateAPIKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update Api key not found response a status code equal to that given
func (o *UpdateAPIKeyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update Api key not found response
func (o *UpdateAPIKeyNotFound) Code() int {
	return 404
}

func (o *UpdateAPIKeyNotFound) Error() string {
	return fmt.Sprintf("[PUT /auth/api_keys/{id}][%d] updateApiKeyNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAPIKeyNotFound) String() string {
	return fmt.Sprintf("[PUT /auth/api_keys/{id}][%d] updateApiKeyNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAPIKeyNotFound) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *UpdateAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateAPIKeyOKBody update API key o k body
swagger:model UpdateAPIKeyOKBody
*/
type UpdateAPIKeyOKBody struct {

	// data
	Data *models.APIKey `json:"data,omitempty"`
}

// Validate validates this update API key o k body
func (o *UpdateAPIKeyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateAPIKeyOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateApiKeyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateApiKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update API key o k body based on the context it is used
func (o *UpdateAPIKeyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateAPIKeyOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateApiKeyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateApiKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateAPIKeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateAPIKeyOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateAPIKeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
