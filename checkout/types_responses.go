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
	case "no_allowed_origins":
		obj := NoAllowedOrigins{}
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
	case "seller_deactivated":
		obj := SellerIsDeactivated{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "seller_not_found":
		obj := SellerNotFound{}
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
	}
	return nil, nil
}

/**
*	Generated when the configuration used to initialize Adyen contains at least one invalid field.
*
 */
type AdyenBadConfig struct {
	// \`error:psp:bad_config`
	Type string `json:"type"`
	// Some fields are invalid
	Message string `json:"message"`
	// An object containing `PSP` and an array of objects with the invalid fields.
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
*	Generated when Adyen cannot be initialized.
*
 */
type AdyenInitError struct {
	// \`error:psp:bad_config`
	Type string `json:"type"`
	// Could not initialise Adyen
	Message string `json:"message"`
	// An object containing `PSP`, `clientKey`, and `enviroment`.
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
*	Generated when a timeout error occurs while initializing Adyen.
*
 */
type AdyenTimeout struct {
	// \`error:psp:bad_config`
	Type string `json:"type"`
	// Timeout while initialising Adyen
	Message string `json:"message"`
	// An object containing `PSP`, `clientKey`, and `enviroment`.
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
*	Generated when the requested application is deactivated. Activate the application in the Merchant Center to continue.
*
 */
type ApplicationDeactivated struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// Application `{appId}` for seller `{sellerId}` is disabled
	Message string `json:"message"`
	// An object containing `sellerId` and `appId`.
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
*	Generated when the requested application is not found. The application may have been deleted or its configuration is incorrect.
*
 */
type ApplicationNotFound struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// Application `{appId}` for seller `{sellerId}` not found
	Message string `json:"message"`
	// An object containing `sellerId` and `appId`.
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
*	Generated when the Cart was emptied during the checkout. It is not possible to recover from this, the customer must restart the checkout process.
*
 */
type CartEmptiedDuringCheckout struct {
	// \`error:cart:empty`
	Type string `json:"type"`
	// Cart `{cartId}` was emptied during checkout
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
*	Generated when the Cart for the current checkout is empty. The Cart must contain at least one Line Item.
*
 */
type CartEmpty struct {
	// \`error:cart:empty`
	Type string `json:"type"`
	// Cart `{cartId}` is empty
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
*	Generated when the Cart is not found. A valid Cart with at least one Line Item is required to start the checkout.
*
 */
type CartNotFound struct {
	// \`error:cart:not_found`
	Type string `json:"type"`
	// Cart `{cartId}` for seller `{sellerId}` and application `{appId}` not found
	Message string `json:"message"`
	// An object containing `sellerId`, `appId`, and `cartId`.
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
*	Generated when the customer cancels the checkout.
*
 */
type CheckoutCancelled struct {
	// \`info:app:status`
	Type string `json:"type"`
	// Checkout cancelled
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
*	Generated when the customer completes the checkout.
*
 */
type CheckoutCompleted struct {
	// \`info:app:status`
	Type string `json:"type"`
	// Checkout for `{orderId}` completed
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
*	Generated when the checkout was loaded successfully and it is waiting for the configuration parameters passed through the `checkoutConfig` object.
*
 */
type CheckoutLoaded struct {
	// \`info:app:status`
	Type string `json:"type"`
	// Checkout loaded
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
*	Generated when the checkout receives the configuration parameters passed through the `checkoutConfig` object and starts successfully.
*
 */
type CheckoutStarted struct {
	// \`info:app:status`
	Type string `json:"type"`
	// Checkout started
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
*	Generated when the checkout has not received the configuration parameters passed through the `checkoutConfig` object on time.
*
 */
type InitTimeout struct {
	// \`error:init:timeout`
	Type string `json:"type"`
	// Timeout error, no init message received
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
*	Generated when the `checkoutConfig` object contains one or more invalid fields.
*
 */
type BadInputData struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// Some fields are invalid
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
	// The token is invalid
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
*	Generated when there are no allowed origins configured for the current application. Add at least one **Origin URL** in your [application settings in the Merchant Center](/configuring-checkout#applications).
*
 */
type NoAllowedOrigins struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// There are no allowed origins configured
	Message string `json:"message"`
	// An object containing additional data about the event.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NoAllowedOrigins) MarshalJSON() ([]byte, error) {
	type Alias NoAllowedOrigins
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "no_allowed_origins", Alias: (*Alias)(&obj)})
}

/**
*	Generated when there are no payment methods available. Add at least one **Payment method** in your [application settings in the Merchant Center](/configuring-checkout#payment-connector).
*
 */
type NoPaymentMethods struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// There are no payment methods configured
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
*	Generated when no Shipping Method is available for the shipping address of the Cart. This may indicate an incomplete configuration.
*
 */
type NoShippingMethods struct {
	// \`warn:init:bad_config`
	Type string `json:"type"`
	// There are no shipping methods matching cart
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
*	Generated when the checkout successfully creates an Order.
*
 */
type OrderCreated struct {
	// \`info:order:status`
	Type string `json:"type"`
	// Order `{orderId}` created
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
*	Generated when the seller is deactivated and the checkout cannot be initialized. Contact commercetools support.
*
 */
type SellerIsDeactivated struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// Seller `{sellerId}` is deactivated
	Message string `json:"message"`
	// An object containing `sellerId`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SellerIsDeactivated) MarshalJSON() ([]byte, error) {
	type Alias SellerIsDeactivated
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "seller_deactivated", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the seller is not found. The seller may have been deleted or the configuration may not be correct. Contact commercetools support.
*
 */
type SellerNotFound struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// Seller `{sellerId}` not found
	Message string `json:"message"`
	// An object containing `sellerId`.
	Payload *interface{} `json:"payload,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SellerNotFound) MarshalJSON() ([]byte, error) {
	type Alias SellerNotFound
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "seller_not_found", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the checkout is initialised with [`skipShipping` set to `true`](/installing-checkout#placeholder-values). You must populate the `shippingAddress` field of the Cart with at least the shipping country.
*
 */
type ShippingAddressMissing struct {
	// \`error:cart:field_required`
	Type string `json:"type"`
	// The shippingAddress field is missing for cart `{cartId}`
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
*	Generated when the currently used origin URL is not in the list of the [**Origin URLs** configured in the Merchant Centre](/configuring-checkout#create-an-application) and the checkout cannot be initialized. Add the **Origin URL** in your application settings in the Merchant Center.
*
 */
type UnallowedOrigin struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// Unallowed origin `{origin}`, allowed origins `{allowedOrigins}`
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
*	Generated when the provided locale is not [available for localization](/installing-checkout#locales). The checkout falls back to English.
*
 */
type UnavailableLocale struct {
	// \`error:init:bad_config`
	Type string `json:"type"`
	// The provided locale `{locale}` is not available for translated definitions
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
