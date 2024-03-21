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

type ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGetInput
}

func (r *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGet) Expand(v []string) *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGet) WithQueryParams(input ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGetInput) *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGet) Execute(ctx context.Context) (result *BusinessUnit, err error) {
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
