// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// PaymentURLPath is the commercetools API path.
const PaymentURLPath = "payments"

// PaymentCreate creates a new instance of type Payment
func (client *Client) PaymentCreate(ctx context.Context, draft *PaymentDraft, opts ...RequestOption) (result *Payment, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Create(ctx, PaymentURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentQuery allows querying for type Payment
func (client *Client) PaymentQuery(ctx context.Context, input *QueryInput) (result *PaymentPagedQueryResponse, err error) {
	err = client.Query(ctx, PaymentURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentDeleteWithKey for type Payment
func (client *Client) PaymentDeleteWithKey(ctx context.Context, key string, version int, dataErasure bool, opts ...RequestOption) (result *Payment, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("payments/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentGetWithKey for type Payment
func (client *Client) PaymentGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Payment, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("payments/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentUpdateWithKeyInput is input for function PaymentUpdateWithKey
type PaymentUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []PaymentUpdateAction
}

// PaymentUpdateWithKey for type Payment
func (client *Client) PaymentUpdateWithKey(ctx context.Context, input *PaymentUpdateWithKeyInput, opts ...RequestOption) (result *Payment, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("payments/key={key}", "{key}", input.Key, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentDeleteWithID for type Payment
func (client *Client) PaymentDeleteWithID(ctx context.Context, ID string, version int, dataErasure bool, opts ...RequestOption) (result *Payment, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("payments/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentGetWithID for type Payment
func (client *Client) PaymentGetWithID(ctx context.Context, ID string, opts ...RequestOption) (result *Payment, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("payments/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentUpdateWithIDInput is input for function PaymentUpdateWithID
type PaymentUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []PaymentUpdateAction
}

// PaymentUpdateWithID for type Payment
func (client *Client) PaymentUpdateWithID(ctx context.Context, input *PaymentUpdateWithIDInput, opts ...RequestOption) (result *Payment, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("payments/{ID}", "{ID}", input.ID, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
