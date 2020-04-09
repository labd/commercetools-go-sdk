// Automatically generated, do not edit

package commercetools

// CartsReplicateURLPath is the commercetools API path.
const CartsReplicateURLPath = "carts/replicate"

// CartsReplicateCreate creates a new instance of type Cart
func (client *Client) CartsReplicateCreate(draft *ReplicaCartDraft) (result *Cart, err error) {
	err = client.Create(CartsReplicateURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
