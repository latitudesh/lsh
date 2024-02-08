package cli

import (
	"fmt"
	"os"

	"github.com/latitudesh/lsh/client/servers"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

type DestroyServerOperation struct {
	Command *cobra.Command
	Flags   cmdflag.Flags
	Params  *servers.CreateServerParams
}

// makeOperationServersDestroyServerCmd returns a cmd to handle operation destroyServer
func makeOperationServersDestroyServerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroy",
		Short: `Remove a server.`,
	}

	registerDestroyServerOperation(cmd)

	return cmd, nil
}

func registerDestroyServerOperation(cmd *cobra.Command) {
	op := DestroyServerOperation{Command: cmd}
	op.Command.RunE = op.run
	op.RegisterFlags()
}

func (o *DestroyServerOperation) RegisterFlags() {
	o.Flags = cmdflag.Flags{
		{
			Name:         "id",
			Description:  "Required. The server ID",
			DefaultValue: "",
			Type:         "string",
		},
	}

	o.Flags.Register(o.Command.Flags())
}

func (o *DestroyServerOperation) AssignID(params interface{}) error {
	return nil
}

func (o *DestroyServerOperation) promptID(params *servers.DestroyServerParams) string {
	prompt := promptui.Prompt{
		Label: fmt.Sprintf("ID from the Server you want to delete"),
	}

	value, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return value
}

func (o *DestroyServerOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := servers.NewDestroyServerParams()
	id, err := o.Command.Flags().GetString("id")
	if err != nil {
		return err
	}

	if len(id) > 0 {
		params.SetID(id)
	} else {
		params.SetID(o.promptID(params))
	}

	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.Servers.DestroyServer(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		response.Render()
	}
	return nil
}
