package cli

import (
	"github.com/latitudesh/lsh/client/ssh_keys"
	"github.com/latitudesh/lsh/internal/cmdflag"
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
	PathParamFlags cmdflag.Flags
}

func (o *DeleteSSHKeyOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:    "destroy",
		Short:  `Allow you remove SSH Keys in a project. Remove a SSH Key from the project won't revoke the SSH Keys access for previously deploy and reinstall actions.`,
		RunE:   o.run,
		PreRun: o.preRun,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *DeleteSSHKeyOperation) registerFlags(cmd *cobra.Command) {
	o.PathParamFlags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "id",
			Label:       "ID from the SSH Key you want to update",
			Description: "ID from the SSH Key you want to update",
			Required:    true,
		},
		&cmdflag.String{
			Name:        "project",
			Label:       "Project ID or Slug",
			Description: "Project ID or Slug",
			Required:    true,
		},
	}

	o.PathParamFlags.Register(schema)
}

func (o *DeleteSSHKeyOperation) preRun(cmd *cobra.Command, args []string) {
	o.PathParamFlags.PreRun(cmd, args)
}

func (o *DeleteSSHKeyOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := ssh_keys.NewDeleteProjectSSHKeyParams()
	o.PathParamFlags.AssignValues(params)

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
