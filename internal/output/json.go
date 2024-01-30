package output

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func RenderJSON(JSONResult []byte) error {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, JSONResult, "", "    "); err != nil {
		return err
	}

	fmt.Println(prettyJSON.String())
	return nil
}
