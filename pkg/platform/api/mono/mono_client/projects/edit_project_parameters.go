// Code generated by go-swagger; DO NOT EDIT.

package projects

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

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// NewEditProjectParams creates a new EditProjectParams object
// with the default values initialized.
func NewEditProjectParams() *EditProjectParams {
	var ()
	return &EditProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditProjectParamsWithTimeout creates a new EditProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditProjectParamsWithTimeout(timeout time.Duration) *EditProjectParams {
	var ()
	return &EditProjectParams{

		timeout: timeout,
	}
}

// NewEditProjectParamsWithContext creates a new EditProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditProjectParamsWithContext(ctx context.Context) *EditProjectParams {
	var ()
	return &EditProjectParams{

		Context: ctx,
	}
}

// NewEditProjectParamsWithHTTPClient creates a new EditProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditProjectParamsWithHTTPClient(client *http.Client) *EditProjectParams {
	var ()
	return &EditProjectParams{
		HTTPClient: client,
	}
}

/*EditProjectParams contains all the parameters to send to the API endpoint
for the edit project operation typically these are written to a http.Request
*/
type EditProjectParams struct {

	/*OrganizationName
	  organizationName of desired organization

	*/
	OrganizationName string
	/*Project*/
	Project *mono_models.Project
	/*ProjectName
	  projectName of desired project

	*/
	ProjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit project params
func (o *EditProjectParams) WithTimeout(timeout time.Duration) *EditProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit project params
func (o *EditProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit project params
func (o *EditProjectParams) WithContext(ctx context.Context) *EditProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit project params
func (o *EditProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit project params
func (o *EditProjectParams) WithHTTPClient(client *http.Client) *EditProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit project params
func (o *EditProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationName adds the organizationName to the edit project params
func (o *EditProjectParams) WithOrganizationName(organizationName string) *EditProjectParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the edit project params
func (o *EditProjectParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WithProject adds the project to the edit project params
func (o *EditProjectParams) WithProject(project *mono_models.Project) *EditProjectParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the edit project params
func (o *EditProjectParams) SetProject(project *mono_models.Project) {
	o.Project = project
}

// WithProjectName adds the projectName to the edit project params
func (o *EditProjectParams) WithProjectName(projectName string) *EditProjectParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the edit project params
func (o *EditProjectParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WriteToRequest writes these params to a swagger request
func (o *EditProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	if o.Project != nil {
		if err := r.SetBodyParam(o.Project); err != nil {
			return err
		}
	}

	// path param projectName
	if err := r.SetPathParam("projectName", o.ProjectName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
