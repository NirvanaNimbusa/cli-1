// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organizations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organizations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddOrganization creates a new organization

Create a new organization
*/
func (a *Client) AddOrganization(params *AddOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*AddOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addOrganization",
		Method:             "POST",
		PathPattern:        "/organizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddOrganizationOK), nil

}

/*
DeleteInvite invites a user to an organization

Invite a user to an organization's roster
*/
func (a *Client) DeleteInvite(params *DeleteInviteParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteInviteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInviteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteInvite",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organizationName}/invitations/{email}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteInviteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteInviteOK), nil

}

/*
DeleteOrganization deletes an organization

Delete an organization
*/
func (a *Client) DeleteOrganization(params *DeleteOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrganization",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organizationName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOrganizationOK), nil

}

/*
EditOrganization edits an organization

Edit an organization
*/
func (a *Client) EditOrganization(params *EditOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*EditOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editOrganization",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EditOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditOrganizationOK), nil

}

/*
GetOrganization retrieves an organization

Return a specific organization matching organization name
*/
func (a *Client) GetOrganization(params *GetOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganization",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrganizationOK), nil

}

/*
GetOrganizationInvitations organizations invitations

Return a list of pending invitations
*/
func (a *Client) GetOrganizationInvitations(params *GetOrganizationInvitationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationInvitationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationInvitationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationInvitations",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationName}/invitations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationInvitationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrganizationInvitationsOK), nil

}

/*
GetOrganizationMembers organizations membership

Return a list of users who are members of the organization
*/
func (a *Client) GetOrganizationMembers(params *GetOrganizationMembersParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationMembers",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationName}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOrganizationMembersOK), nil

}

/*
InviteOrganization invites a user to an organization

Invite a user to an organization's roster
*/
func (a *Client) InviteOrganization(params *InviteOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*InviteOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInviteOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "inviteOrganization",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationName}/invitations/{email}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InviteOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InviteOrganizationOK), nil

}

/*
JoinOrganization joins a user to an organization

Add a user to an organization's roster
*/
func (a *Client) JoinOrganization(params *JoinOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*JoinOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJoinOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "joinOrganization",
		Method:             "PUT",
		PathPattern:        "/organizations/{organizationName}/members/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JoinOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*JoinOrganizationOK), nil

}

/*
ListOrganizations lists of visible organizations

Retrieve all organizations from the system that the user has access to
*/
func (a *Client) ListOrganizations(params *ListOrganizationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrganizationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listOrganizations",
		Method:             "GET",
		PathPattern:        "/organizations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOrganizationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListOrganizationsOK), nil

}

/*
QuitOrganization drops a user from an organization

Remove a user from an organization's roster
*/
func (a *Client) QuitOrganization(params *QuitOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*QuitOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuitOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "quitOrganization",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organizationName}/members/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuitOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuitOrganizationOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}