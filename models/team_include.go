// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TeamInclude team include
//
// swagger:model team_include
type TeamInclude struct {

	// address
	Address string `json:"address,omitempty"`

	// currency
	Currency interface{} `json:"currency,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this team include
func (m *TeamInclude) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this team include based on context it is used
func (m *TeamInclude) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TeamInclude) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamInclude) UnmarshalBinary(b []byte) error {
	var res TeamInclude
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
