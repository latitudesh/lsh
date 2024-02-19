package cli

import (
	"github.com/latitudesh/lsh/client/virtual_network_assignments"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
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
	Flags cmdflag.Flags
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

func (o *CreateVirtualNetworkAssignmentOperation) PromptAttributes(attributes interface{}) {
	p := prompt.New(
		prompt.NewInputText("server_id", "Server ID"),
		prompt.NewInputText("virtual_network_id", "Virtual Network ID"),
	)

	p.Run(attributes)
}

func (o *CreateVirtualNetworkAssignmentOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "server_id",
			CustomFlagName:   "server",
			Description:      "The Server ID (Required).",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
		{
			Name:             "virtual_network_id",
			CustomFlagName:   "virtual_network",
			Description:      "The Virtual Network ID (Required).",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
	}

	o.Flags.Register(schema)
}

func (o *CreateVirtualNetworkAssignmentOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *CreateVirtualNetworkAssignmentOperation) PromptQueryParams(params interface{}) {
}

func (o *CreateVirtualNetworkAssignmentOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := virtual_network_assignments.NewAssignServerVirtualNetworkParams()
	operation.AssignBodyAttributes(o, params.Body.Data.Attributes)

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
