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

type ByProjectKeyApiClientsRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyApiClientsRequestMethodHeadInput
}

func (r *ByProjectKeyApiClientsRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyApiClientsRequestMethodHeadInput struct {
	Where []string
}

func (input *ByProjectKeyApiClientsRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyApiClientsRequestMethodHead) Where(v []string) *ByProjectKeyApiClientsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyApiClientsRequestMethodHeadInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyApiClientsRequestMethodHead) WithQueryParams(input ByProjectKeyApiClientsRequestMethodHeadInput) *ByProjectKeyApiClientsRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyApiClientsRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyApiClientsRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if an API Client exists for the provided query predicate. Returns a `200 OK` status if any API Clients match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyApiClientsRequestMethodHead) Execute(ctx context.Context) error {
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
