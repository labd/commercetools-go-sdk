// Automatically generated, do not edit

package commercetools

// ProductProjectionsSearchURLPath is the commercetools API path.
const ProductProjectionsSearchURLPath = "product-projections/search"

// ProductProjectionsSearchCreate creates a new instance of type ProductProjectionPagedSearchResponse
func (client *Client) ProductProjectionsSearchCreate() (result *ProductProjectionPagedSearchResponse, err error) {
	err = client.Create(ProductProjectionsSearchURLPath, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductProjectionsSearchQuery allows querying for type ProductProjectionPagedSearchResponse
func (client *Client) ProductProjectionsSearchQuery(input *QueryInput) (result *ProductProjectionPagedSearchResponse, err error) {
	err = client.Query(ProductProjectionsSearchURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
