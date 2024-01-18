package internal

import (
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type ErrorSource struct {
	Pointer string `json:"pointer"`
}

type ErrorMeta struct {
	Attribute string `json:"attribute"`
	Message   string `json:"message"`
}

type UnprocessableEntityErrorData struct {
	Code   string       `json:"code"`
	Status string       `json:"status"`
	Title  string       `json:"title"`
	Detail string       `json:"detail"`
	Source *ErrorSource `json:"source"`
	Meta   *ErrorMeta   `json:"meta"`
}

type ErrorResponse struct {
	Errors []UnprocessableEntityErrorData `json:"errors,omitempty"`
}

func NewUnprocessableEntityError() *UnprocessableEntityError {
	return &UnprocessableEntityError{}
}

type UnprocessableEntityError struct {
	Payload *ErrorResponse
}

func (o *UnprocessableEntityError) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
