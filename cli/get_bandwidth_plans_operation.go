package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/latitudesh/lsh/client/plans"
	"github.com/latitudesh/lsh/cmd/lsh"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

// makeOperationPlansGetBandwidthPlansCmd returns a cmd to handle operation getBandwidthPlans
func makeOperationPlansGetBandwidthPlansCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list-bandwidth",
		Short: "List bandwidth package plans",
		RunE:  runOperationPlansGetBandwidthPlans,
	}

	if err := registerOperationPlansGetBandwidthPlansParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPlansGetBandwidthPlans uses cmd flags to call endpoint api
func runOperationPlansGetBandwidthPlans(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plans.NewGetBandwidthPlansParams()
	if err, _ := retrieveOperationPlansGetBandwidthPlansFilterIDFlag(params, "", cmd); err != nil {
		return err
	}
	if lsh.DryRun {

		lsh.LogDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Plans.GetBandwidthPlans(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !lsh.Debug {
		utils.Render(response.GetData())
	}
	return nil
}

// registerOperationPlansGetBandwidthPlansParamFlags registers all flags needed to fill params
func registerOperationPlansGetBandwidthPlansParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPlansGetBandwidthPlansFilterIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPlansGetBandwidthPlansFilterIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterIdDescription := `The plan ID to filter by`

	var filterIdFlagName string
	if cmdPrefix == "" {
		filterIdFlagName = "id"
	} else {
		filterIdFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var filterIdFlagDefault string

	_ = cmd.PersistentFlags().String(filterIdFlagName, filterIdFlagDefault, filterIdDescription)

	return nil
}

func retrieveOperationPlansGetBandwidthPlansFilterIDFlag(m *plans.GetBandwidthPlansParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var filterIdFlagName string
		if cmdPrefix == "" {
			filterIdFlagName = "id"
		} else {
			filterIdFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		filterIdFlagValue, err := cmd.Flags().GetString(filterIdFlagName)
		if err != nil {
			return err, false
		}
		m.FilterID = &filterIdFlagValue

	}
	return nil, retAdded
}
