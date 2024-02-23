package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/latitudesh/lsh/cli"
	"github.com/spf13/cobra"
)

const Ignored = 0
const General = 1

func main() {
	f, err := os.Create("commands.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rootCmd, err := cli.MakeRootCmd()
	if err != nil {
		log.Fatal("Cmd construction error: ", err)
		os.Exit(1)
	}

	for _, c := range rootCmd.Commands() {
		cmdSection := buildSection(c)
		f.WriteString(cmdSection)
	}
}

func buildSection(cmd *cobra.Command) string {
	var section strings.Builder
	header := fmt.Sprintf("## %s\n", cmd.Name())
	section.WriteString(strings.Title(strings.Replace(header, "_", " ", 1)))

	for _, subCmd := range cmd.Commands() {
		section.WriteString(fmt.Sprintf("**%s**\n", subCmd.Short))
		section.WriteString(fmt.Sprintf("```Shell\nlsh %s %s\n```\n", cmd.Name(), subCmd.Name()))
	}
	return section.String()
}
