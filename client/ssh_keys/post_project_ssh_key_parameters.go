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

// NewPostProjectSSHKeyParams creates a new PostProjectSSHKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostProjectSSHKeyParams() *PostProjectSSHKeyParams {
	return &PostProjectSSHKeyParams{
		timeout: cr.DefaultTimeout,
		Body: PostProjectSSHKeyBody{
			Data: &PostProjectSSHKeyParamsBodyData{
				Type:       &sshKeysType,
				Attributes: &PostProjectSSHKeyParamsBodyDataAttributes{},
			},
		},
	}
}

// NewPostProjectSSHKeyParamsWithTimeout creates a new PostProjectSSHKeyParams object
// with the ability to set a timeout on a request.
func NewPostProjectSSHKeyParamsWithTimeout(timeout time.Duration) *PostProjectSSHKeyParams {
	return &PostProjectSSHKeyParams{
		timeout: timeout,
	}
}

// NewPostProjectSSHKeyParamsWithContext creates a new PostProjectSSHKeyParams object
// with the ability to set a context for a request.
func NewPostProjectSSHKeyParamsWithContext(ctx context.Context) *PostProjectSSHKeyParams {
	return &PostProjectSSHKeyParams{
		Context: ctx,
	}
}

// NewPostProjectSSHKeyParamsWithHTTPClient creates a new PostProjectSSHKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostProjectSSHKeyParamsWithHTTPClient(client *http.Client) *PostProjectSSHKeyParams {
	return &PostProjectSSHKeyParams{
		HTTPClient: client,
	}
}

/*
PostProjectSSHKeyParams contains all the parameters to send to the API endpoint

	for the post project ssh key operation.

	Typically these are written to a http.Request.
*/
type PostProjectSSHKeyParams struct {

	// Body.
	Body PostProjectSSHKeyBody

	// ProjectIDOrSlug.
	ProjectIDOrSlug string `json:"project"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post project ssh key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProjectSSHKeyParams) WithDefaults() *PostProjectSSHKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post project ssh key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProjectSSHKeyParams) SetDefaults() {
	val := PostProjectSSHKeyParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post project ssh key params
func (o *PostProjectSSHKeyParams) WithTimeout(timeout time.Duration) *PostProjectSSHKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post project ssh key params
func (o *PostProjectSSHKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post project ssh key params
func (o *PostProjectSSHKeyParams) WithContext(ctx context.Context) *PostProjectSSHKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post project ssh key params
func (o *PostProjectSSHKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post project ssh key params
func (o *PostProjectSSHKeyParams) WithHTTPClient(client *http.Client) *PostProjectSSHKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post project ssh key params
func (o *PostProjectSSHKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post project ssh key params
func (o *PostProjectSSHKeyParams) WithBody(body PostProjectSSHKeyBody) *PostProjectSSHKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post project ssh key params
func (o *PostProjectSSHKeyParams) SetBody(body PostProjectSSHKeyBody) {
	o.Body = body
}

// WithProjectIDOrSlug adds the projectIDOrSlug to the post project ssh key params
func (o *PostProjectSSHKeyParams) WithProjectIDOrSlug(projectIDOrSlug string) *PostProjectSSHKeyParams {
	o.SetProjectIDOrSlug(projectIDOrSlug)
	return o
}

// SetProjectIDOrSlug adds the projectIdOrSlug to the post project ssh key params
func (o *PostProjectSSHKeyParams) SetProjectIDOrSlug(projectIDOrSlug string) {
	o.ProjectIDOrSlug = projectIDOrSlug
}

// WriteToRequest writes these params to a swagger request
func (o *PostProjectSSHKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
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
