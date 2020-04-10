// Automatically generated, do not edit

package commercetools

// LoginURLPath is the commercetools API path.
const LoginURLPath = "login"

// LoginCreate creates a new instance of type CustomerSignInResult
func (client *Client) LoginCreate(draft *CustomerSignin) (result *CustomerSignInResult, err error) {
	err = client.Create(LoginURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
