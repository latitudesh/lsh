// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/virtual_networks"
	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVirtualNetworksDestroyVirtualNetworkCmd returns a cmd to handle operation destroyVirtualNetwork
func makeOperationVirtualNetworksDestroyVirtualNetworkCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "destroy-virtual-network",
		Short: `Delete virtual network
`,
		RunE: runOperationVirtualNetworksDestroyVirtualNetwork,
	}

	if err := registerOperationVirtualNetworksDestroyVirtualNetworkParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVirtualNetworksDestroyVirtualNetwork uses cmd flags to call endpoint api
func runOperationVirtualNetworksDestroyVirtualNetwork(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := virtual_networks.NewDestroyVirtualNetworkParams()
	if err, _ := retrieveOperationVirtualNetworksDestroyVirtualNetworkAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVirtualNetworksDestroyVirtualNetworkIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.VirtualNetworks.DestroyVirtualNetwork(params, nil)
	if err != nil {
		api.RenderErrorOutput(err)
		return nil
	}

	msgStr, err := parseOperationVirtualNetworksDestroyVirtualNetworkResult(result)
	if err != nil {
		return err
	}
	if !debug {

		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationVirtualNetworksDestroyVirtualNetworkParamFlags registers all flags needed to fill params
func registerOperationVirtualNetworksDestroyVirtualNetworkParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVirtualNetworksDestroyVirtualNetworkAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVirtualNetworksDestroyVirtualNetworkIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVirtualNetworksDestroyVirtualNetworkAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	apiVersionDescription := ``

	var apiVersionFlagName string
	if cmdPrefix == "" {
		apiVersionFlagName = "API-Version"
	} else {
		apiVersionFlagName = fmt.Sprintf("%v.API-Version", cmdPrefix)
	}

	var apiVersionFlagDefault string = "2023-06-01"

	_ = cmd.PersistentFlags().String(apiVersionFlagName, apiVersionFlagDefault, apiVersionDescription)

	return nil
}
func registerOperationVirtualNetworksDestroyVirtualNetworkIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. The virtual network ID`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)
	cmd.MarkPersistentFlagRequired(idFlagName)

	return nil
}

func retrieveOperationVirtualNetworksDestroyVirtualNetworkAPIVersionFlag(m *virtual_networks.DestroyVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("API-Version") {

		var apiVersionFlagName string
		if cmdPrefix == "" {
			apiVersionFlagName = "API-Version"
		} else {
			apiVersionFlagName = fmt.Sprintf("%v.API-Version", cmdPrefix)
		}

		apiVersionFlagValue, err := cmd.Flags().GetString(apiVersionFlagName)
		if err != nil {
			return err, false
		}
		m.APIVersion = &apiVersionFlagValue

	}
	return nil, retAdded
}
func retrieveOperationVirtualNetworksDestroyVirtualNetworkIDFlag(m *virtual_networks.DestroyVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// parseOperationVirtualNetworksDestroyVirtualNetworkResult parses request result and return the string content
func parseOperationVirtualNetworksDestroyVirtualNetworkResult(resp0 *virtual_networks.DestroyVirtualNetworkNoContent) (string, error) {
	// warning: non schema response destroyVirtualNetworkNoContent is not supported by go-swagger cli yet.

	return "", nil
}
