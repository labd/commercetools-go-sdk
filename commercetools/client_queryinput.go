package commercetools

import (
	"net/url"
	"strconv"
)

// QueryInput provides the data required to query types.
type QueryInput struct {
	// The queryable APIs support ad-hoc filtering of resources through flexible
	// predicates. They do so via the where query parameter that accepts a
	// predicate expression to determine whether a specific resource
	// representation should be included in the result. The structure of
	// predicates and the names of the fields follow the structure and naming of
	// the fields in the documented response representation of the query
	// results.
	// https://docs.commercetools.com/http-api-query-predicates.html
	Where string

	// A query endpoint that supports sorting does so through the sort query
	// parameter. The provided value must be a valid sort expression. The
	// default sort direction is ASC. The allowed sort paths are typically
	// listed on the specific query endpoints. If multiple sort expressions are
	// specified via multiple sort parameters, they are combined into a composed
	// sort where the results are first sorted by the first expression, followed
	// by equal values being sorted according to the second expression, and so
	// on.
	// https://docs.commercetools.com/http-api.html#sorting
	Sort []string

	// Reference expansion is a feature of the resources listed in the table
	// below that enables clients to request server-side expansion of Reference
	// resources, thereby reducing the number of required client-server
	// roundtrips to obtain the data that a client needs for a specific
	// use-case. Reference expansion can be used when creating, updating,
	// querying, and deleting these resources.
	// https://docs.commercetools.com/http-api.html#reference-expansion
	Expand string

	Limit  int
	Offset int
}

func (qi QueryInput) toParams() (values url.Values) {
	values = url.Values{}

	if qi.Where != "" {
		values.Set("where", qi.Where)
	}

	for i := range qi.Sort {
		values.Add("sort", qi.Sort[i])
	}

	if qi.Expand != "" {
		values.Set("expand", qi.Expand)
	}

	if qi.Limit != 0 {
		values.Set("limit", strconv.Itoa(qi.Limit))
	}

	if qi.Offset != 0 {
		values.Set("offset", strconv.Itoa(qi.Offset))
	}

	return
}
