// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User user
// swagger:model User
type User struct {

	// e u l a accepted
	EULAAccepted *strfmt.DateTime `json:"EULAAccepted,omitempty"`

	// t o t p enabled
	TOTPEnabled *bool `json:"TOTPEnabled,omitempty"`

	// acl
	ACL map[string]bool `json:"acl,omitempty"`

	// added
	Added strfmt.DateTime `json:"added,omitempty"`

	// datetime format
	DatetimeFormat string `json:"datetimeFormat,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// expires
	Expires *strfmt.DateTime `json:"expires,omitempty"`

	// last login
	LastLogin *strfmt.DateTime `json:"lastLogin,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organizations
	Organizations []*UserOrganizationsItems `json:"organizations"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// user ID
	UserID strfmt.UUID `json:"userID,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEULAAccepted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAdded(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExpires(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastLogin(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateEULAAccepted(formats strfmt.Registry) error {

	if swag.IsZero(m.EULAAccepted) { // not required
		return nil
	}

	if err := validate.FormatOf("EULAAccepted", "body", "date-time", m.EULAAccepted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateAdded(formats strfmt.Registry) error {

	if swag.IsZero(m.Added) { // not required
		return nil
	}

	if err := validate.FormatOf("added", "body", "date-time", m.Added.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateExpires(formats strfmt.Registry) error {

	if swag.IsZero(m.Expires) { // not required
		return nil
	}

	if err := validate.FormatOf("expires", "body", "date-time", m.Expires.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateLastLogin(formats strfmt.Registry) error {

	if swag.IsZero(m.LastLogin) { // not required
		return nil
	}

	if err := validate.FormatOf("lastLogin", "body", "date-time", m.LastLogin.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateOrganizations(formats strfmt.Registry) error {

	if swag.IsZero(m.Organizations) { // not required
		return nil
	}

	for i := 0; i < len(m.Organizations); i++ {

		if swag.IsZero(m.Organizations[i]) { // not required
			continue
		}

		if m.Organizations[i] != nil {

			if err := m.Organizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizations" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *User) validateUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.UserID) { // not required
		return nil
	}

	if err := validate.FormatOf("userID", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
