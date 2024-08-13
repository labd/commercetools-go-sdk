package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

/**
*	The data representation for a Discount Code to be imported that is persisted as a [Discount Code](/../api/projects/discountCodes#discountcode) in the Project.
*
 */
type DiscountCodeImport struct {
	// User-defined unique identifier. If a [Discount Code](/../api/projects/discountCodes#discountcode) with this `key` exists, it will be updated with the imported data.
	Key string `json:"key"`
	// Maps to `DiscountCode.name`.
	Name *LocalizedString `json:"name,omitempty"`
	// Maps to `DiscountCode.description`.
	Description *LocalizedString `json:"description,omitempty"`
	// User-defined unique identifier of the DiscountCode that is used by the customer to apply the discount.
	Code string `json:"code"`
	// Reference to CartDiscounts that can be applied to the Cart once the DiscountCode is applied.
	CartDiscounts []CartDiscountKeyReference `json:"cartDiscounts"`
	// DiscountCode can only be applied to Carts that match this predicate.
	CartPredicate *string `json:"cartPredicate,omitempty"`
	// Indicates if the DiscountCode is active and can be applied to the Cart.
	IsActive bool `json:"isActive"`
	// Number of times the DiscountCode can be applied. DiscountCode application is counted at the time of Order creation or update. However, Order cancellation or deletion does not decrement the count.
	MaxApplications *int `json:"maxApplications,omitempty"`
	// Number of times the DiscountCode can be applied per Customer (anonymous Carts are not supported). DiscountCode application is counted at the time of Order creation or update. However, Order cancellation or deletion does not decrement the count.
	MaxApplicationsPerCustomer *int `json:"maxApplicationsPerCustomer,omitempty"`
	// Groups to which the DiscountCode belongs.
	Groups []string `json:"groups"`
	// Date and time (UTC) from which the DiscountCode is effective.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time (UTC) until which the DiscountCode is effective.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Custom Fields of the DiscountCode.
	Custom *Custom `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeImport) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeImport
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["groups"] == nil {
		delete(raw, "groups")
	}

	return json.Marshal(raw)

}
