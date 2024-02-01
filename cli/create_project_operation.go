// Code generated by go-swagger; DO NOT EDIT.

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

// makeOperationProjectsCreateProjectCmd returns a cmd to handle operation createProject
func makeOperationProjectsCreateProjectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "create",
		Short: `Creates a project`,
		RunE:  runOperationProjectsCreateProject,
	}

	if err := registerOperationProjectsCreateProjectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationProjectsCreateProject uses cmd flags to call endpoint api
func runOperationProjectsCreateProject(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := projects.NewCreateProjectParams()
	if err, _ := retrieveOperationProjectsCreateProjectBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Projects.CreateProject(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}

// registerOperationProjectsCreateProjectParamFlags registers all flags needed to fill params
func registerOperationProjectsCreateProjectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationProjectsCreateProjectBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationProjectsCreateProjectBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelCreateProjectBodyFlags(0, "", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationProjectsCreateProjectBodyFlag(m *projects.CreateProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := projects.CreateProjectBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in CreateProjectBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = projects.CreateProjectBody{}
	}
	err, added := retrieveModelCreateProjectBodyFlags(0, &bodyValueModel, "", cmd)
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

// register flags to command
func registerModelCreateProjectBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateProjectBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateProjectBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName = ""

	if err := registerModelCreateProjectParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateProjectBodyFlags(depth int, m *projects.CreateProjectBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveCreateProjectBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveCreateProjectBodyDataFlags(depth int, m *projects.CreateProjectBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("data")
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data CreateProjectParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &projects.CreateProjectParamsBodyData{}
	}

	dataFlagName = ""
	err, dataAdded := retrieveModelCreateProjectParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
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
func registerModelCreateProjectCreatedBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateProjectCreatedBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateProjectCreatedBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName = ""

	if err := registerModelProjectFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateProjectCreatedBodyFlags(depth int, m *projects.CreateProjectCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveCreateProjectCreatedBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveCreateProjectCreatedBodyDataFlags(depth int, m *projects.CreateProjectCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func registerModelCreateProjectParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateProjectParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateProjectParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName = ""

	if err := registerModelCreateProjectParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateProjectParamsBodyDataFlags(depth int, m *projects.CreateProjectParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveCreateProjectParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	return nil, retAdded
}

func retrieveCreateProjectParamsBodyDataAttributesFlags(depth int, m *projects.CreateProjectParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%vattributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes CreateProjectParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &projects.CreateProjectParamsBodyDataAttributes{}
	}

	attributesFlagName = ""
	err, attributesAdded := retrieveModelCreateProjectParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

// register flags to command
func registerModelCreateProjectParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateProjectParamsBodyDataAttributesDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateProjectParamsBodyDataAttributesEnvironment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateProjectParamsBodyDataAttributesName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateProjectParamsBodyDataAttributesProvisioningType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateProjectParamsBodyDataAttributesDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `The project description.`

	var descriptionFlagName = "description"

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerCreateProjectParamsBodyDataAttributesEnvironment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	environmentDescription := `Enum: ["Development","Staging","Production"]. `

	var environmentFlagName = "environment"

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

func registerCreateProjectParamsBodyDataAttributesName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. The project name. Must be unique.`

	var nameFlagName = "name"

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)
	cmd.MarkPersistentFlagRequired(nameFlagName)

	return nil
}

func registerCreateProjectParamsBodyDataAttributesProvisioningType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	provisioningTypeDescription := `Enum: ["reserved","on_demand"]. Required. The provisioning type of the project. Default: on_demand`

	var provisioningTypeFlagName = "provisioning_type"

	_ = cmd.PersistentFlags().String(provisioningTypeFlagName, "on_demand", provisioningTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(provisioningTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["reserved","on_demand"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateProjectParamsBodyDataAttributesFlags(depth int, m *projects.CreateProjectParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveCreateProjectParamsBodyDataAttributesDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, environmentAdded := retrieveCreateProjectParamsBodyDataAttributesEnvironmentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || environmentAdded

	err, nameAdded := retrieveCreateProjectParamsBodyDataAttributesNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, provisioningTypeAdded := retrieveCreateProjectParamsBodyDataAttributesProvisioningTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || provisioningTypeAdded

	return nil, retAdded
}

func retrieveCreateProjectParamsBodyDataAttributesDescriptionFlags(depth int, m *projects.CreateProjectParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var descriptionFlagName = "description"
	if cmd.Flags().Changed(descriptionFlagName) {

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateProjectParamsBodyDataAttributesEnvironmentFlags(depth int, m *projects.CreateProjectParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var environmentFlagName = "environment"
	if cmd.Flags().Changed(environmentFlagName) {

		environmentFlagValue, err := cmd.Flags().GetString(environmentFlagName)
		if err != nil {
			return err, false
		}
		m.Environment = environmentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateProjectParamsBodyDataAttributesNameFlags(depth int, m *projects.CreateProjectParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateProjectParamsBodyDataAttributesProvisioningTypeFlags(depth int, m *projects.CreateProjectParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var provisioningTypeFlagName string

	if cmdPrefix == "" {
		provisioningTypeFlagName = "provisioning_type"
	} else {
		provisioningTypeFlagName = fmt.Sprintf("%v.provisioning_type", cmdPrefix)
	}

	provisioningTypeFlagValue, err := cmd.Flags().GetString(provisioningTypeFlagName)
	if err != nil {
		return err, false
	}
	m.ProvisioningType = &provisioningTypeFlagValue

	retAdded = true

	return nil, retAdded
}
