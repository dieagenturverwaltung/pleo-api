package tags

type PageOrder string

const (
	PageOrderAsc            PageOrder = "ASC"
	PageOrderAscNullsFirst  PageOrder = "ASC_NULLS_FIRST"
	PageOrderAscNullsLast   PageOrder = "ASC_NULLS_LAST"
	PageOrderDesc           PageOrder = "DESC"
	PageOrderDescNullsFirst PageOrder = "DESC_NULLS_FIRST"
	PageOrderDescNullsLast  PageOrder = "DESC_NULLS_LAST"
)

type TagGroupCreateModel struct {
	Archived bool              `json:"archived"`
	Code     string            `json:"code"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Name     string            `json:"name"`
}

type TagGroupUpdateModel struct {
	Archived bool              `json:"archived"`
	Code     string            `json:"code"`
	Metadata map[string]string `json:"metadata"`
	Name     string            `json:"name"`
}

type TagGroupDimensionCreateModel struct {
	Code         string `json:"code"`
	DisplayOrder int    `json:"displayOrder"`
	Name         string `json:"name"`
	Visible      bool   `json:"visible"`
}

type TagGroupDimensionUpdateModel struct {
	Code         string `json:"code"`
	DisplayOrder int    `json:"displayOrder"`
	Name         string `json:"name"`
	Visible      bool   `json:"visible"`
}

type TagCreateModel struct {
	Archived bool   `json:"archived"`
	Code     string `json:"code"`
	Name     string `json:"name,omitempty"`
}

type TagUpdateModel struct {
	Archived bool   `json:"archived"`
	Code     string `json:"code"`
	Name     string `json:"name,omitempty"`
}

type TagDimensionValueCreateModel struct {
	Value string `json:"value"`
}

type TagDimensionValueUpdateModel struct {
	Value string `json:"value"`
}

type AggregatedTagSearchRequest struct {
	TagIDs     []string `json:"tagIds,omitempty"`
	TextSearch string   `json:"textSearch,omitempty"`
}

type TagGroupModel struct {
	Archived  bool              `json:"archived"`
	Code      string            `json:"code"`
	CompanyID string            `json:"companyId"`
	CreatedAt string            `json:"createdAt"`
	ID        string            `json:"id"`
	Metadata  map[string]string `json:"metadata"`
	Name      string            `json:"name"`
	UpdatedAt string            `json:"updatedAt"`
}

type TagGroupDimensionModel struct {
	Code         string `json:"code"`
	CreatedAt    string `json:"createdAt"`
	DisplayOrder int    `json:"displayOrder"`
	GroupID      string `json:"groupId"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	UpdatedAt    string `json:"updatedAt"`
	Visible      bool   `json:"visible"`
}

type TagModel struct {
	Archived  bool   `json:"archived"`
	Code      string `json:"code"`
	CreatedAt string `json:"createdAt"`
	GroupID   string `json:"groupId"`
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	UpdatedAt string `json:"updatedAt"`
}

type TagDimensionValueModel struct {
	CreatedAt   string `json:"createdAt"`
	DimensionID string `json:"dimensionId"`
	TagID       string `json:"tagId"`
	UpdatedAt   string `json:"updatedAt"`
	Value       string `json:"value"`
}

type AggregatedTagGroupModel struct {
	Archived   bool                               `json:"archived"`
	Code       string                             `json:"code"`
	CompanyID  string                             `json:"companyId"`
	CreatedAt  string                             `json:"createdAt"`
	Dimensions []AggregatedTagGroupDimensionModel `json:"dimensions"`
	ID         string                             `json:"id"`
	Name       string                             `json:"name"`
	UpdatedAt  string                             `json:"updatedAt"`
}

type AggregatedTagGroupDimensionModel struct {
	Code         string `json:"code"`
	DisplayOrder int    `json:"displayOrder"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	Visible      bool   `json:"visible"`
}

type AggregatedTagModel struct {
	Archived   bool                           `json:"archived"`
	Code       string                         `json:"code"`
	CompanyID  string                         `json:"companyId"`
	CreatedAt  string                         `json:"createdAt"`
	Dimensions []AggregatedTagDimensionModel  `json:"dimensions"`
	Group      AggregatedTagGroupSummaryModel `json:"group"`
	ID         string                         `json:"id"`
	Name       string                         `json:"name,omitempty"`
	UpdatedAt  string                         `json:"updatedAt"`
}

type AggregatedTagDimensionModel struct {
	Code         string `json:"code"`
	DisplayOrder int    `json:"displayOrder"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	Value        string `json:"value"`
	Visible      bool   `json:"visible"`
}

type AggregatedTagGroupSummaryModel struct {
	Code string `json:"code"`
	ID   string `json:"id"`
	Name string `json:"name"`
}
