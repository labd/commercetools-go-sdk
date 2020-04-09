// Automatically generated, do not edit

package commercetools

// MePasswordResetURLPath is the commercetools API path.
const MePasswordResetURLPath = "me/password/reset"

// MePasswordResetCreate creates a new instance of type
func (client *Client) MePasswordResetCreate() (err error) {
	err = client.Create(MePasswordResetURLPath, nil, nil, nil)
	return err
}
