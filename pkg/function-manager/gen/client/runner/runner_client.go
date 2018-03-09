///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new runner API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for runner API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetFunctionRuns gets function runs that are being executed
*/
func (a *Client) GetFunctionRuns(params *GetFunctionRunsParams, authInfo runtime.ClientAuthInfoWriter) (*GetFunctionRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFunctionRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFunctionRuns",
		Method:             "GET",
		PathPattern:        "/function/{functionName}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFunctionRunsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFunctionRunsOK), nil

}

/*
GetRun gets function run by its name
*/
func (a *Client) GetRun(params *GetRunParams, authInfo runtime.ClientAuthInfoWriter) (*GetRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRun",
		Method:             "GET",
		PathPattern:        "/function/{functionName}/runs/{runName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRunOK), nil

}

/*
GetRuns gets function runs that are being executed
*/
func (a *Client) GetRuns(params *GetRunsParams, authInfo runtime.ClientAuthInfoWriter) (*GetRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRuns",
		Method:             "GET",
		PathPattern:        "/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRunsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRunsOK), nil

}

/*
RunFunction runs a function
*/
func (a *Client) RunFunction(params *RunFunctionParams, authInfo runtime.ClientAuthInfoWriter) (*RunFunctionOK, *RunFunctionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunFunctionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "runFunction",
		Method:             "POST",
		PathPattern:        "/function/{functionName}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunFunctionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *RunFunctionOK:
		return value, nil, nil
	case *RunFunctionAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
