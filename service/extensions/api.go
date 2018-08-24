package extensions

import (
	"fmt"
	"net/url"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/mitchellh/mapstructure"
)

type DeleteInput struct {
	ID      string
	Version int
}

type UpdateInput struct {
	ID      string
	Version int
	Actions commercetools.UpdateActions
}

func (svc *Service) GetByID(id string) (result *Extension, err error) {
	err = svc.client.Get(fmt.Sprintf("extensions/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates a new API extension.
// Currently, a maximum of 25 extensions can be created per project.
func (svc *Service) Create(draft *ExtensionDraft) (result *Extension, err error) {
	err = svc.client.Create("extensions", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *Service) Update(input *UpdateInput) (result *Extension, err error) {
	endpoint := fmt.Sprintf("products/%s", input.ID)
	err = svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *Service) DeleteByID(id string, version int) (result *Extension, err error) {
	endpoint := fmt.Sprintf("extensions/%s", id)
	params := url.Values{}
	params.Set("version", fmt.Sprintf("%d", version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *Service) DeleteByKey(key string, version int) (result *Extension, err error) {
	endpoint := fmt.Sprintf("extensions/key=%s", key)
	params := url.Values{}
	params.Set("version", string(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func destinationMapping(input Destination) Destination {
	inputType := input.(map[string]interface{})["type"]
	switch inputType {
	case "HTTP":
		new := DestinationHTTP{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
