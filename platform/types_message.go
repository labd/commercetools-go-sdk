// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Message interface{}

func mapDiscriminatorMessage(input interface{}) (Message, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "CategoryCreated":
		new := CategoryCreatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CategorySlugChanged":
		new := CategorySlugChangedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomLineItemStateTransition":
		new := CustomLineItemStateTransitionMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressAdded":
		new := CustomerAddressAddedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressChanged":
		new := CustomerAddressChangedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressRemoved":
		new := CustomerAddressRemovedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerCompanyNameSet":
		new := CustomerCompanyNameSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerCreated":
		new := CustomerCreatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerDateOfBirthSet":
		new := CustomerDateOfBirthSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerEmailChanged":
		new := CustomerEmailChangedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerEmailVerified":
		new := CustomerEmailVerifiedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerGroupSet":
		new := CustomerGroupSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerPasswordUpdated":
		new := CustomerPasswordUpdatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryAdded":
		new := DeliveryAddedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryAddressSet":
		new := DeliveryAddressSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryItemsUpdated":
		new := DeliveryItemsUpdatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryRemoved":
		new := DeliveryRemovedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InventoryEntryCreated":
		new := InventoryEntryCreatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InventoryEntryDeleted":
		new := InventoryEntryDeletedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InventoryEntryQuantitySet":
		new := InventoryEntryQuantitySetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LineItemStateTransition":
		new := LineItemStateTransitionMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderBillingAddressSet":
		new := OrderBillingAddressSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCreated":
		new := OrderCreatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomLineItemDiscountSet":
		new := OrderCustomLineItemDiscountSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerEmailSet":
		new := OrderCustomerEmailSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerGroupSet":
		new := OrderCustomerGroupSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerSet":
		new := OrderCustomerSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDeleted":
		new := OrderDeletedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeAdded":
		new := OrderDiscountCodeAddedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeRemoved":
		new := OrderDiscountCodeRemovedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeStateSet":
		new := OrderDiscountCodeStateSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderEditApplied":
		new := OrderEditAppliedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderImported":
		new := OrderImportedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderLineItemAdded":
		new := OrderLineItemAddedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderLineItemDiscountSet":
		new := OrderLineItemDiscountSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderPaymentStateChanged":
		new := OrderPaymentStateChangedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReturnInfoAdded":
		new := OrderReturnInfoAddedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderReturnShipmentStateChanged":
		new := OrderReturnShipmentStateChangedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShipmentStateChanged":
		new := OrderShipmentStateChangedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingAddressSet":
		new := OrderShippingAddressSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingInfoSet":
		new := OrderShippingInfoSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingRateInputSet":
		new := OrderShippingRateInputSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderStateChanged":
		new := OrderStateChangedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderStateTransition":
		new := OrderStateTransitionMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderStoreSet":
		new := OrderStoreSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelAddedToDelivery":
		new := ParcelAddedToDeliveryMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelItemsUpdated":
		new := ParcelItemsUpdatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelMeasurementsUpdated":
		new := ParcelMeasurementsUpdatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelRemovedFromDelivery":
		new := ParcelRemovedFromDeliveryMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelTrackingDataUpdated":
		new := ParcelTrackingDataUpdatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentCreated":
		new := PaymentCreatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentInteractionAdded":
		new := PaymentInteractionAddedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentStatusInterfaceCodeSet":
		new := PaymentStatusInterfaceCodeSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentStatusStateTransition":
		new := PaymentStatusStateTransitionMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentTransactionAdded":
		new := PaymentTransactionAddedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentTransactionStateChanged":
		new := PaymentTransactionStateChangedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductAddedToCategory":
		new := ProductAddedToCategoryMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductCreated":
		new := ProductCreatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductDeleted":
		new := ProductDeletedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductImageAdded":
		new := ProductImageAddedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPriceDiscountsSet":
		new := ProductPriceDiscountsSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPriceExternalDiscountSet":
		new := ProductPriceExternalDiscountSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPublished":
		new := ProductPublishedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductRemovedFromCategory":
		new := ProductRemovedFromCategoryMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductRevertedStagedChanges":
		new := ProductRevertedStagedChangesMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductSlugChanged":
		new := ProductSlugChangedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductStateTransition":
		new := ProductStateTransitionMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductUnpublished":
		new := ProductUnpublishedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductVariantAdded":
		new := ProductVariantAddedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductVariantDeleted":
		new := ProductVariantDeletedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReviewCreated":
		new := ReviewCreatedMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReviewRatingSet":
		new := ReviewRatingSetMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReviewStateTransition":
		new := ReviewStateTransitionMessage{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CategoryCreatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Category                        Category                 `json:"category"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CategoryCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias CategoryCreatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CategoryCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias CategoryCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CategoryCreated", Alias: (*Alias)(&obj)})
}

type CategorySlugChangedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Slug                            LocalizedString          `json:"slug"`
	OldSlug                         *LocalizedString         `json:"oldSlug,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CategorySlugChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias CategorySlugChangedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CategorySlugChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias CategorySlugChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CategorySlugChanged", Alias: (*Alias)(&obj)})
}

type CustomLineItemStateTransitionMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	CustomLineItemId                string                   `json:"customLineItemId"`
	TransitionDate                  time.Time                `json:"transitionDate"`
	Quantity                        int                      `json:"quantity"`
	FromState                       StateReference           `json:"fromState"`
	ToState                         StateReference           `json:"toState"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomLineItemStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomLineItemStateTransitionMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CustomLineItemStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomLineItemStateTransition", Alias: (*Alias)(&obj)})
}

type CustomerAddressAddedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Address                         Address                  `json:"address"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerAddressAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerAddressAddedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerAddressAdded", Alias: (*Alias)(&obj)})
}

type CustomerAddressChangedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Address                         Address                  `json:"address"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerAddressChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerAddressChangedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerAddressChanged", Alias: (*Alias)(&obj)})
}

type CustomerAddressRemovedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Address                         Address                  `json:"address"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerAddressRemovedMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerAddressRemovedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerAddressRemoved", Alias: (*Alias)(&obj)})
}

type CustomerCompanyNameSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	CompanyName                     string                   `json:"companyName"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerCompanyNameSetMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerCompanyNameSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CustomerCompanyNameSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerCompanyNameSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerCompanyNameSet", Alias: (*Alias)(&obj)})
}

type CustomerCreatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Customer                        Customer                 `json:"customer"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerCreatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CustomerCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerCreated", Alias: (*Alias)(&obj)})
}

type CustomerDateOfBirthSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	DateOfBirth                     Date                     `json:"dateOfBirth"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerDateOfBirthSetMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerDateOfBirthSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CustomerDateOfBirthSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerDateOfBirthSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerDateOfBirthSet", Alias: (*Alias)(&obj)})
}

type CustomerEmailChangedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Email                           string                   `json:"email"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerEmailChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerEmailChangedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CustomerEmailChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerEmailChanged", Alias: (*Alias)(&obj)})
}

type CustomerEmailVerifiedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerEmailVerifiedMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerEmailVerifiedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CustomerEmailVerifiedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailVerifiedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerEmailVerified", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	CustomerGroup                   CustomerGroupReference   `json:"customerGroup"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerGroupSetMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerGroupSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerGroupSet", Alias: (*Alias)(&obj)})
}

type CustomerPasswordUpdatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	// true, if password has been updated during Customer's Password Reset workflow.
	Reset bool `json:"reset"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerPasswordUpdatedMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerPasswordUpdatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CustomerPasswordUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerPasswordUpdatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerPasswordUpdated", Alias: (*Alias)(&obj)})
}

type DeliveryAddedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Delivery                        Delivery                 `json:"delivery"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeliveryAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias DeliveryAddedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryAdded", Alias: (*Alias)(&obj)})
}

type DeliveryAddressSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	DeliveryId                      string                   `json:"deliveryId"`
	Address                         *Address                 `json:"address,omitEmpty"`
	OldAddress                      *Address                 `json:"oldAddress,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeliveryAddressSetMessage) UnmarshalJSON(data []byte) error {
	type Alias DeliveryAddressSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryAddressSetMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddressSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryAddressSet", Alias: (*Alias)(&obj)})
}

type DeliveryItemsUpdatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	DeliveryId                      string                   `json:"deliveryId"`
	Items                           []DeliveryItem           `json:"items"`
	OldItems                        []DeliveryItem           `json:"oldItems"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeliveryItemsUpdatedMessage) UnmarshalJSON(data []byte) error {
	type Alias DeliveryItemsUpdatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryItemsUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryItemsUpdatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryItemsUpdated", Alias: (*Alias)(&obj)})
}

type DeliveryRemovedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Delivery                        Delivery                 `json:"delivery"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DeliveryRemovedMessage) UnmarshalJSON(data []byte) error {
	type Alias DeliveryRemovedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryRemoved", Alias: (*Alias)(&obj)})
}

type InventoryEntryCreatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	InventoryEntry                  InventoryEntry           `json:"inventoryEntry"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InventoryEntryCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias InventoryEntryCreatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "InventoryEntryCreated", Alias: (*Alias)(&obj)})
}

type InventoryEntryDeletedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Sku                             string                   `json:"sku"`
	SupplyChannel                   ChannelReference         `json:"supplyChannel"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InventoryEntryDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias InventoryEntryDeletedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "InventoryEntryDeleted", Alias: (*Alias)(&obj)})
}

type InventoryEntryQuantitySetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	OldQuantityOnStock              int                      `json:"oldQuantityOnStock"`
	NewQuantityOnStock              int                      `json:"newQuantityOnStock"`
	OldAvailableQuantity            int                      `json:"oldAvailableQuantity"`
	NewAvailableQuantity            int                      `json:"newAvailableQuantity"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InventoryEntryQuantitySetMessage) UnmarshalJSON(data []byte) error {
	type Alias InventoryEntryQuantitySetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryQuantitySetMessage) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryQuantitySetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "InventoryEntryQuantitySet", Alias: (*Alias)(&obj)})
}

type LineItemStateTransitionMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	LineItemId                      string                   `json:"lineItemId"`
	TransitionDate                  time.Time                `json:"transitionDate"`
	Quantity                        int                      `json:"quantity"`
	FromState                       StateReference           `json:"fromState"`
	ToState                         StateReference           `json:"toState"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *LineItemStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias LineItemStateTransitionMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj LineItemStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias LineItemStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "LineItemStateTransition", Alias: (*Alias)(&obj)})
}

type MessageConfiguration struct {
	Enabled                 bool `json:"enabled"`
	DeleteDaysAfterCreation int  `json:"deleteDaysAfterCreation,omitEmpty"`
}

type MessageConfigurationDraft struct {
	Enabled                 bool `json:"enabled"`
	DeleteDaysAfterCreation int  `json:"deleteDaysAfterCreation"`
}

type MessagePagedQueryResponse struct {
	Limit   int       `json:"limit"`
	Count   int       `json:"count"`
	Total   int       `json:"total,omitEmpty"`
	Offset  int       `json:"offset"`
	Results []Message `json:"results"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MessagePagedQueryResponse) UnmarshalJSON(data []byte) error {
	type Alias MessagePagedQueryResponse
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type OrderBillingAddressSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Address                         *Address                 `json:"address,omitEmpty"`
	OldAddress                      *Address                 `json:"oldAddress,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderBillingAddressSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderBillingAddressSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderBillingAddressSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderBillingAddressSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderBillingAddressSet", Alias: (*Alias)(&obj)})
}

type OrderCreatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Order                           Order                    `json:"order"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderCreatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCreated", Alias: (*Alias)(&obj)})
}

type OrderCustomLineItemDiscountSetMessage struct {
	Id                              string                               `json:"id"`
	Version                         int                                  `json:"version"`
	CreatedAt                       time.Time                            `json:"createdAt"`
	LastModifiedAt                  time.Time                            `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy                      `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy                           `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                                  `json:"sequenceNumber"`
	Resource                        Reference                            `json:"resource"`
	ResourceVersion                 int                                  `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers             `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	CustomLineItemId                string                               `json:"customLineItemId"`
	DiscountedPricePerQuantity      []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	TaxedPrice                      *TaxedItemPrice                      `json:"taxedPrice,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderCustomLineItemDiscountSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderCustomLineItemDiscountSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomLineItemDiscountSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemDiscountSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerEmailSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Email                           string                   `json:"email,omitEmpty"`
	OldEmail                        string                   `json:"oldEmail,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderCustomerEmailSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderCustomerEmailSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerEmailSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerEmailSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomerEmailSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerGroupSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	CustomerGroup                   CustomerGroupReference   `json:"customerGroup,omitEmpty"`
	OldCustomerGroup                CustomerGroupReference   `json:"oldCustomerGroup,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderCustomerGroupSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderCustomerGroupSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerGroupSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerGroupSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomerGroupSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Customer                        CustomerReference        `json:"customer,omitEmpty"`
	CustomerGroup                   CustomerGroupReference   `json:"customerGroup,omitEmpty"`
	OldCustomer                     CustomerReference        `json:"oldCustomer,omitEmpty"`
	OldCustomerGroup                CustomerGroupReference   `json:"oldCustomerGroup,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderCustomerSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderCustomerSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomerSet", Alias: (*Alias)(&obj)})
}

type OrderDeletedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Order                           Order                    `json:"order"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderDeletedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDeleted", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeAddedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	DiscountCode                    DiscountCodeReference    `json:"discountCode"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderDiscountCodeAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderDiscountCodeAddedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDiscountCodeAdded", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeRemovedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	DiscountCode                    DiscountCodeReference    `json:"discountCode"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderDiscountCodeRemovedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderDiscountCodeRemovedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDiscountCodeRemoved", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeStateSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	DiscountCode                    DiscountCodeReference    `json:"discountCode"`
	State                           DiscountCodeState        `json:"state"`
	OldState                        DiscountCodeState        `json:"oldState,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderDiscountCodeStateSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderDiscountCodeStateSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeStateSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeStateSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDiscountCodeStateSet", Alias: (*Alias)(&obj)})
}

type OrderEditAppliedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Edit                            OrderEditReference       `json:"edit"`
	Result                          OrderEditApplied         `json:"result"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditAppliedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderEditAppliedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditAppliedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderEditAppliedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderEditApplied", Alias: (*Alias)(&obj)})
}

type OrderImportedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Order                           Order                    `json:"order"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderImportedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderImportedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderImportedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderImportedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderImported", Alias: (*Alias)(&obj)})
}

type OrderLineItemAddedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	LineItem                        LineItem                 `json:"lineItem"`
	AddedQuantity                   int                      `json:"addedQuantity"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderLineItemAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderLineItemAddedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderLineItemAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLineItemAdded", Alias: (*Alias)(&obj)})
}

type OrderLineItemDiscountSetMessage struct {
	Id                              string                               `json:"id"`
	Version                         int                                  `json:"version"`
	CreatedAt                       time.Time                            `json:"createdAt"`
	LastModifiedAt                  time.Time                            `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy                      `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy                           `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                                  `json:"sequenceNumber"`
	Resource                        Reference                            `json:"resource"`
	ResourceVersion                 int                                  `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers             `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	LineItemId                      string                               `json:"lineItemId"`
	DiscountedPricePerQuantity      []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	TotalPrice                      Money                                `json:"totalPrice"`
	TaxedPrice                      *TaxedItemPrice                      `json:"taxedPrice,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderLineItemDiscountSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderLineItemDiscountSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderLineItemDiscountSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemDiscountSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

type OrderPaymentStateChangedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	PaymentState                    PaymentState             `json:"paymentState"`
	OldPaymentState                 PaymentState             `json:"oldPaymentState,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderPaymentStateChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderPaymentStateChangedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderPaymentStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderPaymentStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderPaymentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderReturnInfoAddedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	ReturnInfo                      ReturnInfo               `json:"returnInfo"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderReturnInfoAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderReturnInfoAddedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderReturnInfoAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnInfoAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReturnInfoAdded", Alias: (*Alias)(&obj)})
}

type OrderReturnShipmentStateChangedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	ReturnItemId                    string                   `json:"returnItemId"`
	ReturnShipmentState             ReturnShipmentState      `json:"returnShipmentState"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderReturnShipmentStateChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderReturnShipmentStateChangedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderReturnShipmentStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnShipmentStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderReturnShipmentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderShipmentStateChangedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	ShipmentState                   ShipmentState            `json:"shipmentState"`
	OldShipmentState                ShipmentState            `json:"oldShipmentState,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderShipmentStateChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderShipmentStateChangedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderShipmentStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShipmentStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShipmentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderShippingAddressSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Address                         *Address                 `json:"address,omitEmpty"`
	OldAddress                      *Address                 `json:"oldAddress,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderShippingAddressSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderShippingAddressSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingAddressSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingAddressSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShippingAddressSet", Alias: (*Alias)(&obj)})
}

type OrderShippingInfoSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	ShippingInfo                    *ShippingInfo            `json:"shippingInfo,omitEmpty"`
	OldShippingInfo                 *ShippingInfo            `json:"oldShippingInfo,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderShippingInfoSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderShippingInfoSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingInfoSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingInfoSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShippingInfoSet", Alias: (*Alias)(&obj)})
}

type OrderShippingRateInputSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	ShippingRateInput               ShippingRateInput        `json:"shippingRateInput,omitEmpty"`
	OldShippingRateInput            ShippingRateInput        `json:"oldShippingRateInput,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderShippingRateInputSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderShippingRateInputSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}
	if obj.OldShippingRateInput != nil {
		var err error
		obj.OldShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.OldShippingRateInput)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingRateInputSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingRateInputSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShippingRateInputSet", Alias: (*Alias)(&obj)})
}

type OrderStateChangedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	OrderState                      OrderState               `json:"orderState"`
	OldOrderState                   OrderState               `json:"oldOrderState"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderStateChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderStateChangedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderStateChanged", Alias: (*Alias)(&obj)})
}

type OrderStateTransitionMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	State                           StateReference           `json:"state"`
	OldState                        StateReference           `json:"oldState,omitEmpty"`
	Force                           bool                     `json:"force"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderStateTransitionMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderStateTransition", Alias: (*Alias)(&obj)})
}

type OrderStoreSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Store                           StoreKeyReference        `json:"store"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderStoreSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderStoreSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderStoreSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderStoreSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderStoreSet", Alias: (*Alias)(&obj)})
}

type ParcelAddedToDeliveryMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Delivery                        Delivery                 `json:"delivery"`
	Parcel                          Parcel                   `json:"parcel"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ParcelAddedToDeliveryMessage) UnmarshalJSON(data []byte) error {
	type Alias ParcelAddedToDeliveryMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ParcelAddedToDeliveryMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelAddedToDeliveryMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelAddedToDelivery", Alias: (*Alias)(&obj)})
}

type ParcelItemsUpdatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	ParcelId                        string                   `json:"parcelId"`
	DeliveryId                      string                   `json:"deliveryId,omitEmpty"`
	Items                           []DeliveryItem           `json:"items"`
	OldItems                        []DeliveryItem           `json:"oldItems"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ParcelItemsUpdatedMessage) UnmarshalJSON(data []byte) error {
	type Alias ParcelItemsUpdatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ParcelItemsUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelItemsUpdatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelItemsUpdated", Alias: (*Alias)(&obj)})
}

type ParcelMeasurementsUpdatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	DeliveryId                      string                   `json:"deliveryId"`
	ParcelId                        string                   `json:"parcelId"`
	Measurements                    *ParcelMeasurements      `json:"measurements,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ParcelMeasurementsUpdatedMessage) UnmarshalJSON(data []byte) error {
	type Alias ParcelMeasurementsUpdatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ParcelMeasurementsUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelMeasurementsUpdatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelMeasurementsUpdated", Alias: (*Alias)(&obj)})
}

type ParcelRemovedFromDeliveryMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	DeliveryId                      string                   `json:"deliveryId"`
	Parcel                          Parcel                   `json:"parcel"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ParcelRemovedFromDeliveryMessage) UnmarshalJSON(data []byte) error {
	type Alias ParcelRemovedFromDeliveryMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ParcelRemovedFromDeliveryMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelRemovedFromDeliveryMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelRemovedFromDelivery", Alias: (*Alias)(&obj)})
}

type ParcelTrackingDataUpdatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	DeliveryId                      string                   `json:"deliveryId"`
	ParcelId                        string                   `json:"parcelId"`
	TrackingData                    *TrackingData            `json:"trackingData,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ParcelTrackingDataUpdatedMessage) UnmarshalJSON(data []byte) error {
	type Alias ParcelTrackingDataUpdatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ParcelTrackingDataUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelTrackingDataUpdatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelTrackingDataUpdated", Alias: (*Alias)(&obj)})
}

type PaymentCreatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Payment                         Payment                  `json:"payment"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias PaymentCreatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj PaymentCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentCreated", Alias: (*Alias)(&obj)})
}

type PaymentInteractionAddedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Interaction                     CustomFields             `json:"interaction"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentInteractionAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias PaymentInteractionAddedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj PaymentInteractionAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentInteractionAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentInteractionAdded", Alias: (*Alias)(&obj)})
}

type PaymentStatusInterfaceCodeSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	PaymentId                       string                   `json:"paymentId"`
	InterfaceCode                   string                   `json:"interfaceCode"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentStatusInterfaceCodeSetMessage) UnmarshalJSON(data []byte) error {
	type Alias PaymentStatusInterfaceCodeSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj PaymentStatusInterfaceCodeSetMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusInterfaceCodeSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentStatusInterfaceCodeSet", Alias: (*Alias)(&obj)})
}

type PaymentStatusStateTransitionMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	State                           StateReference           `json:"state"`
	Force                           bool                     `json:"force"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentStatusStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias PaymentStatusStateTransitionMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj PaymentStatusStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentStatusStateTransition", Alias: (*Alias)(&obj)})
}

type PaymentTransactionAddedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Transaction                     Transaction              `json:"transaction"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentTransactionAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias PaymentTransactionAddedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransactionAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentTransactionAdded", Alias: (*Alias)(&obj)})
}

type PaymentTransactionStateChangedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	TransactionId                   string                   `json:"transactionId"`
	State                           TransactionState         `json:"state"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentTransactionStateChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias PaymentTransactionStateChangedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransactionStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentTransactionStateChanged", Alias: (*Alias)(&obj)})
}

type ProductAddedToCategoryMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Category                        CategoryReference        `json:"category"`
	Staged                          bool                     `json:"staged"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductAddedToCategoryMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductAddedToCategoryMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddedToCategoryMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductAddedToCategoryMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductAddedToCategory", Alias: (*Alias)(&obj)})
}

type ProductCreatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	ProductProjection               ProductProjection        `json:"productProjection"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductCreatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductCreated", Alias: (*Alias)(&obj)})
}

type ProductDeletedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	RemovedImageUrls                []string                 `json:"removedImageUrls"`
	CurrentProjection               ProductProjection        `json:"currentProjection"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductDeletedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductDeleted", Alias: (*Alias)(&obj)})
}

type ProductImageAddedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	VariantId                       int                      `json:"variantId"`
	Image                           Image                    `json:"image"`
	Staged                          bool                     `json:"staged"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductImageAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductImageAddedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductImageAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductImageAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductImageAdded", Alias: (*Alias)(&obj)})
}

type ProductPriceDiscountsSetMessage struct {
	Id                              string                                 `json:"id"`
	Version                         int                                    `json:"version"`
	CreatedAt                       time.Time                              `json:"createdAt"`
	LastModifiedAt                  time.Time                              `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy                        `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy                             `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                                    `json:"sequenceNumber"`
	Resource                        Reference                              `json:"resource"`
	ResourceVersion                 int                                    `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers               `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	UpdatedPrices                   []ProductPriceDiscountsSetUpdatedPrice `json:"updatedPrices"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductPriceDiscountsSetMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductPriceDiscountsSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductPriceDiscountsSetMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductPriceDiscountsSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductPriceDiscountsSet", Alias: (*Alias)(&obj)})
}

type ProductPriceDiscountsSetUpdatedPrice struct {
	VariantId  int              `json:"variantId"`
	VariantKey string           `json:"variantKey,omitEmpty"`
	Sku        string           `json:"sku,omitEmpty"`
	PriceId    string           `json:"priceId"`
	Discounted *DiscountedPrice `json:"discounted,omitEmpty"`
	Staged     bool             `json:"staged"`
}

type ProductPriceExternalDiscountSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	VariantId                       int                      `json:"variantId"`
	VariantKey                      string                   `json:"variantKey,omitEmpty"`
	Sku                             string                   `json:"sku,omitEmpty"`
	PriceId                         string                   `json:"priceId"`
	Discounted                      *DiscountedPrice         `json:"discounted,omitEmpty"`
	Staged                          bool                     `json:"staged"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductPriceExternalDiscountSetMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductPriceExternalDiscountSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductPriceExternalDiscountSetMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductPriceExternalDiscountSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductPriceExternalDiscountSet", Alias: (*Alias)(&obj)})
}

type ProductPublishedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	RemovedImageUrls                []string                 `json:"removedImageUrls"`
	ProductProjection               ProductProjection        `json:"productProjection"`
	Scope                           ProductPublishScope      `json:"scope"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductPublishedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductPublishedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductPublishedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductPublishedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductPublished", Alias: (*Alias)(&obj)})
}

type ProductRemovedFromCategoryMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Category                        CategoryReference        `json:"category"`
	Staged                          bool                     `json:"staged"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductRemovedFromCategoryMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductRemovedFromCategoryMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemovedFromCategoryMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductRemovedFromCategoryMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductRemovedFromCategory", Alias: (*Alias)(&obj)})
}

type ProductRevertedStagedChangesMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	RemovedImageUrls                []string                 `json:"removedImageUrls"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductRevertedStagedChangesMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductRevertedStagedChangesMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductRevertedStagedChangesMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertedStagedChangesMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductRevertedStagedChanges", Alias: (*Alias)(&obj)})
}

type ProductSlugChangedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Slug                            LocalizedString          `json:"slug"`
	OldSlug                         *LocalizedString         `json:"oldSlug,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSlugChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductSlugChangedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductSlugChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductSlugChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSlugChanged", Alias: (*Alias)(&obj)})
}

type ProductStateTransitionMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	State                           StateReference           `json:"state"`
	Force                           bool                     `json:"force"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductStateTransitionMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductStateTransition", Alias: (*Alias)(&obj)})
}

type ProductUnpublishedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductUnpublishedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductUnpublishedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductUnpublishedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductUnpublishedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductUnpublished", Alias: (*Alias)(&obj)})
}

type ProductVariantAddedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Variant                         ProductVariant           `json:"variant"`
	Staged                          bool                     `json:"staged"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductVariantAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductVariantAddedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductVariantAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductVariantAdded", Alias: (*Alias)(&obj)})
}

type ProductVariantDeletedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Variant                         ProductVariant           `json:"variant"`
	RemovedImageUrls                []string                 `json:"removedImageUrls"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductVariantDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductVariantDeletedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ProductVariantDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductVariantDeleted", Alias: (*Alias)(&obj)})
}

type ReviewCreatedMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Review                          Review                   `json:"review"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias ReviewCreatedMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ReviewCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ReviewCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewCreated", Alias: (*Alias)(&obj)})
}

type ReviewRatingSetMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	OldRating                       float64                  `json:"oldRating,omitEmpty"`
	NewRating                       float64                  `json:"newRating,omitEmpty"`
	IncludedInStatistics            bool                     `json:"includedInStatistics"`
	Target                          Reference                `json:"target,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewRatingSetMessage) UnmarshalJSON(data []byte) error {
	type Alias ReviewRatingSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorReference(obj.Target)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ReviewRatingSetMessage) MarshalJSON() ([]byte, error) {
	type Alias ReviewRatingSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewRatingSet", Alias: (*Alias)(&obj)})
}

type ReviewStateTransitionMessage struct {
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitEmpty"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitEmpty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	OldState                        StateReference           `json:"oldState"`
	NewState                        StateReference           `json:"newState"`
	OldIncludedInStatistics         bool                     `json:"oldIncludedInStatistics"`
	NewIncludedInStatistics         bool                     `json:"newIncludedInStatistics"`
	Target                          Reference                `json:"target"`
	Force                           bool                     `json:"force"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias ReviewStateTransitionMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorReference(obj.Target)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ReviewStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias ReviewStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewStateTransition", Alias: (*Alias)(&obj)})
}

type UserProvidedIdentifiers struct {
	Key            string           `json:"key,omitEmpty"`
	ExternalId     string           `json:"externalId,omitEmpty"`
	OrderNumber    string           `json:"orderNumber,omitEmpty"`
	CustomerNumber string           `json:"customerNumber,omitEmpty"`
	Sku            string           `json:"sku,omitEmpty"`
	Slug           *LocalizedString `json:"slug,omitEmpty"`
}

type MessagePayload interface{}

func mapDiscriminatorMessagePayload(input interface{}) (MessagePayload, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "CategoryCreated":
		new := CategoryCreatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CategorySlugChanged":
		new := CategorySlugChangedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomLineItemStateTransition":
		new := CustomLineItemStateTransitionMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressAdded":
		new := CustomerAddressAddedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressChanged":
		new := CustomerAddressChangedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressRemoved":
		new := CustomerAddressRemovedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerCompanyNameSet":
		new := CustomerCompanyNameSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerCreated":
		new := CustomerCreatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerDateOfBirthSet":
		new := CustomerDateOfBirthSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerEmailChanged":
		new := CustomerEmailChangedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerEmailVerified":
		new := CustomerEmailVerifiedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerGroupSet":
		new := CustomerGroupSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerPasswordUpdated":
		new := CustomerPasswordUpdatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryAdded":
		new := DeliveryAddedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryAddressSet":
		new := DeliveryAddressSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryItemsUpdated":
		new := DeliveryItemsUpdatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryRemoved":
		new := DeliveryRemovedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InventoryEntryCreated":
		new := InventoryEntryCreatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InventoryEntryDeleted":
		new := InventoryEntryDeletedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InventoryEntryQuantitySet":
		new := InventoryEntryQuantitySetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LineItemStateTransition":
		new := LineItemStateTransitionMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderBillingAddressSet":
		new := OrderBillingAddressSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCreated":
		new := OrderCreatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomLineItemDiscountSet":
		new := OrderCustomLineItemDiscountSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerEmailSet":
		new := OrderCustomerEmailSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerGroupSet":
		new := OrderCustomerGroupSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerSet":
		new := OrderCustomerSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDeleted":
		new := OrderDeletedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeAdded":
		new := OrderDiscountCodeAddedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeRemoved":
		new := OrderDiscountCodeRemovedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeStateSet":
		new := OrderDiscountCodeStateSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderEditApplied":
		new := OrderEditAppliedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderImported":
		new := OrderImportedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderLineItemAdded":
		new := OrderLineItemAddedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderLineItemDiscountSet":
		new := OrderLineItemDiscountSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderPaymentStateChanged":
		new := OrderPaymentStateChangedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReturnInfoAdded":
		new := OrderReturnInfoAddedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderReturnShipmentStateChanged":
		new := OrderReturnShipmentStateChangedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShipmentStateChanged":
		new := OrderShipmentStateChangedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingAddressSet":
		new := OrderShippingAddressSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingInfoSet":
		new := OrderShippingInfoSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingRateInputSet":
		new := OrderShippingRateInputSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderStateChanged":
		new := OrderStateChangedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderStateTransition":
		new := OrderStateTransitionMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderStoreSet":
		new := OrderStoreSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelAddedToDelivery":
		new := ParcelAddedToDeliveryMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelItemsUpdated":
		new := ParcelItemsUpdatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelMeasurementsUpdated":
		new := ParcelMeasurementsUpdatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelRemovedFromDelivery":
		new := ParcelRemovedFromDeliveryMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelTrackingDataUpdated":
		new := ParcelTrackingDataUpdatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentCreated":
		new := PaymentCreatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentInteractionAdded":
		new := PaymentInteractionAddedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentStatusInterfaceCodeSet":
		new := PaymentStatusInterfaceCodeSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentStatusStateTransition":
		new := PaymentStatusStateTransitionMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentTransactionAdded":
		new := PaymentTransactionAddedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentTransactionStateChanged":
		new := PaymentTransactionStateChangedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductAddedToCategory":
		new := ProductAddedToCategoryMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductCreated":
		new := ProductCreatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductDeleted":
		new := ProductDeletedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductImageAdded":
		new := ProductImageAddedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPriceDiscountsSet":
		new := ProductPriceDiscountsSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPriceExternalDiscountSet":
		new := ProductPriceExternalDiscountSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPublished":
		new := ProductPublishedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductRemovedFromCategory":
		new := ProductRemovedFromCategoryMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductRevertedStagedChanges":
		new := ProductRevertedStagedChangesMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductSlugChanged":
		new := ProductSlugChangedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductStateTransition":
		new := ProductStateTransitionMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductUnpublished":
		new := ProductUnpublishedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductVariantAdded":
		new := ProductVariantAddedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductVariantDeleted":
		new := ProductVariantDeletedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReviewCreated":
		new := ReviewCreatedMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReviewRatingSet":
		new := ReviewRatingSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReviewStateTransition":
		new := ReviewStateTransitionMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ShoppingListStoreSet":
		new := ShoppingListStoreSetMessagePayload{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CategoryCreatedMessagePayload struct {
	Category Category `json:"category"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CategoryCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CategoryCreated", Alias: (*Alias)(&obj)})
}

type CategorySlugChangedMessagePayload struct {
	Slug    LocalizedString  `json:"slug"`
	OldSlug *LocalizedString `json:"oldSlug,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySlugChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CategorySlugChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CategorySlugChanged", Alias: (*Alias)(&obj)})
}

type CustomLineItemStateTransitionMessagePayload struct {
	CustomLineItemId string         `json:"customLineItemId"`
	TransitionDate   time.Time      `json:"transitionDate"`
	Quantity         int            `json:"quantity"`
	FromState        StateReference `json:"fromState"`
	ToState          StateReference `json:"toState"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomLineItemStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomLineItemStateTransition", Alias: (*Alias)(&obj)})
}

type CustomerAddressAddedMessagePayload struct {
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerAddressAdded", Alias: (*Alias)(&obj)})
}

type CustomerAddressChangedMessagePayload struct {
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerAddressChanged", Alias: (*Alias)(&obj)})
}

type CustomerAddressRemovedMessagePayload struct {
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerAddressRemoved", Alias: (*Alias)(&obj)})
}

type CustomerCompanyNameSetMessagePayload struct {
	CompanyName string `json:"companyName"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerCompanyNameSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerCompanyNameSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerCompanyNameSet", Alias: (*Alias)(&obj)})
}

type CustomerCreatedMessagePayload struct {
	Customer Customer `json:"customer"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerCreated", Alias: (*Alias)(&obj)})
}

type CustomerDateOfBirthSetMessagePayload struct {
	DateOfBirth Date `json:"dateOfBirth"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerDateOfBirthSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerDateOfBirthSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerDateOfBirthSet", Alias: (*Alias)(&obj)})
}

type CustomerEmailChangedMessagePayload struct {
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerEmailChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerEmailChanged", Alias: (*Alias)(&obj)})
}

type CustomerEmailVerifiedMessagePayload struct {
}

// MarshalJSON override to set the discriminator value
func (obj CustomerEmailVerifiedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailVerifiedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerEmailVerified", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetMessagePayload struct {
	CustomerGroup CustomerGroupReference `json:"customerGroup"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerGroupSet", Alias: (*Alias)(&obj)})
}

type CustomerPasswordUpdatedMessagePayload struct {
	// true, if password has been updated during Customer's Password Reset workflow.
	Reset bool `json:"reset"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerPasswordUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerPasswordUpdatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerPasswordUpdated", Alias: (*Alias)(&obj)})
}

type DeliveryAddedMessagePayload struct {
	Delivery Delivery `json:"delivery"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryAdded", Alias: (*Alias)(&obj)})
}

type DeliveryAddressSetMessagePayload struct {
	DeliveryId string   `json:"deliveryId"`
	Address    *Address `json:"address,omitEmpty"`
	OldAddress *Address `json:"oldAddress,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddressSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryAddressSet", Alias: (*Alias)(&obj)})
}

type DeliveryItemsUpdatedMessagePayload struct {
	DeliveryId string         `json:"deliveryId"`
	Items      []DeliveryItem `json:"items"`
	OldItems   []DeliveryItem `json:"oldItems"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryItemsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryItemsUpdatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryItemsUpdated", Alias: (*Alias)(&obj)})
}

type DeliveryRemovedMessagePayload struct {
	Delivery Delivery `json:"delivery"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryRemoved", Alias: (*Alias)(&obj)})
}

type InventoryEntryCreatedMessagePayload struct {
	InventoryEntry InventoryEntry `json:"inventoryEntry"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "InventoryEntryCreated", Alias: (*Alias)(&obj)})
}

type InventoryEntryDeletedMessagePayload struct {
	Sku           string           `json:"sku"`
	SupplyChannel ChannelReference `json:"supplyChannel"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "InventoryEntryDeleted", Alias: (*Alias)(&obj)})
}

type InventoryEntryQuantitySetMessagePayload struct {
	OldQuantityOnStock   int `json:"oldQuantityOnStock"`
	NewQuantityOnStock   int `json:"newQuantityOnStock"`
	OldAvailableQuantity int `json:"oldAvailableQuantity"`
	NewAvailableQuantity int `json:"newAvailableQuantity"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryQuantitySetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryQuantitySetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "InventoryEntryQuantitySet", Alias: (*Alias)(&obj)})
}

type LineItemStateTransitionMessagePayload struct {
	LineItemId     string         `json:"lineItemId"`
	TransitionDate time.Time      `json:"transitionDate"`
	Quantity       int            `json:"quantity"`
	FromState      StateReference `json:"fromState"`
	ToState        StateReference `json:"toState"`
}

// MarshalJSON override to set the discriminator value
func (obj LineItemStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias LineItemStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "LineItemStateTransition", Alias: (*Alias)(&obj)})
}

type OrderBillingAddressSetMessagePayload struct {
	Address    *Address `json:"address,omitEmpty"`
	OldAddress *Address `json:"oldAddress,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderBillingAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderBillingAddressSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderBillingAddressSet", Alias: (*Alias)(&obj)})
}

type OrderCreatedMessagePayload struct {
	Order Order `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCreated", Alias: (*Alias)(&obj)})
}

type OrderCustomLineItemDiscountSetMessagePayload struct {
	CustomLineItemId           string                               `json:"customLineItemId"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomLineItemDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemDiscountSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerEmailSetMessagePayload struct {
	Email    string `json:"email,omitEmpty"`
	OldEmail string `json:"oldEmail,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerEmailSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerEmailSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomerEmailSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerGroupSetMessagePayload struct {
	CustomerGroup    CustomerGroupReference `json:"customerGroup,omitEmpty"`
	OldCustomerGroup CustomerGroupReference `json:"oldCustomerGroup,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerGroupSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerGroupSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomerGroupSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerSetMessagePayload struct {
	Customer         CustomerReference      `json:"customer,omitEmpty"`
	CustomerGroup    CustomerGroupReference `json:"customerGroup,omitEmpty"`
	OldCustomer      CustomerReference      `json:"oldCustomer,omitEmpty"`
	OldCustomerGroup CustomerGroupReference `json:"oldCustomerGroup,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomerSet", Alias: (*Alias)(&obj)})
}

type OrderDeletedMessagePayload struct {
	Order Order `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDeleted", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeAddedMessagePayload struct {
	DiscountCode DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDiscountCodeAdded", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeRemovedMessagePayload struct {
	DiscountCode DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDiscountCodeRemoved", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeStateSetMessagePayload struct {
	DiscountCode DiscountCodeReference `json:"discountCode"`
	State        DiscountCodeState     `json:"state"`
	OldState     DiscountCodeState     `json:"oldState,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeStateSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeStateSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDiscountCodeStateSet", Alias: (*Alias)(&obj)})
}

type OrderEditAppliedMessagePayload struct {
	Edit   OrderEditReference `json:"edit"`
	Result OrderEditApplied   `json:"result"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditAppliedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderEditAppliedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderEditApplied", Alias: (*Alias)(&obj)})
}

type OrderImportedMessagePayload struct {
	Order Order `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderImportedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderImportedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderImported", Alias: (*Alias)(&obj)})
}

type OrderLineItemAddedMessagePayload struct {
	LineItem      LineItem `json:"lineItem"`
	AddedQuantity int      `json:"addedQuantity"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderLineItemAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLineItemAdded", Alias: (*Alias)(&obj)})
}

type OrderLineItemDiscountSetMessagePayload struct {
	LineItemId                 string                               `json:"lineItemId"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	TotalPrice                 Money                                `json:"totalPrice"`
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderLineItemDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemDiscountSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

type OrderPaymentStateChangedMessagePayload struct {
	PaymentState    PaymentState `json:"paymentState"`
	OldPaymentState PaymentState `json:"oldPaymentState,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderPaymentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderPaymentStateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderPaymentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderReturnInfoAddedMessagePayload struct {
	ReturnInfo ReturnInfo `json:"returnInfo"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderReturnInfoAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnInfoAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReturnInfoAdded", Alias: (*Alias)(&obj)})
}

type OrderReturnShipmentStateChangedMessagePayload struct {
	ReturnItemId        string              `json:"returnItemId"`
	ReturnShipmentState ReturnShipmentState `json:"returnShipmentState"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderReturnShipmentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnShipmentStateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderReturnShipmentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderShipmentStateChangedMessagePayload struct {
	ShipmentState    ShipmentState `json:"shipmentState"`
	OldShipmentState ShipmentState `json:"oldShipmentState,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShipmentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShipmentStateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShipmentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderShippingAddressSetMessagePayload struct {
	Address    *Address `json:"address,omitEmpty"`
	OldAddress *Address `json:"oldAddress,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingAddressSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShippingAddressSet", Alias: (*Alias)(&obj)})
}

type OrderShippingInfoSetMessagePayload struct {
	ShippingInfo    *ShippingInfo `json:"shippingInfo,omitEmpty"`
	OldShippingInfo *ShippingInfo `json:"oldShippingInfo,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingInfoSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingInfoSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShippingInfoSet", Alias: (*Alias)(&obj)})
}

type OrderShippingRateInputSetMessagePayload struct {
	ShippingRateInput    ShippingRateInput `json:"shippingRateInput,omitEmpty"`
	OldShippingRateInput ShippingRateInput `json:"oldShippingRateInput,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderShippingRateInputSetMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias OrderShippingRateInputSetMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}
	if obj.OldShippingRateInput != nil {
		var err error
		obj.OldShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.OldShippingRateInput)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingRateInputSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingRateInputSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShippingRateInputSet", Alias: (*Alias)(&obj)})
}

type OrderStateChangedMessagePayload struct {
	OrderState    OrderState `json:"orderState"`
	OldOrderState OrderState `json:"oldOrderState"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderStateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderStateChanged", Alias: (*Alias)(&obj)})
}

type OrderStateTransitionMessagePayload struct {
	State StateReference `json:"state"`
	Force bool           `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderStateTransition", Alias: (*Alias)(&obj)})
}

type OrderStoreSetMessagePayload struct {
	Store StoreKeyReference `json:"store"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderStoreSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderStoreSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderStoreSet", Alias: (*Alias)(&obj)})
}

type ParcelAddedToDeliveryMessagePayload struct {
	Delivery Delivery `json:"delivery"`
	Parcel   Parcel   `json:"parcel"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelAddedToDeliveryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelAddedToDeliveryMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelAddedToDelivery", Alias: (*Alias)(&obj)})
}

type ParcelItemsUpdatedMessagePayload struct {
	ParcelId   string         `json:"parcelId"`
	DeliveryId string         `json:"deliveryId,omitEmpty"`
	Items      []DeliveryItem `json:"items"`
	OldItems   []DeliveryItem `json:"oldItems"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelItemsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelItemsUpdatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelItemsUpdated", Alias: (*Alias)(&obj)})
}

type ParcelMeasurementsUpdatedMessagePayload struct {
	DeliveryId   string              `json:"deliveryId"`
	ParcelId     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelMeasurementsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelMeasurementsUpdatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelMeasurementsUpdated", Alias: (*Alias)(&obj)})
}

type ParcelRemovedFromDeliveryMessagePayload struct {
	DeliveryId string `json:"deliveryId"`
	Parcel     Parcel `json:"parcel"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelRemovedFromDeliveryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelRemovedFromDeliveryMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelRemovedFromDelivery", Alias: (*Alias)(&obj)})
}

type ParcelTrackingDataUpdatedMessagePayload struct {
	DeliveryId   string        `json:"deliveryId"`
	ParcelId     string        `json:"parcelId"`
	TrackingData *TrackingData `json:"trackingData,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelTrackingDataUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelTrackingDataUpdatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelTrackingDataUpdated", Alias: (*Alias)(&obj)})
}

type PaymentCreatedMessagePayload struct {
	Payment Payment `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentCreated", Alias: (*Alias)(&obj)})
}

type PaymentInteractionAddedMessagePayload struct {
	Interaction CustomFields `json:"interaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentInteractionAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentInteractionAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentInteractionAdded", Alias: (*Alias)(&obj)})
}

type PaymentStatusInterfaceCodeSetMessagePayload struct {
	PaymentId     string `json:"paymentId"`
	InterfaceCode string `json:"interfaceCode"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentStatusInterfaceCodeSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusInterfaceCodeSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentStatusInterfaceCodeSet", Alias: (*Alias)(&obj)})
}

type PaymentStatusStateTransitionMessagePayload struct {
	State StateReference `json:"state"`
	Force bool           `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentStatusStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentStatusStateTransition", Alias: (*Alias)(&obj)})
}

type PaymentTransactionAddedMessagePayload struct {
	Transaction Transaction `json:"transaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransactionAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentTransactionAdded", Alias: (*Alias)(&obj)})
}

type PaymentTransactionStateChangedMessagePayload struct {
	TransactionId string           `json:"transactionId"`
	State         TransactionState `json:"state"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransactionStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionStateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentTransactionStateChanged", Alias: (*Alias)(&obj)})
}

type ProductAddedToCategoryMessagePayload struct {
	Category CategoryReference `json:"category"`
	Staged   bool              `json:"staged"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddedToCategoryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductAddedToCategoryMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductAddedToCategory", Alias: (*Alias)(&obj)})
}

type ProductCreatedMessagePayload struct {
	ProductProjection ProductProjection `json:"productProjection"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductCreated", Alias: (*Alias)(&obj)})
}

type ProductDeletedMessagePayload struct {
	RemovedImageUrls  []string          `json:"removedImageUrls"`
	CurrentProjection ProductProjection `json:"currentProjection"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductDeleted", Alias: (*Alias)(&obj)})
}

type ProductImageAddedMessagePayload struct {
	VariantId int   `json:"variantId"`
	Image     Image `json:"image"`
	Staged    bool  `json:"staged"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductImageAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductImageAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductImageAdded", Alias: (*Alias)(&obj)})
}

type ProductPriceDiscountsSetMessagePayload struct {
	UpdatedPrices []ProductPriceDiscountsSetUpdatedPrice `json:"updatedPrices"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPriceDiscountsSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductPriceDiscountsSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductPriceDiscountsSet", Alias: (*Alias)(&obj)})
}

type ProductPriceExternalDiscountSetMessagePayload struct {
	VariantId  int              `json:"variantId"`
	VariantKey string           `json:"variantKey,omitEmpty"`
	Sku        string           `json:"sku,omitEmpty"`
	PriceId    string           `json:"priceId"`
	Discounted *DiscountedPrice `json:"discounted,omitEmpty"`
	Staged     bool             `json:"staged"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPriceExternalDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductPriceExternalDiscountSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductPriceExternalDiscountSet", Alias: (*Alias)(&obj)})
}

type ProductPublishedMessagePayload struct {
	RemovedImageUrls  []string            `json:"removedImageUrls"`
	ProductProjection ProductProjection   `json:"productProjection"`
	Scope             ProductPublishScope `json:"scope"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPublishedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductPublishedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductPublished", Alias: (*Alias)(&obj)})
}

type ProductRemovedFromCategoryMessagePayload struct {
	Category CategoryReference `json:"category"`
	Staged   bool              `json:"staged"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemovedFromCategoryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductRemovedFromCategoryMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductRemovedFromCategory", Alias: (*Alias)(&obj)})
}

type ProductRevertedStagedChangesMessagePayload struct {
	RemovedImageUrls []string `json:"removedImageUrls"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRevertedStagedChangesMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertedStagedChangesMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductRevertedStagedChanges", Alias: (*Alias)(&obj)})
}

type ProductSlugChangedMessagePayload struct {
	Slug    LocalizedString  `json:"slug"`
	OldSlug *LocalizedString `json:"oldSlug,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSlugChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductSlugChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSlugChanged", Alias: (*Alias)(&obj)})
}

type ProductStateTransitionMessagePayload struct {
	State StateReference `json:"state"`
	Force bool           `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductStateTransition", Alias: (*Alias)(&obj)})
}

type ProductUnpublishedMessagePayload struct {
}

// MarshalJSON override to set the discriminator value
func (obj ProductUnpublishedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductUnpublishedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductUnpublished", Alias: (*Alias)(&obj)})
}

type ProductVariantAddedMessagePayload struct {
	Variant ProductVariant `json:"variant"`
	Staged  bool           `json:"staged"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductVariantAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductVariantAdded", Alias: (*Alias)(&obj)})
}

type ProductVariantDeletedMessagePayload struct {
	Variant          ProductVariant `json:"variant"`
	RemovedImageUrls []string       `json:"removedImageUrls"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductVariantDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductVariantDeleted", Alias: (*Alias)(&obj)})
}

type ReviewCreatedMessagePayload struct {
	Review Review `json:"review"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewCreated", Alias: (*Alias)(&obj)})
}

type ReviewRatingSetMessagePayload struct {
	OldRating            float64   `json:"oldRating,omitEmpty"`
	NewRating            float64   `json:"newRating,omitEmpty"`
	IncludedInStatistics bool      `json:"includedInStatistics"`
	Target               Reference `json:"target,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewRatingSetMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias ReviewRatingSetMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorReference(obj.Target)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ReviewRatingSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewRatingSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewRatingSet", Alias: (*Alias)(&obj)})
}

type ReviewStateTransitionMessagePayload struct {
	OldState                StateReference `json:"oldState"`
	NewState                StateReference `json:"newState"`
	OldIncludedInStatistics bool           `json:"oldIncludedInStatistics"`
	NewIncludedInStatistics bool           `json:"newIncludedInStatistics"`
	Target                  Reference      `json:"target"`
	Force                   bool           `json:"force"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewStateTransitionMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias ReviewStateTransitionMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorReference(obj.Target)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ReviewStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewStateTransition", Alias: (*Alias)(&obj)})
}

type ShoppingListStoreSetMessagePayload struct {
	Store StoreKeyReference `json:"store"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListStoreSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListStoreSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ShoppingListStoreSet", Alias: (*Alias)(&obj)})
}
