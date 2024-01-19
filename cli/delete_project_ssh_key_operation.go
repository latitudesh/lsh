// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/cli/client/ssh_keys"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationSSHKeysDeleteProjectSSHKeyCmd returns a cmd to handle operation deleteProjectSshKey
func makeOperationSSHKeysDeleteProjectSSHKeyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "destroy",
		Short: `Allow you remove SSH Keys in a project. Remove a SSH Key from the project won't revoke the SSH Keys access for previously deploy and reinstall actions.
`,
		RunE: runOperationSSHKeysDeleteProjectSSHKey,
	}

	if err := registerOperationSSHKeysDeleteProjectSSHKeyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSSHKeysDeleteProjectSSHKey uses cmd flags to call endpoint api
func runOperationSSHKeysDeleteProjectSSHKey(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ssh_keys.NewDeleteProjectSSHKeyParams()
	if err, _ := retrieveOperationSSHKeysDeleteProjectSSHKeyAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSSHKeysDeleteProjectSSHKeyProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSSHKeysDeleteProjectSSHKeySSHKeyIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSSHKeysDeleteProjectSSHKeyResult(appCli.SSHKeys.DeleteProjectSSHKey(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationSSHKeysDeleteProjectSSHKeyParamFlags registers all flags needed to fill params
func registerOperationSSHKeysDeleteProjectSSHKeyParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSSHKeysDeleteProjectSSHKeyAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysDeleteProjectSSHKeyProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysDeleteProjectSSHKeySSHKeyIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSSHKeysDeleteProjectSSHKeyAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSSHKeysDeleteProjectSSHKeyProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSSHKeysDeleteProjectSSHKeySSHKeyIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sshKeyIdDescription := `Required. `

	var sshKeyIdFlagName string
	if cmdPrefix == "" {
		sshKeyIdFlagName = "ssh_key_id"
	} else {
		sshKeyIdFlagName = fmt.Sprintf("%v.ssh_key_id", cmdPrefix)
	}

	var sshKeyIdFlagDefault string

	_ = cmd.PersistentFlags().String(sshKeyIdFlagName, sshKeyIdFlagDefault, sshKeyIdDescription)

	return nil
}

func retrieveOperationSSHKeysDeleteProjectSSHKeyAPIVersionFlag(m *ssh_keys.DeleteProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSSHKeysDeleteProjectSSHKeyProjectIDOrSlugFlag(m *ssh_keys.DeleteProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSSHKeysDeleteProjectSSHKeySSHKeyIDFlag(m *ssh_keys.DeleteProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSSHKeysDeleteProjectSSHKeyResult parses request result and return the string content
func parseOperationSSHKeysDeleteProjectSSHKeyResult(resp0 *ssh_keys.DeleteProjectSSHKeyOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteProjectSshKeyOK is not supported

		// Non schema case: warning deleteProjectSshKeyNotFound is not supported

		return "", respErr
	}

	// warning: non schema response deleteProjectSshKeyOK is not supported by go-swagger cli yet.

	return "", nil
}
