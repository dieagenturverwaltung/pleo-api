package tags

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type GetTagGroupDimensionsExec struct {
	config  *shared.Config
	ctx     context.Context
	groupID string
}

func (c *TagGroupsClient) GetTagGroupDimensions(ctx context.Context, groupID string) *GetTagGroupDimensionsExec {
	return &GetTagGroupDimensionsExec{ctx: ctx, config: c.config, groupID: groupID}
}

func (e *GetTagGroupDimensionsExec) WithContext(ctx context.Context) *GetTagGroupDimensionsExec {
	e.ctx = ctx
	return e
}

func (e *GetTagGroupDimensionsExec) Execute() (*shared.ListResponse[TagGroupDimensionModel], error) {
	var out shared.ListResponse[TagGroupDimensionModel]
	path := basePath + "/tag-groups/" + url.PathEscape(e.groupID) + "/dimensions"
	_, _, err := e.config.SendRequest(e.ctx, "GET", path, nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type CreateTagGroupDimensionExec struct {
	config  *shared.Config
	ctx     context.Context
	groupID string
	body    TagGroupDimensionCreateModel
}

func (c *TagGroupsClient) CreateTagGroupDimension(ctx context.Context, groupID string) *CreateTagGroupDimensionExec {
	return &CreateTagGroupDimensionExec{ctx: ctx, config: c.config, groupID: groupID}
}

func (e *CreateTagGroupDimensionExec) WithContext(ctx context.Context) *CreateTagGroupDimensionExec {
	e.ctx = ctx
	return e
}

func (e *CreateTagGroupDimensionExec) WithBody(body TagGroupDimensionCreateModel) *CreateTagGroupDimensionExec {
	e.body = body
	return e
}

func (e *CreateTagGroupDimensionExec) WithCode(code string) *CreateTagGroupDimensionExec {
	e.body.Code = code
	return e
}

func (e *CreateTagGroupDimensionExec) WithDisplayOrder(displayOrder int) *CreateTagGroupDimensionExec {
	e.body.DisplayOrder = displayOrder
	return e
}

func (e *CreateTagGroupDimensionExec) WithName(name string) *CreateTagGroupDimensionExec {
	e.body.Name = name
	return e
}

func (e *CreateTagGroupDimensionExec) WithVisible(visible bool) *CreateTagGroupDimensionExec {
	e.body.Visible = visible
	return e
}

func (e *CreateTagGroupDimensionExec) Execute() (*TagGroupDimensionModel, error) {
	var out shared.Response[TagGroupDimensionModel]
	path := basePath + "/tag-groups/" + url.PathEscape(e.groupID) + "/dimensions"
	_, _, err := e.config.SendRequest(e.ctx, "POST", path, e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type GetTagGroupDimensionExec struct {
	config      *shared.Config
	ctx         context.Context
	groupID     string
	dimensionID string
}

func (c *TagGroupsClient) GetTagGroupDimension(ctx context.Context, groupID, dimensionID string) *GetTagGroupDimensionExec {
	return &GetTagGroupDimensionExec{ctx: ctx, config: c.config, groupID: groupID, dimensionID: dimensionID}
}

func (e *GetTagGroupDimensionExec) WithContext(ctx context.Context) *GetTagGroupDimensionExec {
	e.ctx = ctx
	return e
}

func (e *GetTagGroupDimensionExec) Execute() (*TagGroupDimensionModel, error) {
	var out shared.Response[TagGroupDimensionModel]
	path := basePath + "/tag-groups/" + url.PathEscape(e.groupID) + "/dimensions/" + url.PathEscape(e.dimensionID)
	_, _, err := e.config.SendRequest(e.ctx, "GET", path, nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type UpdateTagGroupDimensionExec struct {
	config      *shared.Config
	ctx         context.Context
	groupID     string
	dimensionID string
	body        TagGroupDimensionUpdateModel
}

func (c *TagGroupsClient) UpdateTagGroupDimension(ctx context.Context, groupID, dimensionID string) *UpdateTagGroupDimensionExec {
	return &UpdateTagGroupDimensionExec{ctx: ctx, config: c.config, groupID: groupID, dimensionID: dimensionID}
}

func (e *UpdateTagGroupDimensionExec) WithContext(ctx context.Context) *UpdateTagGroupDimensionExec {
	e.ctx = ctx
	return e
}

func (e *UpdateTagGroupDimensionExec) WithBody(body TagGroupDimensionUpdateModel) *UpdateTagGroupDimensionExec {
	e.body = body
	return e
}

func (e *UpdateTagGroupDimensionExec) WithCode(code string) *UpdateTagGroupDimensionExec {
	e.body.Code = code
	return e
}

func (e *UpdateTagGroupDimensionExec) WithDisplayOrder(displayOrder int) *UpdateTagGroupDimensionExec {
	e.body.DisplayOrder = displayOrder
	return e
}

func (e *UpdateTagGroupDimensionExec) WithName(name string) *UpdateTagGroupDimensionExec {
	e.body.Name = name
	return e
}

func (e *UpdateTagGroupDimensionExec) WithVisible(visible bool) *UpdateTagGroupDimensionExec {
	e.body.Visible = visible
	return e
}

func (e *UpdateTagGroupDimensionExec) Execute() (*TagGroupDimensionModel, error) {
	var out shared.Response[TagGroupDimensionModel]
	path := basePath + "/tag-groups/" + url.PathEscape(e.groupID) + "/dimensions/" + url.PathEscape(e.dimensionID)
	_, _, err := e.config.SendRequest(e.ctx, "PUT", path, e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type DeleteTagGroupDimensionExec struct {
	config      *shared.Config
	ctx         context.Context
	groupID     string
	dimensionID string
}

func (c *TagGroupsClient) DeleteTagGroupDimension(ctx context.Context, groupID, dimensionID string) *DeleteTagGroupDimensionExec {
	return &DeleteTagGroupDimensionExec{ctx: ctx, config: c.config, groupID: groupID, dimensionID: dimensionID}
}

func (e *DeleteTagGroupDimensionExec) WithContext(ctx context.Context) *DeleteTagGroupDimensionExec {
	e.ctx = ctx
	return e
}

func (e *DeleteTagGroupDimensionExec) Execute() error {
	path := basePath + "/tag-groups/" + url.PathEscape(e.groupID) + "/dimensions/" + url.PathEscape(e.dimensionID)
	_, _, err := e.config.SendRequest(e.ctx, "DELETE", path, nil, nil)
	return err
}
