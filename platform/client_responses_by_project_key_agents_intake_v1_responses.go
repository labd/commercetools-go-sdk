package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAgentsIntakeV1ResponsesRequestBuilder struct {
	projectKey string
	client     *Client
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
func (rb *ByProjectKeyAgentsIntakeV1ResponsesRequestBuilder) Post(body AgentResponsesRequest) *ByProjectKeyAgentsIntakeV1ResponsesRequestMethodPost {
	return &ByProjectKeyAgentsIntakeV1ResponsesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/agents/intake/v1/responses", rb.projectKey),
		client: rb.client,
	}
}
