// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewJoinOrganizationParams creates a new JoinOrganizationParams object
// with the default values initialized.
func NewJoinOrganizationParams() *JoinOrganizationParams {
	var ()
	return &JoinOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJoinOrganizationParamsWithTimeout creates a new JoinOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJoinOrganizationParamsWithTimeout(timeout time.Duration) *JoinOrganizationParams {
	var ()
	return &JoinOrganizationParams{

		timeout: timeout,
	}
}

// NewJoinOrganizationParamsWithContext creates a new JoinOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewJoinOrganizationParamsWithContext(ctx context.Context) *JoinOrganizationParams {
	var ()
	return &JoinOrganizationParams{

		Context: ctx,
	}
}

// NewJoinOrganizationParamsWithHTTPClient creates a new JoinOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJoinOrganizationParamsWithHTTPClient(client *http.Client) *JoinOrganizationParams {
	var ()
	return &JoinOrganizationParams{
		HTTPClient: client,
	}
}

/*JoinOrganizationParams contains all the parameters to send to the API endpoint
for the join organization operation typically these are written to a http.Request
*/
type JoinOrganizationParams struct {

	/*InviteCode
	  Invite code

	*/
	InviteCode *models.JoinOrganizationParamsBody
	/*OrganizationName
	  organizationID of desired organization

	*/
	OrganizationName string
	/*Username
	  username to join

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the join organization params
func (o *JoinOrganizationParams) WithTimeout(timeout time.Duration) *JoinOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the join organization params
func (o *JoinOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the join organization params
func (o *JoinOrganizationParams) WithContext(ctx context.Context) *JoinOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the join organization params
func (o *JoinOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the join organization params
func (o *JoinOrganizationParams) WithHTTPClient(client *http.Client) *JoinOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the join organization params
func (o *JoinOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInviteCode adds the inviteCode to the join organization params
func (o *JoinOrganizationParams) WithInviteCode(inviteCode *models.JoinOrganizationParamsBody) *JoinOrganizationParams {
	o.SetInviteCode(inviteCode)
	return o
}

// SetInviteCode adds the inviteCode to the join organization params
func (o *JoinOrganizationParams) SetInviteCode(inviteCode *models.JoinOrganizationParamsBody) {
	o.InviteCode = inviteCode
}

// WithOrganizationName adds the organizationName to the join organization params
func (o *JoinOrganizationParams) WithOrganizationName(organizationName string) *JoinOrganizationParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the join organization params
func (o *JoinOrganizationParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WithUsername adds the username to the join organization params
func (o *JoinOrganizationParams) WithUsername(username string) *JoinOrganizationParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the join organization params
func (o *JoinOrganizationParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *JoinOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InviteCode != nil {
		if err := r.SetBodyParam(o.InviteCode); err != nil {
			return err
		}
	}

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
