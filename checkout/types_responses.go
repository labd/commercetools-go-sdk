package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

type Message interface{}

func mapDiscriminatorMessage(input interface{}) (Message, error) {
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
	case "checkout_cancelled":
		obj := CheckoutCancelled{}
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
	case "error_loading_all_payment_integrations":
		obj := ErrorLoadingAllPaymentIntegrations{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "expired_session":
		obj := ExpiredSession{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "express_payment_interrupted":
		obj := ExpressPaymentInterrupted{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "external_terms_and_conditions_pending":
		obj := ExternalTermsAndConditionsPending{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "failed_to_refresh_session":
		obj := FailedToRefreshSession{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "gift_card_balance_error":
		obj := GiftCardBalanceError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "gift_card_balance_removed":
		obj := GiftCardBalanceRemoved{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "gift_card_balance_started":
		obj := GiftCardBalanceStarted{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "gift_card_redeem_error":
		obj := GiftCardRedeemError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "gift_card_redeem_started":
		obj := GiftCardRedeemStarted{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "gift_card_redeem_success":
		obj := GiftCardRedeemSuccess{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "init_error":
		obj := InitError{}
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
	case "invalid_mode":
		obj := InvalidMode{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "multiple_vendor_button_containers":
		obj := MultipleVendorButtonContainers{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "no_express_payment_integrations":
		obj := NoExpressPaymentIntegrations{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "no_payment_integrations":
		obj := NoPaymentIntegrations{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order_verification_started":
		obj := OrderVerificationStarted{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order_verification_timeout":
		obj := OrderVerificationTimeout{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_validation_passed":
		obj := PaymentValidationPassed{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_validation_started":
		obj := PaymentValidationStarted{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "shipping_method_does_not_match_cart":
		obj := ShippingMethodDoesNotMatchCart{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Generated when the customer cancels the checkout process.
*
 */
type CheckoutCancelled struct {
	// \`info`
	Severity string `json:"severity"`
	// Checkout cancelled.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
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
*	Generated when Checkout is loaded and waiting for the configuration properties to be passed with the `checkoutFlow` or `paymentFlow` [method](/sdk).
*
 */
type CheckoutLoaded struct {
	// \`info`
	Severity string `json:"severity"`
	// Checkout loaded.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
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
*	Generated when the configuration properties are passed successfully with the `checkoutFlow` or `paymentFlow` [method](/sdk) and the checkout process starts.
*
 */
type CheckoutStarted struct {
	// \`info`
	Severity string `json:"severity"`
	// Checkout started.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
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
*	Generated when the loading of all the payment integrations have failed.
*
 */
type ErrorLoadingAllPaymentIntegrations struct {
	// \`error`
	Severity string `json:"severity"`
	// Error loading all payment integrations.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ErrorLoadingAllPaymentIntegrations) MarshalJSON() ([]byte, error) {
	type Alias ErrorLoadingAllPaymentIntegrations
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "error_loading_all_payment_integrations", Alias: (*Alias)(&obj)})
}

func (obj ErrorLoadingAllPaymentIntegrations) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ErrorLoadingAllPaymentIntegrations: failed to parse error response"
}

/**
*	Generated when the [Checkout Session](/installing-checkout#create-checkout-sessions) is expired.
*
 */
type ExpiredSession struct {
	// \`error`
	Severity string `json:"severity"`
	// Session is expired.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExpiredSession) MarshalJSON() ([]byte, error) {
	type Alias ExpiredSession
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "expired_session", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the Express Payments process is interrupted by the seller's callback.
*
 */
type ExpressPaymentInterrupted struct {
	// \`info`
	Severity string `json:"severity"`
	// Express Payments process was interrupted by a callback.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExpressPaymentInterrupted) MarshalJSON() ([]byte, error) {
	type Alias ExpressPaymentInterrupted
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "express_payment_interrupted", Alias: (*Alias)(&obj)})
}

/**
*	Generated in Payment Only [mode](/checkout/overview#complete-checkout-and-payment-only-modes) when the customer hasn't accepted the terms and conditions yet.
*
 */
type ExternalTermsAndConditionsPending struct {
	// \`info`
	Severity string `json:"severity"`
	// External terms and conditions pending.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExternalTermsAndConditionsPending) MarshalJSON() ([]byte, error) {
	type Alias ExternalTermsAndConditionsPending
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "external_terms_and_conditions_pending", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the [Checkout Session](/installing-checkout#create-checkout-sessions) fails to refresh.
*
 */
type FailedToRefreshSession struct {
	// \`error`
	Severity string `json:"severity"`
	// Failed to refresh session.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj FailedToRefreshSession) MarshalJSON() ([]byte, error) {
	type Alias FailedToRefreshSession
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "failed_to_refresh_session", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an error occurs while retrieving the balance of a gift card.
*
 */
type GiftCardBalanceError struct {
	// \`error`
	Severity string `json:"severity"`
	// Gift card balance failed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GiftCardBalanceError) MarshalJSON() ([]byte, error) {
	type Alias GiftCardBalanceError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "gift_card_balance_error", Alias: (*Alias)(&obj)})
}

func (obj GiftCardBalanceError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown GiftCardBalanceError: failed to parse error response"
}

/**
*	Generated when the customer removes a gift card's balance that was initially applied as a payment integration.
*
 */
type GiftCardBalanceRemoved struct {
	// \`info`
	Severity string `json:"severity"`
	// Gift card balance removed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GiftCardBalanceRemoved) MarshalJSON() ([]byte, error) {
	type Alias GiftCardBalanceRemoved
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "gift_card_balance_removed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the customer enters the gift card's code for the payment, and Checkout initiates the retrieving of the gift card's balance.
*
 */
type GiftCardBalanceStarted struct {
	// \`info`
	Severity string `json:"severity"`
	// Gift card balance started.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GiftCardBalanceStarted) MarshalJSON() ([]byte, error) {
	type Alias GiftCardBalanceStarted
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "gift_card_balance_started", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the redemption of a gift card's balance fails.
*
 */
type GiftCardRedeemError struct {
	// \`error`
	Severity string `json:"severity"`
	// Gift card redeem failed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GiftCardRedeemError) MarshalJSON() ([]byte, error) {
	type Alias GiftCardRedeemError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "gift_card_redeem_error", Alias: (*Alias)(&obj)})
}

func (obj GiftCardRedeemError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown GiftCardRedeemError: failed to parse error response"
}

/**
*	Generated when the customer confirms the payment by gift card and Checkout initiates the redemption of the gift card's balance.
*
 */
type GiftCardRedeemStarted struct {
	// \`info`
	Severity string `json:"severity"`
	// Gift card redeem started.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GiftCardRedeemStarted) MarshalJSON() ([]byte, error) {
	type Alias GiftCardRedeemStarted
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "gift_card_redeem_started", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the redemption of a gift card's balance is successful.
*
 */
type GiftCardRedeemSuccess struct {
	// \`info`
	Severity string `json:"severity"`
	// Gift card redeem success.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GiftCardRedeemSuccess) MarshalJSON() ([]byte, error) {
	type Alias GiftCardRedeemSuccess
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "gift_card_redeem_success", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an error occurs during Checkout's initialization.
*
 */
type InitError struct {
	// \`error`
	Severity string `json:"severity"`
	// Error during initialization.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InitError) MarshalJSON() ([]byte, error) {
	type Alias InitError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "init_error", Alias: (*Alias)(&obj)})
}

func (obj InitError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown InitError: failed to parse error response"
}

/**
*	Generated when Checkout does not receive the configuration properties with the `checkoutFlow` or `paymentFlow` [method](/sdk) on time.
*
 */
type InitTimeout struct {
	// \`error`
	Severity string `json:"severity"`
	// Timeout error, no init message received.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
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
*	Generated when the Application does not support the requested Checkout [mode](/checkout/overview#complete-checkout-and-payment-only-modes).
*
 */
type InvalidMode struct {
	// \`info`
	Severity string `json:"severity"`
	// \`{mode}` mode requires an application with mode type `{modeTypeRequested}`. Current mode `{modeTypeReceived}`.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidMode) MarshalJSON() ([]byte, error) {
	type Alias InvalidMode
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "invalid_mode", Alias: (*Alias)(&obj)})
}

/**
*	Generated when multiple containers for [vendor payment buttons](/sdk#display-custom-and-vendor-payment-buttons) are found.
*
 */
type MultipleVendorButtonContainers struct {
	// \`warn`
	Severity string `json:"severity"`
	// Multiple vendor button containers detected, this may cause issues.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MultipleVendorButtonContainers) MarshalJSON() ([]byte, error) {
	type Alias MultipleVendorButtonContainers
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "multiple_vendor_button_containers", Alias: (*Alias)(&obj)})
}

/**
*	Generated when no express payment integrations are configured.
*
 */
type NoExpressPaymentIntegrations struct {
	// \`error`
	Severity string `json:"severity"`
	// There are no Express Payments integrations configured.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NoExpressPaymentIntegrations) MarshalJSON() ([]byte, error) {
	type Alias NoExpressPaymentIntegrations
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "no_express_payment_integrations", Alias: (*Alias)(&obj)})
}

/**
*	Generated when no payment integration is set up for an [Application](/connectors-and-applications#applications). Add at least one Payment integration to the Application in the Merchant Center.
*
 */
type NoPaymentIntegrations struct {
	// \`error`
	Severity string `json:"severity"`
	// There are no payment integrations configured.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NoPaymentIntegrations) MarshalJSON() ([]byte, error) {
	type Alias NoPaymentIntegrations
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "no_payment_integrations", Alias: (*Alias)(&obj)})
}

/**
*	Generated when Checkout starts verifying the [Order](/api/projects/orders).
*
 */
type OrderVerificationStarted struct {
	// \`info`
	Severity string `json:"severity"`
	// Order verification started.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderVerificationStarted) MarshalJSON() ([]byte, error) {
	type Alias OrderVerificationStarted
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "order_verification_started", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the verification of the [Order](/api/projects/orders) times out.
*
 */
type OrderVerificationTimeout struct {
	// \`error`
	Severity string `json:"severity"`
	// Order verification timeout.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderVerificationTimeout) MarshalJSON() ([]byte, error) {
	type Alias OrderVerificationTimeout
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "order_verification_timeout", Alias: (*Alias)(&obj)})
}

/**
*	Generated when Checkout and the payment service provider (PSP) validate the payment information entered by the customer.
*
 */
type PaymentValidationPassed struct {
	// \`info`
	Severity string `json:"severity"`
	// Payment validation passed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentValidationPassed) MarshalJSON() ([]byte, error) {
	type Alias PaymentValidationPassed
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_validation_passed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the validation of the payment information entered by the customer starts.
*
 */
type PaymentValidationStarted struct {
	// \`info`
	Severity string `json:"severity"`
	// Payment validation started.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentValidationStarted) MarshalJSON() ([]byte, error) {
	type Alias PaymentValidationStarted
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_validation_started", Alias: (*Alias)(&obj)})
}

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
	case "add_discount_code_error":
		obj := AddDiscountCodeError{}
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
	case "invalid_fields":
		obj := BadInputData{}
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
	case "cart_with_existing_payment":
		obj := CartWithExistingPayment{}
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
	case "connector_error":
		obj := ConnectorError{}
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
	case "discount_code_not_applicable":
		obj := DiscountCodeNotApplicable{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "express_container_not_found":
		obj := ExpressContainerNotFound{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "express_multiple_containers_found":
		obj := ExpressMultipleContainersFound{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "express_payment_cancelled":
		obj := ExpressPaymentCancelled{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "express_payment_completed":
		obj := ExpressPaymentCompleted{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "express_payment_failed":
		obj := ExpressPaymentFailed{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "express_payment_integration_not_available":
		obj := ExpressPaymentIntegrationNotAvailable{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "express_payment_started":
		obj := ExpressPaymentStarted{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "gift_card_balance_success":
		obj := GiftCardBalanceSuccess{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "invalid_locale":
		obj := InvalidLocale{}
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
	case "non_orderable_cart_error":
		obj := NonOrderableCartError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "not_applicable_discount_code_removed":
		obj := NotApplicableDiscountCodeRemoved{}
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
	case "order_creation_error":
		obj := OrderCreationError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "order_verification_retry_error":
		obj := OrderVerificationRetryError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_cancelled":
		obj := PaymentCancelled{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_failed":
		obj := PaymentFailed{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_integration_loaded":
		obj := PaymentIntegrationLoaded{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_integration_loading":
		obj := PaymentIntegrationLoading{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_integration_loading_error":
		obj := PaymentIntegrationLoadingError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_integration_not_available":
		obj := PaymentIntegrationNotAvailable{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_integration_selected":
		obj := PaymentIntegrationSelected{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_integration_selection_confirmation":
		obj := PaymentIntegrationSelectionConfirmation{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_integration_selection_confirmation_failed":
		obj := PaymentIntegrationSelectionConfirmationFailed{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_integrations_received":
		obj := PaymentIntegrationsReceived{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_started":
		obj := PaymentStarted{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "payment_validation_failed":
		obj := PaymentValidationFailed{}
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
	case "remove_discount_code_error":
		obj := RemoveDiscountCodeError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "set_shipping_address_error":
		obj := SetShippingAddressError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "shipping_address_missing":
		obj := ShippingAddressMissingError{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "shipping_method_selected":
		obj := ShippingMethodSelected{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "shipping_method_selection_confirmation":
		obj := ShippingMethodSelectionConfirmation{}
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
	case "unsupported_country":
		obj := UnsupportedCountry{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "updated_fields":
		obj := UpdatedFields{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Generated when there was an error adding a Discount Code.
*
 */
type AddDiscountCodeError struct {
	// \`error`
	Severity string `json:"severity"`
	// Error adding discount code.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `error` object.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddDiscountCodeError) MarshalJSON() ([]byte, error) {
	type Alias AddDiscountCodeError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "add_discount_code_error", Alias: (*Alias)(&obj)})
}

func (obj AddDiscountCodeError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown AddDiscountCodeError: failed to parse error response"
}

/**
*	Generated when the requested [Application](/connectors-and-applications#applications) is deactivated. Activate the Application in the Merchant Center to continue.
*
 */
type ApplicationDeactivated struct {
	// \`error`
	Severity string `json:"severity"`
	// Application `{applicationKey}` for `{projectKey}` is disabled.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `project` and `application` objects with the related `key` property.
	Payload interface{} `json:"payload"`
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
*	Generated when the [Checkout properties](/sdk#configuration-properties) contains invalid fields.
*
 */
type BadInputData struct {
	// \`error`
	Severity string `json:"severity"`
	// Some fields are invalid.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `invalidFields` array of objects with the related `schema`, `path`, `value` and `message` properties.
	Payload interface{} `json:"payload"`
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
*	Generated when the [Cart](/api/projects/carts) was emptied during the checkout process. It is not possible to recover from this, the customer must restart the checkout process.
*
 */
type CartEmptiedDuringCheckout struct {
	// \`error`
	Severity string `json:"severity"`
	// Cart `{cartId}` was emptied during checkout.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `cart` object with the `id` property.
	Payload interface{} `json:"payload"`
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
*	Generated when the [Cart](/api/projects/carts) for the current checkout is empty. The Cart must contain at least one [Line Item](/api/carts-orders-overview#line-items).
*
 */
type CartEmpty struct {
	// \`error`
	Severity string `json:"severity"`
	// Cart `{cartId}` is empty.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `cart` object with the `id` property.
	Payload interface{} `json:"payload"`
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
*	Generated when the [Cart](/api/projects/carts) is not found. To start the checkout process, a valid Cart with at least one [Line Item](/api/carts-orders-overview#line-items) is required.
*
 */
type CartNotFound struct {
	// \`error`
	Severity string `json:"severity"`
	// Cart for `{projectKey}` and session `{sessionId}` not found.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `project` and `session` objects with the related `key` and `id` properties.
	Payload interface{} `json:"payload"`
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
*	Generated when trying to add a [Payment](/api/projects/payments) to a [Cart](/api/projects/carts) that already references an approved Payment.
*
 */
type CartWithExistingPayment struct {
	// \`error`
	Severity string `json:"severity"`
	// Cart with existing approved payment.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `cart` object with the `id` property.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartWithExistingPayment) MarshalJSON() ([]byte, error) {
	type Alias CartWithExistingPayment
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "cart_with_existing_payment", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the customer completes the checkout process.
*
 */
type CheckoutCompleted struct {
	// \`info`
	Severity string `json:"severity"`
	// Checkout for `{orderId}` completed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `order` object with the `id` property.
	Payload interface{} `json:"payload"`
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
*	Generated when the [Connector](/checkout/connectors-and-applications#payment-connectors) triggers an error.
*
 */
type ConnectorError struct {
	// \`error`
	Severity string `json:"severity"`
	// Connector error.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `connector` object with the `id` property and optional `error`, `message`, and `data` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConnectorError) MarshalJSON() ([]byte, error) {
	type Alias ConnectorError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "connector_error", Alias: (*Alias)(&obj)})
}

func (obj ConnectorError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ConnectorError: failed to parse error response"
}

/**
*	Generated when the [Checkout properties](/sdk#configuration-properties) contains one or more deprecated fields.
*
 */
type DeprecatedFields struct {
	// \`warn`
	Severity string `json:"severity"`
	// Some fields are deprecated.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `deprecatedFields` array.
	Payload interface{} `json:"payload"`
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
*	Generated when the [Discount Code](/api/projects/discountCodes) is not applicable for the current Cart.
*
 */
type DiscountCodeNotApplicable struct {
	// \`info`
	Severity string `json:"severity"`
	// Discount code not applicable.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `cartId` and `discountCode` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeNotApplicable) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeNotApplicable
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "discount_code_not_applicable", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the container element specified in the `expressId` option is not found in the DOM.
*
 */
type ExpressContainerNotFound struct {
	// \`error`
	Severity string `json:"severity"`
	// Express container not found.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `id` of the container that was not found.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExpressContainerNotFound) MarshalJSON() ([]byte, error) {
	type Alias ExpressContainerNotFound
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "express_container_not_found", Alias: (*Alias)(&obj)})
}

/**
*	Generated when multiple elements with the `data-ctc-express` attribute are found without an `expressId` specified. Use the `expressId` option to specify which container to use.
*
 */
type ExpressMultipleContainersFound struct {
	// \`error`
	Severity string `json:"severity"`
	// Multiple elements with data-ctc-express attribute found. Use expressId to specify which container to use.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `count` property indicating the number of containers found.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExpressMultipleContainersFound) MarshalJSON() ([]byte, error) {
	type Alias ExpressMultipleContainersFound
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "express_multiple_containers_found", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the customer cancels the Express Payments process.
*
 */
type ExpressPaymentCancelled struct {
	// \`info`
	Severity string `json:"severity"`
	// Express Payments process cancelled.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `paymentReference` and `method` object with the `type` property (if available).
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExpressPaymentCancelled) MarshalJSON() ([]byte, error) {
	type Alias ExpressPaymentCancelled
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "express_payment_cancelled", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the Express Payments process is completed.
*
 */
type ExpressPaymentCompleted struct {
	// \`info`
	Severity string `json:"severity"`
	// Express payment for `{orderId}` completed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `order` object with the `id` property.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExpressPaymentCompleted) MarshalJSON() ([]byte, error) {
	type Alias ExpressPaymentCompleted
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "express_payment_completed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the Express Payments process fails.
*
 */
type ExpressPaymentFailed struct {
	// \`error`
	Severity string `json:"severity"`
	// Express payment failed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `paymentReference`, `error`, and `method` object with the `type` property (if available).
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExpressPaymentFailed) MarshalJSON() ([]byte, error) {
	type Alias ExpressPaymentFailed
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "express_payment_failed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when a requested Express Payments integration is not available. This can happen when using `mountMethod()` with a payment method that is not configured or not available for the current session.
*
 */
type ExpressPaymentIntegrationNotAvailable struct {
	// \`warn`
	Severity string `json:"severity"`
	// Express Payments integration not available.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `integration` object with the `type` property indicating the requested payment integration type, and `availableMethods` array listing the available payment methods.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExpressPaymentIntegrationNotAvailable) MarshalJSON() ([]byte, error) {
	type Alias ExpressPaymentIntegrationNotAvailable
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "express_payment_integration_not_available", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the customer clicks an Express Payments button and the Express Payment process starts.
*
 */
type ExpressPaymentStarted struct {
	// \`info`
	Severity string `json:"severity"`
	// Express payment started.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `integration` object with the `type` property.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExpressPaymentStarted) MarshalJSON() ([]byte, error) {
	type Alias ExpressPaymentStarted
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "express_payment_started", Alias: (*Alias)(&obj)})
}

/**
*	Generated when Checkout successfully retrieves a gift card's balance.
*
 */
type GiftCardBalanceSuccess struct {
	// \`info`
	Severity string `json:"severity"`
	// Gift card balance started.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `amount` and  `isBalanceSufficient` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj GiftCardBalanceSuccess) MarshalJSON() ([]byte, error) {
	type Alias GiftCardBalanceSuccess
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "gift_card_balance_success", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the provided locale is invalid.
*
 */
type InvalidLocale struct {
	// \`warn`
	Severity string `json:"severity"`
	// The provided `{locale}` is invalid.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the provided locale string, the related locale field (language or currency), and the used fallback one.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InvalidLocale) MarshalJSON() ([]byte, error) {
	type Alias InvalidLocale
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "invalid_locale", Alias: (*Alias)(&obj)})
}

/**
*	Generated when no [Shipping Method](/api/projects/shippingMethods) is available for the shipping address of the [Cart](/api/projects/carts). This may indicate an incomplete configuration.
*
 */
type NoShippingMethods struct {
	// \`error`
	Severity string `json:"severity"`
	// There are no shipping methods matching cart.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `cart` object with the `id` property.
	Payload interface{} `json:"payload"`
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
*	Generated when the Order could not be created due to inconsistencies in the [Cart](/api/projects/carts).
*
 */
type NonOrderableCartError struct {
	// \`error`
	Severity string `json:"severity"`
	// This cart is not orderable.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `cart` object with the `id` and `errors` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NonOrderableCartError) MarshalJSON() ([]byte, error) {
	type Alias NonOrderableCartError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "non_orderable_cart_error", Alias: (*Alias)(&obj)})
}

func (obj NonOrderableCartError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown NonOrderableCartError: failed to parse error response"
}

/**
*	Generated when using the `paymentFlow` [method](/checkout/browser-sdk#paymentflow-method) and Checkout removes a [Discount Code](/api/projects/discountCodes) from the Cart because it does not apply to the Cart. Checkout removes the Discount Code to avoid an order creation error when converting the Cart to an Order.
*
 */
type NotApplicableDiscountCodeRemoved struct {
	// \`warn`
	Severity string `json:"severity"`
	// Not applicable discount code removed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `cartId` and `discountCode` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NotApplicableDiscountCodeRemoved) MarshalJSON() ([]byte, error) {
	type Alias NotApplicableDiscountCodeRemoved
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "not_applicable_discount_code_removed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an [Order](/api/projects/orders) is created after a successful checkout process.
*
 */
type OrderCreated struct {
	// \`info`
	Severity string `json:"severity"`
	// Order `{orderId}` created.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `order` object with the `id` property.
	Payload interface{} `json:"payload"`
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
*	Generated when an [Order](/api/projects/orders) that references an approved [Payment](/api/projects/payments) cannot be created.
*
 */
type OrderCreationError struct {
	// \`error`
	Severity string `json:"severity"`
	// Order creation failed with approved payment.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains `paymentReference`, `sessionId`, and the `errors` array of objects, with the related `code` and `message` properties.
	Payload interface{} `json:"payload"`
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
*	Generated when retrying the verification of the [Order](/api/projects/orders) results in an error.
*
 */
type OrderVerificationRetryError struct {
	// \`error`
	Severity string `json:"severity"`
	// Order verification retry error.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `error` property that can be either `orderReferenceNotAvailable` or `orderVerificationOngoing`.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderVerificationRetryError) MarshalJSON() ([]byte, error) {
	type Alias OrderVerificationRetryError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "order_verification_retry_error", Alias: (*Alias)(&obj)})
}

func (obj OrderVerificationRetryError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown OrderVerificationRetryError: failed to parse error response"
}

/**
*	Generated when the customer cancels the payment (for example, by closing the browser's window).
*
 */
type PaymentCancelled struct {
	// \`info`
	Severity string `json:"severity"`
	// Payment cancelled.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the reference data of a Payment.
	Payload PaymentReference `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentCancelled) MarshalJSON() ([]byte, error) {
	type Alias PaymentCancelled
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_cancelled", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the payment fails.
*
 */
type PaymentFailed struct {
	// \`error`
	Severity string `json:"severity"`
	// Payment failed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the reference data of a Payment.
	Payload PaymentReference `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentFailed) MarshalJSON() ([]byte, error) {
	type Alias PaymentFailed
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_failed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the selected payment integration is loaded.
*
 */
type PaymentIntegrationLoaded struct {
	// \`info`
	Severity string `json:"severity"`
	// Payment integration loaded.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `integration` object with the `type` and `connectorId` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntegrationLoaded) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntegrationLoaded
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_integration_loaded", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the selected payment integration is loading.
*
 */
type PaymentIntegrationLoading struct {
	// \`info`
	Severity string `json:"severity"`
	// Payment integration loading.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `integration` object with the `type` and `connectorId` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntegrationLoading) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntegrationLoading
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_integration_loading", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the loading of the selected payment integration fails.
*
 */
type PaymentIntegrationLoadingError struct {
	// \`error`
	Severity string `json:"severity"`
	// Payment integration loading failed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains an `error` object and the `integration` object with the `type` and `connectorId` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntegrationLoadingError) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntegrationLoadingError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_integration_loading_error", Alias: (*Alias)(&obj)})
}

func (obj PaymentIntegrationLoadingError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown PaymentIntegrationLoadingError: failed to parse error response"
}

/**
*	Generated when there is an error with the selected payment integration and the payment integration is unavailable.
*
 */
type PaymentIntegrationNotAvailable struct {
	// \`warn`
	Severity string `json:"severity"`
	// Payment integration not available.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `method` object with the `type`, `id`, and `connectorId` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntegrationNotAvailable) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntegrationNotAvailable
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_integration_not_available", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the customer selects the payment integration.
*
 */
type PaymentIntegrationSelected struct {
	// \`info`
	Severity string `json:"severity"`
	// Payment integration selected.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `integration` object with the `type` and `hasVendorButton` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntegrationSelected) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntegrationSelected
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_integration_selected", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the customer has entered the payment integration information and moves to the next step.
*
 */
type PaymentIntegrationSelectionConfirmation struct {
	// \`info`
	Severity string `json:"severity"`
	// Payment integration selected.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `integration` object with the `type` and `hasVendorButton` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntegrationSelectionConfirmation) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntegrationSelectionConfirmation
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_integration_selection_confirmation", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the selection of the payment integration by the customer is not successful.
*
 */
type PaymentIntegrationSelectionConfirmationFailed struct {
	// \`info`
	Severity string `json:"severity"`
	// Payment integration selection failed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `integration` object with the `type` property.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntegrationSelectionConfirmationFailed) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntegrationSelectionConfirmationFailed
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_integration_selection_confirmation_failed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when Checkout receives and loads the payment integrations configured for the Application.
*
 */
type PaymentIntegrationsReceived struct {
	// \`info`
	Severity string `json:"severity"`
	// Payment integrations received.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `paymentIntegrations` array of objects with the `type` and `connectorId` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntegrationsReceived) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntegrationsReceived
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_integrations_received", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the payment starts.
*
 */
type PaymentStarted struct {
	// \`info`
	Severity string `json:"severity"`
	// Payment started.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `integration` object with the `type` property.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentStarted) MarshalJSON() ([]byte, error) {
	type Alias PaymentStarted
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_started", Alias: (*Alias)(&obj)})
}

/**
*	Generated when Checkout and the payment service provider (PSP) do not validate the payment information entered by the customer.
*
 */
type PaymentValidationFailed struct {
	// \`info`
	Severity string `json:"severity"`
	// Payment validation failed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the reference data of a Payment.
	Payload PaymentReference `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentValidationFailed) MarshalJSON() ([]byte, error) {
	type Alias PaymentValidationFailed
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "payment_validation_failed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the Checkout [`projectKey`](/sdk) is deactivated and cannot be initialized. To activate it, contact the [commercetools support team](https://support.commercetools.com/).
*
 */
type ProjectIsDeactivated struct {
	// \`error`
	Severity string `json:"severity"`
	// Project `{projectKey}` is deactivated.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `project` object with the `key` property.
	Payload interface{} `json:"payload"`
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
*	Generated when there was an error removing the [Discount Code](/api/projects/discountCodes).
*
 */
type RemoveDiscountCodeError struct {
	// \`error`
	Severity string `json:"severity"`
	// Error removing discount code.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `error` object.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveDiscountCodeError) MarshalJSON() ([]byte, error) {
	type Alias RemoveDiscountCodeError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "remove_discount_code_error", Alias: (*Alias)(&obj)})
}

func (obj RemoveDiscountCodeError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown RemoveDiscountCodeError: failed to parse error response"
}

/**
*	Generated when the [Cart](/api/projects/carts) cannot be updated with the shipping address.
*
 */
type SetShippingAddressError struct {
	// \`error`
	Severity string `json:"severity"`
	// Error setting shipping address.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `errors` array of objects, with the related `code` and `message` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetShippingAddressError) MarshalJSON() ([]byte, error) {
	type Alias SetShippingAddressError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "set_shipping_address_error", Alias: (*Alias)(&obj)})
}

func (obj SetShippingAddressError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown SetShippingAddressError: failed to parse error response"
}

/**
*	Generated when the shipping address is missing for the given Cart.
*
 */
type ShippingAddressMissingError struct {
	// \`error`
	Severity string `json:"severity"`
	// The shippingAddress field is missing for cart `{cartId}`.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `cart` object with the `id` property.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingAddressMissingError) MarshalJSON() ([]byte, error) {
	type Alias ShippingAddressMissingError
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "shipping_address_missing", Alias: (*Alias)(&obj)})
}

func (obj ShippingAddressMissingError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ShippingAddressMissingError: failed to parse error response"
}

/**
*	Generated when the selected Shipping Method does not match the Cart anymore.
*
 */
type ShippingMethodDoesNotMatchCart struct {
	// \`info`
	Severity string `json:"severity"`
	// Selected shipping method no longer matches cart.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodDoesNotMatchCart) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodDoesNotMatchCart
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "shipping_method_does_not_match_cart", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the customer selects a Shipping Method that is different from the default option.
*
 */
type ShippingMethodSelected struct {
	// \`info`
	Severity string `json:"severity"`
	// Shipping Method selected.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `method` object with the `name` and `id` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodSelected) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSelected
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "shipping_method_selected", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the customer selects a Shipping Method and moves to the next step of the checkout process.
*
 */
type ShippingMethodSelectionConfirmation struct {
	// \`info`
	Severity string `json:"severity"`
	// Shipping Method selection confirmed.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `method` object with the `name` and `id` properties.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodSelectionConfirmation) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSelectionConfirmation
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "shipping_method_selection_confirmation", Alias: (*Alias)(&obj)})
}

/**
*	Generated when the provided `locale` is not [available for localization](/installing-checkout#locales). The localization falls back to English.
*
 */
type UnavailableLocale struct {
	// \`warn`
	Severity string `json:"severity"`
	// The provided locale `{locale}` is not available for translated definitions.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `locale` and the `fallback` properties.
	Payload interface{} `json:"payload"`
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
*	Generated when the country of the shipping address and/or billing address associated with the [Cart](/api/projects/carts) does not match the countries set for the [Application](/connectors-and-applications#applications).
*
 */
type UnsupportedCountry struct {
	// \`error`
	Severity string `json:"severity"`
	// Cart `{cartId}` has unsupported country.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains the `cart` object with the `id` property, `shippingCountry`, `billingCountry`, and the `supportedCountries` array.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj UnsupportedCountry) MarshalJSON() ([]byte, error) {
	type Alias UnsupportedCountry
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "unsupported_country", Alias: (*Alias)(&obj)})
}

/**
*	Generated when some fields provided in the Checkout initialization request are invalid and have been updated to match our schema. An array of updates will be provided with a sequence of edits used to transform an invalid value into a valid one.
*
 */
type UpdatedFields struct {
	// \`warn`
	Severity string `json:"severity"`
	// Some fields are invalid and have been updated.
	Message string `json:"message"`
	// Unique identifier of the event.
	CorrelationId string `json:"correlationId"`
	// Contains two arrays:
	// - `invalidFields` contains a list of invalid fields with the related `schema`, `path`, `value` and `message` properties.
	// - `updatedFields` contains the updated fields with its related `type`, `path` and `value` properties, where `type` can be `update`|`insert`|`delete`.
	Payload interface{} `json:"payload"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj UpdatedFields) MarshalJSON() ([]byte, error) {
	type Alias UpdatedFields
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "updated_fields", Alias: (*Alias)(&obj)})
}
