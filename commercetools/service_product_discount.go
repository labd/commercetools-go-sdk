// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// ProductDiscountURLPath is the commercetools API path.
const ProductDiscountURLPath = "product-discounts"

func (client *Client) ProductDiscountCreate(draft *ProductDiscountDraft) (result *ProductDiscount, err error) {
	err = client.Create(ProductDiscountURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductDiscountQuery(input *QueryInput) (result *ProductDiscountPagedQueryResponse, err error) {
	err = client.Query(ProductDiscountURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductDiscountDeleteWithKey(key string, version int) (result *ProductDiscount, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("product-discounts/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductDiscountGetWithKey(key string) (result *ProductDiscount, err error) {
	err = client.Get(strings.Replace("product-discounts/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ProductDiscountUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ProductDiscountUpdateAction
}

func (client *Client) ProductDiscountUpdateWithKey(input *ProductDiscountUpdateWithKeyInput) (result *ProductDiscount, err error) {
	err = client.Update(strings.Replace("product-discounts/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductDiscountDeleteWithId(ID string, version int) (result *ProductDiscount, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("product-discounts/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ProductDiscountGetWithId(ID string) (result *ProductDiscount, err error) {
	err = client.Get(strings.Replace("product-discounts/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ProductDiscountUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []ProductDiscountUpdateAction
}

func (client *Client) ProductDiscountUpdateWithId(input *ProductDiscountUpdateWithIdInput) (result *ProductDiscount, err error) {
	err = client.Update(strings.Replace("product-discounts/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
