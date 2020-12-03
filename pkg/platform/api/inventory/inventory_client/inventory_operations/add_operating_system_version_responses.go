// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// AddOperatingSystemVersionReader is a Reader for the AddOperatingSystemVersion structure.
type AddOperatingSystemVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOperatingSystemVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddOperatingSystemVersionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddOperatingSystemVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddOperatingSystemVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddOperatingSystemVersionCreated creates a AddOperatingSystemVersionCreated with default headers values
func NewAddOperatingSystemVersionCreated() *AddOperatingSystemVersionCreated {
	return &AddOperatingSystemVersionCreated{}
}

/*AddOperatingSystemVersionCreated handles this case with default header values.

The added operating system version
*/
type AddOperatingSystemVersionCreated struct {
	Payload *inventory_models.OperatingSystemVersion
}

func (o *AddOperatingSystemVersionCreated) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems/{operating_system_id}/versions][%d] addOperatingSystemVersionCreated  %+v", 201, o.Payload)
}

func (o *AddOperatingSystemVersionCreated) GetPayload() *inventory_models.OperatingSystemVersion {
	return o.Payload
}

func (o *AddOperatingSystemVersionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.OperatingSystemVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOperatingSystemVersionBadRequest creates a AddOperatingSystemVersionBadRequest with default headers values
func NewAddOperatingSystemVersionBadRequest() *AddOperatingSystemVersionBadRequest {
	return &AddOperatingSystemVersionBadRequest{}
}

/*AddOperatingSystemVersionBadRequest handles this case with default header values.

If the operating system version is invalid
*/
type AddOperatingSystemVersionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddOperatingSystemVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems/{operating_system_id}/versions][%d] addOperatingSystemVersionBadRequest  %+v", 400, o.Payload)
}

func (o *AddOperatingSystemVersionBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddOperatingSystemVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOperatingSystemVersionDefault creates a AddOperatingSystemVersionDefault with default headers values
func NewAddOperatingSystemVersionDefault(code int) *AddOperatingSystemVersionDefault {
	return &AddOperatingSystemVersionDefault{
		_statusCode: code,
	}
}

/*AddOperatingSystemVersionDefault handles this case with default header values.

If there is an error processing the request
*/
type AddOperatingSystemVersionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add operating system version default response
func (o *AddOperatingSystemVersionDefault) Code() int {
	return o._statusCode
}

func (o *AddOperatingSystemVersionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems/{operating_system_id}/versions][%d] addOperatingSystemVersion default  %+v", o._statusCode, o.Payload)
}

func (o *AddOperatingSystemVersionDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddOperatingSystemVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
