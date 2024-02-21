package frontend

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ApiBuildUploadRequestMethodPost struct {
	body    BuildUpload
	url     string
	client  *Client
	headers http.Header
}

func (r *ApiBuildUploadRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ApiBuildUploadRequestMethodPost) WithHeaders(headers http.Header) *ApiBuildUploadRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Uploads your backend builds from your custom CI environment to the Studio.
 */
func (rb *ApiBuildUploadRequestMethodPost) Execute(ctx context.Context) (result *interface{}, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	queryParams := url.Values{}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
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
