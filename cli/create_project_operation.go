package cli

import (
	"github.com/latitudesh/lsh/client/projects"
	"github.com/latitudesh/lsh/internal/api/resource"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
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
	Flags cmdflag.Flags
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

func (o *CreateProjectOperation) PromptAttributes(attributes interface{}) {
	project := resource.NewProjectResource()

	p := prompt.New(
		prompt.NewInputText("name", "Name"),
		prompt.NewInputText("description", "Description"),
		prompt.NewInputSelect("environment", "Environment", project.SupportedEnvironments),
		prompt.NewInputSelect("provisioning_type", "Provisioning Type", project.SupportedProvisioningTypes),
	)

	p.Run(attributes)
}

func (o *CreateProjectOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "name",
			Description:      "Required. The project name. Must be unique.",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
		{
			Name:             "description",
			Description:      "The project description",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
		{
			Name:             "environment",
			Description:      `Enum: ["Development","Staging","Production"].`,
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
		{
			Name:             "provisioning_type",
			Description:      `Enum: ["reserved","on_demand"]. Required. The provisioning type of the project. Default: on_demand`,
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
	}

	o.Flags.Register(schema)
}

func (o *CreateProjectOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *CreateProjectOperation) PromptQueryParams(params interface{}) {
}

func (o *CreateProjectOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := projects.NewCreateProjectParams()
	operation.AssignBodyAttributes(o, params.Body.Data.Attributes)

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
