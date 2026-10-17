package checkout

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyRecurringPaymentsKeyByKeyRequestMethodPost struct {
	body    RecurringPaymentUpdateActions
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyRecurringPaymentsKeyByKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyRecurringPaymentsKeyByKeyRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyRecurringPaymentsKeyByKeyRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Updates a Recurring Payment with a given `key`. Specific Error Codes:
*	- [InvalidInput](ctp:checkout:type:InvalidInputError)
*	- [ReferencedResourceNotFound](ctp:checkout:type:ReferencedResourceNotFoundError)
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError)
*
 */
func (rb *ByProjectKeyRecurringPaymentsKeyByKeyRequestMethodPost) Execute(ctx context.Context) (result *RecurringPayment, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	queryParams := url.Values{}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
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
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
