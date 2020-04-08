// Automatically generated, do not edit

package commercetools

// CustomersPasswordResetURLPath is the commercetools API path.
const CustomersPasswordResetURLPath = "customers/password/reset"

// CustomersPasswordResetCreate creates a new instance of type
func (client *Client) CustomersPasswordResetCreate(draft *CustomerResetPassword) (err error) {
	err = client.Create(CustomersPasswordResetURLPath, nil, draft, nil)
	return err
}
