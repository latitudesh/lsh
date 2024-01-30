package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/lsh/client/projects"
	"github.com/latitudesh/lsh/internal/utils"
	"github.com/latitudesh/lsh/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationProjectsUpdateProjectCmd returns a cmd to handle operation updateProject
func makeOperationProjectsUpdateProjectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "update",
		Short: `Updates a project`,
		RunE:  runOperationProjectsUpdateProject,
	}

	if err := registerOperationProjectsUpdateProjectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationProjectsUpdateProject uses cmd flags to call endpoint api
func runOperationProjectsUpdateProject(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := projects.NewUpdateProjectParams()
	if err, _ := retrieveOperationProjectsUpdateProjectAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectsUpdateProjectBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectsUpdateProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Projects.UpdateProject(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}

// registerOperationProjectsUpdateProjectParamFlags registers all flags needed to fill params
func registerOperationProjectsUpdateProjectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationProjectsUpdateProjectAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectsUpdateProjectBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectsUpdateProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationProjectsUpdateProjectAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationProjectsUpdateProjectBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelUpdateProjectBodyFlags(0, "", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationProjectsUpdateProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idOrSlugDescription := `Required. The project ID or Slug`

	var idOrSlugFlagName string
	if cmdPrefix == "" {
		idOrSlugFlagName = "id_or_slug"
	} else {
		idOrSlugFlagName = fmt.Sprintf("%v.id_or_slug", cmdPrefix)
	}

	var idOrSlugFlagDefault string

	_ = cmd.PersistentFlags().String(idOrSlugFlagName, idOrSlugFlagDefault, idOrSlugDescription)
	cmd.MarkPersistentFlagRequired(idOrSlugFlagName)

	return nil
}

func retrieveOperationProjectsUpdateProjectAPIVersionFlag(m *projects.UpdateProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationProjectsUpdateProjectBodyFlag(m *projects.UpdateProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := projects.UpdateProjectBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in UpdateProjectBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = projects.UpdateProjectBody{}
	}
	err, added := retrieveModelUpdateProjectBodyFlags(0, &bodyValueModel, "", cmd)
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
func retrieveOperationProjectsUpdateProjectIDOrSlugFlag(m *projects.UpdateProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id_or_slug") {

		var idOrSlugFlagName string
		if cmdPrefix == "" {
			idOrSlugFlagName = "id_or_slug"
		} else {
			idOrSlugFlagName = fmt.Sprintf("%v.id_or_slug", cmdPrefix)
		}

		idOrSlugFlagValue, err := cmd.Flags().GetString(idOrSlugFlagName)
		if err != nil {
			return err, false
		}
		m.IDOrSlug = idOrSlugFlagValue

	}
	return nil, retAdded
}

// register flags to command
func registerModelUpdateProjectBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateProjectBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateProjectBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName = ""

	if err := registerModelUpdateProjectParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUpdateProjectBodyFlags(depth int, m *projects.UpdateProjectBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveUpdateProjectBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveUpdateProjectBodyDataFlags(depth int, m *projects.UpdateProjectBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("data")
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data UpdateProjectParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &projects.UpdateProjectParamsBodyData{}
	}

	dataFlagName = ""
	err, dataAdded := retrieveModelUpdateProjectParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
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
func registerModelUpdateProjectOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateProjectOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateProjectOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelProjectFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUpdateProjectOKBodyFlags(depth int, m *projects.UpdateProjectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveUpdateProjectOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveUpdateProjectOKBodyDataFlags(depth int, m *projects.UpdateProjectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data models.Project is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.Project{}
	}

	err, dataAdded := retrieveModelProjectFlags(depth+1, dataFlagValue, dataFlagName, cmd)
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
func registerModelUpdateProjectParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateProjectParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateProjectParamsBodyDataID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateProjectParamsBodyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateProjectParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName = ""

	if err := registerModelUpdateProjectParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateProjectParamsBodyDataID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerUpdateProjectParamsBodyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["projects"]. Required. `

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)
	cmd.MarkPersistentFlagRequired(typeFlagName)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["projects"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUpdateProjectParamsBodyDataFlags(depth int, m *projects.UpdateProjectParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveUpdateProjectParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveUpdateProjectParamsBodyDataIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, typeAdded := retrieveUpdateProjectParamsBodyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveUpdateProjectParamsBodyDataAttributesFlags(depth int, m *projects.UpdateProjectParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%vattributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes UpdateProjectParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &projects.UpdateProjectParamsBodyDataAttributes{}
	}

	attributesFlagName = ""
	err, attributesAdded := retrieveModelUpdateProjectParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveUpdateProjectParamsBodyDataIDFlags(depth int, m *projects.UpdateProjectParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var idFlagName = "id"
	if cmd.Flags().Changed(idFlagName) {

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateProjectParamsBodyDataTypeFlags(depth int, m *projects.UpdateProjectParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var typeFlagName = "type"
	if cmd.Flags().Changed(typeFlagName) {

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
func registerModelUpdateProjectParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateProjectParamsBodyDataAttributesBandwidthAlert(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateProjectParamsBodyDataAttributesDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateProjectParamsBodyDataAttributesEnvironment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateProjectParamsBodyDataAttributesName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateProjectParamsBodyDataAttributesBandwidthAlert(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	bandwidthAlertDescription := ``

	var bandwidthAlertFlagName string
	if cmdPrefix == "" {
		bandwidthAlertFlagName = "bandwidth_alert"
	} else {
		bandwidthAlertFlagName = fmt.Sprintf("%v.bandwidth_alert", cmdPrefix)
	}

	var bandwidthAlertFlagDefault bool

	_ = cmd.PersistentFlags().Bool(bandwidthAlertFlagName, bandwidthAlertFlagDefault, bandwidthAlertDescription)

	return nil
}

func registerUpdateProjectParamsBodyDataAttributesDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := ``

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

func registerUpdateProjectParamsBodyDataAttributesEnvironment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	environmentDescription := `Enum: ["Development","Staging","Production"]. `

	var environmentFlagName string
	if cmdPrefix == "" {
		environmentFlagName = "environment"
	} else {
		environmentFlagName = fmt.Sprintf("%v.environment", cmdPrefix)
	}

	var environmentFlagDefault string

	_ = cmd.PersistentFlags().String(environmentFlagName, environmentFlagDefault, environmentDescription)

	if err := cmd.RegisterFlagCompletionFunc(environmentFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["Development","Staging","Production"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerUpdateProjectParamsBodyDataAttributesName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

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
func retrieveModelUpdateProjectParamsBodyDataAttributesFlags(depth int, m *projects.UpdateProjectParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, bandwidthAlertAdded := retrieveUpdateProjectParamsBodyDataAttributesBandwidthAlertFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bandwidthAlertAdded

	err, descriptionAdded := retrieveUpdateProjectParamsBodyDataAttributesDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, environmentAdded := retrieveUpdateProjectParamsBodyDataAttributesEnvironmentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || environmentAdded

	err, nameAdded := retrieveUpdateProjectParamsBodyDataAttributesNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrieveUpdateProjectParamsBodyDataAttributesBandwidthAlertFlags(depth int, m *projects.UpdateProjectParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bandwidthAlertFlagName := fmt.Sprintf("%v.bandwidth_alert", cmdPrefix)
	if cmd.Flags().Changed(bandwidthAlertFlagName) {

		var bandwidthAlertFlagName = "bandwidth_alert"

		bandwidthAlertFlagValue, err := cmd.Flags().GetBool(bandwidthAlertFlagName)
		if err != nil {
			return err, false
		}
		m.BandwidthAlert = bandwidthAlertFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateProjectParamsBodyDataAttributesDescriptionFlags(depth int, m *projects.UpdateProjectParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName = "description"

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateProjectParamsBodyDataAttributesEnvironmentFlags(depth int, m *projects.UpdateProjectParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	environmentFlagName := fmt.Sprintf("%v.environment", cmdPrefix)
	if cmd.Flags().Changed(environmentFlagName) {

		var environmentFlagName = "environment"

		environmentFlagValue, err := cmd.Flags().GetString(environmentFlagName)
		if err != nil {
			return err, false
		}
		m.Environment = environmentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateProjectParamsBodyDataAttributesNameFlags(depth int, m *projects.UpdateProjectParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var nameFlagName = "name"
	if cmd.Flags().Changed(nameFlagName) {

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
