package commercetools

import (
	"fmt"
	"net/url"
)

// ExtensionDeleteInput provides the data required to delete an extension.
type ExtensionDeleteInput struct {
	ID      string
	Version int
}

// ExtensionUpdateInput provides the data required to update an extension.
type ExtensionUpdateInput struct {
	ID      string
	Version int
	Actions []ExtensionUpdateAction
}

// ExtensionGetByID will return an extension matching the provided ID. OAuth2 Scopes:
// manage_extensions:{projectKey}
func (client *Client) ExtensionGetByID(id string) (result *Extension, err error) {
	err = client.Get(fmt.Sprintf("extensions/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionCreate creates a new API extension. Currently, a maximum of 25 extensions can
// be created per project. OAuth2 Scopes: manage_extensions:{projectKey}
func (client *Client) ExtensionCreate(draft *ExtensionDraft) (result *Extension, err error) {
	err = client.Create("extensions", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionUpdate will update an extension matching the provided ID with the defined
// ExtensionUpdateActions. OAuth2 Scopes: manage_extensions:{projectKey}
func (client *Client) ExtensionUpdate(input *ExtensionUpdateInput) (result *Extension, err error) {
	endpoint := fmt.Sprintf("extensions/%s", input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionDeleteByID will delete an extension matching the provided ID. OAuth2 Scopes:
// manage_extensions:{projectKey}
func (client *Client) ExtensionDeleteByID(id string, version int) (result *Extension, err error) {
	endpoint := fmt.Sprintf("extensions/%s", id)
	params := url.Values{}
	params.Set("version", fmt.Sprintf("%d", version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionDeleteByKey will delete an extension matching the provided key. OAuth2
// Scopes: manage_extensions:{projectKey}
func (client *Client) ExtensionDeleteByKey(key string, version int) (result *Extension, err error) {
	endpoint := fmt.Sprintf("extensions/key=%s", key)
	params := url.Values{}
	params.Set("version", string(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
