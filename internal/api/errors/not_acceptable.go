package apierrors

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

func NewNotAcceptable() *NotAcceptable {
	return &NotAcceptable{}
}

type NotAcceptable struct {
	Payload *ErrorPayload
}

func (o *NotAcceptable) IsSuccess() bool {
	return false
}

func (o *NotAcceptable) IsRedirect() bool {
	return false
}

func (o *NotAcceptable) IsNotAcceptable() bool {
	return true
}

func (o *NotAcceptable) IsServerError() bool {
	return false
}

func (o *NotAcceptable) IsCode(code int) bool {
	return code == 406
}

func (o *NotAcceptable) Code() int {
	return 406
}

func (o *NotAcceptable) Error() string {
	return fmt.Sprintf("[%d] NotAcceptable  %+v", 406, o.Payload)
}

func (o *NotAcceptable) String() string {
	return fmt.Sprintf("[%d] NotAcceptable  %+v", 406, o.Payload)
}

func (o *NotAcceptable) GetPayload() *ErrorPayload {
	return o.Payload
}

func (o *NotAcceptable) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
