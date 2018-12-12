package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// ZoneDeleteInput provides the data required to delete a shipping zone.
type ZoneDeleteInput struct {
	ID      string
	Version int
}

// ZoneUpdateInput provides the data required to update a shipping zone.
type ZoneUpdateInput struct {
	ID string

	// The expected version of the type on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict
	// will be returned.
	Version int

	// The list of update actions to be performed on the type.
	Actions []ZoneUpdateAction
}

// ZoneGetByID will return a shipping zone matching the provided ID. OAuth2 Scopes:
// view_products:{projectKey}
func (client *Client) ZoneGetByID(id string) (result *Zone, err error) {
	err = client.Get(fmt.Sprintf("zones/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneCreate will create a new shipping zone from a draft, and return the newly
// created shipping zone. OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) ZoneCreate(draft *ZoneDraft) (result *Zone, err error) {
	err = client.Create("zones", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneUpdate will update a shipping zone matching the provided ID with the defined
// ZoneUpdateActions. OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) ZoneUpdate(input *ZoneUpdateInput) (result *Zone, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("zones/%s", input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneDeleteByID will delete a shipping zone matching the provided ID. OAuth2
// Scopes: manage_products:{projectKey}
func (client *Client) ZoneDeleteByID(id string, version int) (result *Zone, err error) {
	endpoint := fmt.Sprintf("zones/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
