package cli

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationServersUpdateServerCmd() (*cobra.Command, error) {
	operation := UpdateServerOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type UpdateServerOperation struct {
	Flags cmdflag.Flags
}

func (o *UpdateServerOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update server information.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
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
		prompt.NewInputSelect("billing", "Billing", []string{"SKIP", "hourly", "monthly", "yearly"}),
	)

	p.Run(attributes)
}

func (o *UpdateServerOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
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

	o.Flags.Register(schema)
}

func (o *UpdateServerOperation) GetFlags() cmdflag.Flags {
	return o.Flags
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
	operation.AssignResourceID(o, params)
	operation.AssignBodyAttributes(o, params.Body.Data.Attributes)

	params.Body.Data.ID = params.ID

	if swag.IsZero(*params.Body.Data.Attributes) {
		fmt.Println("Skipped action: no params provided")
		return nil
	}

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
