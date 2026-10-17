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

type ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGet) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGet) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves a [Category](ctp:api:type:Category) by its `id` if it is assigned to the specified [Store](ctp:api:type:Store) or is global.
*
*	For global Categories, use the [Get Category by ID](ctp:api:endpoint:/{projectKey}/categories/{id}:GET) endpoint.
*
*	If the Category does not exist in the Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGet) Execute(ctx context.Context) (result *Category, err error) {
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
