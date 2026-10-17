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

type ByProjectKeyApplicationsKeyByKeyRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyApplicationsKeyByKeyRequestMethodDeleteInput
}

func (r *ByProjectKeyApplicationsKeyByKeyRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyApplicationsKeyByKeyRequestMethodDeleteInput struct {
	Version int
}

func (input *ByProjectKeyApplicationsKeyByKeyRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	return values
}

func (rb *ByProjectKeyApplicationsKeyByKeyRequestMethodDelete) Version(v int) *ByProjectKeyApplicationsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyApplicationsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyApplicationsKeyByKeyRequestMethodDelete) WithQueryParams(input ByProjectKeyApplicationsKeyByKeyRequestMethodDeleteInput) *ByProjectKeyApplicationsKeyByKeyRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyApplicationsKeyByKeyRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyApplicationsKeyByKeyRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyApplicationsKeyByKeyRequestMethodDelete) Execute(ctx context.Context) (result *Application, err error) {
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
