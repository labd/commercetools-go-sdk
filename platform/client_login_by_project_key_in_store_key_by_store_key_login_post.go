package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyLoginRequestMethodPost struct {
	body    CustomerSignin
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyInStoreKeyByStoreKeyLoginRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyLoginRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyLoginRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Authenticates a Customer associated with a Store. For more information, see [Global versus Store-specific Customers](/../api/customers-overview#global-versus-store-specific-customers).
*
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns an [InvalidCredentials](ctp:api:type:InvalidCredentialsError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyLoginRequestMethodPost) Execute(ctx context.Context) (result *CustomerSignInResult, err error) {
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
