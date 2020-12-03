// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlatformCreate Platform (create)
//
// A platform upon which a build can be built. This model is only used when a platform is created to allow for both referencing existing platform components or defining new ones.
//
// swagger:model platformCreate
type PlatformCreate struct {

	// The ID of the CPU architecture used by this platform.
	// Required: true
	// Format: uuid
	CPUArchitectureID *strfmt.UUID `json:"cpu_architecture_id"`

	// The IDs of the CPU extensions guaranteed to be present on this platform (if any).
	// Required: true
	CPUExtensions []strfmt.UUID `json:"cpu_extensions"`

	// The last day on which this platform will be supported. Can be omitted if no last day has yet been determined.
	// Format: date
	EndOfSupportDate strfmt.Date `json:"end_of_support_date,omitempty"`

	// The ID of the GPU architecture installed on this platform. Can be omitted if the platform doesn't have a GPU.
	// Format: uuid
	GpuArchitectureID strfmt.UUID `json:"gpu_architecture_id,omitempty"`

	// If true, the platform should be shown to the user as a selectable platform for an order. If false, the platform should be hidden from the user.
	IsUserVisible *bool `json:"is_user_visible,omitempty"`

	// The ID of the kernel version running on this platform's operating system.
	// Required: true
	// Format: uuid
	KernelVersionID *strfmt.UUID `json:"kernel_version_id"`

	// The ID of the libc version installed on this platform. Can be omitted if the platform's operating system does not have a libc.
	// Format: uuid
	LibcVersionID strfmt.UUID `json:"libc_version_id,omitempty"`

	// The ID of the operation system version running on this platform.
	// Required: true
	// Format: uuid
	OperatingSystemVersionID *strfmt.UUID `json:"operating_system_version_id"`
}

// Validate validates this platform create
func (m *PlatformCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUArchitectureID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndOfSupportDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpuArchitectureID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKernelVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLibcVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingSystemVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformCreate) validateCPUArchitectureID(formats strfmt.Registry) error {

	if err := validate.Required("cpu_architecture_id", "body", m.CPUArchitectureID); err != nil {
		return err
	}

	if err := validate.FormatOf("cpu_architecture_id", "body", "uuid", m.CPUArchitectureID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PlatformCreate) validateCPUExtensions(formats strfmt.Registry) error {

	if err := validate.Required("cpu_extensions", "body", m.CPUExtensions); err != nil {
		return err
	}

	for i := 0; i < len(m.CPUExtensions); i++ {

		if err := validate.FormatOf("cpu_extensions"+"."+strconv.Itoa(i), "body", "uuid", m.CPUExtensions[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *PlatformCreate) validateEndOfSupportDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndOfSupportDate) { // not required
		return nil
	}

	if err := validate.FormatOf("end_of_support_date", "body", "date", m.EndOfSupportDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PlatformCreate) validateGpuArchitectureID(formats strfmt.Registry) error {

	if swag.IsZero(m.GpuArchitectureID) { // not required
		return nil
	}

	if err := validate.FormatOf("gpu_architecture_id", "body", "uuid", m.GpuArchitectureID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PlatformCreate) validateKernelVersionID(formats strfmt.Registry) error {

	if err := validate.Required("kernel_version_id", "body", m.KernelVersionID); err != nil {
		return err
	}

	if err := validate.FormatOf("kernel_version_id", "body", "uuid", m.KernelVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PlatformCreate) validateLibcVersionID(formats strfmt.Registry) error {

	if swag.IsZero(m.LibcVersionID) { // not required
		return nil
	}

	if err := validate.FormatOf("libc_version_id", "body", "uuid", m.LibcVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PlatformCreate) validateOperatingSystemVersionID(formats strfmt.Registry) error {

	if err := validate.Required("operating_system_version_id", "body", m.OperatingSystemVersionID); err != nil {
		return err
	}

	if err := validate.FormatOf("operating_system_version_id", "body", "uuid", m.OperatingSystemVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlatformCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformCreate) UnmarshalBinary(b []byte) error {
	var res PlatformCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
