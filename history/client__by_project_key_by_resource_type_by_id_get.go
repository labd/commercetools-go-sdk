// Generated file, please do not change!!!
package history

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

type ByProjectKeyByResourceTypeByIdRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyByResourceTypeByIdRequestMethodGetInput
}

func (r *ByProjectKeyByResourceTypeByIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyByResourceTypeByIdRequestMethodGetInput struct {
	DateFrom   *interface{}
	DateTo     *interface{}
	Limit      *int
	Offset     *int
	UserId     *string
	Type       *string
	ClientId   *string
	Source     *string
	Changes    []string
	CustomerId *string
	Expand     *bool
}

func (input *ByProjectKeyByResourceTypeByIdRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.DateFrom != nil {
		values.Add("date.from", fmt.Sprintf("%v", *input.DateFrom))
	}
	if input.DateTo != nil {
		values.Add("date.to", fmt.Sprintf("%v", *input.DateTo))
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.UserId != nil {
		values.Add("userId", fmt.Sprintf("%v", *input.UserId))
	}
	if input.Type != nil {
		values.Add("type", fmt.Sprintf("%v", *input.Type))
	}
	if input.ClientId != nil {
		values.Add("clientId", fmt.Sprintf("%v", *input.ClientId))
	}
	if input.Source != nil {
		values.Add("source", fmt.Sprintf("%v", *input.Source))
	}
	for _, v := range input.Changes {
		values.Add("changes", fmt.Sprintf("%v", v))
	}
	if input.CustomerId != nil {
		values.Add("customerId", fmt.Sprintf("%v", *input.CustomerId))
	}
	if input.Expand != nil {
		if *input.Expand == true {
			values.Add("expand", "true")
		} else {
			values.Add("expand", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyByResourceTypeByIdRequestMethodGet) WithQueryParams(input *ByProjectKeyByResourceTypeByIdRequestMethodGetInput) *ByProjectKeyByResourceTypeByIdRequestMethodGet {
	rb.query = input.Values()
	return rb
}
func (rb *ByProjectKeyByResourceTypeByIdRequestMethodGet) Execute(ctx context.Context) (result *RecordPagedQueryResponse, err error) {
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
	case 400, 401, 403, 500, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
