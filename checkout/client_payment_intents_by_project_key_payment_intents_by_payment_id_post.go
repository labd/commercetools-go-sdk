package checkout

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyPaymentIntentsByPaymentIdRequestMethodPost struct {
	body    Payment
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyPaymentIntentsByPaymentIdRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyPaymentIntentsByPaymentIdRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyPaymentIntentsByPaymentIdRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Specific Error Codes:
*	- [MultipleActionsNotAllowed](ctp:checkout:type:MultipleActionsNotAllowedError)
*	- [RequiredField](ctp:checkout:type:RequiredFieldError)
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*
 */
func (rb *ByProjectKeyPaymentIntentsByPaymentIdRequestMethodPost) Execute(ctx context.Context) (result *interface{}, err error) {
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
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
