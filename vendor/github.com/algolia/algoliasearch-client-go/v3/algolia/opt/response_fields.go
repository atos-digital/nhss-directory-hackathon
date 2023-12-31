// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ResponseFieldsOption is a wrapper for an ResponseFields option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type ResponseFieldsOption struct {
	value []string
}

// ResponseFields wraps the given value into a ResponseFieldsOption.
func ResponseFields(v ...string) *ResponseFieldsOption {
	if v == nil {
		return &ResponseFieldsOption{[]string{}}
	}
	return &ResponseFieldsOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *ResponseFieldsOption) Get() []string {
	if o == nil {
		return []string{"*"}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// ResponseFieldsOption.
func (o ResponseFieldsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// ResponseFieldsOption.
func (o *ResponseFieldsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{"*"}
		return nil
	}
	var s string
	err := json.Unmarshal(data, &s)
	if err == nil {
		o.value = strings.Split(s, ",")
		if len(o.value) == 1 && o.value[0] == "" {
			o.value = []string{}
		}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *ResponseFieldsOption) Equal(o2 *ResponseFieldsOption) bool {
	if o == nil {
		return o2 == nil || reflect.DeepEqual(o2.value, []string{"*"})
	}
	if o2 == nil {
		return o == nil || reflect.DeepEqual(o.value, []string{"*"})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// ResponseFieldsEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func ResponseFieldsEqual(o1, o2 *ResponseFieldsOption) bool {
	return o1.Equal(o2)
}
