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

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/renderer"
	"github.com/latitudesh/lsh/models"
)

// PutProjectSSHKeyReader is a Reader for the PutProjectSSHKey structure.
type PutProjectSSHKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProjectSSHKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutProjectSSHKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := apierrors.NewBadRequest()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
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
	case 422:
		result := apierrors.NewUnprocessableEntity()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}] put-project-ssh-key", response, response.Code())
	}
}

// NewPutProjectSSHKeyOK creates a PutProjectSSHKeyOK with default headers values
func NewPutProjectSSHKeyOK() *PutProjectSSHKeyOK {
	return &PutProjectSSHKeyOK{}
}

/*
PutProjectSSHKeyOK describes a response with status code 200, with default header values.

Success
*/
type PutProjectSSHKeyOK struct {
	Payload *PutProjectSSHKeyOKBody
}

// IsSuccess returns true when this put project Ssh key o k response has a 2xx status code
func (o *PutProjectSSHKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put project Ssh key o k response has a 3xx status code
func (o *PutProjectSSHKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put project Ssh key o k response has a 4xx status code
func (o *PutProjectSSHKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put project Ssh key o k response has a 5xx status code
func (o *PutProjectSSHKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put project Ssh key o k response a status code equal to that given
func (o *PutProjectSSHKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put project Ssh key o k response
func (o *PutProjectSSHKeyOK) Code() int {
	return 200
}

func (o *PutProjectSSHKeyOK) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}][%d] putProjectSshKeyOK  %+v", 200, o.Payload)
}

func (o *PutProjectSSHKeyOK) String() string {
	return fmt.Sprintf("[PATCH /projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}][%d] putProjectSshKeyOK  %+v", 200, o.Payload)
}

func (o *PutProjectSSHKeyOK) GetPayload() *PutProjectSSHKeyOKBody {
	return o.Payload
}

func (o *PutProjectSSHKeyOK) GetData() []renderer.ResponseData {
	return []renderer.ResponseData{o.Payload.Data}
}

func (o *PutProjectSSHKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutProjectSSHKeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutProjectSSHKeyBody put project SSH key body
swagger:model PutProjectSSHKeyBody
*/
type PutProjectSSHKeyBody struct {

	// data
	// Required: true
	Data *PutProjectSSHKeyParamsBodyData `json:"data"`
}

// Validate validates this put project SSH key body
func (o *PutProjectSSHKeyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutProjectSSHKeyBody) validateData(formats strfmt.Registry) error {

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

// ContextValidate validate this put project SSH key body based on the context it is used
func (o *PutProjectSSHKeyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutProjectSSHKeyBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (o *PutProjectSSHKeyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutProjectSSHKeyBody) UnmarshalBinary(b []byte) error {
	var res PutProjectSSHKeyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutProjectSSHKeyOKBody put project SSH key o k body
swagger:model PutProjectSSHKeyOKBody
*/
type PutProjectSSHKeyOKBody struct {

	// data
	Data *models.SSHKeyData `json:"data,omitempty"`
}

// Validate validates this put project SSH key o k body
func (o *PutProjectSSHKeyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutProjectSSHKeyOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putProjectSshKeyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putProjectSshKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put project SSH key o k body based on the context it is used
func (o *PutProjectSSHKeyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutProjectSSHKeyOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putProjectSshKeyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putProjectSshKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutProjectSSHKeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutProjectSSHKeyOKBody) UnmarshalBinary(b []byte) error {
	var res PutProjectSSHKeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutProjectSSHKeyParamsBodyData put project SSH key params body data
swagger:model PutProjectSSHKeyParamsBodyData
*/
type PutProjectSSHKeyParamsBodyData struct {

	// attributes
	Attributes *PutProjectSSHKeyParamsBodyDataAttributes `json:"attributes,omitempty"`

	// id
	// Required: true
	ID string `json:"id"`

	// type
	// Required: true
	// Enum: [ssh_keys]
	Type *string `json:"type"`
}

// Validate validates this put project SSH key params body data
func (o *PutProjectSSHKeyParamsBodyData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
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

func (o *PutProjectSSHKeyParamsBodyData) validateAttributes(formats strfmt.Registry) error {
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

func (o *PutProjectSSHKeyParamsBodyData) validateID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

var putProjectSshKeyParamsBodyDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ssh_keys"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		putProjectSshKeyParamsBodyDataTypeTypePropEnum = append(putProjectSshKeyParamsBodyDataTypeTypePropEnum, v)
	}
}

const (

	// PutProjectSSHKeyParamsBodyDataTypeSSHKeys captures enum value "ssh_keys"
	PutProjectSSHKeyParamsBodyDataTypeSSHKeys string = "ssh_keys"
)

// prop value enum
func (o *PutProjectSSHKeyParamsBodyData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, putProjectSshKeyParamsBodyDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PutProjectSSHKeyParamsBodyData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"data"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this put project SSH key params body data based on the context it is used
func (o *PutProjectSSHKeyParamsBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutProjectSSHKeyParamsBodyData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (o *PutProjectSSHKeyParamsBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutProjectSSHKeyParamsBodyData) UnmarshalBinary(b []byte) error {
	var res PutProjectSSHKeyParamsBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutProjectSSHKeyParamsBodyDataAttributes put project SSH key params body data attributes
swagger:model PutProjectSSHKeyParamsBodyDataAttributes
*/
type PutProjectSSHKeyParamsBodyDataAttributes struct {

	// Name of the SSH Key
	Name string `json:"name,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

// Validate validates this put project SSH key params body data attributes
func (o *PutProjectSSHKeyParamsBodyDataAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put project SSH key params body data attributes based on context it is used
func (o *PutProjectSSHKeyParamsBodyDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutProjectSSHKeyParamsBodyDataAttributes) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutProjectSSHKeyParamsBodyDataAttributes) UnmarshalBinary(b []byte) error {
	var res PutProjectSSHKeyParamsBodyDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
