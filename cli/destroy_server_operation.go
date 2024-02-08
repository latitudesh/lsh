package cli

import (
	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type DestroyServerOperation struct {
	Name             string
	ShortDescription string
	FlagsSchema      cmdflag.FlagsSchema
	FlagSet          *pflag.FlagSet
}

func (o *DestroyServerOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   o.Name,
		Short: o.ShortDescription,
		RunE:  o.run,
	}

	o.RegisterFlags(cmd)

	return cmd, nil
}

func (o *DestroyServerOperation) ResourceIDFlag() cmdflag.FlagSchema {
	return o.FlagsSchema[0]
}

func (o *DestroyServerOperation) PromptID(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "ID from the Server you want to destroy"),
	)

	p.Run(params)
}

func (o *DestroyServerOperation) PromptAttributes(attributes interface{}) {
}

func (o *DestroyServerOperation) GetFlagsDefinition() cmdflag.FlagsSchema {
	return o.FlagsSchema
}

func (o *DestroyServerOperation) GetFlagSet() *pflag.FlagSet {
	return o.FlagSet
}

func makeOperationServersDestroyServerCmd() (*cobra.Command, error) {
	o := DestroyServerOperation{
		Name:             "destroy",
		ShortDescription: "Remove a server.",
	}

	cmd, err := o.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *DestroyServerOperation) RegisterFlags(cmd *cobra.Command) {
	o.FlagSet = cmd.Flags()
	o.FlagsSchema = cmdflag.FlagsSchema{
		{
			Name:         "id",
			Description:  "Required. The server ID",
			DefaultValue: "",
			Type:         "string",
		},
	}

	o.FlagsSchema.Register(o.FlagSet)
}

func (o *DestroyServerOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := servers.NewDestroyServerParams()
	operation.AssignResourceID(o, params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Servers.DestroyServer(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}
