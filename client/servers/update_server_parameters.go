package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewUpdateServerParams creates a new UpdateServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateServerParams() *UpdateServerParams {
	return &UpdateServerParams{
		timeout: cr.DefaultTimeout,
		Body: UpdateServerBody{
			Data: &UpdateServerParamsBodyData{
				Type:       serverType,
				Attributes: &UpdateServerParamsBodyAttributes{},
			},
		},
	}
}

// NewUpdateServerParamsWithTimeout creates a new UpdateServerParams object
// with the ability to set a timeout on a request.
func NewUpdateServerParamsWithTimeout(timeout time.Duration) *UpdateServerParams {
	return &UpdateServerParams{
		timeout: timeout,
	}
}

// NewUpdateServerParamsWithContext creates a new UpdateServerParams object
// with the ability to set a context for a request.
func NewUpdateServerParamsWithContext(ctx context.Context) *UpdateServerParams {
	return &UpdateServerParams{
		Context: ctx,
	}
}

// NewUpdateServerParamsWithHTTPClient creates a new UpdateServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateServerParamsWithHTTPClient(client *http.Client) *UpdateServerParams {
	return &UpdateServerParams{
		HTTPClient: client,
	}
}

/*
UpdateServerParams contains all the parameters to send to the API endpoint

	for the update server operation.

	Typically these are written to a http.Request.
*/
type UpdateServerParams struct {

	// Body.
	Body UpdateServerBody

	ID string `json:"id"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServerParams) WithDefaults() *UpdateServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateServerParams) SetDefaults() {
	val := UpdateServerParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update server params
func (o *UpdateServerParams) WithTimeout(timeout time.Duration) *UpdateServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update server params
func (o *UpdateServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update server params
func (o *UpdateServerParams) WithContext(ctx context.Context) *UpdateServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update server params
func (o *UpdateServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update server params
func (o *UpdateServerParams) WithHTTPClient(client *http.Client) *UpdateServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update server params
func (o *UpdateServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update server params
func (o *UpdateServerParams) WithBody(body UpdateServerBody) *UpdateServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update server params
func (o *UpdateServerParams) SetBody(body UpdateServerBody) {
	o.Body = body
}

// WithID adds the id to the update server params
func (o *UpdateServerParams) WithID(id string) *UpdateServerParams {
	o.SetID(id)
	return o
}

// SetID adds the serverId to the update server params
func (o *UpdateServerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if err := r.SetPathParam("server_id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
