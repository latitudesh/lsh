package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/cmd/lsh"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationServersGetServersCmd returns a cmd to handle operation getServers
func makeOperationServersGetServersCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List servers",
		Long:  `Returns a list of all servers belonging to the team.`,
		RunE:  runOperationServersGetServers,
	}

	if err := registerOperationServersGetServersParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServersGetServers uses cmd flags to call endpoint api
func runOperationServersGetServers(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := servers.NewGetServersParams()
	if err, _ := retrieveOperationServersGetServersExtraFieldsServersFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterCreatedAtGteFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterCreatedAtLteFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterDiskEqlFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterDiskLteFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterDiskGteFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterGpuFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterHostnameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterLabelFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterOperatingSystemFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterPlanFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterProjectFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterRAMEqlFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterRAMGteFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterRAMLteFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterRegionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterStatusFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServersFilterTagsFlag(params, "", cmd); err != nil {
		return err
	}
	if lsh.DryRun {

		lsh.LogDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Servers.GetServers(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !lsh.Debug {
		utils.Render(response.GetData())
	}
	return nil
}

// registerOperationServersGetServersParamFlags registers all flags needed to fill params
func registerOperationServersGetServersParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServersGetServersExtraFieldsServersParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterCreatedAtGteParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterCreatedAtLteParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterDiskEqlParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterDiskLteParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterDiskGteParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterGpuParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterHostnameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterLabelParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterOperatingSystemParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterPlanParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterProjectParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterRAMEqlParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterRAMGteParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterRAMLteParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterRegionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterStatusParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServersFilterTagsParamFlags("", cmd); err != nil {
		return err
	}

	return nil
}

func registerOperationServersGetServersExtraFieldsServersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	extraFieldsServersDescription := `The ` + "`" + `credentials` + "`" + ` are provided as extra attributes that is lazy loaded. To request it, just set ` + "`" + `extra_fields[servers]=credentials` + "`" + ` in the query string.`

	var extraFieldsServersFlagName string
	if cmdPrefix == "" {
		extraFieldsServersFlagName = "extra_fields[servers]"
	} else {
		extraFieldsServersFlagName = fmt.Sprintf("%v.extra_fields[servers]", cmdPrefix)
	}

	var extraFieldsServersFlagDefault string

	_ = cmd.PersistentFlags().String(extraFieldsServersFlagName, extraFieldsServersFlagDefault, extraFieldsServersDescription)

	return nil
}
func registerOperationServersGetServersFilterCreatedAtGteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterCreatedAtGteDescription := `The created at greater than equal date to filter by`

	var filterCreatedAtGteFlagName string
	if cmdPrefix == "" {
		filterCreatedAtGteFlagName = "created_at_gte"
	} else {
		filterCreatedAtGteFlagName = fmt.Sprintf("%v.created_at_gte", cmdPrefix)
	}

	var filterCreatedAtGteFlagDefault string

	_ = cmd.PersistentFlags().String(filterCreatedAtGteFlagName, filterCreatedAtGteFlagDefault, filterCreatedAtGteDescription)

	return nil
}
func registerOperationServersGetServersFilterCreatedAtLteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterCreatedAtLteDescription := `The created at less than equal date to filter by`

	var filterCreatedAtLteFlagName string
	if cmdPrefix == "" {
		filterCreatedAtLteFlagName = "created_at_lte"
	} else {
		filterCreatedAtLteFlagName = fmt.Sprintf("%v.created_at_lte", cmdPrefix)
	}

	var filterCreatedAtLteFlagDefault string

	_ = cmd.PersistentFlags().String(filterCreatedAtLteFlagName, filterCreatedAtLteFlagDefault, filterCreatedAtLteDescription)

	return nil
}
func registerOperationServersGetServersFilterDiskEqlParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterDiskDescription := "Filter servers with disk size in Gigabytes equal the provided value."

	var filterDiskFlagName string
	if cmdPrefix == "" {
		filterDiskFlagName = "disk_eql"
	} else {
		filterDiskFlagName = fmt.Sprintf("%v.disk_eql", cmdPrefix)
	}

	var filterDiskFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterDiskFlagName, filterDiskFlagDefault, filterDiskDescription)

	return nil
}
func registerOperationServersGetServersFilterDiskLteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterDiskDescription := "Filter servers with disk size in Gigabytes less than or equal the provided value."

	var filterDiskFlagName string
	if cmdPrefix == "" {
		filterDiskFlagName = "disk_lte"
	} else {
		filterDiskFlagName = fmt.Sprintf("%v.disk_lte", cmdPrefix)
	}

	var filterDiskFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterDiskFlagName, filterDiskFlagDefault, filterDiskDescription)

	return nil
}
func registerOperationServersGetServersFilterDiskGteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterDiskDescription := "Filter servers with disk size in Gigabytes greater than or equal the provided value."

	var filterDiskFlagName string
	if cmdPrefix == "" {
		filterDiskFlagName = "disk_gte"
	} else {
		filterDiskFlagName = fmt.Sprintf("%v.disk_gte", cmdPrefix)
	}

	var filterDiskFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterDiskFlagName, filterDiskFlagDefault, filterDiskDescription)

	return nil
}
func registerOperationServersGetServersFilterGpuParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterGpuDescription := `Filter by the existence of an associated GPU`

	var filterGpuFlagName string
	if cmdPrefix == "" {
		filterGpuFlagName = "gpu"
	} else {
		filterGpuFlagName = fmt.Sprintf("%v.gpu", cmdPrefix)
	}

	var filterGpuFlagDefault bool

	_ = cmd.PersistentFlags().Bool(filterGpuFlagName, filterGpuFlagDefault, filterGpuDescription)

	return nil
}
func registerOperationServersGetServersFilterHostnameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterHostnameDescription := `The hostname of server to filter by`

	var filterHostnameFlagName string
	if cmdPrefix == "" {
		filterHostnameFlagName = "hostname"
	} else {
		filterHostnameFlagName = fmt.Sprintf("%v.hostname", cmdPrefix)
	}

	var filterHostnameFlagDefault string

	_ = cmd.PersistentFlags().String(filterHostnameFlagName, filterHostnameFlagDefault, filterHostnameDescription)

	return nil
}
func registerOperationServersGetServersFilterLabelParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterLabelDescription := `The label of server to filter by`

	var filterLabelFlagName string
	if cmdPrefix == "" {
		filterLabelFlagName = "label"
	} else {
		filterLabelFlagName = fmt.Sprintf("%v.label", cmdPrefix)
	}

	var filterLabelFlagDefault string

	_ = cmd.PersistentFlags().String(filterLabelFlagName, filterLabelFlagDefault, filterLabelDescription)

	return nil
}
func registerOperationServersGetServersFilterOperatingSystemParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterOperatingSystemDescription := `The operating system name or slug of the server to filter by`

	var filterOperatingSystemFlagName string
	if cmdPrefix == "" {
		filterOperatingSystemFlagName = "operating_system"
	} else {
		filterOperatingSystemFlagName = fmt.Sprintf("%v.operating_system", cmdPrefix)
	}

	var filterOperatingSystemFlagDefault string

	_ = cmd.PersistentFlags().String(filterOperatingSystemFlagName, filterOperatingSystemFlagDefault, filterOperatingSystemDescription)

	return nil
}
func registerOperationServersGetServersFilterPlanParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterPlanDescription := `The platform/plan name of the server to filter by`

	var filterPlanFlagName string
	if cmdPrefix == "" {
		filterPlanFlagName = "plan"
	} else {
		filterPlanFlagName = fmt.Sprintf("%v.plan", cmdPrefix)
	}

	var filterPlanFlagDefault string

	_ = cmd.PersistentFlags().String(filterPlanFlagName, filterPlanFlagDefault, filterPlanDescription)

	return nil
}
func registerOperationServersGetServersFilterProjectParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterProjectDescription := `The project ID or Slug to filter by`

	var filterProjectFlagName string
	if cmdPrefix == "" {
		filterProjectFlagName = "project"
	} else {
		filterProjectFlagName = fmt.Sprintf("%v.project", cmdPrefix)
	}

	var filterProjectFlagDefault string

	_ = cmd.PersistentFlags().String(filterProjectFlagName, filterProjectFlagDefault, filterProjectDescription)

	return nil
}
func registerOperationServersGetServersFilterRAMEqlParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterRamEqlDescription := `Filter servers with RAM size (in GB) equals the provided value.`

	var filterRamEqlFlagName string
	if cmdPrefix == "" {
		filterRamEqlFlagName = "ram_eql"
	} else {
		filterRamEqlFlagName = fmt.Sprintf("%v.ram_eql", cmdPrefix)
	}

	var filterRamEqlFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterRamEqlFlagName, filterRamEqlFlagDefault, filterRamEqlDescription)

	return nil
}
func registerOperationServersGetServersFilterRAMGteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterRamGteDescription := `Filter servers with RAM size (in GB) greater than or equal the provided value.`

	var filterRamGteFlagName string
	if cmdPrefix == "" {
		filterRamGteFlagName = "ram_gte"
	} else {
		filterRamGteFlagName = fmt.Sprintf("%v.ram_gte", cmdPrefix)
	}

	var filterRamGteFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterRamGteFlagName, filterRamGteFlagDefault, filterRamGteDescription)

	return nil
}
func registerOperationServersGetServersFilterRAMLteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterRamLteDescription := `Filter servers with RAM size (in GB) less than or equal the provided value.`

	var filterRamLteFlagName string
	if cmdPrefix == "" {
		filterRamLteFlagName = "ram_lte"
	} else {
		filterRamLteFlagName = fmt.Sprintf("%v.ram_lte", cmdPrefix)
	}

	var filterRamLteFlagDefault int64

	_ = cmd.PersistentFlags().Int64(filterRamLteFlagName, filterRamLteFlagDefault, filterRamLteDescription)

	return nil
}
func registerOperationServersGetServersFilterRegionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterRegionDescription := `The region Slug to filter by`

	var filterRegionFlagName string
	if cmdPrefix == "" {
		filterRegionFlagName = "region"
	} else {
		filterRegionFlagName = fmt.Sprintf("%v.region", cmdPrefix)
	}

	var filterRegionFlagDefault string

	_ = cmd.PersistentFlags().String(filterRegionFlagName, filterRegionFlagDefault, filterRegionDescription)

	return nil
}
func registerOperationServersGetServersFilterStatusParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterStatusDescription := `The status of server to filter by`

	var filterStatusFlagName string
	if cmdPrefix == "" {
		filterStatusFlagName = "status"
	} else {
		filterStatusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var filterStatusFlagDefault string

	_ = cmd.PersistentFlags().String(filterStatusFlagName, filterStatusFlagDefault, filterStatusDescription)

	return nil
}
func registerOperationServersGetServersFilterTagsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationServersGetServersExtraFieldsServersFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("extra_fields[servers]") {

		var extraFieldsServersFlagName string
		if cmdPrefix == "" {
			extraFieldsServersFlagName = "extra_fields[servers]"
		} else {
			extraFieldsServersFlagName = fmt.Sprintf("%v.extra_fields[servers]", cmdPrefix)
		}

		extraFieldsServersFlagValue, err := cmd.Flags().GetString(extraFieldsServersFlagName)
		if err != nil {
			return err, false
		}
		m.ExtraFieldsServers = &extraFieldsServersFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterCreatedAtGteFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("created_at_gte") {

		var filterCreatedAtGteFlagName string
		if cmdPrefix == "" {
			filterCreatedAtGteFlagName = "created_at_gte"
		} else {
			filterCreatedAtGteFlagName = fmt.Sprintf("%v.created_at_gte", cmdPrefix)
		}

		filterCreatedAtGteFlagValue, err := cmd.Flags().GetString(filterCreatedAtGteFlagName)
		if err != nil {
			return err, false
		}
		m.FilterCreatedAtGte = &filterCreatedAtGteFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterCreatedAtLteFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("created_at_lte") {

		var filterCreatedAtLteFlagName string
		if cmdPrefix == "" {
			filterCreatedAtLteFlagName = "created_at_lte"
		} else {
			filterCreatedAtLteFlagName = fmt.Sprintf("%v.created_at_lte", cmdPrefix)
		}

		filterCreatedAtLteFlagValue, err := cmd.Flags().GetString(filterCreatedAtLteFlagName)
		if err != nil {
			return err, false
		}
		m.FilterCreatedAtLte = &filterCreatedAtLteFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterDiskEqlFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("disk_eql") {

		var filterDiskFlagName string
		if cmdPrefix == "" {
			filterDiskFlagName = "disk_eql"
		} else {
			filterDiskFlagName = fmt.Sprintf("%v.disk_eql", cmdPrefix)
		}

		filterDiskFlagValue, err := cmd.Flags().GetInt64(filterDiskFlagName)
		if err != nil {
			return err, false
		}
		m.FilterDiskEql = &filterDiskFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterDiskLteFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("disk_lte") {

		var filterDiskFlagName string
		if cmdPrefix == "" {
			filterDiskFlagName = "disk_lte"
		} else {
			filterDiskFlagName = fmt.Sprintf("%v.disk_lte", cmdPrefix)
		}

		filterDiskFlagValue, err := cmd.Flags().GetInt64(filterDiskFlagName)
		if err != nil {
			return err, false
		}
		m.FilterDiskLte = &filterDiskFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterDiskGteFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("disk_gte") {

		var filterDiskFlagName string
		if cmdPrefix == "" {
			filterDiskFlagName = "disk_gte"
		} else {
			filterDiskFlagName = fmt.Sprintf("%v.disk_gte", cmdPrefix)
		}

		filterDiskFlagValue, err := cmd.Flags().GetInt64(filterDiskFlagName)
		if err != nil {
			return err, false
		}
		m.FilterDiskGte = &filterDiskFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterGpuFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("gpu") {

		var filterGpuFlagName string
		if cmdPrefix == "" {
			filterGpuFlagName = "gpu"
		} else {
			filterGpuFlagName = fmt.Sprintf("%v.gpu", cmdPrefix)
		}

		filterGpuFlagValue, err := cmd.Flags().GetBool(filterGpuFlagName)
		if err != nil {
			return err, false
		}
		m.FilterGpu = &filterGpuFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterHostnameFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("hostname") {

		var filterHostnameFlagName string
		if cmdPrefix == "" {
			filterHostnameFlagName = "hostname"
		} else {
			filterHostnameFlagName = fmt.Sprintf("%v.hostname", cmdPrefix)
		}

		filterHostnameFlagValue, err := cmd.Flags().GetString(filterHostnameFlagName)
		if err != nil {
			return err, false
		}
		m.FilterHostname = &filterHostnameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterLabelFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("label") {

		var filterLabelFlagName string
		if cmdPrefix == "" {
			filterLabelFlagName = "label"
		} else {
			filterLabelFlagName = fmt.Sprintf("%v.label", cmdPrefix)
		}

		filterLabelFlagValue, err := cmd.Flags().GetString(filterLabelFlagName)
		if err != nil {
			return err, false
		}
		m.FilterLabel = &filterLabelFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterOperatingSystemFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("operating_system") {

		var filterOperatingSystemFlagName string
		if cmdPrefix == "" {
			filterOperatingSystemFlagName = "operating_system"
		} else {
			filterOperatingSystemFlagName = fmt.Sprintf("%v.operating_system", cmdPrefix)
		}

		filterOperatingSystemFlagValue, err := cmd.Flags().GetString(filterOperatingSystemFlagName)
		if err != nil {
			return err, false
		}
		m.FilterOperatingSystem = &filterOperatingSystemFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterPlanFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("plan") {

		var filterPlanFlagName string
		if cmdPrefix == "" {
			filterPlanFlagName = "plan"
		} else {
			filterPlanFlagName = fmt.Sprintf("%v.plan", cmdPrefix)
		}

		filterPlanFlagValue, err := cmd.Flags().GetString(filterPlanFlagName)
		if err != nil {
			return err, false
		}
		m.FilterPlan = &filterPlanFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterProjectFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("project") {

		var filterProjectFlagName string
		if cmdPrefix == "" {
			filterProjectFlagName = "project"
		} else {
			filterProjectFlagName = fmt.Sprintf("%v.project", cmdPrefix)
		}

		filterProjectFlagValue, err := cmd.Flags().GetString(filterProjectFlagName)
		if err != nil {
			return err, false
		}
		m.FilterProject = &filterProjectFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterRAMEqlFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ram_eql") {

		var filterRamEqlFlagName string
		if cmdPrefix == "" {
			filterRamEqlFlagName = "ram_eql"
		} else {
			filterRamEqlFlagName = fmt.Sprintf("%v.ram_eql", cmdPrefix)
		}

		filterRamEqlFlagValue, err := cmd.Flags().GetInt64(filterRamEqlFlagName)
		if err != nil {
			return err, false
		}
		m.FilterRAMEql = &filterRamEqlFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterRAMGteFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ram_gte") {

		var filterRamGteFlagName string
		if cmdPrefix == "" {
			filterRamGteFlagName = "ram_gte"
		} else {
			filterRamGteFlagName = fmt.Sprintf("%v.ram_gte", cmdPrefix)
		}

		filterRamGteFlagValue, err := cmd.Flags().GetInt64(filterRamGteFlagName)
		if err != nil {
			return err, false
		}
		m.FilterRAMGte = &filterRamGteFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterRAMLteFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ram_lte") {

		var filterRamLteFlagName string
		if cmdPrefix == "" {
			filterRamLteFlagName = "ram_lte"
		} else {
			filterRamLteFlagName = fmt.Sprintf("%v.ram_lte", cmdPrefix)
		}

		filterRamLteFlagValue, err := cmd.Flags().GetInt64(filterRamLteFlagName)
		if err != nil {
			return err, false
		}
		m.FilterRAMLte = &filterRamLteFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterRegionFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("region") {

		var filterRegionFlagName string
		if cmdPrefix == "" {
			filterRegionFlagName = "region"
		} else {
			filterRegionFlagName = fmt.Sprintf("%v.region", cmdPrefix)
		}

		filterRegionFlagValue, err := cmd.Flags().GetString(filterRegionFlagName)
		if err != nil {
			return err, false
		}
		m.FilterRegion = &filterRegionFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterStatusFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("status") {

		var filterStatusFlagName string
		if cmdPrefix == "" {
			filterStatusFlagName = "status"
		} else {
			filterStatusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		filterStatusFlagValue, err := cmd.Flags().GetString(filterStatusFlagName)
		if err != nil {
			return err, false
		}
		m.FilterStatus = &filterStatusFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServersGetServersFilterTagsFlag(m *servers.GetServersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
