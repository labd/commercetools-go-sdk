package customize

import (
	"fmt"
	"net/url"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/mitchellh/mapstructure"
)

type ExtensionDeleteInput struct {
	ID      string
	Version int
}

type ExtensionUpdateInput struct {
	ID      string
	Version int
	Actions commercetools.UpdateActions
}

func (svc *Service) ExtensionGetByID(id string) (*Extension, error) {
	var result Extension
	err := svc.client.Get(fmt.Sprintf("extensions/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ExtensionCreate creates a new API extension.
// Currently, a maximum of 25 extensions can be created per project.
func (svc *Service) ExtensionCreate(draft *ExtensionDraft) (*Extension, error) {
	var result Extension
	err := svc.client.Create("extensions", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) ExtensionUpdate(input *ExtensionUpdateInput) (*Extension, error) {
	var result Extension

	endpoint := fmt.Sprintf("products/%s", input.ID)
	err := svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) ExtensionDeleteByID(id string, version int) (*Extension, error) {
	var result Extension
	endpoint := fmt.Sprintf("extensions/%s", id)
	params := url.Values{}
	params.Set("version", string(version))
	err := svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) ExtensionDeleteByKey(key string, version int) (*Extension, error) {
	var result Extension
	endpoint := fmt.Sprintf("extensions/key=%s", key)
	params := url.Values{}
	params.Set("version", string(version))
	err := svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ExtensionDestinationMapping(input ExtensionDestination) ExtensionDestination {
	inputType := input.(map[string]interface{})["type"]
	switch inputType {
	case "HTTP":
		new := ExtensionDestinationHTTP{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
