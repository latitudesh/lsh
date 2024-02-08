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

// NewDestroyServerParams creates a new DestroyServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDestroyServerParams() *DestroyServerParams {
	return &DestroyServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDestroyServerParamsWithTimeout creates a new DestroyServerParams object
// with the ability to set a timeout on a request.
func NewDestroyServerParamsWithTimeout(timeout time.Duration) *DestroyServerParams {
	return &DestroyServerParams{
		timeout: timeout,
	}
}

// NewDestroyServerParamsWithContext creates a new DestroyServerParams object
// with the ability to set a context for a request.
func NewDestroyServerParamsWithContext(ctx context.Context) *DestroyServerParams {
	return &DestroyServerParams{
		Context: ctx,
	}
}

// NewDestroyServerParamsWithHTTPClient creates a new DestroyServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDestroyServerParamsWithHTTPClient(client *http.Client) *DestroyServerParams {
	return &DestroyServerParams{
		HTTPClient: client,
	}
}

/*
DestroyServerParams contains all the parameters to send to the API endpoint

	for the destroy server operation.

	Typically these are written to a http.Request.
*/
type DestroyServerParams struct {
	ID string `json:"id"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the destroy server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DestroyServerParams) WithDefaults() *DestroyServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the destroy server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DestroyServerParams) SetDefaults() {
	val := DestroyServerParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the destroy server params
func (o *DestroyServerParams) WithTimeout(timeout time.Duration) *DestroyServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destroy server params
func (o *DestroyServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destroy server params
func (o *DestroyServerParams) WithContext(ctx context.Context) *DestroyServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destroy server params
func (o *DestroyServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destroy server params
func (o *DestroyServerParams) WithHTTPClient(client *http.Client) *DestroyServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destroy server params
func (o *DestroyServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the destroy server params
func (o *DestroyServerParams) WithID(id string) *DestroyServerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the destroy server params
func (o *DestroyServerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DestroyServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param server_id
	if err := r.SetPathParam("server_id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
