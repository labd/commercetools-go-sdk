// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// DiscountCodeURLPath is the commercetools API path.
const DiscountCodeURLPath = "discount-codes"

// DiscountCodeCreate creates a new instance of type DiscountCode
func (client *Client) DiscountCodeCreate(ctx context.Context, draft *DiscountCodeDraft, opts ...RequestOption) (result *DiscountCode, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, DiscountCodeURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DiscountCodeQuery allows querying for type DiscountCode
func (client *Client) DiscountCodeQuery(ctx context.Context, input *QueryInput) (result *DiscountCodePagedQueryResponse, err error) {
	err = client.Query(ctx, DiscountCodeURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DiscountCodeDeleteWithID for type DiscountCode
func (client *Client) DiscountCodeDeleteWithID(ctx context.Context, id string, version int, dataErasure bool, opts ...RequestOption) (result *DiscountCode, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("discount-codes/%s", id)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DiscountCodeGetWithID for type DiscountCode
func (client *Client) DiscountCodeGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *DiscountCode, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("discount-codes/%s", id)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DiscountCodeUpdateWithIDInput is input for function DiscountCodeUpdateWithID
type DiscountCodeUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []DiscountCodeUpdateAction
}

// DiscountCodeUpdateWithID for type DiscountCode
func (client *Client) DiscountCodeUpdateWithID(ctx context.Context, input *DiscountCodeUpdateWithIDInput, opts ...RequestOption) (result *DiscountCode, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("discount-codes/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
