// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// MeShoppingListsURLPath is the commercetools API path.
const MeShoppingListsURLPath = "me/shopping-lists"

// MeShoppingListsCreate creates a new instance of type MyShoppingList
func (client *Client) MeShoppingListsCreate(draft *MyShoppingListDraft) (result *MyShoppingList, err error) {
	err = client.Create(MeShoppingListsURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MeShoppingListsQuery allows querying for type MyShoppingList
func (client *Client) MeShoppingListsQuery(input *QueryInput) (result *ShoppingListPagedQueryResponse, err error) {
	err = client.Query(MeShoppingListsURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MeShoppingListsDeleteWithID for type MyShoppingList
func (client *Client) MeShoppingListsDeleteWithID(ID string, version int) (result *MyShoppingList, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("me/shopping-lists/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MeShoppingListsGetWithID for type MyShoppingList
func (client *Client) MeShoppingListsGetWithID(ID string) (result *MyShoppingList, err error) {
	err = client.Get(strings.Replace("me/shopping-lists/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyShoppingListUpdateWithIDInput is input for function MeShoppingListsUpdateWithID
type MyShoppingListUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []MyShoppingListUpdateAction
}

// MeShoppingListsUpdateWithID for type MyShoppingList
func (client *Client) MeShoppingListsUpdateWithID(input *MyShoppingListUpdateWithIDInput) (result *MyShoppingList, err error) {
	err = client.Update(strings.Replace("me/shopping-lists/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
