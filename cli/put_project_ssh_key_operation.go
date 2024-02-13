package cli

import (
	"github.com/latitudesh/lsh/client/ssh_keys"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationSSHKeysPutProjectSSHKeyCmd() (*cobra.Command, error) {
	operation := UpdateSSHKeyOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type UpdateSSHKeyOperation struct {
	Flags cmdflag.Flags
}

func (o *UpdateSSHKeyOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "update",
		Short: `Allow you update SSH Key in a project. These keys can be used to access servers after deploy and reinstall actions.`,
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *UpdateSSHKeyOperation) PromptAttributes(attributes interface{}) {
	p := prompt.New(
		prompt.NewInputText("name", "Name of the API Key"),
	)

	p.Run(attributes)
}

func (o *UpdateSSHKeyOperation) PromptPathParams(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "ID from the SSH Key you want to update"),
		prompt.NewInputText("project", "Project ID or Slug"),
	)

	p.Run(params)
}

func (o *UpdateSSHKeyOperation) PromptQueryParams(params interface{}) {
}

func (o *UpdateSSHKeyOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "id",
			Description:      "ID from the SSH Key you want to update",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.PathParam,
		},
		{
			Name:             "project",
			Description:      "Project ID or Slug",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.PathParam,
		},
		{
			Name:             "name",
			Description:      "Name of the SSH Key",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
	}

	o.Flags.Register(schema)
}

func (o *UpdateSSHKeyOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *UpdateSSHKeyOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := ssh_keys.NewPutProjectSSHKeyParams()
	operation.AssignPathParams(o, params)
	operation.AssignBodyAttributes(o, params.Body.Data.Attributes)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.SSHKeys.PutProjectSSHKey(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
