package cli

import (
	"fmt"

	"github.com/latitudesh/lsh/client/virtual_networks"
	"github.com/latitudesh/lsh/internal/cmdflag"
	"github.com/latitudesh/lsh/internal/utils"
	"github.com/latitudesh/lsh/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

type GetVirtualNetworkOperation struct {
	PathParamsFlags cmdflag.Flags
}

func makeOperationVirtualNetworksGetVirtualNetworkCmd() (*cobra.Command, error) {
	operation := GetVirtualNetworkOperation{}

	cmd, err := operation.Register()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (o *GetVirtualNetworkOperation) Register() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve a Virtual Network.",
		RunE:  o.run,
	}

	o.registerFlags(cmd)

	return cmd, nil
}

func (o *GetVirtualNetworkOperation) registerFlags(cmd *cobra.Command) {
	o.PathParamsFlags = cmdflag.Flags{
		FlagSet: cmd.Flags(),
	}

	schema := &cmdflag.FlagsSchema{
		&cmdflag.String{
			Name:        "id",
			Label:       "ID",
			Description: "ID",
		},
	}

	o.PathParamsFlags.Register(schema)
}

func (o *GetVirtualNetworkOperation) run(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}

	params := virtual_networks.NewGetVirtualNetworkParams()
	o.PathParamsFlags.AssignValues(params)

	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.VirtualNetworks.GetVirtualNetwork(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}

// runOperationVirtualNetworksGetVirtualNetwork uses cmd flags to call endpoint api
func runOperationVirtualNetworksGetVirtualNetwork(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := virtual_networks.NewGetVirtualNetworkParams()
	if err, _ := retrieveOperationVirtualNetworksGetVirtualNetworkIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.VirtualNetworks.GetVirtualNetwork(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetData())
	}
	return nil
}

// registerOperationVirtualNetworksGetVirtualNetworkParamFlags registers all flags needed to fill params
func registerOperationVirtualNetworksGetVirtualNetworkParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVirtualNetworksGetVirtualNetworkIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVirtualNetworksGetVirtualNetworkIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)
	cmd.MarkPersistentFlagRequired(idFlagName)

	return nil
}

func retrieveOperationVirtualNetworksGetVirtualNetworkIDFlag(m *virtual_networks.GetVirtualNetworkParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// register flags to command
func registerModelGetVirtualNetworkOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetVirtualNetworkOKBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetVirtualNetworkOKBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelVirtualNetworkFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetVirtualNetworkOKBodyFlags(depth int, m *virtual_networks.GetVirtualNetworkOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrieveGetVirtualNetworkOKBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrieveGetVirtualNetworkOKBodyDataFlags(depth int, m *virtual_networks.GetVirtualNetworkOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data models.VirtualNetwork is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.VirtualNetwork{}
	}

	err, dataAdded := retrieveModelVirtualNetworkFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}
