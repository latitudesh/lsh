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

// Schema cli for VirtualNetwork

// register flags to command
func registerModelVirtualNetworkFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVirtualNetworkAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVirtualNetworkAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelVirtualNetworkAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerVirtualNetworkID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerVirtualNetworkType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["virtual_networks"]. `

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
			if err := json.Unmarshal([]byte(`["virtual_networks"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVirtualNetworkFlags(depth int, m *models.VirtualNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveVirtualNetworkAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveVirtualNetworkIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, typeAdded := retrieveVirtualNetworkTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesFlags(depth int, m *models.VirtualNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes VirtualNetworkAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.VirtualNetworkAttributes{}
	}

	err, attributesAdded := retrieveModelVirtualNetworkAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveVirtualNetworkIDFlags(depth int, m *models.VirtualNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveVirtualNetworkTypeFlags(depth int, m *models.VirtualNetwork, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for VirtualNetworkAttributes

// register flags to command
func registerModelVirtualNetworkAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVirtualNetworkAttributesAssignmentsCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAttributesDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAttributesName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAttributesRegion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAttributesVid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVirtualNetworkAttributesAssignmentsCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	assignmentsCountDescription := `Amount of devices assigned to the virtual network`

	var assignmentsCountFlagName string
	if cmdPrefix == "" {
		assignmentsCountFlagName = "assignments_count"
	} else {
		assignmentsCountFlagName = fmt.Sprintf("%v.assignments_count", cmdPrefix)
	}

	var assignmentsCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(assignmentsCountFlagName, assignmentsCountFlagDefault, assignmentsCountDescription)

	return nil
}

func registerVirtualNetworkAttributesDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `Description of the virtual network`

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerVirtualNetworkAttributesName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name of the virtual network`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerVirtualNetworkAttributesRegion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var regionFlagName string
	if cmdPrefix == "" {
		regionFlagName = "region"
	} else {
		regionFlagName = fmt.Sprintf("%v.region", cmdPrefix)
	}

	if err := registerModelVirtualNetworkAttributesRegionFlags(depth+1, regionFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerVirtualNetworkAttributesVid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vidDescription := `vlan ID of the virtual network`

	var vidFlagName string
	if cmdPrefix == "" {
		vidFlagName = "vid"
	} else {
		vidFlagName = fmt.Sprintf("%v.vid", cmdPrefix)
	}

	var vidFlagDefault int64

	_ = cmd.PersistentFlags().Int64(vidFlagName, vidFlagDefault, vidDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVirtualNetworkAttributesFlags(depth int, m *models.VirtualNetworkAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, assignmentsCountAdded := retrieveVirtualNetworkAttributesAssignmentsCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || assignmentsCountAdded

	err, descriptionAdded := retrieveVirtualNetworkAttributesDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, nameAdded := retrieveVirtualNetworkAttributesNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, regionAdded := retrieveVirtualNetworkAttributesRegionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regionAdded

	err, vidAdded := retrieveVirtualNetworkAttributesVidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vidAdded

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesAssignmentsCountFlags(depth int, m *models.VirtualNetworkAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	assignmentsCountFlagName := fmt.Sprintf("%v.assignments_count", cmdPrefix)
	if cmd.Flags().Changed(assignmentsCountFlagName) {

		var assignmentsCountFlagName string
		if cmdPrefix == "" {
			assignmentsCountFlagName = "assignments_count"
		} else {
			assignmentsCountFlagName = fmt.Sprintf("%v.assignments_count", cmdPrefix)
		}

		assignmentsCountFlagValue, err := cmd.Flags().GetInt64(assignmentsCountFlagName)
		if err != nil {
			return err, false
		}
		m.AssignmentsCount = assignmentsCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesDescriptionFlags(depth int, m *models.VirtualNetworkAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesNameFlags(depth int, m *models.VirtualNetworkAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesRegionFlags(depth int, m *models.VirtualNetworkAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	regionFlagName := fmt.Sprintf("%v.region", cmdPrefix)
	if cmd.Flags().Changed(regionFlagName) {
		// info: complex object region VirtualNetworkAttributesRegion is retrieved outside this Changed() block
	}
	regionFlagValue := m.Region
	if swag.IsZero(regionFlagValue) {
		regionFlagValue = &models.VirtualNetworkAttributesRegion{}
	}

	err, regionAdded := retrieveModelVirtualNetworkAttributesRegionFlags(depth+1, regionFlagValue, regionFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regionAdded
	if regionAdded {
		m.Region = regionFlagValue
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesVidFlags(depth int, m *models.VirtualNetworkAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vidFlagName := fmt.Sprintf("%v.vid", cmdPrefix)
	if cmd.Flags().Changed(vidFlagName) {

		var vidFlagName string
		if cmdPrefix == "" {
			vidFlagName = "vid"
		} else {
			vidFlagName = fmt.Sprintf("%v.vid", cmdPrefix)
		}

		vidFlagValue, err := cmd.Flags().GetInt64(vidFlagName)
		if err != nil {
			return err, false
		}
		m.Vid = vidFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for VirtualNetworkAttributesRegion

// register flags to command
func registerModelVirtualNetworkAttributesRegionFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVirtualNetworkAttributesRegionCity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAttributesRegionCountry(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAttributesRegionSite(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVirtualNetworkAttributesRegionCity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	cityDescription := ``

	var cityFlagName string
	if cmdPrefix == "" {
		cityFlagName = "city"
	} else {
		cityFlagName = fmt.Sprintf("%v.city", cmdPrefix)
	}

	var cityFlagDefault string

	_ = cmd.PersistentFlags().String(cityFlagName, cityFlagDefault, cityDescription)

	return nil
}

func registerVirtualNetworkAttributesRegionCountry(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	countryDescription := ``

	var countryFlagName string
	if cmdPrefix == "" {
		countryFlagName = "country"
	} else {
		countryFlagName = fmt.Sprintf("%v.country", cmdPrefix)
	}

	var countryFlagDefault string

	_ = cmd.PersistentFlags().String(countryFlagName, countryFlagDefault, countryDescription)

	return nil
}

func registerVirtualNetworkAttributesRegionSite(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var siteFlagName string
	if cmdPrefix == "" {
		siteFlagName = "site"
	} else {
		siteFlagName = fmt.Sprintf("%v.site", cmdPrefix)
	}

	if err := registerModelVirtualNetworkAttributesRegionSiteFlags(depth+1, siteFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVirtualNetworkAttributesRegionFlags(depth int, m *models.VirtualNetworkAttributesRegion, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, cityAdded := retrieveVirtualNetworkAttributesRegionCityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cityAdded

	err, countryAdded := retrieveVirtualNetworkAttributesRegionCountryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || countryAdded

	err, siteAdded := retrieveVirtualNetworkAttributesRegionSiteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesRegionCityFlags(depth int, m *models.VirtualNetworkAttributesRegion, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cityFlagName := fmt.Sprintf("%v.city", cmdPrefix)
	if cmd.Flags().Changed(cityFlagName) {

		var cityFlagName string
		if cmdPrefix == "" {
			cityFlagName = "city"
		} else {
			cityFlagName = fmt.Sprintf("%v.city", cmdPrefix)
		}

		cityFlagValue, err := cmd.Flags().GetString(cityFlagName)
		if err != nil {
			return err, false
		}
		m.City = cityFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesRegionCountryFlags(depth int, m *models.VirtualNetworkAttributesRegion, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	countryFlagName := fmt.Sprintf("%v.country", cmdPrefix)
	if cmd.Flags().Changed(countryFlagName) {

		var countryFlagName string
		if cmdPrefix == "" {
			countryFlagName = "country"
		} else {
			countryFlagName = fmt.Sprintf("%v.country", cmdPrefix)
		}

		countryFlagValue, err := cmd.Flags().GetString(countryFlagName)
		if err != nil {
			return err, false
		}
		m.Country = countryFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesRegionSiteFlags(depth int, m *models.VirtualNetworkAttributesRegion, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	siteFlagName := fmt.Sprintf("%v.site", cmdPrefix)
	if cmd.Flags().Changed(siteFlagName) {
		// info: complex object site VirtualNetworkAttributesRegionSite is retrieved outside this Changed() block
	}
	siteFlagValue := m.Site
	if swag.IsZero(siteFlagValue) {
		siteFlagValue = &models.VirtualNetworkAttributesRegionSite{}
	}

	err, siteAdded := retrieveModelVirtualNetworkAttributesRegionSiteFlags(depth+1, siteFlagValue, siteFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded
	if siteAdded {
		m.Site = siteFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for VirtualNetworkAttributesRegionSite

// register flags to command
func registerModelVirtualNetworkAttributesRegionSiteFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVirtualNetworkAttributesRegionSiteFacility(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAttributesRegionSiteID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAttributesRegionSiteName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAttributesRegionSiteSlug(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVirtualNetworkAttributesRegionSiteFacility(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	facilityDescription := ``

	var facilityFlagName string
	if cmdPrefix == "" {
		facilityFlagName = "facility"
	} else {
		facilityFlagName = fmt.Sprintf("%v.facility", cmdPrefix)
	}

	var facilityFlagDefault string

	_ = cmd.PersistentFlags().String(facilityFlagName, facilityFlagDefault, facilityDescription)

	return nil
}

func registerVirtualNetworkAttributesRegionSiteID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerVirtualNetworkAttributesRegionSiteName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerVirtualNetworkAttributesRegionSiteSlug(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	slugDescription := ``

	var slugFlagName string
	if cmdPrefix == "" {
		slugFlagName = "slug"
	} else {
		slugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
	}

	var slugFlagDefault string

	_ = cmd.PersistentFlags().String(slugFlagName, slugFlagDefault, slugDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVirtualNetworkAttributesRegionSiteFlags(depth int, m *models.VirtualNetworkAttributesRegionSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, facilityAdded := retrieveVirtualNetworkAttributesRegionSiteFacilityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || facilityAdded

	err, idAdded := retrieveVirtualNetworkAttributesRegionSiteIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, nameAdded := retrieveVirtualNetworkAttributesRegionSiteNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, slugAdded := retrieveVirtualNetworkAttributesRegionSiteSlugFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || slugAdded

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesRegionSiteFacilityFlags(depth int, m *models.VirtualNetworkAttributesRegionSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	facilityFlagName := fmt.Sprintf("%v.facility", cmdPrefix)
	if cmd.Flags().Changed(facilityFlagName) {

		var facilityFlagName string
		if cmdPrefix == "" {
			facilityFlagName = "facility"
		} else {
			facilityFlagName = fmt.Sprintf("%v.facility", cmdPrefix)
		}

		facilityFlagValue, err := cmd.Flags().GetString(facilityFlagName)
		if err != nil {
			return err, false
		}
		m.Facility = facilityFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesRegionSiteIDFlags(depth int, m *models.VirtualNetworkAttributesRegionSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveVirtualNetworkAttributesRegionSiteNameFlags(depth int, m *models.VirtualNetworkAttributesRegionSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAttributesRegionSiteSlugFlags(depth int, m *models.VirtualNetworkAttributesRegionSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	slugFlagName := fmt.Sprintf("%v.slug", cmdPrefix)
	if cmd.Flags().Changed(slugFlagName) {

		var slugFlagName string
		if cmdPrefix == "" {
			slugFlagName = "slug"
		} else {
			slugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
		}

		slugFlagValue, err := cmd.Flags().GetString(slugFlagName)
		if err != nil {
			return err, false
		}
		m.Slug = slugFlagValue

		retAdded = true
	}

	return nil, retAdded
}
