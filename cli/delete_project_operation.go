package cli

import (
	"github.com/latitudesh/lsh/client/projects"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type DeleteProjectOperation struct {
	Flags cmdflag.Flags
}

func makeOperationProjectsDeleteProjectCmd() (*cobra.Command, error) {
	operation := DeleteProjectOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *DeleteProjectOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroy",
		Short: `Deletes a project from the current team`,
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *DeleteProjectOperation) PromptID(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id_or_slug", "ID or Slug from the Project you want to destroy"),
	)

	p.Run(params)
}

func (o *DeleteProjectOperation) PromptAttributes(attributes interface{}) {
}

func (o *DeleteProjectOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *DeleteProjectOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:         "id_or_slug",
			Description:  "Required. The project ID or Slug",
			DefaultValue: "",
			Type:         "string",
		},
	}

	o.Flags.Register(schema)
}

func (o *DeleteProjectOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := projects.NewDeleteProjectParams()
	operation.AssignResourceID(o, params)
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Projects.DeleteProject(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}
