package tags

import "github.com/dieagenturverwaltung/pleo-api/shared"

const basePath = "/v0"

type Client struct {
	config *shared.Config

	TagGroupsApi *TagGroupsClient
	TagGroupsAPI *TagGroupsClient
	TagsApi      *TagsClient
	TagsAPI      *TagsClient
}

func New(config *shared.Config) *Client {
	tagGroupsClient := &TagGroupsClient{config: config}
	tagsClient := &TagsClient{config: config}

	return &Client{
		config:       config,
		TagGroupsApi: tagGroupsClient,
		TagGroupsAPI: tagGroupsClient,
		TagsApi:      tagsClient,
		TagsAPI:      tagsClient,
	}
}

type TagGroupsClient struct {
	config *shared.Config
}

type TagsClient struct {
	config *shared.Config
}
