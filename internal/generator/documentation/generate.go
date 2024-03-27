package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/latitudesh/lsh/cli"
	"github.com/spf13/cobra"
)

const IGNORED = 0
const GENERAL = 1

// Map to Control the output
var cmdSelection = map[string]int{
	"completion": 0,
	"update":     1,
	"login":      1,
}

func main() {
	rootCmd, err := cli.MakeRootCmd()
	if err != nil {
		log.Fatal("Cmd construction error: ", err)
		os.Exit(1)
	}

	f, err := os.Create("commands.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	generalCmds := []*cobra.Command{}
	lshCmds := []*cobra.Command{}

	for _, c := range rootCmd.Commands() {
		value, ok := cmdSelection[c.Name()]
		if ok {
			if value == IGNORED {
				continue
			}
			if value == GENERAL {
				generalCmds = append(generalCmds, c)
				continue
			}
		}
		lshCmds = append(lshCmds, c)
	}

	// If command is a general command
	if len(generalCmds) != 0 {
		generalSection := buildGeneral(generalCmds)
		f.WriteString(generalSection)
	}

	// If command is a normal CLI command
	if len(lshCmds) != 0 {
		for _, cmd := range lshCmds {
			cmdSection := buildSection(cmd)
			f.WriteString(cmdSection)
		}
	}
}

func buildSection(cmd *cobra.Command) string {
	var section strings.Builder
	header := fmt.Sprintf("## %s\n", cmd.Name())
	section.WriteString(strings.Title(strings.Replace(header, "_", " ", 1)))

	subCmdString := ""
	for _, subCmd := range cmd.Commands() {
		// Verify if command has sub-commands
		if len(subCmd.Commands()) != 0 {
			subCmdString = buildSubSection(subCmd)
			continue
		}

		section.WriteString(fmt.Sprintf("**%s**\n", subCmd.Short))
		section.WriteString(fmt.Sprintf("```Shell\nlsh %s %s\n```\n", cmd.Name(), subCmd.Name()))
	}
	section.WriteString(subCmdString)

	return section.String()
}

func buildSubSection(cmd *cobra.Command) string {
	var section strings.Builder
	header := fmt.Sprintf("## %s %s\n", cmd.Parent().Name(), cmd.Name())
	section.WriteString(strings.Title(strings.Replace(header, "_", " ", 1)))

	for _, subCmd := range cmd.Commands() {
		section.WriteString(fmt.Sprintf("**%s**\n", subCmd.Short))
		section.WriteString(fmt.Sprintf("```Shell\nlsh %s %s %s\n```\n", cmd.Name(), cmd.Parent().Name(), subCmd.Name()))
	}
	return section.String()
}

func buildGeneral(cmds []*cobra.Command) string {
	var section strings.Builder
	header := "## General\n"
	section.WriteString(header)
	for _, cmd := range cmds {
		section.WriteString(fmt.Sprintf("**%s**\n", cmd.Short))
		section.WriteString(fmt.Sprintf("```Shell\nlsh %s\n```\n", cmd.Name()))
	}
	return section.String()
}
