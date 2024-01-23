// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/models"

	"github.com/spf13/cobra"
)

// Schema cli for Membership

// register flags to command
func registerModelMembershipFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMembershipData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMembershipData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelMembershipDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMembershipFlags(depth int, m *models.Membership, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveMembershipDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveMembershipDataFlags(depth int, m *models.Membership, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data MembershipData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.MembershipData{}
	}

	err, dataAdded := retrieveModelMembershipDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for MembershipData

// register flags to command
func registerModelMembershipDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMembershipDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMembershipDataID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMembershipDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName string
	if cmdPrefix == "" {
		attributesFlagName = "attributes"
	} else {
		attributesFlagName = fmt.Sprintf("%v.attributes", cmdPrefix)
	}

	if err := registerModelMembershipDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerMembershipDataID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelMembershipDataFlags(depth int, m *models.MembershipData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveMembershipDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveMembershipDataIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveMembershipDataAttributesFlags(depth int, m *models.MembershipData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes MembershipDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &models.MembershipDataAttributes{}
	}

	err, attributesAdded := retrieveModelMembershipDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

func retrieveMembershipDataIDFlags(depth int, m *models.MembershipData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for MembershipDataAttributes

// register flags to command
func registerModelMembershipDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMembershipDataAttributesEmail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMembershipDataAttributesFirstName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMembershipDataAttributesLastName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMembershipDataAttributesRole(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMembershipDataAttributesEmail(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	emailDescription := ``

	var emailFlagName string
	if cmdPrefix == "" {
		emailFlagName = "email"
	} else {
		emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
	}

	var emailFlagDefault string

	_ = cmd.PersistentFlags().String(emailFlagName, emailFlagDefault, emailDescription)

	return nil
}

func registerMembershipDataAttributesFirstName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	firstNameDescription := ``

	var firstNameFlagName string
	if cmdPrefix == "" {
		firstNameFlagName = "first_name"
	} else {
		firstNameFlagName = fmt.Sprintf("%v.first_name", cmdPrefix)
	}

	var firstNameFlagDefault string

	_ = cmd.PersistentFlags().String(firstNameFlagName, firstNameFlagDefault, firstNameDescription)

	return nil
}

func registerMembershipDataAttributesLastName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lastNameDescription := ``

	var lastNameFlagName string
	if cmdPrefix == "" {
		lastNameFlagName = "last_name"
	} else {
		lastNameFlagName = fmt.Sprintf("%v.last_name", cmdPrefix)
	}

	var lastNameFlagDefault string

	_ = cmd.PersistentFlags().String(lastNameFlagName, lastNameFlagDefault, lastNameDescription)

	return nil
}

func registerMembershipDataAttributesRole(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	roleDescription := `Enum: ["owner","administrator","collaborator","billing"]. `

	var roleFlagName string
	if cmdPrefix == "" {
		roleFlagName = "role"
	} else {
		roleFlagName = fmt.Sprintf("%v.role", cmdPrefix)
	}

	var roleFlagDefault string

	_ = cmd.PersistentFlags().String(roleFlagName, roleFlagDefault, roleDescription)

	if err := cmd.RegisterFlagCompletionFunc(roleFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["owner","administrator","collaborator","billing"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMembershipDataAttributesFlags(depth int, m *models.MembershipDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, emailAdded := retrieveMembershipDataAttributesEmailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || emailAdded

	err, firstNameAdded := retrieveMembershipDataAttributesFirstNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firstNameAdded

	err, lastNameAdded := retrieveMembershipDataAttributesLastNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lastNameAdded

	err, roleAdded := retrieveMembershipDataAttributesRoleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || roleAdded

	return nil, retAdded
}

func retrieveMembershipDataAttributesEmailFlags(depth int, m *models.MembershipDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	emailFlagName := fmt.Sprintf("%v.email", cmdPrefix)
	if cmd.Flags().Changed(emailFlagName) {

		var emailFlagName string
		if cmdPrefix == "" {
			emailFlagName = "email"
		} else {
			emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
		}

		emailFlagValue, err := cmd.Flags().GetString(emailFlagName)
		if err != nil {
			return err, false
		}
		m.Email = emailFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMembershipDataAttributesFirstNameFlags(depth int, m *models.MembershipDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	firstNameFlagName := fmt.Sprintf("%v.first_name", cmdPrefix)
	if cmd.Flags().Changed(firstNameFlagName) {

		var firstNameFlagName string
		if cmdPrefix == "" {
			firstNameFlagName = "first_name"
		} else {
			firstNameFlagName = fmt.Sprintf("%v.first_name", cmdPrefix)
		}

		firstNameFlagValue, err := cmd.Flags().GetString(firstNameFlagName)
		if err != nil {
			return err, false
		}
		m.FirstName = firstNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMembershipDataAttributesLastNameFlags(depth int, m *models.MembershipDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lastNameFlagName := fmt.Sprintf("%v.last_name", cmdPrefix)
	if cmd.Flags().Changed(lastNameFlagName) {

		var lastNameFlagName string
		if cmdPrefix == "" {
			lastNameFlagName = "last_name"
		} else {
			lastNameFlagName = fmt.Sprintf("%v.last_name", cmdPrefix)
		}

		lastNameFlagValue, err := cmd.Flags().GetString(lastNameFlagName)
		if err != nil {
			return err, false
		}
		m.LastName = lastNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMembershipDataAttributesRoleFlags(depth int, m *models.MembershipDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	roleFlagName := fmt.Sprintf("%v.role", cmdPrefix)
	if cmd.Flags().Changed(roleFlagName) {

		var roleFlagName string
		if cmdPrefix == "" {
			roleFlagName = "role"
		} else {
			roleFlagName = fmt.Sprintf("%v.role", cmdPrefix)
		}

		roleFlagValue, err := cmd.Flags().GetString(roleFlagName)
		if err != nil {
			return err, false
		}
		m.Role = roleFlagValue

		retAdded = true
	}

	return nil, retAdded
}
