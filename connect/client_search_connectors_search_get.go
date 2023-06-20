package connect

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

type ConnectorsSearchRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ConnectorsSearchRequestMethodGetInput
}

func (r *ConnectorsSearchRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ConnectorsSearchRequestMethodGetInput struct {
	Text           *string
	Limit          *int
	Offset         *int
	Sort           []string
	Private        *bool
	CreatorCompany *string
	ID             []string
	Key            []string
}

func (input *ConnectorsSearchRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Text != nil {
		values.Add("text", fmt.Sprintf("%v", *input.Text))
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	if input.Private != nil {
		if *input.Private {
			values.Add("private", "true")
		} else {
			values.Add("private", "false")
		}
	}
	if input.CreatorCompany != nil {
		values.Add("creator.company", fmt.Sprintf("%v", *input.CreatorCompany))
	}
	for _, v := range input.ID {
		values.Add("id", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Key {
		values.Add("key", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ConnectorsSearchRequestMethodGet) Text(v string) *ConnectorsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ConnectorsSearchRequestMethodGetInput{}
	}
	rb.params.Text = &v
	return rb
}

func (rb *ConnectorsSearchRequestMethodGet) Limit(v int) *ConnectorsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ConnectorsSearchRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ConnectorsSearchRequestMethodGet) Offset(v int) *ConnectorsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ConnectorsSearchRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ConnectorsSearchRequestMethodGet) Sort(v []string) *ConnectorsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ConnectorsSearchRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ConnectorsSearchRequestMethodGet) Private(v bool) *ConnectorsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ConnectorsSearchRequestMethodGetInput{}
	}
	rb.params.Private = &v
	return rb
}

func (rb *ConnectorsSearchRequestMethodGet) CreatorCompany(v string) *ConnectorsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ConnectorsSearchRequestMethodGetInput{}
	}
	rb.params.CreatorCompany = &v
	return rb
}

func (rb *ConnectorsSearchRequestMethodGet) ID(v []string) *ConnectorsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ConnectorsSearchRequestMethodGetInput{}
	}
	rb.params.ID = v
	return rb
}

func (rb *ConnectorsSearchRequestMethodGet) Key(v []string) *ConnectorsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ConnectorsSearchRequestMethodGetInput{}
	}
	rb.params.Key = v
	return rb
}

func (rb *ConnectorsSearchRequestMethodGet) WithQueryParams(input ConnectorsSearchRequestMethodGetInput) *ConnectorsSearchRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ConnectorsSearchRequestMethodGet) WithHeaders(headers http.Header) *ConnectorsSearchRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves all available Connectors.
 */
func (rb *ConnectorsSearchRequestMethodGet) Execute(ctx context.Context) (result *Paged, err error) {
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
		return result, nil
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
