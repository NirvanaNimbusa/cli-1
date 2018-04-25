// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/internal/api/models"
)

// DeleteEmailReader is a Reader for the DeleteEmail structure.
type DeleteEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteEmailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteEmailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEmailOK creates a DeleteEmailOK with default headers values
func NewDeleteEmailOK() *DeleteEmailOK {
	return &DeleteEmailOK{}
}

/*DeleteEmailOK handles this case with default header values.

Email deleted
*/
type DeleteEmailOK struct {
	Payload *models.Message
}

func (o *DeleteEmailOK) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}/emails][%d] deleteEmailOK  %+v", 200, o.Payload)
}

func (o *DeleteEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEmailBadRequest creates a DeleteEmailBadRequest with default headers values
func NewDeleteEmailBadRequest() *DeleteEmailBadRequest {
	return &DeleteEmailBadRequest{}
}

/*DeleteEmailBadRequest handles this case with default header values.

Bad Request
*/
type DeleteEmailBadRequest struct {
	Payload *models.Message
}

func (o *DeleteEmailBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}/emails][%d] deleteEmailBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEmailForbidden creates a DeleteEmailForbidden with default headers values
func NewDeleteEmailForbidden() *DeleteEmailForbidden {
	return &DeleteEmailForbidden{}
}

/*DeleteEmailForbidden handles this case with default header values.

Forbidden
*/
type DeleteEmailForbidden struct {
	Payload *models.Message
}

func (o *DeleteEmailForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}/emails][%d] deleteEmailForbidden  %+v", 403, o.Payload)
}

func (o *DeleteEmailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEmailNotFound creates a DeleteEmailNotFound with default headers values
func NewDeleteEmailNotFound() *DeleteEmailNotFound {
	return &DeleteEmailNotFound{}
}

/*DeleteEmailNotFound handles this case with default header values.

Not Found
*/
type DeleteEmailNotFound struct {
	Payload *models.Message
}

func (o *DeleteEmailNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}/emails][%d] deleteEmailNotFound  %+v", 404, o.Payload)
}

func (o *DeleteEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEmailInternalServerError creates a DeleteEmailInternalServerError with default headers values
func NewDeleteEmailInternalServerError() *DeleteEmailInternalServerError {
	return &DeleteEmailInternalServerError{}
}

/*DeleteEmailInternalServerError handles this case with default header values.

Server Error
*/
type DeleteEmailInternalServerError struct {
	Payload *models.Message
}

func (o *DeleteEmailInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}/emails][%d] deleteEmailInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteEmailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
