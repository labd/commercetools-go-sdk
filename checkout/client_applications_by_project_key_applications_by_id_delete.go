package checkout

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyApplicationsByIdRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyApplicationsByIdRequestMethodDeleteInput
}

func (r *ByProjectKeyApplicationsByIdRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyApplicationsByIdRequestMethodDeleteInput struct {
	Version int
}

func (input *ByProjectKeyApplicationsByIdRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	return values
}

func (rb *ByProjectKeyApplicationsByIdRequestMethodDelete) Version(v int) *ByProjectKeyApplicationsByIdRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyApplicationsByIdRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyApplicationsByIdRequestMethodDelete) WithQueryParams(input ByProjectKeyApplicationsByIdRequestMethodDeleteInput) *ByProjectKeyApplicationsByIdRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyApplicationsByIdRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyApplicationsByIdRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyApplicationsByIdRequestMethodDelete) Execute(ctx context.Context) (result *Application, err error) {
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
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
