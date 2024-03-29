// Code generated by go-swagger; DO NOT EDIT.

package server_reinstall

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

var serverReinstallType = "reinstalls"

// New creates a new server reinstall API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server reinstall API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateServerReinstall(params *CreateServerReinstallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateServerReinstallCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateServerReinstall runs server reinstall
*/
func (a *Client) CreateServerReinstall(params *CreateServerReinstallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateServerReinstallCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServerReinstallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create-server-reinstall",
		Method:             "POST",
		PathPattern:        "/servers/{server_id}/reinstall",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateServerReinstallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateServerReinstallCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-server-reinstall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
