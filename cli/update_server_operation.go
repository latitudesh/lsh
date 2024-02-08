package cli

import (
	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationServersUpdateServerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "update",
		Short: `Update server information.`,
	}

	registerUpdateServerOperation(cmd)

	return cmd, nil
}

type UpdateServerOperation struct {
	Command *cobra.Command
	Flags   cmdflag.Flags
	Params  *servers.UpdateServerParams
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

func (o *UpdateServerOperation) AssignBodyAttributes(attributes interface{}) error {
	for _, v := range o.Flags {
		if v.Name == "id" {
			continue
		}

		switch v.Type {
		case "string":
			value, err := o.Command.Flags().GetString(v.Name)

			if err != nil {
				return err
			}

			utils.AssignValue(attributes, v.Name, value)
		}
	}

	o.promptAttributes(attributes)

	return nil
}

func (o *UpdateServerOperation) promptID(bodyData interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "ID from the Server you want to update"),
	)

	p.Run(bodyData)
}

func (o *UpdateServerOperation) promptAttributes(attributes interface{}) {
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

	for _, v := range o.Flags {
		switch v.Type {
		case "string":
			o.Command.PersistentFlags().String(v.Name, v.DefaultValue, v.Label)
		}
	}
}

func (o *UpdateServerOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := servers.NewUpdateServerParams()
	o.AssignID(params.Body.Data)
	o.AssignBodyAttributes(params.Body.Data.Attributes)
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
