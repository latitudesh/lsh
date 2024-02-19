package cli

import (
	"github.com/latitudesh/lsh/client/ssh_keys"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationSSHKeysPostProjectSSHKeyCmd() (*cobra.Command, error) {
	operation := CreateSSHKeyOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type CreateSSHKeyOperation struct {
	Flags cmdflag.Flags
}

func (o *CreateSSHKeyOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "create",
		Short: `Allow you create SSH Keys in a project. These keys can be used to access servers after deploy and reinstall actions.`,
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *CreateSSHKeyOperation) PromptAttributes(attributes interface{}) {
	p := prompt.New(
		prompt.NewInputText("name", "Name of the API Key"),
		prompt.NewInputText("public_key", "SSH Public Key"),
	)

	p.Run(attributes)
}

func (o *CreateSSHKeyOperation) PromptQueryParams(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("project", "ID or Slug from the project you want to add the SSH Key"),
	)

	p.Run(params)
}

func (o *CreateSSHKeyOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "project",
			Description:      "ID or Slug from the project you want to add the SSH Key",
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
		{
			Name:             "public_key",
			Description:      "SSH Public Key",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
	}

	o.Flags.Register(schema)
}

func (o *CreateSSHKeyOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *CreateSSHKeyOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := ssh_keys.NewPostProjectSSHKeyParams()
	operation.AssignPathParams(o, params)
	operation.AssignBodyAttributes(o, params.Body.Data.Attributes)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.SSHKeys.PostProjectSSHKey(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
