package renderer

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
)

type JSONRenderer struct{}

func (jr JSONRenderer) Render(data []ResponseData) {
	if !swag.IsZero(data) {
		JSONString, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Could not decode the result as JSON.")
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, JSONString, "", "    "); err != nil {
			fmt.Println("JSON format error")
		}

		fmt.Println(prettyJSON.String())
	}
}
