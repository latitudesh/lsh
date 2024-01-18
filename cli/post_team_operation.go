// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/cli/client/teams"
	"github.com/latitudesh/cli/internal/utils"
	"github.com/latitudesh/cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTeamsPostTeamCmd returns a cmd to handle operation postTeam
func makeOperationTeamsPostTeamCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "post-team",
		Short: ``,
		RunE:  runOperationTeamsPostTeam,
	}

	if err := registerOperationTeamsPostTeamParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTeamsPostTeam uses cmd flags to call endpoint api
func runOperationTeamsPostTeam(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := teams.NewPostTeamParams()
	params.SetDefaults()
	if err, _ := retrieveOperationTeamsPostTeamAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTeamsPostTeamBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTeamsPostTeamResult(appCli.Teams.PostTeam(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(utils.PrettifyJson(msgStr))
	}
	return nil
}

// registerOperationTeamsPostTeamParamFlags registers all flags needed to fill params
func registerOperationTeamsPostTeamParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTeamsPostTeamAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTeamsPostTeamBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTeamsPostTeamAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationTeamsPostTeamBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelPostTeamBodyFlags(0, "postTeamBody", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationTeamsPostTeamAPIVersionFlag(m *teams.PostTeamParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTeamsPostTeamBodyFlag(m *teams.PostTeamParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := teams.PostTeamBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in PostTeamBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = teams.PostTeamBody{}
	}
	err, added := retrieveModelPostTeamBodyFlags(0, &bodyValueModel, "postTeamBody", cmd)
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

// parseOperationTeamsPostTeamResult parses request result and return the string content
func parseOperationTeamsPostTeamResult(resp0 *teams.PostTeamCreated, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*teams.PostTeamCreated)
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
		resp1, ok := iResp1.(*teams.PostTeamNotAcceptable)
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
		resp2, ok := iResp2.(*teams.PostTeamUnprocessableEntity)
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
func registerModelPostTeamBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostTeamBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostTeamBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelPostTeamParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostTeamBodyFlags(depth int, m *teams.PostTeamBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrievePostTeamBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrievePostTeamBodyDataFlags(depth int, m *teams.PostTeamBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data PostTeamParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &teams.PostTeamParamsBodyData{}
	}

	err, dataAdded := retrieveModelPostTeamParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
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
func registerModelPostTeamCreatedBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostTeamCreatedBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostTeamCreatedBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelTeamFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostTeamCreatedBodyFlags(depth int, m *teams.PostTeamCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrievePostTeamCreatedBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrievePostTeamCreatedBodyDataFlags(depth int, m *teams.PostTeamCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data models.Team is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.Team{}
	}

	err, dataAdded := retrieveModelTeamFlags(depth+1, dataFlagValue, dataFlagName, cmd)
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
func registerModelPostTeamParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostTeamParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPostTeamParamsBodyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostTeamParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelPostTeamParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostTeamParamsBodyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["teams"]. Required. `

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
			if err := json.Unmarshal([]byte(`["teams"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostTeamParamsBodyDataFlags(depth int, m *teams.PostTeamParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrievePostTeamParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, typeAdded := retrievePostTeamParamsBodyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrievePostTeamParamsBodyDataAttributesFlags(depth int, m *teams.PostTeamParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes PostTeamParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &teams.PostTeamParamsBodyDataAttributes{}
	}

	err, attributesAdded := retrieveModelPostTeamParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrievePostTeamParamsBodyDataTypeFlags(depth int, m *teams.PostTeamParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func registerModelPostTeamParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostTeamParamsBodyDataAttributesAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPostTeamParamsBodyDataAttributesCurrency(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPostTeamParamsBodyDataAttributesDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPostTeamParamsBodyDataAttributesName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPostTeamParamsBodyDataAttributesReferredCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostTeamParamsBodyDataAttributesAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	addressDescription := ``

	var addressFlagName string
	if cmdPrefix == "" {
		addressFlagName = "address"
	} else {
		addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
	}

	var addressFlagDefault string

	_ = cmd.PersistentFlags().String(addressFlagName, addressFlagDefault, addressDescription)

	return nil
}

func registerPostTeamParamsBodyDataAttributesCurrency(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	currencyDescription := `Enum: ["USD","BRL"]. Required. `

	var currencyFlagName string
	if cmdPrefix == "" {
		currencyFlagName = "currency"
	} else {
		currencyFlagName = fmt.Sprintf("%v.currency", cmdPrefix)
	}

	var currencyFlagDefault string

	_ = cmd.PersistentFlags().String(currencyFlagName, currencyFlagDefault, currencyDescription)

	if err := cmd.RegisterFlagCompletionFunc(currencyFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["USD","BRL"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerPostTeamParamsBodyDataAttributesDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerPostTeamParamsBodyDataAttributesName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. `

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

func registerPostTeamParamsBodyDataAttributesReferredCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	referredCodeDescription := `Supported only for first team creation`

	var referredCodeFlagName string
	if cmdPrefix == "" {
		referredCodeFlagName = "referred_code"
	} else {
		referredCodeFlagName = fmt.Sprintf("%v.referred_code", cmdPrefix)
	}

	var referredCodeFlagDefault string

	_ = cmd.PersistentFlags().String(referredCodeFlagName, referredCodeFlagDefault, referredCodeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostTeamParamsBodyDataAttributesFlags(depth int, m *teams.PostTeamParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressAdded := retrievePostTeamParamsBodyDataAttributesAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, currencyAdded := retrievePostTeamParamsBodyDataAttributesCurrencyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || currencyAdded

	err, descriptionAdded := retrievePostTeamParamsBodyDataAttributesDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, nameAdded := retrievePostTeamParamsBodyDataAttributesNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, referredCodeAdded := retrievePostTeamParamsBodyDataAttributesReferredCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || referredCodeAdded

	return nil, retAdded
}

func retrievePostTeamParamsBodyDataAttributesAddressFlags(depth int, m *teams.PostTeamParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addressFlagName := fmt.Sprintf("%v.address", cmdPrefix)
	if cmd.Flags().Changed(addressFlagName) {

		var addressFlagName string
		if cmdPrefix == "" {
			addressFlagName = "address"
		} else {
			addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
		}

		addressFlagValue, err := cmd.Flags().GetString(addressFlagName)
		if err != nil {
			return err, false
		}
		m.Address = addressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePostTeamParamsBodyDataAttributesCurrencyFlags(depth int, m *teams.PostTeamParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	currencyFlagName := fmt.Sprintf("%v.currency", cmdPrefix)
	if cmd.Flags().Changed(currencyFlagName) {

		var currencyFlagName string
		if cmdPrefix == "" {
			currencyFlagName = "currency"
		} else {
			currencyFlagName = fmt.Sprintf("%v.currency", cmdPrefix)
		}

		currencyFlagValue, err := cmd.Flags().GetString(currencyFlagName)
		if err != nil {
			return err, false
		}
		m.Currency = &currencyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePostTeamParamsBodyDataAttributesDescriptionFlags(depth int, m *teams.PostTeamParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePostTeamParamsBodyDataAttributesNameFlags(depth int, m *teams.PostTeamParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePostTeamParamsBodyDataAttributesReferredCodeFlags(depth int, m *teams.PostTeamParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	referredCodeFlagName := fmt.Sprintf("%v.referred_code", cmdPrefix)
	if cmd.Flags().Changed(referredCodeFlagName) {

		var referredCodeFlagName string
		if cmdPrefix == "" {
			referredCodeFlagName = "referred_code"
		} else {
			referredCodeFlagName = fmt.Sprintf("%v.referred_code", cmdPrefix)
		}

		referredCodeFlagValue, err := cmd.Flags().GetString(referredCodeFlagName)
		if err != nil {
			return err, false
		}
		m.ReferredCode = referredCodeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
