package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// ShippingMethodDeleteInput provides the data required to delete a shipping method
type ShippingMethodDeleteInput struct {
	ID      string
	Version int
}

// ShippingMethodUpdateInput provides the data required to update a shipping method
type ShippingMethodUpdateInput struct {
	ID string

	// The expected version of the type on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict
	// will be returned.
	Version int

	// The list of update actions to be performed on the type.
	Actions []ShippingMethodUpdateAction
}

// ShippingMethodURLPath is the commercetools API shipping method path
const ShippingMethodURLPath = "shipping-methods"

// ShippingMethodGetByID will return a shipping method matching the provided ID.
// OAuth2 Scopes: view_orders:{projectKey} or manage_my_orders:{projectKey}
func (client *Client) ShippingMethodGetByID(id string) (result *ShippingMethod, err error) {
	err = client.Get(fmt.Sprintf("%s/%s", ShippingMethodURLPath, id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodGetByKey will return a shipping method matching the provided Key.
// OAuth2 Scopes: view_orders:{projectKey} or manage_my_orders:{projectKey}
func (client *Client) ShippingMethodGetByKey(key string) (result *ShippingMethod, err error) {
	err = client.Get(fmt.Sprintf("%s/key=%s", ShippingMethodURLPath, key), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodQuery will query the shipping methods.
// OAuth2 Scopes: view_orders:{projectKey} or manage_my_orders:{projectKey}
func (client *Client) ShippingMethodQuery(input *QueryInput) (result *ShippingMethodPagedQueryResponse, err error) {
	err = client.Query(ShippingMethodURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodCreate will create a new shipping method from a draft, and return
// the newly created shipping method.
// OAuth2 Scopes: manage_orders:{projectKey}
func (client *Client) ShippingMethodCreate(draft *ShippingMethodDraft) (result *ShippingMethod, err error) {
	err = client.Create(ShippingMethodURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodUpdate will update a shipping method matching the provided ID with
// the defined ShippingMethodUpdateActions.
// OAuth2 Scopes: manage_orders:{projectKey}
func (client *Client) ShippingMethodUpdate(input *ShippingMethodUpdateInput) (result *ShippingMethod, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("%s/%s", ShippingMethodURLPath, input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodDeleteByID will delete a shipping method matching the provided ID and version
// OAuth2 Scopes: manage_orders:{projectKey}
func (client *Client) ShippingMethodDeleteByID(id string, version int) (result *ShippingMethod, err error) {
	endpoint := fmt.Sprintf("%s/%s", ShippingMethodURLPath, id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
