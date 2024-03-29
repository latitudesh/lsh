package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"github.com/latitudesh/lsh/internal/output/table"
)

// APIKey api key
//
// swagger:model api_key
type APIKey struct {

	// attributes
	Attributes *APIKeyAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

type APIKeys struct {
	Data []*APIKey `json:"data"`
}

// Validate validates this api key
func (m *APIKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIKey) validateAttributes(formats strfmt.Registry) error {
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

// ContextValidate validate this api key based on the context it is used
func (m *APIKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIKey) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *APIKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

func (m *APIKey) TableRow() table.Row {
	attr := m.Attributes

	return table.Row{
		"id": table.Cell{
			Label: "ID",
			Value: table.String(m.ID),
		},
		"name": table.Cell{
			Label: "Name",
			Value: table.String(attr.Name),
		},
		"token_last_slice": table.Cell{
			Label: "Token Last Slice",
			Value: table.String(attr.TokenLastSlice),
		},
		"user": table.Cell{
			Label: "User",
			Value: table.String(attr.User.Email),
		},
		"api_version": table.Cell{
			Label: "API Version",
			Value: table.String(attr.APIVersion),
		},
		"last_used_at": table.Cell{
			Label: "Last Used At",
			Value: table.DateTime(attr.LastUsedAt),
		},
	}
}

// UnmarshalBinary interface implementation
func (m *APIKey) UnmarshalBinary(b []byte) error {
	var res APIKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// APIKeyAttributes API key attributes
//
// swagger:model APIKeyAttributes
type APIKeyAttributes struct {

	// The API version associated with this API Key
	APIVersion string `json:"api_version,omitempty"`

	// The time when the API Key was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The last time a request was made to the API using this API Key
	// Format: date-time
	LastUsedAt *strfmt.DateTime `json:"last_used_at,omitempty"`

	// Name of the API Key
	Name string `json:"name,omitempty"`

	// The last 5 characters of the token created for this API Key
	TokenLastSlice string `json:"token_last_slice,omitempty"`

	// The time when the API Key was updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// user
	User *APIKeyAttributesUser `json:"user,omitempty"`
}

// Validate validates this API key attributes
func (m *APIKeyAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUsedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIKeyAttributes) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIKeyAttributes) validateLastUsedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUsedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"last_used_at", "body", "date-time", m.LastUsedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIKeyAttributes) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIKeyAttributes) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this API key attributes based on the context it is used
func (m *APIKeyAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIKeyAttributes) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {

		if swag.IsZero(m.User) { // not required
			return nil
		}

		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIKeyAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIKeyAttributes) UnmarshalBinary(b []byte) error {
	var res APIKeyAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// APIKeyAttributesUser The owner of the API Key
//
// swagger:model APIKeyAttributesUser
type APIKeyAttributesUser struct {

	// email
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this API key attributes user
func (m *APIKeyAttributesUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this API key attributes user based on context it is used
func (m *APIKeyAttributesUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIKeyAttributesUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIKeyAttributesUser) UnmarshalBinary(b []byte) error {
	var res APIKeyAttributesUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
