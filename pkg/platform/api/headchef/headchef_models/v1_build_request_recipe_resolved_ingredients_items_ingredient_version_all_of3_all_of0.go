// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0 v1 build request recipe resolved ingredients items ingredient version all of3 all of0
// swagger:model v1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0
type V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0 struct {

	// Camel-specific metadata needed to build this ingredient version revision in camel, if there is any.
	CamelExtras interface{} `json:"camel_extras,omitempty"`

	// dependency sets
	DependencySets []*V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0DependencySetsItems `json:"dependency_sets"`

	// Whether or not this is a stable release of the package
	IsStableRelease *bool `json:"is_stable_release,omitempty"`

	// S3 URL where the source distribution is stored for our platform
	// Format: uri
	PlatformSourceURI *strfmt.URI `json:"platform_source_uri,omitempty"`

	// A checksum of the source distribution. The actual type of the checksum (MD5, S3 Etag, etc.) is not specified. It's assumed that the system that populates and uses this data will know how to work with these checksums.
	SourceChecksum *string `json:"source_checksum,omitempty"`
}

// Validate validates this v1 build request recipe resolved ingredients items ingredient version all of3 all of0
func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependencySets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformSourceURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0) validateDependencySets(formats strfmt.Registry) error {

	if swag.IsZero(m.DependencySets) { // not required
		return nil
	}

	for i := 0; i < len(m.DependencySets); i++ {
		if swag.IsZero(m.DependencySets[i]) { // not required
			continue
		}

		if m.DependencySets[i] != nil {
			if err := m.DependencySets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependency_sets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0) validatePlatformSourceURI(formats strfmt.Registry) error {

	if swag.IsZero(m.PlatformSourceURI) { // not required
		return nil
	}

	if err := validate.FormatOf("platform_source_uri", "body", "uri", m.PlatformSourceURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0) UnmarshalBinary(b []byte) error {
	var res V1BuildRequestRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
