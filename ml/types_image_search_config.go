package ml

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type ImageSearchConfigStatus string

const (
	ImageSearchConfigStatusOn  ImageSearchConfigStatus = "on"
	ImageSearchConfigStatusOff ImageSearchConfigStatus = "off"
)

type ImageSearchConfigUpdateAction interface{}

func mapDiscriminatorImageSearchConfigUpdateAction(input interface{}) (ImageSearchConfigUpdateAction, error) {
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
	case "changeStatus":
		obj := ChangeStatusUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ChangeStatusUpdateAction struct {
	Status ImageSearchConfigStatus `json:"status"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeStatusUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias ChangeStatusUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeStatus", Alias: (*Alias)(&obj)})
}

type ImageSearchConfigRequest struct {
	// The list of update actions to be performed on the project.
	Actions []ImageSearchConfigUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ImageSearchConfigRequest) UnmarshalJSON(data []byte) error {
	type Alias ImageSearchConfigRequest
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorImageSearchConfigUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ImageSearchConfigResponse struct {
	// The image search activation status.
	Status         ImageSearchConfigStatus `json:"status"`
	LastModifiedAt time.Time               `json:"lastModifiedAt"`
}
