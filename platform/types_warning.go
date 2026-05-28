package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	Represents a warning related to the returned response.
 */
type WarningObject interface{}

func mapDiscriminatorWarningObject(input interface{}) (WarningObject, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["code"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'code'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "CannotChangeReservationExpiry":
		obj := CannotChangeReservationExpiryWarning{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CannotCreateReservation":
		obj := CannotCreateReservationWarning{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CannotUpdateReservation":
		obj := CannotUpdateReservationWarning{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ImageProcessingOngoing":
		obj := ImageProcessingOngoingWarning{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Contained in responses to a [Cart](/projects/carts) update operation.
*	Indicates that a `ReserveOnCart` [Reservation](ctp:api:type:Reservation) was found for a Line Item but its
*	expiration date and time cannot be changed using the [Set Reservation Expiration In Minutes](/projects/carts#set-reservation-expiration-in-minutes) update action.
*
 */
type CannotChangeReservationExpiryWarning struct {
	// `"The reservation expiration for line item $lineItemId cannot be changed."`
	Message string `json:"message"`
	// The ID of the Line Item.
	LineItemId string `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CannotChangeReservationExpiryWarning) MarshalJSON() ([]byte, error) {
	type Alias CannotChangeReservationExpiryWarning
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "CannotChangeReservationExpiry", Alias: (*Alias)(&obj)})
}

/**
*	Contained in responses to a [Cart](/projects/carts) operation. Indicates that a reservation could not be created for the specified Product.
*
 */
type CannotCreateReservationWarning struct {
	// `"Failed to create a reservation for product $productId (SKU: $sku, Supply Channel: $supplyChannel) with quantity $quantity."`
	Message string `json:"message"`
	// The ID of the Product for which the reservation could not be created.
	ProductId string `json:"productId"`
	// The SKU of the Product Variant for which the reservation could not be created.
	Sku *string `json:"sku,omitempty"`
	// The ID of the Supply Channel for which the reservation could not be created.
	SupplyChannel *string `json:"supplyChannel,omitempty"`
	// The quantity for which the reservation could not be created.
	Quantity float64 `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CannotCreateReservationWarning) MarshalJSON() ([]byte, error) {
	type Alias CannotCreateReservationWarning
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "CannotCreateReservation", Alias: (*Alias)(&obj)})
}

/**
*	Contained in responses to a [Cart](/projects/carts) operation. Indicates that the reservation for a Line Item could not be updated to the requested quantity. The existing reserved amount remains unchanged.
*
 */
type CannotUpdateReservationWarning struct {
	// `"Failed to update the reservation for line item $lineItemId (Product ID: $productId, SKU: $sku, Supply Channel: $supplyChannel) to a quantity of $requestedQuantity. The reserved amount is still $reservedQuantity."`
	Message string `json:"message"`
	// The ID of the Line Item.
	LineItemId string `json:"lineItemId"`
	// The quantity that was requested for the Line Item.
	RequestedQuantity float64 `json:"requestedQuantity"`
	// The quantity that remains reserved for the Line Item.
	ReservedQuantity float64 `json:"reservedQuantity"`
	// The ID of the Product associated with the Line Item.
	ProductId string `json:"productId"`
	// The SKU of the Product Variant associated with the Line Item.
	Sku *string `json:"sku,omitempty"`
	// The ID of the Supply Channel associated with the reservation.
	SupplyChannel *string `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CannotUpdateReservationWarning) MarshalJSON() ([]byte, error) {
	type Alias CannotUpdateReservationWarning
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "CannotUpdateReservation", Alias: (*Alias)(&obj)})
}

/**
*	Contained in responses to an [Upload Product image](/projects/products#upload-product-image) or an [Upload Product Tailoring image](/projects/product-tailoring#upload-product-tailoring-image) request with response status code `202 Accepted`.
*	Indicates that the API is still creating the remaining sizes of the uploaded image. They will be available on the Content Delivery Network (CDN) soon.
*
 */
type ImageProcessingOngoingWarning struct {
	// `"The image processing is still ongoing."`
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ImageProcessingOngoingWarning) MarshalJSON() ([]byte, error) {
	type Alias ImageProcessingOngoingWarning
	return json.Marshal(struct {
		Action string `json:"code"`
		*Alias
	}{Action: "ImageProcessingOngoing", Alias: (*Alias)(&obj)})
}
