package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/client/virtual_network_assignments"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsCmd returns a cmd to handle operation getVirtualNetworksAssignments
func makeOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List assignments",
		Long:  `Returns a list of all servers assigned to virtual networks.`,
		RunE:  runOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignments,
	}

	if err := registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignments uses cmd flags to call endpoint api
func runOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignments(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := virtual_network_assignments.NewGetVirtualNetworksAssignmentsParams()
	if err, _ := retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterServerFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVidFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVirtualNetworkIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.VirtualNetworkAssignments.GetVirtualNetworksAssignments(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}

// registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsParamFlags registers all flags needed to fill params
func registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterServerParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVidParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVirtualNetworkIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterServerParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterServerDescription := `The server ID to filter by`

	var filterServerFlagName string
	if cmdPrefix == "" {
		filterServerFlagName = "server"
	} else {
		filterServerFlagName = fmt.Sprintf("%v.server", cmdPrefix)
	}

	var filterServerFlagDefault string

	_ = cmd.PersistentFlags().String(filterServerFlagName, filterServerFlagDefault, filterServerDescription)

	return nil
}
func registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVidParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterVidDescription := `The vlan ID to filter by`

	var filterVidFlagName string
	if cmdPrefix == "" {
		filterVidFlagName = "vid"
	} else {
		filterVidFlagName = fmt.Sprintf("%v.vid", cmdPrefix)
	}

	var filterVidFlagDefault string

	_ = cmd.PersistentFlags().String(filterVidFlagName, filterVidFlagDefault, filterVidDescription)

	return nil
}
func registerOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVirtualNetworkIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterVirtualNetworkIdDescription := `The virtual network ID to filter by`

	var filterVirtualNetworkIdFlagName string
	if cmdPrefix == "" {
		filterVirtualNetworkIdFlagName = "virtual_network_id"
	} else {
		filterVirtualNetworkIdFlagName = fmt.Sprintf("%v.virtual_network_id", cmdPrefix)
	}

	var filterVirtualNetworkIdFlagDefault string

	_ = cmd.PersistentFlags().String(filterVirtualNetworkIdFlagName, filterVirtualNetworkIdFlagDefault, filterVirtualNetworkIdDescription)

	return nil
}

func retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterServerFlag(m *virtual_network_assignments.GetVirtualNetworksAssignmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("server") {

		var filterServerFlagName string
		if cmdPrefix == "" {
			filterServerFlagName = "server"
		} else {
			filterServerFlagName = fmt.Sprintf("%v.server", cmdPrefix)
		}

		filterServerFlagValue, err := cmd.Flags().GetString(filterServerFlagName)
		if err != nil {
			return err, false
		}
		m.FilterServer = &filterServerFlagValue

	}
	return nil, retAdded
}
func retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVidFlag(m *virtual_network_assignments.GetVirtualNetworksAssignmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("vid") {

		var filterVidFlagName string
		if cmdPrefix == "" {
			filterVidFlagName = "vid"
		} else {
			filterVidFlagName = fmt.Sprintf("%v.vid", cmdPrefix)
		}

		filterVidFlagValue, err := cmd.Flags().GetString(filterVidFlagName)
		if err != nil {
			return err, false
		}
		m.FilterVid = &filterVidFlagValue

	}
	return nil, retAdded
}
func retrieveOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsFilterVirtualNetworkIDFlag(m *virtual_network_assignments.GetVirtualNetworksAssignmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("virtual_network_id") {

		var filterVirtualNetworkIdFlagName string
		if cmdPrefix == "" {
			filterVirtualNetworkIdFlagName = "virtual_network_id"
		} else {
			filterVirtualNetworkIdFlagName = fmt.Sprintf("%v.virtual_network_id", cmdPrefix)
		}

		filterVirtualNetworkIdFlagValue, err := cmd.Flags().GetString(filterVirtualNetworkIdFlagName)
		if err != nil {
			return err, false
		}
		m.FilterVirtualNetworkID = &filterVirtualNetworkIdFlagValue

	}
	return nil, retAdded
}
