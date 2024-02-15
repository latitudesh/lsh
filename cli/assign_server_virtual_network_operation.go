package cli

import (
	"github.com/latitudesh/lsh/client/virtual_network_assignments"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkCmd() (*cobra.Command, error) {
	operation := CreateVirtualNetworkAssignmentOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type CreateVirtualNetworkAssignmentOperation struct {
	BodyAttributesFlags cmdflag.Flags
}

func (o *CreateVirtualNetworkAssignmentOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Assign a server to a virtual network.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *CreateVirtualNetworkAssignmentOperation) registerFlags(cmd *cobra.Command) {
	o.BodyAttributesFlags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "server",
			Label:       "Server ID",
			Description: "The Server ID (Required).",
		},
		&cmdflag.String{
			Name:        "virtual_network",
			Label:       "Virtual Network ID",
			Description: "The Virtual Network ID (Required).",
		},
	}

	o.BodyAttributesFlags.Register(schema)
}

func (o *CreateVirtualNetworkAssignmentOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := virtual_network_assignments.NewAssignServerVirtualNetworkParams()
	o.BodyAttributesFlags.AssignValues(params.Body.Data.Attributes)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.VirtualNetworkAssignments.AssignServerVirtualNetwork(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
