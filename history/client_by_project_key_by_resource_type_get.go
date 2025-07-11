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

type ByProjectKeyByResourceTypeRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyByResourceTypeRequestMethodGetInput
}

func (r *ByProjectKeyByResourceTypeRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyByResourceTypeRequestMethodGetInput struct {
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
	ResourceKey                     *string
	Source                          *Source
	Changes                         []string
	Stores                          []string
	ExcludePlatformInitiatedChanges []PlatformInitiatedChange
	Expand                          *bool
}

func (input *ByProjectKeyByResourceTypeRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) DateFrom(v interface{}) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.DateFrom = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) DateTo(v interface{}) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.DateTo = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) Limit(v int) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) Offset(v int) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) UserId(v string) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.UserId = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) ClientId(v string) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.ClientId = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) CustomerId(v string) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.CustomerId = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) AssociateId(v string) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.AssociateId = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) BusinessUnit(v string) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.BusinessUnit = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) Type(v string) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.Type = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) ResourceKey(v string) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.ResourceKey = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) Source(v Source) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.Source = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) Changes(v []string) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.Changes = v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) Stores(v []string) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.Stores = v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) ExcludePlatformInitiatedChanges(v []PlatformInitiatedChange) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.ExcludePlatformInitiatedChanges = v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) Expand(v bool) *ByProjectKeyByResourceTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeRequestMethodGetInput{}
	}
	rb.params.Expand = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeRequestMethodGet) WithQueryParams(input ByProjectKeyByResourceTypeRequestMethodGetInput) *ByProjectKeyByResourceTypeRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyByResourceTypeRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyByResourceTypeRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	The `view_audit_log:{projectKey}` scope is required, and depending on the [resource type](ctp:history:type:ChangeHistoryResourceType) queried, their respective scopes must be granted.
 */
func (rb *ByProjectKeyByResourceTypeRequestMethodGet) Execute(ctx context.Context) (result *RecordPagedQueryResponse, err error) {
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
