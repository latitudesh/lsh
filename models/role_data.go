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

// RoleData role data
//
// swagger:model role_data
type RoleData struct {

	// attributes
	Attributes *RoleDataAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [roles]
	Type string `json:"type,omitempty"`
}

// Validate validates this role data
func (m *RoleData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
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

func (m *RoleData) validateAttributes(formats strfmt.Registry) error {
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

var roleDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["roles"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		roleDataTypeTypePropEnum = append(roleDataTypeTypePropEnum, v)
	}
}

const (

	// RoleDataTypeRoles captures enum value "roles"
	RoleDataTypeRoles string = "roles"
)

// prop value enum
func (m *RoleData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, roleDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RoleData) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this role data based on the context it is used
func (m *RoleData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RoleData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleData) UnmarshalBinary(b []byte) error {
	var res RoleData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RoleDataAttributes role data attributes
//
// swagger:model RoleDataAttributes
type RoleDataAttributes struct {

	// Name of the Role
	Name string `json:"name,omitempty"`
}

// Validate validates this role data attributes
func (m *RoleDataAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this role data attributes based on context it is used
func (m *RoleDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoleDataAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleDataAttributes) UnmarshalBinary(b []byte) error {
	var res RoleDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
