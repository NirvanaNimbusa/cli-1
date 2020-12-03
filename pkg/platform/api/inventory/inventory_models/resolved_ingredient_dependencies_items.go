// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResolvedIngredientDependenciesItems resolved ingredient dependencies items
//
// swagger:model resolvedIngredientDependenciesItems
type ResolvedIngredientDependenciesItems struct {

	// The type or types of dependencies on this ingredient version.
	// Required: true
	// Min Items: 1
	// Unique: true
	DependencyTypes []DependencyType `json:"dependency_types"`

	// ingredient version id
	// Required: true
	// Format: uuid
	IngredientVersionID *strfmt.UUID `json:"ingredient_version_id"`
}

// Validate validates this resolved ingredient dependencies items
func (m *ResolvedIngredientDependenciesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependencyTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResolvedIngredientDependenciesItems) validateDependencyTypes(formats strfmt.Registry) error {

	if err := validate.Required("dependency_types", "body", m.DependencyTypes); err != nil {
		return err
	}

	iDependencyTypesSize := int64(len(m.DependencyTypes))

	if err := validate.MinItems("dependency_types", "body", iDependencyTypesSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("dependency_types", "body", m.DependencyTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.DependencyTypes); i++ {

		if err := m.DependencyTypes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dependency_types" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ResolvedIngredientDependenciesItems) validateIngredientVersionID(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_version_id", "body", m.IngredientVersionID); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient_version_id", "body", "uuid", m.IngredientVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResolvedIngredientDependenciesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResolvedIngredientDependenciesItems) UnmarshalBinary(b []byte) error {
	var res ResolvedIngredientDependenciesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
