// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ReviewCreate creates a new instance of type Review
func (client *Client) ReviewCreate(ctx context.Context, draft *ReviewDraft, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "reviews"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewQuery allows querying for type Review
func (client *Client) ReviewQuery(ctx context.Context, input *QueryInput) (result *ReviewPagedQueryResponse, err error) {
	endpoint := "reviews"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewDeleteWithID for type Review
func (client *Client) ReviewDeleteWithID(ctx context.Context, id string, version int, dataErasure bool, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("reviews/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewDeleteWithKey for type Review
func (client *Client) ReviewDeleteWithKey(ctx context.Context, key string, version int, dataErasure bool, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("reviews/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewGetWithID for type Review
func (client *Client) ReviewGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("reviews/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewGetWithKey for type Review
func (client *Client) ReviewGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("reviews/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewUpdateWithIDInput is input for function ReviewUpdateWithID
type ReviewUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ReviewUpdateAction
}

func (input *ReviewUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ReviewUpdateWithID for type Review
func (client *Client) ReviewUpdateWithID(ctx context.Context, input *ReviewUpdateWithIDInput, opts ...RequestOption) (result *Review, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("reviews/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewUpdateWithKeyInput is input for function ReviewUpdateWithKey
type ReviewUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ReviewUpdateAction
}

func (input *ReviewUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ReviewUpdateWithKey for type Review
func (client *Client) ReviewUpdateWithKey(ctx context.Context, input *ReviewUpdateWithKeyInput, opts ...RequestOption) (result *Review, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("reviews/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
