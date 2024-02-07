package cli

import (
	"runtime"

	"github.com/latitudesh/lsh/internal/updater"
	"github.com/spf13/cobra"
)

const (
	OS   = runtime.GOOS
	ARCH = runtime.GOARCH
)

func makeOperationUpdateCmd() (*cobra.Command, error) {
	// updateCmd represents the update command
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update the CLI to latest version",
		Long:  `update will verify if a new version is out and update the cli to it.`,
		RunE:  runOperationUpdate,
	}

	return cmd, nil
}

func runOperationUpdate(cmd *cobra.Command, args []string) error {
	err := updater.Update()
	if err != nil {
		return err
	}
	return nil
}
