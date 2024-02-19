package servers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewServerUnscheduleDeletionParams creates a new ServerUnscheduleDeletionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServerUnscheduleDeletionParams() *ServerUnscheduleDeletionParams {
	return &ServerUnscheduleDeletionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServerUnscheduleDeletionParamsWithTimeout creates a new ServerUnscheduleDeletionParams object
// with the ability to set a timeout on a request.
func NewServerUnscheduleDeletionParamsWithTimeout(timeout time.Duration) *ServerUnscheduleDeletionParams {
	return &ServerUnscheduleDeletionParams{
		timeout: timeout,
	}
}

// NewServerUnscheduleDeletionParamsWithContext creates a new ServerUnscheduleDeletionParams object
// with the ability to set a context for a request.
func NewServerUnscheduleDeletionParamsWithContext(ctx context.Context) *ServerUnscheduleDeletionParams {
	return &ServerUnscheduleDeletionParams{
		Context: ctx,
	}
}

// NewServerUnscheduleDeletionParamsWithHTTPClient creates a new ServerUnscheduleDeletionParams object
// with the ability to set a custom HTTPClient for a request.
func NewServerUnscheduleDeletionParamsWithHTTPClient(client *http.Client) *ServerUnscheduleDeletionParams {
	return &ServerUnscheduleDeletionParams{
		HTTPClient: client,
	}
}

/*
ServerUnscheduleDeletionParams contains all the parameters to send to the API endpoint

	for the server unschedule deletion operation.

	Typically these are written to a http.Request.
*/
type ServerUnscheduleDeletionParams struct {

	// ServerID.
	ServerID string `json:"id"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the server unschedule deletion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerUnscheduleDeletionParams) WithDefaults() *ServerUnscheduleDeletionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the server unschedule deletion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerUnscheduleDeletionParams) SetDefaults() {
	val := ServerUnscheduleDeletionParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the server unschedule deletion params
func (o *ServerUnscheduleDeletionParams) WithTimeout(timeout time.Duration) *ServerUnscheduleDeletionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the server unschedule deletion params
func (o *ServerUnscheduleDeletionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the server unschedule deletion params
func (o *ServerUnscheduleDeletionParams) WithContext(ctx context.Context) *ServerUnscheduleDeletionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the server unschedule deletion params
func (o *ServerUnscheduleDeletionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the server unschedule deletion params
func (o *ServerUnscheduleDeletionParams) WithHTTPClient(client *http.Client) *ServerUnscheduleDeletionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the server unschedule deletion params
func (o *ServerUnscheduleDeletionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServerID adds the serverID to the server unschedule deletion params
func (o *ServerUnscheduleDeletionParams) WithServerID(serverID string) *ServerUnscheduleDeletionParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the server unschedule deletion params
func (o *ServerUnscheduleDeletionParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *ServerUnscheduleDeletionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param server_id
	if err := r.SetPathParam("server_id", o.ServerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
