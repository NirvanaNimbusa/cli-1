// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new components API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for components API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetComponent gets a single component record

Retrieve the component by ID
*/
func (a *Client) GetComponent(params *GetComponentParams, authInfo runtime.ClientAuthInfoWriter) (*GetComponentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComponentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getComponent",
		Method:             "GET",
		PathPattern:        "/components/{componentID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComponentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetComponentOK), nil

}

/*
GetComponentsByIdentity alls components seen for specifid identity

Retrieve all components for this identity
*/
func (a *Client) GetComponentsByIdentity(params *GetComponentsByIdentityParams, authInfo runtime.ClientAuthInfoWriter) (*GetComponentsByIdentityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComponentsByIdentityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getComponentsByIdentity",
		Method:             "GET",
		PathPattern:        "/identities/{identityID}/components",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComponentsByIdentityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetComponentsByIdentityOK), nil

}

/*
ListComponents lists of all matching components

Retrieve all components from the system that the user has access to
*/
func (a *Client) ListComponents(params *ListComponentsParams, authInfo runtime.ClientAuthInfoWriter) (*ListComponentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListComponentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listComponents",
		Method:             "GET",
		PathPattern:        "/components",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListComponentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListComponentsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
