package cli

import (
	"github.com/latitudesh/lsh/client/ssh_keys"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type GetSSHKeyOperation struct {
	PathParamsFlags cmdflag.Flags
}

func makeOperationSSHKeysGetProjectSSHKeyCmd() (*cobra.Command, error) {
	operation := GetSSHKeyOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *GetSSHKeyOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Returns a SSH Key in the project. These keys can be used to access servers after deploy and reinstall actions.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *GetSSHKeyOperation) registerFlags(cmd *cobra.Command) {
	o.PathParamsFlags = cmdflag.Flags{
		FlagSet: cmd.Flags(),
	}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "id",
			Label:       "SSH Key ID",
			Description: "SSH Key ID",
		},
		&cmdflag.String{
			Name:        "project",
			Label:       "Project ID or Slug",
			Description: "Project ID or Slug",
		},
	}

	o.PathParamsFlags.Register(schema)
}

func (o *GetSSHKeyOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := ssh_keys.NewGetProjectSSHKeyParams()
	o.PathParamsFlags.AssignValues(params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.SSHKeys.GetProjectSSHKey(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
