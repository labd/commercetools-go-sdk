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
	LineItemId       string  `json:"lineItemId,omitEmpty"`
	CustomLineItemId string  `json:"customLineItemId,omitEmpty"`
	Comment          string  `json:"comment,omitEmpty"`
	// Maps to `ReturnItem.shipmentState`
	ShipmentState ReturnShipmentState `json:"shipmentState"`
}

type ReturnInfo struct {
	Items []ReturnItemDraft `json:"items"`
	// Maps to `ReturnInfo.returnTrackingId`
	ReturnTrackingId string `json:"returnTrackingId,omitEmpty"`
	// Maps to `ReturnInfo.returnDate`
	ReturnDate time.Time `json:"returnDate,omitEmpty"`
}

type DeliveryParcel struct {
	DeliveryId   string              `json:"deliveryId"`
	Measurements *ParcelMeasurements `json:"measurements,omitEmpty"`
	TrackingData *TrackingData       `json:"trackingData,omitEmpty"`
	Items        []DeliveryItem      `json:"items,omitEmpty"`
}

type DeliveryParcelDraft struct {
	Measurements *ParcelMeasurements `json:"measurements,omitEmpty"`
	TrackingData *TrackingData       `json:"trackingData,omitEmpty"`
	Items        []DeliveryItem      `json:"items,omitEmpty"`
}

type DeliveryDraft struct {
	Items   []DeliveryItem        `json:"items"`
	Address *Address              `json:"address,omitEmpty"`
	Parcels []DeliveryParcelDraft `json:"parcels"`
}

type DeliveryAddressDraft struct {
	DeliveryId string   `json:"deliveryId"`
	Address    *Address `json:"address,omitEmpty"`
}

type ParcelMeasurementDraft struct {
	ParcelId     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitEmpty"`
}

type ParcelTrackingData struct {
	ParcelId     string        `json:"parcelId"`
	TrackingData *TrackingData `json:"trackingData,omitEmpty"`
}

type ParcelItems struct {
	ParcelId string         `json:"parcelId"`
	Items    []DeliveryItem `json:"items,omitEmpty"`
}

type RemoveDeliveryDraft struct {
	Id string `json:"id"`
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
	AddReturnInfo *ReturnInfo `json:"addReturnInfo,omitEmpty"`
	// Maps to `Order.delivery`
	AddParcelToDelivery *DeliveryParcel `json:"addParcelToDelivery,omitEmpty"`
	// Maps to `Order.delivery`
	AddDeliveries []DeliveryDraft `json:"addDeliveries,omitEmpty"`
	// Maps to `Order.removeDelivery`
	RemoveDelivery *RemoveDeliveryDraft `json:"removeDelivery,omitEmpty"`
	// Maps to `Order.removeParcelFromDelivery`
	RemoveParcelFromDelivery *RemoveParcelFromDeliveryDraft `json:"removeParcelFromDelivery,omitEmpty"`
	// Maps to `Order.addressDraft`
	SetDeliveryAddress *DeliveryAddressDraft `json:"setDeliveryAddress,omitEmpty"`
	// Maps to `Order.parcelMeasurements`
	SetParcelMeasurements *ParcelMeasurementDraft `json:"setParcelMeasurements,omitEmpty"`
	// Maps to `Order.parcelTrackingData`
	SetParcelTrackingData *ParcelTrackingData `json:"setParcelTrackingData,omitEmpty"`
	// Maps to `Order.parcelItems`
	SetParcelItems []ParcelItems `json:"setParcelItems,omitEmpty"`
}

/**
*	Representation for an update of an [Order](/../api/projects/orders#order). Use this type to import updates for existing
*	[Orders](/../api/projects/orders#order) in a commercetools Project.
*
 */
type OrderPatchImport struct {
	// Maps to `Order.orderNumber`, String that uniquely identifies an order, unique across a project.
	OrderNumber string `json:"orderNumber"`
	// Each field referenced must be defined in an already existing order in the commercetools project or the import operation state is set to `ValidationFailed`.
	Fields OrderField `json:"fields"`
}
