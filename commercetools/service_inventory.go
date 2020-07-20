// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// InventoryEntryURLPath is the commercetools API path.
const InventoryEntryURLPath = "inventory"

// InventoryEntryCreate creates a new instance of type InventoryEntry
func (client *Client) InventoryEntryCreate(ctx context.Context, draft *InventoryEntryDraft, opts ...RequestOption) (result *InventoryEntry, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, InventoryEntryURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InventoryEntryQuery allows querying for type InventoryEntry
func (client *Client) InventoryEntryQuery(ctx context.Context, input *QueryInput) (result *InventoryPagedQueryResponse, err error) {
	err = client.Query(ctx, InventoryEntryURLPath, input.toParams(), &result)
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
	err = client.Delete(ctx, endpoint, params, &result)
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
	err = client.Get(ctx, endpoint, params, &result)
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

// InventoryEntryUpdateWithID for type InventoryEntry
func (client *Client) InventoryEntryUpdateWithID(ctx context.Context, input *InventoryEntryUpdateWithIDInput, opts ...RequestOption) (result *InventoryEntry, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("inventory/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
