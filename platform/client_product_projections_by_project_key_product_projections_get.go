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

type ByProjectKeyProductProjectionsRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyProductProjectionsRequestMethodGetInput
}

func (r *ByProjectKeyProductProjectionsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyProductProjectionsRequestMethodGetInput struct {
	Staged             *bool
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   *string
	StoreProjection    *string
	Expand             *[]string
	Sort               *[]string
	Limit              *int
	Offset             *int
	WithTotal          *bool
	Where              *[]string
	PredicateVar       map[string][]string
}

func (input *ByProjectKeyProductProjectionsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Staged != nil {
		if *input.Staged == true {
			values.Add("staged", "true")
		} else {
			values.Add("staged", "false")
		}
	}
	if input.PriceCurrency != nil {
		values.Add("priceCurrency", *input.PriceCurrency)
	}
	if input.PriceCountry != nil {
		values.Add("priceCountry", *input.PriceCountry)
	}
	if input.PriceCustomerGroup != nil {
		values.Add("priceCustomerGroup", *input.PriceCustomerGroup)
	}
	if input.PriceChannel != nil {
		values.Add("priceChannel", *input.PriceChannel)
	}
	if input.LocaleProjection != nil {
		values.Add("localeProjection", *input.LocaleProjection)
	}
	if input.StoreProjection != nil {
		values.Add("storeProjection", *input.StoreProjection)
	}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
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
	if input.Where != nil {
		for _, v := range *input.Where {
			values.Add("where", v)
		}
	}

	return values
}

func (rb *ByProjectKeyProductProjectionsRequestMethodGet) WithQueryParams(input *ByProjectKeyProductProjectionsRequestMethodGetInput) *ByProjectKeyProductProjectionsRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	You can use the product projections query endpoint to get the current or staged representations of Products.
*	When used with an API client that has the view_published_products:{projectKey} scope,
*	this endpoint only returns published (current) product projections.
*
 */
func (rb *ByProjectKeyProductProjectionsRequestMethodGet) Execute() (result *ProductProjectionPagedQueryResponse, err error) {
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
