// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
	"strings"
)

// AttributesForFacetingOption is a wrapper for an AttributesForFaceting option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type AttributesForFacetingOption struct {
	value []string
}

// AttributesForFaceting wraps the given value into a AttributesForFacetingOption.
func AttributesForFaceting(v ...string) *AttributesForFacetingOption {
	if v == nil {
		return &AttributesForFacetingOption{[]string{}}
	}
	return &AttributesForFacetingOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *AttributesForFacetingOption) Get() []string {
	if o == nil {
		return []string{}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// AttributesForFacetingOption.
func (o AttributesForFacetingOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// AttributesForFacetingOption.
func (o *AttributesForFacetingOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
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
func (o *AttributesForFacetingOption) Equal(o2 *AttributesForFacetingOption) bool {
	if o == nil {
		return o2 == nil || reflect.DeepEqual(o2.value, []string{})
	}
	if o2 == nil {
		return o == nil || reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// AttributesForFacetingEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func AttributesForFacetingEqual(o1, o2 *AttributesForFacetingOption) bool {
	return o1.Equal(o2)
}
