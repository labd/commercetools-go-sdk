// Automatically generated, do not edit

package commercetools

// MeActiveCartURLPath is the commercetools API path.
const MeActiveCartURLPath = "me/active-cart"

// MeActiveCartQuery allows querying for type MyCart
func (client *Client) MeActiveCartQuery(input *QueryInput) (result *MyCart, err error) {
	err = client.Query(MeActiveCartURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
