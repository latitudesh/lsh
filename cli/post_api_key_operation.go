package cli

import (
	"github.com/latitudesh/lsh/client/api_keys"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/operation"
	"github.com/latitudesh/lsh/internal/prompt"
	"github.com/latitudesh/lsh/internal/utils"

	"github.com/spf13/cobra"
)

func makeOperationAPIKeysPostAPIKeyCmd() (*cobra.Command, error) {
	operation := CreateAPIKeyOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

type CreateAPIKeyOperation struct {
	Flags cmdflag.Flags
}

func (o *CreateAPIKeyOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "create",
		Short: `Create a new API Key that is tied to the current user account. The created API key is only listed ONCE upon creation. It can however be regenerated or deleted.`,
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *CreateAPIKeyOperation) PromptAttributes(attributes interface{}) {
	p := prompt.New(
		prompt.NewInputText("name", "Name of the API Key"),
	)

	p.Run(attributes)
}

func (o *CreateAPIKeyOperation) registerFlags(cmd *cobra.Command) {
	o.Flags = cmdflag.Flags{FlagSet: cmd.Flags()}

	schema := &cmdflag.FlagsSchema{
		{
			Name:             "name",
			Description:      "Name of the API Key",
			DefaultValue:     "",
			Type:             "string",
			RequestParamType: cmdflag.BodyParam,
		},
	}

	o.Flags.Register(schema)
}

func (o *CreateAPIKeyOperation) GetFlags() cmdflag.Flags {
	return o.Flags
}

func (o *CreateAPIKeyOperation) PromptPathParams(params interface{}) {
}

func (o *CreateAPIKeyOperation) PromptQueryParams(params interface{}) {
}

func (o *CreateAPIKeyOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := api_keys.NewPostAPIKeyParams()
	operation.AssignBodyAttributes(o, params.Body.Data.Attributes)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.APIKeys.PostAPIKey(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}
