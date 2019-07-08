// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// TaxCategoryURLPath is the commercetools API path.
const TaxCategoryURLPath = "tax-categories"

func (client *Client) TaxCategoryCreate(draft *TaxCategoryDraft) (result *TaxCategory, err error) {
	err = client.Create(TaxCategoryURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) TaxCategoryQuery(input *QueryInput) (result *TaxCategoryPagedQueryResponse, err error) {
	err = client.Query(TaxCategoryURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) TaxCategoryDeleteWithId(ID string, version int) (result *TaxCategory, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("tax-categories/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) TaxCategoryGetWithId(ID string) (result *TaxCategory, err error) {
	err = client.Get(strings.Replace("tax-categories/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type TaxCategoryUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []TaxCategoryUpdateAction
}

func (client *Client) TaxCategoryUpdateWithId(input *TaxCategoryUpdateWithIdInput) (result *TaxCategory, err error) {
	err = client.Update(strings.Replace("tax-categories/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
