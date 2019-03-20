package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// TaxCategoryDeleteInput provides the data required to delete a tax category.
type TaxCategoryDeleteInput struct {
	ID      string
	Version int
}

// TaxCategoryUpdateInput provides the data required to update a tax category.
type TaxCategoryUpdateInput struct {
	ID string

	// The expected version of the type on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict
	// will be returned.
	Version int

	// The list of update actions to be performed on the type.
	Actions []TaxCategoryUpdateAction
}

// TaxCategoryURLPath is the commercetools API tax category path.
const TaxCategoryURLPath = "tax-categories"

// TaxCategoryGetByID will return a tax category matching the provided ID.
// OAuth2 Scopes: view_products:{projectKey}
func (client *Client) TaxCategoryGetByID(id string) (result *TaxCategory, err error) {
	err = client.Get(fmt.Sprintf("%s/%s", TaxCategoryURLPath, id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryCreate will create a new tax category from a draft, and return the
// newly created tax category. OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) TaxCategoryCreate(draft *TaxCategoryDraft) (result *TaxCategory, err error) {
	err = client.Create(TaxCategoryURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryUpdate will update a tax category matching the provided ID with
// the defined TaxCategoryUpdateActions. OAuth2 Scopes:
// manage_products:{projectKey}
func (client *Client) TaxCategoryUpdate(input *TaxCategoryUpdateInput) (result *TaxCategory, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("%s/%s", TaxCategoryURLPath, input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryDeleteByID will delete a tax category matching the provided ID.
// OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) TaxCategoryDeleteByID(id string, version int) (result *TaxCategory, err error) {
	endpoint := fmt.Sprintf("%s/%s", TaxCategoryURLPath, id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryDeleteByKey will delete a tax category matching the provided key.
// OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) TaxCategoryDeleteByKey(key string, version int) (result *TaxCategory, err error) {
	endpoint := fmt.Sprintf("%s/key=%s", TaxCategoryURLPath, key)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryQuery will query the tax categories.
// OAuth2 Scopes: view_products:{projectKey}
func (client *Client) TaxCategoryQuery(input *QueryInput) (result *TaxCategoryPagedQueryResponse, err error) {
	err = client.Query(TaxCategoryURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
