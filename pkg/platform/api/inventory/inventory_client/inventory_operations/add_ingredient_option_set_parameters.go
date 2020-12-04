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

// NewAddIngredientOptionSetParams creates a new AddIngredientOptionSetParams object
// with the default values initialized.
func NewAddIngredientOptionSetParams() *AddIngredientOptionSetParams {
	var ()
	return &AddIngredientOptionSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddIngredientOptionSetParamsWithTimeout creates a new AddIngredientOptionSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddIngredientOptionSetParamsWithTimeout(timeout time.Duration) *AddIngredientOptionSetParams {
	var ()
	return &AddIngredientOptionSetParams{

		timeout: timeout,
	}
}

// NewAddIngredientOptionSetParamsWithContext creates a new AddIngredientOptionSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddIngredientOptionSetParamsWithContext(ctx context.Context) *AddIngredientOptionSetParams {
	var ()
	return &AddIngredientOptionSetParams{

		Context: ctx,
	}
}

// NewAddIngredientOptionSetParamsWithHTTPClient creates a new AddIngredientOptionSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddIngredientOptionSetParamsWithHTTPClient(client *http.Client) *AddIngredientOptionSetParams {
	var ()
	return &AddIngredientOptionSetParams{
		HTTPClient: client,
	}
}

/*AddIngredientOptionSetParams contains all the parameters to send to the API endpoint
for the add ingredient option set operation typically these are written to a http.Request
*/
type AddIngredientOptionSetParams struct {

	/*IngredientOptionSet*/
	IngredientOptionSet *inventory_models.IngredientOptionSetCore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add ingredient option set params
func (o *AddIngredientOptionSetParams) WithTimeout(timeout time.Duration) *AddIngredientOptionSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add ingredient option set params
func (o *AddIngredientOptionSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add ingredient option set params
func (o *AddIngredientOptionSetParams) WithContext(ctx context.Context) *AddIngredientOptionSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add ingredient option set params
func (o *AddIngredientOptionSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add ingredient option set params
func (o *AddIngredientOptionSetParams) WithHTTPClient(client *http.Client) *AddIngredientOptionSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add ingredient option set params
func (o *AddIngredientOptionSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIngredientOptionSet adds the ingredientOptionSet to the add ingredient option set params
func (o *AddIngredientOptionSetParams) WithIngredientOptionSet(ingredientOptionSet *inventory_models.IngredientOptionSetCore) *AddIngredientOptionSetParams {
	o.SetIngredientOptionSet(ingredientOptionSet)
	return o
}

// SetIngredientOptionSet adds the ingredientOptionSet to the add ingredient option set params
func (o *AddIngredientOptionSetParams) SetIngredientOptionSet(ingredientOptionSet *inventory_models.IngredientOptionSetCore) {
	o.IngredientOptionSet = ingredientOptionSet
}

// WriteToRequest writes these params to a swagger request
func (o *AddIngredientOptionSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IngredientOptionSet != nil {
		if err := r.SetBodyParam(o.IngredientOptionSet); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}