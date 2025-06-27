package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

/**
*	Represents the data used to import a DiscountCode. Once imported, this data is persisted as a [DiscountCode](ctp:api:type:DiscountCode) in the Project.
*
 */
type DiscountCodeImport struct {
	// User-defined unique identifier. If a [DiscountCode](ctp:api:type:DiscountCode) with this `key` exists, it is updated with the imported data.
	Key string `json:"key"`
	// Maps to `DiscountCode.name`.
	Name *LocalizedString `json:"name,omitempty"`
	// Maps to `DiscountCode.description`.
	Description *LocalizedString `json:"description,omitempty"`
	// Maps to `DiscountCode.code`. This value cannot be updated. Attempting to update this value will result in an [InvalidFieldsUpdate](/import-export/error#invalidfieldsupdateerror) error.
	Code string `json:"code"`
	// Maps to `DiscountCode.cartDiscounts`. If the referenced [CartDiscounts](ctp:api:type:CartDiscount) do not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced CartDiscounts are created.
	CartDiscounts []CartDiscountKeyReference `json:"cartDiscounts"`
	// Maps to `DiscountCode.cartPredicate`.
	CartPredicate *string `json:"cartPredicate,omitempty"`
	// Maps to `DiscountCode.isActive`.
	IsActive bool `json:"isActive"`
	// Maps to `DiscountCode.maxApplications`.
	MaxApplications *int `json:"maxApplications,omitempty"`
	// Maps to `DiscountCode.maxApplicationsPerCustomer`.
	MaxApplicationsPerCustomer *int `json:"maxApplicationsPerCustomer,omitempty"`
	// Maps to `DiscountCode.groups`.
	Groups []string `json:"groups"`
	// Maps to `DiscountCode.validFrom`.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Maps to `DiscountCode.validUntil`.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Maps to `DiscountCode.custom`.
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
