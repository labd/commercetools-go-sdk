package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// ProductTypeDeleteInput provides the data required to delete a product type.
type ProductTypeDeleteInput struct {
	ID      string
	Version int
}

// ProductTypeUpdateInput provides the data required to update a product type.
type ProductTypeUpdateInput struct {
	ID string

	// The expected version of the product type on which the changes should be
	// applied. If the expected version does not match the actual version, a 409
	// Conflict will be returned.
	Version int

	// The list of update actions to be performed on the product type.
	Actions []ProductTypeUpdateAction
}

// ProductTypeGetByID will return a product type matching the provided ID. OAuth2 Scopes:
// view_products:{projectKey}
func (client *Client) ProductTypeGetByID(id string) (result *ProductType, err error) {
	err = client.Get(fmt.Sprintf("product-types/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeCreate will create a new product type from a draft, and return the
// newly created product type. OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) ProductTypeCreate(draft *ProductTypeDraft) (result *ProductType, err error) {
	err = client.Create("product-types", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeUpdate will update a product type matching the provided ID with
// the defined ProductTypeUpdateActions. OAuth2 Scopes:
// manage_products:{projectKey}
func (client *Client) ProductTypeUpdate(input *ProductTypeUpdateInput) (result *ProductType, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("No valid product type id passed")
	}

	endpoint := fmt.Sprintf("product-types/%s", input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeDeleteByID will delete a product type matching the provided ID.
// This request deletes a product type only if it’s not referenced by a product.
// OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) ProductTypeDeleteByID(id string, version int) (result *ProductType, err error) {
	endpoint := fmt.Sprintf("product-types/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeDeleteByKey will delete a product type matching the provided key.
// This request deletes a product type only if it’s not referenced by a product.
// OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) ProductTypeDeleteByKey(key string, version int) (result *ProductType, err error) {
	endpoint := fmt.Sprintf("product-types/key=%s", key)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
