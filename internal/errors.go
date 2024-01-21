package internal

import (
	"encoding/json"

	"github.com/go-openapi/swag"
)

func ParseNotFoundError(respErr error) (string, error) {
	var iResponse interface{} = respErr
	response, ok := iResponse.(*NotFoundError)

	if ok {
		if !swag.IsZero(response) && !swag.IsZero(response.Payload) {
			msgStr, err := json.Marshal(response.Payload)
			if err != nil {
				return "", err
			}
			return string(msgStr), nil
		}
	}

	return "", nil
}
