package cli

import (
	"github.com/latitudesh/lsh/client/projects"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationProjectsCreateProjectCmd() (*cobra.Command, error) {
	operation := CreateProjectOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type CreateProjectOperation struct {
	BodyAttributesFlags cmdflag.Flags
}

func (o *CreateProjectOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a project.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *CreateProjectOperation) registerFlags(cmd *cobra.Command) {
	o.BodyAttributesFlags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "name",
			Label:       "Name",
			Description: "The project name. Must be unique.",
			Required:    true,
		},
		&cmdflag.String{
			Name:        "provisioning_type",
			Label:       "Provisioning Type",
			Description: `Enum: ["reserved","on_demand"]. The provisioning type of the project. Default: on_demand`,
			Required:    true,
		},
		&cmdflag.String{
			Name:        "description",
			Label:       "Description",
			Description: "The project description",
			Required:    false,
		},
		&cmdflag.String{
			Name:        "environment",
			Label:       "Environment",
			Description: `Enum: ["Development","Staging","Production"].`,
			Required:    false,
		},
	}

	o.BodyAttributesFlags.Register(schema)
}

func (o *CreateProjectOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := projects.NewCreateProjectParams()
	o.BodyAttributesFlags.AssignValues(params.Body.Data.Attributes)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Projects.CreateProject(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
