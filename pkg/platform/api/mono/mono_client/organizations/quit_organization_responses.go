// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// QuitOrganizationReader is a Reader for the QuitOrganization structure.
type QuitOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuitOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuitOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuitOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewQuitOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewQuitOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQuitOrganizationOK creates a QuitOrganizationOK with default headers values
func NewQuitOrganizationOK() *QuitOrganizationOK {
	return &QuitOrganizationOK{}
}

/*QuitOrganizationOK handles this case with default header values.

Membership updated
*/
type QuitOrganizationOK struct {
	Payload []*mono_models.Member
}

func (o *QuitOrganizationOK) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationName}/members/{username}][%d] quitOrganizationOK  %+v", 200, o.Payload)
}

func (o *QuitOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuitOrganizationBadRequest creates a QuitOrganizationBadRequest with default headers values
func NewQuitOrganizationBadRequest() *QuitOrganizationBadRequest {
	return &QuitOrganizationBadRequest{}
}

/*QuitOrganizationBadRequest handles this case with default header values.

Bad Request
*/
type QuitOrganizationBadRequest struct {
	Payload *mono_models.Message
}

func (o *QuitOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationName}/members/{username}][%d] quitOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *QuitOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuitOrganizationForbidden creates a QuitOrganizationForbidden with default headers values
func NewQuitOrganizationForbidden() *QuitOrganizationForbidden {
	return &QuitOrganizationForbidden{}
}

/*QuitOrganizationForbidden handles this case with default header values.

Unauthorized
*/
type QuitOrganizationForbidden struct {
	Payload *mono_models.Message
}

func (o *QuitOrganizationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationName}/members/{username}][%d] quitOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *QuitOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuitOrganizationInternalServerError creates a QuitOrganizationInternalServerError with default headers values
func NewQuitOrganizationInternalServerError() *QuitOrganizationInternalServerError {
	return &QuitOrganizationInternalServerError{}
}

/*QuitOrganizationInternalServerError handles this case with default header values.

Server Error
*/
type QuitOrganizationInternalServerError struct {
	Payload *mono_models.Message
}

func (o *QuitOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationName}/members/{username}][%d] quitOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *QuitOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
