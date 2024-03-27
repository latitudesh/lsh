package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"

	commands "github.com/latitudesh/lsh/internal/generator/commands/utils"
)

func main() {
	cmdToGenerate := os.Args[1:]
	if len(cmdToGenerate) == 0 {
		log.Fatal("No commands passed")
	}

	apiSpec, err := GetApiSpec("https://api.latitude.sh/api-docs/v3/swagger.json")
	if err != nil {
		log.Fatal(err)
	}

	parsedCmds := commands.ParseSpec(cmdToGenerate, apiSpec)

	for _, cmd := range parsedCmds {
		// Create Command Folder
		folderPath := path.Join("cmd", cmd.Root)
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			os.MkdirAll(folderPath, 0700)
		}

		// Generate commands inside command folder
		commands.GenerateCmd(cmd)
	}

	// Generate builder file
	commands.GenerateCmdBuilder(parsedCmds)
}

func GetApiSpec(path string) ([]byte, error) {
	res, err := http.Get(path)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}
