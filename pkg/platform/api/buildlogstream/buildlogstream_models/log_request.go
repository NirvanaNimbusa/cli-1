// Code generated by go-swagger; DO NOT EDIT.

package buildlogstream_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LogRequest log request
//
// swagger:model LogRequest
type LogRequest struct {

	// Request log lines from the build of the specified artifact. Only supported for builds executed by builders.
	// Format: uuid
	ArtifactID strfmt.UUID `json:"artifactID,omitempty"`

	// Request build events and log lines from the specified build request. Only supported for builds executed by camel.
	// Format: uuid
	BuildRequestID strfmt.UUID `json:"buildRequestID,omitempty"`

	// Request build events from the specified recipe. Only supported for builds executed by builders.
	// Format: uuid
	RecipeID strfmt.UUID `json:"recipeID,omitempty"`
}

// Validate validates this log request
func (m *LogRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifactID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogRequest) validateArtifactID(formats strfmt.Registry) error {

	if swag.IsZero(m.ArtifactID) { // not required
		return nil
	}

	if err := validate.FormatOf("artifactID", "body", "uuid", m.ArtifactID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LogRequest) validateBuildRequestID(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildRequestID) { // not required
		return nil
	}

	if err := validate.FormatOf("buildRequestID", "body", "uuid", m.BuildRequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LogRequest) validateRecipeID(formats strfmt.Registry) error {

	if swag.IsZero(m.RecipeID) { // not required
		return nil
	}

	if err := validate.FormatOf("recipeID", "body", "uuid", m.RecipeID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogRequest) UnmarshalBinary(b []byte) error {
	var res LogRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
