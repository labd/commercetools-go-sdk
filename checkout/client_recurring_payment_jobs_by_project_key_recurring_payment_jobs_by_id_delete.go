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

type ByProjectKeyRecurringPaymentJobsByIdRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyRecurringPaymentJobsByIdRequestMethodDeleteInput
}

func (r *ByProjectKeyRecurringPaymentJobsByIdRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyRecurringPaymentJobsByIdRequestMethodDeleteInput struct {
	Version int
}

func (input *ByProjectKeyRecurringPaymentJobsByIdRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	return values
}

func (rb *ByProjectKeyRecurringPaymentJobsByIdRequestMethodDelete) Version(v int) *ByProjectKeyRecurringPaymentJobsByIdRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentJobsByIdRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentJobsByIdRequestMethodDelete) WithQueryParams(input ByProjectKeyRecurringPaymentJobsByIdRequestMethodDeleteInput) *ByProjectKeyRecurringPaymentJobsByIdRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyRecurringPaymentJobsByIdRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyRecurringPaymentJobsByIdRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	Deletes a Recurring Payment Job with a given `id`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError)
*
 */
func (rb *ByProjectKeyRecurringPaymentJobsByIdRequestMethodDelete) Execute(ctx context.Context) (result *RecurringPaymentJob, err error) {
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
