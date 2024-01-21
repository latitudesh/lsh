package api

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/latitudesh/cli/models"
)

// NewErrorResponse creates a ErrorResponse with default headers values
func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}

/*
ErrorResponse describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ErrorResponse struct {
	Payload *models.ErrorObject
}

func (o *ErrorResponse) IsSuccess() bool {
	return false
}

func (o *ErrorResponse) IsRedirect() bool {
	return false
}

func (o *ErrorResponse) IsClientError() bool {
	return true
}

func (o *ErrorResponse) IsServerError() bool {
	return false
}

func (o *ErrorResponse) IsCode(code int) bool {
	return code == 422
}

func (o *ErrorResponse) Code() int {
	return 422
}

func (o *ErrorResponse) Error() string {
	return fmt.Sprintf("[%d] ErrorResponse  %+v", 422, o.Payload)
}

func (o *ErrorResponse) String() string {
	return fmt.Sprintf("[%d] ErrorResponse  %+v", 422, o.Payload)
}

func (o *ErrorResponse) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *ErrorResponse) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func ParseErrorResponse(respErr error) (string, error) {
	var iResponse interface{} = respErr
	response, ok := iResponse.(*ErrorResponse)

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
