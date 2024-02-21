package frontend

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FrontasticPreviewRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *FrontasticPreviewRequestMethodGetInput
}

func (r *FrontasticPreviewRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type FrontasticPreviewRequestMethodGetInput struct {
	Locale    string
	PreviewId string
}

func (input *FrontasticPreviewRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	values.Add("locale", fmt.Sprintf("%v", input.Locale))
	values.Add("previewId", fmt.Sprintf("%v", input.PreviewId))
	return values
}

func (rb *FrontasticPreviewRequestMethodGet) Locale(v string) *FrontasticPreviewRequestMethodGet {
	if rb.params == nil {
		rb.params = &FrontasticPreviewRequestMethodGetInput{}
	}
	rb.params.Locale = v
	return rb
}

func (rb *FrontasticPreviewRequestMethodGet) PreviewId(v string) *FrontasticPreviewRequestMethodGet {
	if rb.params == nil {
		rb.params = &FrontasticPreviewRequestMethodGetInput{}
	}
	rb.params.PreviewId = v
	return rb
}

func (rb *FrontasticPreviewRequestMethodGet) WithQueryParams(input FrontasticPreviewRequestMethodGetInput) *FrontasticPreviewRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *FrontasticPreviewRequestMethodGet) WithHeaders(headers http.Header) *FrontasticPreviewRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Returns the page structure and data for a specific [preview](/../frontend-studio/using-the-page-builder) by its `previewId`. \
*
*	Headers&#58;
*
*	- `Accept` - `application/json` - Required
*
 */
func (rb *FrontasticPreviewRequestMethodGet) Execute(ctx context.Context) (result *PagePreviewDataResponse, err error) {
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
	case 404:
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
