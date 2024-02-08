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

type UpdateServerOperation struct {
	Command *cobra.Command
	Flags   cmdflag.Flags
	Params  *servers.UpdateServerParams
}

func makeOperationServersUpdateServerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "update",
		Short: `Update server information.`,
	}

	registerUpdateServerOperation(cmd)

	return cmd, nil
}

func registerUpdateServerOperation(cmd *cobra.Command) {
	op := UpdateServerOperation{Command: cmd}
	op.Command.RunE = op.run
	op.RegisterFlags()
}

func (o *UpdateServerOperation) AssignID(data interface{}) error {
	id, err := o.Command.Flags().GetString("id")
	if err != nil {
		return err
	}

	utils.AssignValue(data, "id", id)

	o.promptID(data)

	return nil
}

func (o *UpdateServerOperation) promptID(bodyData interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "ID from the Server you want to update"),
	)

	p.Run(bodyData)
}

func (o *UpdateServerOperation) PromptAttributes(attributes interface{}) {
	p := prompt.New(
		prompt.NewInputText("hostname", "Hostname"),
		prompt.NewInputSelect("billing", "Billing", []string{"hourly", "monthly", "yearly"}),
	)

	p.Run(attributes)
}

func (o *UpdateServerOperation) RegisterFlags() {
	o.Flags = cmdflag.Flags{
		{
			Name:         "id",
			Label:        "The Server Id (Required).",
			DefaultValue: "",
			Type:         "string",
		},
		{
			Name:         "hostname",
			Label:        "",
			DefaultValue: "",
			Type:         "string",
		},
		{
			Name:         "billing",
			Label:        "",
			DefaultValue: "",
			Type:         "string",
		},
	}

	o.Flags.Register(o.Command.Flags())
}

func (o *UpdateServerOperation) GetFlagsDefinition() cmdflag.Flags {
	return o.Flags
}

func (o *UpdateServerOperation) GetFlagSet() *pflag.FlagSet {
	return o.Command.Flags()
}

func (o *UpdateServerOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := servers.NewUpdateServerParams()
	o.AssignID(params.Body.Data)
	operation.AssignBodyAttributes(o, params.Body.Data.Attributes)
	params.SetID(params.Body.Data.ID)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Servers.UpdateServer(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}
