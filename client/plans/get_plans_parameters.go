package plans

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

// NewGetPlansParams creates a new GetPlansParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlansParams() *GetPlansParams {
	return &GetPlansParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlansParamsWithTimeout creates a new GetPlansParams object
// with the ability to set a timeout on a request.
func NewGetPlansParamsWithTimeout(timeout time.Duration) *GetPlansParams {
	return &GetPlansParams{
		timeout: timeout,
	}
}

// NewGetPlansParamsWithContext creates a new GetPlansParams object
// with the ability to set a context for a request.
func NewGetPlansParamsWithContext(ctx context.Context) *GetPlansParams {
	return &GetPlansParams{
		Context: ctx,
	}
}

// NewGetPlansParamsWithHTTPClient creates a new GetPlansParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlansParamsWithHTTPClient(client *http.Client) *GetPlansParams {
	return &GetPlansParams{
		HTTPClient: client,
	}
}

/*
GetPlansParams contains all the parameters to send to the API endpoint

	for the get plans operation.

	Typically these are written to a http.Request.
*/
type GetPlansParams struct {
	FilterDiskEql    int64  `json:"disk_eql"`
	FilterDiskLte    int64  `json:"disk_lte"`
	FilterDiskGte    int64  `json:"disk_gte"`
	FilterGpu        bool   `json:"gpu"`
	FilterInStock    bool   `json:"in_stock"`
	FilterLocation   string `json:"location"`
	FilterName       string `json:"name"`
	FilterRAMEql     int64  `json:"ram_eql"`
	FilterRAMLte     int64  `json:"ram_lte"`
	FilterRAMGte     int64  `json:"ram_gte"`
	FilterSlug       string `json:"slug"`
	FilterStockLevel string `json:"stock_level"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get plans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlansParams) WithDefaults() *GetPlansParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get plans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlansParams) SetDefaults() {
	val := GetPlansParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get plans params
func (o *GetPlansParams) WithTimeout(timeout time.Duration) *GetPlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plans params
func (o *GetPlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plans params
func (o *GetPlansParams) WithContext(ctx context.Context) *GetPlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plans params
func (o *GetPlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plans params
func (o *GetPlansParams) WithHTTPClient(client *http.Client) *GetPlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plans params
func (o *GetPlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterDiskEql adds the filterDiskEql to the get plans params
func (o *GetPlansParams) WithFilterDiskEql(filterDiskEql int64) *GetPlansParams {
	o.SetFilterDiskEql(filterDiskEql)
	return o
}

// SetFilterDiskEql adds the filterDiskEql to the get plans params
func (o *GetPlansParams) SetFilterDiskEql(filterDiskEql int64) {
	o.FilterDiskEql = filterDiskEql
}

// WithFilterGpu adds the filterGpu to the get plans params
func (o *GetPlansParams) WithFilterGpu(filterGpu bool) *GetPlansParams {
	o.SetFilterGpu(filterGpu)
	return o
}

// SetFilterGpu adds the filterGpu to the get plans params
func (o *GetPlansParams) SetFilterGpu(filterGpu bool) {
	o.FilterGpu = filterGpu
}

// WithFilterInStock adds the filterInStock to the get plans params
func (o *GetPlansParams) WithFilterInStock(filterInStock bool) *GetPlansParams {
	o.SetFilterInStock(filterInStock)
	return o
}

// SetFilterInStock adds the filterInStock to the get plans params
func (o *GetPlansParams) SetFilterInStock(filterInStock bool) {
	o.FilterInStock = filterInStock
}

// WithFilterLocation adds the filterLocation to the get plans params
func (o *GetPlansParams) WithFilterLocation(filterLocation string) *GetPlansParams {
	o.SetFilterLocation(filterLocation)
	return o
}

// SetFilterLocation adds the filterLocation to the get plans params
func (o *GetPlansParams) SetFilterLocation(filterLocation string) {
	o.FilterLocation = filterLocation
}

// WithFilterName adds the filterName to the get plans params
func (o *GetPlansParams) WithFilterName(filterName string) *GetPlansParams {
	o.SetFilterName(filterName)
	return o
}

// SetFilterName adds the filterName to the get plans params
func (o *GetPlansParams) SetFilterName(filterName string) {
	o.FilterName = filterName
}

// WithFilterRAMEql adds the filterRAMEql to the get plans params
func (o *GetPlansParams) WithFilterRAMEql(filterRAMEql int64) *GetPlansParams {
	o.SetFilterRAMEql(filterRAMEql)
	return o
}

// SetFilterRAMEql adds the filterRamEql to the get plans params
func (o *GetPlansParams) SetFilterRAMEql(filterRAMEql int64) {
	o.FilterRAMEql = filterRAMEql
}

// WithFilterRAMLte adds the filterRAMLte to the get plans params
func (o *GetPlansParams) WithFilterRAMLte(filterRAMLte int64) *GetPlansParams {
	o.SetFilterRAMLte(filterRAMLte)
	return o
}

// SetFilterRAMLte adds the filterRamLte to the get plans params
func (o *GetPlansParams) SetFilterRAMLte(filterRAMLte int64) {
	o.FilterRAMLte = filterRAMLte
}

// WithFilterRAMGte adds the filterRAMGte to the get plans params
func (o *GetPlansParams) WithFilterRAMGte(filterRAMGte int64) *GetPlansParams {
	o.SetFilterRAMGte(filterRAMGte)
	return o
}

// SetFilterRAMGte adds the filterRamGte to the get plans params
func (o *GetPlansParams) SetFilterRAMGte(filterRAMGte int64) {
	o.FilterRAMGte = filterRAMGte
}

// WithFilterSlug adds the filterSlug to the get plans params
func (o *GetPlansParams) WithFilterSlug(filterSlug string) *GetPlansParams {
	o.SetFilterSlug(filterSlug)
	return o
}

// SetFilterSlug adds the filterSlug to the get plans params
func (o *GetPlansParams) SetFilterSlug(filterSlug string) {
	o.FilterSlug = filterSlug
}

// WithFilterStockLevel adds the filterStockLevel to the get plans params
func (o *GetPlansParams) WithFilterStockLevel(filterStockLevel string) *GetPlansParams {
	o.SetFilterStockLevel(filterStockLevel)
	return o
}

// SetFilterStockLevel adds the filterStockLevel to the get plans params
func (o *GetPlansParams) SetFilterStockLevel(filterStockLevel string) {
	o.FilterStockLevel = filterStockLevel
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetQueryParam("page[size]", "100"); err != nil {
		return err
	}

	if !swag.IsZero(o.FilterDiskEql) {
		if err := r.SetQueryParam("filter[disk][eql]", swag.FormatInt64(o.FilterDiskEql)); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterDiskLte) {
		if err := r.SetQueryParam("filter[disk][lte]", swag.FormatInt64(o.FilterDiskLte)); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterDiskGte) {
		if err := r.SetQueryParam("filter[disk][gte]", swag.FormatInt64(o.FilterDiskGte)); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterGpu) {
		if err := r.SetQueryParam("filter[gpu]", swag.FormatBool(o.FilterGpu)); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterInStock) {
		if err := r.SetQueryParam("filter[in_stock]", swag.FormatBool(o.FilterInStock)); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterLocation) {
		if err := r.SetQueryParam("filter[location]", o.FilterLocation); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterName) {
		if err := r.SetQueryParam("filter[name]", o.FilterName); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterRAMEql) {
		if err := r.SetQueryParam("filter[ram][eql]", swag.FormatInt64(o.FilterRAMEql)); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterRAMLte) {
		if err := r.SetQueryParam("filter[ram][lte]", swag.FormatInt64(o.FilterRAMLte)); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterRAMGte) {
		if err := r.SetQueryParam("filter[ram][gte]", swag.FormatInt64(o.FilterRAMGte)); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterSlug) {
		if err := r.SetQueryParam("filter[slug]", o.FilterSlug); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterStockLevel) {
		if err := r.SetQueryParam("filter[stock_level]", o.FilterStockLevel); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
