// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractFacets returns the first found FacetsOption from the
// given variadic arguments or nil otherwise.
func ExtractFacets(opts ...interface{}) *opt.FacetsOption {
	for _, o := range opts {
		if v, ok := o.(*opt.FacetsOption); ok {
			return v
		}
	}
	return nil
}