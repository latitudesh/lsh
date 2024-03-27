package commands

import (
	"fmt"
	"path"
	"strings"

	"github.com/dave/jennifer/jen"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// GenerateCmd generates a command file.
func GenerateCmd(cmd Command) {
	rootName := cmd.Root
	f := jen.NewFile(rootName)

	caser := cases.Title(language.English)
	cmdName := parseCmdName(cmd.Name, rootName)
	casedName := caser.String(cmdName)

	f.Var().Id(fmt.Sprintf("%sCmd", casedName)).Op("=").Op("&").Qual("github.com/spf13/cobra", "Command").Values(jen.Dict{
		jen.Id("Use"):   jen.Lit(cmdName),
		jen.Id("Short"): jen.Lit(cmd.Short),
		jen.Id("Long"):  jen.Lit(cmd.Long),
		jen.Id("Run"): jen.Func().Params(jen.Id("cmd").Op("*").Qual("github.com/spf13/cobra", "Command"), jen.Id("args").Index().String()).Block(
			jen.Qual("fmt", "Println").Call(jen.Lit(fmt.Sprintf("%s tag called", cmdName))),
		),
	})

	filepath := path.Join("cmd", rootName, cmdName+".go")
	err := f.Save(filepath)
	if err != nil {
		fmt.Println(err)
	}
}

// GenerateCmdBuilder generates the Command builder file.
// The Command builder file is responsible for adding all the subcomands to the main command and them add the main command to the CLI.
func GenerateCmdBuilder(cmds []Command) {
	rootName := cmds[0].Root
	root := fmt.Sprintf("%sCmd", rootName)
	caser := cases.Title(language.English)

	f := jen.NewFile("cmd")
	init := f.Func().Id("init").Params()

	jenCmd := []jen.Code{}
	for _, cmd := range cmds {
		cmdName := parseCmdName(cmd.Name, rootName)
		jenCmd = append(jenCmd, jen.Id(root).Dot("AddCommand").Call(
			jen.Qual(
				"github.com/latitudesh/lsh/cmd/tags",
				fmt.Sprintf("%sCmd", caser.String(cmdName)),
			),
		))
	}
	jenCmd = append(jenCmd, jen.Id("rootCmd").Dot("AddCommand").Call(jen.Id(root)))

	block := jen.Block(jenCmd...)
	init.Add(block)

	f.Var().Id(root).Op("=").Op("&").Qual("github.com/spf13/cobra", "Command").Values(jen.Dict{
		jen.Id("Use"): jen.Lit(rootName),
	})

	fileName := fmt.Sprintf("%s_builder.go", rootName)

	err := f.Save(path.Join("cmd", fileName))
	if err != nil {
		fmt.Println(err)
	}
}

func parseCmdName(name, root string) string {
	splitName := strings.Split(name, "-")
	if splitName[0] == "get" && splitName[1] == root {
		splitName[0] = "list"
	}
	return splitName[0]
}
