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
func (client *Client) CustomerCreate(ctx context.Context, draft *CustomerDraft) (result *Customer, err error) {
	err = client.Create(ctx, CustomerURLPath, nil, draft, &result)
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
func (client *Client) CustomerDeleteWithKey(ctx context.Context, key string, version int, dataErasure bool) (result *Customer, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(ctx, strings.Replace("customers/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGetWithKey for type Customer
func (client *Client) CustomerGetWithKey(ctx context.Context, key string) (result *Customer, err error) {
	err = client.Get(ctx, strings.Replace("customers/key={key}", "{key}", key, 1), nil, &result)
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
func (client *Client) CustomerUpdateWithKey(ctx context.Context, input *CustomerUpdateWithKeyInput) (result *Customer, err error) {
	err = client.Update(ctx, strings.Replace("customers/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerDeleteWithID for type Customer
func (client *Client) CustomerDeleteWithID(ctx context.Context, ID string, version int, dataErasure bool) (result *Customer, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(ctx, strings.Replace("customers/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGetWithID for type Customer
func (client *Client) CustomerGetWithID(ctx context.Context, ID string) (result *Customer, err error) {
	err = client.Get(ctx, strings.Replace("customers/{ID}", "{ID}", ID, 1), nil, &result)
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
func (client *Client) CustomerUpdateWithID(ctx context.Context, input *CustomerUpdateWithIDInput) (result *Customer, err error) {
	err = client.Update(ctx, strings.Replace("customers/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
