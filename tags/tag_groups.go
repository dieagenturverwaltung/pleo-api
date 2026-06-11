package tags

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type GetAggregatedTagGroupsExec struct {
	config *shared.Config
	ctx    context.Context

	companyID      *string
	organizationID *string
	tagGroupIDs    []string
}

func (c *TagGroupsClient) GetAggregatedTagGroups(ctx context.Context) *GetAggregatedTagGroupsExec {
	return &GetAggregatedTagGroupsExec{ctx: ctx, config: c.config}
}

func (e *GetAggregatedTagGroupsExec) WithContext(ctx context.Context) *GetAggregatedTagGroupsExec {
	e.ctx = ctx
	return e
}

func (e *GetAggregatedTagGroupsExec) WithCompanyID(companyID string) *GetAggregatedTagGroupsExec {
	e.companyID = &companyID
	return e
}

func (e *GetAggregatedTagGroupsExec) WithOrganizationID(organizationID string) *GetAggregatedTagGroupsExec {
	e.organizationID = &organizationID
	return e
}

func (e *GetAggregatedTagGroupsExec) WithTagGroupIDs(tagGroupIDs ...string) *GetAggregatedTagGroupsExec {
	e.tagGroupIDs = append(e.tagGroupIDs, tagGroupIDs...)
	return e
}

func (e *GetAggregatedTagGroupsExec) Execute() (*shared.ListResponse[AggregatedTagGroupModel], error) {
	queryParams := make(url.Values)
	if e.companyID != nil || e.organizationID == nil {
		shared.AddQueryCompanyID(queryParams, e.companyID)
	}
	shared.AddQueryString(queryParams, "organization_id", e.organizationID)
	shared.AddQueryStrings(queryParams, "tag_group_ids", e.tagGroupIDs)

	var out shared.ListResponse[AggregatedTagGroupModel]
	_, _, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(basePath+"/aggregations/tag-groups", queryParams), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type GetTagGroupsExec struct {
	config *shared.Config
	ctx    context.Context

	companyID      *string
	organizationID *string
}

func (c *TagGroupsClient) GetTagGroups(ctx context.Context) *GetTagGroupsExec {
	return &GetTagGroupsExec{ctx: ctx, config: c.config}
}

func (e *GetTagGroupsExec) WithContext(ctx context.Context) *GetTagGroupsExec {
	e.ctx = ctx
	return e
}

func (e *GetTagGroupsExec) WithCompanyID(companyID string) *GetTagGroupsExec {
	e.companyID = &companyID
	return e
}

func (e *GetTagGroupsExec) WithOrganizationID(organizationID string) *GetTagGroupsExec {
	e.organizationID = &organizationID
	return e
}

func (e *GetTagGroupsExec) Execute() (*shared.ListResponse[TagGroupModel], error) {
	queryParams := make(url.Values)
	if e.companyID != nil || e.organizationID == nil {
		shared.AddQueryCompanyID(queryParams, e.companyID)
	}
	shared.AddQueryString(queryParams, "organization_id", e.organizationID)

	var out shared.ListResponse[TagGroupModel]
	_, _, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(basePath+"/tag-groups", queryParams), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type CreateTagGroupExec struct {
	config *shared.Config
	ctx    context.Context

	companyID *string
	body      TagGroupCreateModel
}

func (c *TagGroupsClient) CreateTagGroup(ctx context.Context) *CreateTagGroupExec {
	return &CreateTagGroupExec{ctx: ctx, config: c.config}
}

func (e *CreateTagGroupExec) WithContext(ctx context.Context) *CreateTagGroupExec {
	e.ctx = ctx
	return e
}

func (e *CreateTagGroupExec) WithCompanyID(companyID string) *CreateTagGroupExec {
	e.companyID = &companyID
	return e
}

func (e *CreateTagGroupExec) WithBody(body TagGroupCreateModel) *CreateTagGroupExec {
	e.body = body
	return e
}

func (e *CreateTagGroupExec) WithArchived(archived bool) *CreateTagGroupExec {
	e.body.Archived = archived
	return e
}

func (e *CreateTagGroupExec) WithCode(code string) *CreateTagGroupExec {
	e.body.Code = code
	return e
}

func (e *CreateTagGroupExec) WithMetadata(metadata map[string]string) *CreateTagGroupExec {
	e.body.Metadata = metadata
	return e
}

func (e *CreateTagGroupExec) WithName(name string) *CreateTagGroupExec {
	e.body.Name = name
	return e
}

func (e *CreateTagGroupExec) Execute() (*TagGroupModel, error) {
	queryParams := make(url.Values)
	shared.AddQueryCompanyID(queryParams, e.companyID)

	var out shared.Response[TagGroupModel]
	_, _, err := e.config.SendRequest(e.ctx, "POST", shared.URLWithQuery(basePath+"/tag-groups", queryParams), e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type GetTagGroupExec struct {
	config  *shared.Config
	ctx     context.Context
	groupID string
}

func (c *TagGroupsClient) GetTagGroup(ctx context.Context, groupID string) *GetTagGroupExec {
	return &GetTagGroupExec{ctx: ctx, config: c.config, groupID: groupID}
}

func (e *GetTagGroupExec) WithContext(ctx context.Context) *GetTagGroupExec {
	e.ctx = ctx
	return e
}

func (e *GetTagGroupExec) Execute() (*TagGroupModel, error) {
	var out shared.Response[TagGroupModel]
	_, _, err := e.config.SendRequest(e.ctx, "GET", basePath+"/tag-groups/"+url.PathEscape(e.groupID), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type UpdateTagGroupExec struct {
	config  *shared.Config
	ctx     context.Context
	groupID string
	body    TagGroupUpdateModel
}

func (c *TagGroupsClient) UpdateTagGroup(ctx context.Context, groupID string) *UpdateTagGroupExec {
	return &UpdateTagGroupExec{ctx: ctx, config: c.config, groupID: groupID}
}

func (e *UpdateTagGroupExec) WithContext(ctx context.Context) *UpdateTagGroupExec {
	e.ctx = ctx
	return e
}

func (e *UpdateTagGroupExec) WithBody(body TagGroupUpdateModel) *UpdateTagGroupExec {
	e.body = body
	return e
}

func (e *UpdateTagGroupExec) WithArchived(archived bool) *UpdateTagGroupExec {
	e.body.Archived = archived
	return e
}

func (e *UpdateTagGroupExec) WithCode(code string) *UpdateTagGroupExec {
	e.body.Code = code
	return e
}

func (e *UpdateTagGroupExec) WithMetadata(metadata map[string]string) *UpdateTagGroupExec {
	e.body.Metadata = metadata
	return e
}

func (e *UpdateTagGroupExec) WithName(name string) *UpdateTagGroupExec {
	e.body.Name = name
	return e
}

func (e *UpdateTagGroupExec) Execute() (*TagGroupModel, error) {
	var out shared.Response[TagGroupModel]
	_, _, err := e.config.SendRequest(e.ctx, "PUT", basePath+"/tag-groups/"+url.PathEscape(e.groupID), e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type DeleteTagGroupExec struct {
	config  *shared.Config
	ctx     context.Context
	groupID string
}

func (c *TagGroupsClient) DeleteTagGroup(ctx context.Context, groupID string) *DeleteTagGroupExec {
	return &DeleteTagGroupExec{ctx: ctx, config: c.config, groupID: groupID}
}

func (e *DeleteTagGroupExec) WithContext(ctx context.Context) *DeleteTagGroupExec {
	e.ctx = ctx
	return e
}

func (e *DeleteTagGroupExec) Execute() (*bool, error) {
	var out bool
	_, _, err := e.config.SendRequest(e.ctx, "DELETE", basePath+"/tag-groups/"+url.PathEscape(e.groupID), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
