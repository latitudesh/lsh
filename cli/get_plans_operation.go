package cli

import (
	"github.com/latitudesh/lsh/client/plans"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type ListPlansOperation struct {
	Flags cmdflag.Flags
}

func makeOperationPlansGetPlansCmd() (*cobra.Command, error) {
	operation := ListPlansOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *ListPlansOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: `Lists all plans. Availability by region is included in ` + "`" + `attributes.regions.locations.available[*]` + "`" + ` node for a given plan.`,
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *ListPlansOperation) PromptPathParams(params interface{}) {
}

func (o *ListPlansOperation) PromptQueryParams(params interface{}) {
	// p := prompt.New(
	// 	prompt.NewInputText("id", "ID"),
	// )

	// p.Run(params)
}

func (o *ListPlansOperation) PromptAttributes(attributes interface{}) {
}

func (o *ListPlansOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *ListPlansOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "disk_eql",
			Description:      "Filter plans with disk size in Gigabytes equals the provided value.",
			DefaultValue:     int64(0),
			Type:             "int64",
			RequestParamType: cmdflag.QueryParam,
		},
		{
			Name:             "disk_gte",
			Description:      "Filter plans with disk size in Gigabytes greater than or equal the provided value.",
			DefaultValue:     int64(0),
			Type:             "int64",
			RequestParamType: cmdflag.QueryParam,
		},
		{
			Name:             "disk_lte",
			Description:      "Filter plans with disk size in Gigabytes less than or equal the provided value.",
			DefaultValue:     int64(0),
			Type:             "int64",
			RequestParamType: cmdflag.QueryParam,
		},
	}

	o.Flags.Register(schema)
}

func (o *ListPlansOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := plans.NewGetPlansParams()
	operation.AssignPathParams(o, params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Plans.GetPlans(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
