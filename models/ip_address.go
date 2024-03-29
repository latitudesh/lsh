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
)

// IPAddress ip address
//
// swagger:model ip_address
type IPAddress struct {

	// attributes
	Attributes *IPAddressAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this ip address
func (m *IPAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAddress) validateAttributes(formats strfmt.Registry) error {
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

// ContextValidate validate this ip address based on the context it is used
func (m *IPAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAddress) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IPAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddress) UnmarshalBinary(b []byte) error {
	var res IPAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPAddressAttributes IP address attributes
//
// swagger:model IPAddressAttributes
type IPAddressAttributes struct {

	// address
	Address string `json:"address,omitempty"`

	// assignment
	Assignment *IPAddressAttributesAssignment `json:"assignment,omitempty"`

	// available
	Available bool `json:"available,omitempty"`

	// cidr
	Cidr *string `json:"cidr,omitempty"`

	// family
	// Enum: [IPv4 IPv6]
	Family string `json:"family,omitempty"`

	// gateway
	Gateway *string `json:"gateway,omitempty"`

	// management
	Management bool `json:"management,omitempty"`

	// netmask
	Netmask string `json:"netmask,omitempty"`

	// project
	Project *IPAddressAttributesProject `json:"project,omitempty"`

	// public
	Public bool `json:"public,omitempty"`

	// region
	Region *IPAddressAttributesRegion `json:"region,omitempty"`

	// type
	// Enum: [Public Private]
	Type string `json:"type,omitempty"`
}

// Validate validates this IP address attributes
func (m *IPAddressAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAddressAttributes) validateAssignment(formats strfmt.Registry) error {
	if swag.IsZero(m.Assignment) { // not required
		return nil
	}

	if m.Assignment != nil {
		if err := m.Assignment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "assignment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "assignment")
			}
			return err
		}
	}

	return nil
}

var ipAddressAttributesTypeFamilyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPv4","IPv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipAddressAttributesTypeFamilyPropEnum = append(ipAddressAttributesTypeFamilyPropEnum, v)
	}
}

const (

	// IPAddressAttributesFamilyIPV4 captures enum value "IPv4"
	IPAddressAttributesFamilyIPV4 string = "IPv4"

	// IPAddressAttributesFamilyIPV6 captures enum value "IPv6"
	IPAddressAttributesFamilyIPV6 string = "IPv6"
)

// prop value enum
func (m *IPAddressAttributes) validateFamilyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ipAddressAttributesTypeFamilyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IPAddressAttributes) validateFamily(formats strfmt.Registry) error {
	if swag.IsZero(m.Family) { // not required
		return nil
	}

	// value enum
	if err := m.validateFamilyEnum("attributes"+"."+"family", "body", m.Family); err != nil {
		return err
	}

	return nil
}

func (m *IPAddressAttributes) validateProject(formats strfmt.Registry) error {
	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "project")
			}
			return err
		}
	}

	return nil
}

func (m *IPAddressAttributes) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "region")
			}
			return err
		}
	}

	return nil
}

var ipAddressAttributesTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Public","Private"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipAddressAttributesTypeTypePropEnum = append(ipAddressAttributesTypeTypePropEnum, v)
	}
}

const (

	// IPAddressAttributesTypePublic captures enum value "Public"
	IPAddressAttributesTypePublic string = "Public"

	// IPAddressAttributesTypePrivate captures enum value "Private"
	IPAddressAttributesTypePrivate string = "Private"
)

// prop value enum
func (m *IPAddressAttributes) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ipAddressAttributesTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IPAddressAttributes) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("attributes"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this IP address attributes based on the context it is used
func (m *IPAddressAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAddressAttributes) contextValidateAssignment(ctx context.Context, formats strfmt.Registry) error {

	if m.Assignment != nil {

		if swag.IsZero(m.Assignment) { // not required
			return nil
		}

		if err := m.Assignment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "assignment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "assignment")
			}
			return err
		}
	}

	return nil
}

func (m *IPAddressAttributes) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	if m.Project != nil {

		if swag.IsZero(m.Project) { // not required
			return nil
		}

		if err := m.Project.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "project")
			}
			return err
		}
	}

	return nil
}

func (m *IPAddressAttributes) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {

		if swag.IsZero(m.Region) { // not required
			return nil
		}

		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAddressAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddressAttributes) UnmarshalBinary(b []byte) error {
	var res IPAddressAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPAddressAttributesAssignment IP address attributes assignment
//
// swagger:model IPAddressAttributesAssignment
type IPAddressAttributesAssignment struct {

	// assigned at
	AssignedAt string `json:"assigned_at,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// server id
	ServerID string `json:"server_id,omitempty"`
}

// Validate validates this IP address attributes assignment
func (m *IPAddressAttributesAssignment) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP address attributes assignment based on context it is used
func (m *IPAddressAttributesAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPAddressAttributesAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddressAttributesAssignment) UnmarshalBinary(b []byte) error {
	var res IPAddressAttributesAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPAddressAttributesProject IP address attributes project
//
// swagger:model IPAddressAttributesProject
type IPAddressAttributesProject struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this IP address attributes project
func (m *IPAddressAttributesProject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP address attributes project based on context it is used
func (m *IPAddressAttributesProject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPAddressAttributesProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddressAttributesProject) UnmarshalBinary(b []byte) error {
	var res IPAddressAttributesProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPAddressAttributesRegion IP address attributes region
//
// swagger:model IPAddressAttributesRegion
type IPAddressAttributesRegion struct {

	// id
	ID string `json:"id,omitempty"`

	// location
	Location *IPAddressAttributesRegionLocation `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this IP address attributes region
func (m *IPAddressAttributesRegion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAddressAttributesRegion) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "region" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "region" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this IP address attributes region based on the context it is used
func (m *IPAddressAttributesRegion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAddressAttributesRegion) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {

		if swag.IsZero(m.Location) { // not required
			return nil
		}

		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "region" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "region" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAddressAttributesRegion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddressAttributesRegion) UnmarshalBinary(b []byte) error {
	var res IPAddressAttributesRegion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPAddressAttributesRegionLocation IP address attributes region location
//
// swagger:model IPAddressAttributesRegionLocation
type IPAddressAttributesRegionLocation struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`
}

// Validate validates this IP address attributes region location
func (m *IPAddressAttributesRegionLocation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP address attributes region location based on context it is used
func (m *IPAddressAttributesRegionLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPAddressAttributesRegionLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAddressAttributesRegionLocation) UnmarshalBinary(b []byte) error {
	var res IPAddressAttributesRegionLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
