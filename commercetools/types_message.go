// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type CategoryCreatedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Category        *Category `json:"category"`
}

type CategoryCreatedMessagePayload struct {
	Category *Category `json:"category"`
}

func (obj CategoryCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CategoryCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CategoryCreated", Alias: (*Alias)(&obj)})
}

type CategorySlugChangedMessage struct {
	Version         int              `json:"version"`
	LastModifiedAt  time.Time        `json:"lastModifiedAt"`
	ID              string           `json:"id"`
	CreatedAt       time.Time        `json:"createdAt"`
	Type            string           `json:"type"`
	SequenceNumber  int              `json:"sequenceNumber"`
	ResourceVersion int              `json:"resourceVersion"`
	Resource        Reference        `json:"resource"`
	Slug            *LocalizedString `json:"slug"`
}

type CategorySlugChangedMessagePayload struct {
	Slug *LocalizedString `json:"slug"`
}

func (obj CategorySlugChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CategorySlugChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CategorySlugChanged", Alias: (*Alias)(&obj)})
}

type CustomLineItemStateTransitionMessage struct {
	Version          int             `json:"version"`
	LastModifiedAt   time.Time       `json:"lastModifiedAt"`
	ID               string          `json:"id"`
	CreatedAt        time.Time       `json:"createdAt"`
	Type             string          `json:"type"`
	SequenceNumber   int             `json:"sequenceNumber"`
	ResourceVersion  int             `json:"resourceVersion"`
	Resource         Reference       `json:"resource"`
	TransitionDate   time.Time       `json:"transitionDate"`
	ToState          *StateReference `json:"toState"`
	Quantity         int             `json:"quantity"`
	FromState        *StateReference `json:"fromState"`
	CustomLineItemID string          `json:"customLineItemId"`
}

type CustomLineItemStateTransitionMessagePayload struct {
	TransitionDate   time.Time       `json:"transitionDate"`
	ToState          *StateReference `json:"toState"`
	Quantity         int             `json:"quantity"`
	FromState        *StateReference `json:"fromState"`
	CustomLineItemID string          `json:"customLineItemId"`
}

func (obj CustomLineItemStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomLineItemStateTransition", Alias: (*Alias)(&obj)})
}

type CustomerAddressAddedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Address         *Address  `json:"address"`
}

type CustomerAddressAddedMessagePayload struct {
	Address *Address `json:"address"`
}

func (obj CustomerAddressAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressAdded", Alias: (*Alias)(&obj)})
}

type CustomerAddressChangedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Address         *Address  `json:"address"`
}

type CustomerAddressChangedMessagePayload struct {
	Address *Address `json:"address"`
}

func (obj CustomerAddressChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressChanged", Alias: (*Alias)(&obj)})
}

type CustomerAddressRemovedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Address         *Address  `json:"address"`
}

type CustomerAddressRemovedMessagePayload struct {
	Address *Address `json:"address"`
}

func (obj CustomerAddressRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressRemovedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressRemoved", Alias: (*Alias)(&obj)})
}

type CustomerCompanyNameSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	CompanyName     string    `json:"companyName"`
}

type CustomerCompanyNameSetMessagePayload struct {
	CompanyName string `json:"companyName"`
}

func (obj CustomerCompanyNameSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerCompanyNameSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerCompanyNameSet", Alias: (*Alias)(&obj)})
}

type CustomerCreatedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Customer        *Customer `json:"customer"`
}

type CustomerCreatedMessagePayload struct {
	Customer *Customer `json:"customer"`
}

func (obj CustomerCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerCreated", Alias: (*Alias)(&obj)})
}

type CustomerDateOfBirthSetMessage struct {
	Version         int         `json:"version"`
	LastModifiedAt  time.Time   `json:"lastModifiedAt"`
	ID              string      `json:"id"`
	CreatedAt       time.Time   `json:"createdAt"`
	Type            string      `json:"type"`
	SequenceNumber  int         `json:"sequenceNumber"`
	ResourceVersion int         `json:"resourceVersion"`
	Resource        Reference   `json:"resource"`
	DateOfBirth     interface{} `json:"dateOfBirth"`
}

type CustomerDateOfBirthSetMessagePayload struct {
	DateOfBirth interface{} `json:"dateOfBirth"`
}

func (obj CustomerDateOfBirthSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerDateOfBirthSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerDateOfBirthSet", Alias: (*Alias)(&obj)})
}

type CustomerEmailChangedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Email           string    `json:"email"`
}

type CustomerEmailChangedMessagePayload struct {
	Email string `json:"email"`
}

func (obj CustomerEmailChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerEmailChanged", Alias: (*Alias)(&obj)})
}

type CustomerEmailVerifiedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
}

type CustomerEmailVerifiedMessagePayload struct{}

func (obj CustomerEmailVerifiedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailVerifiedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerEmailVerified", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetMessage struct {
	Version         int                     `json:"version"`
	LastModifiedAt  time.Time               `json:"lastModifiedAt"`
	ID              string                  `json:"id"`
	CreatedAt       time.Time               `json:"createdAt"`
	Type            string                  `json:"type"`
	SequenceNumber  int                     `json:"sequenceNumber"`
	ResourceVersion int                     `json:"resourceVersion"`
	Resource        Reference               `json:"resource"`
	CustomerGroup   *CustomerGroupReference `json:"customerGroup"`
}

type CustomerGroupSetMessagePayload struct {
	CustomerGroup *CustomerGroupReference `json:"customerGroup"`
}

func (obj CustomerGroupSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerGroupSet", Alias: (*Alias)(&obj)})
}

type DeliveryAddedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Delivery        *Delivery `json:"delivery"`
}

type DeliveryAddedMessagePayload struct {
	Delivery *Delivery `json:"delivery"`
}

func (obj DeliveryAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryAdded", Alias: (*Alias)(&obj)})
}

type DeliveryAddressSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	DeliveryID      string    `json:"deliveryId"`
	Address         *Address  `json:"address,omitempty"`
}

type DeliveryAddressSetMessagePayload struct {
	DeliveryID string   `json:"deliveryId"`
	Address    *Address `json:"address,omitempty"`
}

func (obj DeliveryAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddressSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryAddressSet", Alias: (*Alias)(&obj)})
}

type DeliveryItemsUpdatedMessage struct {
	Version         int            `json:"version"`
	LastModifiedAt  time.Time      `json:"lastModifiedAt"`
	ID              string         `json:"id"`
	CreatedAt       time.Time      `json:"createdAt"`
	Type            string         `json:"type"`
	SequenceNumber  int            `json:"sequenceNumber"`
	ResourceVersion int            `json:"resourceVersion"`
	Resource        Reference      `json:"resource"`
	Items           []DeliveryItem `json:"items"`
	DeliveryID      string         `json:"deliveryId"`
}

type DeliveryItemsUpdatedMessagePayload struct {
	Items      []DeliveryItem `json:"items"`
	DeliveryID string         `json:"deliveryId"`
}

func (obj DeliveryItemsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryItemsUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryItemsUpdated", Alias: (*Alias)(&obj)})
}

type DeliveryRemovedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Delivery        *Delivery `json:"delivery"`
}

type DeliveryRemovedMessagePayload struct {
	Delivery *Delivery `json:"delivery"`
}

func (obj DeliveryRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryRemovedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryRemoved", Alias: (*Alias)(&obj)})
}

type InventoryEntryDeletedMessage struct {
	Version         int               `json:"version"`
	LastModifiedAt  time.Time         `json:"lastModifiedAt"`
	ID              string            `json:"id"`
	CreatedAt       time.Time         `json:"createdAt"`
	Type            string            `json:"type"`
	SequenceNumber  int               `json:"sequenceNumber"`
	ResourceVersion int               `json:"resourceVersion"`
	Resource        Reference         `json:"resource"`
	SupplyChannel   *ChannelReference `json:"supplyChannel"`
	Sku             string            `json:"sku"`
}

type InventoryEntryDeletedMessagePayload struct {
	SupplyChannel *ChannelReference `json:"supplyChannel"`
	Sku           string            `json:"sku"`
}

func (obj InventoryEntryDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "InventoryEntryDeleted", Alias: (*Alias)(&obj)})
}

type LineItemStateTransitionMessage struct {
	Version         int             `json:"version"`
	LastModifiedAt  time.Time       `json:"lastModifiedAt"`
	ID              string          `json:"id"`
	CreatedAt       time.Time       `json:"createdAt"`
	Type            string          `json:"type"`
	SequenceNumber  int             `json:"sequenceNumber"`
	ResourceVersion int             `json:"resourceVersion"`
	Resource        Reference       `json:"resource"`
	TransitionDate  time.Time       `json:"transitionDate"`
	ToState         *StateReference `json:"toState"`
	Quantity        int             `json:"quantity"`
	LineItemID      string          `json:"lineItemId"`
	FromState       *StateReference `json:"fromState"`
}

type LineItemStateTransitionMessagePayload struct {
	TransitionDate time.Time       `json:"transitionDate"`
	ToState        *StateReference `json:"toState"`
	Quantity       int             `json:"quantity"`
	LineItemID     string          `json:"lineItemId"`
	FromState      *StateReference `json:"fromState"`
}

func (obj LineItemStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias LineItemStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "LineItemStateTransition", Alias: (*Alias)(&obj)})
}

type Message struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
}

func (obj *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		obj.Resource = AbstractReferenceDiscriminatorMapping(obj.Resource)
	}

	return nil
}

type MessageConfiguration struct {
	Enabled                 bool    `json:"enabled"`
	DeleteDaysAfterCreation float64 `json:"deleteDaysAfterCreation,omitempty"`
}

type MessageConfigurationDraft struct {
	Enabled                 bool    `json:"enabled"`
	DeleteDaysAfterCreation float64 `json:"deleteDaysAfterCreation"`
}

type MessagePagedQueryResponse struct {
	Total   int       `json:"total,omitempty"`
	Offset  int       `json:"offset"`
	Count   int       `json:"count"`
	Results []Message `json:"results"`
}

type MessagePayload interface{}
type AbstractMessagePayload struct{}

func AbstractMessagePayloadDiscriminatorMapping(input MessagePayload) MessagePayload {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "CategoryCreated":
		new := CategoryCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CategorySlugChanged":
		new := CategorySlugChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomLineItemStateTransition":
		new := CustomLineItemStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerAddressAdded":
		new := CustomerAddressAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerAddressChanged":
		new := CustomerAddressChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerAddressRemoved":
		new := CustomerAddressRemovedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerCompanyNameSet":
		new := CustomerCompanyNameSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerCreated":
		new := CustomerCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerDateOfBirthSet":
		new := CustomerDateOfBirthSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerEmailChanged":
		new := CustomerEmailChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerEmailVerified":
		new := CustomerEmailVerifiedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "CustomerGroupSet":
		new := CustomerGroupSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "DeliveryAdded":
		new := DeliveryAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "DeliveryAddressSet":
		new := DeliveryAddressSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "DeliveryItemsUpdated":
		new := DeliveryItemsUpdatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "DeliveryRemoved":
		new := DeliveryRemovedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "InventoryEntryDeleted":
		new := InventoryEntryDeletedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "LineItemStateTransition":
		new := LineItemStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderBillingAddressSet":
		new := OrderBillingAddressSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderCreated":
		new := OrderCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderCustomLineItemDiscountSet":
		new := OrderCustomLineItemDiscountSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderCustomerEmailSet":
		new := OrderCustomerEmailSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderCustomerSet":
		new := OrderCustomerSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderDeleted":
		new := OrderDeletedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderDiscountCodeAdded":
		new := OrderDiscountCodeAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderDiscountCodeRemoved":
		new := OrderDiscountCodeRemovedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderDiscountCodeStateSet":
		new := OrderDiscountCodeStateSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderEditApplied":
		new := OrderEditAppliedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderImported":
		new := OrderImportedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderLineItemDiscountSet":
		new := OrderLineItemDiscountSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderPaymentStateChanged":
		new := OrderPaymentChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ReturnInfoAdded":
		new := OrderReturnInfoAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderReturnShipmentStateChanged":
		new := OrderReturnShipmentStateChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderShipmentStateChanged":
		new := OrderShipmentStateChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderShippingAddressSet":
		new := OrderShippingAddressSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderShippingInfoSet":
		new := OrderShippingInfoSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderShippingRateInputSet":
		new := OrderShippingRateInputSetMessagePayload{}
		mapstructure.Decode(input, &new)
		if new.OldShippingRateInput != nil {
			new.OldShippingRateInput = AbstractShippingRateInputDiscriminatorMapping(new.OldShippingRateInput)
		}
		if new.ShippingRateInput != nil {
			new.ShippingRateInput = AbstractShippingRateInputDiscriminatorMapping(new.ShippingRateInput)
		}

		return new
	case "OrderStateChanged":
		new := OrderStateChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "OrderStateTransition":
		new := OrderStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ParcelAddedToDelivery":
		new := ParcelAddedToDeliveryMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ParcelItemsUpdated":
		new := ParcelItemsUpdatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ParcelMeasurementsUpdated":
		new := ParcelMeasurementsUpdatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ParcelRemovedFromDelivery":
		new := ParcelRemovedFromDeliveryMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ParcelTrackingDataUpdated":
		new := ParcelTrackingDataUpdatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentCreated":
		new := PaymentCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentInteractionAdded":
		new := PaymentInteractionAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentStatusInterfaceCodeSet":
		new := PaymentStatusInterfaceCodeSetMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentStatusStateTransition":
		new := PaymentStatusStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentTransactionAdded":
		new := PaymentTransactionAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "PaymentTransactionStateChanged":
		new := PaymentTransactionStateChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductCreated":
		new := ProductCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductDeleted":
		new := ProductDeletedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductImageAdded":
		new := ProductImageAddedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductPublished":
		new := ProductPublishedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductRevertedStagedChanges":
		new := ProductRevertedStagedChangesMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductSlugChanged":
		new := ProductSlugChangedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductStateTransition":
		new := ProductStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductUnpublished":
		new := ProductUnpublishedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ProductVariantDeleted":
		new := ProductVariantDeletedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ReviewCreated":
		new := ReviewCreatedMessagePayload{}
		mapstructure.Decode(input, &new)
		return new
	case "ReviewRatingSet":
		new := ReviewRatingSetMessagePayload{}
		mapstructure.Decode(input, &new)
		if new.Target != nil {
			new.Target = AbstractReferenceDiscriminatorMapping(new.Target)
		}

		return new
	case "ReviewStateTransition":
		new := ReviewStateTransitionMessagePayload{}
		mapstructure.Decode(input, &new)
		if new.Target != nil {
			new.Target = AbstractReferenceDiscriminatorMapping(new.Target)
		}

		return new
	}
	return nil
}

type OrderBillingAddressSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Address         *Address  `json:"address"`
}

type OrderBillingAddressSetMessagePayload struct {
	Address *Address `json:"address"`
}

func (obj OrderBillingAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderBillingAddressSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderBillingAddressSet", Alias: (*Alias)(&obj)})
}

type OrderCreatedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Order           *Order    `json:"order"`
}

type OrderCreatedMessagePayload struct {
	Order *Order `json:"order"`
}

func (obj OrderCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCreated", Alias: (*Alias)(&obj)})
}

type OrderCustomLineItemDiscountSetMessage struct {
	Version                    int                                  `json:"version"`
	LastModifiedAt             time.Time                            `json:"lastModifiedAt"`
	ID                         string                               `json:"id"`
	CreatedAt                  time.Time                            `json:"createdAt"`
	Type                       string                               `json:"type"`
	SequenceNumber             int                                  `json:"sequenceNumber"`
	ResourceVersion            int                                  `json:"resourceVersion"`
	Resource                   Reference                            `json:"resource"`
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	CustomLineItemID           string                               `json:"customLineItemId"`
}

type OrderCustomLineItemDiscountSetMessagePayload struct {
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	CustomLineItemID           string                               `json:"customLineItemId"`
}

func (obj OrderCustomLineItemDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemDiscountSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerEmailSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Email           string    `json:"email"`
}

type OrderCustomerEmailSetMessagePayload struct {
	Email string `json:"email"`
}

func (obj OrderCustomerEmailSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerEmailSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomerEmailSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerSetMessage struct {
	Version          int                     `json:"version"`
	LastModifiedAt   time.Time               `json:"lastModifiedAt"`
	ID               string                  `json:"id"`
	CreatedAt        time.Time               `json:"createdAt"`
	Type             string                  `json:"type"`
	SequenceNumber   int                     `json:"sequenceNumber"`
	ResourceVersion  int                     `json:"resourceVersion"`
	Resource         Reference               `json:"resource"`
	OldCustomerGroup *CustomerGroupReference `json:"oldCustomerGroup,omitempty"`
	OldCustomer      *CustomerReference      `json:"oldCustomer,omitempty"`
	CustomerGroup    *CustomerGroupReference `json:"customerGroup,omitempty"`
	Customer         *CustomerReference      `json:"customer,omitempty"`
}

type OrderCustomerSetMessagePayload struct {
	OldCustomerGroup *CustomerGroupReference `json:"oldCustomerGroup,omitempty"`
	OldCustomer      *CustomerReference      `json:"oldCustomer,omitempty"`
	CustomerGroup    *CustomerGroupReference `json:"customerGroup,omitempty"`
	Customer         *CustomerReference      `json:"customer,omitempty"`
}

func (obj OrderCustomerSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomerSet", Alias: (*Alias)(&obj)})
}

type OrderDeletedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Order           *Order    `json:"order"`
}

type OrderDeletedMessagePayload struct {
	Order *Order `json:"order"`
}

func (obj OrderDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDeleted", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeAddedMessage struct {
	Version         int                    `json:"version"`
	LastModifiedAt  time.Time              `json:"lastModifiedAt"`
	ID              string                 `json:"id"`
	CreatedAt       time.Time              `json:"createdAt"`
	Type            string                 `json:"type"`
	SequenceNumber  int                    `json:"sequenceNumber"`
	ResourceVersion int                    `json:"resourceVersion"`
	Resource        Reference              `json:"resource"`
	DiscountCode    *DiscountCodeReference `json:"discountCode"`
}

type OrderDiscountCodeAddedMessagePayload struct {
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

func (obj OrderDiscountCodeAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeAdded", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeRemovedMessage struct {
	Version         int                    `json:"version"`
	LastModifiedAt  time.Time              `json:"lastModifiedAt"`
	ID              string                 `json:"id"`
	CreatedAt       time.Time              `json:"createdAt"`
	Type            string                 `json:"type"`
	SequenceNumber  int                    `json:"sequenceNumber"`
	ResourceVersion int                    `json:"resourceVersion"`
	Resource        Reference              `json:"resource"`
	DiscountCode    *DiscountCodeReference `json:"discountCode"`
}

type OrderDiscountCodeRemovedMessagePayload struct {
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

func (obj OrderDiscountCodeRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeRemovedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeRemoved", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeStateSetMessage struct {
	Version         int                    `json:"version"`
	LastModifiedAt  time.Time              `json:"lastModifiedAt"`
	ID              string                 `json:"id"`
	CreatedAt       time.Time              `json:"createdAt"`
	Type            string                 `json:"type"`
	SequenceNumber  int                    `json:"sequenceNumber"`
	ResourceVersion int                    `json:"resourceVersion"`
	Resource        Reference              `json:"resource"`
	State           DiscountCodeState      `json:"state"`
	OldState        DiscountCodeState      `json:"oldState,omitempty"`
	DiscountCode    *DiscountCodeReference `json:"discountCode"`
}

type OrderDiscountCodeStateSetMessagePayload struct {
	State        DiscountCodeState      `json:"state"`
	OldState     DiscountCodeState      `json:"oldState,omitempty"`
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

func (obj OrderDiscountCodeStateSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeStateSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeStateSet", Alias: (*Alias)(&obj)})
}

type OrderEditAppliedMessage struct {
	Version         int                 `json:"version"`
	LastModifiedAt  time.Time           `json:"lastModifiedAt"`
	ID              string              `json:"id"`
	CreatedAt       time.Time           `json:"createdAt"`
	Type            string              `json:"type"`
	SequenceNumber  int                 `json:"sequenceNumber"`
	ResourceVersion int                 `json:"resourceVersion"`
	Resource        Reference           `json:"resource"`
	Result          *OrderEditApplied   `json:"result"`
	Edit            *OrderEditReference `json:"edit"`
}

type OrderEditAppliedMessagePayload struct {
	Result *OrderEditApplied   `json:"result"`
	Edit   *OrderEditReference `json:"edit"`
}

func (obj OrderEditAppliedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderEditAppliedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderEditApplied", Alias: (*Alias)(&obj)})
}

type OrderImportedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Order           *Order    `json:"order"`
}

type OrderImportedMessagePayload struct {
	Order *Order `json:"order"`
}

func (obj OrderImportedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderImportedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderImported", Alias: (*Alias)(&obj)})
}

type OrderLineItemDiscountSetMessage struct {
	Version                    int                                  `json:"version"`
	LastModifiedAt             time.Time                            `json:"lastModifiedAt"`
	ID                         string                               `json:"id"`
	CreatedAt                  time.Time                            `json:"createdAt"`
	Type                       string                               `json:"type"`
	SequenceNumber             int                                  `json:"sequenceNumber"`
	ResourceVersion            int                                  `json:"resourceVersion"`
	Resource                   Reference                            `json:"resource"`
	TotalPrice                 *Money                               `json:"totalPrice"`
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	LineItemID                 string                               `json:"lineItemId"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
}

type OrderLineItemDiscountSetMessagePayload struct {
	TotalPrice                 *Money                               `json:"totalPrice"`
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	LineItemID                 string                               `json:"lineItemId"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
}

func (obj OrderLineItemDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemDiscountSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

type OrderPaymentChangedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	PaymentState    string    `json:"paymentState"`
}

type OrderPaymentChangedMessagePayload struct {
	PaymentState string `json:"paymentState"`
}

func (obj OrderPaymentChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderPaymentChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderPaymentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderReturnInfoAddedMessage struct {
	Version         int         `json:"version"`
	LastModifiedAt  time.Time   `json:"lastModifiedAt"`
	ID              string      `json:"id"`
	CreatedAt       time.Time   `json:"createdAt"`
	Type            string      `json:"type"`
	SequenceNumber  int         `json:"sequenceNumber"`
	ResourceVersion int         `json:"resourceVersion"`
	Resource        Reference   `json:"resource"`
	ReturnInfo      *ReturnInfo `json:"returnInfo"`
}

type OrderReturnInfoAddedMessagePayload struct {
	ReturnInfo *ReturnInfo `json:"returnInfo"`
}

func (obj OrderReturnInfoAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnInfoAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReturnInfoAdded", Alias: (*Alias)(&obj)})
}

type OrderReturnShipmentStateChangedMessage struct {
	Version             int                 `json:"version"`
	LastModifiedAt      time.Time           `json:"lastModifiedAt"`
	ID                  string              `json:"id"`
	CreatedAt           time.Time           `json:"createdAt"`
	Type                string              `json:"type"`
	SequenceNumber      int                 `json:"sequenceNumber"`
	ResourceVersion     int                 `json:"resourceVersion"`
	Resource            Reference           `json:"resource"`
	ReturnShipmentState ReturnShipmentState `json:"returnShipmentState"`
	ReturnItemID        string              `json:"returnItemId"`
}

type OrderReturnShipmentStateChangedMessagePayload struct {
	ReturnShipmentState ReturnShipmentState `json:"returnShipmentState"`
	ReturnItemID        string              `json:"returnItemId"`
}

func (obj OrderReturnShipmentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnShipmentStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderReturnShipmentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderShipmentStateChangedMessage struct {
	Version         int           `json:"version"`
	LastModifiedAt  time.Time     `json:"lastModifiedAt"`
	ID              string        `json:"id"`
	CreatedAt       time.Time     `json:"createdAt"`
	Type            string        `json:"type"`
	SequenceNumber  int           `json:"sequenceNumber"`
	ResourceVersion int           `json:"resourceVersion"`
	Resource        Reference     `json:"resource"`
	ShipmentState   ShipmentState `json:"shipmentState"`
}

type OrderShipmentStateChangedMessagePayload struct {
	ShipmentState ShipmentState `json:"shipmentState"`
}

func (obj OrderShipmentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShipmentStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShipmentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderShippingAddressSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Address         *Address  `json:"address"`
}

type OrderShippingAddressSetMessagePayload struct {
	Address *Address `json:"address"`
}

func (obj OrderShippingAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingAddressSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingAddressSet", Alias: (*Alias)(&obj)})
}

type OrderShippingInfoSetMessage struct {
	Version         int           `json:"version"`
	LastModifiedAt  time.Time     `json:"lastModifiedAt"`
	ID              string        `json:"id"`
	CreatedAt       time.Time     `json:"createdAt"`
	Type            string        `json:"type"`
	SequenceNumber  int           `json:"sequenceNumber"`
	ResourceVersion int           `json:"resourceVersion"`
	Resource        Reference     `json:"resource"`
	ShippingInfo    *ShippingInfo `json:"shippingInfo,omitempty"`
	OldShippingInfo *ShippingInfo `json:"oldShippingInfo,omitempty"`
}

type OrderShippingInfoSetMessagePayload struct {
	ShippingInfo    *ShippingInfo `json:"shippingInfo,omitempty"`
	OldShippingInfo *ShippingInfo `json:"oldShippingInfo,omitempty"`
}

func (obj OrderShippingInfoSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingInfoSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingInfoSet", Alias: (*Alias)(&obj)})
}

type OrderShippingRateInputSetMessage struct {
	Version              int               `json:"version"`
	LastModifiedAt       time.Time         `json:"lastModifiedAt"`
	ID                   string            `json:"id"`
	CreatedAt            time.Time         `json:"createdAt"`
	Type                 string            `json:"type"`
	SequenceNumber       int               `json:"sequenceNumber"`
	ResourceVersion      int               `json:"resourceVersion"`
	Resource             Reference         `json:"resource"`
	ShippingRateInput    ShippingRateInput `json:"shippingRateInput,omitempty"`
	OldShippingRateInput ShippingRateInput `json:"oldShippingRateInput,omitempty"`
}

func (obj *OrderShippingRateInputSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderShippingRateInputSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.OldShippingRateInput != nil {
		obj.OldShippingRateInput = AbstractShippingRateInputDiscriminatorMapping(obj.OldShippingRateInput)
	}
	if obj.ShippingRateInput != nil {
		obj.ShippingRateInput = AbstractShippingRateInputDiscriminatorMapping(obj.ShippingRateInput)
	}

	return nil
}

type OrderShippingRateInputSetMessagePayload struct {
	ShippingRateInput    ShippingRateInput `json:"shippingRateInput,omitempty"`
	OldShippingRateInput ShippingRateInput `json:"oldShippingRateInput,omitempty"`
}

func (obj OrderShippingRateInputSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingRateInputSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingRateInputSet", Alias: (*Alias)(&obj)})
}
func (obj *OrderShippingRateInputSetMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias OrderShippingRateInputSetMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.OldShippingRateInput != nil {
		obj.OldShippingRateInput = AbstractShippingRateInputDiscriminatorMapping(obj.OldShippingRateInput)
	}
	if obj.ShippingRateInput != nil {
		obj.ShippingRateInput = AbstractShippingRateInputDiscriminatorMapping(obj.ShippingRateInput)
	}

	return nil
}

type OrderStateChangedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	OrderState      string    `json:"orderState"`
}

type OrderStateChangedMessagePayload struct {
	OrderState string `json:"orderState"`
}

func (obj OrderStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderStateChanged", Alias: (*Alias)(&obj)})
}

type OrderStateTransitionMessage struct {
	Version         int             `json:"version"`
	LastModifiedAt  time.Time       `json:"lastModifiedAt"`
	ID              string          `json:"id"`
	CreatedAt       time.Time       `json:"createdAt"`
	Type            string          `json:"type"`
	SequenceNumber  int             `json:"sequenceNumber"`
	ResourceVersion int             `json:"resourceVersion"`
	Resource        Reference       `json:"resource"`
	State           *StateReference `json:"state"`
	Force           bool            `json:"force"`
}

type OrderStateTransitionMessagePayload struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force"`
}

func (obj OrderStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderStateTransition", Alias: (*Alias)(&obj)})
}

type ParcelAddedToDeliveryMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Parcel          *Parcel   `json:"parcel"`
	Delivery        *Delivery `json:"delivery"`
}

type ParcelAddedToDeliveryMessagePayload struct {
	Parcel   *Parcel   `json:"parcel"`
	Delivery *Delivery `json:"delivery"`
}

func (obj ParcelAddedToDeliveryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelAddedToDeliveryMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelAddedToDelivery", Alias: (*Alias)(&obj)})
}

type ParcelItemsUpdatedMessage struct {
	Version         int            `json:"version"`
	LastModifiedAt  time.Time      `json:"lastModifiedAt"`
	ID              string         `json:"id"`
	CreatedAt       time.Time      `json:"createdAt"`
	Type            string         `json:"type"`
	SequenceNumber  int            `json:"sequenceNumber"`
	ResourceVersion int            `json:"resourceVersion"`
	Resource        Reference      `json:"resource"`
	ParcelID        string         `json:"parcelId"`
	Items           []DeliveryItem `json:"items"`
	DeliveryID      string         `json:"deliveryId,omitempty"`
}

type ParcelItemsUpdatedMessagePayload struct {
	ParcelID   string         `json:"parcelId"`
	Items      []DeliveryItem `json:"items"`
	DeliveryID string         `json:"deliveryId,omitempty"`
}

func (obj ParcelItemsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelItemsUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelItemsUpdated", Alias: (*Alias)(&obj)})
}

type ParcelMeasurementsUpdatedMessage struct {
	Version         int                 `json:"version"`
	LastModifiedAt  time.Time           `json:"lastModifiedAt"`
	ID              string              `json:"id"`
	CreatedAt       time.Time           `json:"createdAt"`
	Type            string              `json:"type"`
	SequenceNumber  int                 `json:"sequenceNumber"`
	ResourceVersion int                 `json:"resourceVersion"`
	Resource        Reference           `json:"resource"`
	ParcelID        string              `json:"parcelId"`
	Measurements    *ParcelMeasurements `json:"measurements,omitempty"`
	DeliveryID      string              `json:"deliveryId"`
}

type ParcelMeasurementsUpdatedMessagePayload struct {
	ParcelID     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	DeliveryID   string              `json:"deliveryId"`
}

func (obj ParcelMeasurementsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelMeasurementsUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelMeasurementsUpdated", Alias: (*Alias)(&obj)})
}

type ParcelRemovedFromDeliveryMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Parcel          *Parcel   `json:"parcel"`
	DeliveryID      string    `json:"deliveryId"`
}

type ParcelRemovedFromDeliveryMessagePayload struct {
	Parcel     *Parcel `json:"parcel"`
	DeliveryID string  `json:"deliveryId"`
}

func (obj ParcelRemovedFromDeliveryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelRemovedFromDeliveryMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelRemovedFromDelivery", Alias: (*Alias)(&obj)})
}

type ParcelTrackingDataUpdatedMessage struct {
	Version         int           `json:"version"`
	LastModifiedAt  time.Time     `json:"lastModifiedAt"`
	ID              string        `json:"id"`
	CreatedAt       time.Time     `json:"createdAt"`
	Type            string        `json:"type"`
	SequenceNumber  int           `json:"sequenceNumber"`
	ResourceVersion int           `json:"resourceVersion"`
	Resource        Reference     `json:"resource"`
	TrackingData    *TrackingData `json:"trackingData,omitempty"`
	ParcelID        string        `json:"parcelId"`
	DeliveryID      string        `json:"deliveryId"`
}

type ParcelTrackingDataUpdatedMessagePayload struct {
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	ParcelID     string        `json:"parcelId"`
	DeliveryID   string        `json:"deliveryId"`
}

func (obj ParcelTrackingDataUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelTrackingDataUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelTrackingDataUpdated", Alias: (*Alias)(&obj)})
}

type PaymentCreatedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Payment         *Payment  `json:"payment"`
}

type PaymentCreatedMessagePayload struct {
	Payment *Payment `json:"payment"`
}

func (obj PaymentCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentCreated", Alias: (*Alias)(&obj)})
}

type PaymentInteractionAddedMessage struct {
	Version         int           `json:"version"`
	LastModifiedAt  time.Time     `json:"lastModifiedAt"`
	ID              string        `json:"id"`
	CreatedAt       time.Time     `json:"createdAt"`
	Type            string        `json:"type"`
	SequenceNumber  int           `json:"sequenceNumber"`
	ResourceVersion int           `json:"resourceVersion"`
	Resource        Reference     `json:"resource"`
	Interaction     *CustomFields `json:"interaction"`
}

type PaymentInteractionAddedMessagePayload struct {
	Interaction *CustomFields `json:"interaction"`
}

func (obj PaymentInteractionAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentInteractionAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentInteractionAdded", Alias: (*Alias)(&obj)})
}

type PaymentStatusInterfaceCodeSetMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	PaymentID       string    `json:"paymentId"`
	InterfaceCode   string    `json:"interfaceCode"`
}

type PaymentStatusInterfaceCodeSetMessagePayload struct {
	PaymentID     string `json:"paymentId"`
	InterfaceCode string `json:"interfaceCode"`
}

func (obj PaymentStatusInterfaceCodeSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusInterfaceCodeSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentStatusInterfaceCodeSet", Alias: (*Alias)(&obj)})
}

type PaymentStatusStateTransitionMessage struct {
	Version         int             `json:"version"`
	LastModifiedAt  time.Time       `json:"lastModifiedAt"`
	ID              string          `json:"id"`
	CreatedAt       time.Time       `json:"createdAt"`
	Type            string          `json:"type"`
	SequenceNumber  int             `json:"sequenceNumber"`
	ResourceVersion int             `json:"resourceVersion"`
	Resource        Reference       `json:"resource"`
	State           *StateReference `json:"state"`
	Force           bool            `json:"force"`
}

type PaymentStatusStateTransitionMessagePayload struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force"`
}

func (obj PaymentStatusStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentStatusStateTransition", Alias: (*Alias)(&obj)})
}

type PaymentTransactionAddedMessage struct {
	Version         int          `json:"version"`
	LastModifiedAt  time.Time    `json:"lastModifiedAt"`
	ID              string       `json:"id"`
	CreatedAt       time.Time    `json:"createdAt"`
	Type            string       `json:"type"`
	SequenceNumber  int          `json:"sequenceNumber"`
	ResourceVersion int          `json:"resourceVersion"`
	Resource        Reference    `json:"resource"`
	Transaction     *Transaction `json:"transaction"`
}

type PaymentTransactionAddedMessagePayload struct {
	Transaction *Transaction `json:"transaction"`
}

func (obj PaymentTransactionAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentTransactionAdded", Alias: (*Alias)(&obj)})
}

type PaymentTransactionStateChangedMessage struct {
	Version         int              `json:"version"`
	LastModifiedAt  time.Time        `json:"lastModifiedAt"`
	ID              string           `json:"id"`
	CreatedAt       time.Time        `json:"createdAt"`
	Type            string           `json:"type"`
	SequenceNumber  int              `json:"sequenceNumber"`
	ResourceVersion int              `json:"resourceVersion"`
	Resource        Reference        `json:"resource"`
	TransactionID   string           `json:"transactionId"`
	State           TransactionState `json:"state"`
}

type PaymentTransactionStateChangedMessagePayload struct {
	TransactionID string           `json:"transactionId"`
	State         TransactionState `json:"state"`
}

func (obj PaymentTransactionStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentTransactionStateChanged", Alias: (*Alias)(&obj)})
}

type ProductCreatedMessage struct {
	Version           int                `json:"version"`
	LastModifiedAt    time.Time          `json:"lastModifiedAt"`
	ID                string             `json:"id"`
	CreatedAt         time.Time          `json:"createdAt"`
	Type              string             `json:"type"`
	SequenceNumber    int                `json:"sequenceNumber"`
	ResourceVersion   int                `json:"resourceVersion"`
	Resource          Reference          `json:"resource"`
	ProductProjection *ProductProjection `json:"productProjection"`
}

type ProductCreatedMessagePayload struct {
	ProductProjection *ProductProjection `json:"productProjection"`
}

func (obj ProductCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductCreated", Alias: (*Alias)(&obj)})
}

type ProductDeletedMessage struct {
	Version           int                `json:"version"`
	LastModifiedAt    time.Time          `json:"lastModifiedAt"`
	ID                string             `json:"id"`
	CreatedAt         time.Time          `json:"createdAt"`
	Type              string             `json:"type"`
	SequenceNumber    int                `json:"sequenceNumber"`
	ResourceVersion   int                `json:"resourceVersion"`
	Resource          Reference          `json:"resource"`
	RemovedImageUrls  []interface{}      `json:"removedImageUrls"`
	CurrentProjection *ProductProjection `json:"currentProjection"`
}

type ProductDeletedMessagePayload struct {
	RemovedImageUrls  []interface{}      `json:"removedImageUrls"`
	CurrentProjection *ProductProjection `json:"currentProjection"`
}

func (obj ProductDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductDeleted", Alias: (*Alias)(&obj)})
}

type ProductImageAddedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	VariantID       int       `json:"variantId"`
	Staged          bool      `json:"staged"`
	Image           *Image    `json:"image"`
}

type ProductImageAddedMessagePayload struct {
	VariantID int    `json:"variantId"`
	Staged    bool   `json:"staged"`
	Image     *Image `json:"image"`
}

func (obj ProductImageAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductImageAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductImageAdded", Alias: (*Alias)(&obj)})
}

type ProductPublishedMessage struct {
	Version           int                 `json:"version"`
	LastModifiedAt    time.Time           `json:"lastModifiedAt"`
	ID                string              `json:"id"`
	CreatedAt         time.Time           `json:"createdAt"`
	Type              string              `json:"type"`
	SequenceNumber    int                 `json:"sequenceNumber"`
	ResourceVersion   int                 `json:"resourceVersion"`
	Resource          Reference           `json:"resource"`
	Scope             ProductPublishScope `json:"scope"`
	RemovedImageUrls  []interface{}       `json:"removedImageUrls"`
	ProductProjection *ProductProjection  `json:"productProjection"`
}

type ProductPublishedMessagePayload struct {
	Scope             ProductPublishScope `json:"scope"`
	RemovedImageUrls  []interface{}       `json:"removedImageUrls"`
	ProductProjection *ProductProjection  `json:"productProjection"`
}

func (obj ProductPublishedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductPublishedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductPublished", Alias: (*Alias)(&obj)})
}

type ProductRevertedStagedChangesMessage struct {
	Version          int           `json:"version"`
	LastModifiedAt   time.Time     `json:"lastModifiedAt"`
	ID               string        `json:"id"`
	CreatedAt        time.Time     `json:"createdAt"`
	Type             string        `json:"type"`
	SequenceNumber   int           `json:"sequenceNumber"`
	ResourceVersion  int           `json:"resourceVersion"`
	Resource         Reference     `json:"resource"`
	RemovedImageUrls []interface{} `json:"removedImageUrls"`
}

type ProductRevertedStagedChangesMessagePayload struct {
	RemovedImageUrls []interface{} `json:"removedImageUrls"`
}

func (obj ProductRevertedStagedChangesMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertedStagedChangesMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductRevertedStagedChanges", Alias: (*Alias)(&obj)})
}

type ProductSlugChangedMessage struct {
	Version         int              `json:"version"`
	LastModifiedAt  time.Time        `json:"lastModifiedAt"`
	ID              string           `json:"id"`
	CreatedAt       time.Time        `json:"createdAt"`
	Type            string           `json:"type"`
	SequenceNumber  int              `json:"sequenceNumber"`
	ResourceVersion int              `json:"resourceVersion"`
	Resource        Reference        `json:"resource"`
	Slug            *LocalizedString `json:"slug"`
}

type ProductSlugChangedMessagePayload struct {
	Slug *LocalizedString `json:"slug"`
}

func (obj ProductSlugChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductSlugChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductSlugChanged", Alias: (*Alias)(&obj)})
}

type ProductStateTransitionMessage struct {
	Version         int             `json:"version"`
	LastModifiedAt  time.Time       `json:"lastModifiedAt"`
	ID              string          `json:"id"`
	CreatedAt       time.Time       `json:"createdAt"`
	Type            string          `json:"type"`
	SequenceNumber  int             `json:"sequenceNumber"`
	ResourceVersion int             `json:"resourceVersion"`
	Resource        Reference       `json:"resource"`
	State           *StateReference `json:"state"`
	Force           bool            `json:"force"`
}

type ProductStateTransitionMessagePayload struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force"`
}

func (obj ProductStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductStateTransition", Alias: (*Alias)(&obj)})
}

type ProductUnpublishedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
}

type ProductUnpublishedMessagePayload struct{}

func (obj ProductUnpublishedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductUnpublishedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductUnpublished", Alias: (*Alias)(&obj)})
}

type ProductVariantDeletedMessage struct {
	Version          int             `json:"version"`
	LastModifiedAt   time.Time       `json:"lastModifiedAt"`
	ID               string          `json:"id"`
	CreatedAt        time.Time       `json:"createdAt"`
	Type             string          `json:"type"`
	SequenceNumber   int             `json:"sequenceNumber"`
	ResourceVersion  int             `json:"resourceVersion"`
	Resource         Reference       `json:"resource"`
	Variant          *ProductVariant `json:"variant"`
	RemovedImageUrls []interface{}   `json:"removedImageUrls"`
}

type ProductVariantDeletedMessagePayload struct {
	Variant          *ProductVariant `json:"variant"`
	RemovedImageUrls []interface{}   `json:"removedImageUrls"`
}

func (obj ProductVariantDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductVariantDeleted", Alias: (*Alias)(&obj)})
}

type ReviewCreatedMessage struct {
	Version         int       `json:"version"`
	LastModifiedAt  time.Time `json:"lastModifiedAt"`
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	Type            string    `json:"type"`
	SequenceNumber  int       `json:"sequenceNumber"`
	ResourceVersion int       `json:"resourceVersion"`
	Resource        Reference `json:"resource"`
	Review          *Review   `json:"review"`
}

type ReviewCreatedMessagePayload struct {
	Review *Review `json:"review"`
}

func (obj ReviewCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewCreated", Alias: (*Alias)(&obj)})
}

type ReviewRatingSetMessage struct {
	Version              int       `json:"version"`
	LastModifiedAt       time.Time `json:"lastModifiedAt"`
	ID                   string    `json:"id"`
	CreatedAt            time.Time `json:"createdAt"`
	Type                 string    `json:"type"`
	SequenceNumber       int       `json:"sequenceNumber"`
	ResourceVersion      int       `json:"resourceVersion"`
	Resource             Reference `json:"resource"`
	Target               Reference `json:"target"`
	OldRating            float64   `json:"oldRating"`
	NewRating            float64   `json:"newRating"`
	IncludedInStatistics bool      `json:"includedInStatistics"`
}

func (obj *ReviewRatingSetMessage) UnmarshalJSON(data []byte) error {
	type Alias ReviewRatingSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = AbstractReferenceDiscriminatorMapping(obj.Target)
	}

	return nil
}

type ReviewRatingSetMessagePayload struct {
	Target               Reference `json:"target"`
	OldRating            float64   `json:"oldRating"`
	NewRating            float64   `json:"newRating"`
	IncludedInStatistics bool      `json:"includedInStatistics"`
}

func (obj ReviewRatingSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewRatingSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewRatingSet", Alias: (*Alias)(&obj)})
}
func (obj *ReviewRatingSetMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias ReviewRatingSetMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = AbstractReferenceDiscriminatorMapping(obj.Target)
	}

	return nil
}

type ReviewStateTransitionMessage struct {
	Version                 int             `json:"version"`
	LastModifiedAt          time.Time       `json:"lastModifiedAt"`
	ID                      string          `json:"id"`
	CreatedAt               time.Time       `json:"createdAt"`
	Type                    string          `json:"type"`
	SequenceNumber          int             `json:"sequenceNumber"`
	ResourceVersion         int             `json:"resourceVersion"`
	Resource                Reference       `json:"resource"`
	Target                  Reference       `json:"target"`
	OldState                *StateReference `json:"oldState"`
	OldIncludedInStatistics bool            `json:"oldIncludedInStatistics"`
	NewState                *StateReference `json:"newState"`
	NewIncludedInStatistics bool            `json:"newIncludedInStatistics"`
	Force                   bool            `json:"force"`
}

func (obj *ReviewStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias ReviewStateTransitionMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = AbstractReferenceDiscriminatorMapping(obj.Target)
	}

	return nil
}

type ReviewStateTransitionMessagePayload struct {
	Target                  Reference       `json:"target"`
	OldState                *StateReference `json:"oldState"`
	OldIncludedInStatistics bool            `json:"oldIncludedInStatistics"`
	NewState                *StateReference `json:"newState"`
	NewIncludedInStatistics bool            `json:"newIncludedInStatistics"`
	Force                   bool            `json:"force"`
}

func (obj ReviewStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewStateTransition", Alias: (*Alias)(&obj)})
}
func (obj *ReviewStateTransitionMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias ReviewStateTransitionMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = AbstractReferenceDiscriminatorMapping(obj.Target)
	}

	return nil
}
