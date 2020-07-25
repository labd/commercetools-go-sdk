// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
)

// ProductProjectionQuery allows querying for type ProductProjection
/*
You can use the product projections query endpoint to get the current or staged representations of Products.
When used with an API client that has the view_published_products:{projectKey} scope,
this endpoint only returns published (current) product projections.
*/
func (client *Client) ProductProjectionQuery(ctx context.Context, input *QueryInput) (result *ProductProjectionPagedQueryResponse, err error) {
	endpoint := "product-projections"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
ProductProjectionGetWithID Gets the current or staged representation of a product in a catalog by ID.
When used with an API client that has the view_published_products:{projectKey} scope,
this endpoint only returns published (current) product projections.
*/
func (client *Client) ProductProjectionGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *ProductProjection, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-projections/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
ProductProjectionGetWithKey Gets the current or staged representation of a product found by Key.
When used with an API client that has the view_published_products:{projectKey} scope,
this endpoint only returns published (current) product projections.
*/
func (client *Client) ProductProjectionGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *ProductProjection, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-projections/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductProjectionSearchInput is input for function ProductProjectionSearch
type ProductProjectionSearchInput struct {
	Text                 map[string]string `url:"text"`
	Facet                string            `url:"facet,omitempty"`
	FilterFacets         string            `url:"filter.facets,omitempty"`
	FilterQuery          string            `url:"filter.query,omitempty"`
	Filter               string            `url:"filter,omitempty"`
	Fuzzy                bool              `url:"fuzzy,omitempty"`
	FuzzyLevel           float64           `url:"fuzzyLevel,omitempty"`
	MarkMatchingVariants bool              `url:"markMatchingVariants,omitempty"`
	Staged               bool              `url:"staged,omitempty"`
}

/*
ProductProjectionSearch This endpoint provides high performance search queries over ProductProjections. The query result contains the
ProductProjections for which at least one ProductVariant matches the search query. This means that variants can
be included in the result also for which the search query does not match. To determine which ProductVariants match
the search query, the returned ProductProjections include the additional field isMatchingVariant.
*/
func (client *Client) ProductProjectionSearch(ctx context.Context, value *ProductProjectionSearchInput, opts ...RequestOption) (result *ProductProjectionPagedSearchResponse, err error) {
	params := serializeQueryParams(value)
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "product-projections/search"
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductProjectionSuggestInput is input for function ProductProjectionSuggest
type ProductProjectionSuggestInput struct {
	SearchKeywords map[string]string `url:"searchKeywords"`
	Fuzzy          bool              `url:"fuzzy,omitempty"`
	Staged         bool              `url:"staged,omitempty"`
}

// ProductProjectionSuggest The source of data for suggestions is the searchKeyword field in a product
func (client *Client) ProductProjectionSuggest(ctx context.Context, value *ProductProjectionSuggestInput, opts ...RequestOption) (result *ProductProjection, err error) {
	params := serializeQueryParams(value)
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "product-projections/suggest"
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
