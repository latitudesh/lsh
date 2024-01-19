// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/cli/models"

	"github.com/spf13/cobra"
)

// Schema cli for CreateAPIKey

// register flags to command
func registerModelCreateAPIKeyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateAPIKeyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateAPIKeyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelCreateAPIKeyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateAPIKeyFlags(depth int, m *models.CreateAPIKey, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveCreateAPIKeyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveCreateAPIKeyDataFlags(depth int, m *models.CreateAPIKey, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data CreateAPIKeyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.CreateAPIKeyData{}
	}

	err, dataAdded := retrieveModelCreateAPIKeyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for CreateAPIKeyData

// register flags to command
func registerModelCreateAPIKeyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateAPIKeyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateAPIKeyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateAPIKeyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelCreateAPIKeyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateAPIKeyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["api_keys"]. Required. `

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
			if err := json.Unmarshal([]byte(`["api_keys"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateAPIKeyDataFlags(depth int, m *models.CreateAPIKeyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveCreateAPIKeyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, typeAdded := retrieveCreateAPIKeyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveCreateAPIKeyDataAttributesFlags(depth int, m *models.CreateAPIKeyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes CreateAPIKeyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.CreateAPIKeyDataAttributes{}
	}

	err, attributesAdded := retrieveModelCreateAPIKeyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveCreateAPIKeyDataTypeFlags(depth int, m *models.CreateAPIKeyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for CreateAPIKeyDataAttributes

// register flags to command
func registerModelCreateAPIKeyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateAPIKeyDataAttributesAPIVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateAPIKeyDataAttributesName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateAPIKeyDataAttributesAPIVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	apiVersionDescription := `The API version to use`

	var apiVersionFlagName string
	if cmdPrefix == "" {
		apiVersionFlagName = "api_version"
	} else {
		apiVersionFlagName = fmt.Sprintf("%v.api_version", cmdPrefix)
	}

	var apiVersionFlagDefault string

	_ = cmd.PersistentFlags().String(apiVersionFlagName, apiVersionFlagDefault, apiVersionDescription)

	return nil
}

func registerCreateAPIKeyDataAttributesName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. Name of the API Key`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)
	cmd.MarkPersistentFlagRequired(nameFlagName)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateAPIKeyDataAttributesFlags(depth int, m *models.CreateAPIKeyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, apiVersionAdded := retrieveCreateAPIKeyDataAttributesAPIVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || apiVersionAdded

	err, nameAdded := retrieveCreateAPIKeyDataAttributesNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrieveCreateAPIKeyDataAttributesAPIVersionFlags(depth int, m *models.CreateAPIKeyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	apiVersionFlagName := fmt.Sprintf("%v.api_version", cmdPrefix)
	if cmd.Flags().Changed(apiVersionFlagName) {

		var apiVersionFlagName string
		if cmdPrefix == "" {
			apiVersionFlagName = "api_version"
		} else {
			apiVersionFlagName = fmt.Sprintf("%v.api_version", cmdPrefix)
		}

		apiVersionFlagValue, err := cmd.Flags().GetString(apiVersionFlagName)
		if err != nil {
			return err, false
		}
		m.APIVersion = apiVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateAPIKeyDataAttributesNameFlags(depth int, m *models.CreateAPIKeyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
