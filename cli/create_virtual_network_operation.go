package cli

import (
	"github.com/latitudesh/lsh/client/virtual_networks"
	"github.com/latitudesh/lsh/internal/api/resource"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationVirtualNetworksCreateVirtualNetworkCmd() (*cobra.Command, error) {
	operation := CreateVirtualNetworkOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type CreateVirtualNetworkOperation struct {
	BodyAttributesFlags cmdflag.Flags
}

func (o *CreateVirtualNetworkOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new Virtual Network.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *CreateVirtualNetworkOperation) registerFlags(cmd *cobra.Command) {
	virtualNetwork := resource.NewVirtualNetworkResource()

	o.BodyAttributesFlags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "description",
			Label:       "Description",
			Description: "Description",
			Required:    true,
		},
		&cmdflag.String{
			Name:        "project",
			Label:       "Project ID or slug",
			Description: "Project ID or slug",
			Required:    true,
		},
		&cmdflag.String{
			Name:        "site",
			Label:       "Site ID or slug",
			Description: "Site ID or slug",
			Options:     virtualNetwork.SupportedSites,
			Required:    false,
		},
	}

	o.BodyAttributesFlags.Register(schema)
}

func (o *CreateVirtualNetworkOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := virtual_networks.NewCreateVirtualNetworkParams()
	o.BodyAttributesFlags.AssignValues(params.Body.Data.Attributes)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.VirtualNetworks.CreateVirtualNetwork(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
