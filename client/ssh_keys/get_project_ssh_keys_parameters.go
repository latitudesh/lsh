package ssh_keys

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetProjectSSHKeysParams creates a new GetProjectSSHKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectSSHKeysParams() *GetProjectSSHKeysParams {
	return &GetProjectSSHKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectSSHKeysParamsWithTimeout creates a new GetProjectSSHKeysParams object
// with the ability to set a timeout on a request.
func NewGetProjectSSHKeysParamsWithTimeout(timeout time.Duration) *GetProjectSSHKeysParams {
	return &GetProjectSSHKeysParams{
		timeout: timeout,
	}
}

// NewGetProjectSSHKeysParamsWithContext creates a new GetProjectSSHKeysParams object
// with the ability to set a context for a request.
func NewGetProjectSSHKeysParamsWithContext(ctx context.Context) *GetProjectSSHKeysParams {
	return &GetProjectSSHKeysParams{
		Context: ctx,
	}
}

// NewGetProjectSSHKeysParamsWithHTTPClient creates a new GetProjectSSHKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectSSHKeysParamsWithHTTPClient(client *http.Client) *GetProjectSSHKeysParams {
	return &GetProjectSSHKeysParams{
		HTTPClient: client,
	}
}

/*
GetProjectSSHKeysParams contains all the parameters to send to the API endpoint

	for the get project ssh keys operation.

	Typically these are written to a http.Request.
*/
type GetProjectSSHKeysParams struct {

	// ProjectIDOrSlug.
	ProjectIDOrSlug string `json:"project"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project ssh keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectSSHKeysParams) WithDefaults() *GetProjectSSHKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project ssh keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectSSHKeysParams) SetDefaults() {
	val := GetProjectSSHKeysParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get project ssh keys params
func (o *GetProjectSSHKeysParams) WithTimeout(timeout time.Duration) *GetProjectSSHKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project ssh keys params
func (o *GetProjectSSHKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project ssh keys params
func (o *GetProjectSSHKeysParams) WithContext(ctx context.Context) *GetProjectSSHKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project ssh keys params
func (o *GetProjectSSHKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project ssh keys params
func (o *GetProjectSSHKeysParams) WithHTTPClient(client *http.Client) *GetProjectSSHKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project ssh keys params
func (o *GetProjectSSHKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectIDOrSlug adds the projectIDOrSlug to the get project ssh keys params
func (o *GetProjectSSHKeysParams) WithProjectIDOrSlug(projectIDOrSlug string) *GetProjectSSHKeysParams {
	o.SetProjectIDOrSlug(projectIDOrSlug)
	return o
}

// SetProjectIDOrSlug adds the projectIdOrSlug to the get project ssh keys params
func (o *GetProjectSSHKeysParams) SetProjectIDOrSlug(projectIDOrSlug string) {
	o.ProjectIDOrSlug = projectIDOrSlug
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectSSHKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetQueryParam("page[size]", "100"); err != nil {
		return err
	}

	// path param project_id_or_slug
	if err := r.SetPathParam("project_id_or_slug", o.ProjectIDOrSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
