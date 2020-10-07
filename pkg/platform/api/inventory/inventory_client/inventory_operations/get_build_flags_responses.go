// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// GetBuildFlagsReader is a Reader for the GetBuildFlags structure.
type GetBuildFlagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildFlagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildFlagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBuildFlagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBuildFlagsOK creates a GetBuildFlagsOK with default headers values
func NewGetBuildFlagsOK() *GetBuildFlagsOK {
	return &GetBuildFlagsOK{}
}

/*GetBuildFlagsOK handles this case with default header values.

A paginated list of build flags
*/
type GetBuildFlagsOK struct {
	Payload *inventory_models.V1BuildFlagPagedList
}

func (o *GetBuildFlagsOK) Error() string {
	return fmt.Sprintf("[GET /v1/build-flags][%d] getBuildFlagsOK  %+v", 200, o.Payload)
}

func (o *GetBuildFlagsOK) GetPayload() *inventory_models.V1BuildFlagPagedList {
	return o.Payload
}

func (o *GetBuildFlagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1BuildFlagPagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildFlagsDefault creates a GetBuildFlagsDefault with default headers values
func NewGetBuildFlagsDefault(code int) *GetBuildFlagsDefault {
	return &GetBuildFlagsDefault{
		_statusCode: code,
	}
}

/*GetBuildFlagsDefault handles this case with default header values.

generic error response
*/
type GetBuildFlagsDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get build flags default response
func (o *GetBuildFlagsDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildFlagsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/build-flags][%d] getBuildFlags default  %+v", o._statusCode, o.Payload)
}

func (o *GetBuildFlagsDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetBuildFlagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
