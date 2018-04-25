// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/internal/api/models"
)

// ListTokensReader is a Reader for the ListTokens structure.
type ListTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListTokensOK creates a ListTokensOK with default headers values
func NewListTokensOK() *ListTokensOK {
	return &ListTokensOK{}
}

/*ListTokensOK handles this case with default header values.

Success
*/
type ListTokensOK struct {
	Payload []*models.Token
}

func (o *ListTokensOK) Error() string {
	return fmt.Sprintf("[GET /apikeys][%d] listTokensOK  %+v", 200, o.Payload)
}

func (o *ListTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
