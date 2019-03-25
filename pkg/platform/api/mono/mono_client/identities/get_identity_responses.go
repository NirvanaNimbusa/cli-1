// Code generated by go-swagger; DO NOT EDIT.

package identities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetIdentityReader is a Reader for the GetIdentity structure.
type GetIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetIdentityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIdentityOK creates a GetIdentityOK with default headers values
func NewGetIdentityOK() *GetIdentityOK {
	return &GetIdentityOK{}
}

/*GetIdentityOK handles this case with default header values.

Identity Record
*/
type GetIdentityOK struct {
	Payload *mono_models.Identity
}

func (o *GetIdentityOK) Error() string {
	return fmt.Sprintf("[GET /identities/{identityID}][%d] getIdentityOK  %+v", 200, o.Payload)
}

func (o *GetIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Identity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityNotFound creates a GetIdentityNotFound with default headers values
func NewGetIdentityNotFound() *GetIdentityNotFound {
	return &GetIdentityNotFound{}
}

/*GetIdentityNotFound handles this case with default header values.

Not Found
*/
type GetIdentityNotFound struct {
	Payload *mono_models.Message
}

func (o *GetIdentityNotFound) Error() string {
	return fmt.Sprintf("[GET /identities/{identityID}][%d] getIdentityNotFound  %+v", 404, o.Payload)
}

func (o *GetIdentityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
