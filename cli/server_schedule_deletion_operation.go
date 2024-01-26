// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServersServerScheduleDeletionCmd returns a cmd to handle operation serverScheduleDeletion
func makeOperationServersServerScheduleDeletionCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "schedule-deletion",
		Short: `Schedules the server to be removed at the end of the billing cycle.`,
		RunE:  runOperationServersServerScheduleDeletion,
	}

	if err := registerOperationServersServerScheduleDeletionParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServersServerScheduleDeletion uses cmd flags to call endpoint api
func runOperationServersServerScheduleDeletion(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := servers.NewServerScheduleDeletionParams()
	if err, _ := retrieveOperationServersServerScheduleDeletionAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersServerScheduleDeletionServerIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.Servers.ServerScheduleDeletion(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	msgStr, err := parseOperationServersServerScheduleDeletionResult(result)
	if err != nil {
		return err
	}
	if !debug {

		utils.PrintResult(msgStr, "table")
	}
	return nil
}

// registerOperationServersServerScheduleDeletionParamFlags registers all flags needed to fill params
func registerOperationServersServerScheduleDeletionParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServersServerScheduleDeletionAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersServerScheduleDeletionServerIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServersServerScheduleDeletionAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServersServerScheduleDeletionServerIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationServersServerScheduleDeletionAPIVersionFlag(m *servers.ServerScheduleDeletionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServersServerScheduleDeletionServerIDFlag(m *servers.ServerScheduleDeletionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationServersServerScheduleDeletionResult parses request result and return the string content
func parseOperationServersServerScheduleDeletionResult(resp0 *servers.ServerScheduleDeletionCreated) (string, error) {
	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
