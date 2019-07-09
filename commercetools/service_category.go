// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// CategoryURLPath is the commercetools API path.
const CategoryURLPath = "categories"

func (client *Client) CategoryCreate(draft *CategoryDraft) (result *Category, err error) {
	err = client.Create(CategoryURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CategoryQuery(input *QueryInput) (result *CategoryPagedQueryResponse, err error) {
	err = client.Query(CategoryURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CategoryDeleteWithKey(key string, version int) (result *Category, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("categories/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CategoryGetWithKey(key string) (result *Category, err error) {
	err = client.Get(strings.Replace("categories/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type CategoryUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CategoryUpdateAction
}

func (client *Client) CategoryUpdateWithKey(input *CategoryUpdateWithKeyInput) (result *Category, err error) {
	err = client.Update(strings.Replace("categories/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CategoryDeleteWithID(ID string, version int) (result *Category, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("categories/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CategoryGetWithID(ID string) (result *Category, err error) {
	err = client.Get(strings.Replace("categories/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type CategoryUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CategoryUpdateAction
}

func (client *Client) CategoryUpdateWithID(input *CategoryUpdateWithIDInput) (result *Category, err error) {
	err = client.Update(strings.Replace("categories/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
