package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/client/plans"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationPlansGetPlanCmd returns a cmd to handle operation getPlan
func makeOperationPlansGetPlanCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get",
		Short: `Returns a plan`,
		RunE:  runOperationPlansGetPlan,
	}

	if err := registerOperationPlansGetPlanParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPlansGetPlan uses cmd flags to call endpoint api
func runOperationPlansGetPlan(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plans.NewGetPlanParams()
	if err, _ := retrieveOperationPlansGetPlanAPIVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPlansGetPlanPlanIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	result, err := appCli.Plans.GetPlan(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		result.RenderOutput()
	}
	return nil
}

// registerOperationPlansGetPlanParamFlags registers all flags needed to fill params
func registerOperationPlansGetPlanParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPlansGetPlanAPIVersionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPlansGetPlanPlanIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPlansGetPlanAPIVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	apiVersionDescription := ``

	var apiVersionFlagName string
	if cmdPrefix == "" {
		apiVersionFlagName = "API-Version"
	} else {
		apiVersionFlagName = fmt.Sprintf("%v.API-Version", cmdPrefix)
	}

	var apiVersionFlagDefault string = "2023-06-01"

	_ = cmd.PersistentFlags().String(apiVersionFlagName, apiVersionFlagDefault, apiVersionDescription)

	return nil
}
func registerOperationPlansGetPlanPlanIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	planIdDescription := `Required. `

	var planIdFlagName string
	if cmdPrefix == "" {
		planIdFlagName = "plan_id"
	} else {
		planIdFlagName = fmt.Sprintf("%v.plan_id", cmdPrefix)
	}

	var planIdFlagDefault string

	_ = cmd.PersistentFlags().String(planIdFlagName, planIdFlagDefault, planIdDescription)
	cmd.MarkPersistentFlagRequired(planIdFlagName)

	return nil
}

func retrieveOperationPlansGetPlanAPIVersionFlag(m *plans.GetPlanParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("API-Version") {

		var apiVersionFlagName string
		if cmdPrefix == "" {
			apiVersionFlagName = "API-Version"
		} else {
			apiVersionFlagName = fmt.Sprintf("%v.API-Version", cmdPrefix)
		}

		apiVersionFlagValue, err := cmd.Flags().GetString(apiVersionFlagName)
		if err != nil {
			return err, false
		}
		m.APIVersion = &apiVersionFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPlansGetPlanPlanIDFlag(m *plans.GetPlanParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("plan_id") {

		var planIdFlagName string
		if cmdPrefix == "" {
			planIdFlagName = "plan_id"
		} else {
			planIdFlagName = fmt.Sprintf("%v.plan_id", cmdPrefix)
		}

		planIdFlagValue, err := cmd.Flags().GetString(planIdFlagName)
		if err != nil {
			return err, false
		}
		m.PlanID = planIdFlagValue

	}
	return nil, retAdded
}
