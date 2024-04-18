package commands

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/pb33f/libopenapi"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
)

// var rootName = "tags"

func ParseSpec(commands []string, spec []byte) []Command {
	document, err := libopenapi.NewDocument(spec)
	if err != nil {
		log.Fatal(err)
	}

	// Build document
	model, errors := document.BuildV3Model()
	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])
		}
		log.Fatalf("cannot create v3 model from "+
			"document: %d errors reported",
			len(errors))
	}

	cmdToGenerate := []Command{}

	for _, cmd := range commands {
		// Get paths for command
		// ie: /tags and /tags/{tag_id} belong to the same CLI command
		relatedPaths := []string{}
		allPaths := model.Index.GetAllPaths()
		relatedPaths = append(relatedPaths, "/"+cmd)

		rx := regexp.MustCompile(fmt.Sprintf("%s/{.*}$", cmd))
		for k := range allPaths {
			if rx.MatchString(k) {
				relatedPaths = append(relatedPaths, k)
			}
		}

		// Context for ordered map iteration
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Get info about root command
		for _, subC := range relatedPaths {
			path := model.Model.Paths.PathItems.GetOrZero(subC)

			c := orderedmap.Iterate(ctx, path.GetOperations())
			for pr := range c {

				cmd := Command{
					Name:       pr.Value().OperationId,
					Short:      pr.Value().Summary,
					Long:       pr.Value().Description,
					Method:     pr.Key(),
					Root:       cmd,
					Parameters: parseParameters(pr.Value()),
				}
				cmdToGenerate = append(cmdToGenerate, cmd)
			}
		}
	}
	return cmdToGenerate
}

func parseParameters(params *v3.Operation) []CmdParameter {
	result := []CmdParameter{}

	for _, param := range params.Parameters {
		param := CmdParameter{
			Name:        param.Name,
			Description: param.Description,
			Required:    *param.Required,
		}
		result = append(result, param)
	}

	return result
}
