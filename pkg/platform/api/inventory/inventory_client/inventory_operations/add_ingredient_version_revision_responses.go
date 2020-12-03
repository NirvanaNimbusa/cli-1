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

// AddIngredientVersionRevisionReader is a Reader for the AddIngredientVersionRevision structure.
type AddIngredientVersionRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddIngredientVersionRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddIngredientVersionRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddIngredientVersionRevisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddIngredientVersionRevisionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddIngredientVersionRevisionOK creates a AddIngredientVersionRevisionOK with default headers values
func NewAddIngredientVersionRevisionOK() *AddIngredientVersionRevisionOK {
	return &AddIngredientVersionRevisionOK{}
}

/*AddIngredientVersionRevisionOK handles this case with default header values.

The updated state of the ingredient version
*/
type AddIngredientVersionRevisionOK struct {
	Payload *inventory_models.IngredientVersion
}

func (o *AddIngredientVersionRevisionOK) Error() string {
	return fmt.Sprintf("[POST /v1/ingredients/{ingredient_id}/versions/{ingredient_version_id}/revisions][%d] addIngredientVersionRevisionOK  %+v", 200, o.Payload)
}

func (o *AddIngredientVersionRevisionOK) GetPayload() *inventory_models.IngredientVersion {
	return o.Payload
}

func (o *AddIngredientVersionRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.IngredientVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddIngredientVersionRevisionBadRequest creates a AddIngredientVersionRevisionBadRequest with default headers values
func NewAddIngredientVersionRevisionBadRequest() *AddIngredientVersionRevisionBadRequest {
	return &AddIngredientVersionRevisionBadRequest{}
}

/*AddIngredientVersionRevisionBadRequest handles this case with default header values.

If the ingredient version revision is invalid
*/
type AddIngredientVersionRevisionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddIngredientVersionRevisionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/ingredients/{ingredient_id}/versions/{ingredient_version_id}/revisions][%d] addIngredientVersionRevisionBadRequest  %+v", 400, o.Payload)
}

func (o *AddIngredientVersionRevisionBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddIngredientVersionRevisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddIngredientVersionRevisionDefault creates a AddIngredientVersionRevisionDefault with default headers values
func NewAddIngredientVersionRevisionDefault(code int) *AddIngredientVersionRevisionDefault {
	return &AddIngredientVersionRevisionDefault{
		_statusCode: code,
	}
}

/*AddIngredientVersionRevisionDefault handles this case with default header values.

If there is an error processing the request
*/
type AddIngredientVersionRevisionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add ingredient version revision default response
func (o *AddIngredientVersionRevisionDefault) Code() int {
	return o._statusCode
}

func (o *AddIngredientVersionRevisionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/ingredients/{ingredient_id}/versions/{ingredient_version_id}/revisions][%d] addIngredientVersionRevision default  %+v", o._statusCode, o.Payload)
}

func (o *AddIngredientVersionRevisionDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddIngredientVersionRevisionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
