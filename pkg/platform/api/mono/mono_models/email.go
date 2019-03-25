// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Email email
// swagger:model Email
type Email struct {

	// email
	Email string `json:"email,omitempty"`

	// preferred
	Preferred bool `json:"preferred,omitempty"`

	// verified
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this email
func (m *Email) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Email) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Email) UnmarshalBinary(b []byte) error {
	var res Email
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
