package cli

import (
	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type CreateServerOperation struct {
	FlagsSchema      cmdflag.FlagsSchema
	Name             string
	ShortDescription string
	FlagSet          *pflag.FlagSet
}

func makeOperationServersCreateServerCmd() (*cobra.Command, error) {
	o := DestroyServerOperation{
		Name:             "create",
		ShortDescription: "Deploy a new server.",
	}

	cmd, err := o.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *CreateServerOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   o.Name,
		Short: o.ShortDescription,
		RunE:  o.run,
	}

	o.RegisterFlags(cmd)

	return cmd, nil
}

func (o *CreateServerOperation) PromptAttributes(attributes interface{}) {
	p := prompt.New(
		prompt.NewInputText("project", "Project"),
		prompt.NewInputText("hostname", "Hostname"),
		prompt.NewInputSelect("operating_system", "Operating System", o.supportedOperatingSystems()),
		prompt.NewInputSelect("plan", "Plan", o.supportedPlans()),
		prompt.NewInputSelect("billing", "Billing", o.supportedBillingTypes()),
		prompt.NewInputSelect("site", "Site", o.supportedSites()),
		prompt.NewInputText("ipxe_url", "iPXE URL"),
		prompt.NewInputList("ssh_keys", "SSH Keys"),
		prompt.NewInputText("user_data", "User Data"),
		prompt.NewInputSelect("raid", "RAID Level", o.supportedRAIDLevels()),
	)

	p.Run(attributes)
}

func (o *CreateServerOperation) RegisterFlags(cmd *cobra.Command) {
	o.FlagsSchema = cmdflag.FlagsSchema{
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
			Name:        "user_data",
			Description: "User data to set on the server",
			Type:        "int64",
		},
	}

	o.FlagsSchema.Register(o.FlagSet)
}

func (o *CreateServerOperation) GetFlagsDefinition() cmdflag.FlagsSchema {
	return o.FlagsSchema
}

func (o *CreateServerOperation) PromptID(params interface{}) {
}

func (o *CreateServerOperation) ResourceIDFlag() cmdflag.FlagSchema {
	return o.FlagsSchema[0]
}

func (o *CreateServerOperation) GetFlagSet() *pflag.FlagSet {
	return o.FlagSet
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
		response.Render()
	}
	return nil
}

func (o *CreateServerOperation) supportedOperatingSystems() []string {
	return []string{
		"ipxe",
		"windows_server_2019_std_v1",
		"ubuntu_22_04_x64_lts",
		"debian_11",
		"rockylinux_8",
		"debian_10",
		"rhel8",
		"centos_7_4_x64",
		"centos_8_x64",
		"ubuntu_20_04_x64_lts",
		"debian_12",
		"ubuntu22_ml_in_a_box",
		"windows2022",
	}
}

func (o *CreateServerOperation) supportedPlans() []string {
	return []string{
		"c2-large-x86",
		"c2-medium-x86",
		"c2-small-x86",
		"c3-large-x86",
		"c3-medium-x86",
		"c3-small-x86",
		"c3-xlarge-x86",
		"g3-large-x86",
		"g3-medium-x86",
		"g3-small-x86",
		"g3-xlarge-x86",
		"m3-large-x86",
		"s2-small-x86",
		"s3-large-x86",
	}
}

func (o *CreateServerOperation) supportedSites() []string {
	return []string{
		"ASH",
		"BGT",
		"BUE",
		"CHI",
		"DAL",
		"FRA",
		"LAX",
		"LON",
		"MEX",
		"MEX2",
		"MIA",
		"MIA2",
		"NYC",
		"SAN",
		"SAN2",
		"SAO",
		"SAO2",
		"SYD",
		"TYO",
		"TYO2",
	}
}

func (o *CreateServerOperation) supportedBillingTypes() []string {
	return []string{"hourly", "monthly", "yearly"}
}

func (o *CreateServerOperation) supportedRAIDLevels() []string {
	return []string{"raid-0", "raid-1"}
}
