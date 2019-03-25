// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	secrets_models "github.com/ActiveState/cli/pkg/platform/api/secrets/secrets_models"
)

// GetWhoamiReader is a Reader for the GetWhoami structure.
type GetWhoamiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWhoamiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWhoamiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetWhoamiUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWhoamiOK creates a GetWhoamiOK with default headers values
func NewGetWhoamiOK() *GetWhoamiOK {
	return &GetWhoamiOK{}
}

/*GetWhoamiOK handles this case with default header values.

Success
*/
type GetWhoamiOK struct {
	Payload *secrets_models.UID
}

func (o *GetWhoamiOK) Error() string {
	return fmt.Sprintf("[GET /whoami][%d] getWhoamiOK  %+v", 200, o.Payload)
}

func (o *GetWhoamiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(secrets_models.UID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWhoamiUnauthorized creates a GetWhoamiUnauthorized with default headers values
func NewGetWhoamiUnauthorized() *GetWhoamiUnauthorized {
	return &GetWhoamiUnauthorized{}
}

/*GetWhoamiUnauthorized handles this case with default header values.

Invalid credentials
*/
type GetWhoamiUnauthorized struct {
	Payload *secrets_models.Message
}

func (o *GetWhoamiUnauthorized) Error() string {
	return fmt.Sprintf("[GET /whoami][%d] getWhoamiUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWhoamiUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(secrets_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
