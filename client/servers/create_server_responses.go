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

// CreateServerReader is a Reader for the CreateServer structure.
type CreateServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServerCreated()
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
	case 422:
		result := apierrors.NewUnprocessableEntity()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /servers] create-server", response, response.Code())
	}
}

// NewCreateServerCreated creates a CreateServerCreated with default headers values
func NewCreateServerCreated() *CreateServerCreated {
	return &CreateServerCreated{}
}

/*
CreateServerCreated describes a response with status code 201, with default header values.

Created
*/
type CreateServerCreated struct {
	Payload *models.Server
}

// IsSuccess returns true when this create server created response has a 2xx status code
func (o *CreateServerCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create server created response has a 3xx status code
func (o *CreateServerCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create server created response has a 4xx status code
func (o *CreateServerCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create server created response has a 5xx status code
func (o *CreateServerCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create server created response a status code equal to that given
func (o *CreateServerCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create server created response
func (o *CreateServerCreated) Code() int {
	return 201
}

func (o *CreateServerCreated) Error() string {
	return fmt.Sprintf("[POST /servers][%d] createServerCreated  %+v", 201, o.Payload)
}

func (o *CreateServerCreated) String() string {
	return fmt.Sprintf("[POST /servers][%d] createServerCreated  %+v", 201, o.Payload)
}

func (o *CreateServerCreated) GetPayload() *models.Server {
	return o.Payload
}

type CreateServerTableRow struct {
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

func (o *CreateServerCreated) Render() {
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

func (o *CreateServerCreated) RenderJSON() {
	if !swag.IsZero(o) && !swag.IsZero(o.Payload) {
		JSONString, err := json.Marshal(o.Payload)
		if err != nil {
			fmt.Println("Could not decode the result as JSON.")
		}

		output.RenderJSON(JSONString)
	}
}

func (o *CreateServerCreated) RenderTable() {
	resource := o.Payload.Data

	var rows []CreateServerTableRow

	attributes := resource.Attributes

	row := CreateServerTableRow{
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

func (o *CreateServerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateServerBody create server body
swagger:model CreateServerBody
*/
type CreateServerBody struct {

	// data
	Data *CreateServerParamsBodyData `json:"data,omitempty"`
}

// Validate validates this create server body
func (o *CreateServerBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateServerBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
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

// ContextValidate validate this create server body based on the context it is used
func (o *CreateServerBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateServerBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

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
func (o *CreateServerBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateServerBody) UnmarshalBinary(b []byte) error {
	var res CreateServerBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateServerParamsBodyData create server params body data
swagger:model CreateServerParamsBodyData
*/
type CreateServerParamsBodyData struct {

	// attributes
	Attributes *CreateServerParamsBodyDataAttributes `json:"attributes,omitempty"`

	// type
	// Required: true
	// Enum: [servers]
	Type *string `json:"type"`
}

// Validate validates this create server params body data
func (o *CreateServerParamsBodyData) Validate(formats strfmt.Registry) error {
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

func (o *CreateServerParamsBodyData) validateAttributes(formats strfmt.Registry) error {
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

var createServerParamsBodyDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["servers"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createServerParamsBodyDataTypeTypePropEnum = append(createServerParamsBodyDataTypeTypePropEnum, v)
	}
}

const (

	// CreateServerParamsBodyDataTypeServers captures enum value "servers"
	CreateServerParamsBodyDataTypeServers string = "servers"
)

// prop value enum
func (o *CreateServerParamsBodyData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createServerParamsBodyDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateServerParamsBodyData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"data"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("body"+"."+"data"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create server params body data based on the context it is used
func (o *CreateServerParamsBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateServerParamsBodyData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (o *CreateServerParamsBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateServerParamsBodyData) UnmarshalBinary(b []byte) error {
	var res CreateServerParamsBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateServerParamsBodyDataAttributes create server params body data attributes
swagger:model CreateServerParamsBodyDataAttributes
*/
type CreateServerParamsBodyDataAttributes struct {

	// The server billing type. Accepts `hourly` and `monthly` for on demand projects and `yearly` for reserved projects.
	// Enum: [hourly monthly yearly]
	Billing *string `json:"billing,omitempty"`

	// The server hostname
	Hostname string `json:"hostname,omitempty"`

	// URL where iPXE script is stored on, necessary for custom image deployments.This attribute is required when iPXE is selected as operating system.
	IpxeURL *string `json:"ipxe_url,omitempty"`

	// The operating system for the new server
	// Enum: [ipxe windows_server_2019_std_v1 ubuntu_22_04_x64_lts debian_11 rockylinux_8 debian_10 rhel8 centos_7_4_x64 centos_8_x64 ubuntu_20_04_x64_lts debian_12 ubuntu22_ml_in_a_box windows2022]
	OperatingSystem string `json:"operating_system,omitempty"`

	// The plan to choose server from
	// Enum: [c2-large-x86 c2-medium-x86 c2-small-x86 c3-large-x86 c3-medium-x86 c3-small-x86 c3-xlarge-x86 g3-large-x86 g3-medium-x86 g3-small-x86 g3-xlarge-x86 m3-large-x86 s2-small-x86 s3-large-x86]
	Plan string `json:"plan,omitempty"`

	// The project (ID or Slug) to deploy the server
	Project string `json:"project,omitempty"`

	// RAID mode for the server
	// Enum: [raid-0 raid-1]
	Raid *string `json:"raid,omitempty"`

	// The site to deploy the server
	// Enum: [ASH BGT BUE CHI DAL FRA LAX LON MEX MEX2 MIA MIA2 NYC SAN SAN2 SAO SAO2 SYD TYO TYO2]
	Site string `json:"site,omitempty"`

	// SSH Keys to set on the server
	SSHKeys []string `json:"ssh_keys,omitempty"`

	// User data to set on the server
	UserData *int64 `json:"user_data,omitempty"`
}

// Validate validates this create server params body data attributes
func (o *CreateServerParamsBodyDataAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBilling(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOperatingSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRaid(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createServerParamsBodyDataAttributesTypeBillingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hourly","monthly","yearly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createServerParamsBodyDataAttributesTypeBillingPropEnum = append(createServerParamsBodyDataAttributesTypeBillingPropEnum, v)
	}
}

const (

	// CreateServerParamsBodyDataAttributesBillingHourly captures enum value "hourly"
	CreateServerParamsBodyDataAttributesBillingHourly string = "hourly"

	// CreateServerParamsBodyDataAttributesBillingMonthly captures enum value "monthly"
	CreateServerParamsBodyDataAttributesBillingMonthly string = "monthly"

	// CreateServerParamsBodyDataAttributesBillingYearly captures enum value "yearly"
	CreateServerParamsBodyDataAttributesBillingYearly string = "yearly"
)

// prop value enum
func (o *CreateServerParamsBodyDataAttributes) validateBillingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createServerParamsBodyDataAttributesTypeBillingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateServerParamsBodyDataAttributes) validateBilling(formats strfmt.Registry) error {
	if swag.IsZero(o.Billing) { // not required
		return nil
	}

	// value enum
	if err := o.validateBillingEnum("body"+"."+"data"+"."+"attributes"+"."+"billing", "body", *o.Billing); err != nil {
		return err
	}

	return nil
}

var createServerParamsBodyDataAttributesTypeOperatingSystemPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipxe","windows_server_2019_std_v1","ubuntu_22_04_x64_lts","debian_11","rockylinux_8","debian_10","rhel8","centos_7_4_x64","centos_8_x64","ubuntu_20_04_x64_lts","debian_12","ubuntu22_ml_in_a_box","windows2022"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createServerParamsBodyDataAttributesTypeOperatingSystemPropEnum = append(createServerParamsBodyDataAttributesTypeOperatingSystemPropEnum, v)
	}
}

const (

	// CreateServerParamsBodyDataAttributesOperatingSystemIpxe captures enum value "ipxe"
	CreateServerParamsBodyDataAttributesOperatingSystemIpxe string = "ipxe"

	// CreateServerParamsBodyDataAttributesOperatingSystemWindowsServer2019StdV1 captures enum value "windows_server_2019_std_v1"
	CreateServerParamsBodyDataAttributesOperatingSystemWindowsServer2019StdV1 string = "windows_server_2019_std_v1"

	// CreateServerParamsBodyDataAttributesOperatingSystemUbuntu2204X64Lts captures enum value "ubuntu_22_04_x64_lts"
	CreateServerParamsBodyDataAttributesOperatingSystemUbuntu2204X64Lts string = "ubuntu_22_04_x64_lts"

	// CreateServerParamsBodyDataAttributesOperatingSystemDebian11 captures enum value "debian_11"
	CreateServerParamsBodyDataAttributesOperatingSystemDebian11 string = "debian_11"

	// CreateServerParamsBodyDataAttributesOperatingSystemRockylinux8 captures enum value "rockylinux_8"
	CreateServerParamsBodyDataAttributesOperatingSystemRockylinux8 string = "rockylinux_8"

	// CreateServerParamsBodyDataAttributesOperatingSystemDebian10 captures enum value "debian_10"
	CreateServerParamsBodyDataAttributesOperatingSystemDebian10 string = "debian_10"

	// CreateServerParamsBodyDataAttributesOperatingSystemRhel8 captures enum value "rhel8"
	CreateServerParamsBodyDataAttributesOperatingSystemRhel8 string = "rhel8"

	// CreateServerParamsBodyDataAttributesOperatingSystemCentos74X64 captures enum value "centos_7_4_x64"
	CreateServerParamsBodyDataAttributesOperatingSystemCentos74X64 string = "centos_7_4_x64"

	// CreateServerParamsBodyDataAttributesOperatingSystemCentos8X64 captures enum value "centos_8_x64"
	CreateServerParamsBodyDataAttributesOperatingSystemCentos8X64 string = "centos_8_x64"

	// CreateServerParamsBodyDataAttributesOperatingSystemUbuntu2004X64Lts captures enum value "ubuntu_20_04_x64_lts"
	CreateServerParamsBodyDataAttributesOperatingSystemUbuntu2004X64Lts string = "ubuntu_20_04_x64_lts"

	// CreateServerParamsBodyDataAttributesOperatingSystemDebian12 captures enum value "debian_12"
	CreateServerParamsBodyDataAttributesOperatingSystemDebian12 string = "debian_12"

	// CreateServerParamsBodyDataAttributesOperatingSystemUbuntu22MlInaBox captures enum value "ubuntu22_ml_in_a_box"
	CreateServerParamsBodyDataAttributesOperatingSystemUbuntu22MlInaBox string = "ubuntu22_ml_in_a_box"

	// CreateServerParamsBodyDataAttributesOperatingSystemWindows2022 captures enum value "windows2022"
	CreateServerParamsBodyDataAttributesOperatingSystemWindows2022 string = "windows2022"
)

// prop value enum
func (o *CreateServerParamsBodyDataAttributes) validateOperatingSystemEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createServerParamsBodyDataAttributesTypeOperatingSystemPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateServerParamsBodyDataAttributes) validateOperatingSystem(formats strfmt.Registry) error {
	if swag.IsZero(o.OperatingSystem) { // not required
		return nil
	}

	// value enum
	if err := o.validateOperatingSystemEnum("body"+"."+"data"+"."+"attributes"+"."+"operating_system", "body", o.OperatingSystem); err != nil {
		return err
	}

	return nil
}

var createServerParamsBodyDataAttributesTypePlanPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["c2-large-x86","c2-medium-x86","c2-small-x86","c3-large-x86","c3-medium-x86","c3-small-x86","c3-xlarge-x86","g3-large-x86","g3-medium-x86","g3-small-x86","g3-xlarge-x86","m3-large-x86","s2-small-x86","s3-large-x86"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createServerParamsBodyDataAttributesTypePlanPropEnum = append(createServerParamsBodyDataAttributesTypePlanPropEnum, v)
	}
}

const (

	// CreateServerParamsBodyDataAttributesPlanC2DashLargeDashX86 captures enum value "c2-large-x86"
	CreateServerParamsBodyDataAttributesPlanC2DashLargeDashX86 string = "c2-large-x86"

	// CreateServerParamsBodyDataAttributesPlanC2DashMediumDashX86 captures enum value "c2-medium-x86"
	CreateServerParamsBodyDataAttributesPlanC2DashMediumDashX86 string = "c2-medium-x86"

	// CreateServerParamsBodyDataAttributesPlanC2DashSmallDashX86 captures enum value "c2-small-x86"
	CreateServerParamsBodyDataAttributesPlanC2DashSmallDashX86 string = "c2-small-x86"

	// CreateServerParamsBodyDataAttributesPlanC3DashLargeDashX86 captures enum value "c3-large-x86"
	CreateServerParamsBodyDataAttributesPlanC3DashLargeDashX86 string = "c3-large-x86"

	// CreateServerParamsBodyDataAttributesPlanC3DashMediumDashX86 captures enum value "c3-medium-x86"
	CreateServerParamsBodyDataAttributesPlanC3DashMediumDashX86 string = "c3-medium-x86"

	// CreateServerParamsBodyDataAttributesPlanC3DashSmallDashX86 captures enum value "c3-small-x86"
	CreateServerParamsBodyDataAttributesPlanC3DashSmallDashX86 string = "c3-small-x86"

	// CreateServerParamsBodyDataAttributesPlanC3DashXlargeDashX86 captures enum value "c3-xlarge-x86"
	CreateServerParamsBodyDataAttributesPlanC3DashXlargeDashX86 string = "c3-xlarge-x86"

	// CreateServerParamsBodyDataAttributesPlanG3DashLargeDashX86 captures enum value "g3-large-x86"
	CreateServerParamsBodyDataAttributesPlanG3DashLargeDashX86 string = "g3-large-x86"

	// CreateServerParamsBodyDataAttributesPlanG3DashMediumDashX86 captures enum value "g3-medium-x86"
	CreateServerParamsBodyDataAttributesPlanG3DashMediumDashX86 string = "g3-medium-x86"

	// CreateServerParamsBodyDataAttributesPlanG3DashSmallDashX86 captures enum value "g3-small-x86"
	CreateServerParamsBodyDataAttributesPlanG3DashSmallDashX86 string = "g3-small-x86"

	// CreateServerParamsBodyDataAttributesPlanG3DashXlargeDashX86 captures enum value "g3-xlarge-x86"
	CreateServerParamsBodyDataAttributesPlanG3DashXlargeDashX86 string = "g3-xlarge-x86"

	// CreateServerParamsBodyDataAttributesPlanM3DashLargeDashX86 captures enum value "m3-large-x86"
	CreateServerParamsBodyDataAttributesPlanM3DashLargeDashX86 string = "m3-large-x86"

	// CreateServerParamsBodyDataAttributesPlanS2DashSmallDashX86 captures enum value "s2-small-x86"
	CreateServerParamsBodyDataAttributesPlanS2DashSmallDashX86 string = "s2-small-x86"

	// CreateServerParamsBodyDataAttributesPlanS3DashLargeDashX86 captures enum value "s3-large-x86"
	CreateServerParamsBodyDataAttributesPlanS3DashLargeDashX86 string = "s3-large-x86"
)

// prop value enum
func (o *CreateServerParamsBodyDataAttributes) validatePlanEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createServerParamsBodyDataAttributesTypePlanPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateServerParamsBodyDataAttributes) validatePlan(formats strfmt.Registry) error {
	if swag.IsZero(o.Plan) { // not required
		return nil
	}

	// value enum
	if err := o.validatePlanEnum("body"+"."+"data"+"."+"attributes"+"."+"plan", "body", o.Plan); err != nil {
		return err
	}

	return nil
}

var createServerParamsBodyDataAttributesTypeRaidPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["raid-0","raid-1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createServerParamsBodyDataAttributesTypeRaidPropEnum = append(createServerParamsBodyDataAttributesTypeRaidPropEnum, v)
	}
}

const (

	// CreateServerParamsBodyDataAttributesRaidRaidDash0 captures enum value "raid-0"
	CreateServerParamsBodyDataAttributesRaidRaidDash0 string = "raid-0"

	// CreateServerParamsBodyDataAttributesRaidRaidDash1 captures enum value "raid-1"
	CreateServerParamsBodyDataAttributesRaidRaidDash1 string = "raid-1"
)

// prop value enum
func (o *CreateServerParamsBodyDataAttributes) validateRaidEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createServerParamsBodyDataAttributesTypeRaidPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateServerParamsBodyDataAttributes) validateRaid(formats strfmt.Registry) error {
	if swag.IsZero(o.Raid) { // not required
		return nil
	}

	// value enum
	if err := o.validateRaidEnum("body"+"."+"data"+"."+"attributes"+"."+"raid", "body", *o.Raid); err != nil {
		return err
	}

	return nil
}

var createServerParamsBodyDataAttributesTypeSitePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ASH","BGT","BUE","CHI","DAL","FRA","LAX","LON","MEX","MEX2","MIA","MIA2","NYC","SAN","SAN2","SAO","SAO2","SYD","TYO","TYO2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createServerParamsBodyDataAttributesTypeSitePropEnum = append(createServerParamsBodyDataAttributesTypeSitePropEnum, v)
	}
}

const (

	// CreateServerParamsBodyDataAttributesSiteASH captures enum value "ASH"
	CreateServerParamsBodyDataAttributesSiteASH string = "ASH"

	// CreateServerParamsBodyDataAttributesSiteBGT captures enum value "BGT"
	CreateServerParamsBodyDataAttributesSiteBGT string = "BGT"

	// CreateServerParamsBodyDataAttributesSiteBUE captures enum value "BUE"
	CreateServerParamsBodyDataAttributesSiteBUE string = "BUE"

	// CreateServerParamsBodyDataAttributesSiteCHI captures enum value "CHI"
	CreateServerParamsBodyDataAttributesSiteCHI string = "CHI"

	// CreateServerParamsBodyDataAttributesSiteDAL captures enum value "DAL"
	CreateServerParamsBodyDataAttributesSiteDAL string = "DAL"

	// CreateServerParamsBodyDataAttributesSiteFRA captures enum value "FRA"
	CreateServerParamsBodyDataAttributesSiteFRA string = "FRA"

	// CreateServerParamsBodyDataAttributesSiteLAX captures enum value "LAX"
	CreateServerParamsBodyDataAttributesSiteLAX string = "LAX"

	// CreateServerParamsBodyDataAttributesSiteLON captures enum value "LON"
	CreateServerParamsBodyDataAttributesSiteLON string = "LON"

	// CreateServerParamsBodyDataAttributesSiteMEX captures enum value "MEX"
	CreateServerParamsBodyDataAttributesSiteMEX string = "MEX"

	// CreateServerParamsBodyDataAttributesSiteMEX2 captures enum value "MEX2"
	CreateServerParamsBodyDataAttributesSiteMEX2 string = "MEX2"

	// CreateServerParamsBodyDataAttributesSiteMIA captures enum value "MIA"
	CreateServerParamsBodyDataAttributesSiteMIA string = "MIA"

	// CreateServerParamsBodyDataAttributesSiteMIA2 captures enum value "MIA2"
	CreateServerParamsBodyDataAttributesSiteMIA2 string = "MIA2"

	// CreateServerParamsBodyDataAttributesSiteNYC captures enum value "NYC"
	CreateServerParamsBodyDataAttributesSiteNYC string = "NYC"

	// CreateServerParamsBodyDataAttributesSiteSAN captures enum value "SAN"
	CreateServerParamsBodyDataAttributesSiteSAN string = "SAN"

	// CreateServerParamsBodyDataAttributesSiteSAN2 captures enum value "SAN2"
	CreateServerParamsBodyDataAttributesSiteSAN2 string = "SAN2"

	// CreateServerParamsBodyDataAttributesSiteSAO captures enum value "SAO"
	CreateServerParamsBodyDataAttributesSiteSAO string = "SAO"

	// CreateServerParamsBodyDataAttributesSiteSAO2 captures enum value "SAO2"
	CreateServerParamsBodyDataAttributesSiteSAO2 string = "SAO2"

	// CreateServerParamsBodyDataAttributesSiteSYD captures enum value "SYD"
	CreateServerParamsBodyDataAttributesSiteSYD string = "SYD"

	// CreateServerParamsBodyDataAttributesSiteTYO captures enum value "TYO"
	CreateServerParamsBodyDataAttributesSiteTYO string = "TYO"

	// CreateServerParamsBodyDataAttributesSiteTYO2 captures enum value "TYO2"
	CreateServerParamsBodyDataAttributesSiteTYO2 string = "TYO2"
)

// prop value enum
func (o *CreateServerParamsBodyDataAttributes) validateSiteEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createServerParamsBodyDataAttributesTypeSitePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateServerParamsBodyDataAttributes) validateSite(formats strfmt.Registry) error {
	if swag.IsZero(o.Site) { // not required
		return nil
	}

	// value enum
	if err := o.validateSiteEnum("body"+"."+"data"+"."+"attributes"+"."+"site", "body", o.Site); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create server params body data attributes based on context it is used
func (o *CreateServerParamsBodyDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateServerParamsBodyDataAttributes) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateServerParamsBodyDataAttributes) UnmarshalBinary(b []byte) error {
	var res CreateServerParamsBodyDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
