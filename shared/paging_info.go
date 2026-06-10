package shared

import (
	"net/url"
	"strconv"
)

type PagingInfo struct {
	Before       *string
	After        *string
	Offset       *int64
	Limit        *int
	SortingKeys  []string
	SortingOrder []string
}

func (pi PagingInfo) Apply(values url.Values) {
	if pi.Before != nil {
		values.Add("before", *pi.Before)
	}

	if pi.After != nil {
		values.Add("after", *pi.After)
	}

	if pi.Offset != nil {
		values.Add("offset", strconv.FormatInt(*pi.Offset, 10))
	}

	if pi.Limit != nil {
		values.Add("limit", strconv.Itoa(*pi.Limit))
	}

	for _, sortingKey := range pi.SortingKeys {
		values.Add("sorting_keys", sortingKey)
	}

	for _, sortingOrder := range pi.SortingOrder {
		values.Add("sorting_order", sortingOrder)
	}
}

func (pi PagingInfo) ToQueryString() string {
	var values = make(url.Values)
	pi.Apply(values)
	return values.Encode()
}
