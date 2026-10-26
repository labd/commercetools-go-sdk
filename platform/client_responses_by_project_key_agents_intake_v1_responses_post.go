package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyAgentsIntakeV1ResponsesRequestMethodPost struct {
	body    AgentResponsesRequest
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyAgentsIntakeV1ResponsesRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyAgentsIntakeV1ResponsesRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyAgentsIntakeV1ResponsesRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Turns a natural-language prompt, optionally with supporting files, into a created [Cart](ctp:api:type:Cart), [QuoteRequest](ctp:api:type:QuoteRequest), or [ShoppingList](ctp:api:type:ShoppingList), returned in commercetools REST representation. The entity to create is set by `outputType` and is never inferred from the input.
*
*	A [ShoppingList](ctp:api:type:ShoppingList) additionally requires a name for the list in the input. Intake never invents one. Unlike a Cart or Quote Request, a Shopping List does not require a country.
*
*	Non-fatal issues, such as requested Products that could not be matched to the catalog or a file that failed to parse, are reported as `warnings` alongside a successful `201` response rather than failing the request.
*
*	Accepts either an `application/json` body or a `multipart/form-data` request. An `application/json` body requires `prompt`. A `multipart/form-data` request requires `prompt`, an uploaded file, or both. See [Multipart form data](/api/projects/intake#multipart-form-data) for the file upload format.
*
*	If Intake is not enabled for the Project, a [FeatureDisabled](ctp:api:type:AgentFeatureDisabledError) error is returned.
*
*	For a list of possible errors returned by Intake, see [Intake Errors](/api/errors#intake) type.
*
 */
func (rb *ByProjectKeyAgentsIntakeV1ResponsesRequestMethodPost) Execute(ctx context.Context) (result *AgentResponsesSuccess, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	queryParams := url.Values{}
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
		errorObj := AgentResponsesErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 401:
		errorObj := AgentResponsesAuthError{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 403:
		errorObj := AgentResponsesAuthError{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 409:
		errorObj := AgentResponsesErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 413:
		errorObj := AgentResponsesErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 500:
		errorObj := AgentResponsesErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 502:
		errorObj := AgentResponsesErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 503:
		errorObj := AgentResponsesErrorResponse{}
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
