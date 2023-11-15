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

type ByProjectKeyBusinessUnitsRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyBusinessUnitsRequestMethodHeadInput
}

func (r *ByProjectKeyBusinessUnitsRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyBusinessUnitsRequestMethodHeadInput struct {
	Where []string
}

func (input *ByProjectKeyBusinessUnitsRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyBusinessUnitsRequestMethodHead) Where(v []string) *ByProjectKeyBusinessUnitsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyBusinessUnitsRequestMethodHeadInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyBusinessUnitsRequestMethodHead) WithQueryParams(input ByProjectKeyBusinessUnitsRequestMethodHeadInput) *ByProjectKeyBusinessUnitsRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyBusinessUnitsRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyBusinessUnitsRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if a BusinessUnit exists for a given Query Predicate. Returns a `200 OK` status if any BusinessUnits match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyBusinessUnitsRequestMethodHead) Execute(ctx context.Context) error {
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
