// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

type ByProjectKeyProductProjectionsSuggestRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyProductProjectionsSuggestRequestMethodGetInput
}

func (r *ByProjectKeyProductProjectionsSuggestRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyProductProjectionsSuggestRequestMethodGetInput struct {
	Fuzzy          *bool
	Staged         *bool
	SearchKeywords map[string][]string
	Sort           *[]string
	Limit          *int
	Offset         *int
	WithTotal      *bool
}

func (input *ByProjectKeyProductProjectionsSuggestRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Fuzzy != nil {
		if *input.Fuzzy == true {
			values.Add("fuzzy", "true")
		} else {
			values.Add("fuzzy", "false")
		}
	}
	if input.Staged != nil {
		if *input.Staged == true {
			values.Add("staged", "true")
		} else {
			values.Add("staged", "false")
		}
	}

	if input.Sort != nil {
		for _, v := range *input.Sort {
			values.Add("sort", v)
		}
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal == true {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyProductProjectionsSuggestRequestMethodGet) WithQueryParams(input *ByProjectKeyProductProjectionsSuggestRequestMethodGetInput) *ByProjectKeyProductProjectionsSuggestRequestMethodGet {
	rb.query = input.Values()
	return rb
}
func (rb *ByProjectKeyProductProjectionsSuggestRequestMethodGet) Execute() (result *interface{}, err error) {
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
	case 400, 401, 403, 500, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, fmt.Errorf("StatusCode %d returend", resp.StatusCode)
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
