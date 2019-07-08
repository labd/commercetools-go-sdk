// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// PaymentURLPath is the commercetools API path.
const PaymentURLPath = "payments"

func (client *Client) PaymentCreate(draft *PaymentDraft) (result *Payment, err error) {
	err = client.Create(PaymentURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) PaymentQuery(input *QueryInput) (result *PaymentPagedQueryResponse, err error) {
	err = client.Query(PaymentURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) PaymentDeleteWithKey(key string, version int, dataErasure bool) (result *Payment, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("payments/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) PaymentGetWithKey(key string) (result *Payment, err error) {
	err = client.Get(strings.Replace("payments/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type PaymentUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []PaymentUpdateAction
}

func (client *Client) PaymentUpdateWithKey(input *PaymentUpdateWithKeyInput) (result *Payment, err error) {
	err = client.Update(strings.Replace("payments/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) PaymentDeleteWithId(ID string, version int, dataErasure bool) (result *Payment, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("payments/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) PaymentGetWithId(ID string) (result *Payment, err error) {
	err = client.Get(strings.Replace("payments/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type PaymentUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []PaymentUpdateAction
}

func (client *Client) PaymentUpdateWithId(input *PaymentUpdateWithIdInput) (result *Payment, err error) {
	err = client.Update(strings.Replace("payments/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
