package api

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/latitudesh/cli/models"
)

func NewNotFound() *NotFound {
	return &NotFound{}
}

type NotFound struct {
	Payload *models.ErrorObject
}

func (o *NotFound) IsSuccess() bool {
	return false
}

func (o *NotFound) IsRedirect() bool {
	return false
}

func (o *NotFound) IsNotFound() bool {
	return true
}

func (o *NotFound) IsServerError() bool {
	return false
}

func (o *NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *NotFound) Code() int {
	return 404
}

func (o *NotFound) Error() string {
	return fmt.Sprintf("[%d] NotFound  %+v", 404, o.Payload)
}

func (o *NotFound) String() string {
	return fmt.Sprintf("[%d] NotFound  %+v", 404, o.Payload)
}

func (o *NotFound) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *NotFound) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
