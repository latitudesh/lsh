package cli

import (
	"github.com/latitudesh/lsh/client/api_keys"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationAPIKeysDeleteAPIKeyCmd() (*cobra.Command, error) {
	operation := DeleteAPIKeyOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type DeleteAPIKeyOperation struct {
	Flags cmdflag.Flags
}

func (o *DeleteAPIKeyOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroy",
		Short: `Delete an existing API Key. Once deleted, the API Key can no longer be used to access the API.`,
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *DeleteAPIKeyOperation) PromptAttributes(attributes interface{}) {
}

func (o *DeleteAPIKeyOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:         "id",
			Description:  "ID",
			DefaultValue: "",
			Type:         "string",
		},
	}

	o.Flags.Register(schema)
}

func (o *DeleteAPIKeyOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *DeleteAPIKeyOperation) PromptID(params interface{}) {
	p := prompt.New(
		prompt.NewInputText("id", "ID from the API Key you want to delete"),
	)

	p.Run(params)
}

func (o *DeleteAPIKeyOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := api_keys.NewUpdateAPIKeyParams()
	operation.AssignResourceID(o, params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.APIKeys.UpdateAPIKey(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
