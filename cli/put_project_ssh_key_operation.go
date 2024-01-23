// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/ssh_keys"
	"github.com/latitudesh/cli/internal"
	"github.com/latitudesh/cli/internal/utils"
	"github.com/latitudesh/cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSSHKeysPutProjectSSHKeyCmd returns a cmd to handle operation putProjectSshKey
func makeOperationSSHKeysPutProjectSSHKeyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "put-project-ssh-key",
		Short: `Allow you update SSH Key in a project. These keys can be used to access servers after deploy and reinstall actions.
`,
		RunE: runOperationSSHKeysPutProjectSSHKey,
	}

	if err := registerOperationSSHKeysPutProjectSSHKeyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSSHKeysPutProjectSSHKey uses cmd flags to call endpoint api
func runOperationSSHKeysPutProjectSSHKey(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ssh_keys.NewPutProjectSSHKeyParams()
	if err, _ := retrieveOperationSSHKeysPutProjectSSHKeyAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSSHKeysPutProjectSSHKeyBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSSHKeysPutProjectSSHKeyProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSSHKeysPutProjectSSHKeySSHKeyIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSSHKeysPutProjectSSHKeyResult(appCli.SSHKeys.PutProjectSSHKey(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationSSHKeysPutProjectSSHKeyParamFlags registers all flags needed to fill params
func registerOperationSSHKeysPutProjectSSHKeyParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSSHKeysPutProjectSSHKeyAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysPutProjectSSHKeyBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysPutProjectSSHKeyProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysPutProjectSSHKeySSHKeyIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSSHKeysPutProjectSSHKeyAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSSHKeysPutProjectSSHKeyBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelPutProjectSSHKeyBodyFlags(0, "putProjectSshKeyBody", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationSSHKeysPutProjectSSHKeyProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	projectIdOrSlugDescription := `Required. `

	var projectIdOrSlugFlagName string
	if cmdPrefix == "" {
		projectIdOrSlugFlagName = "project_id_or_slug"
	} else {
		projectIdOrSlugFlagName = fmt.Sprintf("%v.project_id_or_slug", cmdPrefix)
	}

	var projectIdOrSlugFlagDefault string

	_ = cmd.PersistentFlags().String(projectIdOrSlugFlagName, projectIdOrSlugFlagDefault, projectIdOrSlugDescription)
	cmd.MarkPersistentFlagRequired(projectIdOrSlugFlagName)

	return nil
}
func registerOperationSSHKeysPutProjectSSHKeySSHKeyIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sshKeyIdDescription := `Required. `

	var sshKeyIdFlagName string
	if cmdPrefix == "" {
		sshKeyIdFlagName = "ssh_key_id"
	} else {
		sshKeyIdFlagName = fmt.Sprintf("%v.ssh_key_id", cmdPrefix)
	}

	var sshKeyIdFlagDefault string

	_ = cmd.PersistentFlags().String(sshKeyIdFlagName, sshKeyIdFlagDefault, sshKeyIdDescription)
	cmd.MarkPersistentFlagRequired(sshKeyIdFlagName)

	return nil
}

func retrieveOperationSSHKeysPutProjectSSHKeyAPIVersionFlag(m *ssh_keys.PutProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSSHKeysPutProjectSSHKeyBodyFlag(m *ssh_keys.PutProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := ssh_keys.PutProjectSSHKeyBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in PutProjectSSHKeyBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = ssh_keys.PutProjectSSHKeyBody{}
	}
	err, added := retrieveModelPutProjectSSHKeyBodyFlags(0, &bodyValueModel, "putProjectSshKeyBody", cmd)
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
func retrieveOperationSSHKeysPutProjectSSHKeyProjectIDOrSlugFlag(m *ssh_keys.PutProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSSHKeysPutProjectSSHKeySSHKeyIDFlag(m *ssh_keys.PutProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ssh_key_id") {

		var sshKeyIdFlagName string
		if cmdPrefix == "" {
			sshKeyIdFlagName = "ssh_key_id"
		} else {
			sshKeyIdFlagName = fmt.Sprintf("%v.ssh_key_id", cmdPrefix)
		}

		sshKeyIdFlagValue, err := cmd.Flags().GetString(sshKeyIdFlagName)
		if err != nil {
			return err, false
		}
		m.SSHKeyID = sshKeyIdFlagValue

	}
	return nil, retAdded
}

// parseOperationSSHKeysPutProjectSSHKeyResult parses request result and return the string content
func parseOperationSSHKeysPutProjectSSHKeyResult(resp0 *ssh_keys.PutProjectSSHKeyOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*ssh_keys.PutProjectSSHKeyOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		notFoundErrorMessage, err := internal.ParseNotFoundError(respErr)

		if err != nil {
			return "", err
		}

		if len(notFoundErrorMessage) > 0 {
			return notFoundErrorMessage, nil
		}

		// Non schema case: warning putProjectSshKeyBadRequest is not supported

		// Non schema case: warning putProjectSshKeyUnprocessableEntity is not supported

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
func registerModelPutProjectSSHKeyBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPutProjectSSHKeyBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPutProjectSSHKeyBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelPutProjectSSHKeyParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPutProjectSSHKeyBodyFlags(depth int, m *ssh_keys.PutProjectSSHKeyBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrievePutProjectSSHKeyBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrievePutProjectSSHKeyBodyDataFlags(depth int, m *ssh_keys.PutProjectSSHKeyBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data PutProjectSSHKeyParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &ssh_keys.PutProjectSSHKeyParamsBodyData{}
	}

	err, dataAdded := retrieveModelPutProjectSSHKeyParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
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
func registerModelPutProjectSSHKeyOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPutProjectSSHKeyOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPutProjectSSHKeyOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelPutProjectSSHKeyOKBodyFlags(depth int, m *ssh_keys.PutProjectSSHKeyOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrievePutProjectSSHKeyOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrievePutProjectSSHKeyOKBodyDataFlags(depth int, m *ssh_keys.PutProjectSSHKeyOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func registerModelPutProjectSSHKeyParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPutProjectSSHKeyParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPutProjectSSHKeyParamsBodyDataID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPutProjectSSHKeyParamsBodyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPutProjectSSHKeyParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelPutProjectSSHKeyParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPutProjectSSHKeyParamsBodyDataID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

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

func registerPutProjectSSHKeyParamsBodyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelPutProjectSSHKeyParamsBodyDataFlags(depth int, m *ssh_keys.PutProjectSSHKeyParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrievePutProjectSSHKeyParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrievePutProjectSSHKeyParamsBodyDataIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, typeAdded := retrievePutProjectSSHKeyParamsBodyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrievePutProjectSSHKeyParamsBodyDataAttributesFlags(depth int, m *ssh_keys.PutProjectSSHKeyParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes PutProjectSSHKeyParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &ssh_keys.PutProjectSSHKeyParamsBodyDataAttributes{}
	}

	err, attributesAdded := retrieveModelPutProjectSSHKeyParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrievePutProjectSSHKeyParamsBodyDataIDFlags(depth int, m *ssh_keys.PutProjectSSHKeyParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

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
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePutProjectSSHKeyParamsBodyDataTypeFlags(depth int, m *ssh_keys.PutProjectSSHKeyParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func registerModelPutProjectSSHKeyParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPutProjectSSHKeyParamsBodyDataAttributesName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPutProjectSSHKeyParamsBodyDataAttributesName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPutProjectSSHKeyParamsBodyDataAttributesFlags(depth int, m *ssh_keys.PutProjectSSHKeyParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nameAdded := retrievePutProjectSSHKeyParamsBodyDataAttributesNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrievePutProjectSSHKeyParamsBodyDataAttributesNameFlags(depth int, m *ssh_keys.PutProjectSSHKeyParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
