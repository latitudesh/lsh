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

// Events events
//
// swagger:model events
type Events struct {

	// attributes
	Attributes *EventsAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this events
func (m *Events) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Events) validateAttributes(formats strfmt.Registry) error {
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

// ContextValidate validate this events based on the context it is used
func (m *Events) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Events) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Events) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Events) UnmarshalBinary(b []byte) error {
	var res Events
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EventsAttributes events attributes
//
// swagger:model EventsAttributes
type EventsAttributes struct {

	// action
	Action string `json:"action,omitempty"`

	// author
	Author *EventsAttributesAuthor `json:"author,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// project
	Project *EventsAttributesProject `json:"project,omitempty"`

	// target
	Target *EventsAttributesTarget `json:"target,omitempty"`

	// team
	Team *EventsAttributesTeam `json:"team,omitempty"`
}

// Validate validates this events attributes
func (m *EventsAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
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

func (m *EventsAttributes) validateAuthor(formats strfmt.Registry) error {
	if swag.IsZero(m.Author) { // not required
		return nil
	}

	if m.Author != nil {
		if err := m.Author.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "author")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "author")
			}
			return err
		}
	}

	return nil
}

func (m *EventsAttributes) validateProject(formats strfmt.Registry) error {
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

func (m *EventsAttributes) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "target")
			}
			return err
		}
	}

	return nil
}

func (m *EventsAttributes) validateTeam(formats strfmt.Registry) error {
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

// ContextValidate validate this events attributes based on the context it is used
func (m *EventsAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
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

func (m *EventsAttributes) contextValidateAuthor(ctx context.Context, formats strfmt.Registry) error {

	if m.Author != nil {

		if swag.IsZero(m.Author) { // not required
			return nil
		}

		if err := m.Author.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "author")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "author")
			}
			return err
		}
	}

	return nil
}

func (m *EventsAttributes) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EventsAttributes) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "target")
			}
			return err
		}
	}

	return nil
}

func (m *EventsAttributes) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EventsAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventsAttributes) UnmarshalBinary(b []byte) error {
	var res EventsAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EventsAttributesAuthor events attributes author
//
// swagger:model EventsAttributesAuthor
type EventsAttributesAuthor struct {

	// email
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this events attributes author
func (m *EventsAttributesAuthor) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this events attributes author based on context it is used
func (m *EventsAttributesAuthor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventsAttributesAuthor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventsAttributesAuthor) UnmarshalBinary(b []byte) error {
	var res EventsAttributesAuthor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EventsAttributesProject events attributes project
//
// swagger:model EventsAttributesProject
type EventsAttributesProject struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`
}

// Validate validates this events attributes project
func (m *EventsAttributesProject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this events attributes project based on context it is used
func (m *EventsAttributesProject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventsAttributesProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventsAttributesProject) UnmarshalBinary(b []byte) error {
	var res EventsAttributesProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EventsAttributesTarget events attributes target
//
// swagger:model EventsAttributesTarget
type EventsAttributesTarget struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this events attributes target
func (m *EventsAttributesTarget) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this events attributes target based on context it is used
func (m *EventsAttributesTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventsAttributesTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventsAttributesTarget) UnmarshalBinary(b []byte) error {
	var res EventsAttributesTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EventsAttributesTeam events attributes team
//
// swagger:model EventsAttributesTeam
type EventsAttributesTeam struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this events attributes team
func (m *EventsAttributesTeam) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this events attributes team based on context it is used
func (m *EventsAttributesTeam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventsAttributesTeam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventsAttributesTeam) UnmarshalBinary(b []byte) error {
	var res EventsAttributesTeam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
