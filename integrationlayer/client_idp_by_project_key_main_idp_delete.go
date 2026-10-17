package integrationlayer

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyMainIdpRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyMainIdpRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyMainIdpRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyMainIdpRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMainIdpRequestMethodDelete) Execute(ctx context.Context) error {
	queryParams := url.Values{}
	resp, err := rb.client.delete(
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
	case 204:
		return nil
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 404:
		return ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return result
	}

}
