package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyMeEmailConfirmRequestMethodPost struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyMeEmailConfirmRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyMeEmailConfirmRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyMeEmailConfirmRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMeEmailConfirmRequestMethodPost) Execute(ctx context.Context) error {
	queryParams := url.Values{}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
	case 400, 401, 403, 500, 502, 503:
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
