// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// ShippingMethodURLPath is the commercetools API path.
const ShippingMethodURLPath = "shipping-methods"

func (client *Client) ShippingMethodCreate(draft *ShippingMethodDraft) (result *ShippingMethod, err error) {
	err = client.Create(ShippingMethodURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ShippingMethodQuery(input *QueryInput) (result *ShippingMethodPagedQueryResponse, err error) {
	err = client.Query(ShippingMethodURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ShippingMethodDeleteWithKey(key string, version int) (result *ShippingMethod, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("shipping-methods/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ShippingMethodGetWithKey(key string) (result *ShippingMethod, err error) {
	err = client.Get(strings.Replace("shipping-methods/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ShippingMethodUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ShippingMethodUpdateAction
}

func (client *Client) ShippingMethodUpdateWithKey(input *ShippingMethodUpdateWithKeyInput) (result *ShippingMethod, err error) {
	err = client.Update(strings.Replace("shipping-methods/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ShippingMethodDeleteWithId(ID string, version int) (result *ShippingMethod, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("shipping-methods/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ShippingMethodGetWithId(ID string) (result *ShippingMethod, err error) {
	err = client.Get(strings.Replace("shipping-methods/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ShippingMethodUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []ShippingMethodUpdateAction
}

func (client *Client) ShippingMethodUpdateWithId(input *ShippingMethodUpdateWithIdInput) (result *ShippingMethod, err error) {
	err = client.Update(strings.Replace("shipping-methods/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
