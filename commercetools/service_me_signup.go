// Automatically generated, do not edit

package commercetools

// MeSignupURLPath is the commercetools API path.
const MeSignupURLPath = "me/signup"

// MeSignupCreate creates a new instance of type CustomerSignInResult
func (client *Client) MeSignupCreate(draft *MyCustomerDraft) (result *CustomerSignInResult, err error) {
	err = client.Create(MeSignupURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
