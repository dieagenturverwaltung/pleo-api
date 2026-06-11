package shared

import (
	"net/url"
	"strconv"
)

func URLWithQuery(path string, values url.Values) string {
	if len(values) == 0 {
		return path
	}

	return path + "?" + values.Encode()
}

func AddQueryString(values url.Values, name string, value *string) {
	if value != nil {
		values.Add(name, *value)
	}
}

func AddQueryStringValue(values url.Values, name string, value string) {
	if value != "" {
		values.Add(name, value)
	}
}

func AddQueryStrings[T ~string](values url.Values, name string, items []T) {
	for _, item := range items {
		values.Add(name, string(item))
	}
}

func AddQueryBool(values url.Values, name string, value *bool) {
	if value != nil {
		values.Add(name, strconv.FormatBool(*value))
	}
}

func AddQueryInt64(values url.Values, name string, value *int64) {
	if value != nil {
		values.Add(name, strconv.FormatInt(*value, 10))
	}
}

func AddQueryCompanyID(values url.Values, companyID *string) {
	if companyID != nil {
		values.Add("company_id", *companyID)
	}
}
