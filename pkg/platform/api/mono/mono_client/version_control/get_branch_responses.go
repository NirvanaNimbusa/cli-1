// Code generated by go-swagger; DO NOT EDIT.

package version_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetBranchReader is a Reader for the GetBranch structure.
type GetBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBranchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetBranchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBranchOK creates a GetBranchOK with default headers values
func NewGetBranchOK() *GetBranchOK {
	return &GetBranchOK{}
}

/*GetBranchOK handles this case with default header values.

Get details about the branch
*/
type GetBranchOK struct {
	Payload *mono_models.Branch
}

func (o *GetBranchOK) Error() string {
	return fmt.Sprintf("[GET /vcs/branch/{branchID}][%d] getBranchOK  %+v", 200, o.Payload)
}

func (o *GetBranchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Branch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBranchNotFound creates a GetBranchNotFound with default headers values
func NewGetBranchNotFound() *GetBranchNotFound {
	return &GetBranchNotFound{}
}

/*GetBranchNotFound handles this case with default header values.

branch was not found
*/
type GetBranchNotFound struct {
	Payload *mono_models.Message
}

func (o *GetBranchNotFound) Error() string {
	return fmt.Sprintf("[GET /vcs/branch/{branchID}][%d] getBranchNotFound  %+v", 404, o.Payload)
}

func (o *GetBranchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
