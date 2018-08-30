package taxcategories

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// DeleteInput provides the data required to delete a tax category.
type DeleteInput struct {
	ID      string
	Version int
}

// UpdateInput provides the data required to update a tax category.
type UpdateInput struct {
	ID string

	// The expected version of the type on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict
	// will be returned.
	Version int

	// The list of update actions to be performed on the type.
	Actions commercetools.UpdateActions
}

// GetByID will return a tax category matching the provided ID. OAuth2 Scopes:
// view_products:{projectKey}
func (svc *Service) GetByID(id string) (result *TaxCategory, err error) {
	err = svc.client.Get(fmt.Sprintf("tax-categories/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create will create a new tax category from a draft, and return the newly
// created tax category. OAuth2 Scopes: manage_products:{projectKey}
func (svc *Service) Create(draft *TaxCategoryDraft) (result *TaxCategory, err error) {
	err = svc.client.Create("tax-categories", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update will update a tax category matching the provided ID with the defined
// UpdateActions. OAuth2 Scopes: manage_products:{projectKey}
func (svc *Service) Update(input *UpdateInput) (result *TaxCategory, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("tax-categories/%s", input.ID)
	err = svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByID will delete a tax category matching the provided ID. OAuth2
// Scopes: manage_products:{projectKey}
func (svc *Service) DeleteByID(id string, version int) (result *TaxCategory, err error) {
	endpoint := fmt.Sprintf("tax-categories/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByKey will delete a tax category matching the provided key. OAuth2
// Scopes: manage_products:{projectKey}
func (svc *Service) DeleteByKey(key string, version int) (result *TaxCategory, err error) {
	endpoint := fmt.Sprintf("tax-categories/key=%s", key)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
