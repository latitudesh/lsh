// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/api"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServersGetServerCmd returns a cmd to handle operation getServer
func makeOperationServersGetServerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get",
		Short: `Returns a server that belongs to the team.`,
		RunE:  runOperationServersGetServer,
	}

	if err := registerOperationServersGetServerParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServersGetServer uses cmd flags to call endpoint api
func runOperationServersGetServer(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := servers.NewGetServerParams()
	if err, _ := retrieveOperationServersGetServerAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServerExtraFieldsServersFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersGetServerServerIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.Servers.GetServer(params, nil)
	if err != nil {
		api.RenderErrorOutput(err)
		return nil
	}

	msgStr, err := parseOperationServersGetServerResult(result)
	if err != nil {
		return err
	}
	if !debug {

		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationServersGetServerParamFlags registers all flags needed to fill params
func registerOperationServersGetServerParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServersGetServerAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServerExtraFieldsServersParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersGetServerServerIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServersGetServerAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServersGetServerExtraFieldsServersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServersGetServerServerIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	serverIdDescription := `Required. The Server ID`

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

func retrieveOperationServersGetServerAPIVersionFlag(m *servers.GetServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServersGetServerExtraFieldsServersFlag(m *servers.GetServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServersGetServerServerIDFlag(m *servers.GetServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationServersGetServerResult parses request result and return the string content
func parseOperationServersGetServerResult(resp0 *servers.GetServerOK) (string, error) {
	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
