// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"github.com/latitudesh/lsh/internal/output/table"
)

// Project project
//
// swagger:model project
type Project struct {

	// attributes
	Attributes *ProjectAttributes `json:"attributes,omitempty"`

	// The project ID
	ID string `json:"id,omitempty"`
}

// Validate validates this project
func (m *Project) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Project) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this project based on the context it is used
func (m *Project) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Project) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.Attributes != nil {

		if swag.IsZero(m.Attributes) { // not required
			return nil
		}

		if err := m.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Project) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Project) UnmarshalBinary(b []byte) error {
	var res Project
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func (m *Project) TableRow() table.Row {
	attr := m.Attributes

	return table.Row{
		"id": table.Cell{
			Label: "ID",
			Value: table.String(m.ID),
		},
		"tags": table.Cell{
			Label: "Tags",
			Value: table.StringList(tags(attr.Tags), ","),
		},
		"name": table.Cell{
			Label: "Name",
			Value: table.String(attr.Name),
		},
		"description": table.Cell{
			Label: "Description",
			Value: table.String(attr.Description),
		},
		"slug": table.Cell{
			Label: "Slug",
			Value: table.String(attr.Slug),
		},
		"ips": table.Cell{
			Label: "IPs",
			Value: table.Float(attr.Stats.IPAddresses),
		},
		"servers": table.Cell{
			Label: "Servers",
			Value: table.Float(attr.Stats.Servers),
		},
		"vlans": table.Cell{
			Label: "VLANs",
			Value: table.Float(attr.Stats.Vlans),
		},
		"team": table.Cell{
			Label: "Team",
			Value: table.String(attr.Team.Name),
		},
		"environment": table.Cell{
			Label: "Environment",
			Value: table.String(attr.Environment),
		},
		"provisioning_type": table.Cell{
			Label: "Provisioning Type",
			Value: table.String(attr.ProvisiongType),
		},
	}
}

// ProjectAttributes project attributes
//
// swagger:model ProjectAttributes
type ProjectAttributes struct {

	Tags []*TagsIncude `json:"tags,omitempty"`
	
	// billing
	Billing *ProjectAttributesBilling `json:"billing,omitempty"`

	// billing method
	// Enum: [Normal 95th percentile]
	BillingMethod string `json:"billing_method,omitempty"`

	// billing type
	// Enum: [Yearly Monthly Hourly Normal]
	BillingType string `json:"billing_type,omitempty"`

	// cost
	Cost string `json:"cost,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// The project description
	Description string `json:"description,omitempty"`

	// environment
	// Enum: [Development Staging Production]
	Environment string `json:"environment,omitempty"`

	// The project name
	Name string `json:"name,omitempty"`

	// A unique project identifier
	Slug string `json:"slug,omitempty"`

	ProvisiongType string `json:"provisioning_type,omitempty"`

	// stats
	Stats *ProjectAttributesStats `json:"stats,omitempty"`

	// team
	Team *TeamInclude `json:"team,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this project attributes
func (m *ProjectAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBilling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectAttributes) validateBilling(formats strfmt.Registry) error {
	if swag.IsZero(m.Billing) { // not required
		return nil
	}

	if m.Billing != nil {
		if err := m.Billing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "billing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "billing")
			}
			return err
		}
	}

	return nil
}

var projectAttributesTypeBillingMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Normal","95th percentile"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectAttributesTypeBillingMethodPropEnum = append(projectAttributesTypeBillingMethodPropEnum, v)
	}
}

const (

	// ProjectAttributesBillingMethodNormal captures enum value "Normal"
	ProjectAttributesBillingMethodNormal string = "Normal"

	// ProjectAttributesBillingMethodNr95thPercentile captures enum value "95th percentile"
	ProjectAttributesBillingMethodNr95thPercentile string = "95th percentile"
)

// prop value enum
func (m *ProjectAttributes) validateBillingMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, projectAttributesTypeBillingMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProjectAttributes) validateBillingMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingMethodEnum("attributes"+"."+"billing_method", "body", m.BillingMethod); err != nil {
		return err
	}

	return nil
}

var projectAttributesTypeBillingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Yearly","Monthly","Hourly","Normal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectAttributesTypeBillingTypePropEnum = append(projectAttributesTypeBillingTypePropEnum, v)
	}
}

const (

	// ProjectAttributesBillingTypeYearly captures enum value "Yearly"
	ProjectAttributesBillingTypeYearly string = "Yearly"

	// ProjectAttributesBillingTypeMonthly captures enum value "Monthly"
	ProjectAttributesBillingTypeMonthly string = "Monthly"

	// ProjectAttributesBillingTypeHourly captures enum value "Hourly"
	ProjectAttributesBillingTypeHourly string = "Hourly"

	// ProjectAttributesBillingTypeNormal captures enum value "Normal"
	ProjectAttributesBillingTypeNormal string = "Normal"
)

// prop value enum
func (m *ProjectAttributes) validateBillingTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, projectAttributesTypeBillingTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProjectAttributes) validateBillingType(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingTypeEnum("attributes"+"."+"billing_type", "body", m.BillingType); err != nil {
		return err
	}

	return nil
}

var projectAttributesTypeEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Development","Staging","Production"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectAttributesTypeEnvironmentPropEnum = append(projectAttributesTypeEnvironmentPropEnum, v)
	}
}

const (

	// ProjectAttributesEnvironmentDevelopment captures enum value "Development"
	ProjectAttributesEnvironmentDevelopment string = "Development"

	// ProjectAttributesEnvironmentStaging captures enum value "Staging"
	ProjectAttributesEnvironmentStaging string = "Staging"

	// ProjectAttributesEnvironmentProduction captures enum value "Production"
	ProjectAttributesEnvironmentProduction string = "Production"
)

// prop value enum
func (m *ProjectAttributes) validateEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, projectAttributesTypeEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProjectAttributes) validateEnvironment(formats strfmt.Registry) error {
	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnvironmentEnum("attributes"+"."+"environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *ProjectAttributes) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "stats")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectAttributes) validateTeam(formats strfmt.Registry) error {
	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "team")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this project attributes based on the context it is used
func (m *ProjectAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBilling(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectAttributes) contextValidateBilling(ctx context.Context, formats strfmt.Registry) error {

	if m.Billing != nil {

		if swag.IsZero(m.Billing) { // not required
			return nil
		}

		if err := m.Billing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "billing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "billing")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectAttributes) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {

		if swag.IsZero(m.Stats) { // not required
			return nil
		}

		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "stats")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectAttributes) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

	if m.Team != nil {

		if swag.IsZero(m.Team) { // not required
			return nil
		}

		if err := m.Team.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "team")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectAttributes) UnmarshalBinary(b []byte) error {
	var res ProjectAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ProjectAttributesBilling project attributes billing
//
// swagger:model ProjectAttributesBilling
type ProjectAttributesBilling struct {

	// method
	Method string `json:"method,omitempty"`

	// subscription id
	SubscriptionID string `json:"subscription_id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this project attributes billing
func (m *ProjectAttributesBilling) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project attributes billing based on context it is used
func (m *ProjectAttributesBilling) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectAttributesBilling) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectAttributesBilling) UnmarshalBinary(b []byte) error {
	var res ProjectAttributesBilling
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ProjectAttributesStats project attributes stats
//
// swagger:model ProjectAttributesStats
type ProjectAttributesStats struct {

	// The number of IP addresses assigned to the project
	IPAddresses float64 `json:"ip_addresses,omitempty"`

	// The IP address prefixes in the project
	Prefixes float64 `json:"prefixes,omitempty"`

	// The number of servers assigned to the project
	Servers float64 `json:"servers,omitempty"`

	// The number of VLANs assigned to the project
	Vlans float64 `json:"vlans,omitempty"`
}

// Validate validates this project attributes stats
func (m *ProjectAttributesStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project attributes stats based on context it is used
func (m *ProjectAttributesStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectAttributesStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectAttributesStats) UnmarshalBinary(b []byte) error {
	var res ProjectAttributesStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
