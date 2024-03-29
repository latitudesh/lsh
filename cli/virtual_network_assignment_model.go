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

// Schema cli for VirtualNetworkAssignment

// register flags to command
func registerModelVirtualNetworkAssignmentFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVirtualNetworkAssignmentAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAssignmentID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVirtualNetworkAssignmentAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelVirtualNetworkAssignmentAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerVirtualNetworkAssignmentID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelVirtualNetworkAssignmentFlags(depth int, m *models.VirtualNetworkAssignment, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveVirtualNetworkAssignmentAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveVirtualNetworkAssignmentIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveVirtualNetworkAssignmentAttributesFlags(depth int, m *models.VirtualNetworkAssignment, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes VirtualNetworkAssignmentAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.VirtualNetworkAssignmentAttributes{}
	}

	err, attributesAdded := retrieveModelVirtualNetworkAssignmentAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAssignmentIDFlags(depth int, m *models.VirtualNetworkAssignment, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for VirtualNetworkAssignmentAttributes

// register flags to command
func registerModelVirtualNetworkAssignmentAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVirtualNetworkAssignmentAttributesDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAssignmentAttributesServerID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAssignmentAttributesStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAssignmentAttributesVid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVirtualNetworkAssignmentAttributesVirtualNetworkID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVirtualNetworkAssignmentAttributesDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerVirtualNetworkAssignmentAttributesServerID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	serverIdDescription := ``

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

func registerVirtualNetworkAssignmentAttributesStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := ``

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	return nil
}

func registerVirtualNetworkAssignmentAttributesVid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vidDescription := ``

	var vidFlagName string
	if cmdPrefix == "" {
		vidFlagName = "vid"
	} else {
		vidFlagName = fmt.Sprintf("%v.vid", cmdPrefix)
	}

	var vidFlagDefault int64

	_ = cmd.PersistentFlags().Int64(vidFlagName, vidFlagDefault, vidDescription)

	return nil
}

func registerVirtualNetworkAssignmentAttributesVirtualNetworkID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	virtualNetworkIdDescription := ``

	var virtualNetworkIdFlagName string
	if cmdPrefix == "" {
		virtualNetworkIdFlagName = "virtual_network_id"
	} else {
		virtualNetworkIdFlagName = fmt.Sprintf("%v.virtual_network_id", cmdPrefix)
	}

	var virtualNetworkIdFlagDefault string

	_ = cmd.PersistentFlags().String(virtualNetworkIdFlagName, virtualNetworkIdFlagDefault, virtualNetworkIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVirtualNetworkAssignmentAttributesFlags(depth int, m *models.VirtualNetworkAssignmentAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveVirtualNetworkAssignmentAttributesDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, serverIdAdded := retrieveVirtualNetworkAssignmentAttributesServerIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serverIdAdded

	err, statusAdded := retrieveVirtualNetworkAssignmentAttributesStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, vidAdded := retrieveVirtualNetworkAssignmentAttributesVidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vidAdded

	err, virtualNetworkIdAdded := retrieveVirtualNetworkAssignmentAttributesVirtualNetworkIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || virtualNetworkIdAdded

	return nil, retAdded
}

func retrieveVirtualNetworkAssignmentAttributesDescriptionFlags(depth int, m *models.VirtualNetworkAssignmentAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveVirtualNetworkAssignmentAttributesServerIDFlags(depth int, m *models.VirtualNetworkAssignmentAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serverIdFlagName := fmt.Sprintf("%v.server_id", cmdPrefix)
	if cmd.Flags().Changed(serverIdFlagName) {

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

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAssignmentAttributesStatusFlags(depth int, m *models.VirtualNetworkAssignmentAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetString(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = statusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAssignmentAttributesVidFlags(depth int, m *models.VirtualNetworkAssignmentAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vidFlagName := fmt.Sprintf("%v.vid", cmdPrefix)
	if cmd.Flags().Changed(vidFlagName) {

		var vidFlagName string
		if cmdPrefix == "" {
			vidFlagName = "vid"
		} else {
			vidFlagName = fmt.Sprintf("%v.vid", cmdPrefix)
		}

		vidFlagValue, err := cmd.Flags().GetInt64(vidFlagName)
		if err != nil {
			return err, false
		}
		m.Vid = vidFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVirtualNetworkAssignmentAttributesVirtualNetworkIDFlags(depth int, m *models.VirtualNetworkAssignmentAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	virtualNetworkIdFlagName := fmt.Sprintf("%v.virtual_network_id", cmdPrefix)
	if cmd.Flags().Changed(virtualNetworkIdFlagName) {

		var virtualNetworkIdFlagName string
		if cmdPrefix == "" {
			virtualNetworkIdFlagName = "virtual_network_id"
		} else {
			virtualNetworkIdFlagName = fmt.Sprintf("%v.virtual_network_id", cmdPrefix)
		}

		virtualNetworkIdFlagValue, err := cmd.Flags().GetString(virtualNetworkIdFlagName)
		if err != nil {
			return err, false
		}
		m.VirtualNetworkID = virtualNetworkIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
