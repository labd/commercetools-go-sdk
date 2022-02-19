// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Custom Objects are grouped into containers, which can be used like namespaces. Within a given container, a user-defined key can be used to uniquely identify resources.
*
 */
type ContainerAndKey struct {
	// User-defined identifier that is unique within the given container.
	Key string `json:"key"`
	// Namespace to group Custom Objects.
	Container string `json:"container"`
}

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
		obj := CategoryCreatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CategorySlugChanged":
		obj := CategorySlugChangedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerAddressAdded":
		obj := CustomerAddressAddedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerAddressChanged":
		obj := CustomerAddressChangedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerAddressRemoved":
		obj := CustomerAddressRemovedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerCompanyNameSet":
		obj := CustomerCompanyNameSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerCreated":
		obj := CustomerCreatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerDateOfBirthSet":
		obj := CustomerDateOfBirthSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerDeleted":
		obj := CustomerDeletedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerEmailChanged":
		obj := CustomerEmailChangedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerEmailVerified":
		obj := CustomerEmailVerifiedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerFirstNameSet":
		obj := CustomerFirstNameSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerGroupSet":
		obj := CustomerGroupSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerLastNameSet":
		obj := CustomerLastNameSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerPasswordUpdated":
		obj := CustomerPasswordUpdatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "CustomerTitleSet":
		obj := CustomerTitleSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "InventoryEntryCreated":
		obj := InventoryEntryCreatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "InventoryEntryDeleted":
		obj := InventoryEntryDeletedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "InventoryEntryQuantitySet":
		obj := InventoryEntryQuantitySetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderPaymentAdded":
		obj := OrderPaymentAddedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "PaymentCreated":
		obj := PaymentCreatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "PaymentInteractionAdded":
		obj := PaymentInteractionAddedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "PaymentStatusInterfaceCodeSet":
		obj := PaymentStatusInterfaceCodeSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "PaymentStatusStateTransition":
		obj := PaymentStatusStateTransitionMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "PaymentTransactionAdded":
		obj := PaymentTransactionAddedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "PaymentTransactionStateChanged":
		obj := PaymentTransactionStateChangedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductAddedToCategory":
		obj := ProductAddedToCategoryMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductCreated":
		obj := ProductCreatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductDeleted":
		obj := ProductDeletedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductImageAdded":
		obj := ProductImageAddedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductPriceDiscountsSet":
		obj := ProductPriceDiscountsSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductPriceExternalDiscountSet":
		obj := ProductPriceExternalDiscountSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductPublished":
		obj := ProductPublishedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductRemovedFromCategory":
		obj := ProductRemovedFromCategoryMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductRevertedStagedChanges":
		obj := ProductRevertedStagedChangesMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductSelectionCreated":
		obj := ProductSelectionCreatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		if obj.ProductSelection != nil {
			var err error
			obj.ProductSelection, err = mapDiscriminatorProductSelectionType(obj.ProductSelection)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductSelectionDeleted":
		obj := ProductSelectionDeletedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductSelectionProductAdded":
		obj := ProductSelectionProductAddedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductSelectionProductRemoved":
		obj := ProductSelectionProductRemovedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductSlugChanged":
		obj := ProductSlugChangedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductStateTransition":
		obj := ProductStateTransitionMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductUnpublished":
		obj := ProductUnpublishedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductVariantAdded":
		obj := ProductVariantAddedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductVariantDeleted":
		obj := ProductVariantDeletedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ReviewCreated":
		obj := ReviewCreatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ReviewRatingSet":
		obj := ReviewRatingSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		if obj.Target != nil {
			var err error
			obj.Target, err = mapDiscriminatorReference(obj.Target)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ReviewStateTransition":
		obj := ReviewStateTransitionMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		if obj.Target != nil {
			var err error
			obj.Target, err = mapDiscriminatorReference(obj.Target)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "StoreCreated":
		obj := StoreCreatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "StoreDeleted":
		obj := StoreDeletedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "StoreProductSelectionsChanged":
		obj := StoreProductSelectionsChangedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

type CategoryCreatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias CategoryCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CategoryCreated", Alias: (*Alias)(&obj)})
}

type CategorySlugChangedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Slug                            LocalizedString          `json:"slug"`
	OldSlug                         *LocalizedString         `json:"oldSlug,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySlugChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias CategorySlugChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CategorySlugChanged", Alias: (*Alias)(&obj)})
}

type CustomerAddressAddedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddressAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerAddressAdded", Alias: (*Alias)(&obj)})
}

type CustomerAddressChangedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddressChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerAddressChanged", Alias: (*Alias)(&obj)})
}

type CustomerAddressRemovedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddressRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerAddressRemoved", Alias: (*Alias)(&obj)})
}

type CustomerCompanyNameSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	CompanyName                     *string                  `json:"companyName,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerCompanyNameSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerCompanyNameSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerCompanyNameSet", Alias: (*Alias)(&obj)})
}

type CustomerCreatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerCreated", Alias: (*Alias)(&obj)})
}

type CustomerDateOfBirthSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	DateOfBirth                     *Date                    `json:"dateOfBirth,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerDateOfBirthSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerDateOfBirthSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerDateOfBirthSet", Alias: (*Alias)(&obj)})
}

type CustomerDeletedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerDeletedMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerDeleted", Alias: (*Alias)(&obj)})
}

type CustomerEmailChangedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerEmailChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerEmailChanged", Alias: (*Alias)(&obj)})
}

type CustomerEmailVerifiedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerEmailVerifiedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailVerifiedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerEmailVerified", Alias: (*Alias)(&obj)})
}

type CustomerFirstNameSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	FirstName                       *string                  `json:"firstName,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerFirstNameSetMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerFirstNameSetMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerFirstNameSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerFirstNameSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerFirstNameSet", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerGroupSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerGroupSet", Alias: (*Alias)(&obj)})
}

type CustomerLastNameSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	LastName                        *string                  `json:"lastName,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerLastNameSetMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerLastNameSetMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerLastNameSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerLastNameSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerLastNameSet", Alias: (*Alias)(&obj)})
}

type CustomerPasswordUpdatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerPasswordUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerPasswordUpdatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerPasswordUpdated", Alias: (*Alias)(&obj)})
}

type CustomerTitleSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Title                           *string                  `json:"title,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerTitleSetMessage) UnmarshalJSON(data []byte) error {
	type Alias CustomerTitleSetMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerTitleSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerTitleSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerTitleSet", Alias: (*Alias)(&obj)})
}

type InventoryEntryCreatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "InventoryEntryCreated", Alias: (*Alias)(&obj)})
}

type InventoryEntryDeletedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Sku                             string                   `json:"sku"`
	// [Reference](/../api/types#reference) to a [Channel](ctp:api:type:Channel).
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "InventoryEntryDeleted", Alias: (*Alias)(&obj)})
}

type InventoryEntryQuantitySetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	OldQuantityOnStock              int                      `json:"oldQuantityOnStock"`
	NewQuantityOnStock              int                      `json:"newQuantityOnStock"`
	OldAvailableQuantity            int                      `json:"oldAvailableQuantity"`
	NewAvailableQuantity            int                      `json:"newAvailableQuantity"`
	// [Reference](/../api/types#reference) to a [Channel](ctp:api:type:Channel).
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryQuantitySetMessage) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryQuantitySetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "InventoryEntryQuantitySet", Alias: (*Alias)(&obj)})
}

type MessagePagedQueryResponse struct {
	Limit   int       `json:"limit"`
	Count   int       `json:"count"`
	Total   *int      `json:"total,omitempty"`
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

type MessagesConfiguration struct {
	Enabled                 bool `json:"enabled"`
	DeleteDaysAfterCreation *int `json:"deleteDaysAfterCreation,omitempty"`
}

type MessagesConfigurationDraft struct {
	Enabled                 bool `json:"enabled"`
	DeleteDaysAfterCreation int  `json:"deleteDaysAfterCreation"`
}

type OrderMessage interface{}

func mapDiscriminatorOrderMessage(input interface{}) (OrderMessage, error) {

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
	case "CustomLineItemStateTransition":
		obj := CustomLineItemStateTransitionMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "DeliveryAdded":
		obj := DeliveryAddedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "DeliveryAddressSet":
		obj := DeliveryAddressSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "DeliveryItemsUpdated":
		obj := DeliveryItemsUpdatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "DeliveryRemoved":
		obj := DeliveryRemovedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "LineItemStateTransition":
		obj := LineItemStateTransitionMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderBillingAddressSet":
		obj := OrderBillingAddressSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderCreated":
		obj := OrderCreatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderCustomLineItemDiscountSet":
		obj := OrderCustomLineItemDiscountSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderCustomerEmailSet":
		obj := OrderCustomerEmailSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderCustomerGroupSet":
		obj := OrderCustomerGroupSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderCustomerSet":
		obj := OrderCustomerSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderDeleted":
		obj := OrderDeletedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderDiscountCodeAdded":
		obj := OrderDiscountCodeAddedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderDiscountCodeRemoved":
		obj := OrderDiscountCodeRemovedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderDiscountCodeStateSet":
		obj := OrderDiscountCodeStateSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderEditApplied":
		obj := OrderEditAppliedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderImported":
		obj := OrderImportedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderLineItemAdded":
		obj := OrderLineItemAddedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderLineItemDiscountSet":
		obj := OrderLineItemDiscountSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderLineItemDistributionChannelSet":
		obj := OrderLineItemDistributionChannelSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderLineItemRemoved":
		obj := OrderLineItemRemovedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		if obj.NewTotalPrice != nil {
			var err error
			obj.NewTotalPrice, err = mapDiscriminatorTypedMoney(obj.NewTotalPrice)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderPaymentStateChanged":
		obj := OrderPaymentStateChangedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ReturnInfoAdded":
		obj := OrderReturnInfoAddedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ReturnInfoSet":
		obj := OrderReturnInfoSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderReturnShipmentStateChanged":
		obj := OrderReturnShipmentStateChangedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderShipmentStateChanged":
		obj := OrderShipmentStateChangedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderShippingAddressSet":
		obj := OrderShippingAddressSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderShippingInfoSet":
		obj := OrderShippingInfoSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderShippingRateInputSet":
		obj := OrderShippingRateInputSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		if obj.ShippingRateInput != nil {
			var err error
			obj.ShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.ShippingRateInput)
			if err != nil {
				return nil, err
			}
		}
		if obj.OldShippingRateInput != nil {
			var err error
			obj.OldShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.OldShippingRateInput)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderStateChanged":
		obj := OrderStateChangedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderStateTransition":
		obj := OrderStateTransitionMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderStoreSet":
		obj := OrderStoreSetMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ParcelAddedToDelivery":
		obj := ParcelAddedToDeliveryMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ParcelItemsUpdated":
		obj := ParcelItemsUpdatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ParcelMeasurementsUpdated":
		obj := ParcelMeasurementsUpdatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ParcelRemovedFromDelivery":
		obj := ParcelRemovedFromDeliveryMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ParcelTrackingDataUpdated":
		obj := ParcelTrackingDataUpdatedMessage{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

type CustomLineItemStateTransitionMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	CustomLineItemId                string                   `json:"customLineItemId"`
	TransitionDate                  time.Time                `json:"transitionDate"`
	Quantity                        int                      `json:"quantity"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	FromState StateReference `json:"fromState"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	ToState StateReference `json:"toState"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomLineItemStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomLineItemStateTransition", Alias: (*Alias)(&obj)})
}

type DeliveryAddedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeliveryAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryAdded", Alias: (*Alias)(&obj)})
}

type DeliveryAddressSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	DeliveryId                      string                   `json:"deliveryId"`
	Address                         *Address                 `json:"address,omitempty"`
	OldAddress                      *Address                 `json:"oldAddress,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeliveryAddressSetMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddressSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryAddressSet", Alias: (*Alias)(&obj)})
}

type DeliveryItemsUpdatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeliveryItemsUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryItemsUpdatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryItemsUpdated", Alias: (*Alias)(&obj)})
}

type DeliveryRemovedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeliveryRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryRemoved", Alias: (*Alias)(&obj)})
}

type LineItemStateTransitionMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	LineItemId                      string                   `json:"lineItemId"`
	TransitionDate                  time.Time                `json:"transitionDate"`
	Quantity                        int                      `json:"quantity"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	FromState StateReference `json:"fromState"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	ToState StateReference `json:"toState"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LineItemStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias LineItemStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "LineItemStateTransition", Alias: (*Alias)(&obj)})
}

type OrderBillingAddressSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Address                         *Address                 `json:"address,omitempty"`
	OldAddress                      *Address                 `json:"oldAddress,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderBillingAddressSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderBillingAddressSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderBillingAddressSet", Alias: (*Alias)(&obj)})
}

type OrderCreatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCreated", Alias: (*Alias)(&obj)})
}

type OrderCustomLineItemDiscountSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy                           `json:"createdBy,omitempty"`
	SequenceNumber                  int                                  `json:"sequenceNumber"`
	Resource                        Reference                            `json:"resource"`
	ResourceVersion                 int                                  `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers             `json:"resourceUserProvidedIdentifiers,omitempty"`
	CustomLineItemId                string                               `json:"customLineItemId"`
	DiscountedPricePerQuantity      []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	TaxedPrice                      *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCustomLineItemDiscountSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemDiscountSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerEmailSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Email                           *string                  `json:"email,omitempty"`
	OldEmail                        *string                  `json:"oldEmail,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCustomerEmailSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerEmailSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomerEmailSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerGroupSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	OldCustomerGroup *CustomerGroupReference `json:"oldCustomerGroup,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCustomerGroupSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerGroupSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomerGroupSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Customer                        *CustomerReference       `json:"customer,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	OldCustomer   *CustomerReference      `json:"oldCustomer,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	OldCustomerGroup *CustomerGroupReference `json:"oldCustomerGroup,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCustomerSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomerSet", Alias: (*Alias)(&obj)})
}

type OrderDeletedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDeleted", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeAddedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderDiscountCodeAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDiscountCodeAdded", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeRemovedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderDiscountCodeRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDiscountCodeRemoved", Alias: (*Alias)(&obj)})
}

type OrderDiscountCodeStateSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	DiscountCode                    DiscountCodeReference    `json:"discountCode"`
	State                           DiscountCodeState        `json:"state"`
	OldState                        *DiscountCodeState       `json:"oldState,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderDiscountCodeStateSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeStateSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderDiscountCodeStateSet", Alias: (*Alias)(&obj)})
}

type OrderEditAppliedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditAppliedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderEditAppliedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderEditApplied", Alias: (*Alias)(&obj)})
}

type OrderImportedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderImportedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderImportedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderImported", Alias: (*Alias)(&obj)})
}

type OrderLineItemAddedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderLineItemAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLineItemAdded", Alias: (*Alias)(&obj)})
}

type OrderLineItemDiscountSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy                           `json:"createdBy,omitempty"`
	SequenceNumber                  int                                  `json:"sequenceNumber"`
	Resource                        Reference                            `json:"resource"`
	ResourceVersion                 int                                  `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers             `json:"resourceUserProvidedIdentifiers,omitempty"`
	LineItemId                      string                               `json:"lineItemId"`
	DiscountedPricePerQuantity      []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	// Draft type that stores amounts in cent precision for the specified currency.
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	TotalPrice Money           `json:"totalPrice"`
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderLineItemDiscountSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemDiscountSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

type OrderLineItemDistributionChannelSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	LineItemId                      string                   `json:"lineItemId"`
	// [Reference](/../api/types#reference) to a [Channel](ctp:api:type:Channel).
	DistributionChannel *ChannelReference `json:"distributionChannel,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderLineItemDistributionChannelSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderLineItemDistributionChannelSetMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderLineItemDistributionChannelSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemDistributionChannelSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLineItemDistributionChannelSet", Alias: (*Alias)(&obj)})
}

type OrderLineItemRemovedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	LineItemId                      string                   `json:"lineItemId"`
	RemovedQuantity                 int                      `json:"removedQuantity"`
	NewQuantity                     int                      `json:"newQuantity"`
	NewState                        []ItemState              `json:"newState"`
	// Base polymorphic read-only Money type which is stored in cent precision or high precision. The actual type is determined by the `type` field.
	NewTotalPrice     TypedMoney           `json:"newTotalPrice"`
	NewTaxedPrice     *TaxedItemPrice      `json:"newTaxedPrice,omitempty"`
	NewPrice          *Price               `json:"newPrice,omitempty"`
	NewShippingDetail *ItemShippingDetails `json:"newShippingDetail,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderLineItemRemovedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderLineItemRemovedMessage
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
	if obj.NewTotalPrice != nil {
		var err error
		obj.NewTotalPrice, err = mapDiscriminatorTypedMoney(obj.NewTotalPrice)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderLineItemRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLineItemRemoved", Alias: (*Alias)(&obj)})
}

type OrderPaymentAddedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Payment                         PaymentReference         `json:"payment"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderPaymentAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderPaymentAddedMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderPaymentAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderPaymentAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderPaymentAdded", Alias: (*Alias)(&obj)})
}

type OrderPaymentStateChangedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	PaymentState                    PaymentState             `json:"paymentState"`
	OldPaymentState                 *PaymentState            `json:"oldPaymentState,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderPaymentStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderPaymentStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderPaymentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderReturnInfoAddedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderReturnInfoAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnInfoAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReturnInfoAdded", Alias: (*Alias)(&obj)})
}

type OrderReturnInfoSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	ReturnInfo                      []ReturnInfo             `json:"returnInfo"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderReturnInfoSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderReturnInfoSetMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderReturnInfoSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnInfoSetMessage
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReturnInfoSet", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["returnInfo"] == nil {
		delete(target, "returnInfo")
	}

	return json.Marshal(target)
}

type OrderReturnShipmentStateChangedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderReturnShipmentStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnShipmentStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderReturnShipmentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderShipmentStateChangedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	ShipmentState                   ShipmentState            `json:"shipmentState"`
	OldShipmentState                *ShipmentState           `json:"oldShipmentState,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderShipmentStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShipmentStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShipmentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderShippingAddressSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Address                         *Address                 `json:"address,omitempty"`
	OldAddress                      *Address                 `json:"oldAddress,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderShippingAddressSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingAddressSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShippingAddressSet", Alias: (*Alias)(&obj)})
}

type OrderShippingInfoSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	ShippingInfo                    *ShippingInfo            `json:"shippingInfo,omitempty"`
	OldShippingInfo                 *ShippingInfo            `json:"oldShippingInfo,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderShippingInfoSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingInfoSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShippingInfoSet", Alias: (*Alias)(&obj)})
}

type OrderShippingRateInputSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	ShippingRateInput               ShippingRateInput        `json:"shippingRateInput,omitempty"`
	OldShippingRateInput            ShippingRateInput        `json:"oldShippingRateInput,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderShippingRateInputSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingRateInputSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShippingRateInputSet", Alias: (*Alias)(&obj)})
}

type OrderStateChangedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderStateChanged", Alias: (*Alias)(&obj)})
}

type OrderStateTransitionMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	State StateReference `json:"state"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	OldState *StateReference `json:"oldState,omitempty"`
	Force    bool            `json:"force"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderStateTransition", Alias: (*Alias)(&obj)})
}

type OrderStoreSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderStoreSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderStoreSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderStoreSet", Alias: (*Alias)(&obj)})
}

type ParcelAddedToDeliveryMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ParcelAddedToDeliveryMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelAddedToDeliveryMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelAddedToDelivery", Alias: (*Alias)(&obj)})
}

type ParcelItemsUpdatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	ParcelId                        string                   `json:"parcelId"`
	DeliveryId                      *string                  `json:"deliveryId,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ParcelItemsUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelItemsUpdatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelItemsUpdated", Alias: (*Alias)(&obj)})
}

type ParcelMeasurementsUpdatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	DeliveryId                      string                   `json:"deliveryId"`
	ParcelId                        string                   `json:"parcelId"`
	Measurements                    *ParcelMeasurements      `json:"measurements,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ParcelMeasurementsUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelMeasurementsUpdatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelMeasurementsUpdated", Alias: (*Alias)(&obj)})
}

type ParcelRemovedFromDeliveryMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ParcelRemovedFromDeliveryMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelRemovedFromDeliveryMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelRemovedFromDelivery", Alias: (*Alias)(&obj)})
}

type ParcelTrackingDataUpdatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	DeliveryId                      string                   `json:"deliveryId"`
	ParcelId                        string                   `json:"parcelId"`
	TrackingData                    *TrackingData            `json:"trackingData,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ParcelTrackingDataUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelTrackingDataUpdatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelTrackingDataUpdated", Alias: (*Alias)(&obj)})
}

type PaymentCreatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentCreated", Alias: (*Alias)(&obj)})
}

type PaymentInteractionAddedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Serves as value of the `custom` field on a resource or data type customized with a [Type](ctp:api:type:Type).
	Interaction CustomFields `json:"interaction"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentInteractionAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentInteractionAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentInteractionAdded", Alias: (*Alias)(&obj)})
}

type PaymentStatusInterfaceCodeSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentStatusInterfaceCodeSetMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusInterfaceCodeSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentStatusInterfaceCodeSet", Alias: (*Alias)(&obj)})
}

type PaymentStatusStateTransitionMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	State StateReference `json:"state"`
	Force bool           `json:"force"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentStatusStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentStatusStateTransition", Alias: (*Alias)(&obj)})
}

type PaymentTransactionAddedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentTransactionAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentTransactionAdded", Alias: (*Alias)(&obj)})
}

type PaymentTransactionStateChangedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentTransactionStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentTransactionStateChanged", Alias: (*Alias)(&obj)})
}

type ProductAddedToCategoryMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductAddedToCategoryMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductAddedToCategoryMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductAddedToCategory", Alias: (*Alias)(&obj)})
}

type ProductCreatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductCreated", Alias: (*Alias)(&obj)})
}

type ProductDeletedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductDeleted", Alias: (*Alias)(&obj)})
}

type ProductImageAddedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductImageAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductImageAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductImageAdded", Alias: (*Alias)(&obj)})
}

type ProductPriceDiscountsSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy                             `json:"createdBy,omitempty"`
	SequenceNumber                  int                                    `json:"sequenceNumber"`
	Resource                        Reference                              `json:"resource"`
	ResourceVersion                 int                                    `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers               `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductPriceDiscountsSetMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductPriceDiscountsSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductPriceDiscountsSet", Alias: (*Alias)(&obj)})
}

type ProductPriceDiscountsSetUpdatedPrice struct {
	VariantId  int              `json:"variantId"`
	VariantKey *string          `json:"variantKey,omitempty"`
	Sku        *string          `json:"sku,omitempty"`
	PriceId    string           `json:"priceId"`
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	Staged     bool             `json:"staged"`
}

type ProductPriceExternalDiscountSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	VariantId                       int                      `json:"variantId"`
	VariantKey                      *string                  `json:"variantKey,omitempty"`
	Sku                             *string                  `json:"sku,omitempty"`
	PriceId                         string                   `json:"priceId"`
	Discounted                      *DiscountedPrice         `json:"discounted,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductPriceExternalDiscountSetMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductPriceExternalDiscountSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductPriceExternalDiscountSet", Alias: (*Alias)(&obj)})
}

type ProductPublishedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductPublishedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductPublishedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductPublished", Alias: (*Alias)(&obj)})
}

type ProductRemovedFromCategoryMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductRemovedFromCategoryMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductRemovedFromCategoryMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductRemovedFromCategory", Alias: (*Alias)(&obj)})
}

type ProductRevertedStagedChangesMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductRevertedStagedChangesMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertedStagedChangesMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductRevertedStagedChanges", Alias: (*Alias)(&obj)})
}

type ProductSelectionCreatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	ProductSelection                ProductSelectionType     `json:"productSelection"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionCreatedMessage
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
	if obj.ProductSelection != nil {
		var err error
		obj.ProductSelection, err = mapDiscriminatorProductSelectionType(obj.ProductSelection)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSelectionCreated", Alias: (*Alias)(&obj)})
}

type ProductSelectionDeletedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Name                            LocalizedString          `json:"name"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionDeletedMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSelectionDeleted", Alias: (*Alias)(&obj)})
}

type ProductSelectionProductAddedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Product                         ProductReference         `json:"product"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionProductAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionProductAddedMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionProductAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionProductAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSelectionProductAdded", Alias: (*Alias)(&obj)})
}

type ProductSelectionProductRemovedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Product                         ProductReference         `json:"product"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionProductRemovedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionProductRemovedMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionProductRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionProductRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSelectionProductRemoved", Alias: (*Alias)(&obj)})
}

type ProductSlugChangedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Slug                            LocalizedString          `json:"slug"`
	OldSlug                         *LocalizedString         `json:"oldSlug,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSlugChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductSlugChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSlugChanged", Alias: (*Alias)(&obj)})
}

type ProductStateTransitionMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	State StateReference `json:"state"`
	Force bool           `json:"force"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductStateTransition", Alias: (*Alias)(&obj)})
}

type ProductUnpublishedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductUnpublishedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductUnpublishedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductUnpublished", Alias: (*Alias)(&obj)})
}

type ProductVariantAddedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductVariantAdded", Alias: (*Alias)(&obj)})
}

type ProductVariantDeletedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductVariantDeleted", Alias: (*Alias)(&obj)})
}

type ReviewCreatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ReviewCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewCreated", Alias: (*Alias)(&obj)})
}

type ReviewRatingSetMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	OldRating                       *float64                 `json:"oldRating,omitempty"`
	NewRating                       *float64                 `json:"newRating,omitempty"`
	IncludedInStatistics            bool                     `json:"includedInStatistics"`
	Target                          Reference                `json:"target,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewRatingSetMessage) MarshalJSON() ([]byte, error) {
	type Alias ReviewRatingSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewRatingSet", Alias: (*Alias)(&obj)})
}

type ReviewStateTransitionMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	OldState StateReference `json:"oldState"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	NewState                StateReference `json:"newState"`
	OldIncludedInStatistics bool           `json:"oldIncludedInStatistics"`
	NewIncludedInStatistics bool           `json:"newIncludedInStatistics"`
	Target                  Reference      `json:"target"`
	Force                   bool           `json:"force"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias ReviewStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewStateTransition", Alias: (*Alias)(&obj)})
}

type StoreCreatedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy                `json:"createdBy,omitempty"`
	SequenceNumber                  int                       `json:"sequenceNumber"`
	Resource                        Reference                 `json:"resource"`
	ResourceVersion                 int                       `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers  `json:"resourceUserProvidedIdentifiers,omitempty"`
	Name                            *LocalizedString          `json:"name,omitempty"`
	Languages                       []string                  `json:"languages"`
	DistributionChannels            []ChannelReference        `json:"distributionChannels"`
	SupplyChannels                  []ChannelReference        `json:"supplyChannels"`
	ProductSelections               []ProductSelectionSetting `json:"productSelections"`
	// Serves as value of the `custom` field on a resource or data type customized with a [Type](ctp:api:type:Type).
	Custom *CustomFields `json:"custom,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias StoreCreatedMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias StoreCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreCreated", Alias: (*Alias)(&obj)})
}

type StoreDeletedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	Resource                        Reference                `json:"resource"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias StoreDeletedMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias StoreDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreDeleted", Alias: (*Alias)(&obj)})
}

type StoreProductSelectionsChangedMessage struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy                       *CreatedBy                `json:"createdBy,omitempty"`
	SequenceNumber                  int                       `json:"sequenceNumber"`
	Resource                        Reference                 `json:"resource"`
	ResourceVersion                 int                       `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers  `json:"resourceUserProvidedIdentifiers,omitempty"`
	AddedProductSelections          []ProductSelectionSetting `json:"addedProductSelections"`
	RemovedProductSelections        []ProductSelectionSetting `json:"removedProductSelections"`
	UpdatedProductSelections        []ProductSelectionSetting `json:"updatedProductSelections"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreProductSelectionsChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias StoreProductSelectionsChangedMessage
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreProductSelectionsChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias StoreProductSelectionsChangedMessage
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreProductSelectionsChanged", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["addedProductSelections"] == nil {
		delete(target, "addedProductSelections")
	}

	if target["removedProductSelections"] == nil {
		delete(target, "removedProductSelections")
	}

	if target["updatedProductSelections"] == nil {
		delete(target, "updatedProductSelections")
	}

	return json.Marshal(target)
}

type UserProvidedIdentifiers struct {
	Key            *string          `json:"key,omitempty"`
	ExternalId     *string          `json:"externalId,omitempty"`
	OrderNumber    *string          `json:"orderNumber,omitempty"`
	CustomerNumber *string          `json:"customerNumber,omitempty"`
	Sku            *string          `json:"sku,omitempty"`
	Slug           *LocalizedString `json:"slug,omitempty"`
	// Custom Objects are grouped into containers, which can be used like namespaces. Within a given container, a user-defined key can be used to uniquely identify resources.
	ContainerAndKey *ContainerAndKey `json:"containerAndKey,omitempty"`
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
		obj := CategoryCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CategorySlugChanged":
		obj := CategorySlugChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerAddressAdded":
		obj := CustomerAddressAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerAddressChanged":
		obj := CustomerAddressChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerAddressRemoved":
		obj := CustomerAddressRemovedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerCompanyNameSet":
		obj := CustomerCompanyNameSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerCreated":
		obj := CustomerCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerDateOfBirthSet":
		obj := CustomerDateOfBirthSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerDeleted":
		obj := CustomerDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerEmailChanged":
		obj := CustomerEmailChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerEmailVerified":
		obj := CustomerEmailVerifiedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerFirstNameSet":
		obj := CustomerFirstNameSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerGroupSet":
		obj := CustomerGroupSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerLastNameSet":
		obj := CustomerLastNameSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerPasswordUpdated":
		obj := CustomerPasswordUpdatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CustomerTitleSet":
		obj := CustomerTitleSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InventoryEntryCreated":
		obj := InventoryEntryCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InventoryEntryDeleted":
		obj := InventoryEntryDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "InventoryEntryQuantitySet":
		obj := InventoryEntryQuantitySetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderPaymentAdded":
		obj := OrderPaymentAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PaymentCreated":
		obj := PaymentCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PaymentInteractionAdded":
		obj := PaymentInteractionAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PaymentStatusInterfaceCodeSet":
		obj := PaymentStatusInterfaceCodeSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PaymentStatusStateTransition":
		obj := PaymentStatusStateTransitionMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PaymentTransactionAdded":
		obj := PaymentTransactionAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PaymentTransactionStateChanged":
		obj := PaymentTransactionStateChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductAddedToCategory":
		obj := ProductAddedToCategoryMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductCreated":
		obj := ProductCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductDeleted":
		obj := ProductDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductImageAdded":
		obj := ProductImageAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductPriceDiscountsSet":
		obj := ProductPriceDiscountsSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductPriceExternalDiscountSet":
		obj := ProductPriceExternalDiscountSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductPublished":
		obj := ProductPublishedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductRemovedFromCategory":
		obj := ProductRemovedFromCategoryMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductRevertedStagedChanges":
		obj := ProductRevertedStagedChangesMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductSelectionCreated":
		obj := ProductSelectionCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.ProductSelection != nil {
			var err error
			obj.ProductSelection, err = mapDiscriminatorProductSelectionType(obj.ProductSelection)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductSelectionDeleted":
		obj := ProductSelectionDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductSelectionProductAdded":
		obj := ProductSelectionProductAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductSelectionProductRemoved":
		obj := ProductSelectionProductRemovedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductSlugChanged":
		obj := ProductSlugChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductStateTransition":
		obj := ProductStateTransitionMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductUnpublished":
		obj := ProductUnpublishedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductVariantAdded":
		obj := ProductVariantAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductVariantDeleted":
		obj := ProductVariantDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ReviewCreated":
		obj := ReviewCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ReviewRatingSet":
		obj := ReviewRatingSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Target != nil {
			var err error
			obj.Target, err = mapDiscriminatorReference(obj.Target)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ReviewStateTransition":
		obj := ReviewStateTransitionMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Target != nil {
			var err error
			obj.Target, err = mapDiscriminatorReference(obj.Target)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ShoppingListStoreSet":
		obj := ShoppingListStoreSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StoreCreated":
		obj := StoreCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StoreDeleted":
		obj := StoreDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StoreProductSelectionsChanged":
		obj := StoreProductSelectionsChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CategoryCreatedMessagePayload struct {
	Category Category `json:"category"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CategoryCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CategoryCreated", Alias: (*Alias)(&obj)})
}

type CategorySlugChangedMessagePayload struct {
	Slug    LocalizedString  `json:"slug"`
	OldSlug *LocalizedString `json:"oldSlug,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySlugChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CategorySlugChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CategorySlugChanged", Alias: (*Alias)(&obj)})
}

type CustomerAddressAddedMessagePayload struct {
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerAddressRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerAddressRemoved", Alias: (*Alias)(&obj)})
}

type CustomerCompanyNameSetMessagePayload struct {
	CompanyName *string `json:"companyName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerCreated", Alias: (*Alias)(&obj)})
}

type CustomerDateOfBirthSetMessagePayload struct {
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerDateOfBirthSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerDateOfBirthSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerDateOfBirthSet", Alias: (*Alias)(&obj)})
}

type CustomerDeletedMessagePayload struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerDeleted", Alias: (*Alias)(&obj)})
}

type CustomerEmailChangedMessagePayload struct {
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerEmailChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerEmailChanged", Alias: (*Alias)(&obj)})
}

type CustomerEmailVerifiedMessagePayload struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerEmailVerifiedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailVerifiedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerEmailVerified", Alias: (*Alias)(&obj)})
}

type CustomerFirstNameSetMessagePayload struct {
	FirstName *string `json:"firstName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerFirstNameSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerFirstNameSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerFirstNameSet", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetMessagePayload struct {
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerGroupSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerGroupSet", Alias: (*Alias)(&obj)})
}

type CustomerLastNameSetMessagePayload struct {
	LastName *string `json:"lastName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerLastNameSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerLastNameSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerLastNameSet", Alias: (*Alias)(&obj)})
}

type CustomerPasswordUpdatedMessagePayload struct {
	// true, if password has been updated during Customer's Password Reset workflow.
	Reset bool `json:"reset"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerPasswordUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerPasswordUpdatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerPasswordUpdated", Alias: (*Alias)(&obj)})
}

type CustomerTitleSetMessagePayload struct {
	Title *string `json:"title,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerTitleSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerTitleSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomerTitleSet", Alias: (*Alias)(&obj)})
}

type InventoryEntryCreatedMessagePayload struct {
	InventoryEntry InventoryEntry `json:"inventoryEntry"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "InventoryEntryCreated", Alias: (*Alias)(&obj)})
}

type InventoryEntryDeletedMessagePayload struct {
	Sku string `json:"sku"`
	// [Reference](/../api/types#reference) to a [Channel](ctp:api:type:Channel).
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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
	// [Reference](/../api/types#reference) to a [Channel](ctp:api:type:Channel).
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryQuantitySetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryQuantitySetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "InventoryEntryQuantitySet", Alias: (*Alias)(&obj)})
}

type OrderMessagePayload interface{}

func mapDiscriminatorOrderMessagePayload(input interface{}) (OrderMessagePayload, error) {

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
	case "CustomLineItemStateTransition":
		obj := CustomLineItemStateTransitionMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DeliveryAdded":
		obj := DeliveryAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DeliveryAddressSet":
		obj := DeliveryAddressSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DeliveryItemsUpdated":
		obj := DeliveryItemsUpdatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DeliveryRemoved":
		obj := DeliveryRemovedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "LineItemStateTransition":
		obj := LineItemStateTransitionMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderBillingAddressSet":
		obj := OrderBillingAddressSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderCreated":
		obj := OrderCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderCustomLineItemDiscountSet":
		obj := OrderCustomLineItemDiscountSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderCustomerEmailSet":
		obj := OrderCustomerEmailSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderCustomerGroupSet":
		obj := OrderCustomerGroupSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderCustomerSet":
		obj := OrderCustomerSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderDeleted":
		obj := OrderDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderDiscountCodeAdded":
		obj := OrderDiscountCodeAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderDiscountCodeRemoved":
		obj := OrderDiscountCodeRemovedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderDiscountCodeStateSet":
		obj := OrderDiscountCodeStateSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderEditApplied":
		obj := OrderEditAppliedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderImported":
		obj := OrderImportedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderLineItemAdded":
		obj := OrderLineItemAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderLineItemDiscountSet":
		obj := OrderLineItemDiscountSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderLineItemDistributionChannelSet":
		obj := OrderLineItemDistributionChannelSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderLineItemRemoved":
		obj := OrderLineItemRemovedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.NewTotalPrice != nil {
			var err error
			obj.NewTotalPrice, err = mapDiscriminatorTypedMoney(obj.NewTotalPrice)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderPaymentStateChanged":
		obj := OrderPaymentStateChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ReturnInfoAdded":
		obj := OrderReturnInfoAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ReturnInfoSet":
		obj := OrderReturnInfoSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderReturnShipmentStateChanged":
		obj := OrderReturnShipmentStateChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderShipmentStateChanged":
		obj := OrderShipmentStateChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderShippingAddressSet":
		obj := OrderShippingAddressSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderShippingInfoSet":
		obj := OrderShippingInfoSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderShippingRateInputSet":
		obj := OrderShippingRateInputSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.ShippingRateInput != nil {
			var err error
			obj.ShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.ShippingRateInput)
			if err != nil {
				return nil, err
			}
		}
		if obj.OldShippingRateInput != nil {
			var err error
			obj.OldShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.OldShippingRateInput)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "OrderStateChanged":
		obj := OrderStateChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderStateTransition":
		obj := OrderStateTransitionMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderStoreSet":
		obj := OrderStoreSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ParcelAddedToDelivery":
		obj := ParcelAddedToDeliveryMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ParcelItemsUpdated":
		obj := ParcelItemsUpdatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ParcelMeasurementsUpdated":
		obj := ParcelMeasurementsUpdatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ParcelRemovedFromDelivery":
		obj := ParcelRemovedFromDeliveryMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ParcelTrackingDataUpdated":
		obj := ParcelTrackingDataUpdatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CustomLineItemStateTransitionMessagePayload struct {
	CustomLineItemId string    `json:"customLineItemId"`
	TransitionDate   time.Time `json:"transitionDate"`
	Quantity         int       `json:"quantity"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	FromState StateReference `json:"fromState"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	ToState StateReference `json:"toState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomLineItemStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomLineItemStateTransition", Alias: (*Alias)(&obj)})
}

type DeliveryAddedMessagePayload struct {
	Delivery Delivery `json:"delivery"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeliveryAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryAdded", Alias: (*Alias)(&obj)})
}

type DeliveryAddressSetMessagePayload struct {
	DeliveryId string   `json:"deliveryId"`
	Address    *Address `json:"address,omitempty"`
	OldAddress *Address `json:"oldAddress,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeliveryRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DeliveryRemoved", Alias: (*Alias)(&obj)})
}

type LineItemStateTransitionMessagePayload struct {
	LineItemId     string    `json:"lineItemId"`
	TransitionDate time.Time `json:"transitionDate"`
	Quantity       int       `json:"quantity"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	FromState StateReference `json:"fromState"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	ToState StateReference `json:"toState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LineItemStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias LineItemStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "LineItemStateTransition", Alias: (*Alias)(&obj)})
}

type OrderBillingAddressSetMessagePayload struct {
	Address    *Address `json:"address,omitempty"`
	OldAddress *Address `json:"oldAddress,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCustomLineItemDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemDiscountSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerEmailSetMessagePayload struct {
	Email    *string `json:"email,omitempty"`
	OldEmail *string `json:"oldEmail,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCustomerEmailSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerEmailSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomerEmailSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerGroupSetMessagePayload struct {
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	OldCustomerGroup *CustomerGroupReference `json:"oldCustomerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCustomerGroupSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerGroupSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomerGroupSet", Alias: (*Alias)(&obj)})
}

type OrderCustomerSetMessagePayload struct {
	Customer *CustomerReference `json:"customer,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	OldCustomer   *CustomerReference      `json:"oldCustomer,omitempty"`
	// [Reference](/types#reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	OldCustomerGroup *CustomerGroupReference `json:"oldCustomerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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
	OldState     *DiscountCodeState    `json:"oldState,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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
	// Draft type that stores amounts in cent precision for the specified currency.
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	TotalPrice Money           `json:"totalPrice"`
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderLineItemDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemDiscountSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

type OrderLineItemDistributionChannelSetMessagePayload struct {
	LineItemId string `json:"lineItemId"`
	// [Reference](/../api/types#reference) to a [Channel](ctp:api:type:Channel).
	DistributionChannel *ChannelReference `json:"distributionChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderLineItemDistributionChannelSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemDistributionChannelSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLineItemDistributionChannelSet", Alias: (*Alias)(&obj)})
}

type OrderLineItemRemovedMessagePayload struct {
	LineItemId      string      `json:"lineItemId"`
	RemovedQuantity int         `json:"removedQuantity"`
	NewQuantity     int         `json:"newQuantity"`
	NewState        []ItemState `json:"newState"`
	// Base polymorphic read-only Money type which is stored in cent precision or high precision. The actual type is determined by the `type` field.
	NewTotalPrice     TypedMoney           `json:"newTotalPrice"`
	NewTaxedPrice     *TaxedItemPrice      `json:"newTaxedPrice,omitempty"`
	NewPrice          *Price               `json:"newPrice,omitempty"`
	NewShippingDetail *ItemShippingDetails `json:"newShippingDetail,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderLineItemRemovedMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias OrderLineItemRemovedMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.NewTotalPrice != nil {
		var err error
		obj.NewTotalPrice, err = mapDiscriminatorTypedMoney(obj.NewTotalPrice)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderLineItemRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderLineItemRemoved", Alias: (*Alias)(&obj)})
}

type OrderPaymentAddedMessagePayload struct {
	Payment PaymentReference `json:"payment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderPaymentAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderPaymentAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderPaymentAdded", Alias: (*Alias)(&obj)})
}

type OrderPaymentStateChangedMessagePayload struct {
	PaymentState    PaymentState  `json:"paymentState"`
	OldPaymentState *PaymentState `json:"oldPaymentState,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderReturnInfoAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnInfoAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReturnInfoAdded", Alias: (*Alias)(&obj)})
}

type OrderReturnInfoSetMessagePayload struct {
	ReturnInfo []ReturnInfo `json:"returnInfo"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderReturnInfoSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnInfoSetMessagePayload
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReturnInfoSet", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["returnInfo"] == nil {
		delete(target, "returnInfo")
	}

	return json.Marshal(target)
}

type OrderReturnShipmentStateChangedMessagePayload struct {
	ReturnItemId        string              `json:"returnItemId"`
	ReturnShipmentState ReturnShipmentState `json:"returnShipmentState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderReturnShipmentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnShipmentStateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderReturnShipmentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderShipmentStateChangedMessagePayload struct {
	ShipmentState    ShipmentState  `json:"shipmentState"`
	OldShipmentState *ShipmentState `json:"oldShipmentState,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderShipmentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShipmentStateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShipmentStateChanged", Alias: (*Alias)(&obj)})
}

type OrderShippingAddressSetMessagePayload struct {
	Address    *Address `json:"address,omitempty"`
	OldAddress *Address `json:"oldAddress,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderShippingAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingAddressSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShippingAddressSet", Alias: (*Alias)(&obj)})
}

type OrderShippingInfoSetMessagePayload struct {
	ShippingInfo    *ShippingInfo `json:"shippingInfo,omitempty"`
	OldShippingInfo *ShippingInfo `json:"oldShippingInfo,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderShippingInfoSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingInfoSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderShippingInfoSet", Alias: (*Alias)(&obj)})
}

type OrderShippingRateInputSetMessagePayload struct {
	ShippingRateInput    ShippingRateInput `json:"shippingRateInput,omitempty"`
	OldShippingRateInput ShippingRateInput `json:"oldShippingRateInput,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderStateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderStateChanged", Alias: (*Alias)(&obj)})
}

type OrderStateTransitionMessagePayload struct {
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	State StateReference `json:"state"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	OldState *StateReference `json:"oldState,omitempty"`
	Force    bool            `json:"force"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ParcelAddedToDeliveryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelAddedToDeliveryMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ParcelAddedToDelivery", Alias: (*Alias)(&obj)})
}

type ParcelItemsUpdatedMessagePayload struct {
	ParcelId   string         `json:"parcelId"`
	DeliveryId *string        `json:"deliveryId,omitempty"`
	Items      []DeliveryItem `json:"items"`
	OldItems   []DeliveryItem `json:"oldItems"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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
	TrackingData *TrackingData `json:"trackingData,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentCreated", Alias: (*Alias)(&obj)})
}

type PaymentInteractionAddedMessagePayload struct {
	// Serves as value of the `custom` field on a resource or data type customized with a [Type](ctp:api:type:Type).
	Interaction CustomFields `json:"interaction"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentStatusInterfaceCodeSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusInterfaceCodeSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PaymentStatusInterfaceCodeSet", Alias: (*Alias)(&obj)})
}

type PaymentStatusStateTransitionMessagePayload struct {
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	State StateReference `json:"state"`
	Force bool           `json:"force"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductPriceDiscountsSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductPriceDiscountsSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductPriceDiscountsSet", Alias: (*Alias)(&obj)})
}

type ProductPriceExternalDiscountSetMessagePayload struct {
	VariantId  int              `json:"variantId"`
	VariantKey *string          `json:"variantKey,omitempty"`
	Sku        *string          `json:"sku,omitempty"`
	PriceId    string           `json:"priceId"`
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	Staged     bool             `json:"staged"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductRevertedStagedChangesMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertedStagedChangesMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductRevertedStagedChanges", Alias: (*Alias)(&obj)})
}

type ProductSelectionCreatedMessagePayload struct {
	ProductSelection ProductSelectionType `json:"productSelection"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionCreatedMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionCreatedMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ProductSelection != nil {
		var err error
		obj.ProductSelection, err = mapDiscriminatorProductSelectionType(obj.ProductSelection)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSelectionCreated", Alias: (*Alias)(&obj)})
}

type ProductSelectionDeletedMessagePayload struct {
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSelectionDeleted", Alias: (*Alias)(&obj)})
}

type ProductSelectionProductAddedMessagePayload struct {
	Product ProductReference `json:"product"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionProductAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionProductAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSelectionProductAdded", Alias: (*Alias)(&obj)})
}

type ProductSelectionProductRemovedMessagePayload struct {
	Product ProductReference `json:"product"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionProductRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionProductRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSelectionProductRemoved", Alias: (*Alias)(&obj)})
}

type ProductSlugChangedMessagePayload struct {
	Slug    LocalizedString  `json:"slug"`
	OldSlug *LocalizedString `json:"oldSlug,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSlugChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductSlugChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSlugChanged", Alias: (*Alias)(&obj)})
}

type ProductStateTransitionMessagePayload struct {
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	State StateReference `json:"state"`
	Force bool           `json:"force"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductStateTransition", Alias: (*Alias)(&obj)})
}

type ProductUnpublishedMessagePayload struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewCreated", Alias: (*Alias)(&obj)})
}

type ReviewRatingSetMessagePayload struct {
	OldRating            *float64  `json:"oldRating,omitempty"`
	NewRating            *float64  `json:"newRating,omitempty"`
	IncludedInStatistics bool      `json:"includedInStatistics"`
	Target               Reference `json:"target,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewRatingSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewRatingSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReviewRatingSet", Alias: (*Alias)(&obj)})
}

type ReviewStateTransitionMessagePayload struct {
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
	OldState StateReference `json:"oldState"`
	// [Reference](/../api/types#reference) to a [State](ctp:api:type:State).
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShoppingListStoreSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListStoreSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ShoppingListStoreSet", Alias: (*Alias)(&obj)})
}

type StoreCreatedMessagePayload struct {
	Name                 *LocalizedString          `json:"name,omitempty"`
	Languages            []string                  `json:"languages"`
	DistributionChannels []ChannelReference        `json:"distributionChannels"`
	SupplyChannels       []ChannelReference        `json:"supplyChannels"`
	ProductSelections    []ProductSelectionSetting `json:"productSelections"`
	// Serves as value of the `custom` field on a resource or data type customized with a [Type](ctp:api:type:Type).
	Custom *CustomFields `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StoreCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreCreated", Alias: (*Alias)(&obj)})
}

type StoreDeletedMessagePayload struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StoreDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreDeleted", Alias: (*Alias)(&obj)})
}

type StoreProductSelectionsChangedMessagePayload struct {
	AddedProductSelections   []ProductSelectionSetting `json:"addedProductSelections"`
	RemovedProductSelections []ProductSelectionSetting `json:"removedProductSelections"`
	UpdatedProductSelections []ProductSelectionSetting `json:"updatedProductSelections"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreProductSelectionsChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StoreProductSelectionsChangedMessagePayload
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreProductSelectionsChanged", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["addedProductSelections"] == nil {
		delete(target, "addedProductSelections")
	}

	if target["removedProductSelections"] == nil {
		delete(target, "removedProductSelections")
	}

	if target["updatedProductSelections"] == nil {
		delete(target, "updatedProductSelections")
	}

	return json.Marshal(target)
}
