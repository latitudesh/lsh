// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

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

// PostProjectSSHKeyReader is a Reader for the PostProjectSSHKey structure.
type PostProjectSSHKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProjectSSHKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostProjectSSHKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostProjectSSHKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := api.NewNotFound()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /projects/{project_id_or_slug}/ssh_keys] post-project-ssh-key", response, response.Code())
	}
}

// NewPostProjectSSHKeyCreated creates a PostProjectSSHKeyCreated with default headers values
func NewPostProjectSSHKeyCreated() *PostProjectSSHKeyCreated {
	return &PostProjectSSHKeyCreated{}
}

/*
PostProjectSSHKeyCreated describes a response with status code 201, with default header values.

Created
*/
type PostProjectSSHKeyCreated struct {
	Payload *PostProjectSSHKeyCreatedBody
}

// IsSuccess returns true when this post project Ssh key created response has a 2xx status code
func (o *PostProjectSSHKeyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post project Ssh key created response has a 3xx status code
func (o *PostProjectSSHKeyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post project Ssh key created response has a 4xx status code
func (o *PostProjectSSHKeyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post project Ssh key created response has a 5xx status code
func (o *PostProjectSSHKeyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post project Ssh key created response a status code equal to that given
func (o *PostProjectSSHKeyCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post project Ssh key created response
func (o *PostProjectSSHKeyCreated) Code() int {
	return 201
}

func (o *PostProjectSSHKeyCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id_or_slug}/ssh_keys][%d] postProjectSshKeyCreated  %+v", 201, o.Payload)
}

func (o *PostProjectSSHKeyCreated) String() string {
	return fmt.Sprintf("[POST /projects/{project_id_or_slug}/ssh_keys][%d] postProjectSshKeyCreated  %+v", 201, o.Payload)
}

func (o *PostProjectSSHKeyCreated) GetPayload() *PostProjectSSHKeyCreatedBody {
	return o.Payload
}

func (o *PostProjectSSHKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostProjectSSHKeyCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProjectSSHKeyBadRequest creates a PostProjectSSHKeyBadRequest with default headers values
func NewPostProjectSSHKeyBadRequest() *PostProjectSSHKeyBadRequest {
	return &PostProjectSSHKeyBadRequest{}
}

/*
PostProjectSSHKeyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostProjectSSHKeyBadRequest struct {
}

// IsSuccess returns true when this post project Ssh key bad request response has a 2xx status code
func (o *PostProjectSSHKeyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post project Ssh key bad request response has a 3xx status code
func (o *PostProjectSSHKeyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post project Ssh key bad request response has a 4xx status code
func (o *PostProjectSSHKeyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post project Ssh key bad request response has a 5xx status code
func (o *PostProjectSSHKeyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post project Ssh key bad request response a status code equal to that given
func (o *PostProjectSSHKeyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post project Ssh key bad request response
func (o *PostProjectSSHKeyBadRequest) Code() int {
	return 400
}

func (o *PostProjectSSHKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id_or_slug}/ssh_keys][%d] postProjectSshKeyBadRequest ", 400)
}

func (o *PostProjectSSHKeyBadRequest) String() string {
	return fmt.Sprintf("[POST /projects/{project_id_or_slug}/ssh_keys][%d] postProjectSshKeyBadRequest ", 400)
}

func (o *PostProjectSSHKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectSSHKeyUnprocessableEntity creates a PostProjectSSHKeyUnprocessableEntity with default headers values
func NewPostProjectSSHKeyUnprocessableEntity() *PostProjectSSHKeyUnprocessableEntity {
	return &PostProjectSSHKeyUnprocessableEntity{}
}

/*
PostProjectSSHKeyUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PostProjectSSHKeyUnprocessableEntity struct {
}

// IsSuccess returns true when this post project Ssh key unprocessable entity response has a 2xx status code
func (o *PostProjectSSHKeyUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post project Ssh key unprocessable entity response has a 3xx status code
func (o *PostProjectSSHKeyUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post project Ssh key unprocessable entity response has a 4xx status code
func (o *PostProjectSSHKeyUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post project Ssh key unprocessable entity response has a 5xx status code
func (o *PostProjectSSHKeyUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post project Ssh key unprocessable entity response a status code equal to that given
func (o *PostProjectSSHKeyUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the post project Ssh key unprocessable entity response
func (o *PostProjectSSHKeyUnprocessableEntity) Code() int {
	return 422
}

func (o *PostProjectSSHKeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id_or_slug}/ssh_keys][%d] postProjectSshKeyUnprocessableEntity ", 422)
}

func (o *PostProjectSSHKeyUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /projects/{project_id_or_slug}/ssh_keys][%d] postProjectSshKeyUnprocessableEntity ", 422)
}

func (o *PostProjectSSHKeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
PostProjectSSHKeyBody post project SSH key body
swagger:model PostProjectSSHKeyBody
*/
type PostProjectSSHKeyBody struct {

	// data
	// Required: true
	Data *PostProjectSSHKeyParamsBodyData `json:"data"`
}

// Validate validates this post project SSH key body
func (o *PostProjectSSHKeyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostProjectSSHKeyBody) validateData(formats strfmt.Registry) error {

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

// ContextValidate validate this post project SSH key body based on the context it is used
func (o *PostProjectSSHKeyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostProjectSSHKeyBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (o *PostProjectSSHKeyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostProjectSSHKeyBody) UnmarshalBinary(b []byte) error {
	var res PostProjectSSHKeyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostProjectSSHKeyCreatedBody post project SSH key created body
swagger:model PostProjectSSHKeyCreatedBody
*/
type PostProjectSSHKeyCreatedBody struct {

	// data
	Data *models.SSHKeyData `json:"data,omitempty"`
}

// Validate validates this post project SSH key created body
func (o *PostProjectSSHKeyCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostProjectSSHKeyCreatedBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postProjectSshKeyCreated" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postProjectSshKeyCreated" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post project SSH key created body based on the context it is used
func (o *PostProjectSSHKeyCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostProjectSSHKeyCreatedBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postProjectSshKeyCreated" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postProjectSshKeyCreated" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostProjectSSHKeyCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostProjectSSHKeyCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostProjectSSHKeyCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostProjectSSHKeyParamsBodyData post project SSH key params body data
swagger:model PostProjectSSHKeyParamsBodyData
*/
type PostProjectSSHKeyParamsBodyData struct {

	// attributes
	Attributes *PostProjectSSHKeyParamsBodyDataAttributes `json:"attributes,omitempty"`

	// type
	// Required: true
	// Enum: [ssh_keys]
	Type *string `json:"type"`
}

// Validate validates this post project SSH key params body data
func (o *PostProjectSSHKeyParamsBodyData) Validate(formats strfmt.Registry) error {
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

func (o *PostProjectSSHKeyParamsBodyData) validateAttributes(formats strfmt.Registry) error {
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

var postProjectSshKeyParamsBodyDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ssh_keys"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postProjectSshKeyParamsBodyDataTypeTypePropEnum = append(postProjectSshKeyParamsBodyDataTypeTypePropEnum, v)
	}
}

const (

	// PostProjectSSHKeyParamsBodyDataTypeSSHKeys captures enum value "ssh_keys"
	PostProjectSSHKeyParamsBodyDataTypeSSHKeys string = "ssh_keys"
)

// prop value enum
func (o *PostProjectSSHKeyParamsBodyData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postProjectSshKeyParamsBodyDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostProjectSSHKeyParamsBodyData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"data"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post project SSH key params body data based on the context it is used
func (o *PostProjectSSHKeyParamsBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostProjectSSHKeyParamsBodyData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (o *PostProjectSSHKeyParamsBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostProjectSSHKeyParamsBodyData) UnmarshalBinary(b []byte) error {
	var res PostProjectSSHKeyParamsBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostProjectSSHKeyParamsBodyDataAttributes post project SSH key params body data attributes
swagger:model PostProjectSSHKeyParamsBodyDataAttributes
*/
type PostProjectSSHKeyParamsBodyDataAttributes struct {

	// Name of the SSH Key
	Name string `json:"name,omitempty"`

	// SSH Public Key
	PublicKey string `json:"public_key,omitempty"`
}

// Validate validates this post project SSH key params body data attributes
func (o *PostProjectSSHKeyParamsBodyDataAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post project SSH key params body data attributes based on context it is used
func (o *PostProjectSSHKeyParamsBodyDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostProjectSSHKeyParamsBodyDataAttributes) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostProjectSSHKeyParamsBodyDataAttributes) UnmarshalBinary(b []byte) error {
	var res PostProjectSSHKeyParamsBodyDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
