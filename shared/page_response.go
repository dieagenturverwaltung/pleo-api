package shared

type Response[T any] struct {
	Data T `json:"data"`
}

type PageResponse[T any] struct {
	Response[T]
	Pagination Pagination `json:"pagination"`
}

type ListResponse[T any] struct {
	Data []T `json:"data"`
}

type CursorPageResponse[T any] struct {
	Data       []T              `json:"data"`
	Pagination CursorPagination `json:"pagination"`
}

type Pagination struct {
	CurrentRequestPagination struct {
		Parameters   map[string]any `json:"parameters"`
		After        string         `json:"after"`
		Before       string         `json:"before"`
		Limit        int            `json:"limit"`
		Offset       int            `json:"offset"`
		SortingKeys  []string       `json:"sortingKeys"`
		SortingOrder []string       `json:"sortingOrder"`
	} `json:"currentRequestPagination"`
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	EndCursor       string `json:"endCursor"`
	StartCursor     string `json:"startCursor"`
	Total           int    `json:"total"`
}

type CursorPagination struct {
	CurrentRequestPagination CursorPageCurrentRequestInfo `json:"currentRequestPagination"`
	HasNextPage              bool                         `json:"hasNextPage"`
	HasPreviousPage          bool                         `json:"hasPreviousPage"`
	EndCursor                string                       `json:"endCursor"`
	StartCursor              string                       `json:"startCursor"`
	Total                    int64                        `json:"total"`
}

type CursorPageCurrentRequestInfo struct {
	Parameters   map[string][]string `json:"parameters"`
	After        string              `json:"after"`
	Before       string              `json:"before"`
	Limit        int                 `json:"limit"`
	Offset       int64               `json:"offset"`
	SortingKeys  []string            `json:"sortingKeys"`
	SortingOrder []string            `json:"sortingOrder"`
}
