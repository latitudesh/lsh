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
	FlagsSchema      cmdflag.FlagsSchema
	Name             string
	ShortDescription string
	FlagSet          *pflag.FlagSet
}

func makeOperationServersUpdateServerCmd() (*cobra.Command, error) {
	o := DestroyServerOperation{
		Name:             "update",
		ShortDescription: "Update server information.",
	}

	cmd, err := o.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *UpdateServerOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   o.Name,
		Short: o.ShortDescription,
		RunE:  o.run,
	}

	o.RegisterFlags(cmd)

	return cmd, nil
}

func (o *UpdateServerOperation) AssignID(data interface{}) error {
	id, err := o.FlagSet.GetString("id")
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

func (o *UpdateServerOperation) RegisterFlags(cmd *cobra.Command) {
	o.FlagsSchema = cmdflag.FlagsSchema{
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

	o.FlagsSchema.Register(o.FlagSet)
}

func (o *UpdateServerOperation) GetFlagsDefinition() cmdflag.FlagsSchema {
	return o.FlagsSchema
}

func (o *UpdateServerOperation) GetFlagSet() *pflag.FlagSet {
	return o.FlagSet
}

func (o *UpdateServerOperation) ResourceIDFlag() cmdflag.FlagSchema {
	return o.FlagsSchema[0]
}

func (o *UpdateServerOperation) PromptID(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "ID from the Server you want to update"),
	)

	p.Run(params)
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
