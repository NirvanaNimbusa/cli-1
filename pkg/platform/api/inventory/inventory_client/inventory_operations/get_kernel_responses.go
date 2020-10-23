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

// GetKernelReader is a Reader for the GetKernel structure.
type GetKernelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKernelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKernelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetKernelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKernelOK creates a GetKernelOK with default headers values
func NewGetKernelOK() *GetKernelOK {
	return &GetKernelOK{}
}

/*GetKernelOK handles this case with default header values.

The retrieved kernel
*/
type GetKernelOK struct {
	Payload *inventory_models.V1Kernel
}

func (o *GetKernelOK) Error() string {
	return fmt.Sprintf("[GET /v1/kernels/{kernel_id}][%d] getKernelOK  %+v", 200, o.Payload)
}

func (o *GetKernelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1Kernel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKernelDefault creates a GetKernelDefault with default headers values
func NewGetKernelDefault(code int) *GetKernelDefault {
	return &GetKernelDefault{
		_statusCode: code,
	}
}

/*GetKernelDefault handles this case with default header values.

generic error response
*/
type GetKernelDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get kernel default response
func (o *GetKernelDefault) Code() int {
	return o._statusCode
}

func (o *GetKernelDefault) Error() string {
	return fmt.Sprintf("[GET /v1/kernels/{kernel_id}][%d] getKernel default  %+v", o._statusCode, o.Payload)
}

func (o *GetKernelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
