package connect

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ConnectorsDraftsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ConnectorsDraftsRequestMethodGetInput
}

func (r *ConnectorsDraftsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ConnectorsDraftsRequestMethodGetInput struct {
	Limit  *int
	Offset *int
	Sort   []string
}

func (input *ConnectorsDraftsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ConnectorsDraftsRequestMethodGet) Limit(v int) *ConnectorsDraftsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ConnectorsDraftsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ConnectorsDraftsRequestMethodGet) Offset(v int) *ConnectorsDraftsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ConnectorsDraftsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ConnectorsDraftsRequestMethodGet) Sort(v []string) *ConnectorsDraftsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ConnectorsDraftsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ConnectorsDraftsRequestMethodGet) WithQueryParams(input ConnectorsDraftsRequestMethodGetInput) *ConnectorsDraftsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ConnectorsDraftsRequestMethodGet) WithHeaders(headers http.Header) *ConnectorsDraftsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ConnectorsDraftsRequestMethodGet) Execute(ctx context.Context) (result *ConnectorStagedPagedQueryResponse, err error) {
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
