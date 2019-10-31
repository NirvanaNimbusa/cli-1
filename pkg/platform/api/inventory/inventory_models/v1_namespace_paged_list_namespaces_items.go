// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1NamespacePagedListNamespacesItems namespace
//
// The full namespace data model
// swagger:model v1NamespacePagedListNamespacesItems
type V1NamespacePagedListNamespacesItems struct {
	V1NamespacePagedListNamespacesItemsAllOf0

	V1NamespacePagedListNamespacesItemsAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1NamespacePagedListNamespacesItems) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1NamespacePagedListNamespacesItemsAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1NamespacePagedListNamespacesItemsAllOf0 = aO0

	// AO1
	var aO1 V1NamespacePagedListNamespacesItemsAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1NamespacePagedListNamespacesItemsAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1NamespacePagedListNamespacesItems) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1NamespacePagedListNamespacesItemsAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1NamespacePagedListNamespacesItemsAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 namespace paged list namespaces items
func (m *V1NamespacePagedListNamespacesItems) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1NamespacePagedListNamespacesItemsAllOf0
	if err := m.V1NamespacePagedListNamespacesItemsAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1NamespacePagedListNamespacesItemsAllOf1
	if err := m.V1NamespacePagedListNamespacesItemsAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1NamespacePagedListNamespacesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NamespacePagedListNamespacesItems) UnmarshalBinary(b []byte) error {
	var res V1NamespacePagedListNamespacesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
