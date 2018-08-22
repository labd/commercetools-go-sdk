package producttypes

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// DeleteInput provides the data required to delete a product type.
type DeleteInput struct {
	ID      string
	Version int
}

// UpdateInput provides the data required to update a product type.
type UpdateInput struct {
	ID string
	// The expected version of the product type on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict will be returned.
	Version int
	// The list of update actions to be performed on the product type.
	Actions commercetools.UpdateActions
}

// GetByID will return a product type matching the provided ID.
// OAuth2 Scopes: view_products:{projectKey}
func (svc *Service) GetByID(id string) (result *ProductType, err error) {
	err = svc.client.Get(fmt.Sprintf("product-types/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create will create a new product type from a draft, and return the newly created product type.
// OAuth2 Scopes: manage_products:{projectKey}
func (svc *Service) Create(draft *ProductTypeDraft) (result *ProductType, err error) {
	err = svc.client.Create("product-types", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update will update a product type matching the provided ID with the defined UpdateActions.
// OAuth2 Scopes: manage_products:{projectKey}
func (svc *Service) Update(input *UpdateInput) (result *ProductType, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("No valid product type id passed")
	}

	endpoint := fmt.Sprintf("product-types/%s", input.ID)
	err = svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByID will delete a product type matching the provided ID.
// This request deletes a product type only if it’s not referenced by a product.
// OAuth2 Scopes: manage_products:{projectKey}
func (svc *Service) DeleteByID(id string, version int) (result *ProductType, err error) {
	endpoint := fmt.Sprintf("product-types/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByKey will delete a product type matching the provided key.
// This request deletes a product type only if it’s not referenced by a product.
// OAuth2 Scopes: manage_products:{projectKey}
func (svc *Service) DeleteByKey(key string, version int) (result *ProductType, err error) {
	endpoint := fmt.Sprintf("product-types/key=%s", key)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
