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

// V1OrderPlatformBuildFlagsAdditionalPropertiesItems v1 order platform build flags additional properties items
//
// swagger:model v1OrderPlatformBuildFlagsAdditionalPropertiesItems
type V1OrderPlatformBuildFlagsAdditionalPropertiesItems struct {

	// The name of the build flag being passed. Must match the name of an existing build flag.
	// Required: true
	Name *string `json:"name"`

	// The value to pass for the build flag. If omitted, the first of the build flag's default values whose conditions are met is used.
	Value string `json:"value,omitempty"`
}

// Validate validates this v1 order platform build flags additional properties items
func (m *V1OrderPlatformBuildFlagsAdditionalPropertiesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OrderPlatformBuildFlagsAdditionalPropertiesItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OrderPlatformBuildFlagsAdditionalPropertiesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OrderPlatformBuildFlagsAdditionalPropertiesItems) UnmarshalBinary(b []byte) error {
	var res V1OrderPlatformBuildFlagsAdditionalPropertiesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
