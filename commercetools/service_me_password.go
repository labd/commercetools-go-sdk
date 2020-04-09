// Automatically generated, do not edit

package commercetools

// MePasswordURLPath is the commercetools API path.
const MePasswordURLPath = "me/password"

// MePasswordCreate creates a new instance of type
func (client *Client) MePasswordCreate() (err error) {
	err = client.Create(MePasswordURLPath, nil, nil, nil)
	return err
}
