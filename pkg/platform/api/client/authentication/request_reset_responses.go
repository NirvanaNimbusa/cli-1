// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/pkg/platform/api/models"
)

// RequestResetReader is a Reader for the RequestReset structure.
type RequestResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRequestResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRequestResetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRequestResetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRequestResetOK creates a RequestResetOK with default headers values
func NewRequestResetOK() *RequestResetOK {
	return &RequestResetOK{}
}

/*RequestResetOK handles this case with default header values.

Success
*/
type RequestResetOK struct {
	Payload *models.Message
}

func (o *RequestResetOK) Error() string {
	return fmt.Sprintf("[POST /request-reset/{email}][%d] requestResetOK  %+v", 200, o.Payload)
}

func (o *RequestResetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestResetBadRequest creates a RequestResetBadRequest with default headers values
func NewRequestResetBadRequest() *RequestResetBadRequest {
	return &RequestResetBadRequest{}
}

/*RequestResetBadRequest handles this case with default header values.

Bad Request
*/
type RequestResetBadRequest struct {
	Payload *models.Message
}

func (o *RequestResetBadRequest) Error() string {
	return fmt.Sprintf("[POST /request-reset/{email}][%d] requestResetBadRequest  %+v", 400, o.Payload)
}

func (o *RequestResetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestResetInternalServerError creates a RequestResetInternalServerError with default headers values
func NewRequestResetInternalServerError() *RequestResetInternalServerError {
	return &RequestResetInternalServerError{}
}

/*RequestResetInternalServerError handles this case with default header values.

Server Error
*/
type RequestResetInternalServerError struct {
	Payload *models.Message
}

func (o *RequestResetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /request-reset/{email}][%d] requestResetInternalServerError  %+v", 500, o.Payload)
}

func (o *RequestResetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
