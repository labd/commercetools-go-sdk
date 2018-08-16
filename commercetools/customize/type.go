package customize

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

type TypeDeleteInput struct {
	ID      string
	Version int
}

type TypeUpdateInput struct {
	ID      string
	Version int
	Actions commercetools.UpdateActions
}

// OAuth2 Scopes: view_types:{projectKey}
func (svc *Service) TypeGetByID(id string) (result *Type, err error) {
	err = svc.client.Get(fmt.Sprintf("types/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OAuth2 Scopes: manage_types:{projectKey}
func (svc *Service) TypeCreate(draft *TypeDraft) (result *Type, err error) {
	err = svc.client.Create("types", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// The expected version of the type on which the changes should be applied.
// If the expected version does not match the actual version, a 409 Conflict will be returned.
// OAuth2 Scopes: manage_types:{projectKey}
func (svc *Service) TypeUpdate(input *TypeUpdateInput) (result *Type, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("types/%s", input.ID)
	err = svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// These requests delete a type only if it’s not referenced by other entities.
// OAuth2 Scopes: manage_types:{projectKey}
func (svc *Service) TypeDeleteByID(id string, version int) (result *Type, err error) {
	endpoint := fmt.Sprintf("types/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// These requests delete a type only if it’s not referenced by other entities.
// OAuth2 Scopes: manage_types:{projectKey}
func (svc *Service) TypeDeleteByKey(key string, version int) (result *Type, err error) {
	endpoint := fmt.Sprintf("types/key=%s", key)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
