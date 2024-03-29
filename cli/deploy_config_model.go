// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/models"

	"github.com/spf13/cobra"
)

// Schema cli for DeployConfig

// register flags to command
func registerModelDeployConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeployConfigData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeployConfigData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelDeployConfigDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeployConfigFlags(depth int, m *models.DeployConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveDeployConfigDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveDeployConfigDataFlags(depth int, m *models.DeployConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data DeployConfigData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.DeployConfigData{}
	}

	err, dataAdded := retrieveModelDeployConfigDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for DeployConfigData

// register flags to command
func registerModelDeployConfigDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeployConfigDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeployConfigDataID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeployConfigDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelDeployConfigDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeployConfigDataID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeployConfigDataFlags(depth int, m *models.DeployConfigData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveDeployConfigDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveDeployConfigDataIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveDeployConfigDataAttributesFlags(depth int, m *models.DeployConfigData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes DeployConfigDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.DeployConfigDataAttributes{}
	}

	err, attributesAdded := retrieveModelDeployConfigDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveDeployConfigDataIDFlags(depth int, m *models.DeployConfigData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for DeployConfigDataAttributes

// register flags to command
func registerModelDeployConfigDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeployConfigDataAttributesHostname(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeployConfigDataAttributesOperatingSystem(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeployConfigDataAttributesRaid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeployConfigDataAttributesSSHKeys(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeployConfigDataAttributesHostname(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerDeployConfigDataAttributesOperatingSystem(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	operatingSystemDescription := ``

	var operatingSystemFlagName string
	if cmdPrefix == "" {
		operatingSystemFlagName = "operating_system"
	} else {
		operatingSystemFlagName = fmt.Sprintf("%v.operating_system", cmdPrefix)
	}

	var operatingSystemFlagDefault string

	_ = cmd.PersistentFlags().String(operatingSystemFlagName, operatingSystemFlagDefault, operatingSystemDescription)

	return nil
}

func registerDeployConfigDataAttributesRaid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	raidDescription := ``

	var raidFlagName string
	if cmdPrefix == "" {
		raidFlagName = "raid"
	} else {
		raidFlagName = fmt.Sprintf("%v.raid", cmdPrefix)
	}

	var raidFlagDefault string

	_ = cmd.PersistentFlags().String(raidFlagName, raidFlagDefault, raidDescription)

	return nil
}

func registerDeployConfigDataAttributesSSHKeys(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ssh_keys []int64 array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeployConfigDataAttributesFlags(depth int, m *models.DeployConfigDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, hostnameAdded := retrieveDeployConfigDataAttributesHostnameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostnameAdded

	err, operatingSystemAdded := retrieveDeployConfigDataAttributesOperatingSystemFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || operatingSystemAdded

	err, raidAdded := retrieveDeployConfigDataAttributesRaidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || raidAdded

	err, sshKeysAdded := retrieveDeployConfigDataAttributesSSHKeysFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sshKeysAdded

	return nil, retAdded
}

func retrieveDeployConfigDataAttributesHostnameFlags(depth int, m *models.DeployConfigDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeployConfigDataAttributesOperatingSystemFlags(depth int, m *models.DeployConfigDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	operatingSystemFlagName := fmt.Sprintf("%v.operating_system", cmdPrefix)
	if cmd.Flags().Changed(operatingSystemFlagName) {

		var operatingSystemFlagName string
		if cmdPrefix == "" {
			operatingSystemFlagName = "operating_system"
		} else {
			operatingSystemFlagName = fmt.Sprintf("%v.operating_system", cmdPrefix)
		}

		operatingSystemFlagValue, err := cmd.Flags().GetString(operatingSystemFlagName)
		if err != nil {
			return err, false
		}
		m.OperatingSystem = operatingSystemFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeployConfigDataAttributesRaidFlags(depth int, m *models.DeployConfigDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	raidFlagName := fmt.Sprintf("%v.raid", cmdPrefix)
	if cmd.Flags().Changed(raidFlagName) {

		var raidFlagName string
		if cmdPrefix == "" {
			raidFlagName = "raid"
		} else {
			raidFlagName = fmt.Sprintf("%v.raid", cmdPrefix)
		}

		raidFlagValue, err := cmd.Flags().GetString(raidFlagName)
		if err != nil {
			return err, false
		}
		m.Raid = raidFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeployConfigDataAttributesSSHKeysFlags(depth int, m *models.DeployConfigDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sshKeysFlagName := fmt.Sprintf("%v.ssh_keys", cmdPrefix)
	if cmd.Flags().Changed(sshKeysFlagName) {
		// warning: ssh_keys array type []int64 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
