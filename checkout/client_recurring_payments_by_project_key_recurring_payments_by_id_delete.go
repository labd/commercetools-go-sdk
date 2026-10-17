package checkout

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyRecurringPaymentsByIdRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyRecurringPaymentsByIdRequestMethodDeleteInput
}

func (r *ByProjectKeyRecurringPaymentsByIdRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyRecurringPaymentsByIdRequestMethodDeleteInput struct {
	Version int
}

func (input *ByProjectKeyRecurringPaymentsByIdRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	return values
}

func (rb *ByProjectKeyRecurringPaymentsByIdRequestMethodDelete) Version(v int) *ByProjectKeyRecurringPaymentsByIdRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentsByIdRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentsByIdRequestMethodDelete) WithQueryParams(input ByProjectKeyRecurringPaymentsByIdRequestMethodDeleteInput) *ByProjectKeyRecurringPaymentsByIdRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyRecurringPaymentsByIdRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyRecurringPaymentsByIdRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	Deletes a Recurring Payment with a given `id`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError)
*
 */
func (rb *ByProjectKeyRecurringPaymentsByIdRequestMethodDelete) Execute(ctx context.Context) (result *RecurringPayment, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
