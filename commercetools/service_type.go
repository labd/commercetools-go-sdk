package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// TypeDeleteInput provides the data required to delete a type.
type TypeDeleteInput struct {
	ID      string
	Version int
}

// TypeUpdateInput provides the data required to update a type.
type TypeUpdateInput struct {
	ID string

	// The expected version of the type on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict
	// will be returned.
	Version int

	// The list of update actions to be performed on the type.
	Actions []TypeUpdateAction
}

// TypeURLPath is the commercetools API type path.
const TypeURLPath = "types"

// TypeGetByID will return a type matching the provided ID. OAuth2 Scopes:
// view_types:{projectKey}
func (client *Client) TypeGetByID(id string) (result *Type, err error) {
	err = client.Get(fmt.Sprintf("%s/%s", TypeURLPath, id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeCreate will create a new type from a draft, and return the newly created
// type. OAuth2 Scopes: manage_types:{projectKey}
func (client *Client) TypeCreate(draft *TypeDraft) (result *Type, err error) {
	err = client.Create(TypeURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeUpdate will update a type matching the provided ID with the defined
// TypeUpdateActions. OAuth2 Scopes: manage_types:{projectKey}
func (client *Client) TypeUpdate(input *TypeUpdateInput) (result *Type, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("%s/%s", TypeURLPath, input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeDeleteByID will delete a type matching the provided ID. These requests delete
// a type only if it’s not referenced by other entities. OAuth2 Scopes:
// manage_types:{projectKey}
func (client *Client) TypeDeleteByID(id string, version int) (result *Type, err error) {
	endpoint := fmt.Sprintf("%s/%s", TypeURLPath, id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeDeleteByKey will delete a type matching the provided key. These requests
// delete a type only if it’s not referenced by other entities. OAuth2 Scopes:
// manage_types:{projectKey}
func (client *Client) TypeDeleteByKey(key string, version int) (result *Type, err error) {
	endpoint := fmt.Sprintf("%s/key=%s", TypeURLPath, key)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeQuery will query types.
// OAuth2 Scopes: view_types:{projectKey}
func (client *Client) TypeQuery(input *QueryInput) (result *TypePagedQueryResponse, err error) {
	err = client.Query(TypeURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
