package checkout

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyRecurringPaymentsRequestMethodPost struct {
	body    RecurringPaymentDraft
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyRecurringPaymentsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyRecurringPaymentsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyRecurringPaymentsRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates a new [RecurringPayment](ctp:checkout:type:RecurringPayment). Specific Error Codes:
*	- [InvalidInput](ctp:checkout:type:InvalidInputError)
*	- [ReferencedResourceNotFound](ctp:checkout:type:ReferencedResourceNotFoundError)
*	- [RequiredField](ctp:checkout:type:RequiredFieldError)
*
 */
func (rb *ByProjectKeyRecurringPaymentsRequestMethodPost) Execute(ctx context.Context) (result *RecurringPayment, err error) {
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
	case 201:
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
