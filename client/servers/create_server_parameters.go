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

// NewCreateServerParams creates a new CreateServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateServerParams() *CreateServerParams {
	return &CreateServerParams{
		timeout: cr.DefaultTimeout,
		Body: CreateServerBody{
			Data: &CreateServerParamsBodyData{
				Type:       &serverType,
				Attributes: &CreateServerParamsBodyDataAttributes{},
			},
		},
	}
}

// NewCreateServerParamsWithTimeout creates a new CreateServerParams object
// with the ability to set a timeout on a request.
func NewCreateServerParamsWithTimeout(timeout time.Duration) *CreateServerParams {
	return &CreateServerParams{
		timeout: timeout,
	}
}

// NewCreateServerParamsWithContext creates a new CreateServerParams object
// with the ability to set a context for a request.
func NewCreateServerParamsWithContext(ctx context.Context) *CreateServerParams {
	return &CreateServerParams{
		Context: ctx,
	}
}

// NewCreateServerParamsWithHTTPClient creates a new CreateServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateServerParamsWithHTTPClient(client *http.Client) *CreateServerParams {
	return &CreateServerParams{
		HTTPClient: client,
	}
}

/*
CreateServerParams contains all the parameters to send to the API endpoint

	for the create server operation.

	Typically these are written to a http.Request.
*/
type CreateServerParams struct {

	// Body.
	Body CreateServerBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServerParams) WithDefaults() *CreateServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServerParams) SetDefaults() {
	val := CreateServerParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create server params
func (o *CreateServerParams) WithTimeout(timeout time.Duration) *CreateServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create server params
func (o *CreateServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create server params
func (o *CreateServerParams) WithContext(ctx context.Context) *CreateServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create server params
func (o *CreateServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create server params
func (o *CreateServerParams) WithHTTPClient(client *http.Client) *CreateServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create server params
func (o *CreateServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create server params
func (o *CreateServerParams) WithBody(body CreateServerBody) *CreateServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create server params
func (o *CreateServerParams) SetBody(body CreateServerBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
