package projects

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

// NewGetProjectsParams creates a new GetProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectsParams() *GetProjectsParams {
	return &GetProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsParamsWithTimeout creates a new GetProjectsParams object
// with the ability to set a timeout on a request.
func NewGetProjectsParamsWithTimeout(timeout time.Duration) *GetProjectsParams {
	return &GetProjectsParams{
		timeout: timeout,
	}
}

// NewGetProjectsParamsWithContext creates a new GetProjectsParams object
// with the ability to set a context for a request.
func NewGetProjectsParamsWithContext(ctx context.Context) *GetProjectsParams {
	return &GetProjectsParams{
		Context: ctx,
	}
}

// NewGetProjectsParamsWithHTTPClient creates a new GetProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectsParamsWithHTTPClient(client *http.Client) *GetProjectsParams {
	return &GetProjectsParams{
		HTTPClient: client,
	}
}

type GetProjectsParams struct {
	ExtraFieldsProjects string `json:"extra_fields"`
	FilterBillingType   string `json:"billing_type"`
	FilterDescription   string `json:"description"`
	FilterEnvironment   string `json:"environment"`
	FilterName          string `json:"name"`
	FilterSlug          string `json:"slug"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsParams) WithDefaults() *GetProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsParams) SetDefaults() {
	val := GetProjectsParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get projects params
func (o *GetProjectsParams) WithTimeout(timeout time.Duration) *GetProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects params
func (o *GetProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects params
func (o *GetProjectsParams) WithContext(ctx context.Context) *GetProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects params
func (o *GetProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects params
func (o *GetProjectsParams) WithHTTPClient(client *http.Client) *GetProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects params
func (o *GetProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtraFieldsProjects adds the extraFieldsProjects to the get projects params
func (o *GetProjectsParams) WithExtraFieldsProjects(extraFieldsProjects string) *GetProjectsParams {
	o.SetExtraFieldsProjects(extraFieldsProjects)
	return o
}

// SetExtraFieldsProjects adds the extraFieldsProjects to the get projects params
func (o *GetProjectsParams) SetExtraFieldsProjects(extraFieldsProjects string) {
	o.ExtraFieldsProjects = extraFieldsProjects
}

// WithFilterBillingType adds the filterBillingType to the get projects params
func (o *GetProjectsParams) WithFilterBillingType(filterBillingType string) *GetProjectsParams {
	o.SetFilterBillingType(filterBillingType)
	return o
}

// SetFilterBillingType adds the filterBillingType to the get projects params
func (o *GetProjectsParams) SetFilterBillingType(filterBillingType string) {
	o.FilterBillingType = filterBillingType
}

// WithFilterDescription adds the filterDescription to the get projects params
func (o *GetProjectsParams) WithFilterDescription(filterDescription string) *GetProjectsParams {
	o.SetFilterDescription(filterDescription)
	return o
}

// SetFilterDescription adds the filterDescription to the get projects params
func (o *GetProjectsParams) SetFilterDescription(filterDescription string) {
	o.FilterDescription = filterDescription
}

// WithFilterEnvironment adds the filterEnvironment to the get projects params
func (o *GetProjectsParams) WithFilterEnvironment(filterEnvironment string) *GetProjectsParams {
	o.SetFilterEnvironment(filterEnvironment)
	return o
}

// SetFilterEnvironment adds the filterEnvironment to the get projects params
func (o *GetProjectsParams) SetFilterEnvironment(filterEnvironment string) {
	o.FilterEnvironment = filterEnvironment
}

// WithFilterName adds the filterName to the get projects params
func (o *GetProjectsParams) WithFilterName(filterName string) *GetProjectsParams {
	o.SetFilterName(filterName)
	return o
}

// SetFilterName adds the filterName to the get projects params
func (o *GetProjectsParams) SetFilterName(filterName string) {
	o.FilterName = filterName
}

// WithFilterSlug adds the filterSlug to the get projects params
func (o *GetProjectsParams) WithFilterSlug(filterSlug string) *GetProjectsParams {
	o.SetFilterSlug(filterSlug)
	return o
}

// SetFilterSlug adds the filterSlug to the get projects params
func (o *GetProjectsParams) SetFilterSlug(filterSlug string) {
	o.FilterSlug = filterSlug
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetQueryParam("page[size]", "100"); err != nil {
		return err
	}

	if !swag.IsZero(o.ExtraFieldsProjects) {
		if err := r.SetQueryParam("extra_fields[projects]", o.ExtraFieldsProjects); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterBillingType) {
		if err := r.SetQueryParam("filter[billing_type]", o.FilterBillingType); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterDescription) {
		if err := r.SetQueryParam("filter[description]", o.FilterDescription); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterEnvironment) {
		if err := r.SetQueryParam("filter[environment]", o.FilterEnvironment); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterName) {
		if err := r.SetQueryParam("filter[name]", o.FilterName); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterSlug) {
		if err := r.SetQueryParam("filter[slug]", o.FilterSlug); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
