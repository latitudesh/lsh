package output

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func RenderJSON(str string) error {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return err
	}

	fmt.Println(prettyJSON.String())
	return nil
}
