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

// V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf1AllOf0 v1 ingredient create all of0 versions items all of2 all of1 all of1 all of0
// swagger:model v1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf1AllOf0
type V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf1AllOf0 struct {

	// provided features
	// Required: true
	ProvidedFeatures []*V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf1AllOf0ProvidedFeaturesItems `json:"provided_features"`
}

// Validate validates this v1 ingredient create all of0 versions items all of2 all of1 all of1 all of0
func (m *V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf1AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvidedFeatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf1AllOf0) validateProvidedFeatures(formats strfmt.Registry) error {

	if err := validate.Required("provided_features", "body", m.ProvidedFeatures); err != nil {
		return err
	}

	for i := 0; i < len(m.ProvidedFeatures); i++ {
		if swag.IsZero(m.ProvidedFeatures[i]) { // not required
			continue
		}

		if m.ProvidedFeatures[i] != nil {
			if err := m.ProvidedFeatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("provided_features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf1AllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf1AllOf0) UnmarshalBinary(b []byte) error {
	var res V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf1AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
