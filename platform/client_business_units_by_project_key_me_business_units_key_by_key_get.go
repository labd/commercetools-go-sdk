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

type ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGetInput
}

func (r *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet) Expand(v []string) *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet) WithQueryParams(input ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGetInput) *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Returns a Business Unit for a given `key`. Returns a `200 OK` status if the Business Unit exists and the Customer has access to it, or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet) Execute(ctx context.Context) (result *BusinessUnit, err error) {
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
	case 404:
		return nil, ErrNotFound
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
