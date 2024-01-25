// Code generated by go-swagger; DO NOT EDIT.

package plans

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

	// APIVersion.
	//
	// Default: "2023-06-01"
	APIVersion *string

	/* FilterDisk.

	   The disk size in Gigabytes to filter by, should be used with the following options:
	                            [eql] to filter for values equal to the provided value.
	                            [gte] to filter for values greater or equal to the provided value.
	                            [lte] to filter by values lower or equal to the provided value.
	*/
	FilterDiskEql *int64
	FilterDiskLte *int64
	FilterDiskGte *int64

	/* FilterGpu.

	   Filter by the existence of an associated GPU
	*/
	FilterGpu *bool

	/* FilterInStock.

	   The stock available at the site to filter by
	*/
	FilterInStock *bool

	/* FilterLocation.

	   The location of the site to filter by
	*/
	FilterLocation *string

	/* FilterName.

	   The plan name to filter by
	*/
	FilterName *string

	/* FilterRAM.

	   The ram size in Gigabytes to filter by, should be used with the following options:
	                            [eql] to filter for values equal to the provided value.
	                            [gte] to filter for values greater or equal to the provided value.
	                            [lte] to filter by values lower or equal to the provided value.
	*/
	FilterRAM *int64

	/* FilterSlug.

	   The plan slug to filter by
	*/
	FilterSlug *string

	/* FilterStockLevel.

	   The stock level at the site to filter by
	*/
	FilterStockLevel *string

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
	var (
		aPIVersionDefault = string("2023-06-01")
	)

	val := GetPlansParams{
		APIVersion: &aPIVersionDefault,
	}

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

// WithAPIVersion adds the aPIVersion to the get plans params
func (o *GetPlansParams) WithAPIVersion(aPIVersion *string) *GetPlansParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get plans params
func (o *GetPlansParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithFilterDiskEql adds the filterDiskEql to the get plans params
func (o *GetPlansParams) WithFilterDiskEql(filterDiskEql *int64) *GetPlansParams {
	o.SetFilterDiskEql(filterDiskEql)
	return o
}

// SetFilterDiskEql adds the filterDiskEql to the get plans params
func (o *GetPlansParams) SetFilterDiskEql(filterDiskEql *int64) {
	o.FilterDiskEql = filterDiskEql
}

// WithFilterGpu adds the filterGpu to the get plans params
func (o *GetPlansParams) WithFilterGpu(filterGpu *bool) *GetPlansParams {
	o.SetFilterGpu(filterGpu)
	return o
}

// SetFilterGpu adds the filterGpu to the get plans params
func (o *GetPlansParams) SetFilterGpu(filterGpu *bool) {
	o.FilterGpu = filterGpu
}

// WithFilterInStock adds the filterInStock to the get plans params
func (o *GetPlansParams) WithFilterInStock(filterInStock *bool) *GetPlansParams {
	o.SetFilterInStock(filterInStock)
	return o
}

// SetFilterInStock adds the filterInStock to the get plans params
func (o *GetPlansParams) SetFilterInStock(filterInStock *bool) {
	o.FilterInStock = filterInStock
}

// WithFilterLocation adds the filterLocation to the get plans params
func (o *GetPlansParams) WithFilterLocation(filterLocation *string) *GetPlansParams {
	o.SetFilterLocation(filterLocation)
	return o
}

// SetFilterLocation adds the filterLocation to the get plans params
func (o *GetPlansParams) SetFilterLocation(filterLocation *string) {
	o.FilterLocation = filterLocation
}

// WithFilterName adds the filterName to the get plans params
func (o *GetPlansParams) WithFilterName(filterName *string) *GetPlansParams {
	o.SetFilterName(filterName)
	return o
}

// SetFilterName adds the filterName to the get plans params
func (o *GetPlansParams) SetFilterName(filterName *string) {
	o.FilterName = filterName
}

// WithFilterRAM adds the filterRAM to the get plans params
func (o *GetPlansParams) WithFilterRAM(filterRAM *int64) *GetPlansParams {
	o.SetFilterRAM(filterRAM)
	return o
}

// SetFilterRAM adds the filterRam to the get plans params
func (o *GetPlansParams) SetFilterRAM(filterRAM *int64) {
	o.FilterRAM = filterRAM
}

// WithFilterSlug adds the filterSlug to the get plans params
func (o *GetPlansParams) WithFilterSlug(filterSlug *string) *GetPlansParams {
	o.SetFilterSlug(filterSlug)
	return o
}

// SetFilterSlug adds the filterSlug to the get plans params
func (o *GetPlansParams) SetFilterSlug(filterSlug *string) {
	o.FilterSlug = filterSlug
}

// WithFilterStockLevel adds the filterStockLevel to the get plans params
func (o *GetPlansParams) WithFilterStockLevel(filterStockLevel *string) *GetPlansParams {
	o.SetFilterStockLevel(filterStockLevel)
	return o
}

// SetFilterStockLevel adds the filterStockLevel to the get plans params
func (o *GetPlansParams) SetFilterStockLevel(filterStockLevel *string) {
	o.FilterStockLevel = filterStockLevel
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// header param API-Version
		if err := r.SetHeaderParam("API-Version", *o.APIVersion); err != nil {
			return err
		}
	}

	if o.FilterDiskEql != nil {

		// query param filter[disk][eql]
		var qrFilterDiskEql int64

		if o.FilterDiskEql != nil {
			qrFilterDiskEql = *o.FilterDiskEql
		}
		qFilterDiskEql := swag.FormatInt64(qrFilterDiskEql)
		if qFilterDiskEql != "" {

			if err := r.SetQueryParam("filter[disk][eql]", qFilterDiskEql); err != nil {
				return err
			}
		}
	}

	if o.FilterDiskLte != nil {

		// query param filter[disk][lte]
		var qrFilterDiskLte int64

		if o.FilterDiskLte != nil {
			qrFilterDiskLte = *o.FilterDiskLte
		}
		qFilterDiskLte := swag.FormatInt64(qrFilterDiskLte)
		if qFilterDiskLte != "" {

			if err := r.SetQueryParam("filter[disk][lte]", qFilterDiskLte); err != nil {
				return err
			}
		}
	}

	if o.FilterDiskGte != nil {

		// query param filter[disk][gte]
		var qrFilterDiskGte int64

		if o.FilterDiskGte != nil {
			qrFilterDiskGte = *o.FilterDiskGte
		}
		qFilterDiskGte := swag.FormatInt64(qrFilterDiskGte)
		if qFilterDiskGte != "" {

			if err := r.SetQueryParam("filter[disk][gte]", qFilterDiskGte); err != nil {
				return err
			}
		}
	}

	if o.FilterGpu != nil {

		// query param filter[gpu]
		var qrFilterGpu bool

		if o.FilterGpu != nil {
			qrFilterGpu = *o.FilterGpu
		}
		qFilterGpu := swag.FormatBool(qrFilterGpu)
		if qFilterGpu != "" {

			if err := r.SetQueryParam("filter[gpu]", qFilterGpu); err != nil {
				return err
			}
		}
	}

	if o.FilterInStock != nil {

		// query param filter[in_stock]
		var qrFilterInStock bool

		if o.FilterInStock != nil {
			qrFilterInStock = *o.FilterInStock
		}
		qFilterInStock := swag.FormatBool(qrFilterInStock)
		if qFilterInStock != "" {

			if err := r.SetQueryParam("filter[in_stock]", qFilterInStock); err != nil {
				return err
			}
		}
	}

	if o.FilterLocation != nil {

		// query param filter[location]
		var qrFilterLocation string

		if o.FilterLocation != nil {
			qrFilterLocation = *o.FilterLocation
		}
		qFilterLocation := qrFilterLocation
		if qFilterLocation != "" {

			if err := r.SetQueryParam("filter[location]", qFilterLocation); err != nil {
				return err
			}
		}
	}

	if o.FilterName != nil {

		// query param filter[name]
		var qrFilterName string

		if o.FilterName != nil {
			qrFilterName = *o.FilterName
		}
		qFilterName := qrFilterName
		if qFilterName != "" {

			if err := r.SetQueryParam("filter[name]", qFilterName); err != nil {
				return err
			}
		}
	}

	if o.FilterRAM != nil {

		// query param filter[ram]
		var qrFilterRAM int64

		if o.FilterRAM != nil {
			qrFilterRAM = *o.FilterRAM
		}
		qFilterRAM := swag.FormatInt64(qrFilterRAM)
		if qFilterRAM != "" {

			if err := r.SetQueryParam("filter[ram]", qFilterRAM); err != nil {
				return err
			}
		}
	}

	if o.FilterSlug != nil {

		// query param filter[slug]
		var qrFilterSlug string

		if o.FilterSlug != nil {
			qrFilterSlug = *o.FilterSlug
		}
		qFilterSlug := qrFilterSlug
		if qFilterSlug != "" {

			if err := r.SetQueryParam("filter[slug]", qFilterSlug); err != nil {
				return err
			}
		}
	}

	if o.FilterStockLevel != nil {

		// query param filter[stock_level]
		var qrFilterStockLevel string

		if o.FilterStockLevel != nil {
			qrFilterStockLevel = *o.FilterStockLevel
		}
		qFilterStockLevel := qrFilterStockLevel
		if qFilterStockLevel != "" {

			if err := r.SetQueryParam("filter[stock_level]", qFilterStockLevel); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
