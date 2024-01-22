package api

import (
	"encoding/json"
	"errors"

	"github.com/go-openapi/swag"
)

func ParseErrorResponse(respErr error) (string, error) {
	switch e := respErr.(type) {
	case *Forbidden:
		return parseForbiddenError(e)
	case *NotFound:
		return parseNotFoundError(e)
	case *UnprocessableEntity:
		return parseUnprocessableEntityError(e)
	default:
		return "", errors.New("Unknown error type")
	}
}

func parseUnprocessableEntityError(respErr *UnprocessableEntity) (string, error) {
	if !swag.IsZero(respErr) && !swag.IsZero(respErr.Payload) {
		msgStr, err := json.Marshal(respErr.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

func parseNotFoundError(respErr *NotFound) (string, error) {
	if !swag.IsZero(respErr) && !swag.IsZero(respErr.Payload) {
		msgStr, err := json.Marshal(respErr.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

func parseForbiddenError(respErr *Forbidden) (string, error) {
	if !swag.IsZero(respErr) && !swag.IsZero(respErr.Payload) {
		msgStr, err := json.Marshal(respErr.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
