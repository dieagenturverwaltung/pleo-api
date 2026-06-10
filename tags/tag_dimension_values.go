package tags

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type GetTagDimensionValuesExec struct {
	config *shared.Config
	ctx    context.Context
	tagID  string
}

func (c *TagsClient) GetTagDimensionValues(ctx context.Context, tagID string) *GetTagDimensionValuesExec {
	return &GetTagDimensionValuesExec{ctx: ctx, config: c.config, tagID: tagID}
}

func (e *GetTagDimensionValuesExec) WithContext(ctx context.Context) *GetTagDimensionValuesExec {
	e.ctx = ctx
	return e
}

func (e *GetTagDimensionValuesExec) Execute() (*shared.ListResponse[TagDimensionValueModel], error) {
	var out shared.ListResponse[TagDimensionValueModel]
	_, _, err := e.config.SendRequest(e.ctx, "GET", basePath+"/tags/"+url.PathEscape(e.tagID)+"/dimensions", nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type GetDimensionValueExec struct {
	config      *shared.Config
	ctx         context.Context
	tagID       string
	dimensionID string
}

func (c *TagsClient) GetDimensionValue(ctx context.Context, tagID, dimensionID string) *GetDimensionValueExec {
	return &GetDimensionValueExec{ctx: ctx, config: c.config, tagID: tagID, dimensionID: dimensionID}
}

func (e *GetDimensionValueExec) WithContext(ctx context.Context) *GetDimensionValueExec {
	e.ctx = ctx
	return e
}

func (e *GetDimensionValueExec) Execute() (*TagDimensionValueModel, error) {
	var out shared.Response[TagDimensionValueModel]
	path := basePath + "/tags/" + url.PathEscape(e.tagID) + "/dimensions/" + url.PathEscape(e.dimensionID)
	_, _, err := e.config.SendRequest(e.ctx, "GET", path, nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type CreateDimensionValueExec struct {
	config      *shared.Config
	ctx         context.Context
	tagID       string
	dimensionID string
	body        TagDimensionValueCreateModel
}

func (c *TagsClient) CreateDimensionValue(ctx context.Context, tagID, dimensionID string) *CreateDimensionValueExec {
	return &CreateDimensionValueExec{ctx: ctx, config: c.config, tagID: tagID, dimensionID: dimensionID}
}

func (e *CreateDimensionValueExec) WithContext(ctx context.Context) *CreateDimensionValueExec {
	e.ctx = ctx
	return e
}

func (e *CreateDimensionValueExec) WithBody(body TagDimensionValueCreateModel) *CreateDimensionValueExec {
	e.body = body
	return e
}

func (e *CreateDimensionValueExec) WithValue(value string) *CreateDimensionValueExec {
	e.body.Value = value
	return e
}

func (e *CreateDimensionValueExec) Execute() (*TagDimensionValueModel, error) {
	var out shared.Response[TagDimensionValueModel]
	path := basePath + "/tags/" + url.PathEscape(e.tagID) + "/dimensions/" + url.PathEscape(e.dimensionID)
	_, _, err := e.config.SendRequest(e.ctx, "POST", path, e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type UpdateDimensionValueExec struct {
	config      *shared.Config
	ctx         context.Context
	tagID       string
	dimensionID string
	body        TagDimensionValueUpdateModel
}

func (c *TagsClient) UpdateDimensionValue(ctx context.Context, tagID, dimensionID string) *UpdateDimensionValueExec {
	return &UpdateDimensionValueExec{ctx: ctx, config: c.config, tagID: tagID, dimensionID: dimensionID}
}

func (e *UpdateDimensionValueExec) WithContext(ctx context.Context) *UpdateDimensionValueExec {
	e.ctx = ctx
	return e
}

func (e *UpdateDimensionValueExec) WithBody(body TagDimensionValueUpdateModel) *UpdateDimensionValueExec {
	e.body = body
	return e
}

func (e *UpdateDimensionValueExec) WithValue(value string) *UpdateDimensionValueExec {
	e.body.Value = value
	return e
}

func (e *UpdateDimensionValueExec) Execute() (*TagDimensionValueModel, error) {
	var out shared.Response[TagDimensionValueModel]
	path := basePath + "/tags/" + url.PathEscape(e.tagID) + "/dimensions/" + url.PathEscape(e.dimensionID)
	_, _, err := e.config.SendRequest(e.ctx, "PUT", path, e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type DeleteDimensionValueExec struct {
	config      *shared.Config
	ctx         context.Context
	tagID       string
	dimensionID string
}

func (c *TagsClient) DeleteDimensionValue(ctx context.Context, tagID, dimensionID string) *DeleteDimensionValueExec {
	return &DeleteDimensionValueExec{ctx: ctx, config: c.config, tagID: tagID, dimensionID: dimensionID}
}

func (e *DeleteDimensionValueExec) WithContext(ctx context.Context) *DeleteDimensionValueExec {
	e.ctx = ctx
	return e
}

func (e *DeleteDimensionValueExec) Execute() (*TagDimensionValueModel, error) {
	var out shared.Response[TagDimensionValueModel]
	path := basePath + "/tags/" + url.PathEscape(e.tagID) + "/dimensions/" + url.PathEscape(e.dimensionID)
	_, _, err := e.config.SendRequest(e.ctx, "DELETE", path, nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}
