package cli

import (
	"github.com/latitudesh/lsh/client/projects"
	"github.com/latitudesh/lsh/internal/api/resource"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type ListProjectsOperation struct {
	QueryParamsFlags cmdflag.Flags
}

func makeOperationProjectsGetProjectsCmd() (*cobra.Command, error) {
	operation := ListProjectsOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *ListProjectsOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Returns a list of all projects for the current team",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *ListProjectsOperation) registerFlags(cmd *cobra.Command) {
	project := resource.NewProjectResource()

	o.QueryParamsFlags = cmdflag.Flags{
		FlagSet:           cmd.Flags(),
		PromptDescription: "Filter projects: (press ENTER to skip a filter)",
	}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "name",
			Label:       "Name",
			Description: "The project name to filter by",
		},
		&cmdflag.String{
			Name:        "slug",
			Label:       "Slug",
			Description: "The project slug to filter by",
		},
		&cmdflag.String{
			Name:        "description",
			Label:       "Description",
			Description: "The project description to filter by",
		},
		&cmdflag.String{
			Name:        "billing_type",
			Label:       "Billing Type",
			Description: "The billing type to filter by",
		},
		&cmdflag.String{
			Name:        "environment",
			Label:       "Environment",
			Description: "The environment to filter by",
			Options:     project.SupportedEnvironments,
		},
	}

	o.QueryParamsFlags.Register(schema)
}

func (o *ListProjectsOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := projects.NewGetProjectsParams()
	o.QueryParamsFlags.AssignValues(params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Projects.GetProjects(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
