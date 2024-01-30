// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/latitudesh/lsh/client/api_keys"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationAPIKeysGetAPIKeysCmd returns a cmd to handle operation getApiKeys
func makeOperationAPIKeysGetAPIKeysCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: `Returns a list of all API keys from the team members`,
		RunE:  runOperationAPIKeysGetAPIKeys,
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
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.APIKeys.GetAPIKeys(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	msgStr, err := parseOperationAPIKeysGetAPIKeysResult(result)
	if err != nil {
		return err
	}
	if !debug {
		utils.PrintResult(msgStr)
	}
	return nil
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
