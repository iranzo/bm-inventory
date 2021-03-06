// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the events client
type API interface {
	/*
	   ListEvents lists events for an entity id*/
	ListEvents(ctx context.Context, params *ListEventsParams) (*ListEventsOK, error)
}

// New creates a new events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
ListEvents lists events for an entity id
*/
func (a *Client) ListEvents(ctx context.Context, params *ListEventsParams) (*ListEventsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListEvents",
		Method:             "GET",
		PathPattern:        "/events/{entity_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListEventsReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListEventsOK), nil

}
