// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// ZoneURLPath is the commercetools API path.
const ZoneURLPath = "zones"

func (client *Client) ZoneCreate(draft *ZoneDraft) (result *Zone, err error) {
	err = client.Create(ZoneURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ZoneQuery(input *QueryInput) (result *ZonePagedQueryResponse, err error) {
	err = client.Query(ZoneURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ZoneDeleteWithKey(key string, version int) (result *Zone, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("zones/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ZoneGetWithKey(key string) (result *Zone, err error) {
	err = client.Get(strings.Replace("zones/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ZoneUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ZoneUpdateAction
}

func (client *Client) ZoneUpdateWithKey(input *ZoneUpdateWithKeyInput) (result *Zone, err error) {
	err = client.Update(strings.Replace("zones/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ZoneDeleteWithID(ID string, version int) (result *Zone, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("zones/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ZoneGetWithID(ID string) (result *Zone, err error) {
	err = client.Get(strings.Replace("zones/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ZoneUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ZoneUpdateAction
}

func (client *Client) ZoneUpdateWithID(input *ZoneUpdateWithIDInput) (result *Zone, err error) {
	err = client.Update(strings.Replace("zones/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
