package commercetools

import (
	"fmt"
	"net/url"
)

// DeleteInput provides the data required to delete an extension.
type ExtensionDeleteInput struct {
	ID      string
	Version int
}

// UpdateInput provides the data required to update an extension.
type ExtensionUpdateInput struct {
	ID      string
	Version int
	Actions []ExtensionUpdateAction
}

// Service contains client information and bundles all actions.
type ExtensionService struct {
	client *Client
}

// GetByID will return an extension matching the provided ID. OAuth2 Scopes:
// manage_extensions:{projectKey}
func (svc *ExtensionService) GetByID(id string) (result *Extension, err error) {
	err = svc.client.Get(fmt.Sprintf("extensions/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates a new API extension. Currently, a maximum of 25 extensions can
// be created per project. OAuth2 Scopes: manage_extensions:{projectKey}
func (svc *ExtensionService) Create(draft *ExtensionDraft) (result *Extension, err error) {
	err = svc.client.Create("extensions", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update will update an extension matching the provided ID with the defined
// UpdateActions. OAuth2 Scopes: manage_extensions:{projectKey}
func (svc *ExtensionService) Update(input *ExtensionUpdateInput) (result *Extension, err error) {
	endpoint := fmt.Sprintf("extensions/%s", input.ID)
	err = svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByID will delete an extension matching the provided ID. OAuth2 Scopes:
// manage_extensions:{projectKey}
func (svc *ExtensionService) DeleteByID(id string, version int) (result *Extension, err error) {
	endpoint := fmt.Sprintf("extensions/%s", id)
	params := url.Values{}
	params.Set("version", fmt.Sprintf("%d", version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByKey will delete an extension matching the provided key. OAuth2
// Scopes: manage_extensions:{projectKey}
func (svc *ExtensionService) DeleteByKey(key string, version int) (result *Extension, err error) {
	endpoint := fmt.Sprintf("extensions/key=%s", key)
	params := url.Values{}
	params.Set("version", string(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
