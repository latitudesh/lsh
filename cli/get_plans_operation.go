package cli

import (
	"github.com/latitudesh/lsh/client/plans"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type ListPlansOperation struct {
	QueryParamsFlags cmdflag.Flags
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

func (o *ListPlansOperation) registerFlags(cmd *cobra.Command) {
	o.QueryParamsFlags = cmdflag.Flags{
		FlagSet:           cmd.Flags(),
		PromptDescription: "Filter plans: (press ENTER to skip a filter)",
	}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.Int64{
			Name:        "disk_eql",
			Label:       "Disk size (GB)",
			Description: "Filter plans with disk size in Gigabytes equals the provided value.",
		},
		&cmdflag.Int64{
			Name:        "disk_gte",
			Label:       "Disk (GB) greater than or equal to",
			Description: "Filter plans with disk size in Gigabytes greater than or equal the provided value.",
		},
		&cmdflag.Int64{
			Name:        "disk_lte",
			Label:       "Disk (GB) less than or equal to",
			Description: "Filter plans with disk size in Gigabytes less than or equal the provided value.",
		},
		&cmdflag.Bool{
			Name:        "gpu",
			Label:       "With GPU",
			Description: "Filter by the existence of an associated GPU",
		},
		&cmdflag.Bool{
			Name:        "in_stock",
			Label:       "In Stock",
			Description: "The stock available at the site to filter by",
		},
		&cmdflag.String{
			Name:        "location",
			Label:       "Location",
			Description: "The location of the site to filter by",
		},
		&cmdflag.String{
			Name:        "name",
			Label:       "Name",
			Description: "The plan name to filter by",
		},
		&cmdflag.Int64{
			Name:        "ram_eql",
			Label:       "RAM (GB)",
			Description: "Filter plans with RAM size (in GB) equals the provided value.",
		},
		&cmdflag.Int64{
			Name:        "ram_gte",
			Label:       "RAM (GB) greater than or equal to",
			Description: "Filter plans with RAM size (in GB) greater than or equal the provided value.",
		},
		&cmdflag.Int64{
			Name:        "ram_lte",
			Label:       "RAM (GB) less than or equal to",
			Description: "Filter plans with RAM size (in GB) less than or equal the provided value.",
		},
		&cmdflag.String{
			Name:        "slug",
			Label:       "Slug",
			Description: "The plan slug to filter by",
		},
		&cmdflag.String{
			Name:        "stock_level",
			Label:       "Stock Level",
			Description: `Enum: ["Unavailable","Low","Medium","High","Unique"]. The stock level at the site to filter by`,
		},
	}

	o.QueryParamsFlags.Register(schema)
}

func (o *ListPlansOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := plans.NewGetPlansParams()
	o.QueryParamsFlags.AssignValues(params)

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
