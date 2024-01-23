// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/servers"
	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServersDestroyServerCmd returns a cmd to handle operation destroyServer
func makeOperationServersDestroyServerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroy-server",
		Short: ``,
		RunE:  runOperationServersDestroyServer,
	}

	if err := registerOperationServersDestroyServerParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServersDestroyServer uses cmd flags to call endpoint api
func runOperationServersDestroyServer(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := servers.NewDestroyServerParams()
	if err, _ := retrieveOperationServersDestroyServerAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersDestroyServerServerIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.Servers.DestroyServer(params, nil)
	if err != nil {
		api.RenderErrorOutput(err)
		return nil
	}

	msgStr, err := parseOperationServersDestroyServerResult(result)
	if err != nil {
		return err
	}

	if !debug {
		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationServersDestroyServerParamFlags registers all flags needed to fill params
func registerOperationServersDestroyServerParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServersDestroyServerAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersDestroyServerServerIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServersDestroyServerAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServersDestroyServerServerIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	serverIdDescription := `Required. The server ID`

	var serverIdFlagName string
	if cmdPrefix == "" {
		serverIdFlagName = "server_id"
	} else {
		serverIdFlagName = fmt.Sprintf("%v.server_id", cmdPrefix)
	}

	var serverIdFlagDefault string

	_ = cmd.PersistentFlags().String(serverIdFlagName, serverIdFlagDefault, serverIdDescription)
	cmd.MarkPersistentFlagRequired(serverIdFlagName)

	return nil
}

func retrieveOperationServersDestroyServerAPIVersionFlag(m *servers.DestroyServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServersDestroyServerServerIDFlag(m *servers.DestroyServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("server_id") {

		var serverIdFlagName string
		if cmdPrefix == "" {
			serverIdFlagName = "server_id"
		} else {
			serverIdFlagName = fmt.Sprintf("%v.server_id", cmdPrefix)
		}

		serverIdFlagValue, err := cmd.Flags().GetString(serverIdFlagName)
		if err != nil {
			return err, false
		}
		m.ServerID = serverIdFlagValue

	}
	return nil, retAdded
}

// parseOperationServersDestroyServerResult parses request result and return the string content
func parseOperationServersDestroyServerResult(resp0 *servers.DestroyServerNoContent) (string, error) {
	// warning: non schema response destroyServerNoContent is not supported by go-swagger cli yet.

	return "", nil
}
