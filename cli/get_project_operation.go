// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/projects"
	"github.com/latitudesh/cli/internal/utils"
	"github.com/latitudesh/cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationProjectsGetProjectCmd returns a cmd to handle operation getProject
func makeOperationProjectsGetProjectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get-project",
		Short: ``,
		RunE:  runOperationProjectsGetProject,
	}

	if err := registerOperationProjectsGetProjectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationProjectsGetProject uses cmd flags to call endpoint api
func runOperationProjectsGetProject(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := projects.NewGetProjectParams()
	params.SetDefaults()
	if err, _ := retrieveOperationProjectsGetProjectAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectsGetProjectExtraFieldsProjectsFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationProjectsGetProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationProjectsGetProjectResult(appCli.Projects.GetProject(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationProjectsGetProjectParamFlags registers all flags needed to fill params
func registerOperationProjectsGetProjectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationProjectsGetProjectAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectsGetProjectExtraFieldsProjectsParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationProjectsGetProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationProjectsGetProjectAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationProjectsGetProjectExtraFieldsProjectsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	extraFieldsProjectsDescription := `The ` + "`" + `last_renewal_date` + "`" + ` and ` + "`" + `next_renewal_date` + "`" + ` are provided as extra attributes that show previous and future billing cycle dates. To request it, just set ` + "`" + `extra_fields[projects]=last_renewal_date,next_renewal_date` + "`" + ` in the query string.`

	var extraFieldsProjectsFlagName string
	if cmdPrefix == "" {
		extraFieldsProjectsFlagName = "extra_fields[projects]"
	} else {
		extraFieldsProjectsFlagName = fmt.Sprintf("%v.extra_fields[projects]", cmdPrefix)
	}

	var extraFieldsProjectsFlagDefault string

	_ = cmd.PersistentFlags().String(extraFieldsProjectsFlagName, extraFieldsProjectsFlagDefault, extraFieldsProjectsDescription)

	return nil
}
func registerOperationProjectsGetProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idOrSlugDescription := `Required. The project ID or Slug`

	var idOrSlugFlagName string
	if cmdPrefix == "" {
		idOrSlugFlagName = "id_or_slug"
	} else {
		idOrSlugFlagName = fmt.Sprintf("%v.id_or_slug", cmdPrefix)
	}

	var idOrSlugFlagDefault string

	_ = cmd.PersistentFlags().String(idOrSlugFlagName, idOrSlugFlagDefault, idOrSlugDescription)

	return nil
}

func retrieveOperationProjectsGetProjectAPIVersionFlag(m *projects.GetProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationProjectsGetProjectExtraFieldsProjectsFlag(m *projects.GetProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("extra_fields[projects]") {

		var extraFieldsProjectsFlagName string
		if cmdPrefix == "" {
			extraFieldsProjectsFlagName = "extra_fields[projects]"
		} else {
			extraFieldsProjectsFlagName = fmt.Sprintf("%v.extra_fields[projects]", cmdPrefix)
		}

		extraFieldsProjectsFlagValue, err := cmd.Flags().GetString(extraFieldsProjectsFlagName)
		if err != nil {
			return err, false
		}
		m.ExtraFieldsProjects = &extraFieldsProjectsFlagValue

	}
	return nil, retAdded
}
func retrieveOperationProjectsGetProjectIDOrSlugFlag(m *projects.GetProjectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationProjectsGetProjectResult parses request result and return the string content
func parseOperationProjectsGetProjectResult(resp0 *projects.GetProjectOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*projects.GetProjectOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*projects.GetProjectNotFound)
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
func registerModelGetProjectOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetProjectOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetProjectOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelGetProjectOKBodyFlags(depth int, m *projects.GetProjectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveGetProjectOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetProjectOKBodyDataFlags(depth int, m *projects.GetProjectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
