package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyAssociateRolesRequestMethodPost struct {
	body    AssociateRoleDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyAssociateRolesRequestMethodPostInput
}

func (r *ByProjectKeyAssociateRolesRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyAssociateRolesRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyAssociateRolesRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyAssociateRolesRequestMethodPost) Expand(v []string) *ByProjectKeyAssociateRolesRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyAssociateRolesRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyAssociateRolesRequestMethodPost) WithQueryParams(input ByProjectKeyAssociateRolesRequestMethodPostInput) *ByProjectKeyAssociateRolesRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyAssociateRolesRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyAssociateRolesRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creating a Associate Role generates the [AssociateRoleCreated](ctp:api:type:AssociateRoleCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyAssociateRolesRequestMethodPost) Execute(ctx context.Context) (result *AssociateRole, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
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
	case 201:
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
	case 502:
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
