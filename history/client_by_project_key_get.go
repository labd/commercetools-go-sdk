package history

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyRequestMethodGetInput
}

func (r *ByProjectKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyRequestMethodGetInput struct {
	ResourceTypes                   []ChangeHistoryResourceType
	DateFrom                        *interface{}
	DateTo                          *interface{}
	Limit                           *int
	Offset                          *int
	UserId                          *string
	ClientId                        *string
	CustomerId                      *string
	AssociateId                     *string
	BusinessUnit                    *string
	Type                            *string
	ResourceId                      *string
	ResourceKey                     *string
	Source                          *string
	Changes                         []string
	Stores                          []string
	ExcludePlatformInitiatedChanges []PlatformInitiatedChange
	Expand                          *bool
}

func (input *ByProjectKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.ResourceTypes {
		values.Add("resourceTypes", fmt.Sprintf("%v", v))
	}
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
	if input.ClientId != nil {
		values.Add("clientId", fmt.Sprintf("%v", *input.ClientId))
	}
	if input.CustomerId != nil {
		values.Add("customerId", fmt.Sprintf("%v", *input.CustomerId))
	}
	if input.AssociateId != nil {
		values.Add("associateId", fmt.Sprintf("%v", *input.AssociateId))
	}
	if input.BusinessUnit != nil {
		values.Add("businessUnit", fmt.Sprintf("%v", *input.BusinessUnit))
	}
	if input.Type != nil {
		values.Add("type", fmt.Sprintf("%v", *input.Type))
	}
	if input.ResourceId != nil {
		values.Add("resourceId", fmt.Sprintf("%v", *input.ResourceId))
	}
	if input.ResourceKey != nil {
		values.Add("resourceKey", fmt.Sprintf("%v", *input.ResourceKey))
	}
	if input.Source != nil {
		values.Add("source", fmt.Sprintf("%v", *input.Source))
	}
	for _, v := range input.Changes {
		values.Add("changes", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Stores {
		values.Add("stores", fmt.Sprintf("%v", v))
	}
	for _, v := range input.ExcludePlatformInitiatedChanges {
		values.Add("excludePlatformInitiatedChanges", fmt.Sprintf("%v", v))
	}
	if input.Expand != nil {
		if *input.Expand {
			values.Add("expand", "true")
		} else {
			values.Add("expand", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyRequestMethodGet) ResourceTypes(v []ChangeHistoryResourceType) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.ResourceTypes = v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) DateFrom(v interface{}) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.DateFrom = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) DateTo(v interface{}) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.DateTo = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) Limit(v int) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) Offset(v int) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) UserId(v string) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.UserId = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) ClientId(v string) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.ClientId = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) CustomerId(v string) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.CustomerId = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) AssociateId(v string) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.AssociateId = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) BusinessUnit(v string) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.BusinessUnit = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) Type(v string) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.Type = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) ResourceId(v string) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.ResourceId = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) ResourceKey(v string) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.ResourceKey = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) Source(v string) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.Source = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) Changes(v []string) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.Changes = v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) Stores(v []string) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.Stores = v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) ExcludePlatformInitiatedChanges(v []PlatformInitiatedChange) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.ExcludePlatformInitiatedChanges = v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) Expand(v bool) *ByProjectKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRequestMethodGetInput{}
	}
	rb.params.Expand = &v
	return rb
}

func (rb *ByProjectKeyRequestMethodGet) WithQueryParams(input ByProjectKeyRequestMethodGetInput) *ByProjectKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	The `view_audit_log:{projectKey}` scope is required, and depending on the [resource type](ctp:history:type:ChangeHistoryResourceType) queried, their respective scopes must be granted.
 */
func (rb *ByProjectKeyRequestMethodGet) Execute(ctx context.Context) (result *RecordPagedQueryResponse, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
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
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
