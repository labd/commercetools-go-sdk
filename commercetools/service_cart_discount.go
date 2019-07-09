// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// CartDiscountURLPath is the commercetools API path.
const CartDiscountURLPath = "cart-discounts"

func (client *Client) CartDiscountCreate(draft *CartDiscountDraft) (result *CartDiscount, err error) {
	err = client.Create(CartDiscountURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CartDiscountQuery(input *QueryInput) (result *CartDiscountPagedQueryResponse, err error) {
	err = client.Query(CartDiscountURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CartDiscountDeleteWithKey(key string, version int) (result *CartDiscount, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("cart-discounts/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CartDiscountGetWithKey(key string) (result *CartDiscount, err error) {
	err = client.Get(strings.Replace("cart-discounts/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type CartDiscountUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CartDiscountUpdateAction
}

func (client *Client) CartDiscountUpdateWithKey(input *CartDiscountUpdateWithKeyInput) (result *CartDiscount, err error) {
	err = client.Update(strings.Replace("cart-discounts/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CartDiscountDeleteWithID(ID string, version int) (result *CartDiscount, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("cart-discounts/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CartDiscountGetWithID(ID string) (result *CartDiscount, err error) {
	err = client.Get(strings.Replace("cart-discounts/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type CartDiscountUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CartDiscountUpdateAction
}

func (client *Client) CartDiscountUpdateWithID(input *CartDiscountUpdateWithIDInput) (result *CartDiscount, err error) {
	err = client.Update(strings.Replace("cart-discounts/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
