// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// ReviewURLPath is the commercetools API path.
const ReviewURLPath = "reviews"

// ReviewCreate creates a new instance of type Review
func (client *Client) ReviewCreate(ctx context.Context, draft *ReviewDraft) (result *Review, err error) {
	err = client.Create(ctx, ReviewURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewQuery allows querying for type Review
func (client *Client) ReviewQuery(ctx context.Context, input *QueryInput) (result *ReviewPagedQueryResponse, err error) {
	err = client.Query(ctx, ReviewURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewDeleteWithKey for type Review
func (client *Client) ReviewDeleteWithKey(ctx context.Context, key string, version int, dataErasure bool) (result *Review, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(ctx, strings.Replace("reviews/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewGetWithKey for type Review
func (client *Client) ReviewGetWithKey(ctx context.Context, key string) (result *Review, err error) {
	err = client.Get(ctx, strings.Replace("reviews/key={key}", "{key}", key, 1), nil, &result)
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

// ReviewUpdateWithKey for type Review
func (client *Client) ReviewUpdateWithKey(ctx context.Context, input *ReviewUpdateWithKeyInput) (result *Review, err error) {
	err = client.Update(ctx, strings.Replace("reviews/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewDeleteWithID for type Review
func (client *Client) ReviewDeleteWithID(ctx context.Context, ID string, version int, dataErasure bool) (result *Review, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(ctx, strings.Replace("reviews/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewGetWithID for type Review
func (client *Client) ReviewGetWithID(ctx context.Context, ID string) (result *Review, err error) {
	err = client.Get(ctx, strings.Replace("reviews/{ID}", "{ID}", ID, 1), nil, &result)
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

// ReviewUpdateWithID for type Review
func (client *Client) ReviewUpdateWithID(ctx context.Context, input *ReviewUpdateWithIDInput) (result *Review, err error) {
	err = client.Update(ctx, strings.Replace("reviews/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
