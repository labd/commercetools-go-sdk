package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// DeleteInput provides the data required to delete a type.
type TypeDeleteInput struct {
	ID      string
	Version int
}

// UpdateInput provides the data required to update a type.
type TypeUpdateInput struct {
	ID string

	// The expected version of the type on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict
	// will be returned.
	Version int

	// The list of update actions to be performed on the type.
	Actions []TypeUpdateAction
}

// GetByID will return a type matching the provided ID. OAuth2 Scopes:
// view_types:{projectKey}
func (client *Client) TypeGetByID(id string) (result *Type, err error) {
	err = client.Get(fmt.Sprintf("types/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create will create a new type from a draft, and return the newly created
// type. OAuth2 Scopes: manage_types:{projectKey}
func (client *Client) TypeCreate(draft *TypeDraft) (result *Type, err error) {
	err = client.Create("types", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update will update a type matching the provided ID with the defined
// UpdateActions. OAuth2 Scopes: manage_types:{projectKey}
func (client *Client) TypeUpdate(input *TypeUpdateInput) (result *Type, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("types/%s", input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByID will delete a type matching the provided ID. These requests delete
// a type only if it’s not referenced by other entities. OAuth2 Scopes:
// manage_types:{projectKey}
func (client *Client) TypeDeleteByID(id string, version int) (result *Type, err error) {
	endpoint := fmt.Sprintf("types/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByKey will delete a type matching the provided key. These requests
// delete a type only if it’s not referenced by other entities. OAuth2 Scopes:
// manage_types:{projectKey}
func (client *Client) TypeDeleteByKey(key string, version int) (result *Type, err error) {
	endpoint := fmt.Sprintf("types/key=%s", key)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
