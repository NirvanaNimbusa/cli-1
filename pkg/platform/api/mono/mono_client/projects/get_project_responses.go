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

// GetProjectReader is a Reader for the GetProject structure.
type GetProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectOK creates a GetProjectOK with default headers values
func NewGetProjectOK() *GetProjectOK {
	return &GetProjectOK{}
}

/*GetProjectOK handles this case with default header values.

Success
*/
type GetProjectOK struct {
	Payload *mono_models.Project
}

func (o *GetProjectOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}][%d] getProjectOK  %+v", 200, o.Payload)
}

func (o *GetProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectNotFound creates a GetProjectNotFound with default headers values
func NewGetProjectNotFound() *GetProjectNotFound {
	return &GetProjectNotFound{}
}

/*GetProjectNotFound handles this case with default header values.

Not Found
*/
type GetProjectNotFound struct {
	Payload *mono_models.Message
}

func (o *GetProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}][%d] getProjectNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectInternalServerError creates a GetProjectInternalServerError with default headers values
func NewGetProjectInternalServerError() *GetProjectInternalServerError {
	return &GetProjectInternalServerError{}
}

/*GetProjectInternalServerError handles this case with default header values.

Server Error
*/
type GetProjectInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}][%d] getProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
