// Automatically generated, do not edit

package commercetools

// CustomersEmailConfirmURLPath is the commercetools API path.
const CustomersEmailConfirmURLPath = "customers/email/confirm"

// CustomersEmailConfirmCreate creates a new instance of type
func (client *Client) CustomersEmailConfirmCreate(draft *CustomerEmailVerify) (err error) {
	err = client.Create(CustomersEmailConfirmURLPath, nil, draft, nil)
	return err
}
