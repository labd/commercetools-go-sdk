// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// OrderURLPath is the commercetools API path.
const OrderURLPath = "orders"

func (client *Client) OrderCreate(draft *OrderFromCartDraft) (result *Order, err error) {
	err = client.Create(OrderURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) OrderQuery(input *QueryInput) (result *OrderPagedQueryResponse, err error) {
	err = client.Query(OrderURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) OrderDeleteWithId(ID string, version int, dataErasure bool) (result *Order, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("orders/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) OrderGetWithId(ID string) (result *Order, err error) {
	err = client.Get(strings.Replace("orders/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type OrderUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []OrderUpdateAction
}

func (client *Client) OrderUpdateWithId(input *OrderUpdateWithIdInput) (result *Order, err error) {
	err = client.Update(strings.Replace("orders/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
