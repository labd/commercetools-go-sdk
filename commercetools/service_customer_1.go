// Automatically generated, do not edit

package commercetools

// CustomerSignInResultURLPath is the commercetools API path.
const CustomerSignInResultURLPath = "login"

// CustomerSignInResultCreate creates a new instance of type CustomerSignInResult
func (client *Client) CustomerSignInResultCreate(draft *CustomerSignin) (result *CustomerSignInResult, err error) {
	err = client.Create(CustomerSignInResultURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
