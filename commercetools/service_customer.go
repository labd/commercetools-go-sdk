// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// CustomerCreate creates a new instance of type Customer
func (client *Client) CustomerCreate(ctx context.Context, draft *CustomerDraft, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "customers"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerQuery allows querying for type Customer
func (client *Client) CustomerQuery(ctx context.Context, input *QueryInput) (result *CustomerPagedQueryResponse, err error) {
	endpoint := "customers"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerDeleteWithID for type Customer
func (client *Client) CustomerDeleteWithID(ctx context.Context, id string, version int, dataErasure bool, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("customers/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerDeleteWithKey for type Customer
func (client *Client) CustomerDeleteWithKey(ctx context.Context, key string, version int, dataErasure bool, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("customers/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGetWithEmailToken for type Customer
func (client *Client) CustomerGetWithEmailToken(ctx context.Context, emailToken string, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("customers/email-token=%s", emailToken)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGetWithID for type Customer
func (client *Client) CustomerGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("customers/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGetWithKey for type Customer
func (client *Client) CustomerGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("customers/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGetWithPasswordToken for type Customer
func (client *Client) CustomerGetWithPasswordToken(ctx context.Context, passwordToken string, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("customers/password-token=%s", passwordToken)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerUpdateWithIDInput is input for function CustomerUpdateWithID
type CustomerUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CustomerUpdateAction
}

func (input *CustomerUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// CustomerUpdateWithID for type Customer
func (client *Client) CustomerUpdateWithID(ctx context.Context, input *CustomerUpdateWithIDInput, opts ...RequestOption) (result *Customer, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("customers/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerUpdateWithKeyInput is input for function CustomerUpdateWithKey
type CustomerUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CustomerUpdateAction
}

func (input *CustomerUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// CustomerUpdateWithKey for type Customer
func (client *Client) CustomerUpdateWithKey(ctx context.Context, input *CustomerUpdateWithKeyInput, opts ...RequestOption) (result *Customer, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("customers/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
CustomerEmailToken To verify a customer's email, an email token can be created. This should be embedded in a link and sent to the
customer via email. When the customer clicks on the link, the "verify customer's email" endpoint should be called,
which sets customer's isVerifiedEmail field to true.
*/
func (client *Client) CustomerEmailToken(ctx context.Context, value *CustomerCreateEmailToken, opts ...RequestOption) (result *CustomerToken, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "customers/email-token"
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerEmailconfirm for type CustomerEmailVerify
func (client *Client) CustomerEmailconfirm(ctx context.Context, value *CustomerEmailVerify, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "customers/email/confirm"
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerPassword for type CustomerChangePassword
func (client *Client) CustomerPassword(ctx context.Context, value *CustomerChangePassword, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "customers/password"
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
CustomerPasswordToken The following workflow can be used to reset the customer's password:

* Create a password reset token and send it embedded in a link to the customer.
* When the customer clicks on the link, the customer is retrieved with the token.
* The customer enters a new password and the "reset customer's password" endpoint is called.
*/
func (client *Client) CustomerPasswordToken(ctx context.Context, value *CustomerCreatePasswordResetToken, opts ...RequestOption) (result *CustomerToken, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "customers/password-token"
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerPasswordreset for type CustomerResetPassword
func (client *Client) CustomerPasswordreset(ctx context.Context, value *CustomerResetPassword, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "customers/password/reset"
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
