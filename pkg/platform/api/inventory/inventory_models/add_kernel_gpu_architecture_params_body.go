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

// AddKernelGpuArchitectureParamsBody add kernel gpu architecture params body
// swagger:model addKernelGpuArchitectureParamsBody
type AddKernelGpuArchitectureParamsBody struct {

	// The ID of the GPU architecture that can be used with this kernel
	// Required: true
	// Format: uuid
	GpuArchitectureID *strfmt.UUID `json:"gpu_architecture_id"`
}

// Validate validates this add kernel gpu architecture params body
func (m *AddKernelGpuArchitectureParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGpuArchitectureID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddKernelGpuArchitectureParamsBody) validateGpuArchitectureID(formats strfmt.Registry) error {

	if err := validate.Required("gpu_architecture_id", "body", m.GpuArchitectureID); err != nil {
		return err
	}

	if err := validate.FormatOf("gpu_architecture_id", "body", "uuid", m.GpuArchitectureID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddKernelGpuArchitectureParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddKernelGpuArchitectureParamsBody) UnmarshalBinary(b []byte) error {
	var res AddKernelGpuArchitectureParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
