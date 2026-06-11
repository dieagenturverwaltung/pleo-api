package tags

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type SearchAggregatedTagsExec struct {
	config *shared.Config
	ctx    context.Context

	companyID  *string
	pagingInfo shared.PagingInfo
	body       AggregatedTagSearchRequest
}

func (c *TagsClient) SearchAggregatedTags(ctx context.Context) *SearchAggregatedTagsExec {
	return &SearchAggregatedTagsExec{ctx: ctx, config: c.config}
}

func (e *SearchAggregatedTagsExec) WithContext(ctx context.Context) *SearchAggregatedTagsExec {
	e.ctx = ctx
	return e
}

func (e *SearchAggregatedTagsExec) WithCompanyID(companyID string) *SearchAggregatedTagsExec {
	e.companyID = &companyID
	return e
}

func (e *SearchAggregatedTagsExec) WithPagingInfo(pagingInfo shared.PagingInfo) *SearchAggregatedTagsExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *SearchAggregatedTagsExec) WithBody(body AggregatedTagSearchRequest) *SearchAggregatedTagsExec {
	e.body = body
	return e
}

func (e *SearchAggregatedTagsExec) WithTagIDs(tagIDs ...string) *SearchAggregatedTagsExec {
	e.body.TagIDs = append(e.body.TagIDs, tagIDs...)
	return e
}

func (e *SearchAggregatedTagsExec) WithTextSearch(textSearch string) *SearchAggregatedTagsExec {
	e.body.TextSearch = textSearch
	return e
}

func (e *SearchAggregatedTagsExec) Execute() (*shared.CursorPageResponse[AggregatedTagModel], error) {
	queryParams := make(url.Values)
	shared.AddQueryCompanyID(queryParams, e.companyID)
	e.pagingInfo.Apply(queryParams)

	var out shared.CursorPageResponse[AggregatedTagModel]
	_, _, err := e.config.SendRequest(e.ctx, "POST", shared.URLWithQuery(basePath+"/aggregations/tags", queryParams), e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type SearchTagsExec struct {
	config *shared.Config
	ctx    context.Context

	companyID  *string
	textSearch *string
	pagingInfo shared.PagingInfo
}

func (c *TagsClient) SearchTags(ctx context.Context) *SearchTagsExec {
	return &SearchTagsExec{ctx: ctx, config: c.config}
}

func (e *SearchTagsExec) WithContext(ctx context.Context) *SearchTagsExec {
	e.ctx = ctx
	return e
}

func (e *SearchTagsExec) WithCompanyID(companyID string) *SearchTagsExec {
	e.companyID = &companyID
	return e
}

func (e *SearchTagsExec) WithTextSearch(textSearch string) *SearchTagsExec {
	e.textSearch = &textSearch
	return e
}

func (e *SearchTagsExec) WithPagingInfo(pagingInfo shared.PagingInfo) *SearchTagsExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *SearchTagsExec) Execute() (*shared.CursorPageResponse[TagModel], error) {
	queryParams := make(url.Values)
	shared.AddQueryCompanyID(queryParams, e.companyID)
	shared.AddQueryString(queryParams, "text_search", e.textSearch)
	e.pagingInfo.Apply(queryParams)

	var out shared.CursorPageResponse[TagModel]
	_, _, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(basePath+"/tags", queryParams), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type GetTagExec struct {
	config *shared.Config
	ctx    context.Context
	tagID  string
}

func (c *TagsClient) GetTag(ctx context.Context, tagID string) *GetTagExec {
	return &GetTagExec{ctx: ctx, config: c.config, tagID: tagID}
}

func (e *GetTagExec) WithContext(ctx context.Context) *GetTagExec {
	e.ctx = ctx
	return e
}

func (e *GetTagExec) Execute() (*TagModel, error) {
	var out shared.Response[TagModel]
	_, _, err := e.config.SendRequest(e.ctx, "GET", basePath+"/tags/"+url.PathEscape(e.tagID), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type UpdateTagExec struct {
	config *shared.Config
	ctx    context.Context
	tagID  string
	body   TagUpdateModel
}

func (c *TagsClient) UpdateTag(ctx context.Context, tagID string) *UpdateTagExec {
	return &UpdateTagExec{ctx: ctx, config: c.config, tagID: tagID}
}

func (e *UpdateTagExec) WithContext(ctx context.Context) *UpdateTagExec {
	e.ctx = ctx
	return e
}

func (e *UpdateTagExec) WithBody(body TagUpdateModel) *UpdateTagExec {
	e.body = body
	return e
}

func (e *UpdateTagExec) WithArchived(archived bool) *UpdateTagExec {
	e.body.Archived = archived
	return e
}

func (e *UpdateTagExec) WithCode(code string) *UpdateTagExec {
	e.body.Code = code
	return e
}

func (e *UpdateTagExec) WithName(name string) *UpdateTagExec {
	e.body.Name = name
	return e
}

func (e *UpdateTagExec) Execute() (*TagModel, error) {
	var out shared.Response[TagModel]
	_, _, err := e.config.SendRequest(e.ctx, "PUT", basePath+"/tags/"+url.PathEscape(e.tagID), e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type DeleteTagExec struct {
	config *shared.Config
	ctx    context.Context
	tagID  string
}

func (c *TagsClient) DeleteTag(ctx context.Context, tagID string) *DeleteTagExec {
	return &DeleteTagExec{ctx: ctx, config: c.config, tagID: tagID}
}

func (e *DeleteTagExec) WithContext(ctx context.Context) *DeleteTagExec {
	e.ctx = ctx
	return e
}

func (e *DeleteTagExec) Execute() error {
	_, _, err := e.config.SendRequest(e.ctx, "DELETE", basePath+"/tags/"+url.PathEscape(e.tagID), nil, nil)
	return err
}
