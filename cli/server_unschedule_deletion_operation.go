package cli

import (
	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationServersServerUnscheduleDeletionCmd() (*cobra.Command, error) {
	operation := UnscheduleServerDeletionOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type UnscheduleServerDeletionOperation struct {
	Flags cmdflag.Flags
}

func (o *UnscheduleServerDeletionOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "unschedule-deletion",
		Short: "Unschedules the server removal at the end of the billing cycle.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *UnscheduleServerDeletionOperation) PromptAttributes(attributes interface{}) {
}

func (o *UnscheduleServerDeletionOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "id",
			Description:      "The Server Id (Required).",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.PathParam,
		},
	}

	o.Flags.Register(schema)
}

func (o *UnscheduleServerDeletionOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *UnscheduleServerDeletionOperation) PromptQueryParams(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "Server ID"),
	)

	p.Run(params)
}

func (o *UnscheduleServerDeletionOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := servers.NewServerUnscheduleDeletionParams()

	operation.AssignPathParams(o, params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Servers.ServerUnscheduleDeletion(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}
