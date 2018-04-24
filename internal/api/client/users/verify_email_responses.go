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

// VerifyEmailReader is a Reader for the VerifyEmail structure.
type VerifyEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifyEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVerifyEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewVerifyEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewVerifyEmailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewVerifyEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewVerifyEmailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVerifyEmailOK creates a VerifyEmailOK with default headers values
func NewVerifyEmailOK() *VerifyEmailOK {
	return &VerifyEmailOK{}
}

/*VerifyEmailOK handles this case with default header values.

Email updated
*/
type VerifyEmailOK struct {
	Payload *models.Email
}

func (o *VerifyEmailOK) Error() string {
	return fmt.Sprintf("[POST /users/{username}/emails/{email}/verification/check][%d] verifyEmailOK  %+v", 200, o.Payload)
}

func (o *VerifyEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Email)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyEmailBadRequest creates a VerifyEmailBadRequest with default headers values
func NewVerifyEmailBadRequest() *VerifyEmailBadRequest {
	return &VerifyEmailBadRequest{}
}

/*VerifyEmailBadRequest handles this case with default header values.

Invalid Code
*/
type VerifyEmailBadRequest struct {
	Payload *models.Message
}

func (o *VerifyEmailBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{username}/emails/{email}/verification/check][%d] verifyEmailBadRequest  %+v", 400, o.Payload)
}

func (o *VerifyEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyEmailForbidden creates a VerifyEmailForbidden with default headers values
func NewVerifyEmailForbidden() *VerifyEmailForbidden {
	return &VerifyEmailForbidden{}
}

/*VerifyEmailForbidden handles this case with default header values.

Forbidden
*/
type VerifyEmailForbidden struct {
	Payload *models.Message
}

func (o *VerifyEmailForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{username}/emails/{email}/verification/check][%d] verifyEmailForbidden  %+v", 403, o.Payload)
}

func (o *VerifyEmailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyEmailNotFound creates a VerifyEmailNotFound with default headers values
func NewVerifyEmailNotFound() *VerifyEmailNotFound {
	return &VerifyEmailNotFound{}
}

/*VerifyEmailNotFound handles this case with default header values.

Not Found
*/
type VerifyEmailNotFound struct {
	Payload *models.Message
}

func (o *VerifyEmailNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{username}/emails/{email}/verification/check][%d] verifyEmailNotFound  %+v", 404, o.Payload)
}

func (o *VerifyEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyEmailInternalServerError creates a VerifyEmailInternalServerError with default headers values
func NewVerifyEmailInternalServerError() *VerifyEmailInternalServerError {
	return &VerifyEmailInternalServerError{}
}

/*VerifyEmailInternalServerError handles this case with default header values.

Server Error
*/
type VerifyEmailInternalServerError struct {
	Payload *models.Message
}

func (o *VerifyEmailInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/{username}/emails/{email}/verification/check][%d] verifyEmailInternalServerError  %+v", 500, o.Payload)
}

func (o *VerifyEmailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
