// Automatically generated, do not edit

package commercetools

// MeEmailConfirmURLPath is the commercetools API path.
const MeEmailConfirmURLPath = "me/email/confirm"

// MeEmailConfirmCreate creates a new instance of type
func (client *Client) MeEmailConfirmCreate() (err error) {
	err = client.Create(MeEmailConfirmURLPath, nil, nil, nil)
	return err
}
