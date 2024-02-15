package cli

import (
	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/api/resource"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationServersCreateServerCmd() (*cobra.Command, error) {
	operation := CreateServerOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type CreateServerOperation struct {
	Flags cmdflag.Flags
}

func (o *CreateServerOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Deploy a new server.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *CreateServerOperation) PromptAttributes(attributes interface{}) {
	server := resource.NewServerResource()

	p := prompt.New(
		prompt.NewInputText("project", "Project"),
		prompt.NewInputText("hostname", "Hostname"),
		prompt.NewInputSelect("operating_system", "Operating System", server.SupportedOperatingSystems),
		prompt.NewInputSelect("plan", "Plan", server.SupportedPlans),
		prompt.NewInputSelect("billing", "Billing", server.SupportedBillingTypes),
		prompt.NewInputSelect("site", "Site", server.SupportedSites),
		prompt.NewInputText("ipxe_url", "iPXE URL"),
		prompt.NewInputStringList("ssh_keys", "SSH Keys"),
		prompt.NewInputNumber("user_data", "User Data"),
		prompt.NewInputSelect("raid", "RAID Level", server.SupportedRAIDLevels),
	)

	p.Run(attributes)
}

func (o *CreateServerOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:         "hostname",
			Description:  "The server hostname",
			DefaultValue: "",
			Type:         "string",
		},
		{
			Name:         "billing",
			Description:  "The server billing type. Accepts 'hourly', 'monthly' or 'yearly'.",
			DefaultValue: "",
			Type:         "string",
		},
		{
			Name:         "ipxe_url",
			Description:  "URL where iPXE script is stored on, necessary for custom image deployments.This attribute is required when iPXE is selected as operating system.",
			DefaultValue: "",
			Type:         "string",
		},
		{
			Name:         "operating_system",
			Description:  `Enum: ["ipxe","windows_server_2019_std_v1","ubuntu_22_04_x64_lts","debian_11","rockylinux_8","debian_10","rhel8","centos_7_4_x64","centos_8_x64","ubuntu_20_04_x64_lts","debian_12","ubuntu22_ml_in_a_box","windows2022"]. The operating system for the new server`,
			DefaultValue: "",
			Type:         "string",
		},
		{
			Name:         "plan",
			Description:  `Enum: ["c2-large-x86","c2-medium-x86","c2-small-x86","c3-large-x86","c3-medium-x86","c3-small-x86","c3-xlarge-x86","g3-large-x86","g3-medium-x86","g3-small-x86","g3-xlarge-x86","m3-large-x86","s2-small-x86","s3-large-x86"]. The plan to choose server from`,
			DefaultValue: "",
			Type:         "string",
		},
		{
			Name:         "project",
			Description:  "The project (ID or Slug) to deploy the server",
			DefaultValue: "",
			Type:         "string",
		},
		{
			Name:         "raid",
			Description:  `Enum: ["raid-0","raid-1"]. RAID mode for the server`,
			DefaultValue: "",
			Type:         "string",
		},
		{
			Name:         "site",
			Description:  `Enum: ["ASH","BGT","BUE","CHI","DAL","FRA","LAX","LON","MEX","MEX2","MIA","MIA2","NYC","SAN","SAN2","SAO","SAO2","SYD","TYO","TYO2"]. The site to deploy the server`,
			DefaultValue: "",
			Type:         "string",
		},
		{
			Name:         "ssh_keys",
			Description:  "The SSH Keys to set on the server",
			DefaultValue: []string{},
			Type:         "stringSlice",
		},
		{
			Name:         "user_data",
			Description:  "User data to set on the server",
			DefaultValue: int64(0),
			Type:         "int64",
		},
	}

	o.Flags.Register(schema)
}

func (o *CreateServerOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *CreateServerOperation) PromptID(params interface{}) {
}

func (o *CreateServerOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := servers.NewCreateServerParams()
	operation.AssignBodyAttributes(o, params.Body.Data.Attributes)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Servers.CreateServer(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
