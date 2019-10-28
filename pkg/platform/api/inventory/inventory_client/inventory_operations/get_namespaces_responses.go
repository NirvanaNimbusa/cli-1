// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	inventory_models "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// GetNamespacesReader is a Reader for the GetNamespaces structure.
type GetNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNamespacesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNamespacesOK creates a GetNamespacesOK with default headers values
func NewGetNamespacesOK() *GetNamespacesOK {
	return &GetNamespacesOK{}
}

/*GetNamespacesOK handles this case with default header values.

A paginated list of namespaces
*/
type GetNamespacesOK struct {
	Payload *inventory_models.V1NamespacePagedList
}

func (o *GetNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /v1/namespaces][%d] getNamespacesOK  %+v", 200, o.Payload)
}

func (o *GetNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1NamespacePagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespacesDefault creates a GetNamespacesDefault with default headers values
func NewGetNamespacesDefault(code int) *GetNamespacesDefault {
	return &GetNamespacesDefault{
		_statusCode: code,
	}
}

/*GetNamespacesDefault handles this case with default header values.

generic error response
*/
type GetNamespacesDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get namespaces default response
func (o *GetNamespacesDefault) Code() int {
	return o._statusCode
}

func (o *GetNamespacesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/namespaces][%d] getNamespaces default  %+v", o._statusCode, o.Payload)
}

func (o *GetNamespacesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}