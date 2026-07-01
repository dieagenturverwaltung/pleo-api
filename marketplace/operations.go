package marketplace

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type CompleteInstallationExec struct {
	config *shared.Config
	ctx    context.Context
}

func (c *Client) CompleteInstallation(ctx context.Context) *CompleteInstallationExec {
	return &CompleteInstallationExec{ctx: ctx, config: c.config}
}

func (e *CompleteInstallationExec) WithContext(ctx context.Context) *CompleteInstallationExec {
	e.ctx = ctx
	return e
}

func (e *CompleteInstallationExec) Execute() (*InstallationResponse, *shared.ResponseExtra, error) {
	var out InstallationResponse
	_, response, err := e.config.SendRequest(e.ctx, "PUT", basePath+"/installations/completions", nil, &out)
	extra := shared.ResponseExtraFromResponse(response)
	if err != nil {
		return nil, extra, err
	}

	return &out, extra, nil
}

type GetClientInstallationExec struct {
	config *shared.Config
	ctx    context.Context

	companyID *string
}

func (c *Client) GetClientInstallation(ctx context.Context) *GetClientInstallationExec {
	return &GetClientInstallationExec{ctx: ctx, config: c.config}
}

func (e *GetClientInstallationExec) WithContext(ctx context.Context) *GetClientInstallationExec {
	e.ctx = ctx
	return e
}

func (e *GetClientInstallationExec) WithCompanyID(companyID string) *GetClientInstallationExec {
	e.companyID = &companyID
	return e
}

func (e *GetClientInstallationExec) Execute() (*InstallationResponse, *shared.ResponseExtra, error) {
	queryParams := make(url.Values)
	shared.AddQueryCompanyID(queryParams, e.companyID)

	var out InstallationResponse
	_, response, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(basePath+"/installations/me", queryParams), nil, &out)
	extra := shared.ResponseExtraFromResponse(response)
	if err != nil {
		return nil, extra, err
	}

	return &out, extra, nil
}

type CreateClientInstallationExec struct {
	config *shared.Config
	ctx    context.Context
	body   CreateClientInstallationRequest
}

func (c *Client) CreateClientInstallation(ctx context.Context) *CreateClientInstallationExec {
	return &CreateClientInstallationExec{ctx: ctx, config: c.config}
}

func (e *CreateClientInstallationExec) WithContext(ctx context.Context) *CreateClientInstallationExec {
	e.ctx = ctx
	return e
}

func (e *CreateClientInstallationExec) WithBody(body CreateClientInstallationRequest) *CreateClientInstallationExec {
	e.body = body
	return e
}

func (e *CreateClientInstallationExec) WithMetadata(metadata map[string]any) *CreateClientInstallationExec {
	e.body.Metadata = metadata
	return e
}

func (e *CreateClientInstallationExec) WithStatus(status InstallationStatus) *CreateClientInstallationExec {
	e.body.Status = status
	return e
}

func (e *CreateClientInstallationExec) Execute() (*InstallationResponse, *shared.ResponseExtra, error) {
	var out InstallationResponse
	_, response, err := e.config.SendRequest(e.ctx, "POST", basePath+"/installations/me", e.body, &out)
	extra := shared.ResponseExtraFromResponse(response)
	if err != nil {
		return nil, extra, err
	}

	return &out, extra, nil
}

type UpdateClientInstallationExec struct {
	config *shared.Config
	ctx    context.Context
	body   UpdateInstallationRequest
}

func (c *Client) UpdateClientInstallation(ctx context.Context) *UpdateClientInstallationExec {
	return &UpdateClientInstallationExec{ctx: ctx, config: c.config}
}

func (e *UpdateClientInstallationExec) WithContext(ctx context.Context) *UpdateClientInstallationExec {
	e.ctx = ctx
	return e
}

func (e *UpdateClientInstallationExec) WithBody(body UpdateInstallationRequest) *UpdateClientInstallationExec {
	e.body = body
	return e
}

func (e *UpdateClientInstallationExec) WithErrorCode(errorCode InstallationErrorCode) *UpdateClientInstallationExec {
	e.body.ErrorCode = &errorCode
	return e
}

func (e *UpdateClientInstallationExec) WithMetadata(metadata map[string]any) *UpdateClientInstallationExec {
	e.body.Metadata = metadata
	return e
}

func (e *UpdateClientInstallationExec) WithStatus(status InstallationStatus) *UpdateClientInstallationExec {
	e.body.Status = status
	return e
}

func (e *UpdateClientInstallationExec) Execute() (*InstallationResponse, *shared.ResponseExtra, error) {
	var out InstallationResponse
	_, response, err := e.config.SendRequest(e.ctx, "PUT", basePath+"/installations/me", e.body, &out)
	extra := shared.ResponseExtraFromResponse(response)
	if err != nil {
		return nil, extra, err
	}

	return &out, extra, nil
}

type DeleteClientInstallationExec struct {
	config *shared.Config
	ctx    context.Context
}

func (c *Client) DeleteClientInstallation(ctx context.Context) *DeleteClientInstallationExec {
	return &DeleteClientInstallationExec{ctx: ctx, config: c.config}
}

func (e *DeleteClientInstallationExec) WithContext(ctx context.Context) *DeleteClientInstallationExec {
	e.ctx = ctx
	return e
}

func (e *DeleteClientInstallationExec) Execute() (*shared.ResponseExtra, error) {
	_, response, err := e.config.SendRequest(e.ctx, "DELETE", basePath+"/installations/me", nil, nil)
	return shared.ResponseExtraFromResponse(response), err
}

type ActivateMyInstallationExec struct {
	config *shared.Config
	ctx    context.Context
}

func (c *Client) ActivateMyInstallation(ctx context.Context) *ActivateMyInstallationExec {
	return &ActivateMyInstallationExec{ctx: ctx, config: c.config}
}

func (e *ActivateMyInstallationExec) WithContext(ctx context.Context) *ActivateMyInstallationExec {
	e.ctx = ctx
	return e
}

func (e *ActivateMyInstallationExec) Execute() (*InstallationResponse, *shared.ResponseExtra, error) {
	var out InstallationResponse
	_, response, err := e.config.SendRequest(e.ctx, "POST", basePath+"/installations/me:activate", nil, &out)
	extra := shared.ResponseExtraFromResponse(response)
	if err != nil {
		return nil, extra, err
	}

	return &out, extra, nil
}

type IntrospectExec struct {
	config *shared.Config
	ctx    context.Context
}

func (c *Client) Introspect(ctx context.Context) *IntrospectExec {
	return &IntrospectExec{ctx: ctx, config: c.config}
}

func (e *IntrospectExec) WithContext(ctx context.Context) *IntrospectExec {
	e.ctx = ctx
	return e
}

func (e *IntrospectExec) Execute() (*ExternalClientModel, *shared.ResponseExtra, error) {
	var out ExternalClientModelDetails
	_, response, err := e.config.SendRequest(e.ctx, "GET", basePath+"/introspect", nil, &out)
	extra := shared.ResponseExtraFromResponse(response)
	if err != nil {
		return nil, extra, err
	}

	return &out.Data, extra, nil
}
