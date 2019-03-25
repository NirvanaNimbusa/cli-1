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

// GetOrganizationMembersReader is a Reader for the GetOrganizationMembers structure.
type GetOrganizationMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganizationMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetOrganizationMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrganizationMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetOrganizationMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationMembersOK creates a GetOrganizationMembersOK with default headers values
func NewGetOrganizationMembersOK() *GetOrganizationMembersOK {
	return &GetOrganizationMembersOK{}
}

/*GetOrganizationMembersOK handles this case with default header values.

Success
*/
type GetOrganizationMembersOK struct {
	Payload []*mono_models.Member
}

func (o *GetOrganizationMembersOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/members][%d] getOrganizationMembersOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationMembersForbidden creates a GetOrganizationMembersForbidden with default headers values
func NewGetOrganizationMembersForbidden() *GetOrganizationMembersForbidden {
	return &GetOrganizationMembersForbidden{}
}

/*GetOrganizationMembersForbidden handles this case with default header values.

Forbidden
*/
type GetOrganizationMembersForbidden struct {
	Payload *mono_models.Message
}

func (o *GetOrganizationMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/members][%d] getOrganizationMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationMembersNotFound creates a GetOrganizationMembersNotFound with default headers values
func NewGetOrganizationMembersNotFound() *GetOrganizationMembersNotFound {
	return &GetOrganizationMembersNotFound{}
}

/*GetOrganizationMembersNotFound handles this case with default header values.

Not Found
*/
type GetOrganizationMembersNotFound struct {
	Payload *mono_models.Message
}

func (o *GetOrganizationMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/members][%d] getOrganizationMembersNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationMembersInternalServerError creates a GetOrganizationMembersInternalServerError with default headers values
func NewGetOrganizationMembersInternalServerError() *GetOrganizationMembersInternalServerError {
	return &GetOrganizationMembersInternalServerError{}
}

/*GetOrganizationMembersInternalServerError handles this case with default header values.

Server Error
*/
type GetOrganizationMembersInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetOrganizationMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/members][%d] getOrganizationMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
