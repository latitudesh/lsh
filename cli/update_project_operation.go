package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/client/projects"
	"github.com/latitudesh/lsh/internal/api/resource"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

func makeOperationProjectsUpdateProjectCmd() (*cobra.Command, error) {
	operation := UpdateProjectOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type UpdateProjectOperation struct {
	PathParamFlags      cmdflag.Flags
	BodyAttributesFlags cmdflag.Flags
}

func (o *UpdateProjectOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates a project.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *UpdateProjectOperation) registerFlags(cmd *cobra.Command) {
	project := resource.NewProjectResource()

	o.PathParamFlags = cmdflag.Flags{FlagSet: cmd.Flags()}
	o.BodyAttributesFlags = cmdflag.Flags{FlagSet: cmd.Flags()}

	pathParamsSchema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "id_or_slug",
			Label:       "ID or Slug from the Project you want to update",
			Description: "The project ID or Slug",
			Required:    true,
		},
	}

	bodyFlagsSchema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "name",
			Label:       "Name",
			Description: "The project name. Must be unique.",
			Required:    false,
		},
		&cmdflag.String{
			Name:        "description",
			Label:       "Description",
			Description: "The project description",
			Required:    false,
		},
		&cmdflag.Bool{
			Name:        "bandwidth_alert",
			Label:       "Bandwidth Alert",
			Description: "Enable Bandwidth Alert",
			Required:    false,
		},
		&cmdflag.String{
			Name:        "environment",
			Label:       "Environment",
			Description: `Enum: ["Development","Staging","Production"].`,
			Options:     project.SupportedEnvironments,
			Required:    false,
		},
	}

	o.PathParamFlags.Register(pathParamsSchema)
	o.BodyAttributesFlags.Register(bodyFlagsSchema)
}

func (o *UpdateProjectOperation) PromptQueryParams(params interface{}) {
}

func (o *UpdateProjectOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := projects.NewUpdateProjectParams()
	o.PathParamFlags.AssignValues(params)
	o.BodyAttributesFlags.AssignValues(params.Body.Data.Attributes)
	params.Body.Data.ID = params.IDOrSlug

	if swag.IsZero(*params.Body.Data.Attributes) {
		fmt.Println("Skipped action: no params provided")
		return nil
	}

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Projects.UpdateProject(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
