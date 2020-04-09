// Automatically generated, do not edit

package commercetools

// MeLoginURLPath is the commercetools API path.
const MeLoginURLPath = "me/login"

// MeLoginCreate creates a new instance of type
func (client *Client) MeLoginCreate() (err error) {
	err = client.Create(MeLoginURLPath, nil, nil, nil)
	return err
}
