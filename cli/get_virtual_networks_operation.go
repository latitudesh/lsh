package cli

import (
	"github.com/latitudesh/lsh/client/virtual_networks"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type ListVirtualNetworksOperation struct {
	QueryParamsFlags cmdflag.Flags
}

func makeOperationVirtualNetworksGetVirtualNetworksCmd() (*cobra.Command, error) {
	operation := ListVirtualNetworksOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *ListVirtualNetworksOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists virtual networks assigned to a project.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *ListVirtualNetworksOperation) registerFlags(cmd *cobra.Command) {
	o.QueryParamsFlags = cmdflag.Flags{
		FlagSet:           cmd.Flags(),
		PromptDescription: "Filter Virtual Networks: (press ENTER to skip a filter)",
	}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "location",
			Label:       "Location",
			Description: "The location slug to filter by",
		},
		&cmdflag.String{
			Name:        "project",
			Label:       "Project",
			Description: "The project id or slug to filter by",
		},
	}

	o.QueryParamsFlags.Register(schema)
}

func (o *ListVirtualNetworksOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := virtual_networks.NewGetVirtualNetworksParams()
	o.QueryParamsFlags.AssignValues(params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.VirtualNetworks.GetVirtualNetworks(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
