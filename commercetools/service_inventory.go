// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// InventoryEntryURLPath is the commercetools API path.
const InventoryEntryURLPath = "inventory"

func (client *Client) InventoryEntryCreate(draft *InventoryEntryDraft) (result *InventoryEntry, err error) {
	err = client.Create(InventoryEntryURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) InventoryEntryQuery(input *QueryInput) (result *InventoryPagedQueryResponse, err error) {
	err = client.Query(InventoryEntryURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) InventoryEntryDeleteWithId(ID string, version int) (result *InventoryEntry, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("inventory/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) InventoryEntryGetWithId(ID string) (result *InventoryEntry, err error) {
	err = client.Get(strings.Replace("inventory/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type InventoryEntryUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []InventoryUpdateAction
}

func (client *Client) InventoryEntryUpdateWithId(input *InventoryEntryUpdateWithIdInput) (result *InventoryEntry, err error) {
	err = client.Update(strings.Replace("inventory/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
