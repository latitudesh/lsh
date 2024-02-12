package cli

import (
	"github.com/latitudesh/lsh/client/ssh_keys"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationSSHKeysDeleteProjectSSHKeyCmd() (*cobra.Command, error) {
	operation := DeleteSSHKeyOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type DeleteSSHKeyOperation struct {
	Flags cmdflag.Flags
}

func (o *DeleteSSHKeyOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroy",
		Short: `Allow you remove SSH Keys in a project. Remove a SSH Key from the project won't revoke the SSH Keys access for previously deploy and reinstall actions.`,
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *DeleteSSHKeyOperation) PromptAttributes(attributes interface{}) {
}

func (o *DeleteSSHKeyOperation) PromptQueryParams(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "ID from the SSH Key you want to update"),
		prompt.NewInputText("project", "Project ID or Slug"),
	)

	p.Run(params)
}

func (o *DeleteSSHKeyOperation) registerFlags(cmd *cobra.Command) {
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
	}

	o.Flags.Register(schema)
}

func (o *DeleteSSHKeyOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *DeleteSSHKeyOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := ssh_keys.NewDeleteProjectSSHKeyParams()
	operation.AssignPathParams(o, params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.SSHKeys.DeleteProjectSSHKey(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}
