// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/latitudesh/cli/internal/api"
	"github.com/latitudesh/cli/models"
)

// GetTeamReader is a Reader for the GetTeam structure.
type GetTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := api.NewUnauthorized()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /team] get-team", response, response.Code())
	}
}

// NewGetTeamOK creates a GetTeamOK with default headers values
func NewGetTeamOK() *GetTeamOK {
	return &GetTeamOK{}
}

/*
GetTeamOK describes a response with status code 200, with default header values.

Success
*/
type GetTeamOK struct {
	Payload *models.Teams
}

// IsSuccess returns true when this get team o k response has a 2xx status code
func (o *GetTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team o k response has a 3xx status code
func (o *GetTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team o k response has a 4xx status code
func (o *GetTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team o k response has a 5xx status code
func (o *GetTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team o k response a status code equal to that given
func (o *GetTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team o k response
func (o *GetTeamOK) Code() int {
	return 200
}

func (o *GetTeamOK) Error() string {
	return fmt.Sprintf("[GET /team][%d] getTeamOK  %+v", 200, o.Payload)
}

func (o *GetTeamOK) String() string {
	return fmt.Sprintf("[GET /team][%d] getTeamOK  %+v", 200, o.Payload)
}

func (o *GetTeamOK) GetPayload() *models.Teams {
	return o.Payload
}

func (o *GetTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Teams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
