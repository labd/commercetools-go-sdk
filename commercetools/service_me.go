// Automatically generated, do not edit

package commercetools

// MyCustomerURLPath is the commercetools API path.
const MyCustomerURLPath = "me"

// MyCustomerCreate creates a new instance of type MyCustomer
func (client *Client) MyCustomerCreate(draft *Update) (result *MyCustomer, err error) {
	err = client.Create(MyCustomerURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyCustomerQuery allows querying for type MyCustomer
func (client *Client) MyCustomerQuery(input *QueryInput) (result *MyCustomer, err error) {
	err = client.Query(MyCustomerURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
