package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyMeRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeRequestMethodDeleteInput
}

func (r *ByProjectKeyMeRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeRequestMethodDeleteInput struct {
	Version int
}

func (input *ByProjectKeyMeRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	return values
}

func (rb *ByProjectKeyMeRequestMethodDelete) Version(v int) *ByProjectKeyMeRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyMeRequestMethodDelete) WithQueryParams(input ByProjectKeyMeRequestMethodDeleteInput) *ByProjectKeyMeRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyMeRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMeRequestMethodDelete) Execute(ctx context.Context) (result *Customer, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
	case 409, 400, 401, 403, 500, 502, 503:
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
