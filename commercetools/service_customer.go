// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// CustomerURLPath is the commercetools API path.
const CustomerURLPath = "customers"

// CustomerCreate creates a new instance of type Customer
func (client *Client) CustomerCreate(ctx context.Context, draft *CustomerDraft, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Create(ctx, CustomerURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerQuery allows querying for type Customer
func (client *Client) CustomerQuery(ctx context.Context, input *QueryInput) (result *CustomerPagedQueryResponse, err error) {
	err = client.Query(ctx, CustomerURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerDeleteWithKey for type Customer
func (client *Client) CustomerDeleteWithKey(ctx context.Context, key string, version int, dataErasure bool, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("customers/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGetWithKey for type Customer
func (client *Client) CustomerGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("customers/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerUpdateWithKeyInput is input for function CustomerUpdateWithKey
type CustomerUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CustomerUpdateAction
}

// CustomerUpdateWithKey for type Customer
func (client *Client) CustomerUpdateWithKey(ctx context.Context, input *CustomerUpdateWithKeyInput, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("customers/key={key}", "{key}", input.Key, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerDeleteWithID for type Customer
func (client *Client) CustomerDeleteWithID(ctx context.Context, ID string, version int, dataErasure bool, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("customers/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGetWithID for type Customer
func (client *Client) CustomerGetWithID(ctx context.Context, ID string, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("customers/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerUpdateWithIDInput is input for function CustomerUpdateWithID
type CustomerUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CustomerUpdateAction
}

// CustomerUpdateWithID for type Customer
func (client *Client) CustomerUpdateWithID(ctx context.Context, input *CustomerUpdateWithIDInput, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("customers/{ID}", "{ID}", input.ID, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
