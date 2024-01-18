// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlanData plan data
//
// swagger:model plan_data
type PlanData struct {

	// attributes
	Attributes *PlanDataAttributes `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [plans]
	Type string `json:"type,omitempty"`
}

// Validate validates this plan data
func (m *PlanData) Validate(formats strfmt.Registry) error {
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

func (m *PlanData) validateAttributes(formats strfmt.Registry) error {
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

var planDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["plans"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planDataTypeTypePropEnum = append(planDataTypeTypePropEnum, v)
	}
}

const (

	// PlanDataTypePlans captures enum value "plans"
	PlanDataTypePlans string = "plans"
)

// prop value enum
func (m *PlanData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, planDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlanData) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this plan data based on the context it is used
func (m *PlanData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanData) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PlanData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanData) UnmarshalBinary(b []byte) error {
	var res PlanData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributes plan data attributes
//
// swagger:model PlanDataAttributes
type PlanDataAttributes struct {

	// features
	Features []string `json:"features"`

	// regions
	Regions []*PlanDataAttributesRegionsItems0 `json:"regions"`

	// slug
	Slug string `json:"slug,omitempty"`

	// specs
	Specs *PlanDataAttributesSpecs `json:"specs,omitempty"`
}

// Validate validates this plan data attributes
func (m *PlanDataAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanDataAttributes) validateRegions(formats strfmt.Registry) error {
	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	for i := 0; i < len(m.Regions); i++ {
		if swag.IsZero(m.Regions[i]) { // not required
			continue
		}

		if m.Regions[i] != nil {
			if err := m.Regions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + "regions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + "regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlanDataAttributes) validateSpecs(formats strfmt.Registry) error {
	if swag.IsZero(m.Specs) { // not required
		return nil
	}

	if m.Specs != nil {
		if err := m.Specs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "specs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "specs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plan data attributes based on the context it is used
func (m *PlanDataAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanDataAttributes) contextValidateRegions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Regions); i++ {

		if m.Regions[i] != nil {

			if swag.IsZero(m.Regions[i]) { // not required
				return nil
			}

			if err := m.Regions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + "regions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + "regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlanDataAttributes) contextValidateSpecs(ctx context.Context, formats strfmt.Registry) error {

	if m.Specs != nil {

		if swag.IsZero(m.Specs) { // not required
			return nil
		}

		if err := m.Specs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "specs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "specs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributes) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributesRegionsItems0 plan data attributes regions items0
//
// swagger:model PlanDataAttributesRegionsItems0
type PlanDataAttributesRegionsItems0 struct {

	// deploys instantly
	DeploysInstantly []string `json:"deploys_instantly"`

	// locations
	Locations *PlanDataAttributesRegionsItems0Locations `json:"locations,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pricing
	Pricing *PlanDataAttributesRegionsItems0Pricing `json:"pricing,omitempty"`
}

// Validate validates this plan data attributes regions items0
func (m *PlanDataAttributesRegionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanDataAttributesRegionsItems0) validateLocations(formats strfmt.Registry) error {
	if swag.IsZero(m.Locations) { // not required
		return nil
	}

	if m.Locations != nil {
		if err := m.Locations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locations")
			}
			return err
		}
	}

	return nil
}

func (m *PlanDataAttributesRegionsItems0) validatePricing(formats strfmt.Registry) error {
	if swag.IsZero(m.Pricing) { // not required
		return nil
	}

	if m.Pricing != nil {
		if err := m.Pricing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pricing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pricing")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plan data attributes regions items0 based on the context it is used
func (m *PlanDataAttributesRegionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePricing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanDataAttributesRegionsItems0) contextValidateLocations(ctx context.Context, formats strfmt.Registry) error {

	if m.Locations != nil {

		if swag.IsZero(m.Locations) { // not required
			return nil
		}

		if err := m.Locations.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("locations")
			}
			return err
		}
	}

	return nil
}

func (m *PlanDataAttributesRegionsItems0) contextValidatePricing(ctx context.Context, formats strfmt.Registry) error {

	if m.Pricing != nil {

		if swag.IsZero(m.Pricing) { // not required
			return nil
		}

		if err := m.Pricing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pricing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pricing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributesRegionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributesRegionsItems0) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributesRegionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributesRegionsItems0Locations plan data attributes regions items0 locations
//
// swagger:model PlanDataAttributesRegionsItems0Locations
type PlanDataAttributesRegionsItems0Locations struct {

	// available
	Available []string `json:"available"`

	// in stock
	InStock []string `json:"in_stock"`
}

// Validate validates this plan data attributes regions items0 locations
func (m *PlanDataAttributesRegionsItems0Locations) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plan data attributes regions items0 locations based on context it is used
func (m *PlanDataAttributesRegionsItems0Locations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributesRegionsItems0Locations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributesRegionsItems0Locations) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributesRegionsItems0Locations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributesRegionsItems0Pricing plan data attributes regions items0 pricing
//
// swagger:model PlanDataAttributesRegionsItems0Pricing
type PlanDataAttributesRegionsItems0Pricing struct {

	// b r l
	BRL *PlanDataAttributesRegionsItems0PricingBRL `json:"BRL,omitempty"`

	// u s d
	USD *PlanDataAttributesRegionsItems0PricingUSD `json:"USD,omitempty"`
}

// Validate validates this plan data attributes regions items0 pricing
func (m *PlanDataAttributesRegionsItems0Pricing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBRL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUSD(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanDataAttributesRegionsItems0Pricing) validateBRL(formats strfmt.Registry) error {
	if swag.IsZero(m.BRL) { // not required
		return nil
	}

	if m.BRL != nil {
		if err := m.BRL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pricing" + "." + "BRL")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pricing" + "." + "BRL")
			}
			return err
		}
	}

	return nil
}

func (m *PlanDataAttributesRegionsItems0Pricing) validateUSD(formats strfmt.Registry) error {
	if swag.IsZero(m.USD) { // not required
		return nil
	}

	if m.USD != nil {
		if err := m.USD.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pricing" + "." + "USD")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pricing" + "." + "USD")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plan data attributes regions items0 pricing based on the context it is used
func (m *PlanDataAttributesRegionsItems0Pricing) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBRL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUSD(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanDataAttributesRegionsItems0Pricing) contextValidateBRL(ctx context.Context, formats strfmt.Registry) error {

	if m.BRL != nil {

		if swag.IsZero(m.BRL) { // not required
			return nil
		}

		if err := m.BRL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pricing" + "." + "BRL")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pricing" + "." + "BRL")
			}
			return err
		}
	}

	return nil
}

func (m *PlanDataAttributesRegionsItems0Pricing) contextValidateUSD(ctx context.Context, formats strfmt.Registry) error {

	if m.USD != nil {

		if swag.IsZero(m.USD) { // not required
			return nil
		}

		if err := m.USD.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pricing" + "." + "USD")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pricing" + "." + "USD")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributesRegionsItems0Pricing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributesRegionsItems0Pricing) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributesRegionsItems0Pricing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributesRegionsItems0PricingBRL plan data attributes regions items0 pricing b r l
//
// swagger:model PlanDataAttributesRegionsItems0PricingBRL
type PlanDataAttributesRegionsItems0PricingBRL struct {

	// hour
	Hour float64 `json:"hour,omitempty"`

	// month
	Month float64 `json:"month,omitempty"`

	// year
	Year float64 `json:"year,omitempty"`
}

// Validate validates this plan data attributes regions items0 pricing b r l
func (m *PlanDataAttributesRegionsItems0PricingBRL) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plan data attributes regions items0 pricing b r l based on context it is used
func (m *PlanDataAttributesRegionsItems0PricingBRL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributesRegionsItems0PricingBRL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributesRegionsItems0PricingBRL) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributesRegionsItems0PricingBRL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributesRegionsItems0PricingUSD plan data attributes regions items0 pricing u s d
//
// swagger:model PlanDataAttributesRegionsItems0PricingUSD
type PlanDataAttributesRegionsItems0PricingUSD struct {

	// hour
	Hour float64 `json:"hour,omitempty"`

	// month
	Month float64 `json:"month,omitempty"`

	// year
	Year float64 `json:"year,omitempty"`
}

// Validate validates this plan data attributes regions items0 pricing u s d
func (m *PlanDataAttributesRegionsItems0PricingUSD) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plan data attributes regions items0 pricing u s d based on context it is used
func (m *PlanDataAttributesRegionsItems0PricingUSD) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributesRegionsItems0PricingUSD) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributesRegionsItems0PricingUSD) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributesRegionsItems0PricingUSD
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributesSpecs plan data attributes specs
//
// swagger:model PlanDataAttributesSpecs
type PlanDataAttributesSpecs struct {

	// cpu
	CPU *PlanDataAttributesSpecsCPU `json:"cpu,omitempty"`

	// drives
	Drives []*PlanDataAttributesSpecsDrivesItems0 `json:"drives"`

	// gpu
	Gpu *PlanDataAttributesSpecsGpu `json:"gpu,omitempty"`

	// memory
	Memory *PlanDataAttributesSpecsMemory `json:"memory,omitempty"`

	// nics
	Nics []*PlanDataAttributesSpecsNicsItems0 `json:"nics"`
}

// Validate validates this plan data attributes specs
func (m *PlanDataAttributesSpecs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDrives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanDataAttributesSpecs) validateCPU(formats strfmt.Registry) error {
	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "specs" + "." + "cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "specs" + "." + "cpu")
			}
			return err
		}
	}

	return nil
}

func (m *PlanDataAttributesSpecs) validateDrives(formats strfmt.Registry) error {
	if swag.IsZero(m.Drives) { // not required
		return nil
	}

	for i := 0; i < len(m.Drives); i++ {
		if swag.IsZero(m.Drives[i]) { // not required
			continue
		}

		if m.Drives[i] != nil {
			if err := m.Drives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + "specs" + "." + "drives" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + "specs" + "." + "drives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlanDataAttributesSpecs) validateGpu(formats strfmt.Registry) error {
	if swag.IsZero(m.Gpu) { // not required
		return nil
	}

	if m.Gpu != nil {
		if err := m.Gpu.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "specs" + "." + "gpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "specs" + "." + "gpu")
			}
			return err
		}
	}

	return nil
}

func (m *PlanDataAttributesSpecs) validateMemory(formats strfmt.Registry) error {
	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if m.Memory != nil {
		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "specs" + "." + "memory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "specs" + "." + "memory")
			}
			return err
		}
	}

	return nil
}

func (m *PlanDataAttributesSpecs) validateNics(formats strfmt.Registry) error {
	if swag.IsZero(m.Nics) { // not required
		return nil
	}

	for i := 0; i < len(m.Nics); i++ {
		if swag.IsZero(m.Nics[i]) { // not required
			continue
		}

		if m.Nics[i] != nil {
			if err := m.Nics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + "specs" + "." + "nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + "specs" + "." + "nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this plan data attributes specs based on the context it is used
func (m *PlanDataAttributesSpecs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCPU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDrives(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGpu(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanDataAttributesSpecs) contextValidateCPU(ctx context.Context, formats strfmt.Registry) error {

	if m.CPU != nil {

		if swag.IsZero(m.CPU) { // not required
			return nil
		}

		if err := m.CPU.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "specs" + "." + "cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "specs" + "." + "cpu")
			}
			return err
		}
	}

	return nil
}

func (m *PlanDataAttributesSpecs) contextValidateDrives(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Drives); i++ {

		if m.Drives[i] != nil {

			if swag.IsZero(m.Drives[i]) { // not required
				return nil
			}

			if err := m.Drives[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + "specs" + "." + "drives" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + "specs" + "." + "drives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlanDataAttributesSpecs) contextValidateGpu(ctx context.Context, formats strfmt.Registry) error {

	if m.Gpu != nil {

		if swag.IsZero(m.Gpu) { // not required
			return nil
		}

		if err := m.Gpu.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "specs" + "." + "gpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "specs" + "." + "gpu")
			}
			return err
		}
	}

	return nil
}

func (m *PlanDataAttributesSpecs) contextValidateMemory(ctx context.Context, formats strfmt.Registry) error {

	if m.Memory != nil {

		if swag.IsZero(m.Memory) { // not required
			return nil
		}

		if err := m.Memory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "specs" + "." + "memory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "specs" + "." + "memory")
			}
			return err
		}
	}

	return nil
}

func (m *PlanDataAttributesSpecs) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nics); i++ {

		if m.Nics[i] != nil {

			if swag.IsZero(m.Nics[i]) { // not required
				return nil
			}

			if err := m.Nics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + "specs" + "." + "nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + "specs" + "." + "nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributesSpecs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributesSpecs) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributesSpecs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributesSpecsCPU plan data attributes specs CPU
//
// swagger:model PlanDataAttributesSpecsCPU
type PlanDataAttributesSpecsCPU struct {

	// clock
	Clock float64 `json:"clock,omitempty"`

	// cores
	Cores float64 `json:"cores,omitempty"`

	// count
	Count float64 `json:"count,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this plan data attributes specs CPU
func (m *PlanDataAttributesSpecsCPU) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plan data attributes specs CPU based on context it is used
func (m *PlanDataAttributesSpecsCPU) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributesSpecsCPU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributesSpecsCPU) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributesSpecsCPU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributesSpecsDrivesItems0 plan data attributes specs drives items0
//
// swagger:model PlanDataAttributesSpecsDrivesItems0
type PlanDataAttributesSpecsDrivesItems0 struct {

	// count
	Count float64 `json:"count,omitempty"`

	// size
	Size string `json:"size,omitempty"`

	// type
	// Enum: [SSD HDD]
	Type string `json:"type,omitempty"`
}

// Validate validates this plan data attributes specs drives items0
func (m *PlanDataAttributesSpecsDrivesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var planDataAttributesSpecsDrivesItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSD","HDD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planDataAttributesSpecsDrivesItems0TypeTypePropEnum = append(planDataAttributesSpecsDrivesItems0TypeTypePropEnum, v)
	}
}

const (

	// PlanDataAttributesSpecsDrivesItems0TypeSSD captures enum value "SSD"
	PlanDataAttributesSpecsDrivesItems0TypeSSD string = "SSD"

	// PlanDataAttributesSpecsDrivesItems0TypeHDD captures enum value "HDD"
	PlanDataAttributesSpecsDrivesItems0TypeHDD string = "HDD"
)

// prop value enum
func (m *PlanDataAttributesSpecsDrivesItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, planDataAttributesSpecsDrivesItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlanDataAttributesSpecsDrivesItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plan data attributes specs drives items0 based on context it is used
func (m *PlanDataAttributesSpecsDrivesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributesSpecsDrivesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributesSpecsDrivesItems0) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributesSpecsDrivesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributesSpecsGpu plan data attributes specs gpu
//
// swagger:model PlanDataAttributesSpecsGpu
type PlanDataAttributesSpecsGpu struct {

	// count
	Count float64 `json:"count,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this plan data attributes specs gpu
func (m *PlanDataAttributesSpecsGpu) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plan data attributes specs gpu based on context it is used
func (m *PlanDataAttributesSpecsGpu) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributesSpecsGpu) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributesSpecsGpu) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributesSpecsGpu
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributesSpecsMemory plan data attributes specs memory
//
// swagger:model PlanDataAttributesSpecsMemory
type PlanDataAttributesSpecsMemory struct {

	// total
	Total int `json:"total,omitempty"`
}

// Validate validates this plan data attributes specs memory
func (m *PlanDataAttributesSpecsMemory) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plan data attributes specs memory based on context it is used
func (m *PlanDataAttributesSpecsMemory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributesSpecsMemory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributesSpecsMemory) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributesSpecsMemory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlanDataAttributesSpecsNicsItems0 plan data attributes specs nics items0
//
// swagger:model PlanDataAttributesSpecsNicsItems0
type PlanDataAttributesSpecsNicsItems0 struct {

	// count
	Count float64 `json:"count,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this plan data attributes specs nics items0
func (m *PlanDataAttributesSpecsNicsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plan data attributes specs nics items0 based on context it is used
func (m *PlanDataAttributesSpecsNicsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlanDataAttributesSpecsNicsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDataAttributesSpecsNicsItems0) UnmarshalBinary(b []byte) error {
	var res PlanDataAttributesSpecsNicsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
