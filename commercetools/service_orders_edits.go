// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// OrdersEditsURLPath is the commercetools API path.
const OrdersEditsURLPath = "orders/edits"

// OrdersEditsCreate creates a new instance of type OrderEdit
func (client *Client) OrdersEditsCreate(draft *OrderEditDraft) (result *OrderEdit, err error) {
	err = client.Create(OrdersEditsURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrdersEditsQuery allows querying for type OrderEdit
func (client *Client) OrdersEditsQuery(input *QueryInput) (result *OrderEditPagedQueryResponse, err error) {
	err = client.Query(OrdersEditsURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrdersEditsDeleteWithKey for type OrderEdit
func (client *Client) OrdersEditsDeleteWithKey(key string, version int) (result *OrderEdit, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("orders/edits/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrdersEditsGetWithKey for type OrderEdit
func (client *Client) OrdersEditsGetWithKey(key string) (result *OrderEdit, err error) {
	err = client.Get(strings.Replace("orders/edits/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderEditUpdateWithKeyInput is input for function OrdersEditsUpdateWithKey
type OrderEditUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []OrderEditUpdateAction
}

// OrdersEditsUpdateWithKey for type OrderEdit
func (client *Client) OrdersEditsUpdateWithKey(input *OrderEditUpdateWithKeyInput) (result *OrderEdit, err error) {
	err = client.Update(strings.Replace("orders/edits/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrdersEditsDeleteWithID for type OrderEdit
func (client *Client) OrdersEditsDeleteWithID(ID string, version int) (result *OrderEdit, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("orders/edits/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrdersEditsGetWithID for type OrderEdit
func (client *Client) OrdersEditsGetWithID(ID string) (result *OrderEdit, err error) {
	err = client.Get(strings.Replace("orders/edits/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderEditUpdateWithIDInput is input for function OrdersEditsUpdateWithID
type OrderEditUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []OrderEditUpdateAction
}

// OrdersEditsUpdateWithID for type OrderEdit
func (client *Client) OrdersEditsUpdateWithID(input *OrderEditUpdateWithIDInput) (result *OrderEdit, err error) {
	err = client.Update(strings.Replace("orders/edits/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
