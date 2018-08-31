package types

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/labd/commercetools-go-sdk/commercetools"
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

// DeleteInput provides the data required to delete a type.
type DeleteInput struct {
	ID      string
	Version int
}

// UpdateInput provides the data required to update a type.
type UpdateInput struct {
	ID string
	// The expected version of the type on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict will be returned.
	Version int
	// The list of update actions to be performed on the type.
	Actions commercetools.UpdateActions
}

// GetByID will return a type matching the provided ID.
// OAuth2 Scopes: view_types:{projectKey}
func (svc *Service) GetByID(id string) (result *Type, err error) {
	err = svc.client.Get(fmt.Sprintf("types/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Query will return a list of types matching the provided query input. OAuth2
// Scopes: view_types:{projectKey}
func (svc *Service) Query(input *QueryInput) (result *commercetools.PagedQueryResult, err error) {
	params := input.toParams()
	err = svc.client.Query("types", params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create will create a new type from a draft, and return the newly created type.
// OAuth2 Scopes: manage_types:{projectKey}
func (svc *Service) Create(draft *TypeDraft) (result *Type, err error) {
	err = svc.client.Create("types", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update will update a type matching the provided ID with the defined UpdateActions.
// OAuth2 Scopes: manage_types:{projectKey}
func (svc *Service) Update(input *UpdateInput) (result *Type, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("types/%s", input.ID)
	err = svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByID will delete a type matching the provided ID.
// These requests delete a type only if it’s not referenced by other entities.
// OAuth2 Scopes: manage_types:{projectKey}
func (svc *Service) DeleteByID(id string, version int) (result *Type, err error) {
	endpoint := fmt.Sprintf("types/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByKey will delete a type matching the provided key.
// These requests delete a type only if it’s not referenced by other entities.
// OAuth2 Scopes: manage_types:{projectKey}
func (svc *Service) DeleteByKey(key string, version int) (result *Type, err error) {
	endpoint := fmt.Sprintf("types/key=%s", key)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
