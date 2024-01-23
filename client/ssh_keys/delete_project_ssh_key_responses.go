// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/latitudesh/cli/internal/api"
)

// DeleteProjectSSHKeyReader is a Reader for the DeleteProjectSSHKey structure.
type DeleteProjectSSHKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectSSHKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectSSHKeyOK()
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
	case 404:
		result := api.NewNotFound()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}] delete-project-ssh-key", response, response.Code())
	}
}

// NewDeleteProjectSSHKeyOK creates a DeleteProjectSSHKeyOK with default headers values
func NewDeleteProjectSSHKeyOK() *DeleteProjectSSHKeyOK {
	return &DeleteProjectSSHKeyOK{}
}

/*
DeleteProjectSSHKeyOK describes a response with status code 200, with default header values.

Success
*/
type DeleteProjectSSHKeyOK struct {
}

// IsSuccess returns true when this delete project Ssh key o k response has a 2xx status code
func (o *DeleteProjectSSHKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete project Ssh key o k response has a 3xx status code
func (o *DeleteProjectSSHKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete project Ssh key o k response has a 4xx status code
func (o *DeleteProjectSSHKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete project Ssh key o k response has a 5xx status code
func (o *DeleteProjectSSHKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete project Ssh key o k response a status code equal to that given
func (o *DeleteProjectSSHKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete project Ssh key o k response
func (o *DeleteProjectSSHKeyOK) Code() int {
	return 200
}

func (o *DeleteProjectSSHKeyOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}][%d] deleteProjectSshKeyOK ", 200)
}

func (o *DeleteProjectSSHKeyOK) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_id_or_slug}/ssh_keys/{ssh_key_id}][%d] deleteProjectSshKeyOK ", 200)
}

func (o *DeleteProjectSSHKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
