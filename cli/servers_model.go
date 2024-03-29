// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/models"
	"github.com/spf13/cobra"
)

// Schema cli for Servers

// register flags to command
func registerModelServersFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServersData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServersMeta(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServersData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*ServerData array type is not supported by go-swagger cli yet

	return nil
}

func registerServersMeta(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: meta interface{} map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServersFlags(depth int, m *models.Servers, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveServersDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	err, metaAdded := retrieveServersMetaFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metaAdded

	return nil, retAdded
}

func retrieveServersDataFlags(depth int, m *models.Servers, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// warning: data array type []*ServerData is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveServersMetaFlags(depth int, m *models.Servers, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
