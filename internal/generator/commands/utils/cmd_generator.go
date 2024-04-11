package commands

import (
	"fmt"
	"path"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/latitudesh/lsh/internal/utils"
)

// GenerateCmd generates a command file.
func GenerateCmd(cmd Command) {
	cmdName := parseCmdName(cmd.Name, cmd.Root)
	f := jen.NewFile(cmd.Root)

	genNewFunc(cmd, f)
	genOperation(cmd, f)
	genRun(cmd, f)

	filepath := path.Join("cmd", cmd.Root, cmdName+".go")
	err := f.Save(filepath)
	if err != nil {
		fmt.Println(err)
	}
}

func genNewFunc(cmd Command, f *jen.File) {
	cmdName := parseCmdName(cmd.Name, cmd.Root)
	titledCmd := utils.TitleStr(cmdName)
	titledRoot := utils.TitleStr(cmd.Root)
	opName := fmt.Sprintf("%s%sOperation", titledCmd, titledRoot)

	f.Func().Id(fmt.Sprintf("New%sCmd", titledCmd)).Params().Op("*").Qual("github.com/spf13/cobra", "Command").Block(
		jen.Id("op").Op(":=").Id(opName).Block(),
		jen.Id("cmd").Op(":=").Op("&").Qual("github.com/spf13/cobra", "Command").Values(jen.Dict{
			jen.Id("Use"):   jen.Lit(cmdName),
			jen.Id("Short"): jen.Lit(cmd.Short),
			jen.Id("Long"):  jen.Lit(cmd.Long),
			jen.Id("RunE"):  jen.Id("op").Dot("run"),
		}),
		jen.Return(jen.Id("cmd")),
	)
}
func genOperation(cmd Command, f *jen.File) {
	cmdName := parseCmdName(cmd.Name, cmd.Root)
	titledCmd := utils.TitleStr(cmdName)
	titledRoot := utils.TitleStr(cmd.Root)
	opName := fmt.Sprintf("%s%sOperation", titledCmd, titledRoot)

	f.Type().Id(opName).Struct()
}

func genRun(cmd Command, f *jen.File) {
	cmdName := parseCmdName(cmd.Name, cmd.Root)
	titledCmd := utils.TitleStr(cmdName)
	titledRoot := utils.TitleStr(cmd.Root)
	opName := fmt.Sprintf("%s%sOperation", titledCmd, titledRoot)

	f.Func().Params(
		jen.Id("op").Op("*").Id(opName),
	).Id("run").Params(
		jen.Id("cmd").Op("*").Qual("github.com/spf13/cobra", "Command"),
		jen.Id("args").Index().String(),
	).Error().Block(

		jen.Id("c").Op(":=").Qual("github.com/latitudesh/lsh/cmd/lsh", "NewClient").Call(),
		jen.Line(),
		jen.If(jen.Qual("github.com/latitudesh/lsh/cmd/lsh", "DryRun")).Block(
			jen.Qual("github.com/latitudesh/lsh/cmd/lsh", "LogDebugf").Call(jen.Lit("dry-run flag specified. Skip sending request.")),
			jen.Return(jen.Nil()),
		),
		jen.Line(),
		jen.List(jen.Id(cmd.Root), jen.Id("_"), jen.Id("err")).Op(":=").Id("c").Dot(titledRoot).Dot(titledCmd).Call(jen.Nil()),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Qual("github.com/latitudesh/lsh/internal/utils", "PrintError").Call(jen.Id("err")),
			jen.Return(jen.Id("err")),
		),
		jen.Line(),
		jen.Id("lshData").Op(":=").Index().Op("*").Id(utils.Singular(titledRoot)).Block(),
		jen.For().List(jen.Id("_"), jen.Id(utils.Singular(cmd.Root))).Op(":=").Range().Id(cmd.Root).Block(
			jen.Id(fmt.Sprintf("lsh%s", utils.Singular(titledRoot))).Op(":=").Id(utils.Singular(titledRoot)).Values(
				jen.Dict{
					jen.Id("Attributes"): jen.Id(utils.Singular(cmd.Root)),
				},
			),
			jen.Id("lshData").Op("=").Append(jen.Id("lshData"), jen.Op("&").Id(fmt.Sprintf("lsh%s", utils.Singular(titledRoot)))),
		),
		jen.Line(),
		jen.Id(fmt.Sprintf("lsh%s", titledRoot)).Op(":=").Id(titledRoot).Values(
			jen.Dict{
				jen.Id("Data"): jen.Id("lshData"),
			},
		),
		jen.Line(),
		jen.If(jen.Op("!").Qual("github.com/latitudesh/lsh/cmd/lsh", "Debug")).Block(
			jen.Qual("github.com/latitudesh/lsh/internal/utils", "Render").Call(jen.Id(fmt.Sprintf("lsh%s", titledRoot)).Dot("GetData").Call()),
		),
		jen.Line(),
		jen.Return(jen.Nil()),
	)

}

func parseCmdName(name, root string) string {
	splitName := strings.Split(name, "-")
	if splitName[0] == "get" && splitName[1] == root {
		splitName[0] = "list"
	}
	return splitName[0]
}
