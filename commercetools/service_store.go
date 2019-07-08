// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// StoreURLPath is the commercetools API path.
const StoreURLPath = "stores"

func (client *Client) StoreCreate(draft *StoreDraft) (result *Store, err error) {
	err = client.Create(StoreURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) StoreQuery(input *QueryInput) (result *StorePagedQueryResponse, err error) {
	err = client.Query(StoreURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) StoreDeleteWithKey(key string, version int) (result *Store, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("stores/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) StoreGetWithKey(key string) (result *Store, err error) {
	err = client.Get(strings.Replace("stores/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type StoreUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []StoreUpdateAction
}

func (client *Client) StoreUpdateWithKey(input *StoreUpdateWithKeyInput) (result *Store, err error) {
	err = client.Update(strings.Replace("stores/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) StoreDeleteWithId(ID string, version int) (result *Store, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("stores/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) StoreGetWithId(ID string) (result *Store, err error) {
	err = client.Get(strings.Replace("stores/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type StoreUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []StoreUpdateAction
}

func (client *Client) StoreUpdateWithId(input *StoreUpdateWithIdInput) (result *Store, err error) {
	err = client.Update(strings.Replace("stores/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
