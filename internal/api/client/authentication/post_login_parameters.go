// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/internal/api/models"
)

// NewPostLoginParams creates a new PostLoginParams object
// with the default values initialized.
func NewPostLoginParams() *PostLoginParams {
	var ()
	return &PostLoginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLoginParamsWithTimeout creates a new PostLoginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLoginParamsWithTimeout(timeout time.Duration) *PostLoginParams {
	var ()
	return &PostLoginParams{

		timeout: timeout,
	}
}

// NewPostLoginParamsWithContext creates a new PostLoginParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLoginParamsWithContext(ctx context.Context) *PostLoginParams {
	var ()
	return &PostLoginParams{

		Context: ctx,
	}
}

// NewPostLoginParamsWithHTTPClient creates a new PostLoginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLoginParamsWithHTTPClient(client *http.Client) *PostLoginParams {
	var ()
	return &PostLoginParams{
		HTTPClient: client,
	}
}

/*PostLoginParams contains all the parameters to send to the API endpoint
for the post login operation typically these are written to a http.Request
*/
type PostLoginParams struct {

	/*Credentials*/
	Credentials *models.Credentials

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post login params
func (o *PostLoginParams) WithTimeout(timeout time.Duration) *PostLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post login params
func (o *PostLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post login params
func (o *PostLoginParams) WithContext(ctx context.Context) *PostLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post login params
func (o *PostLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post login params
func (o *PostLoginParams) WithHTTPClient(client *http.Client) *PostLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post login params
func (o *PostLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentials adds the credentials to the post login params
func (o *PostLoginParams) WithCredentials(credentials *models.Credentials) *PostLoginParams {
	o.SetCredentials(credentials)
	return o
}

// SetCredentials adds the credentials to the post login params
func (o *PostLoginParams) SetCredentials(credentials *models.Credentials) {
	o.Credentials = credentials
}

// WriteToRequest writes these params to a swagger request
func (o *PostLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credentials != nil {
		if err := r.SetBodyParam(o.Credentials); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
