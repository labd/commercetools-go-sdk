package checkout

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyRecurringPaymentsByIdRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyRecurringPaymentsByIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyRecurringPaymentsByIdRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyRecurringPaymentsByIdRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves a Recurring Payment with a given `id`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyRecurringPaymentsByIdRequestMethodGet) Execute(ctx context.Context) (result *RecurringPayment, err error) {
	queryParams := url.Values{}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
