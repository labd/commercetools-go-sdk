// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// InventoryEntryCreate creates a new instance of type InventoryEntry
func (client *Client) InventoryEntryCreate(ctx context.Context, draft *InventoryEntryDraft, opts ...RequestOption) (result *InventoryEntry, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "inventory"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InventoryEntryQuery allows querying for type InventoryEntry
func (client *Client) InventoryEntryQuery(ctx context.Context, input *QueryInput) (result *InventoryPagedQueryResponse, err error) {
	endpoint := "inventory"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InventoryEntryDeleteWithID for type InventoryEntry
func (client *Client) InventoryEntryDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *InventoryEntry, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("inventory/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InventoryEntryGetWithID for type InventoryEntry
func (client *Client) InventoryEntryGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *InventoryEntry, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("inventory/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InventoryEntryUpdateWithIDInput is input for function InventoryEntryUpdateWithID
type InventoryEntryUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []InventoryEntryUpdateAction
}

func (input *InventoryEntryUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// InventoryEntryUpdateWithID for type InventoryEntry
func (client *Client) InventoryEntryUpdateWithID(ctx context.Context, input *InventoryEntryUpdateWithIDInput, opts ...RequestOption) (result *InventoryEntry, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("inventory/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
