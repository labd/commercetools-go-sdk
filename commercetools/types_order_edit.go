// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type OrderEdit struct {
	Version        int                       `json:"version"`
	ID             string                    `json:"id"`
	StagedActions  []StagedOrderUpdateAction `json:"stagedActions"`
	Result         OrderEditResult           `json:"result"`
	Resource       *OrderReference           `json:"resource"`
	LastModifiedAt time.Time                 `json:"lastModifiedAt"`
	Key            string                    `json:"key,omitempty"`
	Custom         *CustomFields             `json:"custom,omitempty"`
	CreatedAt      time.Time                 `json:"createdAt"`
	Comment        string                    `json:"comment,omitempty"`
}

func (obj *OrderEdit) UnmarshalJSON(data []byte) error {
	type Alias OrderEdit
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Result != nil {
		obj.Result = AbstractOrderEditResultDiscriminatorMapping(obj.Result)
	}
	for i := range obj.StagedActions {
		obj.StagedActions[i] = AbstractStagedOrderUpdateActionDiscriminatorMapping(obj.StagedActions[i])
	}

	return nil
}

type OrderEditAddStagedActionAction struct {
	StagedAction StagedOrderUpdateAction `json:"stagedAction"`
}

func (obj OrderEditAddStagedActionAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditAddStagedActionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addStagedAction", Alias: (*Alias)(&obj)})
}
func (obj *OrderEditAddStagedActionAction) UnmarshalJSON(data []byte) error {
	type Alias OrderEditAddStagedActionAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.StagedAction != nil {
		obj.StagedAction = AbstractStagedOrderUpdateActionDiscriminatorMapping(obj.StagedAction)
	}

	return nil
}

type OrderEditApplied struct {
	ExcerptBeforeEdit *OrderExcerpt `json:"excerptBeforeEdit"`
	ExcerptAfterEdit  *OrderExcerpt `json:"excerptAfterEdit"`
	AppliedAt         time.Time     `json:"appliedAt"`
}

func (obj OrderEditApplied) MarshalJSON() ([]byte, error) {
	type Alias OrderEditApplied
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "Applied", Alias: (*Alias)(&obj)})
}

type OrderEditApply struct {
	ResourceVersion int `json:"resourceVersion"`
	EditVersion     int `json:"editVersion"`
}

type OrderEditDraft struct {
	StagedActions []StagedOrderUpdateAction `json:"stagedActions,omitempty"`
	Resource      *OrderReference           `json:"resource"`
	Key           string                    `json:"key,omitempty"`
	DryRun        bool                      `json:"dryRun,omitempty"`
	Custom        *CustomFieldsDraft        `json:"custom,omitempty"`
	Comment       string                    `json:"comment,omitempty"`
}

func (obj *OrderEditDraft) UnmarshalJSON(data []byte) error {
	type Alias OrderEditDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.StagedActions {
		obj.StagedActions[i] = AbstractStagedOrderUpdateActionDiscriminatorMapping(obj.StagedActions[i])
	}

	return nil
}

type OrderEditNotProcessed struct{}

func (obj OrderEditNotProcessed) MarshalJSON() ([]byte, error) {
	type Alias OrderEditNotProcessed
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "NotProcessed", Alias: (*Alias)(&obj)})
}

type OrderEditPagedQueryResponse struct {
	Total   int         `json:"total,omitempty"`
	Offset  int         `json:"offset"`
	Count   int         `json:"count"`
	Results []OrderEdit `json:"results"`
}

type OrderEditPreviewFailure struct {
	Errors []ErrorObject `json:"errors"`
}

func (obj OrderEditPreviewFailure) MarshalJSON() ([]byte, error) {
	type Alias OrderEditPreviewFailure
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PreviewFailure", Alias: (*Alias)(&obj)})
}
func (obj *OrderEditPreviewFailure) UnmarshalJSON(data []byte) error {
	type Alias OrderEditPreviewFailure
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Errors {
		obj.Errors[i] = AbstractErrorObjectDiscriminatorMapping(obj.Errors[i])
	}

	return nil
}

type OrderEditPreviewSuccess struct {
	Preview         *StagedOrder     `json:"preview"`
	MessagePayloads []MessagePayload `json:"messagePayloads"`
}

func (obj OrderEditPreviewSuccess) MarshalJSON() ([]byte, error) {
	type Alias OrderEditPreviewSuccess
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PreviewSuccess", Alias: (*Alias)(&obj)})
}
func (obj *OrderEditPreviewSuccess) UnmarshalJSON(data []byte) error {
	type Alias OrderEditPreviewSuccess
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.MessagePayloads {
		obj.MessagePayloads[i] = AbstractMessagePayloadDiscriminatorMapping(obj.MessagePayloads[i])
	}

	return nil
}

type OrderEditReference struct {
	Key string     `json:"key,omitempty"`
	ID  string     `json:"id,omitempty"`
	Obj *OrderEdit `json:"obj,omitempty"`
}

func (obj OrderEditReference) MarshalJSON() ([]byte, error) {
	type Alias OrderEditReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "order-edit", Alias: (*Alias)(&obj)})
}

type OrderEditResult interface{}
type AbstractOrderEditResult struct{}

func AbstractOrderEditResultDiscriminatorMapping(input OrderEditResult) OrderEditResult {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "Applied":
		new := OrderEditApplied{}
		mapstructure.Decode(input, &new)
		return new
	case "NotProcessed":
		new := OrderEditNotProcessed{}
		mapstructure.Decode(input, &new)
		return new
	case "PreviewFailure":
		new := OrderEditPreviewFailure{}
		mapstructure.Decode(input, &new)
		for i := range new.Errors {
			new.Errors[i] = AbstractErrorObjectDiscriminatorMapping(new.Errors[i])
		}

		return new
	case "PreviewSuccess":
		new := OrderEditPreviewSuccess{}
		mapstructure.Decode(input, &new)
		for i := range new.MessagePayloads {
			new.MessagePayloads[i] = AbstractMessagePayloadDiscriminatorMapping(new.MessagePayloads[i])
		}

		return new
	}
	return nil
}

type OrderEditSetCommentAction struct {
	Comment string `json:"comment,omitempty"`
}

func (obj OrderEditSetCommentAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetCommentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setComment", Alias: (*Alias)(&obj)})
}

type OrderEditSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj OrderEditSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type OrderEditSetCustomTypeAction struct {
	Type   *TypeReference `json:"type,omitempty"`
	Fields interface{}    `json:"fields,omitempty"`
}

func (obj OrderEditSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type OrderEditSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

func (obj OrderEditSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type OrderEditSetStagedActionsAction struct {
	StagedActions []StagedOrderUpdateAction `json:"stagedActions"`
}

func (obj OrderEditSetStagedActionsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetStagedActionsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStagedActions", Alias: (*Alias)(&obj)})
}
func (obj *OrderEditSetStagedActionsAction) UnmarshalJSON(data []byte) error {
	type Alias OrderEditSetStagedActionsAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.StagedActions {
		obj.StagedActions[i] = AbstractStagedOrderUpdateActionDiscriminatorMapping(obj.StagedActions[i])
	}

	return nil
}

type OrderEditUpdate struct {
	Version int                     `json:"version"`
	DryRun  bool                    `json:"dryRun"`
	Actions []OrderEditUpdateAction `json:"actions"`
}

func (obj *OrderEditUpdate) UnmarshalJSON(data []byte) error {
	type Alias OrderEditUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractOrderEditUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type OrderEditUpdateAction interface{}
type AbstractOrderEditUpdateAction struct{}

func AbstractOrderEditUpdateActionDiscriminatorMapping(input OrderEditUpdateAction) OrderEditUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addStagedAction":
		new := OrderEditAddStagedActionAction{}
		mapstructure.Decode(input, &new)
		if new.StagedAction != nil {
			new.StagedAction = AbstractStagedOrderUpdateActionDiscriminatorMapping(new.StagedAction)
		}

		return new
	case "setComment":
		new := OrderEditSetCommentAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := OrderEditSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := OrderEditSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := OrderEditSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setStagedActions":
		new := OrderEditSetStagedActionsAction{}
		mapstructure.Decode(input, &new)
		for i := range new.StagedActions {
			new.StagedActions[i] = AbstractStagedOrderUpdateActionDiscriminatorMapping(new.StagedActions[i])
		}

		return new
	}
	return nil
}

type OrderExcerpt struct {
	Version    int         `json:"version"`
	TotalPrice *Money      `json:"totalPrice"`
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitempty"`
}

type StagedOrder struct {
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

type StagedOrderAddCustomLineItemAction struct {
	TaxCategory     *TaxCategoryReference `json:"taxCategory,omitempty"`
	Slug            string                `json:"slug"`
	Quantity        float64               `json:"quantity,omitempty"`
	Name            *LocalizedString      `json:"name"`
	Money           *Money                `json:"money"`
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	Custom          *CustomFieldsDraft    `json:"custom,omitempty"`
}

func (obj StagedOrderAddCustomLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddCustomLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addCustomLineItem", Alias: (*Alias)(&obj)})
}

type StagedOrderAddDeliveryAction struct {
	Parcels []ParcelDraft  `json:"parcels,omitempty"`
	Items   []DeliveryItem `json:"items,omitempty"`
	Address *Address       `json:"address,omitempty"`
}

func (obj StagedOrderAddDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDelivery", Alias: (*Alias)(&obj)})
}

type StagedOrderAddDiscountCodeAction struct {
	Code string `json:"code"`
}

func (obj StagedOrderAddDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDiscountCode", Alias: (*Alias)(&obj)})
}

type StagedOrderAddItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

func (obj StagedOrderAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addItemShippingAddress", Alias: (*Alias)(&obj)})
}

type StagedOrderAddLineItemAction struct {
	VariantID           int                         `json:"variantId,omitempty"`
	SupplyChannel       *ChannelReference           `json:"supplyChannel,omitempty"`
	Sku                 string                      `json:"sku,omitempty"`
	ShippingDetails     *ItemShippingDetailsDraft   `json:"shippingDetails,omitempty"`
	Quantity            float64                     `json:"quantity,omitempty"`
	ProductID           string                      `json:"productId,omitempty"`
	ExternalTotalPrice  *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalTaxRate     *ExternalTaxRateDraft       `json:"externalTaxRate,omitempty"`
	ExternalPrice       *Money                      `json:"externalPrice,omitempty"`
	DistributionChannel *ChannelReference           `json:"distributionChannel,omitempty"`
	Custom              *CustomFieldsDraft          `json:"custom,omitempty"`
}

func (obj StagedOrderAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

type StagedOrderAddParcelToDeliveryAction struct {
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	Items        []DeliveryItem      `json:"items,omitempty"`
	DeliveryID   string              `json:"deliveryId"`
}

func (obj StagedOrderAddParcelToDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddParcelToDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addParcelToDelivery", Alias: (*Alias)(&obj)})
}

type StagedOrderAddPaymentAction struct {
	Payment *PaymentReference `json:"payment"`
}

func (obj StagedOrderAddPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPayment", Alias: (*Alias)(&obj)})
}

type StagedOrderAddReturnInfoAction struct {
	ReturnTrackingID string            `json:"returnTrackingId,omitempty"`
	ReturnDate       time.Time         `json:"returnDate,omitempty"`
	Items            []ReturnItemDraft `json:"items"`
}

func (obj StagedOrderAddReturnInfoAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddReturnInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addReturnInfo", Alias: (*Alias)(&obj)})
}

type StagedOrderAddShoppingListAction struct {
	SupplyChannel       *ChannelReference      `json:"supplyChannel,omitempty"`
	ShoppingList        *ShoppingListReference `json:"shoppingList"`
	DistributionChannel *ChannelReference      `json:"distributionChannel,omitempty"`
}

func (obj StagedOrderAddShoppingListAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddShoppingListAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShoppingList", Alias: (*Alias)(&obj)})
}

type StagedOrderChangeCustomLineItemMoneyAction struct {
	Money            *Money `json:"money"`
	CustomLineItemID string `json:"customLineItemId"`
}

func (obj StagedOrderChangeCustomLineItemMoneyAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeCustomLineItemMoneyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomLineItemMoney", Alias: (*Alias)(&obj)})
}

type StagedOrderChangeCustomLineItemQuantityAction struct {
	Quantity         float64 `json:"quantity"`
	CustomLineItemID string  `json:"customLineItemId"`
}

func (obj StagedOrderChangeCustomLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeCustomLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomLineItemQuantity", Alias: (*Alias)(&obj)})
}

type StagedOrderChangeLineItemQuantityAction struct {
	Quantity           float64                     `json:"quantity"`
	LineItemID         string                      `json:"lineItemId"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalPrice      *Money                      `json:"externalPrice,omitempty"`
}

func (obj StagedOrderChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

type StagedOrderChangeOrderStateAction struct {
	OrderState OrderState `json:"orderState"`
}

func (obj StagedOrderChangeOrderStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeOrderStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeOrderState", Alias: (*Alias)(&obj)})
}

type StagedOrderChangePaymentStateAction struct {
	PaymentState PaymentState `json:"paymentState,omitempty"`
}

func (obj StagedOrderChangePaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangePaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePaymentState", Alias: (*Alias)(&obj)})
}

type StagedOrderChangeShipmentStateAction struct {
	ShipmentState ShipmentState `json:"shipmentState,omitempty"`
}

func (obj StagedOrderChangeShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeShipmentState", Alias: (*Alias)(&obj)})
}

type StagedOrderChangeTaxCalculationModeAction struct {
	TaxCalculationMode TaxCalculationMode `json:"taxCalculationMode"`
}

func (obj StagedOrderChangeTaxCalculationModeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeTaxCalculationModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxCalculationMode", Alias: (*Alias)(&obj)})
}

type StagedOrderChangeTaxModeAction struct {
	TaxMode TaxMode `json:"taxMode"`
}

func (obj StagedOrderChangeTaxModeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeTaxModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxMode", Alias: (*Alias)(&obj)})
}

type StagedOrderChangeTaxRoundingModeAction struct {
	TaxRoundingMode RoundingMode `json:"taxRoundingMode"`
}

func (obj StagedOrderChangeTaxRoundingModeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeTaxRoundingModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxRoundingMode", Alias: (*Alias)(&obj)})
}

type StagedOrderImportCustomLineItemStateAction struct {
	State            []ItemState `json:"state"`
	CustomLineItemID string      `json:"customLineItemId"`
}

func (obj StagedOrderImportCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderImportCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importCustomLineItemState", Alias: (*Alias)(&obj)})
}

type StagedOrderImportLineItemStateAction struct {
	State      []ItemState `json:"state"`
	LineItemID string      `json:"lineItemId"`
}

func (obj StagedOrderImportLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderImportLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importLineItemState", Alias: (*Alias)(&obj)})
}

type StagedOrderRemoveCustomLineItemAction struct {
	CustomLineItemID string `json:"customLineItemId"`
}

func (obj StagedOrderRemoveCustomLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveCustomLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeCustomLineItem", Alias: (*Alias)(&obj)})
}

type StagedOrderRemoveDeliveryAction struct {
	DeliveryID string `json:"deliveryId"`
}

func (obj StagedOrderRemoveDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDelivery", Alias: (*Alias)(&obj)})
}

type StagedOrderRemoveDiscountCodeAction struct {
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

func (obj StagedOrderRemoveDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDiscountCode", Alias: (*Alias)(&obj)})
}

type StagedOrderRemoveItemShippingAddressAction struct {
	AddressKey string `json:"addressKey"`
}

func (obj StagedOrderRemoveItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeItemShippingAddress", Alias: (*Alias)(&obj)})
}

type StagedOrderRemoveLineItemAction struct {
	ShippingDetailsToRemove *ItemShippingDetailsDraft   `json:"shippingDetailsToRemove,omitempty"`
	Quantity                float64                     `json:"quantity,omitempty"`
	LineItemID              string                      `json:"lineItemId"`
	ExternalTotalPrice      *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalPrice           *Money                      `json:"externalPrice,omitempty"`
}

func (obj StagedOrderRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

type StagedOrderRemoveParcelFromDeliveryAction struct {
	ParcelID string `json:"parcelId"`
}

func (obj StagedOrderRemoveParcelFromDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveParcelFromDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeParcelFromDelivery", Alias: (*Alias)(&obj)})
}

type StagedOrderRemovePaymentAction struct {
	Payment *PaymentReference `json:"payment"`
}

func (obj StagedOrderRemovePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemovePaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePayment", Alias: (*Alias)(&obj)})
}

type StagedOrderSetBillingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

func (obj StagedOrderSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCountryAction struct {
	Country string `json:"country,omitempty"`
}

func (obj StagedOrderSetCountryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCountryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCountry", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj StagedOrderSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomLineItemCustomFieldAction struct {
	Value            interface{} `json:"value,omitempty"`
	Name             string      `json:"name"`
	CustomLineItemID string      `json:"customLineItemId"`
}

func (obj StagedOrderSetCustomLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomLineItemCustomTypeAction struct {
	Type             *TypeReference  `json:"type,omitempty"`
	Fields           *FieldContainer `json:"fields,omitempty"`
	CustomLineItemID string          `json:"customLineItemId"`
}

func (obj StagedOrderSetCustomLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomType", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomLineItemShippingDetailsAction struct {
	ShippingDetails  *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	CustomLineItemID string                    `json:"customLineItemId"`
}

func (obj StagedOrderSetCustomLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomLineItemTaxAmountAction struct {
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
	CustomLineItemID  string                  `json:"customLineItemId"`
}

func (obj StagedOrderSetCustomLineItemTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemTaxAmount", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomLineItemTaxRateAction struct {
	ExternalTaxRate  *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	CustomLineItemID string                `json:"customLineItemId"`
}

func (obj StagedOrderSetCustomLineItemTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemTaxRate", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomShippingMethodAction struct {
	TaxCategory        *TaxCategoryReference `json:"taxCategory,omitempty"`
	ShippingRate       *ShippingRateDraft    `json:"shippingRate"`
	ShippingMethodName string                `json:"shippingMethodName"`
	ExternalTaxRate    *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

func (obj StagedOrderSetCustomShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomShippingMethod", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

func (obj StagedOrderSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomerEmailAction struct {
	Email string `json:"email,omitempty"`
}

func (obj StagedOrderSetCustomerEmailAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomerEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerEmail", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomerGroupAction struct {
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
}

func (obj StagedOrderSetCustomerGroupAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomerGroupAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerGroup", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomerIdAction struct {
	CustomerID string `json:"customerId,omitempty"`
}

func (obj StagedOrderSetCustomerIdAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomerIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerId", Alias: (*Alias)(&obj)})
}

type StagedOrderSetDeliveryAddressAction struct {
	DeliveryID string   `json:"deliveryId"`
	Address    *Address `json:"address,omitempty"`
}

func (obj StagedOrderSetDeliveryAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetDeliveryAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddress", Alias: (*Alias)(&obj)})
}

type StagedOrderSetDeliveryItemsAction struct {
	Items      []DeliveryItem `json:"items"`
	DeliveryID string         `json:"deliveryId"`
}

func (obj StagedOrderSetDeliveryItemsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetDeliveryItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryItems", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLineItemCustomFieldAction struct {
	Value      interface{} `json:"value,omitempty"`
	Name       string      `json:"name"`
	LineItemID string      `json:"lineItemId"`
}

func (obj StagedOrderSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLineItemCustomTypeAction struct {
	Type       *TypeReference  `json:"type,omitempty"`
	LineItemID string          `json:"lineItemId"`
	Fields     *FieldContainer `json:"fields,omitempty"`
}

func (obj StagedOrderSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLineItemPriceAction struct {
	LineItemID    string `json:"lineItemId"`
	ExternalPrice *Money `json:"externalPrice,omitempty"`
}

func (obj StagedOrderSetLineItemPriceAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemPrice", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLineItemShippingDetailsAction struct {
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	LineItemID      string                    `json:"lineItemId"`
}

func (obj StagedOrderSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLineItemTaxAmountAction struct {
	LineItemID        string                  `json:"lineItemId"`
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
}

func (obj StagedOrderSetLineItemTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTaxAmount", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLineItemTaxRateAction struct {
	LineItemID      string                `json:"lineItemId"`
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

func (obj StagedOrderSetLineItemTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTaxRate", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLineItemTotalPriceAction struct {
	LineItemID         string                      `json:"lineItemId"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
}

func (obj StagedOrderSetLineItemTotalPriceAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemTotalPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTotalPrice", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

func (obj StagedOrderSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type StagedOrderSetOrderNumberAction struct {
	OrderNumber string `json:"orderNumber,omitempty"`
}

func (obj StagedOrderSetOrderNumberAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetOrderNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setOrderNumber", Alias: (*Alias)(&obj)})
}

type StagedOrderSetOrderTotalTaxAction struct {
	ExternalTotalGross  *Money       `json:"externalTotalGross"`
	ExternalTaxPortions []TaxPortion `json:"externalTaxPortions,omitempty"`
}

func (obj StagedOrderSetOrderTotalTaxAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetOrderTotalTaxAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setOrderTotalTax", Alias: (*Alias)(&obj)})
}

type StagedOrderSetParcelItemsAction struct {
	ParcelID string         `json:"parcelId"`
	Items    []DeliveryItem `json:"items"`
}

func (obj StagedOrderSetParcelItemsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetParcelItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelItems", Alias: (*Alias)(&obj)})
}

type StagedOrderSetParcelMeasurementsAction struct {
	ParcelID     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
}

func (obj StagedOrderSetParcelMeasurementsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetParcelMeasurementsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelMeasurements", Alias: (*Alias)(&obj)})
}

type StagedOrderSetParcelTrackingDataAction struct {
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	ParcelID     string        `json:"parcelId"`
}

func (obj StagedOrderSetParcelTrackingDataAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetParcelTrackingDataAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelTrackingData", Alias: (*Alias)(&obj)})
}

type StagedOrderSetReturnPaymentStateAction struct {
	ReturnItemID string             `json:"returnItemId"`
	PaymentState ReturnPaymentState `json:"paymentState"`
}

func (obj StagedOrderSetReturnPaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetReturnPaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnPaymentState", Alias: (*Alias)(&obj)})
}

type StagedOrderSetReturnShipmentStateAction struct {
	ShipmentState ReturnShipmentState `json:"shipmentState"`
	ReturnItemID  string              `json:"returnItemId"`
}

func (obj StagedOrderSetReturnShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetReturnShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnShipmentState", Alias: (*Alias)(&obj)})
}

type StagedOrderSetShippingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

func (obj StagedOrderSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

type StagedOrderSetShippingAddressAndCustomShippingMethodAction struct {
	TaxCategory        *TaxCategoryReference `json:"taxCategory,omitempty"`
	ShippingRate       *ShippingRateDraft    `json:"shippingRate"`
	ShippingMethodName string                `json:"shippingMethodName"`
	ExternalTaxRate    *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	Address            *Address              `json:"address"`
}

func (obj StagedOrderSetShippingAddressAndCustomShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingAddressAndCustomShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressAndCustomShippingMethod", Alias: (*Alias)(&obj)})
}

type StagedOrderSetShippingAddressAndShippingMethodAction struct {
	ShippingMethod  *ShippingMethodReference `json:"shippingMethod,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft    `json:"externalTaxRate,omitempty"`
	Address         *Address                 `json:"address"`
}

func (obj StagedOrderSetShippingAddressAndShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingAddressAndShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressAndShippingMethod", Alias: (*Alias)(&obj)})
}

type StagedOrderSetShippingMethodAction struct {
	ShippingMethod  *TypeReference        `json:"shippingMethod,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

func (obj StagedOrderSetShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethod", Alias: (*Alias)(&obj)})
}

type StagedOrderSetShippingMethodTaxAmountAction struct {
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
}

func (obj StagedOrderSetShippingMethodTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingMethodTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethodTaxAmount", Alias: (*Alias)(&obj)})
}

type StagedOrderSetShippingMethodTaxRateAction struct {
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

func (obj StagedOrderSetShippingMethodTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingMethodTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethodTaxRate", Alias: (*Alias)(&obj)})
}

type StagedOrderSetShippingRateInputAction struct {
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
}

func (obj StagedOrderSetShippingRateInputAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingRateInputAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingRateInput", Alias: (*Alias)(&obj)})
}
func (obj *StagedOrderSetShippingRateInputAction) UnmarshalJSON(data []byte) error {
	type Alias StagedOrderSetShippingRateInputAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		obj.ShippingRateInput = AbstractShippingRateInputDraftDiscriminatorMapping(obj.ShippingRateInput)
	}

	return nil
}

type StagedOrderTransitionCustomLineItemStateAction struct {
	ToState              *StateReference `json:"toState"`
	Quantity             int             `json:"quantity"`
	FromState            *StateReference `json:"fromState"`
	CustomLineItemID     string          `json:"customLineItemId"`
	ActualTransitionDate time.Time       `json:"actualTransitionDate,omitempty"`
}

func (obj StagedOrderTransitionCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderTransitionCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionCustomLineItemState", Alias: (*Alias)(&obj)})
}

type StagedOrderTransitionLineItemStateAction struct {
	ToState              *StateReference `json:"toState"`
	Quantity             int             `json:"quantity"`
	LineItemID           string          `json:"lineItemId"`
	FromState            *StateReference `json:"fromState"`
	ActualTransitionDate time.Time       `json:"actualTransitionDate,omitempty"`
}

func (obj StagedOrderTransitionLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderTransitionLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionLineItemState", Alias: (*Alias)(&obj)})
}

type StagedOrderTransitionStateAction struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force,omitempty"`
}

func (obj StagedOrderTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

type StagedOrderUpdateItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

func (obj StagedOrderUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

type StagedOrderUpdateSyncInfoAction struct {
	SyncedAt   time.Time         `json:"syncedAt,omitempty"`
	ExternalID string            `json:"externalId,omitempty"`
	Channel    *ChannelReference `json:"channel"`
}

func (obj StagedOrderUpdateSyncInfoAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderUpdateSyncInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateSyncInfo", Alias: (*Alias)(&obj)})
}
