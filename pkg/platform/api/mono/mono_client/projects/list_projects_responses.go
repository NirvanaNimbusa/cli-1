// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// ListProjectsReader is a Reader for the ListProjects structure.
type ListProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewListProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListProjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListProjectsOK creates a ListProjectsOK with default headers values
func NewListProjectsOK() *ListProjectsOK {
	return &ListProjectsOK{}
}

/*ListProjectsOK handles this case with default header values.

Success
*/
type ListProjectsOK struct {
	Payload []*mono_models.Project
}

func (o *ListProjectsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects][%d] listProjectsOK  %+v", 200, o.Payload)
}

func (o *ListProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsNotFound creates a ListProjectsNotFound with default headers values
func NewListProjectsNotFound() *ListProjectsNotFound {
	return &ListProjectsNotFound{}
}

/*ListProjectsNotFound handles this case with default header values.

Not Found
*/
type ListProjectsNotFound struct {
	Payload *mono_models.Message
}

func (o *ListProjectsNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects][%d] listProjectsNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectsInternalServerError creates a ListProjectsInternalServerError with default headers values
func NewListProjectsInternalServerError() *ListProjectsInternalServerError {
	return &ListProjectsInternalServerError{}
}

/*ListProjectsInternalServerError handles this case with default header values.

Server Error
*/
type ListProjectsInternalServerError struct {
	Payload *mono_models.Message
}

func (o *ListProjectsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects][%d] listProjectsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
