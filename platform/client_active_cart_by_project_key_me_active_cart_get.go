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

type ByProjectKeyMeActiveCartRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeActiveCartRequestMethodGetInput
}

func (r *ByProjectKeyMeActiveCartRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeActiveCartRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyMeActiveCartRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyMeActiveCartRequestMethodGet) Expand(v []string) *ByProjectKeyMeActiveCartRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeActiveCartRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyMeActiveCartRequestMethodGet) WithQueryParams(input ByProjectKeyMeActiveCartRequestMethodGetInput) *ByProjectKeyMeActiveCartRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeActiveCartRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyMeActiveCartRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves the Customer's most recently modified active Cart.
*	Carts with `Merchant` or `Quote` [CartOrigin](ctp:api:type:CartOrigin) are ignored.
*
*	If no active Cart exists, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyMeActiveCartRequestMethodGet) Execute(ctx context.Context) (result *Cart, err error) {
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
		return result, nil
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
