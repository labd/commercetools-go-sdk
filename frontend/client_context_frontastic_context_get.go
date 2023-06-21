package frontend

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FrontasticContextRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *FrontasticContextRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *FrontasticContextRequestMethodGet) WithHeaders(headers http.Header) *FrontasticContextRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Returns information about the project locales setup and the environment in which the requested host acts in. \
*
*	Headers&#58;
*
*	- `Accept` - `application/json` - Required
*
 */
func (rb *FrontasticContextRequestMethodGet) Execute(ctx context.Context) (result *ProjectContext, err error) {
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
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
