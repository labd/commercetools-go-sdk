// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyMeEmailConfirmRequestMethodPost struct {
	url    string
	client *Client
	query  url.Values
}

func (r *ByProjectKeyMeEmailConfirmRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

func (rb *ByProjectKeyMeEmailConfirmRequestMethodPost) Execute(ctx context.Context) error {
	resp, err := rb.client.post(
		ctx,
		rb.url,
		rb.query,
		nil,
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
