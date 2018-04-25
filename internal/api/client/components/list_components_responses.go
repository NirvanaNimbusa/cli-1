// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/internal/api/models"
)

// ListComponentsReader is a Reader for the ListComponents structure.
type ListComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListComponentsOK creates a ListComponentsOK with default headers values
func NewListComponentsOK() *ListComponentsOK {
	return &ListComponentsOK{}
}

/*ListComponentsOK handles this case with default header values.

Success
*/
type ListComponentsOK struct {
	Payload []*models.Component
}

func (o *ListComponentsOK) Error() string {
	return fmt.Sprintf("[GET /components][%d] listComponentsOK  %+v", 200, o.Payload)
}

func (o *ListComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
