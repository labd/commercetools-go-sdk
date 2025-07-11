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

type ByProjectKeyCustomerGroupsRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomerGroupsRequestMethodHeadInput
}

func (r *ByProjectKeyCustomerGroupsRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomerGroupsRequestMethodHeadInput struct {
	Where []string
}

func (input *ByProjectKeyCustomerGroupsRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomerGroupsRequestMethodHead) Where(v []string) *ByProjectKeyCustomerGroupsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomerGroupsRequestMethodHeadInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyCustomerGroupsRequestMethodHead) WithQueryParams(input ByProjectKeyCustomerGroupsRequestMethodHeadInput) *ByProjectKeyCustomerGroupsRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomerGroupsRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyCustomerGroupsRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if one or more CustomerGroups exist for the provided query predicate. Returns a `200 OK` status if any CustomerGroup match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCustomerGroupsRequestMethodHead) Execute(ctx context.Context) error {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.head(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		return nil
	case 404:
		return ErrNotFound
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return result
	}

}
