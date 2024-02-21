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

type ByProjectKeyByResourceTypeByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyByResourceTypeByIDRequestMethodGetInput
}

func (r *ByProjectKeyByResourceTypeByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyByResourceTypeByIDRequestMethodGetInput struct {
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
	Source                          *string
	Changes                         []string
	Stores                          []string
	ExcludePlatformInitiatedChanges []PlatformInitiatedChange
	Expand                          *bool
}

func (input *ByProjectKeyByResourceTypeByIDRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) DateFrom(v interface{}) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.DateFrom = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) DateTo(v interface{}) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.DateTo = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) Limit(v int) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) Offset(v int) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) UserId(v string) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.UserId = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) ClientId(v string) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.ClientId = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) CustomerId(v string) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.CustomerId = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) AssociateId(v string) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.AssociateId = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) BusinessUnit(v string) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.BusinessUnit = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) Type(v string) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.Type = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) Source(v string) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.Source = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) Changes(v []string) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.Changes = v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) Stores(v []string) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.Stores = v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) ExcludePlatformInitiatedChanges(v []PlatformInitiatedChange) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.ExcludePlatformInitiatedChanges = v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) Expand(v bool) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyByResourceTypeByIDRequestMethodGetInput{}
	}
	rb.params.Expand = &v
	return rb
}

func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) WithQueryParams(input ByProjectKeyByResourceTypeByIDRequestMethodGetInput) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	The `view_audit_log:{projectKey}` scope is required, and depending on the [resource type](ctp:history:type:ChangeHistoryResourceType) queried, their respective scopes must be granted.
 */
func (rb *ByProjectKeyByResourceTypeByIDRequestMethodGet) Execute(ctx context.Context) (result *RecordPagedQueryResponse, err error) {
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
