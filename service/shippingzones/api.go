package shippingzones

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// DeleteInput provides the data required to delete a shipping zone.
type DeleteInput struct {
	ID      string
	Version int
}

// UpdateInput provides the data required to update a shipping zone.
type UpdateInput struct {
	ID string

	// The expected version of the type on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict
	// will be returned.
	Version int

	// The list of update actions to be performed on the type.
	Actions commercetools.UpdateActions
}

// GetByID will return a shipping zone matching the provided ID. OAuth2 Scopes:
// view_products:{projectKey}
func (svc *Service) GetByID(id string) (result *ShippingZone, err error) {
	err = svc.client.Get(fmt.Sprintf("zones/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create will create a new shipping zone from a draft, and return the newly
// created shipping zone. OAuth2 Scopes: manage_products:{projectKey}
func (svc *Service) Create(draft *ShippingZoneDraft) (result *ShippingZone, err error) {
	err = svc.client.Create("zones", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update will update a shipping zone matching the provided ID with the defined
// UpdateActions. OAuth2 Scopes: manage_products:{projectKey}
func (svc *Service) Update(input *UpdateInput) (result *ShippingZone, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("zones/%s", input.ID)
	err = svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByID will delete a shipping zone matching the provided ID. OAuth2
// Scopes: manage_products:{projectKey}
func (svc *Service) DeleteByID(id string, version int) (result *ShippingZone, err error) {
	endpoint := fmt.Sprintf("zones/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
