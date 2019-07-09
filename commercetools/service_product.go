// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// ProductURLPath is the commercetools API path.
const ProductURLPath = "products"

func (client *Client) ProductCreate(draft *ProductDraft) (result *Product, err error) {
	err = client.Create(ProductURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductQuery(input *QueryInput) (result *ProductPagedQueryResponse, err error) {
	err = client.Query(ProductURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductDeleteWithKey(key string, version int) (result *Product, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("products/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductGetWithKey(key string) (result *Product, err error) {
	err = client.Get(strings.Replace("products/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ProductUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ProductUpdateAction
}

func (client *Client) ProductUpdateWithKey(input *ProductUpdateWithKeyInput) (result *Product, err error) {
	err = client.Update(strings.Replace("products/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductDeleteWithID(ID string, version int) (result *Product, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("products/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductGetWithID(ID string) (result *Product, err error) {
	err = client.Get(strings.Replace("products/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ProductUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ProductUpdateAction
}

func (client *Client) ProductUpdateWithID(input *ProductUpdateWithIDInput) (result *Product, err error) {
	err = client.Update(strings.Replace("products/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
