package server_reinstall

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCreateServerReinstallParams creates a new CreateServerReinstallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateServerReinstallParams() *CreateServerReinstallParams {
	return &CreateServerReinstallParams{
		timeout: cr.DefaultTimeout,
		Body: CreateServerReinstallBody{
			Data: &CreateServerReinstallParamsBodyData{
				Type:       &serverReinstallType,
				Attributes: &CreateServerReinstallParamsBodyDataAttributes{},
			},
		},
	}
}

// NewCreateServerReinstallParamsWithTimeout creates a new CreateServerReinstallParams object
// with the ability to set a timeout on a request.
func NewCreateServerReinstallParamsWithTimeout(timeout time.Duration) *CreateServerReinstallParams {
	return &CreateServerReinstallParams{
		timeout: timeout,
	}
}

// NewCreateServerReinstallParamsWithContext creates a new CreateServerReinstallParams object
// with the ability to set a context for a request.
func NewCreateServerReinstallParamsWithContext(ctx context.Context) *CreateServerReinstallParams {
	return &CreateServerReinstallParams{
		Context: ctx,
	}
}

// NewCreateServerReinstallParamsWithHTTPClient creates a new CreateServerReinstallParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateServerReinstallParamsWithHTTPClient(client *http.Client) *CreateServerReinstallParams {
	return &CreateServerReinstallParams{
		HTTPClient: client,
	}
}

/*
CreateServerReinstallParams contains all the parameters to send to the API endpoint

	for the create server reinstall operation.

	Typically these are written to a http.Request.
*/
type CreateServerReinstallParams struct {

	// Body.
	Body CreateServerReinstallBody

	// ServerID.
	ServerID string `json:"id"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create server reinstall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServerReinstallParams) WithDefaults() *CreateServerReinstallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create server reinstall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServerReinstallParams) SetDefaults() {
	val := CreateServerReinstallParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create server reinstall params
func (o *CreateServerReinstallParams) WithTimeout(timeout time.Duration) *CreateServerReinstallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create server reinstall params
func (o *CreateServerReinstallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create server reinstall params
func (o *CreateServerReinstallParams) WithContext(ctx context.Context) *CreateServerReinstallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create server reinstall params
func (o *CreateServerReinstallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create server reinstall params
func (o *CreateServerReinstallParams) WithHTTPClient(client *http.Client) *CreateServerReinstallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create server reinstall params
func (o *CreateServerReinstallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create server reinstall params
func (o *CreateServerReinstallParams) WithBody(body CreateServerReinstallBody) *CreateServerReinstallParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create server reinstall params
func (o *CreateServerReinstallParams) SetBody(body CreateServerReinstallBody) {
	o.Body = body
}

// WithServerID adds the serverID to the create server reinstall params
func (o *CreateServerReinstallParams) WithServerID(serverID string) *CreateServerReinstallParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the create server reinstall params
func (o *CreateServerReinstallParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServerReinstallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param server_id
	if err := r.SetPathParam("server_id", o.ServerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
