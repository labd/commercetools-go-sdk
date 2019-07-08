// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// DiscountCodeURLPath is the commercetools API path.
const DiscountCodeURLPath = "discount-codes"

func (client *Client) DiscountCodeCreate(draft *DiscountCodeDraft) (result *DiscountCode, err error) {
	err = client.Create(DiscountCodeURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) DiscountCodeQuery(input *QueryInput) (result *DiscountCodePagedQueryResponse, err error) {
	err = client.Query(DiscountCodeURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) DiscountCodeDeleteWithId(ID string, version int, dataErasure bool) (result *DiscountCode, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("discount-codes/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) DiscountCodeGetWithId(ID string) (result *DiscountCode, err error) {
	err = client.Get(strings.Replace("discount-codes/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type DiscountCodeUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []DiscountCodeUpdateAction
}

func (client *Client) DiscountCodeUpdateWithId(input *DiscountCodeUpdateWithIdInput) (result *DiscountCode, err error) {
	err = client.Update(strings.Replace("discount-codes/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
