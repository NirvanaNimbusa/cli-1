// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// NewAddBuildScriptParams creates a new AddBuildScriptParams object
// with the default values initialized.
func NewAddBuildScriptParams() *AddBuildScriptParams {
	var ()
	return &AddBuildScriptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddBuildScriptParamsWithTimeout creates a new AddBuildScriptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddBuildScriptParamsWithTimeout(timeout time.Duration) *AddBuildScriptParams {
	var ()
	return &AddBuildScriptParams{

		timeout: timeout,
	}
}

// NewAddBuildScriptParamsWithContext creates a new AddBuildScriptParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddBuildScriptParamsWithContext(ctx context.Context) *AddBuildScriptParams {
	var ()
	return &AddBuildScriptParams{

		Context: ctx,
	}
}

// NewAddBuildScriptParamsWithHTTPClient creates a new AddBuildScriptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddBuildScriptParamsWithHTTPClient(client *http.Client) *AddBuildScriptParams {
	var ()
	return &AddBuildScriptParams{
		HTTPClient: client,
	}
}

/*AddBuildScriptParams contains all the parameters to send to the API endpoint
for the add build script operation typically these are written to a http.Request
*/
type AddBuildScriptParams struct {

	/*BuildScript*/
	BuildScript *inventory_models.BuildScriptCore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add build script params
func (o *AddBuildScriptParams) WithTimeout(timeout time.Duration) *AddBuildScriptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add build script params
func (o *AddBuildScriptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add build script params
func (o *AddBuildScriptParams) WithContext(ctx context.Context) *AddBuildScriptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add build script params
func (o *AddBuildScriptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add build script params
func (o *AddBuildScriptParams) WithHTTPClient(client *http.Client) *AddBuildScriptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add build script params
func (o *AddBuildScriptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildScript adds the buildScript to the add build script params
func (o *AddBuildScriptParams) WithBuildScript(buildScript *inventory_models.BuildScriptCore) *AddBuildScriptParams {
	o.SetBuildScript(buildScript)
	return o
}

// SetBuildScript adds the buildScript to the add build script params
func (o *AddBuildScriptParams) SetBuildScript(buildScript *inventory_models.BuildScriptCore) {
	o.BuildScript = buildScript
}

// WriteToRequest writes these params to a swagger request
func (o *AddBuildScriptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BuildScript != nil {
		if err := r.SetBodyParam(o.BuildScript); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
