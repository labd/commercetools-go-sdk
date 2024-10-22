package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	Represents a warning related to the returned response.
 */
type WarningObject interface{}

func mapDiscriminatorWarningObject(input interface{}) (WarningObject, error) {
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
	case "ImageProcessingOngoing":
		obj := ImageProcessingOngoingWarning{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Contained in responses to an [Upload Product image](/projects/products#upload-product-image) or an [Upload Product Tailoring image](/projects/product-tailoring#upload-product-tailoring-image) request with response status code `202 Accepted`.
*	Indicates that the API is still creating the remaining sizes of the uploaded image. They will be available on the Content Delivery Network (CDN) soon.
*
 */
type ImageProcessingOngoingWarning struct {
	// `"The image processing is still ongoing."`
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ImageProcessingOngoingWarning) MarshalJSON() ([]byte, error) {
	type Alias ImageProcessingOngoingWarning
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ImageProcessingOngoing", Alias: (*Alias)(&obj)})
}
