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

// Schema cli for BandwidthPlan

// register flags to command
func registerModelBandwidthPlanFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerBandwidthPlanAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBandwidthPlanID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBandwidthPlanType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerBandwidthPlanAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelBandwidthPlanAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerBandwidthPlanID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerBandwidthPlanType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["bandwidth_plan"]. `

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
			if err := json.Unmarshal([]byte(`["bandwidth_plan"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelBandwidthPlanFlags(depth int, m *models.BandwidthPlan, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveBandwidthPlanAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveBandwidthPlanIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, typeAdded := retrieveBandwidthPlanTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveBandwidthPlanAttributesFlags(depth int, m *models.BandwidthPlan, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes BandwidthPlanAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.BandwidthPlanAttributes{}
	}

	err, attributesAdded := retrieveModelBandwidthPlanAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveBandwidthPlanIDFlags(depth int, m *models.BandwidthPlan, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveBandwidthPlanTypeFlags(depth int, m *models.BandwidthPlan, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for BandwidthPlanAttributes

// register flags to command
func registerModelBandwidthPlanAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerBandwidthPlanAttributesLocations(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBandwidthPlanAttributesPricing(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBandwidthPlanAttributesRegion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerBandwidthPlanAttributesLocations(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: locations []string array type is not supported by go-swagger cli yet

	return nil
}

func registerBandwidthPlanAttributesPricing(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var pricingFlagName string
	if cmdPrefix == "" {
		pricingFlagName = "pricing"
	} else {
		pricingFlagName = fmt.Sprintf("%v.pricing", cmdPrefix)
	}

	if err := registerModelBandwidthPlanAttributesPricingFlags(depth+1, pricingFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerBandwidthPlanAttributesRegion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	regionDescription := ``

	var regionFlagName string
	if cmdPrefix == "" {
		regionFlagName = "region"
	} else {
		regionFlagName = fmt.Sprintf("%v.region", cmdPrefix)
	}

	var regionFlagDefault string

	_ = cmd.PersistentFlags().String(regionFlagName, regionFlagDefault, regionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelBandwidthPlanAttributesFlags(depth int, m *models.BandwidthPlanAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, locationsAdded := retrieveBandwidthPlanAttributesLocationsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || locationsAdded

	err, pricingAdded := retrieveBandwidthPlanAttributesPricingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pricingAdded

	err, regionAdded := retrieveBandwidthPlanAttributesRegionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regionAdded

	return nil, retAdded
}

func retrieveBandwidthPlanAttributesLocationsFlags(depth int, m *models.BandwidthPlanAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	locationsFlagName := fmt.Sprintf("%v.locations", cmdPrefix)
	if cmd.Flags().Changed(locationsFlagName) {
		// warning: locations array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveBandwidthPlanAttributesPricingFlags(depth int, m *models.BandwidthPlanAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pricingFlagName := fmt.Sprintf("%v.pricing", cmdPrefix)
	if cmd.Flags().Changed(pricingFlagName) {
		// info: complex object pricing BandwidthPlanAttributesPricing is retrieved outside this Changed() block
	}
	pricingFlagValue := m.Pricing
	if swag.IsZero(pricingFlagValue) {
		pricingFlagValue = &models.BandwidthPlanAttributesPricing{}
	}

	err, pricingAdded := retrieveModelBandwidthPlanAttributesPricingFlags(depth+1, pricingFlagValue, pricingFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pricingAdded
	if pricingAdded {
		m.Pricing = pricingFlagValue
	}

	return nil, retAdded
}

func retrieveBandwidthPlanAttributesRegionFlags(depth int, m *models.BandwidthPlanAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	regionFlagName := fmt.Sprintf("%v.region", cmdPrefix)
	if cmd.Flags().Changed(regionFlagName) {

		var regionFlagName string
		if cmdPrefix == "" {
			regionFlagName = "region"
		} else {
			regionFlagName = fmt.Sprintf("%v.region", cmdPrefix)
		}

		regionFlagValue, err := cmd.Flags().GetString(regionFlagName)
		if err != nil {
			return err, false
		}
		m.Region = regionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for BandwidthPlanAttributesPricing

// register flags to command
func registerModelBandwidthPlanAttributesPricingFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerBandwidthPlanAttributesPricingBrl(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBandwidthPlanAttributesPricingUsd(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerBandwidthPlanAttributesPricingBrl(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var brlFlagName string
	if cmdPrefix == "" {
		brlFlagName = "brl"
	} else {
		brlFlagName = fmt.Sprintf("%v.brl", cmdPrefix)
	}

	if err := registerModelBandwidthPlanAttributesPricingBrlFlags(depth+1, brlFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerBandwidthPlanAttributesPricingUsd(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var usdFlagName string
	if cmdPrefix == "" {
		usdFlagName = "usd"
	} else {
		usdFlagName = fmt.Sprintf("%v.usd", cmdPrefix)
	}

	if err := registerModelBandwidthPlanAttributesPricingUsdFlags(depth+1, usdFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelBandwidthPlanAttributesPricingFlags(depth int, m *models.BandwidthPlanAttributesPricing, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, brlAdded := retrieveBandwidthPlanAttributesPricingBrlFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || brlAdded

	err, usdAdded := retrieveBandwidthPlanAttributesPricingUsdFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || usdAdded

	return nil, retAdded
}

func retrieveBandwidthPlanAttributesPricingBrlFlags(depth int, m *models.BandwidthPlanAttributesPricing, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	brlFlagName := fmt.Sprintf("%v.brl", cmdPrefix)
	if cmd.Flags().Changed(brlFlagName) {
		// info: complex object brl BandwidthPlanAttributesPricingBrl is retrieved outside this Changed() block
	}
	brlFlagValue := m.Brl
	if swag.IsZero(brlFlagValue) {
		brlFlagValue = &models.BandwidthPlanAttributesPricingBrl{}
	}

	err, brlAdded := retrieveModelBandwidthPlanAttributesPricingBrlFlags(depth+1, brlFlagValue, brlFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || brlAdded
	if brlAdded {
		m.Brl = brlFlagValue
	}

	return nil, retAdded
}

func retrieveBandwidthPlanAttributesPricingUsdFlags(depth int, m *models.BandwidthPlanAttributesPricing, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	usdFlagName := fmt.Sprintf("%v.usd", cmdPrefix)
	if cmd.Flags().Changed(usdFlagName) {
		// info: complex object usd BandwidthPlanAttributesPricingUsd is retrieved outside this Changed() block
	}
	usdFlagValue := m.Usd
	if swag.IsZero(usdFlagValue) {
		usdFlagValue = &models.BandwidthPlanAttributesPricingUsd{}
	}

	err, usdAdded := retrieveModelBandwidthPlanAttributesPricingUsdFlags(depth+1, usdFlagValue, usdFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || usdAdded
	if usdAdded {
		m.Usd = usdFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for BandwidthPlanAttributesPricingBrl

// register flags to command
func registerModelBandwidthPlanAttributesPricingBrlFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerBandwidthPlanAttributesPricingBrlHourly(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBandwidthPlanAttributesPricingBrlMonthly(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerBandwidthPlanAttributesPricingBrlHourly(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	hourlyDescription := ``

	var hourlyFlagName string
	if cmdPrefix == "" {
		hourlyFlagName = "hourly"
	} else {
		hourlyFlagName = fmt.Sprintf("%v.hourly", cmdPrefix)
	}

	var hourlyFlagDefault int64

	_ = cmd.PersistentFlags().Int64(hourlyFlagName, hourlyFlagDefault, hourlyDescription)

	return nil
}

func registerBandwidthPlanAttributesPricingBrlMonthly(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	monthlyDescription := ``

	var monthlyFlagName string
	if cmdPrefix == "" {
		monthlyFlagName = "monthly"
	} else {
		monthlyFlagName = fmt.Sprintf("%v.monthly", cmdPrefix)
	}

	var monthlyFlagDefault int64

	_ = cmd.PersistentFlags().Int64(monthlyFlagName, monthlyFlagDefault, monthlyDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelBandwidthPlanAttributesPricingBrlFlags(depth int, m *models.BandwidthPlanAttributesPricingBrl, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, hourlyAdded := retrieveBandwidthPlanAttributesPricingBrlHourlyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hourlyAdded

	err, monthlyAdded := retrieveBandwidthPlanAttributesPricingBrlMonthlyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || monthlyAdded

	return nil, retAdded
}

func retrieveBandwidthPlanAttributesPricingBrlHourlyFlags(depth int, m *models.BandwidthPlanAttributesPricingBrl, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hourlyFlagName := fmt.Sprintf("%v.hourly", cmdPrefix)
	if cmd.Flags().Changed(hourlyFlagName) {

		var hourlyFlagName string
		if cmdPrefix == "" {
			hourlyFlagName = "hourly"
		} else {
			hourlyFlagName = fmt.Sprintf("%v.hourly", cmdPrefix)
		}

		hourlyFlagValue, err := cmd.Flags().GetInt64(hourlyFlagName)
		if err != nil {
			return err, false
		}
		m.Hourly = hourlyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveBandwidthPlanAttributesPricingBrlMonthlyFlags(depth int, m *models.BandwidthPlanAttributesPricingBrl, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	monthlyFlagName := fmt.Sprintf("%v.monthly", cmdPrefix)
	if cmd.Flags().Changed(monthlyFlagName) {

		var monthlyFlagName string
		if cmdPrefix == "" {
			monthlyFlagName = "monthly"
		} else {
			monthlyFlagName = fmt.Sprintf("%v.monthly", cmdPrefix)
		}

		monthlyFlagValue, err := cmd.Flags().GetInt64(monthlyFlagName)
		if err != nil {
			return err, false
		}
		m.Monthly = monthlyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for BandwidthPlanAttributesPricingUsd

// register flags to command
func registerModelBandwidthPlanAttributesPricingUsdFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerBandwidthPlanAttributesPricingUsdHourly(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBandwidthPlanAttributesPricingUsdMonthly(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerBandwidthPlanAttributesPricingUsdHourly(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	hourlyDescription := ``

	var hourlyFlagName string
	if cmdPrefix == "" {
		hourlyFlagName = "hourly"
	} else {
		hourlyFlagName = fmt.Sprintf("%v.hourly", cmdPrefix)
	}

	var hourlyFlagDefault int64

	_ = cmd.PersistentFlags().Int64(hourlyFlagName, hourlyFlagDefault, hourlyDescription)

	return nil
}

func registerBandwidthPlanAttributesPricingUsdMonthly(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	monthlyDescription := ``

	var monthlyFlagName string
	if cmdPrefix == "" {
		monthlyFlagName = "monthly"
	} else {
		monthlyFlagName = fmt.Sprintf("%v.monthly", cmdPrefix)
	}

	var monthlyFlagDefault int64

	_ = cmd.PersistentFlags().Int64(monthlyFlagName, monthlyFlagDefault, monthlyDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelBandwidthPlanAttributesPricingUsdFlags(depth int, m *models.BandwidthPlanAttributesPricingUsd, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, hourlyAdded := retrieveBandwidthPlanAttributesPricingUsdHourlyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hourlyAdded

	err, monthlyAdded := retrieveBandwidthPlanAttributesPricingUsdMonthlyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || monthlyAdded

	return nil, retAdded
}

func retrieveBandwidthPlanAttributesPricingUsdHourlyFlags(depth int, m *models.BandwidthPlanAttributesPricingUsd, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hourlyFlagName := fmt.Sprintf("%v.hourly", cmdPrefix)
	if cmd.Flags().Changed(hourlyFlagName) {

		var hourlyFlagName string
		if cmdPrefix == "" {
			hourlyFlagName = "hourly"
		} else {
			hourlyFlagName = fmt.Sprintf("%v.hourly", cmdPrefix)
		}

		hourlyFlagValue, err := cmd.Flags().GetInt64(hourlyFlagName)
		if err != nil {
			return err, false
		}
		m.Hourly = hourlyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveBandwidthPlanAttributesPricingUsdMonthlyFlags(depth int, m *models.BandwidthPlanAttributesPricingUsd, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	monthlyFlagName := fmt.Sprintf("%v.monthly", cmdPrefix)
	if cmd.Flags().Changed(monthlyFlagName) {

		var monthlyFlagName string
		if cmdPrefix == "" {
			monthlyFlagName = "monthly"
		} else {
			monthlyFlagName = fmt.Sprintf("%v.monthly", cmdPrefix)
		}

		monthlyFlagValue, err := cmd.Flags().GetInt64(monthlyFlagName)
		if err != nil {
			return err, false
		}
		m.Monthly = monthlyFlagValue

		retAdded = true
	}

	return nil, retAdded
}
