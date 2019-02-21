// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddressInfo address info
// swagger:model AddressInfo
type AddressInfo struct {

	// address1
	// Required: true
	Address1 *string `json:"address1"`

	// address2
	Address2 *string `json:"address2,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// zip
	Zip string `json:"zip,omitempty"`
}

// Validate validates this address info
func (m *AddressInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress1(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddressInfo) validateAddress1(formats strfmt.Registry) error {

	if err := validate.Required("address1", "body", m.Address1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddressInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressInfo) UnmarshalBinary(b []byte) error {
	var res AddressInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
