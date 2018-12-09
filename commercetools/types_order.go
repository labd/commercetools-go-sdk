// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type CustomLineItemReturnItem struct {
	Type             string              `json:"type"`
	ShipmentState    ReturnShipmentState `json:"shipmentState"`
	Quantity         int                 `json:"quantity"`
	PaymentState     ReturnPaymentState  `json:"paymentState"`
	LastModifiedAt   time.Time           `json:"lastModifiedAt"`
	ID               string              `json:"id"`
	CreatedAt        time.Time           `json:"createdAt"`
	Comment          string              `json:"comment,omitempty"`
	CustomlineItemID string              `json:"customlineItemId"`
}

type Delivery struct {
	Parcels   []Parcel       `json:"parcels"`
	Items     []DeliveryItem `json:"items"`
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	Address   *Address       `json:"address,omitempty"`
}

type DeliveryItem struct {
	Quantity float64 `json:"quantity"`
	ID       string  `json:"id"`
}

type DiscountedLineItemPriceDraft struct {
	Value             *Money                      `json:"value"`
	IncludedDiscounts []DiscountedLineItemPortion `json:"includedDiscounts"`
}

type ItemState struct {
	State    *StateReference `json:"state"`
	Quantity float64         `json:"quantity"`
}

type LineItemImportDraft struct {
	Variant         *ProductVariantImportDraft `json:"variant"`
	TaxRate         *TaxRate                   `json:"taxRate,omitempty"`
	SupplyChannel   *ChannelReference          `json:"supplyChannel,omitempty"`
	State           []ItemState                `json:"state,omitempty"`
	ShippingDetails *ItemShippingDetailsDraft  `json:"shippingDetails,omitempty"`
	Quantity        float64                    `json:"quantity"`
	ProductID       string                     `json:"productId,omitempty"`
	Price           *PriceDraft                `json:"price"`
	Name            *LocalizedString           `json:"name"`
	Custom          *CustomFieldsDraft         `json:"custom,omitempty"`
}

type LineItemReturnItem struct {
	Type           string              `json:"type"`
	ShipmentState  ReturnShipmentState `json:"shipmentState"`
	Quantity       int                 `json:"quantity"`
	PaymentState   ReturnPaymentState  `json:"paymentState"`
	LastModifiedAt time.Time           `json:"lastModifiedAt"`
	ID             string              `json:"id"`
	CreatedAt      time.Time           `json:"createdAt"`
	Comment        string              `json:"comment,omitempty"`
	LineItemID     string              `json:"lineItemId"`
}

type Order struct {
	Version                   int                     `json:"version"`
	LastModifiedAt            time.Time               `json:"lastModifiedAt"`
	ID                        string                  `json:"id"`
	CreatedAt                 time.Time               `json:"createdAt"`
	TotalPrice                *Money                  `json:"totalPrice"`
	TaxedPrice                *TaxedPrice             `json:"taxedPrice,omitempty"`
	TaxRoundingMode           RoundingMode            `json:"taxRoundingMode,omitempty"`
	TaxMode                   TaxMode                 `json:"taxMode,omitempty"`
	TaxCalculationMode        TaxCalculationMode      `json:"taxCalculationMode,omitempty"`
	SyncInfo                  []SyncInfo              `json:"syncInfo"`
	State                     *StateReference         `json:"state,omitempty"`
	ShippingRateInput         ShippingRateInput       `json:"shippingRateInput,omitempty"`
	ShippingInfo              *ShippingInfo           `json:"shippingInfo,omitempty"`
	ShippingAddress           *Address                `json:"shippingAddress,omitempty"`
	ShipmentState             ShipmentState           `json:"shipmentState,omitempty"`
	ReturnInfo                []ReturnInfo            `json:"returnInfo,omitempty"`
	PaymentState              PaymentState            `json:"paymentState,omitempty"`
	PaymentInfo               *PaymentInfo            `json:"paymentInfo,omitempty"`
	Origin                    CartOrigin              `json:"origin"`
	OrderState                OrderState              `json:"orderState"`
	OrderNumber               string                  `json:"orderNumber,omitempty"`
	Locale                    string                  `json:"locale,omitempty"`
	LineItems                 []LineItem              `json:"lineItems"`
	LastMessageSequenceNumber int                     `json:"lastMessageSequenceNumber"`
	ItemShippingAddresses     []Address               `json:"itemShippingAddresses,omitempty"`
	InventoryMode             InventoryMode           `json:"inventoryMode,omitempty"`
	DiscountCodes             []DiscountCodeInfo      `json:"discountCodes,omitempty"`
	CustomerID                string                  `json:"customerId,omitempty"`
	CustomerGroup             *CustomerGroupReference `json:"customerGroup,omitempty"`
	CustomerEmail             string                  `json:"customerEmail,omitempty"`
	CustomLineItems           []CustomLineItem        `json:"customLineItems"`
	Custom                    *CustomFields           `json:"custom,omitempty"`
	Country                   string                  `json:"country,omitempty"`
	CompletedAt               time.Time               `json:"completedAt,omitempty"`
	Cart                      *CartReference          `json:"cart,omitempty"`
	BillingAddress            *Address                `json:"billingAddress,omitempty"`
	AnonymousID               string                  `json:"anonymousId,omitempty"`
}

func (obj *Order) UnmarshalJSON(data []byte) error {
	type Alias Order
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		obj.ShippingRateInput = AbstractShippingRateInputDiscriminatorMapping(obj.ShippingRateInput)
	}

	return nil
}

type OrderAddDeliveryAction struct {
	Parcels []ParcelDraft  `json:"parcels,omitempty"`
	Items   []DeliveryItem `json:"items,omitempty"`
	Address *Address       `json:"address,omitempty"`
}

func (obj OrderAddDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDelivery", Alias: (*Alias)(&obj)})
}

type OrderAddItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

func (obj OrderAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addItemShippingAddress", Alias: (*Alias)(&obj)})
}

type OrderAddParcelToDeliveryAction struct {
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	Items        []DeliveryItem      `json:"items,omitempty"`
	DeliveryID   string              `json:"deliveryId"`
}

func (obj OrderAddParcelToDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddParcelToDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addParcelToDelivery", Alias: (*Alias)(&obj)})
}

type OrderAddPaymentAction struct {
	Payment *PaymentReference `json:"payment"`
}

func (obj OrderAddPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPayment", Alias: (*Alias)(&obj)})
}

type OrderAddReturnInfoAction struct {
	ReturnTrackingID string            `json:"returnTrackingId,omitempty"`
	ReturnDate       time.Time         `json:"returnDate,omitempty"`
	Items            []ReturnItemDraft `json:"items"`
}

func (obj OrderAddReturnInfoAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddReturnInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addReturnInfo", Alias: (*Alias)(&obj)})
}

type OrderChangeOrderStateAction struct {
	OrderState OrderState `json:"orderState"`
}

func (obj OrderChangeOrderStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderChangeOrderStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeOrderState", Alias: (*Alias)(&obj)})
}

type OrderChangePaymentStateAction struct {
	PaymentState PaymentState `json:"paymentState,omitempty"`
}

func (obj OrderChangePaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderChangePaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePaymentState", Alias: (*Alias)(&obj)})
}

type OrderChangeShipmentStateAction struct {
	ShipmentState ShipmentState `json:"shipmentState,omitempty"`
}

func (obj OrderChangeShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderChangeShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeShipmentState", Alias: (*Alias)(&obj)})
}

type OrderFromCartDraft struct {
	Version      int          `json:"version"`
	PaymentState PaymentState `json:"paymentState,omitempty"`
	OrderNumber  string       `json:"orderNumber,omitempty"`
	ID           string       `json:"id"`
}

type OrderImportCustomLineItemStateAction struct {
	State            []ItemState `json:"state"`
	CustomLineItemID string      `json:"customLineItemId"`
}

func (obj OrderImportCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderImportCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importCustomLineItemState", Alias: (*Alias)(&obj)})
}

type OrderImportDraft struct {
	TotalPrice            *Money                  `json:"totalPrice"`
	TaxedPrice            *TaxedPrice             `json:"taxedPrice,omitempty"`
	TaxRoundingMode       RoundingMode            `json:"taxRoundingMode,omitempty"`
	ShippingInfo          *ShippingInfoDraft      `json:"shippingInfo,omitempty"`
	ShippingAddress       *Address                `json:"shippingAddress,omitempty"`
	ShipmentState         ShipmentState           `json:"shipmentState,omitempty"`
	PaymentState          PaymentState            `json:"paymentState,omitempty"`
	OrderState            OrderState              `json:"orderState,omitempty"`
	OrderNumber           string                  `json:"orderNumber,omitempty"`
	LineItems             []LineItemImportDraft   `json:"lineItems,omitempty"`
	ItemShippingAddresses []Address               `json:"itemShippingAddresses,omitempty"`
	InventoryMode         InventoryMode           `json:"inventoryMode,omitempty"`
	CustomerID            string                  `json:"customerId,omitempty"`
	CustomerGroup         *CustomerGroupReference `json:"customerGroup,omitempty"`
	CustomerEmail         string                  `json:"customerEmail,omitempty"`
	CustomLineItems       []CustomLineItemDraft   `json:"customLineItems,omitempty"`
	Custom                *CustomFieldsDraft      `json:"custom,omitempty"`
	Country               string                  `json:"country,omitempty"`
	CompletedAt           time.Time               `json:"completedAt,omitempty"`
	BillingAddress        *Address                `json:"billingAddress,omitempty"`
}

type OrderImportLineItemStateAction struct {
	State      []ItemState `json:"state"`
	LineItemID string      `json:"lineItemId"`
}

func (obj OrderImportLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderImportLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importLineItemState", Alias: (*Alias)(&obj)})
}

type OrderPagedQueryResponse struct {
	Total   int     `json:"total,omitempty"`
	Offset  int     `json:"offset"`
	Count   int     `json:"count"`
	Results []Order `json:"results"`
}

type OrderReference struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
	Obj *Order `json:"obj,omitempty"`
}

func (obj OrderReference) MarshalJSON() ([]byte, error) {
	type Alias OrderReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "order", Alias: (*Alias)(&obj)})
}

type OrderRemoveDeliveryAction struct {
	DeliveryID string `json:"deliveryId"`
}

func (obj OrderRemoveDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemoveDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDelivery", Alias: (*Alias)(&obj)})
}

type OrderRemoveItemShippingAddressAction struct {
	AddressKey string `json:"addressKey"`
}

func (obj OrderRemoveItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemoveItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeItemShippingAddress", Alias: (*Alias)(&obj)})
}

type OrderRemoveParcelFromDeliveryAction struct {
	ParcelID string `json:"parcelId"`
}

func (obj OrderRemoveParcelFromDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemoveParcelFromDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeParcelFromDelivery", Alias: (*Alias)(&obj)})
}

type OrderRemovePaymentAction struct {
	Payment *PaymentReference `json:"payment"`
}

func (obj OrderRemovePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemovePaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePayment", Alias: (*Alias)(&obj)})
}

type OrderSetBillingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

func (obj OrderSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

type OrderSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj OrderSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetCustomLineItemCustomFieldAction struct {
	Value            interface{} `json:"value,omitempty"`
	Name             string      `json:"name"`
	CustomLineItemID string      `json:"customLineItemId"`
}

func (obj OrderSetCustomLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetCustomLineItemCustomTypeAction struct {
	Type             *TypeReference  `json:"type,omitempty"`
	Fields           *FieldContainer `json:"fields,omitempty"`
	CustomLineItemID string          `json:"customLineItemId"`
}

func (obj OrderSetCustomLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetCustomLineItemShippingDetailsAction struct {
	ShippingDetails  *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	CustomLineItemID string                    `json:"customLineItemId"`
}

func (obj OrderSetCustomLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type OrderSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

func (obj OrderSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetCustomerEmailAction struct {
	Email string `json:"email,omitempty"`
}

func (obj OrderSetCustomerEmailAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomerEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerEmail", Alias: (*Alias)(&obj)})
}

type OrderSetCustomerIdAction struct {
	CustomerID string `json:"customerId,omitempty"`
}

func (obj OrderSetCustomerIdAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomerIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerId", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryAddressAction struct {
	DeliveryID string   `json:"deliveryId"`
	Address    *Address `json:"address,omitempty"`
}

func (obj OrderSetDeliveryAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddress", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryItemsAction struct {
	Items      []DeliveryItem `json:"items"`
	DeliveryID string         `json:"deliveryId"`
}

func (obj OrderSetDeliveryItemsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryItems", Alias: (*Alias)(&obj)})
}

type OrderSetLineItemCustomFieldAction struct {
	Value      interface{} `json:"value,omitempty"`
	Name       string      `json:"name"`
	LineItemID string      `json:"lineItemId"`
}

func (obj OrderSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetLineItemCustomTypeAction struct {
	Type       *TypeReference  `json:"type,omitempty"`
	LineItemID string          `json:"lineItemId"`
	Fields     *FieldContainer `json:"fields,omitempty"`
}

func (obj OrderSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetLineItemShippingDetailsAction struct {
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	LineItemID      string                    `json:"lineItemId"`
}

func (obj OrderSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type OrderSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

func (obj OrderSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type OrderSetOrderNumberAction struct {
	OrderNumber string `json:"orderNumber,omitempty"`
}

func (obj OrderSetOrderNumberAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetOrderNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setOrderNumber", Alias: (*Alias)(&obj)})
}

type OrderSetParcelItemsAction struct {
	ParcelID string         `json:"parcelId"`
	Items    []DeliveryItem `json:"items"`
}

func (obj OrderSetParcelItemsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelItems", Alias: (*Alias)(&obj)})
}

type OrderSetParcelMeasurementsAction struct {
	ParcelID     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
}

func (obj OrderSetParcelMeasurementsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelMeasurementsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelMeasurements", Alias: (*Alias)(&obj)})
}

type OrderSetParcelTrackingDataAction struct {
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	ParcelID     string        `json:"parcelId"`
}

func (obj OrderSetParcelTrackingDataAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelTrackingDataAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelTrackingData", Alias: (*Alias)(&obj)})
}

type OrderSetReturnPaymentStateAction struct {
	ReturnItemID string             `json:"returnItemId"`
	PaymentState ReturnPaymentState `json:"paymentState"`
}

func (obj OrderSetReturnPaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetReturnPaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnPaymentState", Alias: (*Alias)(&obj)})
}

type OrderSetReturnShipmentStateAction struct {
	ShipmentState ReturnShipmentState `json:"shipmentState"`
	ReturnItemID  string              `json:"returnItemId"`
}

func (obj OrderSetReturnShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetReturnShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnShipmentState", Alias: (*Alias)(&obj)})
}

type OrderSetShippingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

func (obj OrderSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

type OrderState string

const (
	OrderStateOpen      OrderState = "Open"
	OrderStateConfirmed OrderState = "Confirmed"
	OrderStateComplete  OrderState = "Complete"
	OrderStateCancelled OrderState = "Cancelled"
)

type OrderTransitionCustomLineItemStateAction struct {
	ToState              *StateReference `json:"toState"`
	Quantity             int             `json:"quantity"`
	FromState            *StateReference `json:"fromState"`
	CustomLineItemID     string          `json:"customLineItemId"`
	ActualTransitionDate time.Time       `json:"actualTransitionDate,omitempty"`
}

func (obj OrderTransitionCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderTransitionCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionCustomLineItemState", Alias: (*Alias)(&obj)})
}

type OrderTransitionLineItemStateAction struct {
	ToState              *StateReference `json:"toState"`
	Quantity             int             `json:"quantity"`
	LineItemID           string          `json:"lineItemId"`
	FromState            *StateReference `json:"fromState"`
	ActualTransitionDate time.Time       `json:"actualTransitionDate,omitempty"`
}

func (obj OrderTransitionLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderTransitionLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionLineItemState", Alias: (*Alias)(&obj)})
}

type OrderTransitionStateAction struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force,omitempty"`
}

func (obj OrderTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

type OrderUpdate struct {
	Version int                 `json:"version"`
	Actions []OrderUpdateAction `json:"actions"`
}

func (obj *OrderUpdate) UnmarshalJSON(data []byte) error {
	type Alias OrderUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractOrderUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type OrderUpdateAction interface{}
type AbstractOrderUpdateAction struct{}

func AbstractOrderUpdateActionDiscriminatorMapping(input OrderUpdateAction) OrderUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addDelivery":
		new := OrderAddDeliveryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addItemShippingAddress":
		new := OrderAddItemShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addParcelToDelivery":
		new := OrderAddParcelToDeliveryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addPayment":
		new := OrderAddPaymentAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addReturnInfo":
		new := OrderAddReturnInfoAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeOrderState":
		new := OrderChangeOrderStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changePaymentState":
		new := OrderChangePaymentStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeShipmentState":
		new := OrderChangeShipmentStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "importCustomLineItemState":
		new := OrderImportCustomLineItemStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "importLineItemState":
		new := OrderImportLineItemStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeDelivery":
		new := OrderRemoveDeliveryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeItemShippingAddress":
		new := OrderRemoveItemShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeParcelFromDelivery":
		new := OrderRemoveParcelFromDeliveryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removePayment":
		new := OrderRemovePaymentAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setBillingAddress":
		new := OrderSetBillingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := OrderSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemCustomField":
		new := OrderSetCustomLineItemCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemCustomType":
		new := OrderSetCustomLineItemCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemShippingDetails":
		new := OrderSetCustomLineItemShippingDetailsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := OrderSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomerEmail":
		new := OrderSetCustomerEmailAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomerId":
		new := OrderSetCustomerIdAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDeliveryAddress":
		new := OrderSetDeliveryAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDeliveryItems":
		new := OrderSetDeliveryItemsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemCustomField":
		new := OrderSetLineItemCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemCustomType":
		new := OrderSetLineItemCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemShippingDetails":
		new := OrderSetLineItemShippingDetailsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLocale":
		new := OrderSetLocaleAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setOrderNumber":
		new := OrderSetOrderNumberAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setParcelItems":
		new := OrderSetParcelItemsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setParcelMeasurements":
		new := OrderSetParcelMeasurementsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setParcelTrackingData":
		new := OrderSetParcelTrackingDataAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setReturnPaymentState":
		new := OrderSetReturnPaymentStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setReturnShipmentState":
		new := OrderSetReturnShipmentStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingAddress":
		new := OrderSetShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "transitionCustomLineItemState":
		new := OrderTransitionCustomLineItemStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "transitionLineItemState":
		new := OrderTransitionLineItemStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "transitionState":
		new := OrderTransitionStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "updateItemShippingAddress":
		new := OrderUpdateItemShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "updateSyncInfo":
		new := OrderUpdateSyncInfoAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type OrderUpdateItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

func (obj OrderUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

type OrderUpdateSyncInfoAction struct {
	SyncedAt   time.Time         `json:"syncedAt,omitempty"`
	ExternalID string            `json:"externalId,omitempty"`
	Channel    *ChannelReference `json:"channel"`
}

func (obj OrderUpdateSyncInfoAction) MarshalJSON() ([]byte, error) {
	type Alias OrderUpdateSyncInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateSyncInfo", Alias: (*Alias)(&obj)})
}

type Parcel struct {
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	Items        []DeliveryItem      `json:"items,omitempty"`
	ID           string              `json:"id"`
	CreatedAt    time.Time           `json:"createdAt"`
}

type ParcelDraft struct {
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	Items        []DeliveryItem      `json:"items,omitempty"`
}

type ParcelMeasurements struct {
	WidthInMillimeter  float64 `json:"widthInMillimeter,omitempty"`
	WeightInGram       float64 `json:"weightInGram,omitempty"`
	LengthInMillimeter float64 `json:"lengthInMillimeter,omitempty"`
	HeightInMillimeter float64 `json:"heightInMillimeter,omitempty"`
}

type PaymentInfo struct {
	Payments []PaymentReference `json:"payments"`
}
type PaymentState string

const (
	PaymentStateBalanceDue PaymentState = "BalanceDue"
	PaymentStateFailed     PaymentState = "Failed"
	PaymentStatePending    PaymentState = "Pending"
	PaymentStateCreditOwed PaymentState = "CreditOwed"
	PaymentStatePaid       PaymentState = "Paid"
)

type ProductVariantImportDraft struct {
	Sku        string      `json:"sku,omitempty"`
	Prices     []Price     `json:"prices,omitempty"`
	Images     []Image     `json:"images,omitempty"`
	ID         int         `json:"id,omitempty"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

type ReturnInfo struct {
	ReturnTrackingID string       `json:"returnTrackingId,omitempty"`
	ReturnDate       time.Time    `json:"returnDate,omitempty"`
	Items            []ReturnItem `json:"items"`
}

type ReturnItem struct {
	Type           string              `json:"type"`
	ShipmentState  ReturnShipmentState `json:"shipmentState"`
	Quantity       int                 `json:"quantity"`
	PaymentState   ReturnPaymentState  `json:"paymentState"`
	LastModifiedAt time.Time           `json:"lastModifiedAt"`
	ID             string              `json:"id"`
	CreatedAt      time.Time           `json:"createdAt"`
	Comment        string              `json:"comment,omitempty"`
}

type ReturnItemDraft struct {
	ShipmentState    ReturnShipmentState `json:"shipmentState"`
	Quantity         int                 `json:"quantity"`
	LineItemID       string              `json:"lineItemId,omitempty"`
	CustomLineItemID string              `json:"customLineItemId,omitempty"`
	Comment          string              `json:"comment,omitempty"`
}
type ReturnPaymentState string

const (
	ReturnPaymentStateNonRefundable ReturnPaymentState = "NonRefundable"
	ReturnPaymentStateInitial       ReturnPaymentState = "Initial"
	ReturnPaymentStateRefunded      ReturnPaymentState = "Refunded"
	ReturnPaymentStateNotRefunded   ReturnPaymentState = "NotRefunded"
)

type ReturnShipmentState string

const (
	ReturnShipmentStateAdvised     ReturnShipmentState = "Advised"
	ReturnShipmentStateReturned    ReturnShipmentState = "Returned"
	ReturnShipmentStateBackInStock ReturnShipmentState = "BackInStock"
	ReturnShipmentStateUnusable    ReturnShipmentState = "Unusable"
)

type ShipmentState string

const (
	ShipmentStateShipped   ShipmentState = "Shipped"
	ShipmentStateReady     ShipmentState = "Ready"
	ShipmentStatePending   ShipmentState = "Pending"
	ShipmentStateDelayed   ShipmentState = "Delayed"
	ShipmentStatePartial   ShipmentState = "Partial"
	ShipmentStateBackorder ShipmentState = "Backorder"
)

type ShippingInfoDraft struct {
	TaxedPrice          *TaxedItemPriceDraft          `json:"taxedPrice,omitempty"`
	TaxRate             *TaxRate                      `json:"taxRate,omitempty"`
	TaxCategory         *TaxCategoryReference         `json:"taxCategory,omitempty"`
	ShippingRate        *ShippingRateDraft            `json:"shippingRate"`
	ShippingMethodState ShippingMethodState           `json:"shippingMethodState,omitempty"`
	ShippingMethodName  string                        `json:"shippingMethodName"`
	ShippingMethod      *ShippingMethodReference      `json:"shippingMethod,omitempty"`
	Price               *Money                        `json:"price"`
	DiscountedPrice     *DiscountedLineItemPriceDraft `json:"discountedPrice,omitempty"`
	Deliveries          []Delivery                    `json:"deliveries,omitempty"`
}

type StagedOrderUpdateAction interface{}
type AbstractStagedOrderUpdateAction struct{}

func AbstractStagedOrderUpdateActionDiscriminatorMapping(input StagedOrderUpdateAction) StagedOrderUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addCustomLineItem":
		new := StagedOrderAddCustomLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addDelivery":
		new := StagedOrderAddDeliveryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addDiscountCode":
		new := StagedOrderAddDiscountCodeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addItemShippingAddress":
		new := StagedOrderAddItemShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addLineItem":
		new := StagedOrderAddLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addParcelToDelivery":
		new := StagedOrderAddParcelToDeliveryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addPayment":
		new := StagedOrderAddPaymentAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addReturnInfo":
		new := StagedOrderAddReturnInfoAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addShoppingList":
		new := StagedOrderAddShoppingListAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeCustomLineItemMoney":
		new := StagedOrderChangeCustomLineItemMoneyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeCustomLineItemQuantity":
		new := StagedOrderChangeCustomLineItemQuantityAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeLineItemQuantity":
		new := StagedOrderChangeLineItemQuantityAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeOrderState":
		new := StagedOrderChangeOrderStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changePaymentState":
		new := StagedOrderChangePaymentStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeShipmentState":
		new := StagedOrderChangeShipmentStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTaxCalculationMode":
		new := StagedOrderChangeTaxCalculationModeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTaxMode":
		new := StagedOrderChangeTaxModeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTaxRoundingMode":
		new := StagedOrderChangeTaxRoundingModeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "importCustomLineItemState":
		new := StagedOrderImportCustomLineItemStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "importLineItemState":
		new := StagedOrderImportLineItemStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeCustomLineItem":
		new := StagedOrderRemoveCustomLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeDelivery":
		new := StagedOrderRemoveDeliveryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeDiscountCode":
		new := StagedOrderRemoveDiscountCodeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeItemShippingAddress":
		new := StagedOrderRemoveItemShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeLineItem":
		new := StagedOrderRemoveLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeParcelFromDelivery":
		new := StagedOrderRemoveParcelFromDeliveryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removePayment":
		new := StagedOrderRemovePaymentAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setBillingAddress":
		new := StagedOrderSetBillingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCountry":
		new := StagedOrderSetCountryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := StagedOrderSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemCustomField":
		new := StagedOrderSetCustomLineItemCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemCustomType":
		new := StagedOrderSetCustomLineItemCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemShippingDetails":
		new := StagedOrderSetCustomLineItemShippingDetailsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemTaxAmount":
		new := StagedOrderSetCustomLineItemTaxAmountAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemTaxRate":
		new := StagedOrderSetCustomLineItemTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomShippingMethod":
		new := StagedOrderSetCustomShippingMethodAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := StagedOrderSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomerEmail":
		new := StagedOrderSetCustomerEmailAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomerGroup":
		new := StagedOrderSetCustomerGroupAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomerId":
		new := StagedOrderSetCustomerIdAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDeliveryAddress":
		new := StagedOrderSetDeliveryAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDeliveryItems":
		new := StagedOrderSetDeliveryItemsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemCustomField":
		new := StagedOrderSetLineItemCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemCustomType":
		new := StagedOrderSetLineItemCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemPrice":
		new := StagedOrderSetLineItemPriceAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemShippingDetails":
		new := StagedOrderSetLineItemShippingDetailsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemTaxAmount":
		new := StagedOrderSetLineItemTaxAmountAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemTaxRate":
		new := StagedOrderSetLineItemTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemTotalPrice":
		new := StagedOrderSetLineItemTotalPriceAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLocale":
		new := StagedOrderSetLocaleAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setOrderNumber":
		new := StagedOrderSetOrderNumberAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setOrderTotalTax":
		new := StagedOrderSetOrderTotalTaxAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setParcelItems":
		new := StagedOrderSetParcelItemsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setParcelMeasurements":
		new := StagedOrderSetParcelMeasurementsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setParcelTrackingData":
		new := StagedOrderSetParcelTrackingDataAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setReturnPaymentState":
		new := StagedOrderSetReturnPaymentStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setReturnShipmentState":
		new := StagedOrderSetReturnShipmentStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingAddress":
		new := StagedOrderSetShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingAddressAndCustomShippingMethod":
		new := StagedOrderSetShippingAddressAndCustomShippingMethodAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingAddressAndShippingMethod":
		new := StagedOrderSetShippingAddressAndShippingMethodAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingMethod":
		new := StagedOrderSetShippingMethodAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingMethodTaxAmount":
		new := StagedOrderSetShippingMethodTaxAmountAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingMethodTaxRate":
		new := StagedOrderSetShippingMethodTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingRateInput":
		new := StagedOrderSetShippingRateInputAction{}
		mapstructure.Decode(input, &new)
		if new.ShippingRateInput != nil {
			new.ShippingRateInput = AbstractShippingRateInputDraftDiscriminatorMapping(new.ShippingRateInput)
		}

		return new
	case "transitionCustomLineItemState":
		new := StagedOrderTransitionCustomLineItemStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "transitionLineItemState":
		new := StagedOrderTransitionLineItemStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "transitionState":
		new := StagedOrderTransitionStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "updateItemShippingAddress":
		new := StagedOrderUpdateItemShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "updateSyncInfo":
		new := StagedOrderUpdateSyncInfoAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type SyncInfo struct {
	SyncedAt   time.Time         `json:"syncedAt"`
	ExternalID string            `json:"externalId,omitempty"`
	Channel    *ChannelReference `json:"channel"`
}

type TaxedItemPriceDraft struct {
	TotalNet   *Money `json:"totalNet"`
	TotalGross *Money `json:"totalGross"`
}

type TrackingData struct {
	TrackingID          string `json:"trackingId,omitempty"`
	ProviderTransaction string `json:"providerTransaction,omitempty"`
	Provider            string `json:"provider,omitempty"`
	IsReturn            bool   `json:"isReturn,omitempty"`
	Carrier             string `json:"carrier,omitempty"`
}
