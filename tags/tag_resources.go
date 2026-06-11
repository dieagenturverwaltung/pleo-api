package tags

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type GetTagsForGroupExec struct {
	config *shared.Config
	ctx    context.Context

	groupID         string
	includeArchived *bool
	pagingInfo      shared.PagingInfo
}

func (c *TagGroupsClient) GetTags(ctx context.Context, groupID string) *GetTagsForGroupExec {
	return &GetTagsForGroupExec{ctx: ctx, config: c.config, groupID: groupID}
}

func (e *GetTagsForGroupExec) WithContext(ctx context.Context) *GetTagsForGroupExec {
	e.ctx = ctx
	return e
}

func (e *GetTagsForGroupExec) WithIncludeArchived(includeArchived bool) *GetTagsForGroupExec {
	e.includeArchived = &includeArchived
	return e
}

func (e *GetTagsForGroupExec) WithPagingInfo(pagingInfo shared.PagingInfo) *GetTagsForGroupExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *GetTagsForGroupExec) Execute() (*shared.CursorPageResponse[TagModel], error) {
	queryParams := make(url.Values)
	shared.AddQueryBool(queryParams, "include_archived", e.includeArchived)
	e.pagingInfo.Apply(queryParams)

	path := basePath + "/tag-groups/" + url.PathEscape(e.groupID) + "/tags"
	var out shared.CursorPageResponse[TagModel]
	_, _, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(path, queryParams), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type CreateTagExec struct {
	config  *shared.Config
	ctx     context.Context
	groupID string
	body    TagCreateModel
}

func (c *TagGroupsClient) CreateTag(ctx context.Context, groupID string) *CreateTagExec {
	return &CreateTagExec{ctx: ctx, config: c.config, groupID: groupID}
}

func (e *CreateTagExec) WithContext(ctx context.Context) *CreateTagExec {
	e.ctx = ctx
	return e
}

func (e *CreateTagExec) WithBody(body TagCreateModel) *CreateTagExec {
	e.body = body
	return e
}

func (e *CreateTagExec) WithArchived(archived bool) *CreateTagExec {
	e.body.Archived = archived
	return e
}

func (e *CreateTagExec) WithCode(code string) *CreateTagExec {
	e.body.Code = code
	return e
}

func (e *CreateTagExec) WithName(name string) *CreateTagExec {
	e.body.Name = name
	return e
}

func (e *CreateTagExec) Execute() (*TagModel, error) {
	var out shared.Response[TagModel]
	path := basePath + "/tag-groups/" + url.PathEscape(e.groupID) + "/tags"
	_, _, err := e.config.SendRequest(e.ctx, "POST", path, e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}
