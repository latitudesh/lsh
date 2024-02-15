package virtual_networks

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetVirtualNetworksParams creates a new GetVirtualNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualNetworksParams() *GetVirtualNetworksParams {
	return &GetVirtualNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualNetworksParamsWithTimeout creates a new GetVirtualNetworksParams object
// with the ability to set a timeout on a request.
func NewGetVirtualNetworksParamsWithTimeout(timeout time.Duration) *GetVirtualNetworksParams {
	return &GetVirtualNetworksParams{
		timeout: timeout,
	}
}

// NewGetVirtualNetworksParamsWithContext creates a new GetVirtualNetworksParams object
// with the ability to set a context for a request.
func NewGetVirtualNetworksParamsWithContext(ctx context.Context) *GetVirtualNetworksParams {
	return &GetVirtualNetworksParams{
		Context: ctx,
	}
}

// NewGetVirtualNetworksParamsWithHTTPClient creates a new GetVirtualNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualNetworksParamsWithHTTPClient(client *http.Client) *GetVirtualNetworksParams {
	return &GetVirtualNetworksParams{
		HTTPClient: client,
	}
}

/*
GetVirtualNetworksParams contains all the parameters to send to the API endpoint

	for the get virtual networks operation.

	Typically these are written to a http.Request.
*/
type GetVirtualNetworksParams struct {
	FilterLocation string `json:"location"`
	FilterProject  string `json:"project"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualNetworksParams) WithDefaults() *GetVirtualNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualNetworksParams) SetDefaults() {
	val := GetVirtualNetworksParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get virtual networks params
func (o *GetVirtualNetworksParams) WithTimeout(timeout time.Duration) *GetVirtualNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual networks params
func (o *GetVirtualNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual networks params
func (o *GetVirtualNetworksParams) WithContext(ctx context.Context) *GetVirtualNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual networks params
func (o *GetVirtualNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual networks params
func (o *GetVirtualNetworksParams) WithHTTPClient(client *http.Client) *GetVirtualNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual networks params
func (o *GetVirtualNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterLocation adds the filterLocation to the get virtual networks params
func (o *GetVirtualNetworksParams) WithFilterLocation(filterLocation string) *GetVirtualNetworksParams {
	o.SetFilterLocation(filterLocation)
	return o
}

// SetFilterLocation adds the filterLocation to the get virtual networks params
func (o *GetVirtualNetworksParams) SetFilterLocation(filterLocation string) {
	o.FilterLocation = filterLocation
}

// WithFilterProject adds the filterProject to the get virtual networks params
func (o *GetVirtualNetworksParams) WithFilterProject(filterProject string) *GetVirtualNetworksParams {
	o.SetFilterProject(filterProject)
	return o
}

// SetFilterProject adds the filterProject to the get virtual networks params
func (o *GetVirtualNetworksParams) SetFilterProject(filterProject string) {
	o.FilterProject = filterProject
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetQueryParam("page[size]", "100"); err != nil {
		return err
	}

	if !swag.IsZero(o.FilterLocation) {
		if err := r.SetQueryParam("filter[location]", o.FilterLocation); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterProject) {
		if err := r.SetQueryParam("filter[project]", o.FilterProject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
