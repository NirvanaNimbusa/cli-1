// Code generated by go-swagger; DO NOT EDIT.

package platforms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListPlatformsParams creates a new ListPlatformsParams object
// with the default values initialized.
func NewListPlatformsParams() *ListPlatformsParams {

	return &ListPlatformsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListPlatformsParamsWithTimeout creates a new ListPlatformsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListPlatformsParamsWithTimeout(timeout time.Duration) *ListPlatformsParams {

	return &ListPlatformsParams{

		timeout: timeout,
	}
}

// NewListPlatformsParamsWithContext creates a new ListPlatformsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListPlatformsParamsWithContext(ctx context.Context) *ListPlatformsParams {

	return &ListPlatformsParams{

		Context: ctx,
	}
}

// NewListPlatformsParamsWithHTTPClient creates a new ListPlatformsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListPlatformsParamsWithHTTPClient(client *http.Client) *ListPlatformsParams {

	return &ListPlatformsParams{
		HTTPClient: client,
	}
}

/*ListPlatformsParams contains all the parameters to send to the API endpoint
for the list platforms operation typically these are written to a http.Request
*/
type ListPlatformsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list platforms params
func (o *ListPlatformsParams) WithTimeout(timeout time.Duration) *ListPlatformsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list platforms params
func (o *ListPlatformsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list platforms params
func (o *ListPlatformsParams) WithContext(ctx context.Context) *ListPlatformsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list platforms params
func (o *ListPlatformsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list platforms params
func (o *ListPlatformsParams) WithHTTPClient(client *http.Client) *ListPlatformsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list platforms params
func (o *ListPlatformsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListPlatformsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}