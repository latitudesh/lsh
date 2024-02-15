package servers

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

// NewGetServersParams creates a new GetServersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServersParams() *GetServersParams {
	return &GetServersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServersParamsWithTimeout creates a new GetServersParams object
// with the ability to set a timeout on a request.
func NewGetServersParamsWithTimeout(timeout time.Duration) *GetServersParams {
	return &GetServersParams{
		timeout: timeout,
	}
}

// NewGetServersParamsWithContext creates a new GetServersParams object
// with the ability to set a context for a request.
func NewGetServersParamsWithContext(ctx context.Context) *GetServersParams {
	return &GetServersParams{
		Context: ctx,
	}
}

// NewGetServersParamsWithHTTPClient creates a new GetServersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServersParamsWithHTTPClient(client *http.Client) *GetServersParams {
	return &GetServersParams{
		HTTPClient: client,
	}
}

/*
GetServersParams contains all the parameters to send to the API endpoint

	for the get servers operation.

	Typically these are written to a http.Request.
*/
type GetServersParams struct {
	ExtraFieldsServers    string `json:"extra_fields"`
	FilterCreatedAtGte    string `json:"created_at_gte"`
	FilterCreatedAtLte    string `json:"created_at_lte"`
	FilterDiskEql         int64  `json:"disk_eql"`
	FilterDiskLte         int64  `json:"disk_lte"`
	FilterDiskGte         int64  `json:"disk_gte"`
	FilterGpu             bool   `json:"gpu"`
	FilterHostname        string `json:"hostname"`
	FilterLabel           string `json:"label"`
	FilterOperatingSystem string `json:"operating_system"`
	FilterPlan            string `json:"plan"`
	FilterProject         string `json:"project"`
	FilterRAMEql          int64  `json:"ram_eql"`
	FilterRAMGte          int64  `json:"ram_gte"`
	FilterRAMLte          int64  `json:"ram_lte"`
	FilterRegion          string `json:"region"`
	FilterStatus          string `json:"status"`

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get servers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServersParams) WithDefaults() *GetServersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get servers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServersParams) SetDefaults() {
	val := GetServersParams{}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get servers params
func (o *GetServersParams) WithTimeout(timeout time.Duration) *GetServersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get servers params
func (o *GetServersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get servers params
func (o *GetServersParams) WithContext(ctx context.Context) *GetServersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get servers params
func (o *GetServersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get servers params
func (o *GetServersParams) WithHTTPClient(client *http.Client) *GetServersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get servers params
func (o *GetServersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtraFieldsServers adds the extraFieldsServers to the get servers params
func (o *GetServersParams) WithExtraFieldsServers(extraFieldsServers string) *GetServersParams {
	o.SetExtraFieldsServers(extraFieldsServers)
	return o
}

// SetExtraFieldsServers adds the extraFieldsServers to the get servers params
func (o *GetServersParams) SetExtraFieldsServers(extraFieldsServers string) {
	o.ExtraFieldsServers = extraFieldsServers
}

// WithFilterCreatedAtGte adds the filterCreatedAtGte to the get servers params
func (o *GetServersParams) WithFilterCreatedAtGte(filterCreatedAtGte string) *GetServersParams {
	o.SetFilterCreatedAtGte(filterCreatedAtGte)
	return o
}

// SetFilterCreatedAtGte adds the filterCreatedAtGte to the get servers params
func (o *GetServersParams) SetFilterCreatedAtGte(filterCreatedAtGte string) {
	o.FilterCreatedAtGte = filterCreatedAtGte
}

// WithFilterCreatedAtLte adds the filterCreatedAtLte to the get servers params
func (o *GetServersParams) WithFilterCreatedAtLte(filterCreatedAtLte string) *GetServersParams {
	o.SetFilterCreatedAtLte(filterCreatedAtLte)
	return o
}

// SetFilterCreatedAtLte adds the filterCreatedAtLte to the get servers params
func (o *GetServersParams) SetFilterCreatedAtLte(filterCreatedAtLte string) {
	o.FilterCreatedAtLte = filterCreatedAtLte
}

// WithFilterDiskEql adds the filterDiskEql to the get servers params
func (o *GetServersParams) WithFilterDiskEql(filterDiskEql int64) *GetServersParams {
	o.SetFilterDiskEql(filterDiskEql)
	return o
}

// WithFilterDiskLte adds the filterDiskLte to the get servers params
func (o *GetServersParams) WithFilterDiskLte(filterDiskLte int64) *GetServersParams {
	o.SetFilterDiskLte(filterDiskLte)
	return o
}

// WithFilterDiskGte adds the filterDiskGte to the get servers params
func (o *GetServersParams) WithFilterDiskGte(filterDiskGte int64) *GetServersParams {
	o.SetFilterDiskGte(filterDiskGte)
	return o
}

// SetFilterDiskEql adds the filterDiskEql to the get servers params
func (o *GetServersParams) SetFilterDiskEql(filterDiskEql int64) {
	o.FilterDiskEql = filterDiskEql
}

// SetFilterDiskLte adds the filterDiskLte to the get servers params
func (o *GetServersParams) SetFilterDiskLte(filterDiskLte int64) {
	o.FilterDiskLte = filterDiskLte
}

// SetFilterDiskGte adds the filterDiskGte to the get servers params
func (o *GetServersParams) SetFilterDiskGte(filterDiskGte int64) {
	o.FilterDiskGte = filterDiskGte
}

// WithFilterGpu adds the filterGpu to the get servers params
func (o *GetServersParams) WithFilterGpu(filterGpu bool) *GetServersParams {
	o.SetFilterGpu(filterGpu)
	return o
}

// SetFilterGpu adds the filterGpu to the get servers params
func (o *GetServersParams) SetFilterGpu(filterGpu bool) {
	o.FilterGpu = filterGpu
}

// WithFilterHostname adds the filterHostname to the get servers params
func (o *GetServersParams) WithFilterHostname(filterHostname string) *GetServersParams {
	o.SetFilterHostname(filterHostname)
	return o
}

// SetFilterHostname adds the filterHostname to the get servers params
func (o *GetServersParams) SetFilterHostname(filterHostname string) {
	o.FilterHostname = filterHostname
}

// WithFilterLabel adds the filterLabel to the get servers params
func (o *GetServersParams) WithFilterLabel(filterLabel string) *GetServersParams {
	o.SetFilterLabel(filterLabel)
	return o
}

// SetFilterLabel adds the filterLabel to the get servers params
func (o *GetServersParams) SetFilterLabel(filterLabel string) {
	o.FilterLabel = filterLabel
}

// WithFilterOperatingSystem adds the filterOperatingSystem to the get servers params
func (o *GetServersParams) WithFilterOperatingSystem(filterOperatingSystem string) *GetServersParams {
	o.SetFilterOperatingSystem(filterOperatingSystem)
	return o
}

// SetFilterOperatingSystem adds the filterOperatingSystem to the get servers params
func (o *GetServersParams) SetFilterOperatingSystem(filterOperatingSystem string) {
	o.FilterOperatingSystem = filterOperatingSystem
}

// WithFilterPlan adds the filterPlan to the get servers params
func (o *GetServersParams) WithFilterPlan(filterPlan string) *GetServersParams {
	o.SetFilterPlan(filterPlan)
	return o
}

// SetFilterPlan adds the filterPlan to the get servers params
func (o *GetServersParams) SetFilterPlan(filterPlan string) {
	o.FilterPlan = filterPlan
}

// WithFilterProject adds the filterProject to the get servers params
func (o *GetServersParams) WithFilterProject(filterProject string) *GetServersParams {
	o.SetFilterProject(filterProject)
	return o
}

// SetFilterProject adds the filterProject to the get servers params
func (o *GetServersParams) SetFilterProject(filterProject string) {
	o.FilterProject = filterProject
}

// WithFilterRAMEql adds the filterRAMEql to the get servers params
func (o *GetServersParams) WithFilterRAMEql(filterRAMEql int64) *GetServersParams {
	o.SetFilterRAMEql(filterRAMEql)
	return o
}

// SetFilterRAMEql adds the filterRamEql to the get servers params
func (o *GetServersParams) SetFilterRAMEql(filterRAMEql int64) {
	o.FilterRAMEql = filterRAMEql
}

// WithFilterRAMGte adds the filterRAMGte to the get servers params
func (o *GetServersParams) WithFilterRAMGte(filterRAMGte int64) *GetServersParams {
	o.SetFilterRAMGte(filterRAMGte)
	return o
}

// SetFilterRAMGte adds the filterRamGte to the get servers params
func (o *GetServersParams) SetFilterRAMGte(filterRAMGte int64) {
	o.FilterRAMGte = filterRAMGte
}

// WithFilterRAMLte adds the filterRAMLte to the get servers params
func (o *GetServersParams) WithFilterRAMLte(filterRAMLte int64) *GetServersParams {
	o.SetFilterRAMLte(filterRAMLte)
	return o
}

// SetFilterRAMLte adds the filterRamLte to the get servers params
func (o *GetServersParams) SetFilterRAMLte(filterRAMLte int64) {
	o.FilterRAMLte = filterRAMLte
}

// WithFilterRegion adds the filterRegion to the get servers params
func (o *GetServersParams) WithFilterRegion(filterRegion string) *GetServersParams {
	o.SetFilterRegion(filterRegion)
	return o
}

// SetFilterRegion adds the filterRegion to the get servers params
func (o *GetServersParams) SetFilterRegion(filterRegion string) {
	o.FilterRegion = filterRegion
}

// WithFilterStatus adds the filterStatus to the get servers params
func (o *GetServersParams) WithFilterStatus(filterStatus string) *GetServersParams {
	o.SetFilterStatus(filterStatus)
	return o
}

// SetFilterStatus adds the filterStatus to the get servers params
func (o *GetServersParams) SetFilterStatus(filterStatus string) {
	o.FilterStatus = filterStatus
}

// WriteToRequest writes these params to a swagger request
func (o *GetServersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetQueryParam("page[size]", "100"); err != nil {
		return err
	}

	if !swag.IsZero(o.ExtraFieldsServers) {
		if err := r.SetQueryParam("extra_fields[servers]", o.ExtraFieldsServers); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterCreatedAtGte) {
		if err := r.SetQueryParam("filter[created_at_gte]", o.FilterCreatedAtGte); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterCreatedAtLte) {
		if err := r.SetQueryParam("filter[created_at_lte]", o.FilterCreatedAtLte); err != nil {
			return err
		}
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

	if !swag.IsZero(o.FilterHostname) {
		if err := r.SetQueryParam("filter[hostname]", o.FilterHostname); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterLabel) {
		if err := r.SetQueryParam("filter[label]", o.FilterLabel); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterPlan) {
		if err := r.SetQueryParam("filter[plan]", o.FilterPlan); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterProject) {
		if err := r.SetQueryParam("filter[project]", o.FilterProject); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterRAMEql) {
		if err := r.SetQueryParam("filter[ram][eql]", swag.FormatInt64(o.FilterRAMEql)); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterRAMGte) {
		if err := r.SetQueryParam("filter[ram][gte]", swag.FormatInt64(o.FilterRAMGte)); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterRAMLte) {
		if err := r.SetQueryParam("filter[ram][lte]", swag.FormatInt64(o.FilterRAMLte)); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterRegion) {
		if err := r.SetQueryParam("filter[region]", o.FilterRegion); err != nil {
			return err
		}
	}

	if !swag.IsZero(o.FilterStatus) {
		if err := r.SetQueryParam("filter[status]", o.FilterStatus); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
