// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CPUArchitectureAllOf0 cpu architecture all of0
//
// swagger:model cpuArchitectureAllOf0
type CPUArchitectureAllOf0 struct {

	// cpu architecture id
	// Required: true
	// Format: uuid
	CPUArchitectureID *strfmt.UUID `json:"cpu_architecture_id"`

	// links
	// Required: true
	Links *SelfLink `json:"links"`
}

// Validate validates this cpu architecture all of0
func (m *CPUArchitectureAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUArchitectureID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CPUArchitectureAllOf0) validateCPUArchitectureID(formats strfmt.Registry) error {

	if err := validate.Required("cpu_architecture_id", "body", m.CPUArchitectureID); err != nil {
		return err
	}

	if err := validate.FormatOf("cpu_architecture_id", "body", "uuid", m.CPUArchitectureID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CPUArchitectureAllOf0) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CPUArchitectureAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPUArchitectureAllOf0) UnmarshalBinary(b []byte) error {
	var res CPUArchitectureAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
