// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IdentityEditable identity editable
// swagger:model IdentityEditable
type IdentityEditable struct {

	// identity ID
	// Format: uuid
	IdentityID *strfmt.UUID `json:"identityID,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization ID
	// Format: uuid
	OrganizationID strfmt.UUID `json:"organizationID,omitempty"`
}

// Validate validates this identity editable
func (m *IdentityEditable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityEditable) validateIdentityID(formats strfmt.Registry) error {

	if swag.IsZero(m.IdentityID) { // not required
		return nil
	}

	if err := validate.FormatOf("identityID", "body", "uuid", m.IdentityID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IdentityEditable) validateOrganizationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organizationID", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdentityEditable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityEditable) UnmarshalBinary(b []byte) error {
	var res IdentityEditable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
