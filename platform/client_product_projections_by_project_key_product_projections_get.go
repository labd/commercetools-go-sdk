package platform

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

type ByProjectKeyProductProjectionsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductProjectionsRequestMethodGetInput
}

func (r *ByProjectKeyProductProjectionsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductProjectionsRequestMethodGetInput struct {
	Staged             *bool
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   []string
	StoreProjection    *string
	Expand             []string
	Sort               []string
	Limit              *int
	Offset             *int
	WithTotal          *bool
	Where              []string
	PredicateVar       map[string][]string
}

func (input *ByProjectKeyProductProjectionsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Staged != nil {
		if *input.Staged {
			values.Add("staged", "true")
		} else {
			values.Add("staged", "false")
		}
	}
	if input.PriceCurrency != nil {
		values.Add("priceCurrency", fmt.Sprintf("%v", *input.PriceCurrency))
	}
	if input.PriceCountry != nil {
		values.Add("priceCountry", fmt.Sprintf("%v", *input.PriceCountry))
	}
	if input.PriceCustomerGroup != nil {
		values.Add("priceCustomerGroup", fmt.Sprintf("%v", *input.PriceCustomerGroup))
	}
	if input.PriceChannel != nil {
		values.Add("priceChannel", fmt.Sprintf("%v", *input.PriceChannel))
	}
	for _, v := range input.LocaleProjection {
		values.Add("localeProjection", fmt.Sprintf("%v", v))
	}
	if input.StoreProjection != nil {
		values.Add("storeProjection", fmt.Sprintf("%v", *input.StoreProjection))
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	for k, v := range input.PredicateVar {
		for _, x := range v {
			values.Set(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) Staged(v bool) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) PriceCurrency(v string) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) PriceCountry(v string) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) PriceChannel(v string) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) StoreProjection(v string) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.StoreProjection = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) Expand(v []string) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) Sort(v []string) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) Limit(v int) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) Offset(v int) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) WithTotal(v bool) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) Where(v []string) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyProductProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) WithQueryParams(input ByProjectKeyProductProjectionsRequestMethodGetInput) *ByProjectKeyProductProjectionsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductProjectionsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductProjectionsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Use the Product Projections query endpoint to get the current or staged representations of Products.
*	When used with an API Client that has the `view_published_products:{projectKey}` scope,
*	this endpoint only returns published (current) Product Projections.
*
 */
func (rb *ByProjectKeyProductProjectionsRequestMethodGet) Execute(ctx context.Context) (result *ProductProjectionPagedQueryResponse, err error) {
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
