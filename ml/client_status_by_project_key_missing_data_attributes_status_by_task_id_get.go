// Generated file, please do not change!!!
package ml

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyMissingDataAttributesStatusByTaskIdRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
}

func (r *ByProjectKeyMissingDataAttributesStatusByTaskIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

func (rb *ByProjectKeyMissingDataAttributesStatusByTaskIdRequestMethodGet) Execute() (result *MissingDataTaskStatus, err error) {
	resp, err := rb.client.get(
		context.Background(),
		rb.url,
		rb.query,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
