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
	genFlags(cmd, f)
	genPreRun(cmd, f)
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
			jen.Id("Use"):    jen.Lit(cmdName),
			jen.Id("Short"):  jen.Lit(cmd.Short),
			jen.Id("Long"):   jen.Lit(cmd.Long),
			jen.Id("PreRun"): jen.Id("op").Dot("preRun"),
			jen.Id("RunE"):   jen.Id("op").Dot("run"),
		}),

		jen.Id("op").Dot("registerFlags").Call(jen.Id("cmd")),

		jen.Return(jen.Id("cmd")),
	)
}

func genOperation(cmd Command, f *jen.File) {
	cmdName := parseCmdName(cmd.Name, cmd.Root)
	titledCmd := utils.TitleStr(cmdName)
	titledRoot := utils.TitleStr(cmd.Root)
	opName := fmt.Sprintf("%s%sOperation", titledCmd, titledRoot)

	jenParams := []jen.Code{}

	switch cmdName {
	case "list":
	case "get", "destroy":
		jenParams = append(jenParams, jen.Id("PathParamFlags").Qual("github.com/latitudesh/lsh/internal/cmdflag", "Flags"))
	case "create":
		jenParams = append(jenParams, jen.Id("BodyAttributesFlags").Qual("github.com/latitudesh/lsh/internal/cmdflag", "Flags"))
	default:
		jenParams = append(jenParams,
			jen.Id("PathParamFlags").Qual("github.com/latitudesh/lsh/internal/cmdflag", "Flags"),
			jen.Id("BodyAttributesFlags").Qual("github.com/latitudesh/lsh/internal/cmdflag", "Flags"),
		)
	}

	f.Type().Id(opName).Struct(jenParams...)
}

func genFlags(cmd Command, f *jen.File) {
	cmdName := parseCmdName(cmd.Name, cmd.Root)
	titledCmd := utils.TitleStr(cmdName)
	titledRoot := utils.TitleStr(cmd.Root)
	opName := fmt.Sprintf("%s%sOperation", titledCmd, titledRoot)

	fmt.Println(cmd.Parameters)

	f.Func().Params(
		jen.Id("op").Op("*").Id(opName),
	).Id("registerFlags").Params(
		jen.Id("cmd").Op("*").Qual("github.com/spf13/cobra", "Command"),
	).Block(
		jen.Id("op").Dot("PathParamFlags").Op("=").Qual("github.com/latitudesh/lsh/internal/cmdflag", "Flags").Values(jen.Dict{
			jen.Id("FlagSet"): jen.Id("cmd").Dot("Flags").Call(),
		}),
		jen.Id("schema").Op(":=").Op("&").Qual("github.com/latitudesh/lsh/internal/cmdflag", "FlagsSchema").Values(
			jen.Op("&").Qual("github.com/latitudesh/lsh/internal/cmdflag", "String").Values(jen.Dict{
				jen.Id("Name"):        jen.Lit(cmdName),
				jen.Id("Label"):       jen.Lit(cmdName),
				jen.Id("Description"): jen.Lit(cmdName),
				jen.Id("Required"):    jen.Lit(false),
			}),
		),
		jen.Id("op").Dot("PathParamFlags").Dot("Register").Call(jen.Id("schema")),
	)
}

func genPreRun(cmd Command, f *jen.File) {
	cmdName := parseCmdName(cmd.Name, cmd.Root)
	titledCmd := utils.TitleStr(cmdName)
	titledRoot := utils.TitleStr(cmd.Root)
	opName := fmt.Sprintf("%s%sOperation", titledCmd, titledRoot)

	f.Func().Params(
		jen.Id("op").Op("*").Id(opName),
	).Id("preRun").Params(
		jen.Id("cmd").Op("*").Qual("github.com/spf13/cobra", "Command"),
		jen.Id("args").Index().String(),
	).Block(
		jen.Id("op").Dot("PathParamFlags").Dot("PreRun").Call(jen.Id("cmd"), jen.Id("args")),
	)
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

		genkSdkFunc(cmd),

		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Qual("github.com/latitudesh/lsh/internal/utils", "PrintError").Call(jen.Id("err")),
			jen.Return(jen.Id("err")),
		),
		jen.Line(),

		genDataSection(cmd),

		jen.Line(),

		genRenderSection(cmd),

		jen.Line(),
		jen.Return(jen.Nil()),
	)

}

func genkSdkFunc(cmd Command) *jen.Statement {
	cmdName := parseCmdName(cmd.Name, cmd.Root)
	titledRoot := utils.TitleStr(cmd.Root)
	titledCmd := utils.TitleStr(cmdName)

	sdkSection := jen.Statement{}

	switch cmdName {
	case "get", "update":
		sdkSection = jen.Statement{
			jen.Id("attr").Op(":=").Struct(jen.Id("ID").String().Tag(map[string]string{"json": "id"})).Block(),
			jen.Line(),
			jen.Id("op").Dot("PathParamFlags").Dot("AssignValues").Call(jen.Op("&").Id("attr")),
			jen.Line(),
			jen.List(jen.Id(cmd.Root), jen.Id("_"), jen.Id("err")).Op(":=").Id("c").Dot(titledRoot).Dot(titledCmd).Call(jen.Id("attr").Dot("ID"), jen.Nil()),
		}
	case "destroy":
		sdkSection = jen.Statement{
			jen.Id("attr").Op(":=").Struct(jen.Id("ID").String().Tag(map[string]string{"json": "id"})).Block(),
			jen.Line(),
			jen.Id("op").Dot("PathParamFlags").Dot("AssignValues").Call(jen.Op("&").Id("attr")),
			jen.Line(),
			jen.List(jen.Id("resp"), jen.Id("err")).Op(":=").Id("c").Dot(titledRoot).Dot("Delete").Call(jen.Id("attr").Dot("ID")),
		}
	default:
		sdkSection = jen.Statement{jen.List(jen.Id(cmd.Root), jen.Id("_"), jen.Id("err")).Op(":=").Id("c").Dot(titledRoot).Dot(titledCmd).Call(jen.Nil())}
	}

	return &sdkSection
}

func genDataSection(cmd Command) *jen.Statement {
	cmdName := parseCmdName(cmd.Name, cmd.Root)
	titledRoot := utils.TitleStr(cmd.Root)

	dataSection := jen.Statement{}

	switch cmdName {
	case "list":
		dataSection = jen.Statement{
			jen.Id("lshData").Op(":=").Index().Op("*").Id(utils.Singular(titledRoot)).Block(),
			jen.Line(),
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
		}
	case "destroy":
	default:
		dataSection = jen.Statement{
			jen.Id(fmt.Sprintf("lsh%s", titledRoot)).Op(":=").Id(utils.Singular(titledRoot)).Values(
				jen.Dict{
					jen.Id("Attributes"): jen.Op("*").Id(cmd.Root),
				},
			),
		}
	}

	return &dataSection
}

func genRenderSection(cmd Command) *jen.Statement {
	cmdName := parseCmdName(cmd.Name, cmd.Root)
	titledRoot := utils.TitleStr(cmd.Root)

	renderSection := jen.Statement{}

	switch cmdName {
	case "destroy":
		renderSection = jen.Statement{
			jen.If(jen.Op("!").Qual("github.com/latitudesh/lsh/cmd/lsh", "Debug")).Block(
				jen.If(jen.Id("resp").Dot("StatusCode").Op("!=").Qual("net/http", "StatusNoContent")).Block(
					jen.Qual("log", "Fatal").Call(jen.Lit("Something went wrong while deleting resource.")),
				),
				jen.Qual("fmt", "Printf").Call(jen.Lit(fmt.Sprintf("\n\n%s deleted successfully!\n\n", utils.Singular(titledRoot)))),
			),
		}
	default:
		renderSection = jen.Statement{
			jen.If(jen.Op("!").Qual("github.com/latitudesh/lsh/cmd/lsh", "Debug")).Block(
				jen.Qual("github.com/latitudesh/lsh/internal/utils", "Render").Call(jen.Id(fmt.Sprintf("lsh%s", titledRoot)).Dot("GetData").Call()),
			),
		}
	}

	return &renderSection
}

func parseCmdName(name, root string) string {
	splitName := strings.Split(name, "-")
	if splitName[0] == "get" && splitName[1] == root {
		splitName[0] = "list"
	}
	return splitName[0]
}
