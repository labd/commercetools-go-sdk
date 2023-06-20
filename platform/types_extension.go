package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type Extension struct {
	// Unique identifier of the Extension.
	ID string `json:"id"`
	// Current version of the Extension.
	Version int `json:"version"`
	// Date and time (UTC) the Extension was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Extension was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the Extension.
	Key *string `json:"key,omitempty"`
	// The configuration for the Extension, including its type, location and authentication details.
	Destination ExtensionDestination `json:"destination"`
	// Describes what triggers the Extension.
	Triggers []ExtensionTrigger `json:"triggers"`
	// Maximum time (in milliseconds) that the Extension can respond within.
	// If no timeout is provided, the default value is used for all types of Extensions.
	// The maximum value is 10000 ms (10 seconds) for `payment` Extensions and 2000 ms (2 seconds) for all other Extensions.
	TimeoutInMs *int `json:"timeoutInMs,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Extension) UnmarshalJSON(data []byte) error {
	type Alias Extension
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorExtensionDestination(obj.Destination)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	An Extension gets called during any of the following requests of an API call, but before the result is persisted.
*
 */
type ExtensionAction string

const (
	ExtensionActionCreate ExtensionAction = "Create"
	ExtensionActionUpdate ExtensionAction = "Update"
)

/**
*	Generic type for destinations.
 */
type ExtensionDestination interface{}

func mapDiscriminatorExtensionDestination(input interface{}) (ExtensionDestination, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "AWSLambda":
		obj := AWSLambdaDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "GoogleCloudFunction":
		obj := GoogleCloudFunctionDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "HTTP":
		obj := HttpDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Authentication != nil {
			var err error
			obj.Authentication, err = mapDiscriminatorHttpDestinationAuthentication(obj.Authentication)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	We recommend creating an Identify and Access Management (IAM) user with an `accessKey` and `accessSecret` pair, specifically for each Extension that only has the `lambda:InvokeFunction` permission on this function.
*
 */
type AWSLambdaDestination struct {
	// Amazon Resource Name (ARN) of the Lambda function in the format `arn:aws:lambda:<region>:<accountid>:function:<functionName>`. Use the format `arn:aws:lambda:<region>:<accountid>:function:<functionName>:<functionAlias/version>` to point to a specific version of the function.
	Arn string `json:"arn"`
	// Partially hidden on retrieval for security reasons.
	AccessKey string `json:"accessKey"`
	// Partially hidden on retrieval for security reasons.
	AccessSecret string `json:"accessSecret"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AWSLambdaDestination) MarshalJSON() ([]byte, error) {
	type Alias AWSLambdaDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AWSLambda", Alias: (*Alias)(&obj)})
}

type ExtensionDraft struct {
	// User-defined unique identifier for the Extension.
	Key *string `json:"key,omitempty"`
	// Defines where the Extension can be reached.
	Destination ExtensionDestination `json:"destination"`
	// Describes what triggers the Extension.
	Triggers []ExtensionTrigger `json:"triggers"`
	// Maximum time (in milliseconds) the Extension can respond within.
	// If no timeout is provided, the default value is used for all types of Extensions.
	// The maximum value is 10000 ms (10 seconds) for `payment` Extensions and 2000 ms (2 seconds) for all other Extensions.
	//
	// This limit can be increased per Project after we review the performance impact.
	// Please contact our support via the [Support Portal](https://support.commercetools.com) and provide the Region, Project key, and use case.
	TimeoutInMs *int `json:"timeoutInMs,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionDraft) UnmarshalJSON(data []byte) error {
	type Alias ExtensionDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorExtensionDestination(obj.Destination)
		if err != nil {
			return err
		}
	}

	return nil
}

type ExtensionInput struct {
	// `Create` or `Update` request.
	Action ExtensionAction `json:"action"`
	// Expanded reference to the resource that triggered the Extension.
	Resource Reference `json:"resource"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionInput) UnmarshalJSON(data []byte) error {
	type Alias ExtensionInput
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [Extension](ctp:api:type:Extension).
*
 */
type ExtensionPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// Extensions matching the query.
	Results []Extension `json:"results"`
}

/**
*	Extensions are available for:
*
 */
type ExtensionResourceTypeId string

const (
	ExtensionResourceTypeIdCart         ExtensionResourceTypeId = "cart"
	ExtensionResourceTypeIdOrder        ExtensionResourceTypeId = "order"
	ExtensionResourceTypeIdPayment      ExtensionResourceTypeId = "payment"
	ExtensionResourceTypeIdCustomer     ExtensionResourceTypeId = "customer"
	ExtensionResourceTypeIdQuoteRequest ExtensionResourceTypeId = "quote-request"
	ExtensionResourceTypeIdStagedQuote  ExtensionResourceTypeId = "staged-quote"
	ExtensionResourceTypeIdQuote        ExtensionResourceTypeId = "quote"
	ExtensionResourceTypeIdBusinessUnit ExtensionResourceTypeId = "business-unit"
)

type ExtensionTrigger struct {
	// `cart`, `order`, `payment`, `customer`, `quote-request`, `staged-quote`, `quote`, and `business-unit` are supported.
	ResourceTypeId ExtensionResourceTypeId `json:"resourceTypeId"`
	// `Create` and `Update` requests are supported.
	Actions []ExtensionAction `json:"actions"`
	// Valid [predicate](/../api/predicates/query) that controls the conditions under which the API Extension is called. The Extension is not triggered when the specified condition is not fulfilled.
	Condition *string `json:"condition,omitempty"`
}

type ExtensionUpdate struct {
	// Expected version of the Extension on which the changes should be applied. If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error is returned.
	Version int `json:"version"`
	// Update actions to be performed on the Extension.
	Actions []ExtensionUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionUpdate) UnmarshalJSON(data []byte) error {
	type Alias ExtensionUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorExtensionUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ExtensionUpdateAction interface{}

func mapDiscriminatorExtensionUpdateAction(input interface{}) (ExtensionUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "changeDestination":
		obj := ExtensionChangeDestinationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Destination != nil {
			var err error
			obj.Destination, err = mapDiscriminatorExtensionDestination(obj.Destination)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "changeTriggers":
		obj := ExtensionChangeTriggersAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ExtensionSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTimeoutInMs":
		obj := ExtensionSetTimeoutInMsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	For GoogleCloudFunction destinations, you need to grant permissions to the `extensions@commercetools-platform.iam.gserviceaccount.com` service account to invoke your function. If your function's version is 1st gen, grant the service account the IAM role `Cloud Functions Invoker`. For version 2nd gen, assign the IAM role `Cloud Run Invoker` using the Cloud Run console.
*
 */
type GoogleCloudFunctionDestination struct {
	// URL to the target function.
	Url string `json:"url"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GoogleCloudFunctionDestination) MarshalJSON() ([]byte, error) {
	type Alias GoogleCloudFunctionDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "GoogleCloudFunction", Alias: (*Alias)(&obj)})
}

/**
*	We recommend an encrypted `HTTPS` connection for production setups. However, we also accept unencrypted `HTTP` connections for development purposes. HTTP redirects will not be followed and cache headers will be ignored.
*
 */
type HttpDestination struct {
	// URL to the target destination. If the Project is hosted in the China (AWS, Ningxia) Region, verify that the URL is not blocked due to firewall restrictions.
	Url string `json:"url"`
	// Authentication methods (such as `Basic` or `Bearer`).
	Authentication HttpDestinationAuthentication `json:"authentication,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *HttpDestination) UnmarshalJSON(data []byte) error {
	type Alias HttpDestination
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Authentication != nil {
		var err error
		obj.Authentication, err = mapDiscriminatorHttpDestinationAuthentication(obj.Authentication)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj HttpDestination) MarshalJSON() ([]byte, error) {
	type Alias HttpDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "HTTP", Alias: (*Alias)(&obj)})
}

type HttpDestinationAuthentication interface{}

func mapDiscriminatorHttpDestinationAuthentication(input interface{}) (HttpDestinationAuthentication, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "AuthorizationHeader":
		obj := AuthorizationHeaderAuthentication{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AzureFunctions":
		obj := AzureFunctionsAuthentication{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	The `Authorization` header will be set to the content of `headerValue`. The authentication scheme (such as `Basic` or `Bearer`) should be included in the `headerValue`.
*
*	For example, the `headerValue` for [Basic Authentication](https://datatracker.ietf.org/doc/html/rfc7617) should be set to `Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==`.
*
 */
type AuthorizationHeaderAuthentication struct {
	// Partially hidden on retrieval for security reasons.
	HeaderValue string `json:"headerValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AuthorizationHeaderAuthentication) MarshalJSON() ([]byte, error) {
	type Alias AuthorizationHeaderAuthentication
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AuthorizationHeader", Alias: (*Alias)(&obj)})
}

/**
*	To protect your Azure Function, set its `authLevel` to `function` and provide the function's key to be used inside the `x-functions-key` header. For more information, see the [Azure Functions documentation](https://docs.microsoft.com/en-us/azure/azure-functions/functions-bindings-http-webhook#keys).
*
*	To protect the secret key from being exposed, remove the code parameter and secret key from the URL. For example, use `https://foo.azurewebsites.net/api/bar` instead of
*	`https://foo.azurewebsites.net/api/bar?code=secret`.
*
 */
type AzureFunctionsAuthentication struct {
	// Partially hidden on retrieval for security reasons.
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AzureFunctionsAuthentication) MarshalJSON() ([]byte, error) {
	type Alias AzureFunctionsAuthentication
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AzureFunctions", Alias: (*Alias)(&obj)})
}

type ExtensionChangeDestinationAction struct {
	// New value to set. Must not be empty.
	Destination ExtensionDestination `json:"destination"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionChangeDestinationAction) UnmarshalJSON(data []byte) error {
	type Alias ExtensionChangeDestinationAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorExtensionDestination(obj.Destination)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionChangeDestinationAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeDestinationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDestination", Alias: (*Alias)(&obj)})
}

type ExtensionChangeTriggersAction struct {
	// New value to set. Must not be empty.
	Triggers []ExtensionTrigger `json:"triggers"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionChangeTriggersAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeTriggersAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTriggers", Alias: (*Alias)(&obj)})
}

type ExtensionSetKeyAction struct {
	// Value to set. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ExtensionSetTimeoutInMsAction struct {
	// Value to set. If not defined, the maximum value is used.
	// If no timeout is provided, the default value is used for all types of Extensions.
	// The maximum value is 10000 ms (10 seconds) for `payment` Extensions and 2000 ms (2 seconds) for all other Extensions.
	//
	// This limit can be increased per Project after we review the performance impact.
	// Please contact our support via the [Support Portal](https://support.commercetools.com/) and provide the Region, Project key, and use case.
	TimeoutInMs *int `json:"timeoutInMs,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExtensionSetTimeoutInMsAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionSetTimeoutInMsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTimeoutInMs", Alias: (*Alias)(&obj)})
}
