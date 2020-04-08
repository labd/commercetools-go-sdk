// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// MePaymentsURLPath is the commercetools API path.
const MePaymentsURLPath = "me/payments"

// MePaymentsCreate creates a new instance of type MyPayment
func (client *Client) MePaymentsCreate(draft *MyPaymentDraft) (result *MyPayment, err error) {
	err = client.Create(MePaymentsURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MePaymentsQuery allows querying for type MyPayment
func (client *Client) MePaymentsQuery(input *QueryInput) (result *MyPaymentPagedQueryResponse, err error) {
	err = client.Query(MePaymentsURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MePaymentsDeleteWithID for type MyPayment
func (client *Client) MePaymentsDeleteWithID(ID string, version int) (result *MyPayment, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("me/payments/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MePaymentsGetWithID for type MyPayment
func (client *Client) MePaymentsGetWithID(ID string) (result *MyPayment, err error) {
	err = client.Get(strings.Replace("me/payments/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyPaymentUpdateWithIDInput is input for function MePaymentsUpdateWithID
type MyPaymentUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []MyPaymentUpdateAction
}

// MePaymentsUpdateWithID for type MyPayment
func (client *Client) MePaymentsUpdateWithID(input *MyPaymentUpdateWithIDInput) (result *MyPayment, err error) {
	err = client.Update(strings.Replace("me/payments/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
