package history

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	Returned when the [Query Records](/../api/history/change-history#query-records) request exceeds the rate limit.
*
*	Reduce the date range and resource types in your query to minimize the token usage, or retry the request after some time (indicated in the `Retry-After` header).
*
 */
type TooManyRequestsError struct {
	// `"TooManyRequests"`
	Code string `json:"code"`
	// `"You have made too many requests. Please try again later."`
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TooManyRequestsError) MarshalJSON() ([]byte, error) {
	type Alias TooManyRequestsError
	return json.Marshal(struct {
		Action string `json:"null"`
		*Alias
	}{Action: "TooManyRequests", Alias: (*Alias)(&obj)})
}

func (obj TooManyRequestsError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown TooManyRequestsError: failed to parse error response"
}

/**
*	Represents a single error.
 */
type GraphQLErrorObject interface{}

func mapDiscriminatorGraphQLErrorObject(input interface{}) (GraphQLErrorObject, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["code"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'code'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "TooManyRequests":
		obj := GraphQLTooManyRequestsError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Returned when the [Query Records](/../api/history/change-history#query-records) request exceeds the rate limit.
*
*
*
*	Reduce the date range and resource types in your query to minimize the token usage, or retry the request after some time (indicated in the `Retry-After` header).
*
 */
type GraphQLTooManyRequestsError struct {
	// Error-specific additional fields.
	ExtraValues map[string]interface{} `json:"-"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *GraphQLTooManyRequestsError) UnmarshalJSON(data []byte) error {
	type Alias GraphQLTooManyRequestsError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &obj.ExtraValues); err != nil {
		return err
	}
	delete(obj.ExtraValues, "code")

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GraphQLTooManyRequestsError) MarshalJSON() ([]byte, error) {
	type Alias GraphQLTooManyRequestsError
	data, err := json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "TooManyRequests", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	for key, value := range obj.ExtraValues {
		raw[key] = value
	}

	return json.Marshal(raw)

}

func (obj *GraphQLTooManyRequestsError) DecodeStruct(src map[string]interface{}) error {
	{
		obj.ExtraValues = make(map[string]interface{})
		for key, value := range src {
			//
			if key != "code" {
				obj.ExtraValues[key] = value
			}
		}
	}
	return nil
}
