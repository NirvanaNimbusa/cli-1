// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1RecipeResolvedIngredientsItemsPatchesItemsAllOf0AllOf1 Patch Core
//
// A diff of changes that can be applied to an ingredient's source code. This model only contains mutable fields of a patch.
// swagger:model v1RecipeResolvedIngredientsItemsPatchesItemsAllOf0AllOf1
type V1RecipeResolvedIngredientsItemsPatchesItemsAllOf0AllOf1 struct {

	// The features that must already be present in the recipe for this patch to be used. For example, can be used to create patches that only work on specific operating systems.
	Conditions []*V1RecipeResolvedIngredientsItemsPatchesItemsAllOf0AllOf1ConditionsItems `json:"conditions"`

	// The patch itself
	// Required: true
	Content *string `json:"content"`

	// A concise summary of what this patch can be used for
	// Required: true
	Description *string `json:"description"`
}

// Validate validates this v1 recipe resolved ingredients items patches items all of0 all of1
func (m *V1RecipeResolvedIngredientsItemsPatchesItemsAllOf0AllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RecipeResolvedIngredientsItemsPatchesItemsAllOf0AllOf1) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsPatchesItemsAllOf0AllOf1) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsPatchesItemsAllOf0AllOf1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipeResolvedIngredientsItemsPatchesItemsAllOf0AllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipeResolvedIngredientsItemsPatchesItemsAllOf0AllOf1) UnmarshalBinary(b []byte) error {
	var res V1RecipeResolvedIngredientsItemsPatchesItemsAllOf0AllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
