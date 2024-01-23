package apierrors

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/latitudesh/lsh/models"
)

// NewUnprocessableEntity creates a UnprocessableEntity with default headers values
func NewUnprocessableEntity() *UnprocessableEntity {
	return &UnprocessableEntity{}
}

/*
UnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type UnprocessableEntity struct {
	Payload *models.ErrorObject
}

func (o *UnprocessableEntity) IsSuccess() bool {
	return false
}

func (o *UnprocessableEntity) IsRedirect() bool {
	return false
}

func (o *UnprocessableEntity) IsClientError() bool {
	return true
}

func (o *UnprocessableEntity) IsServerError() bool {
	return false
}

func (o *UnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *UnprocessableEntity) Code() int {
	return 422
}

func (o *UnprocessableEntity) Error() string {
	return fmt.Sprintf("[%d] UnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UnprocessableEntity) String() string {
	return fmt.Sprintf("[%d] UnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UnprocessableEntity) GetPayload() *models.ErrorObject {
	return o.Payload
}

func (o *UnprocessableEntity) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
