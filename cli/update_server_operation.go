package cli

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/api/resource"
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

func (o *UpdateServerOperation) PromptAttributes(attributes interface{}) {
	server := resource.NewServerResource()

	p := prompt.New(
		prompt.NewInputText("hostname", "Hostname"),
		prompt.NewInputSelect("billing", "Billing", server.SupportedBillingTypes),
	)

	p.Run(attributes)
}

func (o *UpdateServerOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "id",
			Description:      "The Server Id (Required).",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.PathParam,
		},
		{
			Name:             "hostname",
			Description:      "",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
		{
			Name:             "billing",
			Description:      "",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
	}

	o.Flags.Register(schema)
}

func (o *UpdateServerOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *UpdateServerOperation) PromptPathParams(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "ID from the Server you want to update"),
	)

	p.Run(params)
}

func (o *UpdateServerOperation) PromptQueryParams(params interface{}) {
}

func (o *UpdateServerOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := servers.NewUpdateServerParams()
	operation.AssignPathParams(o, params)
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
		utils.Render(response.GetData())
	}
	return nil
}
