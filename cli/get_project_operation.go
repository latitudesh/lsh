package cli

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/client/projects"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type GetProjectOperation struct {
	PathParamsFlags cmdflag.Flags
}

func makeOperationProjectsGetProjectCmd() (*cobra.Command, error) {
	operation := GetProjectOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *GetProjectOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Returns a project from the current team",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *GetProjectOperation) registerFlags(cmd *cobra.Command) {
	o.PathParamsFlags = cmdflag.Flags{
		FlagSet: cmd.Flags(),
	}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "id_or_slug",
			Label:       "ID or Slug",
			Description: "ID or Slug",
		},
	}

	o.PathParamsFlags.Register(schema)
}

func (o *GetProjectOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := projects.NewGetProjectParams()
	o.PathParamsFlags.AssignValues(params)

	if swag.IsZero(params.IDOrSlug) {
		fmt.Println("Skipped action: 'id_or_slug' is required")
		return nil
	}

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Projects.GetProject(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
