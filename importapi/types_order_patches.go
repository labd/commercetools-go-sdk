package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

/**
*	Maps to `ReturnItem.shipmentState`
 */
type ReturnShipmentState string

const (
	ReturnShipmentStateAdvised     ReturnShipmentState = "Advised"
	ReturnShipmentStateReturned    ReturnShipmentState = "Returned"
	ReturnShipmentStateBackInStock ReturnShipmentState = "BackInStock"
	ReturnShipmentStateUnusable    ReturnShipmentState = "Unusable"
)

type ReturnItemDraft struct {
	// Number of Line Items or Custom Line Items to return.
	Quantity int `json:"quantity"`
	// `id` of the [LineItem](ctp:api:type:LineItem) to return.
	//
	// Required if Line Items are returned, to create a [LineItemReturnItem](ctp:api:type:LineItemReturnItem).
	LineItemId *string `json:"lineItemId,omitempty"`
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to return.
	//
	// Required if Custom Line Items are returned, to create a [CustomLineItemReturnItem](ctp:api:type:CustomLineItemReturnItem).
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// User-defined description for the return.
	Comment *string `json:"comment,omitempty"`
	// Shipment status of the item to be returned.
	ShipmentState ReturnShipmentState `json:"shipmentState"`
}

type ReturnInfo struct {
	// Information on the Line Items or Custom Line Items returned.
	Items []ReturnItemDraft `json:"items"`
	// User-defined identifier to track the return.
	ReturnTrackingId *string `json:"returnTrackingId,omitempty"`
	// Date and time (UTC) the return is initiated.
	ReturnDate *time.Time `json:"returnDate,omitempty"`
}

type DeliveryParcel struct {
	// Unique identifier of the Delivery.
	DeliveryId string `json:"deliveryId"`
	// Information about the dimensions of the Parcel.
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	// Shipment tracking information of the Parcel.
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	// Line Items or Custom Line Items delivered in this Parcel.
	Items []DeliveryItem `json:"items"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeliveryParcel) MarshalJSON() ([]byte, error) {
	type Alias DeliveryParcel
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

	if raw["items"] == nil {
		delete(raw, "items")
	}

	return json.Marshal(raw)

}

type DeliveryParcelDraft struct {
	// Information about the dimensions for the Parcel.
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	// Shipment tracking information for the Parcel.
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	// Line Items or Custom Line Items delivered in this Parcel.
	Items []DeliveryItem `json:"items"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeliveryParcelDraft) MarshalJSON() ([]byte, error) {
	type Alias DeliveryParcelDraft
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

	if raw["items"] == nil {
		delete(raw, "items")
	}

	return json.Marshal(raw)

}

type DeliveryDraft struct {
	// Line Items or Custom Line Items to deliver. It can also be specified individually for each [Parcel](ctp:api:type:Parcel).
	Items []DeliveryItem `json:"items"`
	// Address to which the Parcels are delivered.
	Address *Address `json:"address,omitempty"`
	// Information regarding the appearance, content, and shipment of a parcel.
	Parcels []DeliveryParcelDraft `json:"parcels"`
}

type DeliveryAddressDraft struct {
	// Unique identifier of the Delivery.
	DeliveryId string `json:"deliveryId"`
	// Address to which Parcels are delivered.
	Address *Address `json:"address,omitempty"`
}

type ParcelMeasurementDraft struct {
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	ParcelId string `json:"parcelId"`
	// Information about the dimensions of the Parcel.
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
}

type ParcelTrackingData struct {
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	ParcelId string `json:"parcelId"`
	// Information that helps track a Parcel.
	TrackingData *TrackingData `json:"trackingData,omitempty"`
}

type ParcelItems struct {
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	ParcelId string `json:"parcelId"`
	// Items in the Parcel.
	Items []DeliveryItem `json:"items"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ParcelItems) MarshalJSON() ([]byte, error) {
	type Alias ParcelItems
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

	if raw["items"] == nil {
		delete(raw, "items")
	}

	return json.Marshal(raw)

}

type RemoveDeliveryDraft struct {
	ID string `json:"id"`
}

type RemoveParcelFromDeliveryDraft struct {
	// `id` of the [Parcel](ctp:api:type:Parcel) to be removed from the Delivery.
	ParcelId string `json:"parcelId"`
}

/**
*	Order fields that needs to be added or updated.
*
 */
type OrderField struct {
	// Maps to `Order.returnInfo`
	AddReturnInfo *ReturnInfo `json:"addReturnInfo,omitempty"`
	// Maps to `Order.delivery`
	AddParcelToDelivery *DeliveryParcel `json:"addParcelToDelivery,omitempty"`
	// Maps to `Order.delivery`
	AddDeliveries []DeliveryDraft `json:"addDeliveries"`
	// Maps to `Order.removeDelivery`
	RemoveDelivery *RemoveDeliveryDraft `json:"removeDelivery,omitempty"`
	// Maps to `Order.removeParcelFromDelivery`
	RemoveParcelFromDelivery *RemoveParcelFromDeliveryDraft `json:"removeParcelFromDelivery,omitempty"`
	// Maps to `Order.addressDraft`
	SetDeliveryAddress *DeliveryAddressDraft `json:"setDeliveryAddress,omitempty"`
	// Maps to `Order.parcelMeasurements`
	SetParcelMeasurements *ParcelMeasurementDraft `json:"setParcelMeasurements,omitempty"`
	// Maps to `Order.parcelTrackingData`
	SetParcelTrackingData *ParcelTrackingData `json:"setParcelTrackingData,omitempty"`
	// Maps to `Order.parcelItems`
	SetParcelItems []ParcelItems `json:"setParcelItems"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderField) MarshalJSON() ([]byte, error) {
	type Alias OrderField
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

	if raw["addDeliveries"] == nil {
		delete(raw, "addDeliveries")
	}

	if raw["setParcelItems"] == nil {
		delete(raw, "setParcelItems")
	}

	return json.Marshal(raw)

}

/**
*	Represents the data used to update an [Order](ctp:api:type:Order) in a Project.
*
 */
type OrderPatchImport struct {
	// User-defined unique identifier. If an [Order](ctp:api:type:Order) with this `orderNumber` exists, it is updated with the imported data.
	OrderNumber string `json:"orderNumber"`
	// Each field referenced must be defined in an existing [Order](ctp:api:type:Order) or the [ImportOperationState](ctp:import:type:ImportOperationState) is set to `validationFailed`.
	Fields OrderField `json:"fields"`
}
