package apierrors

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

func NewUnauthorized() *Unauthorized {
	return &Unauthorized{}
}

type Unauthorized struct {
	Payload *ErrorPayload
}

func (o *Unauthorized) IsSuccess() bool {
	return false
}

func (o *Unauthorized) IsRedirect() bool {
	return false
}

func (o *Unauthorized) IsUnauthorized() bool {
	return true
}

func (o *Unauthorized) IsServerError() bool {
	return false
}

func (o *Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *Unauthorized) Code() int {
	return 401
}

func (o *Unauthorized) Error() string {
	return fmt.Sprintf("[%d] Unauthorized  %+v", 401, o.Payload)
}

func (o *Unauthorized) String() string {
	return fmt.Sprintf("[%d] Unauthorized  %+v", 401, o.Payload)
}

func (o *Unauthorized) GetPayload() *ErrorPayload {
	return o.Payload
}

func (o *Unauthorized) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
