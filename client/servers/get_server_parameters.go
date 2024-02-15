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

// NewGetServerParams creates a new GetServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServerParams() *GetServerParams {
	return &GetServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerParamsWithTimeout creates a new GetServerParams object
// with the ability to set a timeout on a request.
func NewGetServerParamsWithTimeout(timeout time.Duration) *GetServerParams {
	return &GetServerParams{
		timeout: timeout,
	}
}

// NewGetServerParamsWithContext creates a new GetServerParams object
// with the ability to set a context for a request.
func NewGetServerParamsWithContext(ctx context.Context) *GetServerParams {
	return &GetServerParams{
		Context: ctx,
	}
}

// NewGetServerParamsWithHTTPClient creates a new GetServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServerParamsWithHTTPClient(client *http.Client) *GetServerParams {
	return &GetServerParams{
		HTTPClient: client,
	}
}

type GetServerParams struct {
	ExtraFieldsServers *string
	ServerID           string `json:"id"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerParams) WithDefaults() *GetServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServerParams) SetDefaults() {
	val := GetServerParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get server params
func (o *GetServerParams) WithTimeout(timeout time.Duration) *GetServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server params
func (o *GetServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server params
func (o *GetServerParams) WithContext(ctx context.Context) *GetServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server params
func (o *GetServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server params
func (o *GetServerParams) WithHTTPClient(client *http.Client) *GetServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server params
func (o *GetServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtraFieldsServers adds the extraFieldsServers to the get server params
func (o *GetServerParams) WithExtraFieldsServers(extraFieldsServers *string) *GetServerParams {
	o.SetExtraFieldsServers(extraFieldsServers)
	return o
}

// SetExtraFieldsServers adds the extraFieldsServers to the get server params
func (o *GetServerParams) SetExtraFieldsServers(extraFieldsServers *string) {
	o.ExtraFieldsServers = extraFieldsServers
}

// WithServerID adds the serverID to the get server params
func (o *GetServerParams) WithServerID(serverID string) *GetServerParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the get server params
func (o *GetServerParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExtraFieldsServers != nil {

		// query param extra_fields[servers]
		var qrExtraFieldsServers string

		if o.ExtraFieldsServers != nil {
			qrExtraFieldsServers = *o.ExtraFieldsServers
		}
		qExtraFieldsServers := qrExtraFieldsServers
		if qExtraFieldsServers != "" {

			if err := r.SetQueryParam("extra_fields[servers]", qExtraFieldsServers); err != nil {
				return err
			}
		}
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
