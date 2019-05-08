package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// CartUpdateInput provides the data required to update a cart.
type CartUpdateInput struct {
	ID string

	// The expected version of the cart on which the changes should be
	// applied. If the expected version does not match the actual version, a 409
	// Conflict will be returned.
	Version int

	// The list of update actions to be performed on the cart.
	Actions []CartUpdateAction
}

// CartURLPath is the commercetools API cart path.
const CartURLPath = "carts"

// CartGetByID will return a cart matching the provided ID. OAuth2 Scopes:
// view_orders:{projectKey}
func (client *Client) CartGetByID(id string) (result *Cart, err error) {
	err = client.Get(fmt.Sprintf("%s/%s", CartURLPath, id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartCreate will create a new cart from a draft, and return the newly created
// cart. OAuth2 Scopes: manage_orders:{projectKey}
func (client *Client) CartCreate(draft *CartDraft) (result *Cart, err error) {
	err = client.Create(CartURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartUpdate will update a cart matching the provided ID with the defined
// CartUpdateActions. OAuth2 Scopes: manage_orders:{projectKey}
func (client *Client) CartUpdate(input *CartUpdateInput) (result *Cart, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("%s/%s", CartURLPath, input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDelete will delete a type matching the provided ID. These requests
// delete a type only if it’s not referenced by other entities. OAuth2 Scopes:
// manage_orders:{projectKey}
func (client *Client) CartDelete(id string, version int, dataErasure bool) (result *Cart, err error) {
	endpoint := fmt.Sprintf("%s/%s", CartURLPath, id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartQuery will query carts.
// OAuth2 Scopes: view_orders:{projectKey}
func (client *Client) CartQuery(input *QueryInput) (result *CartPagedQueryResponse, err error) {
	err = client.Query(CartURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartReplicate will replicate an existing cart.
// OAuth2 Scopes: manage_orders:{projectKey}
func (client *Client) CartReplicate(input *ReplicaCartDraft) (result *Cart, err error) {
	err = client.Create(fmt.Sprintf("%s/replicate", CartURLPath), nil, input, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
