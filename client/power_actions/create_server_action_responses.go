// Code generated by go-swagger; DO NOT EDIT.

package power_actions

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

	"github.com/latitudesh/lsh/models"
)

// CreateServerActionReader is a Reader for the CreateServerAction structure.
type CreateServerActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServerActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServerActionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateServerActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateServerActionNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /servers/{server_id}/actions] create-server-action", response, response.Code())
	}
}

// NewCreateServerActionCreated creates a CreateServerActionCreated with default headers values
func NewCreateServerActionCreated() *CreateServerActionCreated {
	return &CreateServerActionCreated{}
}

/*
CreateServerActionCreated describes a response with status code 201, with default header values.

Created
*/
type CreateServerActionCreated struct {
	Payload *models.ServerAction
}

// IsSuccess returns true when this create server action created response has a 2xx status code
func (o *CreateServerActionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create server action created response has a 3xx status code
func (o *CreateServerActionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create server action created response has a 4xx status code
func (o *CreateServerActionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create server action created response has a 5xx status code
func (o *CreateServerActionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create server action created response a status code equal to that given
func (o *CreateServerActionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create server action created response
func (o *CreateServerActionCreated) Code() int {
	return 201
}

func (o *CreateServerActionCreated) Error() string {
	return fmt.Sprintf("[POST /servers/{server_id}/actions][%d] createServerActionCreated  %+v", 201, o.Payload)
}

func (o *CreateServerActionCreated) String() string {
	return fmt.Sprintf("[POST /servers/{server_id}/actions][%d] createServerActionCreated  %+v", 201, o.Payload)
}

func (o *CreateServerActionCreated) GetPayload() *models.ServerAction {
	return o.Payload
}

func (o *CreateServerActionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerAction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerActionForbidden creates a CreateServerActionForbidden with default headers values
func NewCreateServerActionForbidden() *CreateServerActionForbidden {
	return &CreateServerActionForbidden{}
}

/*
CreateServerActionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateServerActionForbidden struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this create server action forbidden response has a 2xx status code
func (o *CreateServerActionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create server action forbidden response has a 3xx status code
func (o *CreateServerActionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create server action forbidden response has a 4xx status code
func (o *CreateServerActionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create server action forbidden response has a 5xx status code
func (o *CreateServerActionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create server action forbidden response a status code equal to that given
func (o *CreateServerActionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create server action forbidden response
func (o *CreateServerActionForbidden) Code() int {
	return 403
}

func (o *CreateServerActionForbidden) Error() string {
	return fmt.Sprintf("[POST /servers/{server_id}/actions][%d] createServerActionForbidden  %+v", 403, o.Payload)
}

func (o *CreateServerActionForbidden) String() string {
	return fmt.Sprintf("[POST /servers/{server_id}/actions][%d] createServerActionForbidden  %+v", 403, o.Payload)
}

func (o *CreateServerActionForbidden) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *CreateServerActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServerActionNotAcceptable creates a CreateServerActionNotAcceptable with default headers values
func NewCreateServerActionNotAcceptable() *CreateServerActionNotAcceptable {
	return &CreateServerActionNotAcceptable{}
}

/*
CreateServerActionNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type CreateServerActionNotAcceptable struct {
	Payload *models.ErrorObject
}

// IsSuccess returns true when this create server action not acceptable response has a 2xx status code
func (o *CreateServerActionNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create server action not acceptable response has a 3xx status code
func (o *CreateServerActionNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create server action not acceptable response has a 4xx status code
func (o *CreateServerActionNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this create server action not acceptable response has a 5xx status code
func (o *CreateServerActionNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this create server action not acceptable response a status code equal to that given
func (o *CreateServerActionNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the create server action not acceptable response
func (o *CreateServerActionNotAcceptable) Code() int {
	return 406
}

func (o *CreateServerActionNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /servers/{server_id}/actions][%d] createServerActionNotAcceptable  %+v", 406, o.Payload)
}

func (o *CreateServerActionNotAcceptable) String() string {
	return fmt.Sprintf("[POST /servers/{server_id}/actions][%d] createServerActionNotAcceptable  %+v", 406, o.Payload)
}

func (o *CreateServerActionNotAcceptable) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *CreateServerActionNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateServerActionBody create server action body
swagger:model CreateServerActionBody
*/
type CreateServerActionBody struct {

	// data
	// Required: true
	Data *CreateServerActionParamsBodyData `json:"data"`
}

// Validate validates this create server action body
func (o *CreateServerActionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateServerActionBody) validateData(formats strfmt.Registry) error {

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

// ContextValidate validate this create server action body based on the context it is used
func (o *CreateServerActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateServerActionBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (o *CreateServerActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateServerActionBody) UnmarshalBinary(b []byte) error {
	var res CreateServerActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateServerActionParamsBodyData create server action params body data
swagger:model CreateServerActionParamsBodyData
*/
type CreateServerActionParamsBodyData struct {

	// attributes
	Attributes *CreateServerActionParamsBodyDataAttributes `json:"attributes,omitempty"`

	// type
	// Required: true
	// Enum: [actions]
	Type *string `json:"type"`
}

// Validate validates this create server action params body data
func (o *CreateServerActionParamsBodyData) Validate(formats strfmt.Registry) error {
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

func (o *CreateServerActionParamsBodyData) validateAttributes(formats strfmt.Registry) error {
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

var createServerActionParamsBodyDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["actions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createServerActionParamsBodyDataTypeTypePropEnum = append(createServerActionParamsBodyDataTypeTypePropEnum, v)
	}
}

const (

	// CreateServerActionParamsBodyDataTypeActions captures enum value "actions"
	CreateServerActionParamsBodyDataTypeActions string = "actions"
)

// prop value enum
func (o *CreateServerActionParamsBodyData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createServerActionParamsBodyDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateServerActionParamsBodyData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"data"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create server action params body data based on the context it is used
func (o *CreateServerActionParamsBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateServerActionParamsBodyData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (o *CreateServerActionParamsBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateServerActionParamsBodyData) UnmarshalBinary(b []byte) error {
	var res CreateServerActionParamsBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateServerActionParamsBodyDataAttributes create server action params body data attributes
swagger:model CreateServerActionParamsBodyDataAttributes
*/
type CreateServerActionParamsBodyDataAttributes struct {

	// The action to perform on the server
	// Required: true
	// Enum: [power_on power_off reboot]
	Action *string `json:"action"`
}

// Validate validates this create server action params body data attributes
func (o *CreateServerActionParamsBodyDataAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createServerActionParamsBodyDataAttributesTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["power_on","power_off","reboot"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createServerActionParamsBodyDataAttributesTypeActionPropEnum = append(createServerActionParamsBodyDataAttributesTypeActionPropEnum, v)
	}
}

const (

	// CreateServerActionParamsBodyDataAttributesActionPowerOn captures enum value "power_on"
	CreateServerActionParamsBodyDataAttributesActionPowerOn string = "power_on"

	// CreateServerActionParamsBodyDataAttributesActionPowerOff captures enum value "power_off"
	CreateServerActionParamsBodyDataAttributesActionPowerOff string = "power_off"

	// CreateServerActionParamsBodyDataAttributesActionReboot captures enum value "reboot"
	CreateServerActionParamsBodyDataAttributesActionReboot string = "reboot"
)

// prop value enum
func (o *CreateServerActionParamsBodyDataAttributes) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createServerActionParamsBodyDataAttributesTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateServerActionParamsBodyDataAttributes) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"attributes"+"."+"action", "body", o.Action); err != nil {
		return err
	}

	// value enum
	if err := o.validateActionEnum("body"+"."+"data"+"."+"attributes"+"."+"action", "body", *o.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create server action params body data attributes based on context it is used
func (o *CreateServerActionParamsBodyDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateServerActionParamsBodyDataAttributes) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateServerActionParamsBodyDataAttributes) UnmarshalBinary(b []byte) error {
	var res CreateServerActionParamsBodyDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
