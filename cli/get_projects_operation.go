package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/client/projects"
	"github.com/latitudesh/lsh/cmd/lsh"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationProjectsGetProjectsCmd returns a cmd to handle operation getProjects
func makeOperationProjectsGetProjectsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Long:  `Returns a list of all projects for the current team.`,
		RunE:  runOperationProjectsGetProjects,
	}

	if err := registerOperationProjectsGetProjectsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationProjectsGetProjects uses cmd flags to call endpoint api
func runOperationProjectsGetProjects(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := projects.NewGetProjectsParams()
	if err, _ := retrieveOperationProjectsGetProjectsExtraFieldsProjectsFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectsGetProjectsFilterBillingTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectsGetProjectsFilterDescriptionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectsGetProjectsFilterEnvironmentFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectsGetProjectsFilterNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectsGetProjectsFilterSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectsGetProjectsFilterTagsFlag(params, "", cmd); err != nil {
		return err
	}
	if lsh.DryRun {

		lsh.LogDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Projects.GetProjects(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !lsh.Debug {
		utils.Render(response.GetData())
	}
	return nil
}

// registerOperationProjectsGetProjectsParamFlags registers all flags needed to fill params
func registerOperationProjectsGetProjectsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationProjectsGetProjectsExtraFieldsProjectsParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectsGetProjectsFilterBillingTypeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectsGetProjectsFilterDescriptionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectsGetProjectsFilterEnvironmentParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectsGetProjectsFilterNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectsGetProjectsFilterSlugParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectsGetProjectsFilterTagsFlag("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationProjectsGetProjectsExtraFieldsProjectsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	extraFieldsProjectsDescription := `The ` + "`" + `last_renewal_date` + "`" + ` and ` + "`" + `next_renewal_date` + "`" + ` are provided as extra attributes that show previous and future billing cycle dates. To request it, just set ` + "`" + `extra_fields[projects]=last_renewal_date,next_renewal_date` + "`" + ` in the query string.`

	var extraFieldsProjectsFlagName string
	if cmdPrefix == "" {
		extraFieldsProjectsFlagName = "extra_fields[projects]"
	} else {
		extraFieldsProjectsFlagName = fmt.Sprintf("%v.extra_fields[projects]", cmdPrefix)
	}

	var extraFieldsProjectsFlagDefault string

	_ = cmd.PersistentFlags().String(extraFieldsProjectsFlagName, extraFieldsProjectsFlagDefault, extraFieldsProjectsDescription)

	return nil
}
func registerOperationProjectsGetProjectsFilterBillingTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterBillingTypeDescription := `The billing type to filter by`

	var filterBillingTypeFlagName string
	if cmdPrefix == "" {
		filterBillingTypeFlagName = "billing_type"
	} else {
		filterBillingTypeFlagName = fmt.Sprintf("%v.billing_type", cmdPrefix)
	}

	var filterBillingTypeFlagDefault string

	_ = cmd.PersistentFlags().String(filterBillingTypeFlagName, filterBillingTypeFlagDefault, filterBillingTypeDescription)

	return nil
}
func registerOperationProjectsGetProjectsFilterDescriptionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterDescriptionDescription := `The project description to filter by`

	var filterDescriptionFlagName string
	if cmdPrefix == "" {
		filterDescriptionFlagName = "description"
	} else {
		filterDescriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var filterDescriptionFlagDefault string

	_ = cmd.PersistentFlags().String(filterDescriptionFlagName, filterDescriptionFlagDefault, filterDescriptionDescription)

	return nil
}
func registerOperationProjectsGetProjectsFilterEnvironmentParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterEnvironmentDescription := `The environment to filter by`

	var filterEnvironmentFlagName string
	if cmdPrefix == "" {
		filterEnvironmentFlagName = "environment"
	} else {
		filterEnvironmentFlagName = fmt.Sprintf("%v.environment", cmdPrefix)
	}

	var filterEnvironmentFlagDefault string

	_ = cmd.PersistentFlags().String(filterEnvironmentFlagName, filterEnvironmentFlagDefault, filterEnvironmentDescription)

	return nil
}
func registerOperationProjectsGetProjectsFilterNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterNameDescription := `The project name to filter by`

	var filterNameFlagName string

	if cmdPrefix == "" {
		filterNameFlagName = "name"
	} else {
		filterNameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var filterNameFlagDefault string

	_ = cmd.PersistentFlags().String(filterNameFlagName, filterNameFlagDefault, filterNameDescription)

	return nil
}

func registerOperationProjectsGetProjectsFilterSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterSlugDescription := `The project slug to filter by`

	var filterSlugFlagName string
	if cmdPrefix == "" {
		filterSlugFlagName = "slug"
	} else {
		filterSlugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
	}

	var filterSlugFlagDefault string

	_ = cmd.PersistentFlags().String(filterSlugFlagName, filterSlugFlagDefault, filterSlugDescription)

	return nil
}

func registerOperationProjectsGetProjectsFilterTagsFlag(cmdPrefix string, cmd *cobra.Command) error {

	filterTagsDescription := `The Tags to filter by`

	var filterTagsFlagName string
	if cmdPrefix == "" {
		filterTagsFlagName = "tags"
	} else {
		filterTagsFlagName = fmt.Sprintf("%v.tags", cmdPrefix)
	}

	var filterTagsFlagDefault = ""

	_ = cmd.PersistentFlags().String(filterTagsFlagName, filterTagsFlagDefault, filterTagsDescription)

	return nil
}

func retrieveOperationProjectsGetProjectsExtraFieldsProjectsFlag(m *projects.GetProjectsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("extra_fields[projects]") {

		var extraFieldsProjectsFlagName string
		if cmdPrefix == "" {
			extraFieldsProjectsFlagName = "extra_fields[projects]"
		} else {
			extraFieldsProjectsFlagName = fmt.Sprintf("%v.extra_fields[projects]", cmdPrefix)
		}

		extraFieldsProjectsFlagValue, err := cmd.Flags().GetString(extraFieldsProjectsFlagName)
		if err != nil {
			return err, false
		}
		m.ExtraFieldsProjects = &extraFieldsProjectsFlagValue

	}
	return nil, retAdded
}
func retrieveOperationProjectsGetProjectsFilterBillingTypeFlag(m *projects.GetProjectsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("billing_type") {

		var filterBillingTypeFlagName string
		if cmdPrefix == "" {
			filterBillingTypeFlagName = "billing_type"
		} else {
			filterBillingTypeFlagName = fmt.Sprintf("%v.billing_type", cmdPrefix)
		}

		filterBillingTypeFlagValue, err := cmd.Flags().GetString(filterBillingTypeFlagName)
		if err != nil {
			return err, false
		}
		m.FilterBillingType = &filterBillingTypeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationProjectsGetProjectsFilterDescriptionFlag(m *projects.GetProjectsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("description") {

		var filterDescriptionFlagName string
		if cmdPrefix == "" {
			filterDescriptionFlagName = "description"
		} else {
			filterDescriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		filterDescriptionFlagValue, err := cmd.Flags().GetString(filterDescriptionFlagName)
		if err != nil {
			return err, false
		}
		m.FilterDescription = &filterDescriptionFlagValue

	}
	return nil, retAdded
}
func retrieveOperationProjectsGetProjectsFilterEnvironmentFlag(m *projects.GetProjectsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("environment") {

		var filterEnvironmentFlagName string
		if cmdPrefix == "" {
			filterEnvironmentFlagName = "environment"
		} else {
			filterEnvironmentFlagName = fmt.Sprintf("%v.environment", cmdPrefix)
		}

		filterEnvironmentFlagValue, err := cmd.Flags().GetString(filterEnvironmentFlagName)
		if err != nil {
			return err, false
		}
		m.FilterEnvironment = &filterEnvironmentFlagValue

	}
	return nil, retAdded
}
func retrieveOperationProjectsGetProjectsFilterNameFlag(m *projects.GetProjectsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("name") {

		var filterNameFlagName string
		if cmdPrefix == "" {
			filterNameFlagName = "name"
		} else {
			filterNameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		filterNameFlagValue, err := cmd.Flags().GetString(filterNameFlagName)
		if err != nil {
			return err, false
		}
		m.FilterName = &filterNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationProjectsGetProjectsFilterSlugFlag(m *projects.GetProjectsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("slug") {

		var filterSlugFlagName string
		if cmdPrefix == "" {
			filterSlugFlagName = "slug"
		} else {
			filterSlugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
		}

		filterSlugFlagValue, err := cmd.Flags().GetString(filterSlugFlagName)
		if err != nil {
			return err, false
		}
		m.FilterSlug = &filterSlugFlagValue

	}
	return nil, retAdded
}

func retrieveOperationProjectsGetProjectsFilterTagsFlag(m *projects.GetProjectsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tags") {

		var filterTagsFlagName string
		if cmdPrefix == "" {
			filterTagsFlagName = "tags"
		} else {
			filterTagsFlagName = fmt.Sprintf("%v.tags", cmdPrefix)
		}

		filterTagsFlagValue, err := cmd.Flags().GetString(filterTagsFlagName)
		if err != nil {
			return err, false
		}
		m.FilterTags = &filterTagsFlagValue

	}
	return nil, retAdded
}
