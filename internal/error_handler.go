package internal

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/latitudesh/cli/models"
)

func NewNotFoundError() *NotFoundError {
	return &NotFoundError{}
}

type NotFoundError struct {
	Payload *models.ErrorObject
}

func (o *NotFoundError) IsSuccess() bool {
	return false
}

func (o *NotFoundError) IsRedirect() bool {
	return false
}

func (o *NotFoundError) IsClientError() bool {
	return true
}

func (o *NotFoundError) IsServerError() bool {
	return false
}

func (o *NotFoundError) IsCode(code int) bool {
	return code == 404
}

func (o *NotFoundError) Code() int {
	return 404
}

func (o *NotFoundError) Error() string {
	return fmt.Sprintf("[%d] NotFoundError  %+v", 404, o.Payload)
}

func (o *NotFoundError) String() string {
	return fmt.Sprintf("[%d] NotFoundError  %+v", 404, o.Payload)
}

func (o *NotFoundError) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *NotFoundError) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
