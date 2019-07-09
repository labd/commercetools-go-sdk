// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// ExtensionURLPath is the commercetools API path.
const ExtensionURLPath = "extensions"

func (client *Client) ExtensionCreate(draft *ExtensionDraft) (result *Extension, err error) {
	err = client.Create(ExtensionURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ExtensionQuery(input *QueryInput) (result *ExtensionPagedQueryResponse, err error) {
	err = client.Query(ExtensionURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ExtensionDeleteWithKey(key string, version int) (result *Extension, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("extensions/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ExtensionGetWithKey(key string) (result *Extension, err error) {
	err = client.Get(strings.Replace("extensions/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ExtensionUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ExtensionUpdateAction
}

func (client *Client) ExtensionUpdateWithKey(input *ExtensionUpdateWithKeyInput) (result *Extension, err error) {
	err = client.Update(strings.Replace("extensions/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ExtensionDeleteWithID(ID string, version int) (result *Extension, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("extensions/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ExtensionGetWithID(ID string) (result *Extension, err error) {
	err = client.Get(strings.Replace("extensions/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ExtensionUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ExtensionUpdateAction
}

func (client *Client) ExtensionUpdateWithID(input *ExtensionUpdateWithIDInput) (result *Extension, err error) {
	err = client.Update(strings.Replace("extensions/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
