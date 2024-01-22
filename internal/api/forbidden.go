package api

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/latitudesh/cli/models"
)

func NewForbidden() *Forbidden {
	return &Forbidden{}
}

type Forbidden struct {
	Payload *models.ErrorObject
}

func (o *Forbidden) IsSuccess() bool {
	return false
}

func (o *Forbidden) IsRedirect() bool {
	return false
}

func (o *Forbidden) IsForbidden() bool {
	return true
}

func (o *Forbidden) IsServerError() bool {
	return false
}

func (o *Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *Forbidden) Code() int {
	return 403
}

func (o *Forbidden) Error() string {
	return fmt.Sprintf("[%d] Forbidden  %+v", 403, o.Payload)
}

func (o *Forbidden) String() string {
	return fmt.Sprintf("[%d] Forbidden  %+v", 403, o.Payload)
}

func (o *Forbidden) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *Forbidden) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
