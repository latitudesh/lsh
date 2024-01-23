// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/api_keys"
	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationAPIKeysGetAPIKeysCmd returns a cmd to handle operation getApiKeys
func makeOperationAPIKeysGetAPIKeysCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "get-api-keys",
		Short: `Returns a list of all API keys from the team members
`,
		RunE: runOperationAPIKeysGetAPIKeys,
	}

	if err := registerOperationAPIKeysGetAPIKeysParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAPIKeysGetAPIKeys uses cmd flags to call endpoint api
func runOperationAPIKeysGetAPIKeys(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := api_keys.NewGetAPIKeysParams()
	if err, _ := retrieveOperationAPIKeysGetAPIKeysAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.APIKeys.GetAPIKeys(params, nil)
	if err != nil {
		api.RenderErrorOutput(err)
		return nil
	}

	msgStr, err := parseOperationAPIKeysGetAPIKeysResult(result)
	if err != nil {
		return err
	}
	if !debug {
		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationAPIKeysGetAPIKeysParamFlags registers all flags needed to fill params
func registerOperationAPIKeysGetAPIKeysParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAPIKeysGetAPIKeysAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAPIKeysGetAPIKeysAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationAPIKeysGetAPIKeysAPIVersionFlag(m *api_keys.GetAPIKeysParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationAPIKeysGetAPIKeysResult parses request result and return the string content
func parseOperationAPIKeysGetAPIKeysResult(resp0 *api_keys.GetAPIKeysOK) (string, error) {
	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
