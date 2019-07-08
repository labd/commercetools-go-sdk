// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// CustomerGroupURLPath is the commercetools API path.
const CustomerGroupURLPath = "customer-groups"

func (client *Client) CustomerGroupCreate(draft *CustomerGroupDraft) (result *CustomerGroup, err error) {
	err = client.Create(CustomerGroupURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CustomerGroupQuery(input *QueryInput) (result *CustomerGroupPagedQueryResponse, err error) {
	err = client.Query(CustomerGroupURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CustomerGroupDeleteWithKey(key string, version int) (result *CustomerGroup, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("customer-groups/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CustomerGroupGetWithKey(key string) (result *CustomerGroup, err error) {
	err = client.Get(strings.Replace("customer-groups/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type CustomerGroupUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CustomerGroupUpdateAction
}

func (client *Client) CustomerGroupUpdateWithKey(input *CustomerGroupUpdateWithKeyInput) (result *CustomerGroup, err error) {
	err = client.Update(strings.Replace("customer-groups/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CustomerGroupDeleteWithId(ID string, version int) (result *CustomerGroup, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("customer-groups/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CustomerGroupGetWithId(ID string) (result *CustomerGroup, err error) {
	err = client.Get(strings.Replace("customer-groups/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type CustomerGroupUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []CustomerGroupUpdateAction
}

func (client *Client) CustomerGroupUpdateWithId(input *CustomerGroupUpdateWithIdInput) (result *CustomerGroup, err error) {
	err = client.Update(strings.Replace("customer-groups/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
