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

// UnnormalizedNames Unnormalized Names
//
// A list of unnormalized names.
//
// swagger:model unnormalizedNames
type UnnormalizedNames struct {

	// The unnormalized names
	// Required: true
	Names []string `json:"names"`
}

// Validate validates this unnormalized names
func (m *UnnormalizedNames) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnnormalizedNames) validateNames(formats strfmt.Registry) error {

	if err := validate.Required("names", "body", m.Names); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnnormalizedNames) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnnormalizedNames) UnmarshalBinary(b []byte) error {
	var res UnnormalizedNames
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}