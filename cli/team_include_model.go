// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/models"
	"github.com/spf13/cobra"
)

// Schema cli for TeamInclude

// register flags to command
func registerModelTeamIncludeFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTeamIncludeAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamIncludeCurrency(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamIncludeDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamIncludeID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamIncludeName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamIncludeSlug(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTeamIncludeStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTeamIncludeAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerTeamIncludeCurrency(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: currency interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerTeamIncludeDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerTeamIncludeID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerTeamIncludeName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

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

func registerTeamIncludeSlug(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	slugDescription := ``

	var slugFlagName string
	if cmdPrefix == "" {
		slugFlagName = "slug"
	} else {
		slugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
	}

	var slugFlagDefault string

	_ = cmd.PersistentFlags().String(slugFlagName, slugFlagDefault, slugDescription)

	return nil
}

func registerTeamIncludeStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTeamIncludeFlags(depth int, m *models.TeamInclude, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressAdded := retrieveTeamIncludeAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, currencyAdded := retrieveTeamIncludeCurrencyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || currencyAdded

	err, descriptionAdded := retrieveTeamIncludeDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, idAdded := retrieveTeamIncludeIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, nameAdded := retrieveTeamIncludeNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, slugAdded := retrieveTeamIncludeSlugFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || slugAdded

	err, statusAdded := retrieveTeamIncludeStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	return nil, retAdded
}

func retrieveTeamIncludeAddressFlags(depth int, m *models.TeamInclude, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTeamIncludeCurrencyFlags(depth int, m *models.TeamInclude, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	currencyFlagName := fmt.Sprintf("%v.currency", cmdPrefix)
	if cmd.Flags().Changed(currencyFlagName) {
		// warning: currency map type interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveTeamIncludeDescriptionFlags(depth int, m *models.TeamInclude, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTeamIncludeIDFlags(depth int, m *models.TeamInclude, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTeamIncludeNameFlags(depth int, m *models.TeamInclude, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTeamIncludeSlugFlags(depth int, m *models.TeamInclude, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	slugFlagName := fmt.Sprintf("%v.slug", cmdPrefix)
	if cmd.Flags().Changed(slugFlagName) {

		var slugFlagName string
		if cmdPrefix == "" {
			slugFlagName = "slug"
		} else {
			slugFlagName = fmt.Sprintf("%v.slug", cmdPrefix)
		}

		slugFlagValue, err := cmd.Flags().GetString(slugFlagName)
		if err != nil {
			return err, false
		}
		m.Slug = slugFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTeamIncludeStatusFlags(depth int, m *models.TeamInclude, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
