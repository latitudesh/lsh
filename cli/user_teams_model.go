// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/models"
	"github.com/spf13/cobra"
)

// Schema cli for UserTeams

// register flags to command
func registerModelUserTeamsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUserTeamsData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserTeamsMeta(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserTeamsData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*UserTeam array type is not supported by go-swagger cli yet

	return nil
}

func registerUserTeamsMeta(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: meta interface{} map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUserTeamsFlags(depth int, m *models.UserTeams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveUserTeamsDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	err, metaAdded := retrieveUserTeamsMetaFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metaAdded

	return nil, retAdded
}

func retrieveUserTeamsDataFlags(depth int, m *models.UserTeams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type []*UserTeam is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveUserTeamsMetaFlags(depth int, m *models.UserTeams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	metaFlagName := fmt.Sprintf("%v.meta", cmdPrefix)
	if cmd.Flags().Changed(metaFlagName) {
		// warning: meta map type interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
