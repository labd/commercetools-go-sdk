package connect

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestMethodGet) Execute(ctx context.Context) (result *DeploymentRevision, err error) {
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
