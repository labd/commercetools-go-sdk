// Automatically generated, do not edit

package commercetools

// ProductDiscountsMatchingURLPath is the commercetools API path.
const ProductDiscountsMatchingURLPath = "product-discounts/matching"

// ProductDiscountsMatchingCreate creates a new instance of type
func (client *Client) ProductDiscountsMatchingCreate(draft *ProductDiscountMatchQuery) (err error) {
	err = client.Create(ProductDiscountsMatchingURLPath, nil, draft, nil)
	return err
}
