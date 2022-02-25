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

type ByProjectKeyOrdersEditsByIDApplyRequestMethodPost struct {
	body    OrderEditApply
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyOrdersEditsByIDApplyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyOrdersEditsByIDApplyRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyOrdersEditsByIDApplyRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyOrdersEditsByIDApplyRequestMethodPost) Execute(ctx context.Context) error {
	data, err := serializeInput(rb.body)
	if err != nil {
		return err
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
	case 404:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return result
	default:
		return fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
