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

// V1BuildFlagCoreAllOf1AllOf0DefaultsItemsConditionSetsItems Condition Set Sub Schema
//
// All of the conditions in a condition set must be satisfied in order for the condition set to be considered satisfied (i.e. conditions in a condition set are ANDed together)
// swagger:model v1BuildFlagCoreAllOf1AllOf0DefaultsItemsConditionSetsItems
type V1BuildFlagCoreAllOf1AllOf0DefaultsItemsConditionSetsItems struct {

	// The list of conditions that must all be satisfied for this condition set to be satisfied. These conditions may only condition on platform features, not ingredient features.
	// Required: true
	// Min Items: 1
	Conditions []*V1BuildFlagCoreAllOf1AllOf0DefaultsItemsConditionSetsItemsConditionsItems `json:"conditions"`
}

// Validate validates this v1 build flag core all of1 all of0 defaults items condition sets items
func (m *V1BuildFlagCoreAllOf1AllOf0DefaultsItemsConditionSetsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildFlagCoreAllOf1AllOf0DefaultsItemsConditionSetsItems) validateConditions(formats strfmt.Registry) error {

	if err := validate.Required("conditions", "body", m.Conditions); err != nil {
		return err
	}

	iConditionsSize := int64(len(m.Conditions))

	if err := validate.MinItems("conditions", "body", iConditionsSize, 1); err != nil {
		return err
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

// MarshalBinary interface implementation
func (m *V1BuildFlagCoreAllOf1AllOf0DefaultsItemsConditionSetsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildFlagCoreAllOf1AllOf0DefaultsItemsConditionSetsItems) UnmarshalBinary(b []byte) error {
	var res V1BuildFlagCoreAllOf1AllOf0DefaultsItemsConditionSetsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
