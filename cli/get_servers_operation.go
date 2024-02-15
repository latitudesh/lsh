package cli

import (
	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/api/resource"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type ListServersOperation struct {
	QueryParamsFlags cmdflag.Flags
}

func makeOperationServersGetServersCmd() (*cobra.Command, error) {
	operation := ListServersOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *ListServersOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Returns a list of all servers belonging to the team.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *ListServersOperation) registerFlags(cmd *cobra.Command) {
	server := resource.NewServerResource()

	o.QueryParamsFlags = cmdflag.Flags{
		FlagSet:           cmd.Flags(),
		PromptDescription: "Filter servers: (press ENTER to skip a filter)",
	}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "project",
			Label:       "Project ID or Slug",
			Description: "The project ID or Slug to filter by",
		},
		&cmdflag.String{
			Name:        "region",
			Label:       "Region",
			Description: "The region Slug to filter by",
			Options:     server.SupportedSites,
		},
		&cmdflag.String{
			Name:        "hostname",
			Label:       "Hostname",
			Description: "The hostname of server to filter by",
		},
		&cmdflag.String{
			Name:        "created_at_gte",
			Label:       "Created from",
			Description: "The created at greater than equal date to filter by",
		},
		&cmdflag.String{
			Name:        "created_at_lte",
			Label:       "Created until",
			Description: "The created at less than equal date to filter by",
		},
		&cmdflag.String{
			Name:        "label",
			Label:       "Label",
			Description: "The label of server to filter by",
		},
		&cmdflag.String{
			Name:        "status",
			Label:       "Status",
			Description: "The status of server to filter by",
		},
		&cmdflag.String{
			Name:        "operating_system",
			Label:       "Operating System",
			Description: "The operating system name or slug of the server to filter by",
			Options:     server.SupportedOperatingSystems,
		},
		&cmdflag.String{
			Name:        "plan",
			Label:       "Plan",
			Description: "The platform/plan name of the server to filter by",
			Options:     server.SupportedPlans,
		},
		&cmdflag.Bool{
			Name:        "gpu",
			Label:       "GPU",
			Description: "Filter by the existence of an associated GPU",
		},
		&cmdflag.Int64{
			Name:        "ram_eql",
			Label:       "RAM (GB)",
			Description: "Filter servers with RAM size (in GB) equals the provided value.",
		},
		&cmdflag.Int64{
			Name:        "ram_gte",
			Label:       "RAM (GB) greater than or equal to",
			Description: "Filter servers with RAM size (in GB) greater than or equal the provided value.",
		},
		&cmdflag.Int64{
			Name:        "ram_lte",
			Label:       "RAM (GB) less than or equal to",
			Description: "Filter servers with RAM size (in GB) less than or equal the provided value.",
		},
		&cmdflag.Int64{
			Name:        "disk_eql",
			Label:       "Disk size (GB)",
			Description: "Filter servers with disk size in Gigabytes equals the provided value.",
		},
		&cmdflag.Int64{
			Name:        "disk_gte",
			Label:       "Disk (GB) greater than or equal to",
			Description: "Filter servers with disk size in Gigabytes greater than or equal the provided value.",
		},
		&cmdflag.Int64{
			Name:        "disk_lte",
			Label:       "Disk (GB) less than or equal to",
			Description: "Filter servers with disk size in Gigabytes less than or equal the provided value.",
		},
	}

	o.QueryParamsFlags.Register(schema)
}

func (o *ListServersOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := servers.NewGetServersParams()
	o.QueryParamsFlags.AssignValues(params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Servers.GetServers(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
