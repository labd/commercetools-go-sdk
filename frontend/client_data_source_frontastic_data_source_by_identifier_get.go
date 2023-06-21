package frontend

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FrontasticDataSourceByIdentifierRequestMethodGet struct {
	body    interface{}
	url     string
	client  *Client
	headers http.Header
}

func (r *FrontasticDataSourceByIdentifierRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *FrontasticDataSourceByIdentifierRequestMethodGet) WithHeaders(headers http.Header) *FrontasticDataSourceByIdentifierRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Returns data and preview from the data source.
 */
func (rb *FrontasticDataSourceByIdentifierRequestMethodGet) Execute(ctx context.Context) (result *DataSourceResponse, err error) {
	queryParams := url.Values{}
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
	case 415:
		return nil, nil
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
