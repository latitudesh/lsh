package cli

import (
	"github.com/latitudesh/lsh/client/virtual_networks"
	"github.com/latitudesh/lsh/internal/api/resource"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
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
	Flags cmdflag.Flags
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

func (o *CreateVirtualNetworkOperation) PromptAttributes(attributes interface{}) {
	virtualNetwork := resource.NewVirtualNetworkResource()

	p := prompt.New(
		prompt.NewInputText("description", "Description"),
		prompt.NewInputText("project", "Project ID or slug"),
		prompt.NewInputSelect("site", "Site ID or Slug", virtualNetwork.SupportedSites),
	)

	p.Run(attributes)
}

func (o *CreateVirtualNetworkOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "description",
			Description:      "Description",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
		{
			Name:             "project",
			Description:      "Project ID or slug",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
		{
			Name:             "site",
			Description:      "Site ID or slug",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
	}

	o.Flags.Register(schema)
}

func (o *CreateVirtualNetworkOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *CreateVirtualNetworkOperation) PromptQueryParams(params interface{}) {
}

func (o *CreateVirtualNetworkOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := virtual_networks.NewCreateVirtualNetworkParams()
	operation.AssignBodyAttributes(o, params.Body.Data.Attributes)

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
