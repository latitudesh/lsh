// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/models"

	"github.com/spf13/cobra"
)

// Schema cli for TrafficQuota

// register flags to command
func registerModelTrafficQuotaFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficQuotaData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficQuotaData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelTrafficQuotaDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficQuotaFlags(depth int, m *models.TrafficQuota, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveTrafficQuotaDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveTrafficQuotaDataFlags(depth int, m *models.TrafficQuota, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data TrafficQuotaData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.TrafficQuotaData{}
	}

	err, dataAdded := retrieveModelTrafficQuotaDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for TrafficQuotaData

// register flags to command
func registerModelTrafficQuotaDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficQuotaDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficQuotaDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelTrafficQuotaDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficQuotaDataID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerTrafficQuotaDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["traffic_quota"]. `

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["traffic_quota"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficQuotaDataFlags(depth int, m *models.TrafficQuotaData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveTrafficQuotaDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveTrafficQuotaDataIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, typeAdded := retrieveTrafficQuotaDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesFlags(depth int, m *models.TrafficQuotaData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes TrafficQuotaDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.TrafficQuotaDataAttributes{}
	}

	err, attributesAdded := retrieveModelTrafficQuotaDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataIDFlags(depth int, m *models.TrafficQuotaData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataTypeFlags(depth int, m *models.TrafficQuotaData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for TrafficQuotaDataAttributes

// register flags to command
func registerModelTrafficQuotaDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficQuotaDataAttributesQuotaPerProject(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProject(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: quota_per_project []*TrafficQuotaDataAttributesQuotaPerProjectItems0 array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficQuotaDataAttributesFlags(depth int, m *models.TrafficQuotaDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, quotaPerProjectAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || quotaPerProjectAdded

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectFlags(depth int, m *models.TrafficQuotaDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	quotaPerProjectFlagName := fmt.Sprintf("%v.quota_per_project", cmdPrefix)
	if cmd.Flags().Changed(quotaPerProjectFlagName) {
		// warning: quota_per_project array type []*TrafficQuotaDataAttributesQuotaPerProjectItems0 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

// Extra schema cli for TrafficQuotaDataAttributesQuotaPerProjectItems0

// register flags to command
func registerModelTrafficQuotaDataAttributesQuotaPerProjectItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0BillingMethod(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0ProjectID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0ProjectSlug(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0BillingMethod(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	billingMethodDescription := ``

	var billingMethodFlagName string
	if cmdPrefix == "" {
		billingMethodFlagName = "billing_method"
	} else {
		billingMethodFlagName = fmt.Sprintf("%v.billing_method", cmdPrefix)
	}

	var billingMethodFlagDefault string

	_ = cmd.PersistentFlags().String(billingMethodFlagName, billingMethodFlagDefault, billingMethodDescription)

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0ProjectID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	projectIdDescription := ``

	var projectIdFlagName string
	if cmdPrefix == "" {
		projectIdFlagName = "project_id"
	} else {
		projectIdFlagName = fmt.Sprintf("%v.project_id", cmdPrefix)
	}

	var projectIdFlagDefault string

	_ = cmd.PersistentFlags().String(projectIdFlagName, projectIdFlagDefault, projectIdDescription)

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0ProjectSlug(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	projectSlugDescription := ``

	var projectSlugFlagName string
	if cmdPrefix == "" {
		projectSlugFlagName = "project_slug"
	} else {
		projectSlugFlagName = fmt.Sprintf("%v.project_slug", cmdPrefix)
	}

	var projectSlugFlagDefault string

	_ = cmd.PersistentFlags().String(projectSlugFlagName, projectSlugFlagDefault, projectSlugDescription)

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: quota_per_region []*TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0 array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficQuotaDataAttributesQuotaPerProjectItems0Flags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, billingMethodAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0BillingMethodFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || billingMethodAdded

	err, projectIdAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0ProjectIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || projectIdAdded

	err, projectSlugAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0ProjectSlugFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || projectSlugAdded

	err, quotaPerRegionAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || quotaPerRegionAdded

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0BillingMethodFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	billingMethodFlagName := fmt.Sprintf("%v.billing_method", cmdPrefix)
	if cmd.Flags().Changed(billingMethodFlagName) {

		var billingMethodFlagName string
		if cmdPrefix == "" {
			billingMethodFlagName = "billing_method"
		} else {
			billingMethodFlagName = fmt.Sprintf("%v.billing_method", cmdPrefix)
		}

		billingMethodFlagValue, err := cmd.Flags().GetString(billingMethodFlagName)
		if err != nil {
			return err, false
		}
		m.BillingMethod = billingMethodFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0ProjectIDFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	projectIdFlagName := fmt.Sprintf("%v.project_id", cmdPrefix)
	if cmd.Flags().Changed(projectIdFlagName) {

		var projectIdFlagName string
		if cmdPrefix == "" {
			projectIdFlagName = "project_id"
		} else {
			projectIdFlagName = fmt.Sprintf("%v.project_id", cmdPrefix)
		}

		projectIdFlagValue, err := cmd.Flags().GetString(projectIdFlagName)
		if err != nil {
			return err, false
		}
		m.ProjectID = projectIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0ProjectSlugFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	projectSlugFlagName := fmt.Sprintf("%v.project_slug", cmdPrefix)
	if cmd.Flags().Changed(projectSlugFlagName) {

		var projectSlugFlagName string
		if cmdPrefix == "" {
			projectSlugFlagName = "project_slug"
		} else {
			projectSlugFlagName = fmt.Sprintf("%v.project_slug", cmdPrefix)
		}

		projectSlugFlagValue, err := cmd.Flags().GetString(projectSlugFlagName)
		if err != nil {
			return err, false
		}
		m.ProjectSlug = projectSlugFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	quotaPerRegionFlagName := fmt.Sprintf("%v.quota_per_region", cmdPrefix)
	if cmd.Flags().Changed(quotaPerRegionFlagName) {
		// warning: quota_per_region array type []*TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

// Extra schema cli for TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0

// register flags to command
func registerModelTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTb(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0RegionID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0RegionSlug(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var quotaInMbpsFlagName string
	if cmdPrefix == "" {
		quotaInMbpsFlagName = "quota_in_mbps"
	} else {
		quotaInMbpsFlagName = fmt.Sprintf("%v.quota_in_mbps", cmdPrefix)
	}

	if err := registerModelTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsFlags(depth+1, quotaInMbpsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTb(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var quotaInTbFlagName string
	if cmdPrefix == "" {
		quotaInTbFlagName = "quota_in_tb"
	} else {
		quotaInTbFlagName = fmt.Sprintf("%v.quota_in_tb", cmdPrefix)
	}

	if err := registerModelTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbFlags(depth+1, quotaInTbFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0RegionID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	regionIdDescription := ``

	var regionIdFlagName string
	if cmdPrefix == "" {
		regionIdFlagName = "region_id"
	} else {
		regionIdFlagName = fmt.Sprintf("%v.region_id", cmdPrefix)
	}

	var regionIdFlagDefault string

	_ = cmd.PersistentFlags().String(regionIdFlagName, regionIdFlagDefault, regionIdDescription)

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0RegionSlug(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	regionSlugDescription := ``

	var regionSlugFlagName string
	if cmdPrefix == "" {
		regionSlugFlagName = "region_slug"
	} else {
		regionSlugFlagName = fmt.Sprintf("%v.region_slug", cmdPrefix)
	}

	var regionSlugFlagDefault string

	_ = cmd.PersistentFlags().String(regionSlugFlagName, regionSlugFlagDefault, regionSlugDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0Flags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, quotaInMbpsAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || quotaInMbpsAdded

	err, quotaInTbAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || quotaInTbAdded

	err, regionIdAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0RegionIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regionIdAdded

	err, regionSlugAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0RegionSlugFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regionSlugAdded

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	quotaInMbpsFlagName := fmt.Sprintf("%v.quota_in_mbps", cmdPrefix)
	if cmd.Flags().Changed(quotaInMbpsFlagName) {
		// info: complex object quota_in_mbps TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbps is retrieved outside this Changed() block
	}
	quotaInMbpsFlagValue := m.QuotaInMbps
	if swag.IsZero(quotaInMbpsFlagValue) {
		quotaInMbpsFlagValue = &models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbps{}
	}

	err, quotaInMbpsAdded := retrieveModelTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsFlags(depth+1, quotaInMbpsFlagValue, quotaInMbpsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || quotaInMbpsAdded
	if quotaInMbpsAdded {
		m.QuotaInMbps = quotaInMbpsFlagValue
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	quotaInTbFlagName := fmt.Sprintf("%v.quota_in_tb", cmdPrefix)
	if cmd.Flags().Changed(quotaInTbFlagName) {
		// info: complex object quota_in_tb TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTb is retrieved outside this Changed() block
	}
	quotaInTbFlagValue := m.QuotaInTb
	if swag.IsZero(quotaInTbFlagValue) {
		quotaInTbFlagValue = &models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTb{}
	}

	err, quotaInTbAdded := retrieveModelTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbFlags(depth+1, quotaInTbFlagValue, quotaInTbFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || quotaInTbAdded
	if quotaInTbAdded {
		m.QuotaInTb = quotaInTbFlagValue
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0RegionIDFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	regionIdFlagName := fmt.Sprintf("%v.region_id", cmdPrefix)
	if cmd.Flags().Changed(regionIdFlagName) {

		var regionIdFlagName string
		if cmdPrefix == "" {
			regionIdFlagName = "region_id"
		} else {
			regionIdFlagName = fmt.Sprintf("%v.region_id", cmdPrefix)
		}

		regionIdFlagValue, err := cmd.Flags().GetString(regionIdFlagName)
		if err != nil {
			return err, false
		}
		m.RegionID = regionIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0RegionSlugFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	regionSlugFlagName := fmt.Sprintf("%v.region_slug", cmdPrefix)
	if cmd.Flags().Changed(regionSlugFlagName) {

		var regionSlugFlagName string
		if cmdPrefix == "" {
			regionSlugFlagName = "region_slug"
		} else {
			regionSlugFlagName = fmt.Sprintf("%v.region_slug", cmdPrefix)
		}

		regionSlugFlagValue, err := cmd.Flags().GetString(regionSlugFlagName)
		if err != nil {
			return err, false
		}
		m.RegionSlug = regionSlugFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbps

// register flags to command
func registerModelTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsAdditional(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsGranted(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsTotal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsAdditional(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	additionalDescription := ``

	var additionalFlagName string
	if cmdPrefix == "" {
		additionalFlagName = "additional"
	} else {
		additionalFlagName = fmt.Sprintf("%v.additional", cmdPrefix)
	}

	var additionalFlagDefault int64

	_ = cmd.PersistentFlags().Int64(additionalFlagName, additionalFlagDefault, additionalDescription)

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsGranted(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	grantedDescription := ``

	var grantedFlagName string
	if cmdPrefix == "" {
		grantedFlagName = "granted"
	} else {
		grantedFlagName = fmt.Sprintf("%v.granted", cmdPrefix)
	}

	var grantedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(grantedFlagName, grantedFlagDefault, grantedDescription)

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsTotal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalDescription := ``

	var totalFlagName string
	if cmdPrefix == "" {
		totalFlagName = "total"
	} else {
		totalFlagName = fmt.Sprintf("%v.total", cmdPrefix)
	}

	var totalFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalFlagName, totalFlagDefault, totalDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbps, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, additionalAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsAdditionalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || additionalAdded

	err, grantedAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsGrantedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || grantedAdded

	err, totalAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsTotalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalAdded

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsAdditionalFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbps, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	additionalFlagName := fmt.Sprintf("%v.additional", cmdPrefix)
	if cmd.Flags().Changed(additionalFlagName) {

		var additionalFlagName string
		if cmdPrefix == "" {
			additionalFlagName = "additional"
		} else {
			additionalFlagName = fmt.Sprintf("%v.additional", cmdPrefix)
		}

		additionalFlagValue, err := cmd.Flags().GetInt64(additionalFlagName)
		if err != nil {
			return err, false
		}
		m.Additional = additionalFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsGrantedFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbps, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	grantedFlagName := fmt.Sprintf("%v.granted", cmdPrefix)
	if cmd.Flags().Changed(grantedFlagName) {

		var grantedFlagName string
		if cmdPrefix == "" {
			grantedFlagName = "granted"
		} else {
			grantedFlagName = fmt.Sprintf("%v.granted", cmdPrefix)
		}

		grantedFlagValue, err := cmd.Flags().GetInt64(grantedFlagName)
		if err != nil {
			return err, false
		}
		m.Granted = grantedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbpsTotalFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInMbps, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalFlagName := fmt.Sprintf("%v.total", cmdPrefix)
	if cmd.Flags().Changed(totalFlagName) {

		var totalFlagName string
		if cmdPrefix == "" {
			totalFlagName = "total"
		} else {
			totalFlagName = fmt.Sprintf("%v.total", cmdPrefix)
		}

		totalFlagValue, err := cmd.Flags().GetInt64(totalFlagName)
		if err != nil {
			return err, false
		}
		m.Total = totalFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTb

// register flags to command
func registerModelTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbAdditional(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbGranted(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbTotal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbAdditional(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	additionalDescription := ``

	var additionalFlagName string
	if cmdPrefix == "" {
		additionalFlagName = "additional"
	} else {
		additionalFlagName = fmt.Sprintf("%v.additional", cmdPrefix)
	}

	var additionalFlagDefault int64

	_ = cmd.PersistentFlags().Int64(additionalFlagName, additionalFlagDefault, additionalDescription)

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbGranted(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	grantedDescription := ``

	var grantedFlagName string
	if cmdPrefix == "" {
		grantedFlagName = "granted"
	} else {
		grantedFlagName = fmt.Sprintf("%v.granted", cmdPrefix)
	}

	var grantedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(grantedFlagName, grantedFlagDefault, grantedDescription)

	return nil
}

func registerTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbTotal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalDescription := ``

	var totalFlagName string
	if cmdPrefix == "" {
		totalFlagName = "total"
	} else {
		totalFlagName = fmt.Sprintf("%v.total", cmdPrefix)
	}

	var totalFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalFlagName, totalFlagDefault, totalDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTb, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, additionalAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbAdditionalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || additionalAdded

	err, grantedAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbGrantedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || grantedAdded

	err, totalAdded := retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbTotalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalAdded

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbAdditionalFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTb, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	additionalFlagName := fmt.Sprintf("%v.additional", cmdPrefix)
	if cmd.Flags().Changed(additionalFlagName) {

		var additionalFlagName string
		if cmdPrefix == "" {
			additionalFlagName = "additional"
		} else {
			additionalFlagName = fmt.Sprintf("%v.additional", cmdPrefix)
		}

		additionalFlagValue, err := cmd.Flags().GetInt64(additionalFlagName)
		if err != nil {
			return err, false
		}
		m.Additional = additionalFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbGrantedFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTb, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	grantedFlagName := fmt.Sprintf("%v.granted", cmdPrefix)
	if cmd.Flags().Changed(grantedFlagName) {

		var grantedFlagName string
		if cmdPrefix == "" {
			grantedFlagName = "granted"
		} else {
			grantedFlagName = fmt.Sprintf("%v.granted", cmdPrefix)
		}

		grantedFlagValue, err := cmd.Flags().GetInt64(grantedFlagName)
		if err != nil {
			return err, false
		}
		m.Granted = grantedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTbTotalFlags(depth int, m *models.TrafficQuotaDataAttributesQuotaPerProjectItems0QuotaPerRegionItems0QuotaInTb, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalFlagName := fmt.Sprintf("%v.total", cmdPrefix)
	if cmd.Flags().Changed(totalFlagName) {

		var totalFlagName string
		if cmdPrefix == "" {
			totalFlagName = "total"
		} else {
			totalFlagName = fmt.Sprintf("%v.total", cmdPrefix)
		}

		totalFlagValue, err := cmd.Flags().GetInt64(totalFlagName)
		if err != nil {
			return err, false
		}
		m.Total = totalFlagValue

		retAdded = true
	}

	return nil, retAdded
}
