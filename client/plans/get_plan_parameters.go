package plans

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetPlanParams creates a new GetPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlanParams() *GetPlanParams {
	return &GetPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlanParamsWithTimeout creates a new GetPlanParams object
// with the ability to set a timeout on a request.
func NewGetPlanParamsWithTimeout(timeout time.Duration) *GetPlanParams {
	return &GetPlanParams{
		timeout: timeout,
	}
}

// NewGetPlanParamsWithContext creates a new GetPlanParams object
// with the ability to set a context for a request.
func NewGetPlanParamsWithContext(ctx context.Context) *GetPlanParams {
	return &GetPlanParams{
		Context: ctx,
	}
}

// NewGetPlanParamsWithHTTPClient creates a new GetPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlanParamsWithHTTPClient(client *http.Client) *GetPlanParams {
	return &GetPlanParams{
		HTTPClient: client,
	}
}

/*
GetPlanParams contains all the parameters to send to the API endpoint

	for the get plan operation.

	Typically these are written to a http.Request.
*/
type GetPlanParams struct {

	// PlanID.
	PlanID string `json:"id"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlanParams) WithDefaults() *GetPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlanParams) SetDefaults() {
	val := GetPlanParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get plan params
func (o *GetPlanParams) WithTimeout(timeout time.Duration) *GetPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plan params
func (o *GetPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plan params
func (o *GetPlanParams) WithContext(ctx context.Context) *GetPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plan params
func (o *GetPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plan params
func (o *GetPlanParams) WithHTTPClient(client *http.Client) *GetPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plan params
func (o *GetPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlanID adds the planID to the get plan params
func (o *GetPlanParams) WithPlanID(planID string) *GetPlanParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the get plan params
func (o *GetPlanParams) SetPlanID(planID string) {
	o.PlanID = planID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param plan_id
	if err := r.SetPathParam("plan_id", o.PlanID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
