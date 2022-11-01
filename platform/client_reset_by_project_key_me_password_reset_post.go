package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyMePasswordResetRequestMethodPost struct {
	body    MyCustomerResetPassword
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyMePasswordResetRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyMePasswordResetRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyMePasswordResetRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	This is the last step in the [password reset process of a Customer](/../api/projects/customers#password-reset-of-customer).
*
*	Resetting a password of the Customer produces the [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=true`.
*
 */
func (rb *ByProjectKeyMePasswordResetRequestMethodPost) Execute(ctx context.Context) (result *Customer, err error) {
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
		return result, nil
	case 400, 401, 403, 500, 502, 503:
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
