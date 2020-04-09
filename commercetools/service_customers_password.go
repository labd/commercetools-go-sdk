// Automatically generated, do not edit

package commercetools

// CustomersPasswordURLPath is the commercetools API path.
const CustomersPasswordURLPath = "customers/password"

// CustomersPasswordCreate creates a new instance of type
func (client *Client) CustomersPasswordCreate(draft *CustomerChangePassword) (err error) {
	err = client.Create(CustomersPasswordURLPath, nil, draft, nil)
	return err
}
