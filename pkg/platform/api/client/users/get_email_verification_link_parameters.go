// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetEmailVerificationLinkParams creates a new GetEmailVerificationLinkParams object
// with the default values initialized.
func NewGetEmailVerificationLinkParams() *GetEmailVerificationLinkParams {
	var ()
	return &GetEmailVerificationLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmailVerificationLinkParamsWithTimeout creates a new GetEmailVerificationLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmailVerificationLinkParamsWithTimeout(timeout time.Duration) *GetEmailVerificationLinkParams {
	var ()
	return &GetEmailVerificationLinkParams{

		timeout: timeout,
	}
}

// NewGetEmailVerificationLinkParamsWithContext creates a new GetEmailVerificationLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmailVerificationLinkParamsWithContext(ctx context.Context) *GetEmailVerificationLinkParams {
	var ()
	return &GetEmailVerificationLinkParams{

		Context: ctx,
	}
}

// NewGetEmailVerificationLinkParamsWithHTTPClient creates a new GetEmailVerificationLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmailVerificationLinkParamsWithHTTPClient(client *http.Client) *GetEmailVerificationLinkParams {
	var ()
	return &GetEmailVerificationLinkParams{
		HTTPClient: client,
	}
}

/*GetEmailVerificationLinkParams contains all the parameters to send to the API endpoint
for the get email verification link operation typically these are written to a http.Request
*/
type GetEmailVerificationLinkParams struct {

	/*Email
	  Email address for which to get the verification link

	*/
	Email string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get email verification link params
func (o *GetEmailVerificationLinkParams) WithTimeout(timeout time.Duration) *GetEmailVerificationLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get email verification link params
func (o *GetEmailVerificationLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get email verification link params
func (o *GetEmailVerificationLinkParams) WithContext(ctx context.Context) *GetEmailVerificationLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get email verification link params
func (o *GetEmailVerificationLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get email verification link params
func (o *GetEmailVerificationLinkParams) WithHTTPClient(client *http.Client) *GetEmailVerificationLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get email verification link params
func (o *GetEmailVerificationLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the get email verification link params
func (o *GetEmailVerificationLinkParams) WithEmail(email string) *GetEmailVerificationLinkParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the get email verification link params
func (o *GetEmailVerificationLinkParams) SetEmail(email string) {
	o.Email = email
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmailVerificationLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param email
	if err := r.SetPathParam("email", o.Email); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}