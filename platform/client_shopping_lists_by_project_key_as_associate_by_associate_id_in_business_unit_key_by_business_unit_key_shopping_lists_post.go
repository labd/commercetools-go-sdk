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

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPost struct {
	body    ShoppingListDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPostInput
}

func (r *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPost) Expand(v []string) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPost) WithQueryParams(input ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPostInput) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates a ShoppingList in the [BusinessUnit](ctp:api:type:BusinessUnit) referenced by `businessUnitKey`. As such, the `businessUnit` field on [ShoppingListDraft](ctp:api:type:ShoppingListDraft) is ignored for this request.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsRequestMethodPost) Execute(ctx context.Context) (result *ShoppingList, err error) {
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
	case 404:
		return nil, ErrNotFound
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
