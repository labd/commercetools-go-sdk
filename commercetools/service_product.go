// Automatically generated, do not edit

package commercetools

import "strings"

// ProductProjectionURLPath is the commercetools API path.
const ProductProjectionURLPath = "product-projections"

// ProductProjectionQuery allows querying for type ProductProjection
func (client *Client) ProductProjectionQuery(input *QueryInput) (result *ProductProjectionPagedQueryResponse, err error) {
	err = client.Query(ProductProjectionURLPath, input.toParams(), &result)
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
func (client *Client) ProductProjectionGetWithKey(key string) (result *ProductProjection, err error) {
	err = client.Get(strings.Replace("product-projections/key={key}", "{key}", key, 1), nil, &result)
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
func (client *Client) ProductProjectionGetWithID(ID string) (result *ProductProjection, err error) {
	err = client.Get(strings.Replace("product-projections/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
