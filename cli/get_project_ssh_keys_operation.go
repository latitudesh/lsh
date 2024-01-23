// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/lsh/client/ssh_keys"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSSHKeysGetProjectSSHKeysCmd returns a cmd to handle operation getProjectSshKeys
func makeOperationSSHKeysGetProjectSSHKeysCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "list",
		Short: `List all SSH Keys in the project. These keys can be used to access servers after deploy and reinstall actions.`,
		RunE: runOperationSSHKeysGetProjectSSHKeys,
	}

	if err := registerOperationSSHKeysGetProjectSSHKeysParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSSHKeysGetProjectSSHKeys uses cmd flags to call endpoint api
func runOperationSSHKeysGetProjectSSHKeys(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ssh_keys.NewGetProjectSSHKeysParams()
	if err, _ := retrieveOperationSSHKeysGetProjectSSHKeysAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSSHKeysGetProjectSSHKeysProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSSHKeysGetProjectSSHKeysResult(appCli.SSHKeys.GetProjectSSHKeys(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		utils.PrintOutput(msgStr)
	}
	return nil
}

// registerOperationSSHKeysGetProjectSSHKeysParamFlags registers all flags needed to fill params
func registerOperationSSHKeysGetProjectSSHKeysParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSSHKeysGetProjectSSHKeysAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysGetProjectSSHKeysProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSSHKeysGetProjectSSHKeysAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSSHKeysGetProjectSSHKeysProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSSHKeysGetProjectSSHKeysAPIVersionFlag(m *ssh_keys.GetProjectSSHKeysParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSSHKeysGetProjectSSHKeysProjectIDOrSlugFlag(m *ssh_keys.GetProjectSSHKeysParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSSHKeysGetProjectSSHKeysResult parses request result and return the string content
func parseOperationSSHKeysGetProjectSSHKeysResult(resp0 *ssh_keys.GetProjectSSHKeysOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*ssh_keys.GetProjectSSHKeysOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
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
