// Automatically generated, do not edit

package commercetools

// CustomersPasswordTokenURLPath is the commercetools API path.
const CustomersPasswordTokenURLPath = "customers/password-token"

// CustomersPasswordTokenCreate creates a new instance of type CustomerToken
func (client *Client) CustomersPasswordTokenCreate(draft *CustomerCreatePasswordResetToken) (result *CustomerToken, err error) {
	err = client.Create(CustomersPasswordTokenURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
