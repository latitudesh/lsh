// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/models"
	"github.com/spf13/cobra"
)

// Schema cli for Projects

// register flags to command
func registerModelProjectsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerProjectsData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerProjectsData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*Project array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelProjectsFlags(depth int, m *models.Projects, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveProjectsDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveProjectsDataFlags(depth int, m *models.Projects, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type []*Project is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
