package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type AllowedOrigins struct {
	// Indicates whether all the origin URLs are allowed or not.
	AllowAll bool `json:"allowAll"`
	// Origins to allow. Set to an empty array to remove any existing origins.
	Origins []string `json:"origins"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AllowedOrigins) MarshalJSON() ([]byte, error) {
	type Alias AllowedOrigins
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

	if raw["origins"] == nil {
		delete(raw, "origins")
	}

	return json.Marshal(raw)

}

type Application struct {
	// Unique identifier of the Application.
	ID string `json:"id"`
	// Current version of the Application.
	Version int `json:"version"`
	// User-defined unique identifier of the Application. `MinLength: 2 ​MaxLength: 256 ​Pattern: ^[A-Za-z0-9_-]+$`
	Key string `json:"key"`
	// Indicates whether the Application is active or not.
	Status *ApplicationStatus `json:"status,omitempty"`
	// Name of the Application.
	Name string `json:"name"`
	// Description of the Application.
	Description *LocalizedString `json:"description,omitempty"`
	// Details of the application logo. It is required if the `mode = CompleteFlow`.
	Logo *ApplicationLogo `json:"logo,omitempty"`
	// List of countries that the Application supports. Each country is represented by two-digit country code as per ISO 3166-1 alpha-2.
	Countries []string `json:"countries"`
	// Allowed origin URLs for the Application.
	AllowedOrigins AllowedOrigins `json:"allowedOrigins"`
	// Checkout mode to use for the Application.
	Mode ApplicationMode `json:"mode"`
	// Sets the payment configuration options of the Application.
	PaymentsConfiguration PaymentsConfiguration `json:"paymentsConfiguration"`
	// User agreements to display. It is only applicable if the `mode` is `CompleteFlow`.
	Agreements []ApplicationAgreement `json:"agreements"`
	// Sets the discounts configuration options of the Application. It is only applicable if the `mode` is `CompleteFlow`.
	DiscountsConfiguration *DiscountsConfiguration `json:"discountsConfiguration,omitempty"`
	// Date and time (UTC) the Application was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// IDs and references that created the Application.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Date and time (UTC) the Application was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the Application.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Application) MarshalJSON() ([]byte, error) {
	type Alias Application
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

	if raw["agreements"] == nil {
		delete(raw, "agreements")
	}

	return json.Marshal(raw)

}

type ApplicationAgreement struct {
	// Unique identifier of the ApplicationAgreement.
	ID string `json:"id"`
	// Unique name identifier for the ApplicationAgreement.
	Name string `json:"name"`
	// Type of agreement.
	Type ApplicationAgreementType `json:"type"`
	// Indicates whether the agreement is active or not.
	Status ApplicationAgreementStatus `json:"status"`
	// Text to display in checkout in Markdown format.
	Text LocalizedString `json:"text"`
}

type ApplicationAgreementDraft struct {
	// Unique name for identifying the agreement.
	Name string `json:"name"`
	// Type of agreement.
	Type ApplicationAgreementType `json:"type"`
	// Indicates whether the agreement is active or not.
	Status ApplicationAgreementStatus `json:"status"`
	// Text to display in checkout in Markdown format.
	Text LocalizedString `json:"text"`
}

type ApplicationAgreementStatus string

const (
	ApplicationAgreementStatusActive   ApplicationAgreementStatus = "Active"
	ApplicationAgreementStatusInactive ApplicationAgreementStatus = "Inactive"
)

type ApplicationAgreementType string

const (
	ApplicationAgreementTypeText              ApplicationAgreementType = "Text"
	ApplicationAgreementTypeMandatoryCheckbox ApplicationAgreementType = "MandatoryCheckbox"
)

type ApplicationDraft struct {
	// User-defined unique identifier of the Application. `MinLength: 2 ​MaxLength: 256 ​Pattern: ^[A-Za-z0-9_-]+$`
	Key string `json:"key"`
	// Indicates whether the Application is active or not.
	Status *ApplicationStatus `json:"status,omitempty"`
	// Name of the Application.
	Name string `json:"name"`
	// Description of the Application.
	Description *LocalizedString `json:"description,omitempty"`
	// Details of the Application logo. It is required if the `mode` is `CompleteFlow`.
	Logo *ApplicationLogo `json:"logo,omitempty"`
	// List of countries that the Application supports. Each country is represented by two-digit country code as per ISO 3166-1 alpha-2.
	Countries []string `json:"countries"`
	// Allowed origin URLs for the Application.
	AllowedOrigins AllowedOrigins `json:"allowedOrigins"`
	// Checkout mode to use for the Application.
	Mode ApplicationMode `json:"mode"`
	// Sets the payment configuration options of the Application.
	PaymentsConfiguration PaymentsConfiguration `json:"paymentsConfiguration"`
	// User agreements to display. It is only applicable if the `mode` is `CompleteFlow`.
	Agreements []ApplicationAgreement `json:"agreements"`
	// Sets the discounts configuration options of the Application. It is only applicable if the `mode` is `CompleteFlow`.
	DiscountsConfiguration *DiscountsConfiguration `json:"discountsConfiguration,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApplicationDraft) MarshalJSON() ([]byte, error) {
	type Alias ApplicationDraft
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

	if raw["agreements"] == nil {
		delete(raw, "agreements")
	}

	return json.Marshal(raw)

}

/**
*	Details of the Application logo. It is required if the `mode` is `CompleteFlow`.
*
 */
type ApplicationLogo struct {
	// URL of the logo.
	Url *string `json:"url,omitempty"`
}

type ApplicationMode string

const (
	ApplicationModeCompleteFlow ApplicationMode = "CompleteFlow"
	ApplicationModePaymentOnly  ApplicationMode = "PaymentOnly"
)

/**
*	Reference to an [Application](ctp:checkout:type:Application).
*
 */
type ApplicationReference struct {
	// Unique identifier of the referenced Application.
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApplicationReference) MarshalJSON() ([]byte, error) {
	type Alias ApplicationReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "application", Alias: (*Alias)(&obj)})
}

/**
*	Resource identifier to an [Application](ctp:checkout:type:Application). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](ctp:api:type:InvalidJsonInputError) error is returned.
*
 */
type ApplicationResourceIdentifier struct {
	// Unique identifier of the referenced Application. Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced Application. Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApplicationResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ApplicationResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "application", Alias: (*Alias)(&obj)})
}

type ApplicationStatus string

const (
	ApplicationStatusActive   ApplicationStatus = "Active"
	ApplicationStatusInactive ApplicationStatus = "Inactive"
)

/**
*	Base type for all Application update actions.
*
 */
type ApplicationUpdateAction interface{}

func mapDiscriminatorApplicationUpdateAction(input interface{}) (ApplicationUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "addAllowedOrigin":
		obj := AddAllowedOriginUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addAgreement":
		obj := AddApplicationAgreementUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addCountry":
		obj := AddCountryUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAllowedOrigin":
		obj := RemoveAllowedOriginUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAgreement":
		obj := RemoveApplicationAgreementUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeCountry":
		obj := RemoveCountryUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "reorderAgreement":
		obj := ReorderApplicationAgreementUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setActivePaymentComponentType":
		obj := SetActivePaymentComponentTypeUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAllowAllOrigins":
		obj := SetAllowAllOriginsUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAllowedOrigins":
		obj := SetAllowedOriginsUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAgreementName":
		obj := SetApplicationAgreementNameUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAgreementStatus":
		obj := SetApplicationAgreementStatusUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAgreementText":
		obj := SetApplicationAgreementTextUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAgreementType":
		obj := SetApplicationAgreementTypeUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAgreements":
		obj := SetApplicationAgreementsUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setApplicationLogo":
		obj := SetApplicationLogoUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := SetApplicationNameUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStatus":
		obj := SetApplicationStatusUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCountries":
		obj := SetCountriesUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := SetDescriptionUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDiscountsConfiguration":
		obj := SetDiscountsConfigurationUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPaymentReturnUrl":
		obj := SetPaymentReturnUrlUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPaymentsConfiguration":
		obj := SetPaymentsConfigurationUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ApplicationUpdateActions struct {
	// Expected version of the Application on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Application.
	Actions []ApplicationUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ApplicationUpdateActions) UnmarshalJSON(data []byte) error {
	type Alias ApplicationUpdateActions
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorApplicationUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type DiscountsConfiguration struct {
	// Indicates if the buyer is able to apply discounts to the cart. It is only applicable if the `mode = CompleteFlow`.
	AllowDiscounts bool `json:"allowDiscounts"`
}

/**
*	Paginated result containing Application.
*
 */
type PaginatedApplication struct {
	// Number of results requested.
	Limit int `json:"limit"`
	// Number of elements skipped.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	Total int `json:"total"`
	// Applications matching the query.
	Results []Application `json:"results"`
}

type PaymentsConfiguration struct {
	// URL where the buyer will be redirected back to after an offsite payment.
	PaymentReturnUrl *string `json:"paymentReturnUrl,omitempty"`
	// Indicates the type of payment components that the Application will use.
	ActivePaymentComponentType PaymentComponentType `json:"activePaymentComponentType"`
}

type AddAllowedOriginUpdateAction struct {
	// Origin URL to add to the allowed origins list.
	Origin string `json:"origin"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddAllowedOriginUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias AddAllowedOriginUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAllowedOrigin", Alias: (*Alias)(&obj)})
}

type AddApplicationAgreementUpdateAction struct {
	// Agreement to add to the Application.
	Agreement ApplicationAgreementDraft `json:"agreement"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddApplicationAgreementUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias AddApplicationAgreementUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAgreement", Alias: (*Alias)(&obj)})
}

type AddCountryUpdateAction struct {
	// Country to add. Represented by two-digit country code as per ISO 3166-1 alpha-2.
	Country string `json:"country"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddCountryUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias AddCountryUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addCountry", Alias: (*Alias)(&obj)})
}

type RemoveAllowedOriginUpdateAction struct {
	// Origin URL to remove from the allowed origins list.
	Origin string `json:"origin"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveAllowedOriginUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias RemoveAllowedOriginUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAllowedOrigin", Alias: (*Alias)(&obj)})
}

type RemoveApplicationAgreementUpdateAction struct {
	// ID of the agreement to remove from the Application.
	AgreementId string `json:"agreementId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveApplicationAgreementUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias RemoveApplicationAgreementUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAgreement", Alias: (*Alias)(&obj)})
}

type RemoveCountryUpdateAction struct {
	// Country to remove. Represented by two-digit country code as per ISO 3166-1 alpha-2.
	Country string `json:"country"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveCountryUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias RemoveCountryUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeCountry", Alias: (*Alias)(&obj)})
}

type ReorderApplicationAgreementUpdateAction struct {
	// Array of agreement IDs.
	AgreementIds []string `json:"agreementIds"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReorderApplicationAgreementUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias ReorderApplicationAgreementUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "reorderAgreement", Alias: (*Alias)(&obj)})
}

type SetActivePaymentComponentTypeUpdateAction struct {
	// Type of payment components that the Application will use.
	ActivePaymentComponentType PaymentComponentType `json:"activePaymentComponentType"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetActivePaymentComponentTypeUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetActivePaymentComponentTypeUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setActivePaymentComponentType", Alias: (*Alias)(&obj)})
}

type SetAllowAllOriginsUpdateAction struct {
	// Whether to allow all origins or not.
	AllowAll bool `json:"allowAll"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAllowAllOriginsUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetAllowAllOriginsUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAllowAllOrigins", Alias: (*Alias)(&obj)})
}

type SetAllowedOriginsUpdateAction struct {
	// New allowed origins configuration.
	AllowedOrigins AllowedOrigins `json:"allowedOrigins"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAllowedOriginsUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetAllowedOriginsUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAllowedOrigins", Alias: (*Alias)(&obj)})
}

type SetApplicationAgreementNameUpdateAction struct {
	// ID of the agreement to update.
	AgreementId string `json:"agreementId"`
	// New name for the agreement.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetApplicationAgreementNameUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetApplicationAgreementNameUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAgreementName", Alias: (*Alias)(&obj)})
}

type SetApplicationAgreementStatusUpdateAction struct {
	// ID of the agreement to update.
	AgreementId string `json:"agreementId"`
	// New status for the agreement.
	Status ApplicationAgreementStatus `json:"status"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetApplicationAgreementStatusUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetApplicationAgreementStatusUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAgreementStatus", Alias: (*Alias)(&obj)})
}

type SetApplicationAgreementTextUpdateAction struct {
	// ID of the agreement to update.
	AgreementId string `json:"agreementId"`
	// New text for the agreement.
	Text LocalizedString `json:"text"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetApplicationAgreementTextUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetApplicationAgreementTextUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAgreementText", Alias: (*Alias)(&obj)})
}

type SetApplicationAgreementTypeUpdateAction struct {
	// ID of the agreement to update.
	AgreementId string `json:"agreementId"`
	// New type for the agreement.
	Type ApplicationAgreementType `json:"type"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetApplicationAgreementTypeUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetApplicationAgreementTypeUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAgreementType", Alias: (*Alias)(&obj)})
}

type SetApplicationAgreementsUpdateAction struct {
	// If omitted, any existing value is removed.
	Agreements []ApplicationAgreementDraft `json:"agreements"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetApplicationAgreementsUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetApplicationAgreementsUpdateAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAgreements", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["agreements"] == nil {
		delete(raw, "agreements")
	}

	return json.Marshal(raw)

}

type SetApplicationLogoUpdateAction struct {
	// New logo to assign to the Application, which must be a valid url.
	Logo ApplicationLogo `json:"logo"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetApplicationLogoUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetApplicationLogoUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setApplicationLogo", Alias: (*Alias)(&obj)})
}

type SetApplicationNameUpdateAction struct {
	// New name to assign to the Application.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetApplicationNameUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetApplicationNameUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

type SetApplicationStatusUpdateAction struct {
	// New status to assign to the Application.
	Status ApplicationStatus `json:"status"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetApplicationStatusUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetApplicationStatusUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatus", Alias: (*Alias)(&obj)})
}

type SetCountriesUpdateAction struct {
	// Sets a list of countries to the Application. Each country is represented by two-digit country code as per ISO 3166-1 alpha-2.
	Countries []string `json:"countries"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCountriesUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetCountriesUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCountries", Alias: (*Alias)(&obj)})
}

type SetDescriptionUpdateAction struct {
	// Value to set. Set to an empty object to remove any existing value.
	Description LocalizedString `json:"description"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDescriptionUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetDescriptionUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type SetDiscountsConfigurationUpdateAction struct {
	// New discounts configuration for the Application.
	DiscountsConfiguration DiscountsConfiguration `json:"discountsConfiguration"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDiscountsConfigurationUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetDiscountsConfigurationUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDiscountsConfiguration", Alias: (*Alias)(&obj)})
}

type SetPaymentReturnUrlUpdateAction struct {
	// If omitted, any existing value is removed.
	PaymentReturnUrl *string `json:"paymentReturnUrl,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetPaymentReturnUrlUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetPaymentReturnUrlUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPaymentReturnUrl", Alias: (*Alias)(&obj)})
}

type SetPaymentsConfigurationUpdateAction struct {
	// New payments configuration for the Application.
	PaymentsConfiguration PaymentsConfiguration `json:"paymentsConfiguration"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetPaymentsConfigurationUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetPaymentsConfigurationUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPaymentsConfiguration", Alias: (*Alias)(&obj)})
}
