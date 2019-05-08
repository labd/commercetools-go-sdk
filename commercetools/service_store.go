package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// StoreUpdateInput provides the data required to update a store.
type StoreUpdateInput struct {
	ID string

	// The expected version of the store on which the changes should be
	// applied. If the expected version does not match the actual version, a 409
	// Conflict will be returned.
	Version int

	// The list of update actions to be performed on the store.
	Actions []StoreUpdateAction
}

// StoreUpdateKeyInput provides the data required to update a store.
type StoreUpdateKeyInput struct {
	Key string

	// The expected version of the store on which the changes should be
	// applied. If the expected version does not match the actual version, a 409
	// Conflict will be returned.
	Version int

	// The list of update actions to be performed on the store.
	Actions []StoreUpdateAction
}

// StoreURLPath is the commercetools API store path.
const StoreURLPath = "stores"

// StoreGetByID will return a store matching the provided ID. OAuth2 Scopes:
// view_stores:{projectKey}
func (client *Client) StoreGetByID(id string) (result *Store, err error) {
	err = client.Get(fmt.Sprintf("%s/%s", StoreURLPath, id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreGetByKey will return a store matching the provided key. OAuth2 Scopes:
// view_stores:{projectKey}
func (client *Client) StoreGetByKey(key string) (result *Store, err error) {
	err = client.Get(fmt.Sprintf("%s/key=%s", StoreURLPath, key), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCreate will create a new store from a draft, and return the newly created
// store. OAuth2 Scopes: manage_stores:{projectKey}
func (client *Client) StoreCreate(draft *StoreDraft) (result *Store, err error) {
	err = client.Create(StoreURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreUpdate will update a store matching the provided ID with the defined
// StoreUpdateActions. OAuth2 Scopes: manage_stores:{projectKey}
func (client *Client) StoreUpdate(input *StoreUpdateInput) (result *Store, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("%s/%s", StoreURLPath, input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreUpdateByKey will update a store matching the provided ID with the defined
// StoreUpdateActions. OAuth2 Scopes: manage_stores:{projectKey}
func (client *Client) StoreUpdateByKey(input *StoreUpdateKeyInput) (result *Store, err error) {
	if input.Key == "" {
		return nil, fmt.Errorf("no valid type key passed")
	}

	endpoint := fmt.Sprintf("%s/key=%s", StoreURLPath, input.Key)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreDelete will delete a type matching the provided ID. These requests
// delete a type only if it’s not referenced by other entities. OAuth2 Scopes:
// manage_stores:{projectKey}
func (client *Client) StoreDelete(id string, version int) (result *Store, err error) {
	endpoint := fmt.Sprintf("%s/%s", StoreURLPath, id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreDeleteByKey will delete a type matching the provided key. These requests
// delete a type only if it’s not referenced by other entities. OAuth2 Scopes:
// manage_stores:{projectKey}
func (client *Client) StoreDeleteByKey(key string, version int) (result *Store, err error) {
	endpoint := fmt.Sprintf("%s/key=%s", StoreURLPath, key)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreQuery will query stores.
// OAuth2 Scopes: view_stores:{projectKey}
func (client *Client) StoreQuery(input *QueryInput) (result *StorePagedQueryResponse, err error) {
	err = client.Query(StoreURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
