// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/ssh_keys"
	"github.com/latitudesh/cli/internal/utils"
	"github.com/latitudesh/cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSSHKeysPostProjectSSHKeyCmd returns a cmd to handle operation postProjectSshKey
func makeOperationSSHKeysPostProjectSSHKeyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "create",
		Short: `Allow you create SSH Keys in a project. These keys can be used to access servers after deploy and reinstall actions.
`,
		RunE: runOperationSSHKeysPostProjectSSHKey,
	}

	if err := registerOperationSSHKeysPostProjectSSHKeyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSSHKeysPostProjectSSHKey uses cmd flags to call endpoint api
func runOperationSSHKeysPostProjectSSHKey(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ssh_keys.NewPostProjectSSHKeyParams()
	if err, _ := retrieveOperationSSHKeysPostProjectSSHKeyAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSSHKeysPostProjectSSHKeyBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSSHKeysPostProjectSSHKeyProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSSHKeysPostProjectSSHKeyResult(appCli.SSHKeys.PostProjectSSHKey(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationSSHKeysPostProjectSSHKeyParamFlags registers all flags needed to fill params
func registerOperationSSHKeysPostProjectSSHKeyParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSSHKeysPostProjectSSHKeyAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysPostProjectSSHKeyBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysPostProjectSSHKeyProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSSHKeysPostProjectSSHKeyAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSSHKeysPostProjectSSHKeyBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelPostProjectSSHKeyBodyFlags(0, "postProjectSshKeyBody", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationSSHKeysPostProjectSSHKeyProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	projectIdOrSlugDescription := `Required. `

	var projectIdOrSlugFlagName string
	if cmdPrefix == "" {
		projectIdOrSlugFlagName = "project_id_or_slug"
	} else {
		projectIdOrSlugFlagName = fmt.Sprintf("%v.project_id_or_slug", cmdPrefix)
	}

	var projectIdOrSlugFlagDefault string

	_ = cmd.PersistentFlags().String(projectIdOrSlugFlagName, projectIdOrSlugFlagDefault, projectIdOrSlugDescription)

	return nil
}

func retrieveOperationSSHKeysPostProjectSSHKeyAPIVersionFlag(m *ssh_keys.PostProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSSHKeysPostProjectSSHKeyBodyFlag(m *ssh_keys.PostProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := ssh_keys.PostProjectSSHKeyBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in PostProjectSSHKeyBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = ssh_keys.PostProjectSSHKeyBody{}
	}
	err, added := retrieveModelPostProjectSSHKeyBodyFlags(0, &bodyValueModel, "postProjectSshKeyBody", cmd)
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
func retrieveOperationSSHKeysPostProjectSSHKeyProjectIDOrSlugFlag(m *ssh_keys.PostProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("project_id_or_slug") {

		var projectIdOrSlugFlagName string
		if cmdPrefix == "" {
			projectIdOrSlugFlagName = "project_id_or_slug"
		} else {
			projectIdOrSlugFlagName = fmt.Sprintf("%v.project_id_or_slug", cmdPrefix)
		}

		projectIdOrSlugFlagValue, err := cmd.Flags().GetString(projectIdOrSlugFlagName)
		if err != nil {
			return err, false
		}
		m.ProjectIDOrSlug = projectIdOrSlugFlagValue

	}
	return nil, retAdded
}

// parseOperationSSHKeysPostProjectSSHKeyResult parses request result and return the string content
func parseOperationSSHKeysPostProjectSSHKeyResult(resp0 *ssh_keys.PostProjectSSHKeyCreated, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*ssh_keys.PostProjectSSHKeyCreated)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning postProjectSshKeyBadRequest is not supported

		// Non schema case: warning postProjectSshKeyUnprocessableEntity is not supported

		return "", respErr
	}

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
func registerModelPostProjectSSHKeyBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostProjectSSHKeyBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostProjectSSHKeyBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelPostProjectSSHKeyParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostProjectSSHKeyBodyFlags(depth int, m *ssh_keys.PostProjectSSHKeyBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrievePostProjectSSHKeyBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrievePostProjectSSHKeyBodyDataFlags(depth int, m *ssh_keys.PostProjectSSHKeyBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data PostProjectSSHKeyParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &ssh_keys.PostProjectSSHKeyParamsBodyData{}
	}

	err, dataAdded := retrieveModelPostProjectSSHKeyParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// register flags to command
func registerModelPostProjectSSHKeyCreatedBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostProjectSSHKeyCreatedBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostProjectSSHKeyCreatedBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelSSHKeyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostProjectSSHKeyCreatedBodyFlags(depth int, m *ssh_keys.PostProjectSSHKeyCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrievePostProjectSSHKeyCreatedBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrievePostProjectSSHKeyCreatedBodyDataFlags(depth int, m *ssh_keys.PostProjectSSHKeyCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data models.SSHKeyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.SSHKeyData{}
	}

	err, dataAdded := retrieveModelSSHKeyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// register flags to command
func registerModelPostProjectSSHKeyParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostProjectSSHKeyParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPostProjectSSHKeyParamsBodyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostProjectSSHKeyParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelPostProjectSSHKeyParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostProjectSSHKeyParamsBodyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["ssh_keys"]. Required. `

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["ssh_keys"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostProjectSSHKeyParamsBodyDataFlags(depth int, m *ssh_keys.PostProjectSSHKeyParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrievePostProjectSSHKeyParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, typeAdded := retrievePostProjectSSHKeyParamsBodyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrievePostProjectSSHKeyParamsBodyDataAttributesFlags(depth int, m *ssh_keys.PostProjectSSHKeyParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes PostProjectSSHKeyParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &ssh_keys.PostProjectSSHKeyParamsBodyDataAttributes{}
	}

	err, attributesAdded := retrieveModelPostProjectSSHKeyParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrievePostProjectSSHKeyParamsBodyDataTypeFlags(depth int, m *ssh_keys.PostProjectSSHKeyParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// register flags to command
func registerModelPostProjectSSHKeyParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostProjectSSHKeyParamsBodyDataAttributesName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPostProjectSSHKeyParamsBodyDataAttributesPublicKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostProjectSSHKeyParamsBodyDataAttributesName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name of the SSH Key`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerPostProjectSSHKeyParamsBodyDataAttributesPublicKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	publicKeyDescription := `SSH Public Key`

	var publicKeyFlagName string
	if cmdPrefix == "" {
		publicKeyFlagName = "public_key"
	} else {
		publicKeyFlagName = fmt.Sprintf("%v.public_key", cmdPrefix)
	}

	var publicKeyFlagDefault string

	_ = cmd.PersistentFlags().String(publicKeyFlagName, publicKeyFlagDefault, publicKeyDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostProjectSSHKeyParamsBodyDataAttributesFlags(depth int, m *ssh_keys.PostProjectSSHKeyParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nameAdded := retrievePostProjectSSHKeyParamsBodyDataAttributesNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, publicKeyAdded := retrievePostProjectSSHKeyParamsBodyDataAttributesPublicKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || publicKeyAdded

	return nil, retAdded
}

func retrievePostProjectSSHKeyParamsBodyDataAttributesNameFlags(depth int, m *ssh_keys.PostProjectSSHKeyParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePostProjectSSHKeyParamsBodyDataAttributesPublicKeyFlags(depth int, m *ssh_keys.PostProjectSSHKeyParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	publicKeyFlagName := fmt.Sprintf("%v.public_key", cmdPrefix)
	if cmd.Flags().Changed(publicKeyFlagName) {

		var publicKeyFlagName string
		if cmdPrefix == "" {
			publicKeyFlagName = "public_key"
		} else {
			publicKeyFlagName = fmt.Sprintf("%v.public_key", cmdPrefix)
		}

		publicKeyFlagValue, err := cmd.Flags().GetString(publicKeyFlagName)
		if err != nil {
			return err, false
		}
		m.PublicKey = publicKeyFlagValue

		retAdded = true
	}

	return nil, retAdded
}
