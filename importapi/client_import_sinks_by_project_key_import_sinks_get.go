// Generated file, please do not change!!!
package importapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

type ByProjectKeyImportSinksRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyImportSinksRequestMethodGetInput
}

func (r *ByProjectKeyImportSinksRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyImportSinksRequestMethodGetInput struct {
	Limit  *float64
	Offset *float64
	Sort   []string
}

func (input *ByProjectKeyImportSinksRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Limit != nil {
		values.Add("limit", fmt.Sprintf("%f", *input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", fmt.Sprintf("%f", *input.Offset))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyImportSinksRequestMethodGet) WithQueryParams(input *ByProjectKeyImportSinksRequestMethodGetInput) *ByProjectKeyImportSinksRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	Retrieves all import sinks of a project key.
 */
func (rb *ByProjectKeyImportSinksRequestMethodGet) Execute(ctx context.Context) (result *ImportSinkPagedResponse, err error) {
	resp, err := rb.client.get(
		ctx,
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
