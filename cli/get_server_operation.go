package cli

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

type GetServerOperation struct {
	PathParamsFlags cmdflag.Flags
}

func makeOperationServersGetServerCmd() (*cobra.Command, error) {
	operation := GetServerOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *GetServerOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Returns a server that belongs to the team.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *GetServerOperation) registerFlags(cmd *cobra.Command) {
	o.PathParamsFlags = cmdflag.Flags{
		FlagSet: cmd.Flags(),
	}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "id",
			Label:       "ID",
			Description: "Server ID",
		},
	}

	o.PathParamsFlags.Register(schema)
}

func (o *GetServerOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := servers.NewGetServerParams()
	o.PathParamsFlags.AssignValues(params)

	if swag.IsZero(params.ServerID) {
		fmt.Println("Skipped action: 'id' is required")
		return nil
	}

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Servers.GetServer(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
