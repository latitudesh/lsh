package api

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	apierrors "github.com/latitudesh/cli/internal/api/errors"
)

func NewUnauthorized() *Unauthorized {
	return &Unauthorized{}
}

type Unauthorized struct {
	Payload *apierrors.ErrorResponse
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

func (o *Unauthorized) GetPayload() *apierrors.ErrorResponse {
	return o.Payload
}

func (o *Unauthorized) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apierrors.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
