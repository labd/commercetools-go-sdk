package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyProductProjectionsSuggestRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductProjectionsSuggestRequestMethodGetInput
}

func (r *ByProjectKeyProductProjectionsSuggestRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductProjectionsSuggestRequestMethodGetInput struct {
	SearchKeywords map[string][]string
	Limit          *int
	Fuzzy          *bool
	Staged         *bool
}

func (input *ByProjectKeyProductProjectionsSuggestRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for k, v := range input.SearchKeywords {
		for _, x := range v {
			values.Add(k, x)
		}
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Fuzzy != nil {
		if *input.Fuzzy {
			values.Add("fuzzy", "true")
		} else {
			values.Add("fuzzy", "false")
		}
	}
	if input.Staged != nil {
		if *input.Staged {
			values.Add("staged", "true")
		} else {
			values.Add("staged", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyProductProjectionsSuggestRequestMethodGet) SearchKeywords(v map[string][]string) *ByProjectKeyProductProjectionsSuggestRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSuggestRequestMethodGetInput{}
	}
	rb.params.SearchKeywords = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSuggestRequestMethodGet) Limit(v int) *ByProjectKeyProductProjectionsSuggestRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSuggestRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSuggestRequestMethodGet) Fuzzy(v bool) *ByProjectKeyProductProjectionsSuggestRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSuggestRequestMethodGetInput{}
	}
	rb.params.Fuzzy = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSuggestRequestMethodGet) Staged(v bool) *ByProjectKeyProductProjectionsSuggestRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSuggestRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSuggestRequestMethodGet) WithQueryParams(input ByProjectKeyProductProjectionsSuggestRequestMethodGetInput) *ByProjectKeyProductProjectionsSuggestRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductProjectionsSuggestRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductProjectionsSuggestRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyProductProjectionsSuggestRequestMethodGet) Execute(ctx context.Context) (result *SuggestionResult, err error) {
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
