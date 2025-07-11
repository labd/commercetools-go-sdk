package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Use this method to retrieve a Store-specific Customer's details by using the password reset token during their [password reset process](/../api/customers-overview#customer-password-reset).
*
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet) Execute(ctx context.Context) (result *Customer, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
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
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 503:
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
