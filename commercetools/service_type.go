// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// TypeURLPath is the commercetools API path.
const TypeURLPath = "types"

func (client *Client) TypeCreate(draft *TypeDraft) (result *Type, err error) {
	err = client.Create(TypeURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) TypeQuery(input *QueryInput) (result *TypePagedQueryResponse, err error) {
	err = client.Query(TypeURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) TypeDeleteWithKey(key string, version int) (result *Type, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("types/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) TypeGetWithKey(key string) (result *Type, err error) {
	err = client.Get(strings.Replace("types/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type TypeUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []TypeUpdateAction
}

func (client *Client) TypeUpdateWithKey(input *TypeUpdateWithKeyInput) (result *Type, err error) {
	err = client.Update(strings.Replace("types/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) TypeDeleteWithId(ID string, version int) (result *Type, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("types/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) TypeGetWithId(ID string) (result *Type, err error) {
	err = client.Get(strings.Replace("types/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type TypeUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []TypeUpdateAction
}

func (client *Client) TypeUpdateWithId(input *TypeUpdateWithIdInput) (result *Type, err error) {
	err = client.Update(strings.Replace("types/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
