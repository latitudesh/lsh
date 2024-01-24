// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/api"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationServersServerUnscheduleDeletionCmd returns a cmd to handle operation serverUnscheduleDeletion
func makeOperationServersServerUnscheduleDeletionCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "unschedule-deletion",
		Short: `Unschedules the server removal at the end of the billing cycle.`,
		RunE:  runOperationServersServerUnscheduleDeletion,
	}

	if err := registerOperationServersServerUnscheduleDeletionParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServersServerUnscheduleDeletion uses cmd flags to call endpoint api
func runOperationServersServerUnscheduleDeletion(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := servers.NewServerUnscheduleDeletionParams()
	if err, _ := retrieveOperationServersServerUnscheduleDeletionAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersServerUnscheduleDeletionServerIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.Servers.ServerUnscheduleDeletion(params, nil)
	if err != nil {
		api.RenderErrorOutput(err)
		return nil
	}

	msgStr, err := parseOperationServersServerUnscheduleDeletionResult(result)
	if err != nil {
		return err
	}
	if !debug {

		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationServersServerUnscheduleDeletionParamFlags registers all flags needed to fill params
func registerOperationServersServerUnscheduleDeletionParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServersServerUnscheduleDeletionAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersServerUnscheduleDeletionServerIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServersServerUnscheduleDeletionAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServersServerUnscheduleDeletionServerIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	serverIdDescription := `Required. `

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

func retrieveOperationServersServerUnscheduleDeletionAPIVersionFlag(m *servers.ServerUnscheduleDeletionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServersServerUnscheduleDeletionServerIDFlag(m *servers.ServerUnscheduleDeletionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationServersServerUnscheduleDeletionResult parses request result and return the string content
func parseOperationServersServerUnscheduleDeletionResult(resp0 *servers.ServerUnscheduleDeletionNoContent) (string, error) {
	// warning: non schema response serverUnscheduleDeletionNoContent is not supported by go-swagger cli yet.

	return "", nil
}
