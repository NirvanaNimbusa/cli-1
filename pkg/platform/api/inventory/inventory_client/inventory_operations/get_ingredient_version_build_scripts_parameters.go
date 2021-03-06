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
	"github.com/go-openapi/swag"
)

// NewGetIngredientVersionBuildScriptsParams creates a new GetIngredientVersionBuildScriptsParams object
// with the default values initialized.
func NewGetIngredientVersionBuildScriptsParams() *GetIngredientVersionBuildScriptsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetIngredientVersionBuildScriptsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIngredientVersionBuildScriptsParamsWithTimeout creates a new GetIngredientVersionBuildScriptsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIngredientVersionBuildScriptsParamsWithTimeout(timeout time.Duration) *GetIngredientVersionBuildScriptsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetIngredientVersionBuildScriptsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		timeout: timeout,
	}
}

// NewGetIngredientVersionBuildScriptsParamsWithContext creates a new GetIngredientVersionBuildScriptsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIngredientVersionBuildScriptsParamsWithContext(ctx context.Context) *GetIngredientVersionBuildScriptsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetIngredientVersionBuildScriptsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		Context: ctx,
	}
}

// NewGetIngredientVersionBuildScriptsParamsWithHTTPClient creates a new GetIngredientVersionBuildScriptsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIngredientVersionBuildScriptsParamsWithHTTPClient(client *http.Client) *GetIngredientVersionBuildScriptsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetIngredientVersionBuildScriptsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,
		HTTPClient:    client,
	}
}

/*GetIngredientVersionBuildScriptsParams contains all the parameters to send to the API endpoint
for the get ingredient version build scripts operation typically these are written to a http.Request
*/
type GetIngredientVersionBuildScriptsParams struct {

	/*AllowUnstable
	  Whether to show an unstable revision of a resource if there is an available unstable version newer than the newest available stable version

	*/
	AllowUnstable *bool
	/*IngredientID*/
	IngredientID strfmt.UUID
	/*IngredientVersionID*/
	IngredientVersionID strfmt.UUID
	/*Limit
	  The maximum number of items returned per page

	*/
	Limit *int64
	/*Page
	  The page number returned

	*/
	Page *int64
	/*StateAt
	  Show the state of a resource as it was at the specified timestamp. If omitted, shows the current state of the resource.

	*/
	StateAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) WithTimeout(timeout time.Duration) *GetIngredientVersionBuildScriptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) WithContext(ctx context.Context) *GetIngredientVersionBuildScriptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) WithHTTPClient(client *http.Client) *GetIngredientVersionBuildScriptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowUnstable adds the allowUnstable to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) WithAllowUnstable(allowUnstable *bool) *GetIngredientVersionBuildScriptsParams {
	o.SetAllowUnstable(allowUnstable)
	return o
}

// SetAllowUnstable adds the allowUnstable to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) SetAllowUnstable(allowUnstable *bool) {
	o.AllowUnstable = allowUnstable
}

// WithIngredientID adds the ingredientID to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) WithIngredientID(ingredientID strfmt.UUID) *GetIngredientVersionBuildScriptsParams {
	o.SetIngredientID(ingredientID)
	return o
}

// SetIngredientID adds the ingredientId to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) SetIngredientID(ingredientID strfmt.UUID) {
	o.IngredientID = ingredientID
}

// WithIngredientVersionID adds the ingredientVersionID to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) WithIngredientVersionID(ingredientVersionID strfmt.UUID) *GetIngredientVersionBuildScriptsParams {
	o.SetIngredientVersionID(ingredientVersionID)
	return o
}

// SetIngredientVersionID adds the ingredientVersionId to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) SetIngredientVersionID(ingredientVersionID strfmt.UUID) {
	o.IngredientVersionID = ingredientVersionID
}

// WithLimit adds the limit to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) WithLimit(limit *int64) *GetIngredientVersionBuildScriptsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) WithPage(page *int64) *GetIngredientVersionBuildScriptsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) SetPage(page *int64) {
	o.Page = page
}

// WithStateAt adds the stateAt to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) WithStateAt(stateAt *strfmt.DateTime) *GetIngredientVersionBuildScriptsParams {
	o.SetStateAt(stateAt)
	return o
}

// SetStateAt adds the stateAt to the get ingredient version build scripts params
func (o *GetIngredientVersionBuildScriptsParams) SetStateAt(stateAt *strfmt.DateTime) {
	o.StateAt = stateAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetIngredientVersionBuildScriptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowUnstable != nil {

		// query param allow_unstable
		var qrAllowUnstable bool
		if o.AllowUnstable != nil {
			qrAllowUnstable = *o.AllowUnstable
		}
		qAllowUnstable := swag.FormatBool(qrAllowUnstable)
		if qAllowUnstable != "" {
			if err := r.SetQueryParam("allow_unstable", qAllowUnstable); err != nil {
				return err
			}
		}

	}

	// path param ingredient_id
	if err := r.SetPathParam("ingredient_id", o.IngredientID.String()); err != nil {
		return err
	}

	// path param ingredient_version_id
	if err := r.SetPathParam("ingredient_version_id", o.IngredientVersionID.String()); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.StateAt != nil {

		// query param state_at
		var qrStateAt strfmt.DateTime
		if o.StateAt != nil {
			qrStateAt = *o.StateAt
		}
		qStateAt := qrStateAt.String()
		if qStateAt != "" {
			if err := r.SetQueryParam("state_at", qStateAt); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
