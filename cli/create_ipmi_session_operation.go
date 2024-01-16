// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/ip_m_i_credentials"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationIPmiCredentialsCreateIpmiSessionCmd returns a cmd to handle operation createIpmiSession
func makeOperationIPmiCredentialsCreateIpmiSessionCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "create-ipmi-session",
		Short: `Generates IPMI credentials for a given server. Remote access creates a VPN connection to the internal network of your server so you can connect to its IPMI.
You will have to use a VPN client such as https://openvpn.net to connect. See ` + "`" + `VPN Sessions` + "`" + ` API to create a VPN connection.

Related guide: https://docs.latitude.sh/docs/ipmi
`,
		RunE: runOperationIPmiCredentialsCreateIpmiSession,
	}

	if err := registerOperationIPmiCredentialsCreateIpmiSessionParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationIPmiCredentialsCreateIpmiSession uses cmd flags to call endpoint api
func runOperationIPmiCredentialsCreateIpmiSession(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ip_m_i_credentials.NewCreateIpmiSessionParams()
	if err, _ := retrieveOperationIPmiCredentialsCreateIpmiSessionAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationIPmiCredentialsCreateIpmiSessionIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationIPmiCredentialsCreateIpmiSessionResult(appCli.IPmiCredentials.CreateIpmiSession(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationIPmiCredentialsCreateIpmiSessionParamFlags registers all flags needed to fill params
func registerOperationIPmiCredentialsCreateIpmiSessionParamFlags(cmd *cobra.Command) error {
	if err := registerOperationIPmiCredentialsCreateIpmiSessionAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationIPmiCredentialsCreateIpmiSessionIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationIPmiCredentialsCreateIpmiSessionAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationIPmiCredentialsCreateIpmiSessionIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationIPmiCredentialsCreateIpmiSessionAPIVersionFlag(m *ip_m_i_credentials.CreateIpmiSessionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationIPmiCredentialsCreateIpmiSessionIDFlag(m *ip_m_i_credentials.CreateIpmiSessionParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationIPmiCredentialsCreateIpmiSessionResult parses request result and return the string content
func parseOperationIPmiCredentialsCreateIpmiSessionResult(resp0 *ip_m_i_credentials.CreateIpmiSessionCreated, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*ip_m_i_credentials.CreateIpmiSessionCreated)
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
		resp1, ok := iResp1.(*ip_m_i_credentials.CreateIpmiSessionForbidden)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*ip_m_i_credentials.CreateIpmiSessionNotFound)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*ip_m_i_credentials.CreateIpmiSessionUnprocessableEntity)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
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
