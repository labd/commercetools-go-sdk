// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyOrdersEditsByIdApplyRequestMethodPost struct {
	body   OrderEditApply
	url    string
	client *Client
	query  url.Values
}

func (r *ByProjectKeyOrdersEditsByIdApplyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

func (rb *ByProjectKeyOrdersEditsByIdApplyRequestMethodPost) Execute() error {
	data, err := serializeInput(rb.body)
	if err != nil {
		return err
	}
	resp, err := rb.client.post(
		context.Background(),
		rb.url,
		rb.query,
		data,
	)

	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 400, 401, 403, 500, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 404:
		return fmt.Errorf("StatusCode %d returend", resp.StatusCode)
	default:
		return fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
