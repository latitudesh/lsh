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
	"github.com/latitudesh/cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationAPIKeysUpdateAPIKeyCmd returns a cmd to handle operation updateApiKey
func makeOperationAPIKeysUpdateAPIKeyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "update-api-key",
		Short: `Regenerate an existing API Key that is tied to the current user. This overrides the previous key.
`,
		RunE: runOperationAPIKeysUpdateAPIKey,
	}

	if err := registerOperationAPIKeysUpdateAPIKeyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAPIKeysUpdateAPIKey uses cmd flags to call endpoint api
func runOperationAPIKeysUpdateAPIKey(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := api_keys.NewUpdateAPIKeyParams()
	if err, _ := retrieveOperationAPIKeysUpdateAPIKeyAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAPIKeysUpdateAPIKeyBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAPIKeysUpdateAPIKeyIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.APIKeys.UpdateAPIKey(params, nil)
	if err != nil {
		api.RenderErrorOutput(err)
		return nil
	}

	msgStr, err := parseOperationAPIKeysUpdateAPIKeyResult(result)
	if err != nil {
		return err
	}
	if !debug {
		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationAPIKeysUpdateAPIKeyParamFlags registers all flags needed to fill params
func registerOperationAPIKeysUpdateAPIKeyParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAPIKeysUpdateAPIKeyAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAPIKeysUpdateAPIKeyBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAPIKeysUpdateAPIKeyIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAPIKeysUpdateAPIKeyAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationAPIKeysUpdateAPIKeyBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelUpdateAPIKeyFlags(0, "updateApiKey", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationAPIKeysUpdateAPIKeyIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)
	cmd.MarkPersistentFlagRequired(idFlagName)

	return nil
}

func retrieveOperationAPIKeysUpdateAPIKeyAPIVersionFlag(m *api_keys.UpdateAPIKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationAPIKeysUpdateAPIKeyBodyFlag(m *api_keys.UpdateAPIKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.UpdateAPIKey{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.UpdateAPIKey: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.UpdateAPIKey{}
	}
	err, added := retrieveModelUpdateAPIKeyFlags(0, bodyValueModel, "updateApiKey", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationAPIKeysUpdateAPIKeyIDFlag(m *api_keys.UpdateAPIKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationAPIKeysUpdateAPIKeyResult parses request result and return the string content
func parseOperationAPIKeysUpdateAPIKeyResult(resp0 *api_keys.UpdateAPIKeyOK) (string, error) {
	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

// register flags to command
func registerModelUpdateAPIKeyOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateAPIKeyOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateAPIKeyOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelAPIKeyFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUpdateAPIKeyOKBodyFlags(depth int, m *api_keys.UpdateAPIKeyOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveUpdateAPIKeyOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveUpdateAPIKeyOKBodyDataFlags(depth int, m *api_keys.UpdateAPIKeyOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data models.APIKey is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.APIKey{}
	}

	err, dataAdded := retrieveModelAPIKeyFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}
