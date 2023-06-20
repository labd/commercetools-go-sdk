package connect

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyDeploymentsByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyDeploymentsByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyDeploymentsByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyDeploymentsByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyDeploymentsByIDRequestMethodDelete) Execute(ctx context.Context) (result *Deployment, err error) {
	queryParams := url.Values{}
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
	case 202:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 200:
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
