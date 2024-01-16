package cli

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func makeOperationLoginCmd() (*cobra.Command, error) {
	// loginCmd represents the login command
	cmd := &cobra.Command{
		Use:   "login [api-token]",
		Short: "Set your Auth Token",
		Long: `login will create a configuration file and save your API authentication 
token in it, allowing it to be used when interacting with the API.`,
		Args: cobra.MaximumNArgs(1),
		RunE: runOperationLogin,
	}

	return cmd, nil
}

func runOperationLogin(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	home, err := homedir.Dir()
	cobra.CheckErr(err)

	folderPath := path.Join(home, ".config", exeName)

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0700)
	}

	f, err := os.Create(path.Join(folderPath, "config.json"))
	if err != nil {
		return err
	}
	defer f.Close()

	viper.Set("authorization", args[0])
	viper.WriteConfig()
	fmt.Println("Success, configuration file updated!")

	return nil
}
