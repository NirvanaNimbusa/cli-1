// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Namespace namespace
//
// The full namespace data model
//
// swagger:model namespace
type Namespace struct {
	NamespaceAllOf0

	NamespaceCore
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Namespace) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NamespaceAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NamespaceAllOf0 = aO0

	// AO1
	var aO1 NamespaceCore
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.NamespaceCore = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Namespace) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.NamespaceAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.NamespaceCore)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this namespace
func (m *Namespace) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NamespaceAllOf0
	if err := m.NamespaceAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with NamespaceCore
	if err := m.NamespaceCore.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Namespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Namespace) UnmarshalBinary(b []byte) error {
	var res Namespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
