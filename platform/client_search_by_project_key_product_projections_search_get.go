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

type ByProjectKeyProductProjectionsSearchRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyProductProjectionsSearchRequestMethodGetInput
}

func (r *ByProjectKeyProductProjectionsSearchRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyProductProjectionsSearchRequestMethodGetInput struct {
	Fuzzy                *bool
	FuzzyLevel           *float64
	MarkMatchingVariants bool
	Staged               *bool
	Filter               *[]string
	FilterFacets         *[]string
	FilterQuery          *[]string
	Facet                *[]string
	Text                 map[string][]string
	Sort                 *[]string
	Limit                *int
	Offset               *int
	WithTotal            *bool
	PriceCurrency        *string
	PriceCountry         *string
	PriceCustomerGroup   *string
	PriceChannel         *string
	LocaleProjection     *string
	StoreProjection      *string
	Expand               *[]string
}

func (input *ByProjectKeyProductProjectionsSearchRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Fuzzy != nil {
		if *input.Fuzzy == true {
			values.Add("fuzzy", "true")
		} else {
			values.Add("fuzzy", "false")
		}
	}
	if input.FuzzyLevel != nil {
		values.Add("fuzzyLevel", fmt.Sprintf("%f", *input.FuzzyLevel))
	}
	if input.MarkMatchingVariants == true {
		values.Add("markMatchingVariants", "true")
	} else {
		values.Add("markMatchingVariants", "false")
	}
	if input.Staged != nil {
		if *input.Staged == true {
			values.Add("staged", "true")
		} else {
			values.Add("staged", "false")
		}
	}
	if input.Filter != nil {
		for _, v := range *input.Filter {
			values.Add("filter", v)
		}
	}
	if input.FilterFacets != nil {
		for _, v := range *input.FilterFacets {
			values.Add("filter.facets", v)
		}
	}
	if input.FilterQuery != nil {
		for _, v := range *input.FilterQuery {
			values.Add("filter.query", v)
		}
	}
	if input.Facet != nil {
		for _, v := range *input.Facet {
			values.Add("facet", v)
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
	return values
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) WithQueryParams(input *ByProjectKeyProductProjectionsSearchRequestMethodGetInput) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	Search Product Projection
 */
func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Execute() (result *ProductProjectionPagedSearchResponse, err error) {
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
