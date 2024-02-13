package cli

import (
	"fmt"

	"github.com/latitudesh/lsh/client/virtual_network_assignments"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsCmd() (*cobra.Command, error) {
	operation := DeleteVirtualNetworkAssignmentOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type DeleteVirtualNetworkAssignmentOperation struct {
	Flags cmdflag.Flags
}

func (o *DeleteVirtualNetworkAssignmentOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroy",
		Short: "Allow you to remove a Virtual Network assignment.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *DeleteVirtualNetworkAssignmentOperation) PromptAttributes(attributes interface{}) {
}

func (o *DeleteVirtualNetworkAssignmentOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "id",
			Description:      "Virtual Network Assignment Id (Required).",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.PathParam,
		},
	}

	o.Flags.Register(schema)
}

func (o *DeleteVirtualNetworkAssignmentOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *DeleteVirtualNetworkAssignmentOperation) PromptPathParams(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "Virtual Network Assignment ID"),
	)

	p.Run(params)
}

func (o *DeleteVirtualNetworkAssignmentOperation) PromptQueryParams(params interface{}) {
}

func (o *DeleteVirtualNetworkAssignmentOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := virtual_network_assignments.NewDeleteVirtualNetworksAssignmentsParams()
	operation.AssignPathParams(o, params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.VirtualNetworkAssignments.DeleteVirtualNetworksAssignments(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}

// runOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignments uses cmd flags to call endpoint api
func runOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignments(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := virtual_network_assignments.NewDeleteVirtualNetworksAssignmentsParams()
	if err, _ := retrieveOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAssignmentIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.VirtualNetworkAssignments.DeleteVirtualNetworksAssignments(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}

// registerOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsParamFlags registers all flags needed to fill params
func registerOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAssignmentIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAssignmentIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	assignmentIdDescription := `Virtual Network Assignment Id (Required).`

	var assignmentIdFlagName string
	if cmdPrefix == "" {
		assignmentIdFlagName = "id"
	} else {
		assignmentIdFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var assignmentIdFlagDefault string

	_ = cmd.PersistentFlags().String(assignmentIdFlagName, assignmentIdFlagDefault, assignmentIdDescription)
	cmd.MarkPersistentFlagRequired(assignmentIdFlagName)

	return nil
}

func retrieveOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsAssignmentIDFlag(m *virtual_network_assignments.DeleteVirtualNetworksAssignmentsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var assignmentIdFlagName string
		if cmdPrefix == "" {
			assignmentIdFlagName = "id"
		} else {
			assignmentIdFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		assignmentIdFlagValue, err := cmd.Flags().GetString(assignmentIdFlagName)
		if err != nil {
			return err, false
		}
		m.AssignmentID = assignmentIdFlagValue

	}
	return nil, retAdded
}
