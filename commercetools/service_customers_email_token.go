// Automatically generated, do not edit

package commercetools

// CustomersEmailTokenURLPath is the commercetools API path.
const CustomersEmailTokenURLPath = "customers/email-token"

// CustomersEmailTokenCreate creates a new instance of type
func (client *Client) CustomersEmailTokenCreate(draft *CustomerCreateEmailToken) (err error) {
	err = client.Create(CustomersEmailTokenURLPath, nil, draft, nil)
	return err
}
