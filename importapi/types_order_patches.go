// Generated file, please do not change!!!
package importapi

import (
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
	Quantity         float64 `json:"quantity"`
	LineItemId       *string `json:"lineItemId,omitempty"`
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	Comment          *string `json:"comment,omitempty"`
	// Maps to `ReturnItem.shipmentState`
	ShipmentState ReturnShipmentState `json:"shipmentState"`
}

type ReturnInfo struct {
	Items []ReturnItemDraft `json:"items"`
	// Maps to `ReturnInfo.returnTrackingId`
	ReturnTrackingId *string `json:"returnTrackingId,omitempty"`
	// Maps to `ReturnInfo.returnDate`
	ReturnDate *time.Time `json:"returnDate,omitempty"`
}

type DeliveryParcel struct {
	DeliveryId   string              `json:"deliveryId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Items        []DeliveryItem      `json:"items,omitempty"`
}

type DeliveryParcelDraft struct {
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Items        []DeliveryItem      `json:"items,omitempty"`
}

type DeliveryDraft struct {
	Items   []DeliveryItem        `json:"items"`
	Address *Address              `json:"address,omitempty"`
	Parcels []DeliveryParcelDraft `json:"parcels"`
}

type DeliveryAddressDraft struct {
	DeliveryId string   `json:"deliveryId"`
	Address    *Address `json:"address,omitempty"`
}

type ParcelMeasurementDraft struct {
	ParcelId     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
}

type ParcelTrackingData struct {
	ParcelId     string        `json:"parcelId"`
	TrackingData *TrackingData `json:"trackingData,omitempty"`
}

type ParcelItems struct {
	ParcelId string         `json:"parcelId"`
	Items    []DeliveryItem `json:"items,omitempty"`
}

type RemoveDeliveryDraft struct {
	ID string `json:"id"`
}

type RemoveParcelFromDeliveryDraft struct {
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
	AddDeliveries []DeliveryDraft `json:"addDeliveries,omitempty"`
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
	SetParcelItems []ParcelItems `json:"setParcelItems,omitempty"`
}

/**
*	Representation for an update of an [Order](/../api/projects/orders#order). Use this type to import updates for existing
*	[Orders](/../api/projects/orders#order) in a commercetools Project.
*
 */
type OrderPatchImport struct {
	// Maps to `Order.orderNumber`, String that uniquely identifies an order, unique across a project.
	OrderNumber string `json:"orderNumber"`
	// Each field referenced must be defined in an already existing order in the commercetools project or the import operation state is set to `validationFailed`.
	Fields OrderField `json:"fields"`
}
