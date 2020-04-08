// Automatically generated, do not edit

package commercetools

// OrdersImportURLPath is the commercetools API path.
const OrdersImportURLPath = "orders/import"

// OrdersImportCreate creates a new instance of type Order
func (client *Client) OrdersImportCreate(draft *OrderImportDraft) (result *Order, err error) {
	err = client.Create(OrdersImportURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
