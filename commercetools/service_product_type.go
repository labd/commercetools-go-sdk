// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// ProductTypeURLPath is the commercetools API path.
const ProductTypeURLPath = "product-types"

func (client *Client) ProductTypeCreate(draft *ProductTypeDraft) (result *ProductType, err error) {
	err = client.Create(ProductTypeURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductTypeQuery(input *QueryInput) (result *ProductTypePagedQueryResponse, err error) {
	err = client.Query(ProductTypeURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductTypeDeleteWithKey(key string, version int) (result *ProductType, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("product-types/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductTypeGetWithKey(key string) (result *ProductType, err error) {
	err = client.Get(strings.Replace("product-types/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ProductTypeUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ProductTypeUpdateAction
}

func (client *Client) ProductTypeUpdateWithKey(input *ProductTypeUpdateWithKeyInput) (result *ProductType, err error) {
	err = client.Update(strings.Replace("product-types/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductTypeDeleteWithID(ID string, version int) (result *ProductType, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("product-types/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductTypeGetWithID(ID string) (result *ProductType, err error) {
	err = client.Get(strings.Replace("product-types/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ProductTypeUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ProductTypeUpdateAction
}

func (client *Client) ProductTypeUpdateWithID(input *ProductTypeUpdateWithIDInput) (result *ProductType, err error) {
	err = client.Update(strings.Replace("product-types/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
