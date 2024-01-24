package apierrors

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

func NewBadRequest() *BadRequest {
	return &BadRequest{}
}

type BadRequest struct {
	Payload *ErrorResponse
}

func (o *BadRequest) IsSuccess() bool {
	return false
}

func (o *BadRequest) IsRedirect() bool {
	return false
}

func (o *BadRequest) IsBadRequest() bool {
	return true
}

func (o *BadRequest) IsServerError() bool {
	return false
}

func (o *BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BadRequest) Code() int {
	return 400
}

func (o *BadRequest) Error() string {
	return fmt.Sprintf("[%d] BadRequest  %+v", 400, o.Payload)
}

func (o *BadRequest) String() string {
	return fmt.Sprintf("[%d] BadRequest  %+v", 400, o.Payload)
}

func (o *BadRequest) GetPayload() *ErrorResponse {
	return o.Payload
}

func (o *BadRequest) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
