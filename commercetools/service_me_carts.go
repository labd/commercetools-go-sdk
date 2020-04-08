// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// MeCartsURLPath is the commercetools API path.
const MeCartsURLPath = "me/carts"

// MeCartsCreate creates a new instance of type MyCart
func (client *Client) MeCartsCreate(draft *MyCartDraft) (result *MyCart, err error) {
	err = client.Create(MeCartsURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MeCartsQuery allows querying for type MyCart
func (client *Client) MeCartsQuery(input *QueryInput) (result *CartPagedQueryResponse, err error) {
	err = client.Query(MeCartsURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MeCartsDeleteWithID for type MyCart
func (client *Client) MeCartsDeleteWithID(ID string, version int) (result *MyCart, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("me/carts/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MeCartsGetWithID for type MyCart
func (client *Client) MeCartsGetWithID(ID string) (result *MyCart, err error) {
	err = client.Get(strings.Replace("me/carts/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyCartUpdateWithIDInput is input for function MeCartsUpdateWithID
type MyCartUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []MyCartUpdateAction
}

// MeCartsUpdateWithID for type MyCart
func (client *Client) MeCartsUpdateWithID(input *MyCartUpdateWithIDInput) (result *MyCart, err error) {
	err = client.Update(strings.Replace("me/carts/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
