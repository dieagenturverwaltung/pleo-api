package shared

import (
	"encoding/json"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func (ct *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(time.RFC3339, s)
	return
}

func (ct Time) MarshalJSON() ([]byte, error) {
	format := ct.Time.Format(time.RFC3339)
	return []byte(`"` + format + `"`), nil
}

type NullableTime struct {
	value *Time
	isSet bool
}

func (v NullableTime) Get() *Time {
	return v.value
}

func (v *NullableTime) Set(val *Time) {
	v.value = val
	v.isSet = true
}

func (v NullableTime) IsSet() bool {
	return v.isSet
}

func (v *NullableTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTime(val *Time) *NullableTime {
	return &NullableTime{value: val, isSet: true}
}

func (v NullableTime) MarshalJSON() ([]byte, error) {
	if !v.isSet {
		return []byte("null"), nil
	}

	return v.value.MarshalJSON()
}

func (v *NullableTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
