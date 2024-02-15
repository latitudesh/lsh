package cli

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/client/plans"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type GetPlanOperation struct {
	QueryParamsFlags cmdflag.Flags
}

func makeOperationPlansGetPlanCmd() (*cobra.Command, error) {
	operation := GetPlanOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *GetPlanOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Returns a plan",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *GetPlanOperation) registerFlags(cmd *cobra.Command) {
	o.QueryParamsFlags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "id",
			Label:       "Plan ID",
			Description: "Plan ID",
		},
	}

	o.QueryParamsFlags.Register(schema)
}

func (o *GetPlanOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := plans.NewGetPlanParams()
	o.QueryParamsFlags.AssignValues(params)

	if swag.IsZero(params.PlanID) {
		fmt.Println("Skipped action: 'id' is required")
		return nil
	}

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Plans.GetPlan(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
