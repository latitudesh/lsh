package virtual_network_assignments

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewAssignServerVirtualNetworkParams creates a new AssignServerVirtualNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignServerVirtualNetworkParams() *AssignServerVirtualNetworkParams {
	return &AssignServerVirtualNetworkParams{
		timeout: cr.DefaultTimeout,
		Body: AssignServerVirtualNetworkBody{
			Data: &AssignServerVirtualNetworkParamsBodyData{
				Type:       &virtualNetworkAssignmentType,
				Attributes: &AssignServerVirtualNetworkParamsBodyDataAttributes{},
			},
		},
	}
}

// NewAssignServerVirtualNetworkParamsWithTimeout creates a new AssignServerVirtualNetworkParams object
// with the ability to set a timeout on a request.
func NewAssignServerVirtualNetworkParamsWithTimeout(timeout time.Duration) *AssignServerVirtualNetworkParams {
	return &AssignServerVirtualNetworkParams{
		timeout: timeout,
	}
}

// NewAssignServerVirtualNetworkParamsWithContext creates a new AssignServerVirtualNetworkParams object
// with the ability to set a context for a request.
func NewAssignServerVirtualNetworkParamsWithContext(ctx context.Context) *AssignServerVirtualNetworkParams {
	return &AssignServerVirtualNetworkParams{
		Context: ctx,
	}
}

// NewAssignServerVirtualNetworkParamsWithHTTPClient creates a new AssignServerVirtualNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignServerVirtualNetworkParamsWithHTTPClient(client *http.Client) *AssignServerVirtualNetworkParams {
	return &AssignServerVirtualNetworkParams{
		HTTPClient: client,
	}
}

/*
AssignServerVirtualNetworkParams contains all the parameters to send to the API endpoint

	for the assign server virtual network operation.

	Typically these are written to a http.Request.
*/
type AssignServerVirtualNetworkParams struct {

	// Body.
	Body AssignServerVirtualNetworkBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the assign server virtual network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignServerVirtualNetworkParams) WithDefaults() *AssignServerVirtualNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assign server virtual network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignServerVirtualNetworkParams) SetDefaults() {
	val := AssignServerVirtualNetworkParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the assign server virtual network params
func (o *AssignServerVirtualNetworkParams) WithTimeout(timeout time.Duration) *AssignServerVirtualNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign server virtual network params
func (o *AssignServerVirtualNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign server virtual network params
func (o *AssignServerVirtualNetworkParams) WithContext(ctx context.Context) *AssignServerVirtualNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign server virtual network params
func (o *AssignServerVirtualNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign server virtual network params
func (o *AssignServerVirtualNetworkParams) WithHTTPClient(client *http.Client) *AssignServerVirtualNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign server virtual network params
func (o *AssignServerVirtualNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the assign server virtual network params
func (o *AssignServerVirtualNetworkParams) WithBody(body AssignServerVirtualNetworkBody) *AssignServerVirtualNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the assign server virtual network params
func (o *AssignServerVirtualNetworkParams) SetBody(body AssignServerVirtualNetworkBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AssignServerVirtualNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
