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

type ByProjectKeyPaymentMethodsRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyPaymentMethodsRequestMethodHeadInput
}

func (r *ByProjectKeyPaymentMethodsRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyPaymentMethodsRequestMethodHeadInput struct {
	Where []string
}

func (input *ByProjectKeyPaymentMethodsRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyPaymentMethodsRequestMethodHead) Where(v []string) *ByProjectKeyPaymentMethodsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentMethodsRequestMethodHeadInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyPaymentMethodsRequestMethodHead) WithQueryParams(input ByProjectKeyPaymentMethodsRequestMethodHeadInput) *ByProjectKeyPaymentMethodsRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyPaymentMethodsRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyPaymentMethodsRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if one or more PaymentMethods exist for the provided query predicate.
*	Returns a `200 OK` status if any PaymentMethods match the query predicate; otherwise, returns a [Not Found](/../api/errors#404-not-found).
*
 */
func (rb *ByProjectKeyPaymentMethodsRequestMethodHead) Execute(ctx context.Context) error {
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
