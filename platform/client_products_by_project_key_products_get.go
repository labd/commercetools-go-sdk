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

type ByProjectKeyProductsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductsRequestMethodGetInput
}

func (r *ByProjectKeyProductsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductsRequestMethodGetInput struct {
	Where              []string
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   []string
	Expand             []string
	Sort               []string
	Limit              *int
	Offset             *int
	WithTotal          *bool
	PredicateVar       map[string][]string
}

func (input *ByProjectKeyProductsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
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
	for k, v := range input.PredicateVar {
		for _, x := range v {
			values.Add(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyProductsRequestMethodGet) Where(v []string) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) PriceCurrency(v string) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) PriceCountry(v string) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) PriceChannel(v string) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) Expand(v []string) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) Sort(v []string) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) Limit(v int) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) Offset(v int) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) WithTotal(v bool) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodGet) WithQueryParams(input ByProjectKeyProductsRequestMethodGetInput) *ByProjectKeyProductsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	If [Price selection](ctp:api:type:ProductPriceSelection) query parameters are provided, the selected Prices are added to the response.
 */
func (rb *ByProjectKeyProductsRequestMethodGet) Execute(ctx context.Context) (result *ProductPagedQueryResponse, err error) {
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
