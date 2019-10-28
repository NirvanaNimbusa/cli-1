// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1BuildScriptPagedListLinks Paging Links
//
// Links for a model that include links for paging data
// swagger:model v1BuildScriptPagedListLinks
type V1BuildScriptPagedListLinks struct {

	// The URI of the first page
	// Required: true
	// Format: uri
	First *strfmt.URI `json:"first"`

	// The URI of last page
	// Required: true
	// Format: uri
	Last *strfmt.URI `json:"last"`

	// The URI of the next page
	// Format: uri
	Next strfmt.URI `json:"next,omitempty"`

	// The URI of the previous page
	// Format: uri
	Previous strfmt.URI `json:"previous,omitempty"`

	// The URI of this page
	// Required: true
	// Format: uri
	Self *strfmt.URI `json:"self"`
}

// Validate validates this v1 build script paged list links
func (m *V1BuildScriptPagedListLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildScriptPagedListLinks) validateFirst(formats strfmt.Registry) error {

	if err := validate.Required("first", "body", m.First); err != nil {
		return err
	}

	if err := validate.FormatOf("first", "body", "uri", m.First.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildScriptPagedListLinks) validateLast(formats strfmt.Registry) error {

	if err := validate.Required("last", "body", m.Last); err != nil {
		return err
	}

	if err := validate.FormatOf("last", "body", "uri", m.Last.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildScriptPagedListLinks) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("next", "body", "uri", m.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildScriptPagedListLinks) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(m.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("previous", "body", "uri", m.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildScriptPagedListLinks) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("self", "body", m.Self); err != nil {
		return err
	}

	if err := validate.FormatOf("self", "body", "uri", m.Self.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildScriptPagedListLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildScriptPagedListLinks) UnmarshalBinary(b []byte) error {
	var res V1BuildScriptPagedListLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}