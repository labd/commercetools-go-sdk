// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// CustomerURLPath is the commercetools API path.
const CustomerURLPath = "customers"

func (client *Client) CustomerCreate(draft *CustomerDraft) (result *Customer, err error) {
	err = client.Create(CustomerURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CustomerQuery(input *QueryInput) (result *CustomerPagedQueryResponse, err error) {
	err = client.Query(CustomerURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CustomerDeleteWithId(ID string, version int, dataErasure bool) (result *Customer, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("customers/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) CustomerGetWithId(ID string) (result *Customer, err error) {
	err = client.Get(strings.Replace("customers/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type CustomerUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []CustomerUpdateAction
}

func (client *Client) CustomerUpdateWithId(input *CustomerUpdateWithIdInput) (result *Customer, err error) {
	err = client.Update(strings.Replace("customers/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
