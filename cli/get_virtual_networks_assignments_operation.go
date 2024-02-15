package cli

import (
	"github.com/latitudesh/lsh/client/virtual_network_assignments"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type ListVirtualNetworkAssignmentsOperation struct {
	QueryParamsFlags cmdflag.Flags
}

func makeOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsCmd() (*cobra.Command, error) {
	operation := ListVirtualNetworkAssignmentsOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *ListVirtualNetworkAssignmentsOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Returns a list of all servers assigned to virtual networks.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *ListVirtualNetworkAssignmentsOperation) registerFlags(cmd *cobra.Command) {
	o.QueryParamsFlags = cmdflag.Flags{
		FlagSet:           cmd.Flags(),
		PromptDescription: "Filter virtual network assignments: (press ENTER to skip a filter)",
	}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "server",
			Label:       "Server ID",
			Description: "The server ID to filter by",
		},
		&cmdflag.String{
			Name:        "vid",
			Label:       "VLAN ID",
			Description: "The vlan ID to filter by",
		},
		&cmdflag.String{
			Name:        "virtual_network_id",
			Label:       "Virtual Network ID",
			Description: "The virtual network ID to filter by",
		},
	}

	o.QueryParamsFlags.Register(schema)
}

func (o *ListVirtualNetworkAssignmentsOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := virtual_network_assignments.NewGetVirtualNetworksAssignmentsParams()
	o.QueryParamsFlags.AssignValues(params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.VirtualNetworkAssignments.GetVirtualNetworksAssignments(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
