// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// PaymentCreate creates a new instance of type Payment
func (client *Client) PaymentCreate(ctx context.Context, draft *PaymentDraft, opts ...RequestOption) (result *Payment, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "payments"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentQuery allows querying for type Payment
func (client *Client) PaymentQuery(ctx context.Context, input *QueryInput) (result *PaymentPagedQueryResponse, err error) {
	endpoint := "payments"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentDeleteWithID for type Payment
func (client *Client) PaymentDeleteWithID(ctx context.Context, id string, version int, dataErasure bool, opts ...RequestOption) (result *Payment, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("payments/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
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
	endpoint := fmt.Sprintf("payments/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentGetWithID for type Payment
func (client *Client) PaymentGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Payment, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("payments/%s", id)
	err = client.get(ctx, endpoint, params, &result)
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
	endpoint := fmt.Sprintf("payments/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
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

func (input *PaymentUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// PaymentUpdateWithID for type Payment
func (client *Client) PaymentUpdateWithID(ctx context.Context, input *PaymentUpdateWithIDInput, opts ...RequestOption) (result *Payment, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("payments/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
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

func (input *PaymentUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// PaymentUpdateWithKey for type Payment
func (client *Client) PaymentUpdateWithKey(ctx context.Context, input *PaymentUpdateWithKeyInput, opts ...RequestOption) (result *Payment, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("payments/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
