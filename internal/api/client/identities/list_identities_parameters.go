// Code generated by go-swagger; DO NOT EDIT.

package identities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListIdentitiesParams creates a new ListIdentitiesParams object
// with the default values initialized.
func NewListIdentitiesParams() *ListIdentitiesParams {
	var ()
	return &ListIdentitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListIdentitiesParamsWithTimeout creates a new ListIdentitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListIdentitiesParamsWithTimeout(timeout time.Duration) *ListIdentitiesParams {
	var ()
	return &ListIdentitiesParams{

		timeout: timeout,
	}
}

// NewListIdentitiesParamsWithContext creates a new ListIdentitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListIdentitiesParamsWithContext(ctx context.Context) *ListIdentitiesParams {
	var ()
	return &ListIdentitiesParams{

		Context: ctx,
	}
}

// NewListIdentitiesParamsWithHTTPClient creates a new ListIdentitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListIdentitiesParamsWithHTTPClient(client *http.Client) *ListIdentitiesParams {
	var ()
	return &ListIdentitiesParams{
		HTTPClient: client,
	}
}

/*ListIdentitiesParams contains all the parameters to send to the API endpoint
for the list identities operation typically these are written to a http.Request
*/
type ListIdentitiesParams struct {

	/*IncludeComponents
	  Include array of associated components for the identity

	*/
	IncludeComponents *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list identities params
func (o *ListIdentitiesParams) WithTimeout(timeout time.Duration) *ListIdentitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list identities params
func (o *ListIdentitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list identities params
func (o *ListIdentitiesParams) WithContext(ctx context.Context) *ListIdentitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list identities params
func (o *ListIdentitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list identities params
func (o *ListIdentitiesParams) WithHTTPClient(client *http.Client) *ListIdentitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list identities params
func (o *ListIdentitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeComponents adds the includeComponents to the list identities params
func (o *ListIdentitiesParams) WithIncludeComponents(includeComponents *bool) *ListIdentitiesParams {
	o.SetIncludeComponents(includeComponents)
	return o
}

// SetIncludeComponents adds the includeComponents to the list identities params
func (o *ListIdentitiesParams) SetIncludeComponents(includeComponents *bool) {
	o.IncludeComponents = includeComponents
}

// WriteToRequest writes these params to a swagger request
func (o *ListIdentitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeComponents != nil {

		// query param includeComponents
		var qrIncludeComponents bool
		if o.IncludeComponents != nil {
			qrIncludeComponents = *o.IncludeComponents
		}
		qIncludeComponents := swag.FormatBool(qrIncludeComponents)
		if qIncludeComponents != "" {
			if err := r.SetQueryParam("includeComponents", qIncludeComponents); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
