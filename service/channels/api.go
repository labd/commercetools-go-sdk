package channels

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// DeleteInput provides the data required to delete a channel.
type DeleteInput struct {
	ID      string
	Version int
}

// UpdateInput provides the data required to update a channel.
type UpdateInput struct {
	ID string

	// The expected version of the channel on which the changes should be
	// applied. If the expected version does not match the actual version, a 409
	// Conflict will be returned.
	Version int

	// The list of update actions to be performed on the channel.
	Actions commercetools.UpdateActions
}

// GetByID will return a channel matching the provided ID. OAuth2 Scopes:
// view_products:{projectKey}
func (svc *Service) GetByID(id string) (result *Channel, err error) {
	err = svc.client.Get(fmt.Sprintf("channels/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create will create a new channel from a draft, and return the newly created
// channel. OAuth2 Scopes: manage_products:{projectKey}
func (svc *Service) Create(draft *ChannelDraft) (result *Channel, err error) {
	err = svc.client.Create("channels", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update will update a channel matching the provided ID with the defined
// UpdateActions. OAuth2 Scopes: manage_products:{projectKey}
func (svc *Service) Update(input *UpdateInput) (result *Channel, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("channels/%s", input.ID)
	err = svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete will delete a type matching the provided ID. These requests delete a
// type only if it’s not referenced by other entities. OAuth2 Scopes:
// manage_types:{projectKey}
func (svc *Service) Delete(id string, version int) (result *Channel, err error) {
	endpoint := fmt.Sprintf("channels/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
