package cli

import (
	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationServersDestroyServerCmd() (*cobra.Command, error) {
	operation := DestroyServerOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type DestroyServerOperation struct {
	Flags cmdflag.Flags
}

func (o *DestroyServerOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroy",
		Short: "Remove a server.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *DestroyServerOperation) PromptID(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "ID from the Server you want to destroy"),
	)

	p.Run(params)
}

func (o *DestroyServerOperation) PromptAttributes(attributes interface{}) {
}

func (o *DestroyServerOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *DestroyServerOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := cmdflag.FlagsSchema{
		{
			Name:         "id",
			Description:  "Required. The server ID",
			DefaultValue: "",
			Type:         "string",
		},
	}

	o.Flags.Register(schema)
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
