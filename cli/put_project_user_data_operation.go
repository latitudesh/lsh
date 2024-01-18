// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/user_data"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUserDataPutProjectUserDataCmd returns a cmd to handle operation putProjectUserData
func makeOperationUserDataPutProjectUserDataCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "put-project-user-data",
		Short: `Allow you update User Data in a project.
`,
		RunE: runOperationUserDataPutProjectUserData,
	}

	if err := registerOperationUserDataPutProjectUserDataParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUserDataPutProjectUserData uses cmd flags to call endpoint api
func runOperationUserDataPutProjectUserData(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := user_data.NewPutProjectUserDataParams()
	params.SetDefaults()
	if err, _ := retrieveOperationUserDataPutProjectUserDataAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserDataPutProjectUserDataBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserDataPutProjectUserDataProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUserDataPutProjectUserDataUserDataIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUserDataPutProjectUserDataResult(appCli.UserData.PutProjectUserData(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationUserDataPutProjectUserDataParamFlags registers all flags needed to fill params
func registerOperationUserDataPutProjectUserDataParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUserDataPutProjectUserDataAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserDataPutProjectUserDataBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserDataPutProjectUserDataProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUserDataPutProjectUserDataUserDataIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUserDataPutProjectUserDataAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationUserDataPutProjectUserDataBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelPutProjectUserDataBodyFlags(0, "putProjectUserDataBody", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationUserDataPutProjectUserDataProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationUserDataPutProjectUserDataUserDataIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	userDataIdDescription := `Required. `

	var userDataIdFlagName string
	if cmdPrefix == "" {
		userDataIdFlagName = "user_data_id"
	} else {
		userDataIdFlagName = fmt.Sprintf("%v.user_data_id", cmdPrefix)
	}

	var userDataIdFlagDefault string

	_ = cmd.PersistentFlags().String(userDataIdFlagName, userDataIdFlagDefault, userDataIdDescription)

	return nil
}

func retrieveOperationUserDataPutProjectUserDataAPIVersionFlag(m *user_data.PutProjectUserDataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationUserDataPutProjectUserDataBodyFlag(m *user_data.PutProjectUserDataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := user_data.PutProjectUserDataBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in PutProjectUserDataBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = user_data.PutProjectUserDataBody{}
	}
	err, added := retrieveModelPutProjectUserDataBodyFlags(0, &bodyValueModel, "putProjectUserDataBody", cmd)
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
func retrieveOperationUserDataPutProjectUserDataProjectIDOrSlugFlag(m *user_data.PutProjectUserDataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationUserDataPutProjectUserDataUserDataIDFlag(m *user_data.PutProjectUserDataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("user_data_id") {

		var userDataIdFlagName string
		if cmdPrefix == "" {
			userDataIdFlagName = "user_data_id"
		} else {
			userDataIdFlagName = fmt.Sprintf("%v.user_data_id", cmdPrefix)
		}

		userDataIdFlagValue, err := cmd.Flags().GetString(userDataIdFlagName)
		if err != nil {
			return err, false
		}
		m.UserDataID = userDataIdFlagValue

	}
	return nil, retAdded
}

// parseOperationUserDataPutProjectUserDataResult parses request result and return the string content
func parseOperationUserDataPutProjectUserDataResult(resp0 *user_data.PutProjectUserDataOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*user_data.PutProjectUserDataOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning putProjectUserDataBadRequest is not supported

		// Non schema case: warning putProjectUserDataNotFound is not supported

		// Non schema case: warning putProjectUserDataUnprocessableEntity is not supported

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
func registerModelPutProjectUserDataBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPutProjectUserDataBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPutProjectUserDataBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelPutProjectUserDataParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPutProjectUserDataBodyFlags(depth int, m *user_data.PutProjectUserDataBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrievePutProjectUserDataBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrievePutProjectUserDataBodyDataFlags(depth int, m *user_data.PutProjectUserDataBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data PutProjectUserDataParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &user_data.PutProjectUserDataParamsBodyData{}
	}

	err, dataAdded := retrieveModelPutProjectUserDataParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
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
func registerModelPutProjectUserDataParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPutProjectUserDataParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPutProjectUserDataParamsBodyDataID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPutProjectUserDataParamsBodyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPutProjectUserDataParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelPutProjectUserDataParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPutProjectUserDataParamsBodyDataID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

	return nil
}

func registerPutProjectUserDataParamsBodyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["user_data"]. Required. `

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
			if err := json.Unmarshal([]byte(`["user_data"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPutProjectUserDataParamsBodyDataFlags(depth int, m *user_data.PutProjectUserDataParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrievePutProjectUserDataParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrievePutProjectUserDataParamsBodyDataIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, typeAdded := retrievePutProjectUserDataParamsBodyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrievePutProjectUserDataParamsBodyDataAttributesFlags(depth int, m *user_data.PutProjectUserDataParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes PutProjectUserDataParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &user_data.PutProjectUserDataParamsBodyDataAttributes{}
	}

	err, attributesAdded := retrieveModelPutProjectUserDataParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrievePutProjectUserDataParamsBodyDataIDFlags(depth int, m *user_data.PutProjectUserDataParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePutProjectUserDataParamsBodyDataTypeFlags(depth int, m *user_data.PutProjectUserDataParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func registerModelPutProjectUserDataParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPutProjectUserDataParamsBodyDataAttributesContent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPutProjectUserDataParamsBodyDataAttributesDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPutProjectUserDataParamsBodyDataAttributesContent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	contentDescription := `encoded content of the User Data`

	var contentFlagName string
	if cmdPrefix == "" {
		contentFlagName = "content"
	} else {
		contentFlagName = fmt.Sprintf("%v.content", cmdPrefix)
	}

	var contentFlagDefault string

	_ = cmd.PersistentFlags().String(contentFlagName, contentFlagDefault, contentDescription)

	return nil
}

func registerPutProjectUserDataParamsBodyDataAttributesDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `description dummy user data`

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPutProjectUserDataParamsBodyDataAttributesFlags(depth int, m *user_data.PutProjectUserDataParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, contentAdded := retrievePutProjectUserDataParamsBodyDataAttributesContentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contentAdded

	err, descriptionAdded := retrievePutProjectUserDataParamsBodyDataAttributesDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	return nil, retAdded
}

func retrievePutProjectUserDataParamsBodyDataAttributesContentFlags(depth int, m *user_data.PutProjectUserDataParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	contentFlagName := fmt.Sprintf("%v.content", cmdPrefix)
	if cmd.Flags().Changed(contentFlagName) {

		var contentFlagName string
		if cmdPrefix == "" {
			contentFlagName = "content"
		} else {
			contentFlagName = fmt.Sprintf("%v.content", cmdPrefix)
		}

		contentFlagValue, err := cmd.Flags().GetString(contentFlagName)
		if err != nil {
			return err, false
		}
		m.Content = contentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePutProjectUserDataParamsBodyDataAttributesDescriptionFlags(depth int, m *user_data.PutProjectUserDataParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}
