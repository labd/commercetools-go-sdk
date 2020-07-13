// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// InventoryEntryURLPath is the commercetools API path.
const InventoryEntryURLPath = "inventory"

// InventoryEntryCreate creates a new instance of type InventoryEntry
func (client *Client) InventoryEntryCreate(ctx context.Context, draft *InventoryEntryDraft) (result *InventoryEntry, err error) {
	err = client.Create(ctx, InventoryEntryURLPath, nil, draft, &result)
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
func (client *Client) InventoryEntryDeleteWithID(ctx context.Context, ID string, version int) (result *InventoryEntry, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("inventory/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InventoryEntryGetWithID for type InventoryEntry
func (client *Client) InventoryEntryGetWithID(ctx context.Context, ID string) (result *InventoryEntry, err error) {
	err = client.Get(ctx, strings.Replace("inventory/{ID}", "{ID}", ID, 1), nil, &result)
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
func (client *Client) InventoryEntryUpdateWithID(ctx context.Context, input *InventoryEntryUpdateWithIDInput) (result *InventoryEntry, err error) {
	err = client.Update(ctx, strings.Replace("inventory/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
