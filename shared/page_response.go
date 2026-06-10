package shared

type Response[T any] struct {
	Data T `json:"data"`
}

type PageResponse[T any] struct {
	Response[T]
	Pagination Pagination `json:"pagination"`
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
