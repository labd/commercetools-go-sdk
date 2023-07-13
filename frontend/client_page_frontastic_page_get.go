package frontend

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FrontasticPageRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *FrontasticPageRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *FrontasticPageRequestMethodGet) WithHeaders(headers http.Header) *FrontasticPageRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Returns the structure and data of the published page displayed from commercetools Frontend. \
*
*	Headers&#58;
*
*	- `Frontastic-Path` - `^/.*$` - Required
*	- `Accept` - `application/json` - Required
*	- `Frontastic-Locale` - [Locale](ctp:frontend-api:type:Locale) - Required
*
 */
func (rb *FrontasticPageRequestMethodGet) Execute(ctx context.Context) (result *PageDataResponse, err error) {
	queryParams := url.Values{}
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
		return result, nil
	// case 301:
	// 	errorObj := RedirectResponse{}
	// 	err = json.Unmarshal(content, &errorObj)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return nil, errorObj
	case 404:
		errorObj := Error{}
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
