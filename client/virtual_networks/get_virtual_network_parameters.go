package virtual_networks

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetVirtualNetworkParams creates a new GetVirtualNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualNetworkParams() *GetVirtualNetworkParams {
	return &GetVirtualNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualNetworkParamsWithTimeout creates a new GetVirtualNetworkParams object
// with the ability to set a timeout on a request.
func NewGetVirtualNetworkParamsWithTimeout(timeout time.Duration) *GetVirtualNetworkParams {
	return &GetVirtualNetworkParams{
		timeout: timeout,
	}
}

// NewGetVirtualNetworkParamsWithContext creates a new GetVirtualNetworkParams object
// with the ability to set a context for a request.
func NewGetVirtualNetworkParamsWithContext(ctx context.Context) *GetVirtualNetworkParams {
	return &GetVirtualNetworkParams{
		Context: ctx,
	}
}

// NewGetVirtualNetworkParamsWithHTTPClient creates a new GetVirtualNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualNetworkParamsWithHTTPClient(client *http.Client) *GetVirtualNetworkParams {
	return &GetVirtualNetworkParams{
		HTTPClient: client,
	}
}

/*
GetVirtualNetworkParams contains all the parameters to send to the API endpoint

	for the get virtual network operation.

	Typically these are written to a http.Request.
*/
type GetVirtualNetworkParams struct {
	// ID.
	ID string `json:"id"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualNetworkParams) WithDefaults() *GetVirtualNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualNetworkParams) SetDefaults() {
	val := GetVirtualNetworkParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get virtual network params
func (o *GetVirtualNetworkParams) WithTimeout(timeout time.Duration) *GetVirtualNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual network params
func (o *GetVirtualNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual network params
func (o *GetVirtualNetworkParams) WithContext(ctx context.Context) *GetVirtualNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual network params
func (o *GetVirtualNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual network params
func (o *GetVirtualNetworkParams) WithHTTPClient(client *http.Client) *GetVirtualNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual network params
func (o *GetVirtualNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get virtual network params
func (o *GetVirtualNetworkParams) WithID(id string) *GetVirtualNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get virtual network params
func (o *GetVirtualNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
