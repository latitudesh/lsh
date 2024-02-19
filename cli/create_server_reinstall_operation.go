package cli

import (
	"fmt"

	"github.com/latitudesh/lsh/client/server_reinstall"
	"github.com/latitudesh/lsh/internal/api/resource"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

func makeOperationServerReinstallCmd() (*cobra.Command, error) {
	operation := CreateServerReinstallOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type CreateServerReinstallOperation struct {
	Flags cmdflag.Flags
}

func (o *CreateServerReinstallOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "reinstall",
		Short: "Submit a reinstall request to a server.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *CreateServerReinstallOperation) PromptAttributes(attributes interface{}) {
	server := resource.NewServerResource()

	p := prompt.New(
		prompt.NewInputText("hostname", "Hostname"),
		prompt.NewInputSelect("operating_system", "Operating System", server.SupportedOperatingSystems),
		prompt.NewInputText("ipxe_url", "iPXE URL"),
		prompt.NewInputStringList("ssh_keys", "SSH Keys"),
		prompt.NewInputNumber("user_data", "User Data"),
		prompt.NewInputSelect("raid", "RAID Level", server.SupportedRAIDLevels),
	)

	p.Run(attributes)
}

func (o *CreateServerReinstallOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "id",
			Description:      "The Server Id (Required).",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.PathParam,
		},
		{
			Name:             "operating_system",
			Description:      `Enum: ["ipxe","windows_server_2019_std_v1","ubuntu_22_04_x64_lts","debian_11","rockylinux_8","debian_10","rhel8","centos_7_4_x64","centos_8_x64","ubuntu_20_04_x64_lts","debian_12","ubuntu22_ml_in_a_box","windows2022"]. The operating system for the new server`,
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
			Required:         false,
		},
		{
			Name:             "hostname",
			Description:      "The server hostname",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
			Required:         false,
		},
		{
			Name:             "ssh_keys",
			Description:      "The SSH Keys to set on the server",
			DefaultValue:     []string{},
			Type:             "stringSlice",
			RequestParamType: cmdflag.BodyParam,
			Required:         false,
		},
		{
			Name:             "user_data",
			Description:      "User data to set on the server",
			DefaultValue:     int64(0),
			Type:             "int64",
			RequestParamType: cmdflag.BodyParam,
			Required:         false,
		},
		{
			Name:             "raid",
			Description:      `Enum: ["raid-0","raid-1"]. RAID mode for the server`,
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
			Required:         false,
		},
		{
			Name:             "ipxe_url",
			Description:      "URL where iPXE script is stored on, necessary for custom image deployments.This attribute is required when iPXE is selected as operating system.",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
			Required:         false,
		},
	}

	o.Flags.Register(schema)
}

func (o *CreateServerReinstallOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *CreateServerReinstallOperation) PromptQueryParams(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "Server ID"),
	)

	p.Run(params)
}

func (o *CreateServerReinstallOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := server_reinstall.NewCreateServerReinstallParams()

	operation.AssignPathParams(o, params)
	operation.AssignBodyAttributes(o, params.Body.Data.Attributes)

	if swag.IsZero(*params.Body.Data.Attributes) {
		fmt.Println("Skipped action: no params provided")
		return nil
	}

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.ServerReinstall.CreateServerReinstall(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}
