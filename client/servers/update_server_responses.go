package servers

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
	"github.com/spf13/viper"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/output"
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

// UpdateServerReader is a Reader for the UpdateServer structure.
type UpdateServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServerOK()
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
		return nil, runtime.NewAPIError("[PATCH /servers/{server_id}] update-server", response, response.Code())
	}
}

// NewUpdateServerOK creates a UpdateServerOK with default headers values
func NewUpdateServerOK() *UpdateServerOK {
	return &UpdateServerOK{}
}

/*
UpdateServerOK describes a response with status code 200, with default header values.

Updated
*/
type UpdateServerOK struct {
	Payload *models.Server
}

// IsSuccess returns true when this update server o k response has a 2xx status code
func (o *UpdateServerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update server o k response has a 3xx status code
func (o *UpdateServerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update server o k response has a 4xx status code
func (o *UpdateServerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update server o k response has a 5xx status code
func (o *UpdateServerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update server o k response a status code equal to that given
func (o *UpdateServerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update server o k response
func (o *UpdateServerOK) Code() int {
	return 200
}

func (o *UpdateServerOK) Error() string {
	return fmt.Sprintf("[PATCH /servers/{server_id}][%d] updateServerOK  %+v", 200, o.Payload)
}

func (o *UpdateServerOK) String() string {
	return fmt.Sprintf("[PATCH /servers/{server_id}][%d] updateServerOK  %+v", 200, o.Payload)
}

func (o *UpdateServerOK) GetPayload() *models.Server {
	return o.Payload
}

type UpdateServerTableRow struct {
	ID              string `json:"id,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	PrimaryIPV4     string `json:"primary_ipv4,omitempty"`
	Status          string `json:"status,omitempty"`
	IPMIStatus      string `json:"ipmi_status,omitempty"`
	Location        string `json:"location,omitempty"`
	Project         string `json:"project,omitempty"`
	Team            string `json:"team,omitempty"`
	Plan            string `json:"plan,omitempty"`
	OperatingSystem string `json:"operating_system,omitempty"`
	CPU             string `json:"cpu,omitempty"`
	Disk            string `json:"disk,omitempty"`
	RAM             string `json:"ram,omitempty"`
}

func (o *UpdateServerOK) Render() {
	formatAsJSON := viper.GetBool("json")

	if formatAsJSON {
		o.RenderJSON()
		return
	}

	formatOutputFlag := viper.GetString("output")

	switch formatOutputFlag {
	case "json":
		o.RenderJSON()
	case "table":
		o.RenderTable()
	default:
		fmt.Println("Unsupported output format")
	}
}

func (o *UpdateServerOK) RenderJSON() {
	if !swag.IsZero(o) && !swag.IsZero(o.Payload) {
		JSONString, err := json.Marshal(o.Payload)
		if err != nil {
			fmt.Println("Could not decode the result as JSON.")
		}

		output.RenderJSON(JSONString)
	}
}

func (o *UpdateServerOK) RenderTable() {
	resource := o.Payload.Data

	var rows []UpdateServerTableRow

	attributes := resource.Attributes

	row := UpdateServerTableRow{
		ID:              table.RenderString(resource.ID),
		Hostname:        table.RenderString(attributes.Hostname),
		PrimaryIPV4:     table.RenderString(*attributes.PrimaryIPV4),
		Status:          table.RenderString(attributes.Status),
		IPMIStatus:      table.RenderString(attributes.IpmiStatus),
		Location:        table.RenderString(attributes.Region.Site.Slug),
		Project:         table.RenderString(attributes.Project.Name),
		Team:            table.RenderString(attributes.Team.Name),
		Plan:            table.RenderString(attributes.Plan.Name),
		OperatingSystem: table.RenderString(attributes.OperatingSystem.Slug),
		CPU:             table.RenderString(attributes.Specs.CPU),
		Disk:            table.RenderString(attributes.Specs.Disk),
		RAM:             table.RenderString(attributes.Specs.RAM),
	}
	rows = append(rows, row)

	headers := table.ExtractHeaders(rows[0])

	var values [][]string

	for _, row := range rows {
		var tr []string

		for _, key := range headers {
			value, err := table.GetFieldValue(row, key)
			if err != nil {
				fmt.Printf("Error accessing field %s: %v\n", key, err)
				continue
			}

			tr = append(tr, fmt.Sprintf("%v", value))
		}

		values = append(values, tr)
	}

	table.Render(table.Table{Headers: headers, Rows: values})
}

func (o *UpdateServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateServerBody update server body
swagger:model UpdateServerBody
*/
type UpdateServerBody struct {

	// attributes
	Attributes *UpdateServerParamsBodyAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [servers]
	Type string `json:"type,omitempty"`
}

// Validate validates this update server body
func (o *UpdateServerBody) Validate(formats strfmt.Registry) error {
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

func (o *UpdateServerBody) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(o.Attributes) { // not required
		return nil
	}

	if o.Attributes != nil {
		if err := o.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

var updateServerBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["servers"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateServerBodyTypeTypePropEnum = append(updateServerBodyTypeTypePropEnum, v)
	}
}

const (

	// UpdateServerBodyTypeServers captures enum value "servers"
	UpdateServerBodyTypeServers string = "servers"
)

// prop value enum
func (o *UpdateServerBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateServerBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateServerBody) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update server body based on the context it is used
func (o *UpdateServerBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateServerBody) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if o.Attributes != nil {

		if swag.IsZero(o.Attributes) { // not required
			return nil
		}

		if err := o.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateServerBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateServerBody) UnmarshalBinary(b []byte) error {
	var res UpdateServerBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateServerParamsBodyAttributes update server params body attributes
swagger:model UpdateServerParamsBodyAttributes
*/
type UpdateServerParamsBodyAttributes struct {

	// billing
	Billing string `json:"billing,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`
}

// Validate validates this update server params body attributes
func (o *UpdateServerParamsBodyAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update server params body attributes based on context it is used
func (o *UpdateServerParamsBodyAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateServerParamsBodyAttributes) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateServerParamsBodyAttributes) UnmarshalBinary(b []byte) error {
	var res UpdateServerParamsBodyAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
