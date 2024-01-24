// Code generated by go-swagger; DO NOT EDIT.

package plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/models"
)

// GetPlansReader is a Reader for the GetPlans structure.
type GetPlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlansOK()
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
		return nil, runtime.NewAPIError("[GET /plans] get-plans", response, response.Code())
	}
}

// NewGetPlansOK creates a GetPlansOK with default headers values
func NewGetPlansOK() *GetPlansOK {
	return &GetPlansOK{}
}

/*
GetPlansOK describes a response with status code 200, with default header values.

Success
*/
type GetPlansOK struct {
	Payload *GetPlansOKBody
}

// IsSuccess returns true when this get plans o k response has a 2xx status code
func (o *GetPlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get plans o k response has a 3xx status code
func (o *GetPlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get plans o k response has a 4xx status code
func (o *GetPlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get plans o k response has a 5xx status code
func (o *GetPlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get plans o k response a status code equal to that given
func (o *GetPlansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get plans o k response
func (o *GetPlansOK) Code() int {
	return 200
}

func (o *GetPlansOK) Error() string {
	return fmt.Sprintf("[GET /plans][%d] getPlansOK  %+v", 200, o.Payload)
}

func (o *GetPlansOK) String() string {
	return fmt.Sprintf("[GET /plans][%d] getPlansOK  %+v", 200, o.Payload)
}

func (o *GetPlansOK) GetPayload() *GetPlansOKBody {
	return o.Payload
}

func (o *GetPlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPlansOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetPlansOKBody get plans o k body
swagger:model GetPlansOKBody
*/
type GetPlansOKBody struct {

	// data
	Data []*models.PlanData `json:"data"`
}

// Validate validates this get plans o k body
func (o *GetPlansOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPlansOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPlansOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getPlansOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get plans o k body based on the context it is used
func (o *GetPlansOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPlansOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPlansOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getPlansOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPlansOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPlansOKBody) UnmarshalBinary(b []byte) error {
	var res GetPlansOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
