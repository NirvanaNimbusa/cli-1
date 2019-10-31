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

// V1NamespacePagedListPaging Paging
//
// Paging data
// swagger:model v1NamespacePagedListPaging
type V1NamespacePagedListPaging struct {

	// The number of items on this page
	// Required: true
	// Minimum: 0
	ItemCount *int64 `json:"item_count"`

	// The maximum number of items that could be returned
	// Required: true
	// Minimum: 1
	Limit *int64 `json:"limit"`

	// The page number of this result set
	// Required: true
	// Minimum: 1
	Page *int64 `json:"page"`

	// The total number of pages
	// Required: true
	// Minimum: 1
	PageCount *int64 `json:"page_count"`
}

// Validate validates this v1 namespace paged list paging
func (m *V1NamespacePagedListPaging) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NamespacePagedListPaging) validateItemCount(formats strfmt.Registry) error {

	if err := validate.Required("item_count", "body", m.ItemCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("item_count", "body", int64(*m.ItemCount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *V1NamespacePagedListPaging) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	if err := validate.MinimumInt("limit", "body", int64(*m.Limit), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *V1NamespacePagedListPaging) validatePage(formats strfmt.Registry) error {

	if err := validate.Required("page", "body", m.Page); err != nil {
		return err
	}

	if err := validate.MinimumInt("page", "body", int64(*m.Page), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *V1NamespacePagedListPaging) validatePageCount(formats strfmt.Registry) error {

	if err := validate.Required("page_count", "body", m.PageCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("page_count", "body", int64(*m.PageCount), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1NamespacePagedListPaging) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NamespacePagedListPaging) UnmarshalBinary(b []byte) error {
	var res V1NamespacePagedListPaging
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
