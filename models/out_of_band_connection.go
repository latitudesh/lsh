// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OutOfBandConnection out of band connection
//
// swagger:model out_of_band_connection
type OutOfBandConnection struct {

	// data
	Data *OutOfBandConnectionData `json:"data,omitempty"`
}

// Validate validates this out of band connection
func (m *OutOfBandConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutOfBandConnection) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this out of band connection based on the context it is used
func (m *OutOfBandConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutOfBandConnection) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {

		if swag.IsZero(m.Data) { // not required
			return nil
		}

		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OutOfBandConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutOfBandConnection) UnmarshalBinary(b []byte) error {
	var res OutOfBandConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OutOfBandConnectionData out of band connection data
//
// swagger:model OutOfBandConnectionData
type OutOfBandConnectionData struct {

	// attributes
	Attributes *OutOfBandConnectionDataAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this out of band connection data
func (m *OutOfBandConnectionData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutOfBandConnectionData) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this out of band connection data based on the context it is used
func (m *OutOfBandConnectionData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutOfBandConnectionData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.Attributes != nil {

		if swag.IsZero(m.Attributes) { // not required
			return nil
		}

		if err := m.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OutOfBandConnectionData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutOfBandConnectionData) UnmarshalBinary(b []byte) error {
	var res OutOfBandConnectionData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OutOfBandConnectionDataAttributes out of band connection data attributes
//
// swagger:model OutOfBandConnectionDataAttributes
type OutOfBandConnectionDataAttributes struct {

	// access ip
	AccessIP string `json:"access_ip,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// credentials
	Credentials *OutOfBandConnectionDataAttributesCredentials `json:"credentials,omitempty"`

	// port
	Port string `json:"port,omitempty"`

	// server id
	ServerID string `json:"server_id,omitempty"`

	// ssh key
	SSHKey *OutOfBandConnectionDataAttributesSSHKey `json:"ssh_key,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this out of band connection data attributes
func (m *OutOfBandConnectionDataAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutOfBandConnectionDataAttributes) validateCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes" + "." + "credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes" + "." + "credentials")
			}
			return err
		}
	}

	return nil
}

func (m *OutOfBandConnectionDataAttributes) validateSSHKey(formats strfmt.Registry) error {
	if swag.IsZero(m.SSHKey) { // not required
		return nil
	}

	if m.SSHKey != nil {
		if err := m.SSHKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes" + "." + "ssh_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes" + "." + "ssh_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this out of band connection data attributes based on the context it is used
func (m *OutOfBandConnectionDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSSHKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutOfBandConnectionDataAttributes) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {

		if swag.IsZero(m.Credentials) { // not required
			return nil
		}

		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes" + "." + "credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes" + "." + "credentials")
			}
			return err
		}
	}

	return nil
}

func (m *OutOfBandConnectionDataAttributes) contextValidateSSHKey(ctx context.Context, formats strfmt.Registry) error {

	if m.SSHKey != nil {

		if swag.IsZero(m.SSHKey) { // not required
			return nil
		}

		if err := m.SSHKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "attributes" + "." + "ssh_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "attributes" + "." + "ssh_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OutOfBandConnectionDataAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutOfBandConnectionDataAttributes) UnmarshalBinary(b []byte) error {
	var res OutOfBandConnectionDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OutOfBandConnectionDataAttributesCredentials credentials are valid only when the server is deployed with ssh keys
//
// swagger:model OutOfBandConnectionDataAttributesCredentials
type OutOfBandConnectionDataAttributesCredentials struct {

	// password
	Password string `json:"password,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this out of band connection data attributes credentials
func (m *OutOfBandConnectionDataAttributesCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this out of band connection data attributes credentials based on context it is used
func (m *OutOfBandConnectionDataAttributesCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OutOfBandConnectionDataAttributesCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutOfBandConnectionDataAttributesCredentials) UnmarshalBinary(b []byte) error {
	var res OutOfBandConnectionDataAttributesCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OutOfBandConnectionDataAttributesSSHKey out of band connection data attributes SSH key
//
// swagger:model OutOfBandConnectionDataAttributesSSHKey
type OutOfBandConnectionDataAttributesSSHKey struct {

	// description
	Description string `json:"description,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this out of band connection data attributes SSH key
func (m *OutOfBandConnectionDataAttributesSSHKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this out of band connection data attributes SSH key based on context it is used
func (m *OutOfBandConnectionDataAttributesSSHKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OutOfBandConnectionDataAttributesSSHKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutOfBandConnectionDataAttributesSSHKey) UnmarshalBinary(b []byte) error {
	var res OutOfBandConnectionDataAttributesSSHKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
