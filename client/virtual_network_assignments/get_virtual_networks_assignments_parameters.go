package virtual_network_assignments

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

// NewGetVirtualNetworksAssignmentsParams creates a new GetVirtualNetworksAssignmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualNetworksAssignmentsParams() *GetVirtualNetworksAssignmentsParams {
	return &GetVirtualNetworksAssignmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualNetworksAssignmentsParamsWithTimeout creates a new GetVirtualNetworksAssignmentsParams object
// with the ability to set a timeout on a request.
func NewGetVirtualNetworksAssignmentsParamsWithTimeout(timeout time.Duration) *GetVirtualNetworksAssignmentsParams {
	return &GetVirtualNetworksAssignmentsParams{
		timeout: timeout,
	}
}

// NewGetVirtualNetworksAssignmentsParamsWithContext creates a new GetVirtualNetworksAssignmentsParams object
// with the ability to set a context for a request.
func NewGetVirtualNetworksAssignmentsParamsWithContext(ctx context.Context) *GetVirtualNetworksAssignmentsParams {
	return &GetVirtualNetworksAssignmentsParams{
		Context: ctx,
	}
}

// NewGetVirtualNetworksAssignmentsParamsWithHTTPClient creates a new GetVirtualNetworksAssignmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualNetworksAssignmentsParamsWithHTTPClient(client *http.Client) *GetVirtualNetworksAssignmentsParams {
	return &GetVirtualNetworksAssignmentsParams{
		HTTPClient: client,
	}
}

/*
GetVirtualNetworksAssignmentsParams contains all the parameters to send to the API endpoint

	for the get virtual networks assignments operation.

	Typically these are written to a http.Request.
*/
type GetVirtualNetworksAssignmentsParams struct {
	FilterServer           string `json:"server"`
	FilterVid              string `json:"vid"`
	FilterVirtualNetworkID string `json:"virtual_network_id"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual networks assignments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualNetworksAssignmentsParams) WithDefaults() *GetVirtualNetworksAssignmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual networks assignments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualNetworksAssignmentsParams) SetDefaults() {
	val := GetVirtualNetworksAssignmentsParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) WithTimeout(timeout time.Duration) *GetVirtualNetworksAssignmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) WithContext(ctx context.Context) *GetVirtualNetworksAssignmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) WithHTTPClient(client *http.Client) *GetVirtualNetworksAssignmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterServer adds the filterServer to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) WithFilterServer(filterServer string) *GetVirtualNetworksAssignmentsParams {
	o.SetFilterServer(filterServer)
	return o
}

// SetFilterServer adds the filterServer to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) SetFilterServer(filterServer string) {
	o.FilterServer = filterServer
}

// WithFilterVid adds the filterVid to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) WithFilterVid(filterVid string) *GetVirtualNetworksAssignmentsParams {
	o.SetFilterVid(filterVid)
	return o
}

// SetFilterVid adds the filterVid to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) SetFilterVid(filterVid string) {
	o.FilterVid = filterVid
}

// WithFilterVirtualNetworkID adds the filterVirtualNetworkID to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) WithFilterVirtualNetworkID(filterVirtualNetworkID string) *GetVirtualNetworksAssignmentsParams {
	o.SetFilterVirtualNetworkID(filterVirtualNetworkID)
	return o
}

// SetFilterVirtualNetworkID adds the filterVirtualNetworkId to the get virtual networks assignments params
func (o *GetVirtualNetworksAssignmentsParams) SetFilterVirtualNetworkID(filterVirtualNetworkID string) {
	o.FilterVirtualNetworkID = filterVirtualNetworkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualNetworksAssignmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetQueryParam("page[size]", "100"); err != nil {
		return err
	}

	if !swag.IsZero(o.FilterServer) {
		if err := r.SetQueryParam("filter[server]", o.FilterServer); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterVirtualNetworkID) {
		if err := r.SetQueryParam("filter[virtual_network_id]", o.FilterVirtualNetworkID); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterVid) {
		if err := r.SetQueryParam("filter[vid]", o.FilterVid); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
