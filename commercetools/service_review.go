// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// ReviewURLPath is the commercetools API path.
const ReviewURLPath = "reviews"

func (client *Client) ReviewCreate(draft *ReviewDraft) (result *Review, err error) {
	err = client.Create(ReviewURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ReviewQuery(input *QueryInput) (result *ReviewPagedQueryResponse, err error) {
	err = client.Query(ReviewURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ReviewDeleteWithKey(key string, version int, dataErasure bool) (result *Review, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("reviews/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ReviewGetWithKey(key string) (result *Review, err error) {
	err = client.Get(strings.Replace("reviews/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ReviewUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ReviewUpdateAction
}

func (client *Client) ReviewUpdateWithKey(input *ReviewUpdateWithKeyInput) (result *Review, err error) {
	err = client.Update(strings.Replace("reviews/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ReviewDeleteWithID(ID string, version int, dataErasure bool) (result *Review, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("reviews/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ReviewGetWithID(ID string) (result *Review, err error) {
	err = client.Get(strings.Replace("reviews/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ReviewUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ReviewUpdateAction
}

func (client *Client) ReviewUpdateWithID(input *ReviewUpdateWithIDInput) (result *Review, err error) {
	err = client.Update(strings.Replace("reviews/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
