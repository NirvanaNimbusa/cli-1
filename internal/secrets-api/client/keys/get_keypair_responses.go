// Code generated by go-swagger; DO NOT EDIT.

package keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	models "github.com/ActiveState/cli/internal/secrets-api/models"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// GetKeypairReader is a Reader for the GetKeypair structure.
type GetKeypairReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeypairReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKeypairOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetKeypairUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetKeypairNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetKeypairInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKeypairOK creates a GetKeypairOK with default headers values
func NewGetKeypairOK() *GetKeypairOK {
	return &GetKeypairOK{}
}

/*GetKeypairOK handles this case with default header values.

Success
*/
type GetKeypairOK struct {
	Payload *models.Keypair
}

func (o *GetKeypairOK) Error() string {
	return fmt.Sprintf("[GET /keypair][%d] getKeypairOK  %+v", 200, o.Payload)
}

func (o *GetKeypairOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Keypair)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeypairUnauthorized creates a GetKeypairUnauthorized with default headers values
func NewGetKeypairUnauthorized() *GetKeypairUnauthorized {
	return &GetKeypairUnauthorized{}
}

/*GetKeypairUnauthorized handles this case with default header values.

Invalid credentials
*/
type GetKeypairUnauthorized struct {
	Payload *models.Message
}

func (o *GetKeypairUnauthorized) Error() string {
	return fmt.Sprintf("[GET /keypair][%d] getKeypairUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKeypairUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeypairNotFound creates a GetKeypairNotFound with default headers values
func NewGetKeypairNotFound() *GetKeypairNotFound {
	return &GetKeypairNotFound{}
}

/*GetKeypairNotFound handles this case with default header values.

Not Found
*/
type GetKeypairNotFound struct {
	Payload *models.Message
}

func (o *GetKeypairNotFound) Error() string {
	return fmt.Sprintf("[GET /keypair][%d] getKeypairNotFound  %+v", 404, o.Payload)
}

func (o *GetKeypairNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeypairInternalServerError creates a GetKeypairInternalServerError with default headers values
func NewGetKeypairInternalServerError() *GetKeypairInternalServerError {
	return &GetKeypairInternalServerError{}
}

/*GetKeypairInternalServerError handles this case with default header values.

Server Error
*/
type GetKeypairInternalServerError struct {
	Payload *models.Message
}

func (o *GetKeypairInternalServerError) Error() string {
	return fmt.Sprintf("[GET /keypair][%d] getKeypairInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKeypairInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
