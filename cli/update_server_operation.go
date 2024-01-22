// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/servers"
	"github.com/latitudesh/cli/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServersUpdateServerCmd returns a cmd to handle operation updateServer
func makeOperationServersUpdateServerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "update",
		Short: `Update server information.`,
		RunE:  runOperationServersUpdateServer,
	}

	if err := registerOperationServersUpdateServerParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServersUpdateServer uses cmd flags to call endpoint api
func runOperationServersUpdateServer(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := servers.NewUpdateServerParams()
	if err, _ := retrieveOperationServersUpdateServerAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersUpdateServerBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServersUpdateServerServerIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServersUpdateServerResult(appCli.Servers.UpdateServer(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationServersUpdateServerParamFlags registers all flags needed to fill params
func registerOperationServersUpdateServerParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServersUpdateServerAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersUpdateServerBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServersUpdateServerServerIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServersUpdateServerAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServersUpdateServerBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
if cmdPrefix == "" {
	bodyFlagName = "body"
} else {
	bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelUpdateServerBodyFlags(0, "updateServerBody", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationServersUpdateServerServerIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	serverIdDescription := `Required. `
	
	var serverIdFlagName string
	if cmdPrefix == "" {
	serverIdFlagName = "server_id"
	} else {
	serverIdFlagName = fmt.Sprintf("%v.server_id", cmdPrefix)
	}

	var serverIdFlagDefault string

	_ = cmd.PersistentFlags().String(serverIdFlagName, serverIdFlagDefault, serverIdDescription)

	return nil
}

func retrieveOperationServersUpdateServerAPIVersionFlag(m *servers.UpdateServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServersUpdateServerBodyFlag(m *servers.UpdateServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := servers.UpdateServerBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in UpdateServerBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = servers.UpdateServerBody{}
	}
	err, added := retrieveModelUpdateServerBodyFlags(0, &bodyValueModel, "updateServerBody", cmd)
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
func retrieveOperationServersUpdateServerServerIDFlag(m *servers.UpdateServerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("server_id") {

		var serverIdFlagName string
		if cmdPrefix == "" {
			serverIdFlagName = "server_id"
		} else {
			serverIdFlagName = fmt.Sprintf("%v.server_id", cmdPrefix)
		}

		serverIdFlagValue, err := cmd.Flags().GetString(serverIdFlagName)
		if err != nil {
			return err, false
		}
		m.ServerID = serverIdFlagValue

	}
	return nil, retAdded
}

// parseOperationServersUpdateServerResult parses request result and return the string content
func parseOperationServersUpdateServerResult(resp0 *servers.UpdateServerOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*servers.UpdateServerOK)
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
		resp1, ok := iResp1.(*servers.UpdateServerBadRequest)
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
		resp2, ok := iResp2.(*servers.UpdateServerUnprocessableEntity)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
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
func registerModelUpdateServerBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateServerBodyAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateServerBodyID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateServerBodyType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateServerBodyAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelUpdateServerParamsBodyAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateServerBodyID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

func registerUpdateServerBodyType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["servers"]. `

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
			if err := json.Unmarshal([]byte(`["servers"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUpdateServerBodyFlags(depth int, m *servers.UpdateServerBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveUpdateServerBodyAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveUpdateServerBodyIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, typeAdded := retrieveUpdateServerBodyTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveUpdateServerBodyAttributesFlags(depth int, m *servers.UpdateServerBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes UpdateServerParamsBodyAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &servers.UpdateServerParamsBodyAttributes{}
	}

	err, attributesAdded := retrieveModelUpdateServerParamsBodyAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveUpdateServerBodyIDFlags(depth int, m *servers.UpdateServerBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateServerBodyTypeFlags(depth int, m *servers.UpdateServerBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// register flags to command
func registerModelUpdateServerParamsBodyAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUpdateServerParamsBodyAttributesBilling(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUpdateServerParamsBodyAttributesHostname(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUpdateServerParamsBodyAttributesBilling(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	billingDescription := ``

	var billingFlagName string
	if cmdPrefix == "" {
		billingFlagName = "billing"
	} else {
		billingFlagName = fmt.Sprintf("%v.billing", cmdPrefix)
	}

	var billingFlagDefault string

	_ = cmd.PersistentFlags().String(billingFlagName, billingFlagDefault, billingDescription)

	return nil
}

func registerUpdateServerParamsBodyAttributesHostname(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	hostnameDescription := ``

	var hostnameFlagName string
	if cmdPrefix == "" {
		hostnameFlagName = "hostname"
	} else {
		hostnameFlagName = fmt.Sprintf("%v.hostname", cmdPrefix)
	}

	var hostnameFlagDefault string

	_ = cmd.PersistentFlags().String(hostnameFlagName, hostnameFlagDefault, hostnameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUpdateServerParamsBodyAttributesFlags(depth int, m *servers.UpdateServerParamsBodyAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, billingAdded := retrieveUpdateServerParamsBodyAttributesBillingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || billingAdded

	err, hostnameAdded := retrieveUpdateServerParamsBodyAttributesHostnameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostnameAdded

	return nil, retAdded
}

func retrieveUpdateServerParamsBodyAttributesBillingFlags(depth int, m *servers.UpdateServerParamsBodyAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	billingFlagName := fmt.Sprintf("%v.billing", cmdPrefix)
	if cmd.Flags().Changed(billingFlagName) {

		var billingFlagName string
		if cmdPrefix == "" {
			billingFlagName = "billing"
		} else {
			billingFlagName = fmt.Sprintf("%v.billing", cmdPrefix)
		}

		billingFlagValue, err := cmd.Flags().GetString(billingFlagName)
		if err != nil {
			return err, false
		}
		m.Billing = billingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUpdateServerParamsBodyAttributesHostnameFlags(depth int, m *servers.UpdateServerParamsBodyAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hostnameFlagName := fmt.Sprintf("%v.hostname", cmdPrefix)
	if cmd.Flags().Changed(hostnameFlagName) {

		var hostnameFlagName string
		if cmdPrefix == "" {
			hostnameFlagName = "hostname"
		} else {
			hostnameFlagName = fmt.Sprintf("%v.hostname", cmdPrefix)
		}

		hostnameFlagValue, err := cmd.Flags().GetString(hostnameFlagName)
		if err != nil {
			return err, false
		}
		m.Hostname = hostnameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
