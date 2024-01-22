// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/api_keys"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationAPIKeysDeleteAPIKeyCmd returns a cmd to handle operation deleteApiKey
func makeOperationAPIKeysDeleteAPIKeyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "destroy",
		Short: `Delete an existing API Key. Once deleted, the API Key can no longer be used to access the API.`,
		RunE: runOperationAPIKeysDeleteAPIKey,
	}

	if err := registerOperationAPIKeysDeleteAPIKeyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAPIKeysDeleteAPIKey uses cmd flags to call endpoint api
func runOperationAPIKeysDeleteAPIKey(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := api_keys.NewDeleteAPIKeyParams()
	if err, _ := retrieveOperationAPIKeysDeleteAPIKeyAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAPIKeysDeleteAPIKeyIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAPIKeysDeleteAPIKeyResult(appCli.APIKeys.DeleteAPIKey(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationAPIKeysDeleteAPIKeyParamFlags registers all flags needed to fill params
func registerOperationAPIKeysDeleteAPIKeyParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAPIKeysDeleteAPIKeyAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAPIKeysDeleteAPIKeyIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAPIKeysDeleteAPIKeyAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationAPIKeysDeleteAPIKeyIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. `

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

func retrieveOperationAPIKeysDeleteAPIKeyAPIVersionFlag(m *api_keys.DeleteAPIKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationAPIKeysDeleteAPIKeyIDFlag(m *api_keys.DeleteAPIKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

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

	}
	return nil, retAdded
}

// parseOperationAPIKeysDeleteAPIKeyResult parses request result and return the string content
func parseOperationAPIKeysDeleteAPIKeyResult(resp0 *api_keys.DeleteAPIKeyOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteApiKeyOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*api_keys.DeleteAPIKeyNotFound)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response deleteApiKeyOK is not supported by go-swagger cli yet.

	return "", nil
}
