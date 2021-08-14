// Generated file, please do not change!!!
package history

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyRequestMethodGetInput
}

func (r *ByProjectKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyRequestMethodGetInput struct {
	ResourceType *[]ChangeHistoryResourceType
	DateFrom     *float64
	DateTo       *float64
	Limit        *int
	Offset       *int
	UserId       *string
	Type         *string
	ClientId     *string
	ResourceId   *string
	Source       *string
	Changes      *[]string
	CustomerId   *string
	Expand       *bool
}

func (input *ByProjectKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.ResourceType != nil {
		for _, v := range *input.ResourceType {
			values.Add("resourceType", v)
		}
	}
	if input.DateFrom != nil {
		values.Add("date.from", *input.DateFrom)
	}
	if input.DateTo != nil {
		values.Add("date.to", *input.DateTo)
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.UserId != nil {
		values.Add("userId", *input.UserId)
	}
	if input.Type != nil {
		values.Add("type", *input.Type)
	}
	if input.ClientId != nil {
		values.Add("clientId", *input.ClientId)
	}
	if input.ResourceId != nil {
		values.Add("resourceId", *input.ResourceId)
	}
	if input.Source != nil {
		values.Add("source", *input.Source)
	}
	if input.Changes != nil {
		for _, v := range *input.Changes {
			values.Add("changes", v)
		}
	}
	if input.CustomerId != nil {
		values.Add("customerId", *input.CustomerId)
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

func (rb *ByProjectKeyRequestMethodGet) WithQueryParams(input *ByProjectKeyRequestMethodGetInput) *ByProjectKeyRequestMethodGet {
	rb.query = input.Values()
	return rb
}
func (rb *ByProjectKeyRequestMethodGet) Execute() (result *RecordPagedQueryResponse, err error) {
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
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
