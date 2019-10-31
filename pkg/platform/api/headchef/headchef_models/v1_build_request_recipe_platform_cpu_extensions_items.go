// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1BuildRequestRecipePlatformCPUExtensionsItems CPU Extension
//
// The full CPU extension data model
// swagger:model v1BuildRequestRecipePlatformCpuExtensionsItems
type V1BuildRequestRecipePlatformCPUExtensionsItems struct {
	V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf0

	V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf1

	V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1BuildRequestRecipePlatformCPUExtensionsItems) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf0 = aO0

	// AO1
	var aO1 V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf1 = aO1

	// AO2
	var aO2 V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1BuildRequestRecipePlatformCPUExtensionsItems) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 build request recipe platform Cpu extensions items
func (m *V1BuildRequestRecipePlatformCPUExtensionsItems) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf0
	if err := m.V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf1
	if err := m.V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf2
	if err := m.V1BuildRequestRecipePlatformCPUExtensionsItemsAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformCPUExtensionsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformCPUExtensionsItems) UnmarshalBinary(b []byte) error {
	var res V1BuildRequestRecipePlatformCPUExtensionsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
