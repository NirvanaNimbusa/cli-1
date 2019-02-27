// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListOrganizationsParams creates a new ListOrganizationsParams object
// with the default values initialized.
func NewListOrganizationsParams() *ListOrganizationsParams {
	var ()
	return &ListOrganizationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOrganizationsParamsWithTimeout creates a new ListOrganizationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOrganizationsParamsWithTimeout(timeout time.Duration) *ListOrganizationsParams {
	var ()
	return &ListOrganizationsParams{

		timeout: timeout,
	}
}

// NewListOrganizationsParamsWithContext creates a new ListOrganizationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOrganizationsParamsWithContext(ctx context.Context) *ListOrganizationsParams {
	var ()
	return &ListOrganizationsParams{

		Context: ctx,
	}
}

// NewListOrganizationsParamsWithHTTPClient creates a new ListOrganizationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOrganizationsParamsWithHTTPClient(client *http.Client) *ListOrganizationsParams {
	var ()
	return &ListOrganizationsParams{
		HTTPClient: client,
	}
}

/*ListOrganizationsParams contains all the parameters to send to the API endpoint
for the list organizations operation typically these are written to a http.Request
*/
type ListOrganizationsParams struct {

	/*MemberOnly
	  Return all orgs a user has access to or only the ones a user is a member of

	*/
	MemberOnly *bool
	/*Personal
	  Filter based on if the org is a personal or not

	*/
	Personal *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list organizations params
func (o *ListOrganizationsParams) WithTimeout(timeout time.Duration) *ListOrganizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list organizations params
func (o *ListOrganizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list organizations params
func (o *ListOrganizationsParams) WithContext(ctx context.Context) *ListOrganizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list organizations params
func (o *ListOrganizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list organizations params
func (o *ListOrganizationsParams) WithHTTPClient(client *http.Client) *ListOrganizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list organizations params
func (o *ListOrganizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMemberOnly adds the memberOnly to the list organizations params
func (o *ListOrganizationsParams) WithMemberOnly(memberOnly *bool) *ListOrganizationsParams {
	o.SetMemberOnly(memberOnly)
	return o
}

// SetMemberOnly adds the memberOnly to the list organizations params
func (o *ListOrganizationsParams) SetMemberOnly(memberOnly *bool) {
	o.MemberOnly = memberOnly
}

// WithPersonal adds the personal to the list organizations params
func (o *ListOrganizationsParams) WithPersonal(personal *bool) *ListOrganizationsParams {
	o.SetPersonal(personal)
	return o
}

// SetPersonal adds the personal to the list organizations params
func (o *ListOrganizationsParams) SetPersonal(personal *bool) {
	o.Personal = personal
}

// WriteToRequest writes these params to a swagger request
func (o *ListOrganizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MemberOnly != nil {

		// query param memberOnly
		var qrMemberOnly bool
		if o.MemberOnly != nil {
			qrMemberOnly = *o.MemberOnly
		}
		qMemberOnly := swag.FormatBool(qrMemberOnly)
		if qMemberOnly != "" {
			if err := r.SetQueryParam("memberOnly", qMemberOnly); err != nil {
				return err
			}
		}

	}

	if o.Personal != nil {

		// query param personal
		var qrPersonal bool
		if o.Personal != nil {
			qrPersonal = *o.Personal
		}
		qPersonal := swag.FormatBool(qrPersonal)
		if qPersonal != "" {
			if err := r.SetQueryParam("personal", qPersonal); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}