package cli

import (
	"github.com/latitudesh/lsh/client/virtual_networks"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type DeleteVirtualNetworkOperation struct {
	Flags cmdflag.Flags
}

func makeOperationVirtualNetworksDestroyVirtualNetworkCmd() (*cobra.Command, error) {
	operation := DeleteVirtualNetworkOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *DeleteVirtualNetworkOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroy",
		Short: "Delete virtual network.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *DeleteVirtualNetworkOperation) PromptPathParams(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "ID"),
	)

	p.Run(params)
}

func (o *DeleteVirtualNetworkOperation) PromptQueryParams(params interface{}) {
}

func (o *DeleteVirtualNetworkOperation) PromptAttributes(attributes interface{}) {
}

func (o *DeleteVirtualNetworkOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *DeleteVirtualNetworkOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "id",
			Description:      "The virtual network ID",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.PathParam,
		},
	}

	o.Flags.Register(schema)
}

func (o *DeleteVirtualNetworkOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := virtual_networks.NewDestroyVirtualNetworkParams()
	operation.AssignPathParams(o, params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.VirtualNetworks.DestroyVirtualNetwork(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}
