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

func (svc *Service) TypeGetByID(id string) (result *Type, err error) {
	err = svc.client.Get(fmt.Sprintf("types/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *Service) TypeCreate(draft *TypeDraft) (result *Type, err error) {
	err = svc.client.Create("types", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *Service) TypeUpdate(input *TypeUpdateInput) (result *Type, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("No valid type id passed")
	}

	endpoint := fmt.Sprintf("types/%s", input.ID)
	err = svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

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
