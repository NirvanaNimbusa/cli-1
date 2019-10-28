// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1IngredientVersionRevisionCreate Ingredient Version Revision Create
//
// All updatable fields of an ingredient version revision plus resources that can be linked to the revison when it is created, but that are not part of the ingredient version revision model itself.
// swagger:model v1IngredientVersionRevisionCreate
type V1IngredientVersionRevisionCreate struct {
	V1IngredientVersionRevisionCreateAllOf0

	V1IngredientVersionRevisionCreateAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1IngredientVersionRevisionCreate) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1IngredientVersionRevisionCreateAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1IngredientVersionRevisionCreateAllOf0 = aO0

	// AO1
	var aO1 V1IngredientVersionRevisionCreateAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1IngredientVersionRevisionCreateAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1IngredientVersionRevisionCreate) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1IngredientVersionRevisionCreateAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1IngredientVersionRevisionCreateAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 ingredient version revision create
func (m *V1IngredientVersionRevisionCreate) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1IngredientVersionRevisionCreateAllOf0
	if err := m.V1IngredientVersionRevisionCreateAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1IngredientVersionRevisionCreateAllOf1
	if err := m.V1IngredientVersionRevisionCreateAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientVersionRevisionCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientVersionRevisionCreate) UnmarshalBinary(b []byte) error {
	var res V1IngredientVersionRevisionCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}