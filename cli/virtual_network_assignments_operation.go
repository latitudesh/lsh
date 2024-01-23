package cli

import (
	"github.com/spf13/cobra"
)

func makeOperationGroupVirtualNetworkAssignmentCmd() (*cobra.Command, error) {
	operationGroupVirtualNetworkAssignmentsCmd := &cobra.Command{
		Use:   "assignments",
		Short: `Virtual Network Assignments.`,
		Long:  ``,
	}

	operationAssignServerVirtualNetworkCmd, err := makeOperationVirtualNetworkAssignmentsAssignServerVirtualNetworkCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVirtualNetworkAssignmentsCmd.AddCommand(operationAssignServerVirtualNetworkCmd)

	operationDeleteVirtualNetworksAssignmentsCmd, err := makeOperationVirtualNetworkAssignmentsDeleteVirtualNetworksAssignmentsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVirtualNetworkAssignmentsCmd.AddCommand(operationDeleteVirtualNetworksAssignmentsCmd)

	operationGetVirtualNetworksAssignmentsCmd, err := makeOperationVirtualNetworkAssignmentsGetVirtualNetworksAssignmentsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVirtualNetworkAssignmentsCmd.AddCommand(operationGetVirtualNetworksAssignmentsCmd)

	return operationGroupVirtualNetworkAssignmentsCmd, nil
}
