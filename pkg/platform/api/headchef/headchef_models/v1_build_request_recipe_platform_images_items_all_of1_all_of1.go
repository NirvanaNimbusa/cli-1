// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1 Image Revision Core
//
// The properties of an image revision needed to create a new one
// swagger:model v1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1
type V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1 struct {
	V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf0

	V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf1

	V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf0 = aO0

	// AO1
	var aO1 V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf1 = aO1

	// AO2
	var aO2 V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 build request recipe platform images items all of1 all of1
func (m *V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf0
	if err := m.V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf1
	if err := m.V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf2
	if err := m.V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1AllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1) UnmarshalBinary(b []byte) error {
	var res V1BuildRequestRecipePlatformImagesItemsAllOf1AllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
