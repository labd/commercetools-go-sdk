// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// CustomerGroupURLPath is the commercetools API path.
const CustomerGroupURLPath = "customer-groups"

// CustomerGroupCreate creates a new instance of type CustomerGroup
func (client *Client) CustomerGroupCreate(ctx context.Context, draft *CustomerGroupDraft, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, CustomerGroupURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupQuery allows querying for type CustomerGroup
func (client *Client) CustomerGroupQuery(ctx context.Context, input *QueryInput) (result *CustomerGroupPagedQueryResponse, err error) {
	err = client.Query(ctx, CustomerGroupURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupDeleteWithKey for type CustomerGroup
func (client *Client) CustomerGroupDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("customer-groups/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupGetWithKey Gets a customer group by Key.
func (client *Client) CustomerGroupGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("customer-groups/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupUpdateWithKeyInput is input for function CustomerGroupUpdateWithKey
type CustomerGroupUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CustomerGroupUpdateAction
}

// CustomerGroupUpdateWithKey Updates a customer group by Key.
func (client *Client) CustomerGroupUpdateWithKey(ctx context.Context, input *CustomerGroupUpdateWithKeyInput, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("customer-groups/key={key}", "{key}", input.Key, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupDeleteWithID for type CustomerGroup
func (client *Client) CustomerGroupDeleteWithID(ctx context.Context, ID string, version int, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("customer-groups/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupGetWithID for type CustomerGroup
func (client *Client) CustomerGroupGetWithID(ctx context.Context, ID string, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("customer-groups/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupUpdateWithIDInput is input for function CustomerGroupUpdateWithID
type CustomerGroupUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CustomerGroupUpdateAction
}

// CustomerGroupUpdateWithID for type CustomerGroup
func (client *Client) CustomerGroupUpdateWithID(ctx context.Context, input *CustomerGroupUpdateWithIDInput, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("customer-groups/{ID}", "{ID}", input.ID, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
