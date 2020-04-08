// Automatically generated, do not edit

package commercetools

import "strings"

// MeOrdersURLPath is the commercetools API path.
const MeOrdersURLPath = "me/orders"

// MeOrdersCreate creates a new instance of type MyOrder
func (client *Client) MeOrdersCreate(draft *MyOrderFromCartDraft) (result *MyOrder, err error) {
	err = client.Create(MeOrdersURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MeOrdersQuery allows querying for type MyOrder
func (client *Client) MeOrdersQuery(input *QueryInput) (result *OrderPagedQueryResponse, err error) {
	err = client.Query(MeOrdersURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MeOrdersGetWithID for type MyOrder
func (client *Client) MeOrdersGetWithID(ID string) (result *MyOrder, err error) {
	err = client.Get(strings.Replace("me/orders/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyOrderUpdateWithIDInput is input for function MeOrdersUpdateWithID
type MyOrderUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []MyOrderUpdateAction
}

// MeOrdersUpdateWithID for type MyOrder
func (client *Client) MeOrdersUpdateWithID(input *MyOrderUpdateWithIDInput) (result *MyOrder, err error) {
	err = client.Update(strings.Replace("me/orders/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
