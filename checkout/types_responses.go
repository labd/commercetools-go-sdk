package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

type ResponseMessage interface{}

func mapDiscriminatorResponseMessage(input interface{}) (ResponseMessage, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["code"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'code'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "adyen_bad_config":
		obj := AdyenBadConfig{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "adyen_init_error":
		obj := AdyenInitError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "adyen_timeout":
		obj := AdyenTimeout{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "application_disabled":
		obj := ApplicationDeactivated{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "application_not_found":
		obj := ApplicationNotFound{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "cart_emptied_during_checkout":
		obj := CartEmptiedDuringCheckout{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "cart_empty":
		obj := CartEmpty{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "cart_not_found":
		obj := CartNotFound{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "checkout_cancelled":
		obj := CheckoutCancelled{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "checkout_completed":
		obj := CheckoutCompleted{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "checkout_loaded":
		obj := CheckoutLoaded{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "checkout_started":
		obj := CheckoutStarted{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "init_timeout":
		obj := InitTimeout{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "invalid_fields":
		obj := BadInputData{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "invalid_token":
		obj := InvalidToken{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "no_payment_methods":
		obj := NoPaymentMethods{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "no_shipping_methods":
		obj := NoShippingMethods{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order_created":
		obj := OrderCreated{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "project_deactivated":
		obj := ProjectIsDeactivated{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "shipping_address_missing":
		obj := ShippingAddressMissing{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "unallowed_origin":
		obj := UnallowedOrigin{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "unavailable_locale":
		obj := UnavailableLocale{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "deprecated_fields":
		obj := DeprecatedFields{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order_creation_error":
		obj := OrderCreationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "cart_with_exisiting_payment":
		obj := CartWithExistingPayment{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Generated when the configuration to initialize the [Adyen payment connector](/configuring-adyen) contains at least one invalid field.
*
 */
type AdyenBadConfig struct {
	// \`error:psp:bad_config`
	Type string `json:"type"`
	// Some fields are invalid.
	Message string `json:"message"`
	// An object containing `psp` and an array of objects with the invalid fields.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AdyenBadConfig) MarshalJSON() ([]byte, error) {
	type Alias AdyenBadConfig
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "adyen_bad_config", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the [Adyen payment connector](/configuring-adyen) cannot be initialized.
*
 */
type AdyenInitError struct {
	// \`error:psp:bad_config`
	Type string `json:"type"`
	// Could not initialise Adyen.
	Message string `json:"message"`
	// An object containing `psp`, `clientKey`, and `enviroment`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AdyenInitError) MarshalJSON() ([]byte, error) {
	type Alias AdyenInitError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "adyen_init_error", Alias: (*Alias)(&obj)})
}

func (obj AdyenInitError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AdyenInitError: failed to parse error response"
}

/**
*	Generated when a timeout error occurs while initializing the [Adyen payment connector](/configuring-adyen).
*
 */
type AdyenTimeout struct {
	// \`error:psp:bad_config`
	Type string `json:"type"`
	// Timeout while initialising Adyen.
	Message string `json:"message"`
	// An object containing `psp`, `clientKey`, and `enviroment`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AdyenTimeout) MarshalJSON() ([]byte, error) {
	type Alias AdyenTimeout
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "adyen_timeout", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the requested [application](/configuring-checkout#applications) is deactivated. Activate the application in the Merchant Center to continue.
*
 */
type ApplicationDeactivated struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// Application `{applicationKey}` for commercetools Checkout `{projectKey}` is disabled.
	Message string `json:"message"`
	// An object containing `projectKey` and `applicationKey`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApplicationDeactivated) MarshalJSON() ([]byte, error) {
	type Alias ApplicationDeactivated
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "application_disabled", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the requested [application](/configuring-checkout#applications) is not found. The application may have been deleted or its configuration is incorrect.
*
 */
type ApplicationNotFound struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// Application `{applicationKey}` for commercetools Checkout `{projectKey}` not found.
	Message string `json:"message"`
	// An object containing `projectKey` and `applicationKey`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApplicationNotFound) MarshalJSON() ([]byte, error) {
	type Alias ApplicationNotFound
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "application_not_found", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the [Cart](/../api/projects/carts) was emptied during the checkout process. It is not possible to recover from this, the customer must restart the checkout process.
*
 */
type CartEmptiedDuringCheckout struct {
	// \`error:cart:empty`
	Type string `json:"type"`
	// Cart `{cartId}` was emptied during checkout.
	Message string `json:"message"`
	// An object containing `cartId`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartEmptiedDuringCheckout) MarshalJSON() ([]byte, error) {
	type Alias CartEmptiedDuringCheckout
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "cart_emptied_during_checkout", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the [Cart](/../api/projects/carts) for the current checkout is empty. The Cart must contain at least one [Line Item](/../api/carts-orders-overview#line-items).
*
 */
type CartEmpty struct {
	// \`error:cart:empty`
	Type string `json:"type"`
	// Cart `{cartId}` is empty.
	Message string `json:"message"`
	// An object containing `cartId`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartEmpty) MarshalJSON() ([]byte, error) {
	type Alias CartEmpty
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "cart_empty", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the [Cart](/../api/projects/carts) is not found. To start the checkout process, a valid Cart with at least one [Line Item](/../api/carts-orders-overview#line-items) is required.
*
 */
type CartNotFound struct {
	// \`error:cart:not_found`
	Type string `json:"type"`
	// Cart `{cartId}` for commercetools Checkout `{projectKey}` and application `{applicationKey}` not found.
	Message string `json:"message"`
	// An object containing `projectKey`, `applicationKey`, and `cartId`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartNotFound) MarshalJSON() ([]byte, error) {
	type Alias CartNotFound
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "cart_not_found", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the customer cancels the checkout process.
*
 */
type CheckoutCancelled struct {
	// \`info:app:status`
	Type string `json:"type"`
	// Checkout cancelled.
	Message string `json:"message"`
	// An object containing additional data about the event.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutCancelled) MarshalJSON() ([]byte, error) {
	type Alias CheckoutCancelled
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "checkout_cancelled", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the customer completes the checkout process.
*
 */
type CheckoutCompleted struct {
	// \`info:app:status`
	Type string `json:"type"`
	// Checkout for `{orderId}` completed.
	Message string `json:"message"`
	// An object containing `orderId`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutCompleted) MarshalJSON() ([]byte, error) {
	type Alias CheckoutCompleted
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "checkout_completed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when Checkout has been loaded successfully and is waiting for the configuration parameters passed in the `checkoutConfig` [object](/sdk).
*
 */
type CheckoutLoaded struct {
	// \`info:app:status`
	Type string `json:"type"`
	// Checkout loaded.
	Message string `json:"message"`
	// An object containing additional data about the event.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutLoaded) MarshalJSON() ([]byte, error) {
	type Alias CheckoutLoaded
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "checkout_loaded", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the configuration parameters are passed successfully through the `checkoutConfig` [object](/sdk) and the checkout process starts.
*
 */
type CheckoutStarted struct {
	// \`info:app:status`
	Type string `json:"type"`
	// Checkout started.
	Message string `json:"message"`
	// An object containing additional data about the event.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutStarted) MarshalJSON() ([]byte, error) {
	type Alias CheckoutStarted
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "checkout_started", Alias: (*Alias)(&obj)})
}

/**
*	Generated when Checkout has not received the configuration parameters passed through the `checkoutConfig` [object](/sdk) on time.
*
 */
type InitTimeout struct {
	// \`error:init:timeout`
	Type string `json:"type"`
	// Timeout error, no init message received.
	Message string `json:"message"`
	// An object containing additional data about the event.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InitTimeout) MarshalJSON() ([]byte, error) {
	type Alias InitTimeout
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "init_timeout", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the `checkoutConfig` [object](/sdk) contains one or more invalid fields.
*
 */
type BadInputData struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// Some fields are invalid.
	Message string `json:"message"`
	// An object containing an array of `field`, `message`, and `invalidValue`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BadInputData) MarshalJSON() ([]byte, error) {
	type Alias BadInputData
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "invalid_fields", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the `accessToken` is invalid.
*
 */
type InvalidToken struct {
	// \`error:token:invalid`
	Type string `json:"type"`
	// The token is invalid.
	Message string `json:"message"`
	// An object containing additional data about the event.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidToken) MarshalJSON() ([]byte, error) {
	type Alias InvalidToken
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "invalid_token", Alias: (*Alias)(&obj)})
}

/**
*	Generated when no payment method is set up for an [application](/configuring-checkout#applications). Add at least one **Payment method** to the application in the Merchant Center.
*
 */
type NoPaymentMethods struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// There are no payment methods configured.
	Message string `json:"message"`
	// An object containing additional data about the event.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NoPaymentMethods) MarshalJSON() ([]byte, error) {
	type Alias NoPaymentMethods
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "no_payment_methods", Alias: (*Alias)(&obj)})
}

/**
*	Generated when no [Shipping Method](/../api/projects/shippingMethods) is available for the shipping address of the [Cart](/../api/projects/carts). This may indicate an incomplete configuration.
*
 */
type NoShippingMethods struct {
	// \`warn:init:bad_config`
	Type string `json:"type"`
	// There are no shipping methods matching cart.
	Message string `json:"message"`
	// An object containing `cartId`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NoShippingMethods) MarshalJSON() ([]byte, error) {
	type Alias NoShippingMethods
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "no_shipping_methods", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an [Order](/../api/projects/orders) is created after a successful checkout process.
*
 */
type OrderCreated struct {
	// \`info:order:status`
	Type string `json:"type"`
	// Order `{orderId}` created.
	Message string `json:"message"`
	// An object containing `orderId`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCreated) MarshalJSON() ([]byte, error) {
	type Alias OrderCreated
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "order_created", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the commercetools Checkout [`projectKey`](/sdk) is deactivated and cannot be initialized. Contact support via the [Support Portal](https://commercetools.atlassian.net/servicedesk/customer/portal/30).
*
 */
type ProjectIsDeactivated struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// commercetools Checkout `{projectKey}` is deactivated.
	Message string `json:"message"`
	// An object containing `projectKey`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectIsDeactivated) MarshalJSON() ([]byte, error) {
	type Alias ProjectIsDeactivated
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "project_deactivated", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the checkout is initialised with [`skipShipping` set to `true`](/installing-checkout#placeholder-values). You must populate the `shippingAddress` field of the Cart with at least the shipping country.
*
 */
type ShippingAddressMissing struct {
	// \`error:cart:field_required`
	Type string `json:"type"`
	// The shippingAddress field is missing for cart `{cartId}`.
	Message string `json:"message"`
	// An object containing `cartId`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingAddressMissing) MarshalJSON() ([]byte, error) {
	type Alias ShippingAddressMissing
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "shipping_address_missing", Alias: (*Alias)(&obj)})
}

/**
*	Generated when Checkout cannot be initialized because the URL that is trying to initialize it is not in the list of the allowed URLs for the [application](/configuring-checkout#applications). Add the URL to the **Origin URLs** list in your application settings in the Merchant Center.
*
 */
type UnallowedOrigin struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// Unallowed origin `{origin}`, allowed origins `{allowedOrigins}`.
	Message string `json:"message"`
	// An object containing `origin` and `allowedOrigins`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj UnallowedOrigin) MarshalJSON() ([]byte, error) {
	type Alias UnallowedOrigin
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "unallowed_origin", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the provided `locale` is not [available for localization](/installing-checkout#locales). The localization falls back to English.
*
 */
type UnavailableLocale struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// The provided locale `{locale}` is not available for translated definitions.
	Message string `json:"message"`
	// An object containing `locale`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj UnavailableLocale) MarshalJSON() ([]byte, error) {
	type Alias UnavailableLocale
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "unavailable_locale", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the `checkoutConfig` [object](/sdk) contains one or more deprecated fields.
*
 */
type DeprecatedFields struct {
	// \`warn:init:bad_config`
	Type string `json:"type"`
	// Some fields are deprecated.
	Message string `json:"message"`
	// An object containing an array of `field` and `message`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeprecatedFields) MarshalJSON() ([]byte, error) {
	type Alias DeprecatedFields
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "deprecated_fields", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an [Order](/../api/projects/orders) that references an approved [Payment](/../api/projects/payments) cannot be created.
*
 */
type OrderCreationError struct {
	// \`error:order:failed`
	Type string `json:"type"`
	// Order creation failed with approved payment.
	Message string `json:"message"`
	// An object containing `cartId` and `errors`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCreationError) MarshalJSON() ([]byte, error) {
	type Alias OrderCreationError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "order_creation_error", Alias: (*Alias)(&obj)})
}

func (obj OrderCreationError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown OrderCreationError: failed to parse error response"
}

/**
*	Generated when trying to add a [Payment](/../api/projects/payments) to a [Cart](/../api/projects/carts) that already references an approved Payment.
*
 */
type CartWithExistingPayment struct {
	// \`error:payment:failed`
	Type string `json:"type"`
	// Cart with existing approved payment.
	Message string `json:"message"`
	// An object containing `cartId`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartWithExistingPayment) MarshalJSON() ([]byte, error) {
	type Alias CartWithExistingPayment
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "cart_with_exisiting_payment", Alias: (*Alias)(&obj)})
}
