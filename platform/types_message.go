package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	`ContainerAndKey` is specific to [Custom Objects](ctp:api:type:CustomObject). Custom Objects are grouped into containers, which can be used like namespaces. Within a given container, a user-defined key can be used to uniquely identify resources.
*
 */
type ContainerAndKey struct {
	// User-defined identifier that is unique within the given container.
	Key string `json:"key"`
	// Namespace to group [Custom Objects](ctp:api:type:CustomObject).
	Container string `json:"container"`
}

/**
*	Base representation of a Message containing common fields to all [Message Types](/../api/projects/messages#message-types).
*
 */
type Message interface{}

func mapDiscriminatorMessage(input interface{}) (Message, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "BusinessUnitAddressAdded":
		obj := BusinessUnitAddressAddedMessage{}
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
	case "BusinessUnitAddressChanged":
		obj := BusinessUnitAddressChangedMessage{}
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
	case "BusinessUnitAddressRemoved":
		obj := BusinessUnitAddressRemovedMessage{}
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
	case "BusinessUnitAssociateAdded":
		obj := BusinessUnitAssociateAddedMessage{}
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
	case "BusinessUnitAssociateChanged":
		obj := BusinessUnitAssociateChangedMessage{}
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
	case "BusinessUnitAssociateRemoved":
		obj := BusinessUnitAssociateRemovedMessage{}
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
	case "BusinessUnitAssociatesSet":
		obj := BusinessUnitAssociatesSetMessage{}
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
	case "BusinessUnitBillingAddressAdded":
		obj := BusinessUnitBillingAddressAddedMessage{}
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
	case "BusinessUnitBillingAddressRemoved":
		obj := BusinessUnitBillingAddressRemovedMessage{}
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
	case "BusinessUnitContactEmailSet":
		obj := BusinessUnitContactEmailSetMessage{}
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
	case "BusinessUnitCreated":
		obj := BusinessUnitCreatedMessage{}
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
		if obj.BusinessUnit != nil {
			var err error
			obj.BusinessUnit, err = mapDiscriminatorBusinessUnit(obj.BusinessUnit)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "BusinessUnitDefaultBillingAddressSet":
		obj := BusinessUnitDefaultBillingAddressSetMessage{}
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
	case "BusinessUnitDefaultShippingAddressSet":
		obj := BusinessUnitDefaultShippingAddressSetMessage{}
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
	case "BusinessUnitDeleted":
		obj := BusinessUnitDeletedMessage{}
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
	case "BusinessUnitNameChanged":
		obj := BusinessUnitNameChangedMessage{}
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
	case "BusinessUnitParentUnitChanged":
		obj := BusinessUnitParentUnitChangedMessage{}
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
	case "BusinessUnitShippingAddressAdded":
		obj := BusinessUnitShippingAddressAddedMessage{}
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
	case "BusinessUnitShippingAddressRemoved":
		obj := BusinessUnitShippingAddressRemovedMessage{}
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
	case "BusinessUnitStatusChanged":
		obj := BusinessUnitStatusChangedMessage{}
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
	case "BusinessUnitStoreAdded":
		obj := BusinessUnitStoreAddedMessage{}
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
	case "BusinessUnitStoreModeChanged":
		obj := BusinessUnitStoreModeChangedMessage{}
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
	case "BusinessUnitStoreRemoved":
		obj := BusinessUnitStoreRemovedMessage{}
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
	case "BusinessUnitStoresSet":
		obj := BusinessUnitStoresSetMessage{}
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
		if obj.VariantSelection != nil {
			var err error
			obj.VariantSelection, err = mapDiscriminatorProductVariantSelection(obj.VariantSelection)
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
	case "ProductSelectionVariantSelectionChanged":
		obj := ProductSelectionVariantSelectionChangedMessage{}
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
		if obj.OldVariantSelection != nil {
			var err error
			obj.OldVariantSelection, err = mapDiscriminatorProductVariantSelection(obj.OldVariantSelection)
			if err != nil {
				return nil, err
			}
		}
		if obj.NewVariantSelection != nil {
			var err error
			obj.NewVariantSelection, err = mapDiscriminatorProductVariantSelection(obj.NewVariantSelection)
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
	case "QuoteCreated":
		obj := QuoteCreatedMessage{}
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
	case "QuoteDeleted":
		obj := QuoteDeletedMessage{}
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
	case "QuoteRequestCreated":
		obj := QuoteRequestCreatedMessage{}
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
	case "QuoteRequestDeleted":
		obj := QuoteRequestDeletedMessage{}
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
	case "QuoteRequestStateChanged":
		obj := QuoteRequestStateChangedMessage{}
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
	case "QuoteRequestStateTransition":
		obj := QuoteRequestStateTransitionMessage{}
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
	case "QuoteStateChanged":
		obj := QuoteStateChangedMessage{}
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
	case "QuoteStateTransition":
		obj := QuoteStateTransitionMessage{}
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
	case "StagedQuoteCreated":
		obj := StagedQuoteCreatedMessage{}
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
	case "StagedQuoteDeleted":
		obj := StagedQuoteDeletedMessage{}
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
	case "StagedQuoteSellerCommentSet":
		obj := StagedQuoteSellerCommentSetMessage{}
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
	case "StagedQuoteStateChanged":
		obj := StagedQuoteStateChangedMessage{}
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
	case "StagedQuoteStateTransition":
		obj := StagedQuoteStateTransitionMessage{}
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
	case "StagedQuoteValidToSet":
		obj := StagedQuoteValidToSetMessage{}
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
	case "StandalonePriceActiveChanged":
		obj := StandalonePriceActiveChangedMessage{}
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
	case "StandalonePriceCreated":
		obj := StandalonePriceCreatedMessage{}
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
	case "StandalonePriceDeleted":
		obj := StandalonePriceDeletedMessage{}
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
	case "StandalonePriceDiscountSet":
		obj := StandalonePriceDiscountSetMessage{}
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
	case "StandalonePriceExternalDiscountSet":
		obj := StandalonePriceExternalDiscountSetMessage{}
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
	case "StandalonePriceStagedChangesApplied":
		obj := StandalonePriceStagedChangesAppliedMessage{}
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
	case "StandalonePriceValueChanged":
		obj := StandalonePriceValueChangedMessage{}
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
	case "StoreDistributionChannelsChanged":
		obj := StoreDistributionChannelsChangedMessage{}
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
	case "StoreLanguagesChanged":
		obj := StoreLanguagesChangedMessage{}
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
	case "StoreNameSet":
		obj := StoreNameSetMessage{}
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
	case "StoreSupplyChannelsChanged":
		obj := StoreSupplyChannelsChangedMessage{}
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

/**
*	Generated after a successful [Add Address](ctp:api:type:BusinessUnitAddAddressAction) update action.
*
 */
type BusinessUnitAddressAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The address that was added to the [Business Unit](ctp:api:type:BusinessUnit).
	Address Address `json:"address"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitAddressAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitAddressAddedMessage
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
func (obj BusinessUnitAddressAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAddressAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAddressAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Address](ctp:api:type:BusinessUnitChangeAddressAction) update action.
*
 */
type BusinessUnitAddressChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Updated address of the Business Unit.
	Address Address `json:"address"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitAddressChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitAddressChangedMessage
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
func (obj BusinessUnitAddressChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAddressChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAddressChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Remove Address](ctp:api:type:BusinessUnitRemoveAddressAction) update action.
*
 */
type BusinessUnitAddressRemovedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The address that was removed from the [Business Unit](ctp:api:type:BusinessUnit).
	Address Address `json:"address"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitAddressRemovedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitAddressRemovedMessage
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
func (obj BusinessUnitAddressRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAddressRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAddressRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Add Associate](ctp:api:type:BusinessUnitAddAssociateAction) update action.
*
 */
type BusinessUnitAssociateAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The [Associate](ctp:api:type:Associate) that was added to the [Business Unit](ctp:api:type:BusinessUnit).
	Associate Associate `json:"associate"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitAssociateAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitAssociateAddedMessage
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
func (obj BusinessUnitAssociateAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAssociateAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAssociateAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Associate](ctp:api:type:BusinessUnitChangeAssociateAction) update action.
*
 */
type BusinessUnitAssociateChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The [Associate](ctp:api:type:Associate) that was updated.
	Associate Associate `json:"associate"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitAssociateChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitAssociateChangedMessage
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
func (obj BusinessUnitAssociateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAssociateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAssociateChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Remove Associate](ctp:api:type:BusinessUnitRemoveAssociateAction) update action.
*
 */
type BusinessUnitAssociateRemovedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The [Associate](ctp:api:type:Associate) that was removed from the [Business Unit](ctp:api:type:BusinessUnit).
	Associate Associate `json:"associate"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitAssociateRemovedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitAssociateRemovedMessage
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
func (obj BusinessUnitAssociateRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAssociateRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAssociateRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Associates](ctp:api:type:BusinessUnitSetAssociatesAction) update action.
*
 */
type BusinessUnitAssociatesSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The list of [Associates](ctp:api:type:Associate) that was updated on the [Business Unit](ctp:api:type:BusinessUnit).
	Associates []Associate `json:"associates"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitAssociatesSetMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitAssociatesSetMessage
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
func (obj BusinessUnitAssociatesSetMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAssociatesSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAssociatesSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Add Billing Address Identifier](ctp:api:type:BusinessUnitAddBillingAddressIdAction) update action.
*
 */
type BusinessUnitBillingAddressAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The address that was added to the [Business Unit](ctp:api:type:BusinessUnit) as billing address.
	Address Address `json:"address"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitBillingAddressAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitBillingAddressAddedMessage
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
func (obj BusinessUnitBillingAddressAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitBillingAddressAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitBillingAddressAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Remove Billing Address Identifier](ctp:api:type:BusinessUnitRemoveBillingAddressIdAction) update action.
*
 */
type BusinessUnitBillingAddressRemovedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The address that was removed from the billing addresses of the [Business Unit](ctp:api:type:BusinessUnit).
	Address Address `json:"address"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitBillingAddressRemovedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitBillingAddressRemovedMessage
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
func (obj BusinessUnitBillingAddressRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitBillingAddressRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitBillingAddressRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Contact Email](ctp:api:type:BusinessUnitSetContactEmailAction) update action.
*
 */
type BusinessUnitContactEmailSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The contact email that was updated on the [Business Unit](ctp:api:type:BusinessUnit).
	ContactEmail *string `json:"contactEmail,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitContactEmailSetMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitContactEmailSetMessage
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
func (obj BusinessUnitContactEmailSetMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitContactEmailSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitContactEmailSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Create Business Unit](/projects/business-units#create-businessunit) request.
*
 */
type BusinessUnitCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The [Business Unit](ctp:api:type:BusinessUnit) that was created.
	BusinessUnit BusinessUnit `json:"businessUnit"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitCreatedMessage
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
	if obj.BusinessUnit != nil {
		var err error
		obj.BusinessUnit, err = mapDiscriminatorBusinessUnit(obj.BusinessUnit)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitCreated", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Default Billing Address](ctp:api:type:BusinessUnitSetDefaultBillingAddressAction) update action.
*
 */
type BusinessUnitDefaultBillingAddressSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The address that was set as the default billing address.
	Address *Address `json:"address,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitDefaultBillingAddressSetMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitDefaultBillingAddressSetMessage
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
func (obj BusinessUnitDefaultBillingAddressSetMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitDefaultBillingAddressSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitDefaultBillingAddressSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Default Shipping Address](ctp:api:type:BusinessUnitSetDefaultShippingAddressAction) update action.
*
 */
type BusinessUnitDefaultShippingAddressSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The address that was set as the default shipping address.
	Address *Address `json:"address,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitDefaultShippingAddressSetMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitDefaultShippingAddressSetMessage
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
func (obj BusinessUnitDefaultShippingAddressSetMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitDefaultShippingAddressSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitDefaultShippingAddressSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Delete Business Unit](/projects/business-units#delete-businessunit) request.
*
 */
type BusinessUnitDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitDeletedMessage
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
func (obj BusinessUnitDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitDeleted", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Name](ctp:api:type:BusinessUnitChangeNameAction) update action.
*
 */
type BusinessUnitNameChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Updated name of the [Business Unit](ctp:api:type:BusinessUnit).
	Name string `json:"name"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitNameChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitNameChangedMessage
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
func (obj BusinessUnitNameChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitNameChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitNameChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Parent Unit](ctp:api:type:BusinessUnitChangeParentUnitAction) update action.
*
 */
type BusinessUnitParentUnitChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Parent unit of the [Business Unit](ctp:api:type:BusinessUnit) before the [Change Parent Unit](ctp:api:type:BusinessUnitChangeParentUnitAction) update action.
	OldParentUnit *BusinessUnitKeyReference `json:"oldParentUnit,omitempty"`
	// Parent unit of the [Business Unit](ctp:api:type:BusinessUnit) after the [Change Parent Unit](ctp:api:type:BusinessUnitChangeParentUnitAction) update action.
	NewParentUnit *BusinessUnitKeyReference `json:"newParentUnit,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitParentUnitChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitParentUnitChangedMessage
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
func (obj BusinessUnitParentUnitChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitParentUnitChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitParentUnitChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Add Shipping Address Identifier](ctp:api:type:BusinessUnitAddShippingAddressIdAction) update action.
*
 */
type BusinessUnitShippingAddressAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The address that was added to the [Business Unit](ctp:api:type:BusinessUnit) as shipping address.
	Address Address `json:"address"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitShippingAddressAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitShippingAddressAddedMessage
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
func (obj BusinessUnitShippingAddressAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitShippingAddressAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitShippingAddressAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Remove Shipping Address Identifier](ctp:api:type:BusinessUnitRemoveShippingAddressIdAction) update action.
*
 */
type BusinessUnitShippingAddressRemovedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The address that was removed from shipping addresses of the [Business Unit](ctp:api:type:BusinessUnit).
	Address Address `json:"address"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitShippingAddressRemovedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitShippingAddressRemovedMessage
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
func (obj BusinessUnitShippingAddressRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitShippingAddressRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitShippingAddressRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Status](ctp:api:type:BusinessUnitChangeStatusAction) update action.
*
 */
type BusinessUnitStatusChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Updated status of the [Business Unit](ctp:api:type:BusinessUnit).
	Active BusinessUnitStatus `json:"active"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitStatusChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitStatusChangedMessage
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
func (obj BusinessUnitStatusChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitStatusChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitStatusChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Add Store](ctp:api:type:BusinessUnitAddStoreAction) update action.
*
 */
type BusinessUnitStoreAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The [Store](ctp:api:type:Store) that was added to the [Business Unit](ctp:api:type:BusinessUnit).
	Store StoreKeyReference `json:"store"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitStoreAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitStoreAddedMessage
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
func (obj BusinessUnitStoreAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitStoreAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitStoreAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Store Mode](ctp:api:type:BusinessUnitSetStoreModeAction) update action.
*
 */
type BusinessUnitStoreModeChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Stores](ctp:api:type:Store) of the [Business Unit](ctp:api:type:BusinessUnit) after the [Set Store Mode](ctp:api:type:BusinessUnitSetStoreModeAction) update action.
	Stores []StoreKeyReference `json:"stores"`
	// [BusinessUnitStoreMode](ctp:api:type:BusinessUnitStoreMode) of the Business Unit after the [Set Store Mode](ctp:api:type:BusinessUnitSetStoreModeAction) update action.
	StoreMode BusinessUnitStoreMode `json:"storeMode"`
	// [Stores](ctp:api:type:Store) of the [Business Unit](ctp:api:type:BusinessUnit) before the [Set Store Mode](ctp:api:type:BusinessUnitSetStoreModeAction) update action.
	OldStores []StoreKeyReference `json:"oldStores"`
	// [BusinessUnitStoreMode](ctp:api:type:BusinessUnitStoreMode) of the Business Unit before the [Set Store Mode](ctp:api:type:BusinessUnitSetStoreModeAction) update action.
	OldStoreMode BusinessUnitStoreMode `json:"oldStoreMode"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitStoreModeChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitStoreModeChangedMessage
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
func (obj BusinessUnitStoreModeChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitStoreModeChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitStoreModeChanged", Alias: (*Alias)(&obj)})
}

type BusinessUnitStoreRemovedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The [Store](ctp:api:type:Store) that was removed from the [Business Unit](ctp:api:type:BusinessUnit).
	Store StoreKeyReference `json:"store"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitStoreRemovedMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitStoreRemovedMessage
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
func (obj BusinessUnitStoreRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitStoreRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitStoreRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Stores](ctp:api:type:BusinessUnitSetStoresAction) update action.
*
 */
type BusinessUnitStoresSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Stores](ctp:api:type:Store) of the [Business Unit](ctp:api:type:BusinessUnit) after the [Set Stores](ctp:api:type:BusinessUnitSetStoresAction) update action.
	Stores []StoreKeyReference `json:"stores"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitStoresSetMessage) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitStoresSetMessage
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
func (obj BusinessUnitStoresSetMessage) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitStoresSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitStoresSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Create Category](/../api/projects/categories#create-category) request.
*
 */
type CategoryCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Category](ctp:api:type:Category) that was created.
	Category Category `json:"category"`
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

/**
*	Generated after a successful [Change Slug](ctp:api:type:CategoryChangeSlugAction) update action.
*
 */
type CategorySlugChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The slug of the [Category](ctp:api:type:Category) after the [Change Slug](ctp:api:type:CategoryChangeSlugAction) update action.
	Slug LocalizedString `json:"slug"`
	// The slug of the [Category](ctp:api:type:Category) before the [Change Slug](ctp:api:type:CategoryChangeSlugAction) update action.
	OldSlug *LocalizedString `json:"oldSlug,omitempty"`
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

/**
*	Generated after a successful [Add Address](ctp:api:type:CustomerAddAddressAction) update action.
*
 */
type CustomerAddressAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Address](ctp:api:type:Address) that was added during the [Add Address](ctp:api:type:CustomerAddAddressAction) update action.
	Address Address `json:"address"`
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

/**
*	Generated after a successful [Change Address](ctp:api:type:CustomerChangeAddressAction) update action.
*
 */
type CustomerAddressChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Address](ctp:api:type:Address) that was set during the [Change Address](ctp:api:type:CustomerChangeAddressAction) update action.
	Address Address `json:"address"`
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

/**
*	Generated after a successful [Remove Address](ctp:api:type:CustomerRemoveAddressAction) update action.
*
 */
type CustomerAddressRemovedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Address](ctp:api:type:Address) that was removed during the [Remove Address](ctp:api:type:CustomerRemoveAddressAction) update action.
	Address Address `json:"address"`
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

/**
*	Generated after a successful [Set Company Name](ctp:api:type:CustomerSetCompanyNameAction) update action.
*
 */
type CustomerCompanyNameSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The `companyName` that was set during the [Set Company Name](ctp:api:type:CustomerSetCompanyNameAction) update action.
	CompanyName *string `json:"companyName,omitempty"`
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

/**
*	Generated after a successful [Create Customer](/../api/projects/customers#create-sign-up-customer) request.
*
 */
type CustomerCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Customer](ctp:api:type:Customer) that was created.
	Customer Customer `json:"customer"`
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

/**
*	Generated after a successful [Set Date of Birth](ctp:api:type:CustomerSetDateOfBirthAction) update action.
*
 */
type CustomerDateOfBirthSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The `dateOfBirth` that was set during the [Set Date of Birth](ctp:api:type:CustomerSetDateOfBirthAction) update action.
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
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

/**
*	Generated after a successful [Delete Customer](/../api/projects/customers#delete-customer) request.
*
 */
type CustomerDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
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

/**
*	Generated after a successful [Change Email](ctp:api:type:CustomerChangeEmailAction) update action.
*
 */
type CustomerEmailChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The `email` that was set during the [Change Email](ctp:api:type:CustomerChangeEmailAction) update action.
	Email string `json:"email"`
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

/**
*	Generated after a successful [Verify Customer's Email](/../api/projects/customers#verify-email-of-customer) request.
*
 */
type CustomerEmailVerifiedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
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

/**
*	Generated after a successful [Set First Name](ctp:api:type:CustomerSetFirstNameAction) update action.
*
 */
type CustomerFirstNameSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The `firstName` that was set during the [Set First Name](ctp:api:type:CustomerSetFirstNameAction) update action.
	FirstName *string `json:"firstName,omitempty"`
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

/**
*	Generated after a successful [Set Customer Group](ctp:api:type:CustomerSetCustomerGroupAction) update action.
*
 */
type CustomerGroupSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Customer Group](ctp:api:type:CustomerGroup) that was set during the [Set Customer Group](ctp:api:type:CustomerSetCustomerGroupAction) update action.
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

/**
*	Generated after a successful [Set Last Name](ctp:api:type:CustomerSetLastNameAction) update action.
*
 */
type CustomerLastNameSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The `lastName` that was set during the [Set Last Name](ctp:api:type:CustomerSetLastNameAction) update action.
	LastName *string `json:"lastName,omitempty"`
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

/**
*	Generated after a successful [Reset Customer's Password](/../api/projects/customers#reset-password-of-customer), [Reset Customer's Password in a Store](/../api/projects/customers#reset-password-of-customer-in-store), [Change Customer's Password](/../api/projects/customers#change-password-of-customer), or [Change Customer's Password in a Store](/../api/projects/customers#change-password-of-customer-in-store) request. This Message is also produced during equivalent requests to the [My Customer Profile](/../api/projects/me-profile) endpoint.
*
 */
type CustomerPasswordUpdatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Whether the Customer's password was updated during the [Reset password](/../api/projects/customers#password-reset-of-customer) or [Change password](/../api/projects/customers#change-password-of-customer) flow.
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

/**
*	Generated after a successful [Set Title](ctp:api:type:CustomerSetTitleAction) update action.
*
 */
type CustomerTitleSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The `title` that was set during the [Set Title](ctp:api:type:CustomerSetTitleAction) update action.
	Title *string `json:"title,omitempty"`
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

/**
*	Generated after a successful [Create InventoryEntry](/../api/projects/inventory#create-inventoryentry) request.
*
 */
type InventoryEntryCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [InventoryEntry](ctp:api:type:InventoryEntry) that was created.
	InventoryEntry InventoryEntry `json:"inventoryEntry"`
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

/**
*	Generated after a successful [Delete InventoryEntry](/../api/projects/inventory#delete-inventoryentry) request.
*
 */
type InventoryEntryDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The `sku` of the [InventoryEntry](ctp:api:type:InventoryEntry) that was deleted.
	Sku string `json:"sku"`
	// [Reference](ctp:api:type:Reference) to the [Channel](ctp:api:type:Channel) where the [InventoryEntry](ctp:api:type:InventoryEntry) was deleted.
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

/**
*	Generated after a successful [Add Quantity](ctp:api:type:InventoryEntryAddQuantityAction), [Remove Quantity](ctp:api:type:InventoryEntryRemoveQuantityAction) or [Change Quantity](ctp:api:type:InventoryEntryChangeQuantityAction) update action.
*	Inventory changes as a result of [Order creation](/../api/projects/orders#create-order) do not trigger this message.
*
 */
type InventoryEntryQuantitySetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Quantity on stock for the [InventoryEntry](ctp:api:type:InventoryEntry) before the quantity was updated.
	OldQuantityOnStock int `json:"oldQuantityOnStock"`
	// Quantity on stock for the [InventoryEntry](ctp:api:type:InventoryEntry) after the quantity was updated.
	NewQuantityOnStock int `json:"newQuantityOnStock"`
	// Available quantity for the [InventoryEntry](ctp:api:type:InventoryEntry) before the quantity was updated.
	OldAvailableQuantity int `json:"oldAvailableQuantity"`
	// Available quantity for the [InventoryEntry](ctp:api:type:InventoryEntry) after the quantity was updated.
	NewAvailableQuantity int `json:"newAvailableQuantity"`
	// [Reference](ctp:api:type:Reference) to the [Channel](ctp:api:type:Channel) where the [InventoryEntry](ctp:api:type:InventoryEntry) quantity was set.
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

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [Message](ctp:api:type:Message).
*
 */
type MessagePagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// [Messages](ctp:api:type:Message) matching the query.
	Results []Message `json:"results"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MessagePagedQueryResponse) UnmarshalJSON(data []byte) error {
	type Alias MessagePagedQueryResponse
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Results {
		var err error
		obj.Results[i], err = mapDiscriminatorMessage(obj.Results[i])
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Holds the configuration for the [Messages Query](/../api/projects/messages) feature for the Project.
 */
type MessagesConfiguration struct {
	// When `true`, the [Messages Query](/../api/projects/messages) feature is active.
	Enabled bool `json:"enabled"`
	// Specifies the number of days each Message should be available via the [Messages Query](/../api/projects/messages) API.
	// For Messages older than the specified period, it is not guaranteed that they are still accessible via the API.
	// This field may not be present on Projects created before 8 October 2018.
	DeleteDaysAfterCreation *int `json:"deleteDaysAfterCreation,omitempty"`
}

/**
*	Defines the configuration for the [Messages Query](/../api/projects/messages) feature for the Project.
 */
type MessagesConfigurationDraft struct {
	// Setting to `true` activates the [Messages Query](/../api/projects/messages) feature.
	Enabled bool `json:"enabled"`
	// Specifies the number of days each Message should be available via the [Messages Query](/../api/projects/messages) API. For Messages older than the specified period, it is not guaranteed that they are still accessible via the API.
	DeleteDaysAfterCreation int `json:"deleteDaysAfterCreation"`
}

type OrderMessage interface{}

func mapDiscriminatorOrderMessage(input interface{}) (OrderMessage, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
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
	case "OrderCustomLineItemAdded":
		obj := OrderCustomLineItemAddedMessage{}
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
	case "OrderCustomLineItemQuantityChanged":
		obj := OrderCustomLineItemQuantityChangedMessage{}
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
	case "OrderCustomLineItemRemoved":
		obj := OrderCustomLineItemRemovedMessage{}
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
	case "ReturnInfoAdded":
		obj := ReturnInfoAddedMessage{}
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
		obj := ReturnInfoSetMessage{}
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

/**
*	Generated after a successful [Transition Custom Line Item State](ctp:api:type:OrderTransitionCustomLineItemStateAction) update action.
*
 */
type CustomLineItemStateTransitionMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Custom Line Item](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
	// Date and time (UTC) when the transition of the [Custom Line Item](ctp:api:type:CustomLineItem) [State](ctp:api:type:State) was performed.
	TransitionDate time.Time `json:"transitionDate"`
	// Number of [Custom Line Items](ctp:api:type:CustomLineItem) for which the [State](ctp:api:type:State) was transitioned.
	Quantity int `json:"quantity"`
	// [State](ctp:api:type:State) the [Custom Line Item](ctp:api:type:CustomLineItem) was transitioned from.
	FromState StateReference `json:"fromState"`
	// [State](ctp:api:type:State) the [Custom Line Item](ctp:api:type:CustomLineItem) was transitioned to.
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

/**
*	Generated after a successful [Add Delivery](ctp:api:type:OrderAddDeliveryAction) update action.
*
 */
type DeliveryAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Delivery](ctp:api:type:Delivery) that was added to the [Order](ctp:api:type:Order). The [Delivery](ctp:api:type:Delivery) in the Message body does not contain [Parcels](ctp:api:type:Parcel) if those were part of the initial [Add Delivery](ctp:api:type:OrderAddDeliveryAction) update action. In that case, the update action produces an additional [ParcelAddedToDelivery](ctp:api:type:ParcelAddedToDeliveryMessage) Message containing information about the [Parcels](ctp:api:type:Parcel).
	Delivery Delivery `json:"delivery"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Set Delivery Address](ctp:api:type:OrderSetDeliveryAddressAction) update action.
*
 */
type DeliveryAddressSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Parcel](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// [Address](ctp:api:type:Address) after the [Set Delivery Address](ctp:api:type:OrderSetDeliveryAddressAction) update action.
	Address *Address `json:"address,omitempty"`
	// [Address](ctp:api:type:Address) before the [Set Delivery Address](ctp:api:type:OrderSetDeliveryAddressAction) update action.
	OldAddress *Address `json:"oldAddress,omitempty"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Set Delivery Items](ctp:api:type:OrderSetDeliveryItemsAction) update action.
*
 */
type DeliveryItemsUpdatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// [Delivery Items](ctp:api:type:DeliveryItem) after the [Set Delivery Items](ctp:api:type:OrderSetDeliveryItemsAction) update action.
	Items []DeliveryItem `json:"items"`
	// [Delivery Items](ctp:api:type:DeliveryItem) before the [Set Delivery Items](ctp:api:type:OrderSetDeliveryItemsAction) update action.
	OldItems []DeliveryItem `json:"oldItems"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Remove Delivery](ctp:api:type:OrderRemoveDeliveryAction) update action.
*
 */
type DeliveryRemovedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The [Delivery](ctp:api:type:Delivery) that was removed from the [Order](ctp:api:type:Order).
	Delivery Delivery `json:"delivery"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Transition Line Item State](ctp:api:type:OrderTransitionLineItemStateAction) update action.
*
 */
type LineItemStateTransitionMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Line Item](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
	// Date and time (UTC) when the transition of the [Line Item](ctp:api:type:LineItem) [State](ctp:api:type:State) was performed.
	TransitionDate time.Time `json:"transitionDate"`
	// Number of [Line Items](ctp:api:type:LineItem) for which the [State](ctp:api:type:State) was transitioned.
	Quantity int `json:"quantity"`
	// [State](ctp:api:type:State) the [Line Item](ctp:api:type:LineItem) was transitioned from.
	FromState StateReference `json:"fromState"`
	// [State](ctp:api:type:State) the [Line Item](ctp:api:type:LineItem) was transitioned to.
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

/**
*	Generated after a successful [Set Billing Address](ctp:api:type:OrderSetBillingAddressAction) update action.
*
 */
type OrderBillingAddressSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Billing address on the Order after the [Set Billing Address](ctp:api:type:OrderSetBillingAddressAction) update action.
	Address *Address `json:"address,omitempty"`
	// Billing address on the Order before the [Set Billing Address](ctp:api:type:OrderSetBillingAddressAction) update action.
	OldAddress *Address `json:"oldAddress,omitempty"`
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

/**
*	Generated after a successful [Create Order](/../api/projects/orders#create-order) request.
*
 */
type OrderCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Order](ctp:api:type:Order) that was created.
	Order Order `json:"order"`
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

/**
*	Generated after a successful [Add Custom Line Item](ctp:api:type:StagedOrderAddCustomLineItemAction) update action.
*
 */
type OrderCustomLineItemAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Custom Line Item](ctp:api:type:CustomLineItem) that was added to the [Order](ctp:api:type:Order).
	CustomLineItem CustomLineItem `json:"customLineItem"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderCustomLineItemAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderCustomLineItemAddedMessage
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
func (obj OrderCustomLineItemAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomLineItemAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful recalculation of a Discount on a [Custom Line Item](ctp:api:type:CustomLineItem).
*
 */
type OrderCustomLineItemDiscountSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier for the [Custom Line Item](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
	// Array of [DiscountedLineItemPriceForQuantity](ctp:api:type:DiscountedLineItemPriceForQuantity) after the Discount recalculation.
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	// [TaxedItemPrice](ctp:api:type:TaxedItemPrice) of the [Custom Line Item](ctp:api:type:CustomLineItem) after the Discount recalculation.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
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

/**
*	Generated after a successful [Change Custom Line Item Quantity](ctp:api:type:StagedOrderChangeCustomLineItemQuantityAction) update action.
*
 */
type OrderCustomLineItemQuantityChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Custom Line Item](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
	// [Custom Line Item](ctp:api:type:CustomLineItem) quantity after the [Change Custom Line Item Quantity](ctp:api:type:StagedOrderChangeCustomLineItemQuantityAction) update action.
	Quantity int `json:"quantity"`
	// [Custom Line Item](ctp:api:type:CustomLineItem) quantity before the [Change Custom Line Item Quantity](ctp:api:type:StagedOrderChangeCustomLineItemQuantityAction) update action.
	OldQuantity int `json:"oldQuantity"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderCustomLineItemQuantityChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderCustomLineItemQuantityChangedMessage
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
func (obj OrderCustomLineItemQuantityChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemQuantityChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomLineItemQuantityChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Remove Custom Line Item](ctp:api:type:StagedOrderRemoveCustomLineItemAction) update action.
*
 */
type OrderCustomLineItemRemovedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Custom Line Item](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
	// [Custom Line Item](ctp:api:type:CustomLineItem) that was removed from the [Order](ctp:api:type:Order).
	CustomLineItem CustomLineItem `json:"customLineItem"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderCustomLineItemRemovedMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderCustomLineItemRemovedMessage
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
func (obj OrderCustomLineItemRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemRemovedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomLineItemRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Customer Email](ctp:api:type:OrderSetCustomerEmailAction) update action.
*
 */
type OrderCustomerEmailSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Email address on the [Order](ctp:api:type:Order) after the [Set Customer Email](ctp:api:type:OrderSetCustomerEmailAction) update action.
	Email *string `json:"email,omitempty"`
	// Email address on the [Order](ctp:api:type:Order) before the [Set Customer Email](ctp:api:type:OrderSetCustomerEmailAction) update action.
	OldEmail *string `json:"oldEmail,omitempty"`
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

/**
*	Generated after a successful [Set Customer Group](ctp:api:type:StagedOrderSetCustomerGroupAction) update action.
*
 */
type OrderCustomerGroupSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) on the [Order](ctp:api:type:Order) after the [Set Customer Group](ctp:api:type:StagedOrderSetCustomerGroupAction) update action.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) on the [Order](ctp:api:type:Order) before the [Set Customer Group](ctp:api:type:StagedOrderSetCustomerGroupAction) update action.
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

/**
*	Generated after a successful [Set Customer Id](ctp:api:type:OrderSetCustomerIdAction) update action.
*
 */
type OrderCustomerSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Customer](ctp:api:type:Customer) on the [Order](ctp:api:type:Order) after the [Set Customer Id](ctp:api:type:OrderSetCustomerIdAction) update action.
	Customer *CustomerReference `json:"customer,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) on the [Order](ctp:api:type:Order) after the [Set Customer Id](ctp:api:type:OrderSetCustomerIdAction) update action.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [Customer](ctp:api:type:Customer) on the [Order](ctp:api:type:Order) before the [Set Customer Id](ctp:api:type:OrderSetCustomerIdAction) update action.
	OldCustomer *CustomerReference `json:"oldCustomer,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) on the [Order](ctp:api:type:Order) before the [Set Customer Id](ctp:api:type:OrderSetCustomerIdAction) update action.
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

/**
*	Generated after a successful [Delete Order](/../api/projects/orders#delete-order) request.
*
 */
type OrderDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Order](ctp:api:type:Order) that has been deleted.
	Order Order `json:"order"`
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

/**
*	Generated after a successful [Add Discount Code](ctp:api:type:StagedOrderAddDiscountCodeAction) update action.
*
 */
type OrderDiscountCodeAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [DiscountCode](ctp:api:type:DiscountCode) that was added.
	DiscountCode DiscountCodeReference `json:"discountCode"`
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

/**
*	Generated after a successful [Remove Discount Code](ctp:api:type:StagedOrderRemoveDiscountCodeAction) update action.
*
 */
type OrderDiscountCodeRemovedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [DiscountCode](ctp:api:type:DiscountCode) that was removed.
	DiscountCode DiscountCodeReference `json:"discountCode"`
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

/**
*	Generated after the [DiscountCodeState](ctp:api:type:DiscountCodeState) changes due to a [recalculation](/../api/projects/carts#recalculate).
*
 */
type OrderDiscountCodeStateSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [DiscountCode](ctp:api:type:DiscountCode) that changed due to the recalculation.
	DiscountCode DiscountCodeReference `json:"discountCode"`
	// [DiscountCodeState](ctp:api:type:DiscountCodeState) after the recalculation.
	State DiscountCodeState `json:"state"`
	// [DiscountCodeState](ctp:api:type:DiscountCodeState) before the recalculation.
	OldState *DiscountCodeState `json:"oldState,omitempty"`
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

/**
*	Generated after a successfully applying an [OrderEdit](/../api/projects/order-edits#apply-an-orderedit).
*
 */
type OrderEditAppliedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [OrderEdit](ctp:api:type:OrderEdit) that was applied.
	Edit OrderEdit `json:"edit"`
	// Information about a successfully applied [OrderEdit](ctp:api:type:OrderEdit).
	Result OrderEditApplied `json:"result"`
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

/**
*	Generated after a successful [Order Import](/../api/projects/orders-import#create-an-order-by-import).
*
 */
type OrderImportedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Order](ctp:api:type:Order) that was imported.
	Order Order `json:"order"`
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

/**
*	Generated after a successful [Add Line Item](ctp:api:type:StagedOrderAddLineItemAction) update action.
*
 */
type OrderLineItemAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Line Item](ctp:api:type:LineItem) that was added to the [Order](ctp:api:type:Order).
	LineItem LineItem `json:"lineItem"`
	// Quantity of [Line Items](ctp:api:type:LineItem) that were added to the [Order](ctp:api:type:Order).
	AddedQuantity int `json:"addedQuantity"`
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

/**
*	Generated after a successful recalculation of a Discount on a [Line Item](ctp:api:type:LineItem).
*
 */
type OrderLineItemDiscountSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier for the [Line Item](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
	// Array of [DiscountedLineItemPriceForQuantity](ctp:api:type:DiscountedLineItemPriceForQuantity) after the Discount recalculation.
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	// Total Price of the [Line Item](ctp:api:type:LineItem) after the Discount recalculation.
	TotalPrice Money `json:"totalPrice"`
	// [TaxedItemPrice](ctp:api:type:TaxedItemPrice) of the [Line Item](ctp:api:type:LineItem) after the Discount recalculation.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
	// Taxed price of the Shipping Methods in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode)..
	TaxedPricePortions []MethodTaxedPrice `json:"taxedPricePortions"`
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

/**
*	Generated after a successful [Set Line Item Distribution Channel](/../api/projects/order-edits#set-lineitem-distributionchannel) update action.
*
 */
type OrderLineItemDistributionChannelSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Line Item](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
	// [Distribution Channel](ctp:api:type:Channel) that was set.
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

/**
*	Generated after a successful [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
*
 */
type OrderLineItemRemovedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Line Item](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
	// Quantity of [Line Items](ctp:api:type:LineItem) that were removed during the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	RemovedQuantity int `json:"removedQuantity"`
	// [Line Item](ctp:api:type:LineItem) quantity after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	NewQuantity int `json:"newQuantity"`
	// [ItemStates](ctp:api:type:ItemState) after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	NewState []ItemState `json:"newState"`
	// `totalPrice` of the [Order](ctp:api:type:Order) after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	NewTotalPrice CentPrecisionMoney `json:"newTotalPrice"`
	// [TaxedItemPrice](ctp:api:type:TaxedItemPrice) of the [Order](ctp:api:type:Order) after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	NewTaxedPrice *TaxedItemPrice `json:"newTaxedPrice,omitempty"`
	// [Price](ctp:api:type:Price) of the [Order](ctp:api:type:Order) after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	NewPrice *Price `json:"newPrice,omitempty"`
	// [Shipping Details](ctp:api:type:ItemShippingDetails) of the [Order](ctp:api:type:Order) after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
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

/**
*	Generated after a successful [Add Payment](ctp:api:type:OrderAddPaymentAction) update action or when a [Payment](ctp:api:type:Payment) is added via [Order Edits](ctp:api:type:StagedOrderAddPaymentAction).
*
 */
type OrderPaymentAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Payment](ctp:api:type:Payment) that was added to the [Order](ctp:api:type:Order).
	Payment PaymentReference `json:"payment"`
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

/**
*	Generated after a successful [Change Payment State](ctp:api:type:OrderChangePaymentStateAction) update action.
*
 */
type OrderPaymentStateChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [PaymentState](ctp:api:type:PaymentState) after the [Change Payment State](ctp:api:type:OrderChangePaymentStateAction) update action.
	PaymentState PaymentState `json:"paymentState"`
	// [PaymentState](ctp:api:type:PaymentState) before the [Change Payment State](ctp:api:type:OrderChangePaymentStateAction) update action.
	OldPaymentState *PaymentState `json:"oldPaymentState,omitempty"`
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

/**
*	Generated after a successful [Set Return Shipment State](ctp:api:type:OrderSetReturnShipmentStateAction) update action.
*
 */
type OrderReturnShipmentStateChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [ReturnItem](ctp:api:type:ReturnItem).
	ReturnItemId string `json:"returnItemId"`
	// State of the [ReturnItem](ctp:api:type:ReturnItem) after the [Set Return Shipment State](ctp:api:type:OrderSetReturnShipmentStateAction) update action.
	ReturnShipmentState ReturnShipmentState `json:"returnShipmentState"`
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

/**
*	Generated after a successful [Change Shipment State](ctp:api:type:OrderChangeShipmentStateAction) update action.
*
 */
type OrderShipmentStateChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [ShipmentState](ctp:api:type:ShipmentState) after the [Change Shipment State](ctp:api:type:OrderChangeShipmentStateAction) update action.
	ShipmentState ShipmentState `json:"shipmentState"`
	// [ShipmentState](ctp:api:type:ShipmentState) before the [Change Shipment State](ctp:api:type:OrderChangeShipmentStateAction) update action.
	OldShipmentState ShipmentState `json:"oldShipmentState"`
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

/**
*	Generated after a successful [Set Shipping Address](ctp:api:type:OrderSetShippingAddressAction) update action.
*
 */
type OrderShippingAddressSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Shipping address on the Order after the [Set Shipping Address](ctp:api:type:OrderSetShippingAddressAction) update action.
	Address *Address `json:"address,omitempty"`
	// Shipping address on the Order before the [Set Shipping Address](ctp:api:type:OrderSetShippingAddressAction) update action.
	OldAddress *Address `json:"oldAddress,omitempty"`
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

/**
*	Generated after a successful [Set Shipping Method](ctp:api:type:StagedOrderSetShippingMethodAction) and [Set Custom Shipping Method](ctp:api:type:StagedOrderSetCustomShippingMethodAction) update actions.
*
 */
type OrderShippingInfoSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [ShippingInfo](ctp:api:type:ShippingInfo) after the [Set Shipping Method](ctp:api:type:StagedOrderSetShippingMethodAction) or [Set Custom Shipping Method](ctp:api:type:StagedOrderSetCustomShippingMethodAction) update action.
	ShippingInfo *ShippingInfo `json:"shippingInfo,omitempty"`
	// [ShippingInfo](ctp:api:type:ShippingInfo) before the [Set Shipping Method](ctp:api:type:StagedOrderSetShippingMethodAction) or [Set Custom Shipping Method](ctp:api:type:StagedOrderSetCustomShippingMethodAction) update action.
	OldShippingInfo *ShippingInfo `json:"oldShippingInfo,omitempty"`
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

/**
*	Generated after a successful [Set ShippingRateInput](ctp:api:type:StagedOrderSetShippingRateInputAction) update action.
*
 */
type OrderShippingRateInputSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [ShippingRateInput](ctp:api:type:ShippingRateInput) after the [Set ShippingRateInput](ctp:api:type:StagedOrderSetShippingRateInputAction) update action.
	ShippingRateInput ShippingRateInput `json:"shippingRateInput,omitempty"`
	// [ShippingRateInput](ctp:api:type:ShippingRateInput) before the [Set ShippingRateInput](ctp:api:type:StagedOrderSetShippingRateInputAction) update action.
	OldShippingRateInput ShippingRateInput `json:"oldShippingRateInput,omitempty"`
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

/**
*	Generated after a successful [Change Order State](ctp:api:type:OrderChangeOrderStateAction) update action.
*
 */
type OrderStateChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [OrderState](ctp:api:type:OrderState) after the [Change Order State](ctp:api:type:OrderChangeOrderStateAction) update action.
	OrderState OrderState `json:"orderState"`
	// [OrderState](ctp:api:type:OrderState) before the [Change Order State](ctp:api:type:OrderChangeOrderStateAction) update action.
	OldOrderState OrderState `json:"oldOrderState"`
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

/**
*	Generated after a successful [Transition State](ctp:api:type:OrderTransitionStateAction) update action.
*
 */
type OrderStateTransitionMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [OrderState](ctp:api:type:OrderState) after the [Transition State](ctp:api:type:OrderTransitionStateAction) update action.
	State StateReference `json:"state"`
	// [OrderState](ctp:api:type:OrderState) before the [Transition State](ctp:api:type:OrderTransitionStateAction) update action.
	OldState *StateReference `json:"oldState,omitempty"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:OrderTransitionStateAction) update action.
	Force bool `json:"force"`
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

/**
*	Generated after a successful [Set Store](ctp:api:type:OrderSetStoreAction) update action.
*
 */
type OrderStoreSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Store](ctp:api:type:Store) that was set.
	Store *StoreKeyReference `json:"store,omitempty"`
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

/**
*	Generated after a successful [Add Parcel To Delivery](ctp:api:type:OrderAddParcelToDeliveryAction) update action.
*
 */
type ParcelAddedToDeliveryMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	Delivery Delivery `json:"delivery"`
	// [Parcel](ctp:api:type:Parcel) that was added to the [Delivery](ctp:api:type:Delivery).
	Parcel Parcel `json:"parcel"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Set Parcel Items](ctp:api:type:OrderSetParcelItemsAction) update action.
*
 */
type ParcelItemsUpdatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Parcel](ctp:api:type:Parcel).
	ParcelId string `json:"parcelId"`
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// [Delivery Items](ctp:api:type:DeliveryItem) after the [Set Parcel Items](ctp:api:type:OrderSetParcelItemsAction) update action.
	Items []DeliveryItem `json:"items"`
	// [Delivery Items](ctp:api:type:DeliveryItem) before the [Set Parcel Items](ctp:api:type:OrderSetParcelItemsAction) update action.
	OldItems []DeliveryItem `json:"oldItems"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Set Parcel Measurements](ctp:api:type:OrderSetParcelMeasurementsAction) update action.
*
 */
type ParcelMeasurementsUpdatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// Unique identifier of the [Parcel](ctp:api:type:Parcel).
	ParcelId string `json:"parcelId"`
	// The [Parcel Measurements](ctp:api:type:ParcelMeasurements) that were set on the [Parcel](ctp:api:type:Parcel).
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Remove Parcel From Delivery](ctp:api:type:OrderRemoveParcelFromDeliveryAction) update action.
*
 */
type ParcelRemovedFromDeliveryMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// [Parcel](ctp:api:type:Parcel) that was removed from the [Delivery](ctp:api:type:Delivery).
	Parcel Parcel `json:"parcel"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Set Parcel TrackingData](ctp:api:type:OrderSetParcelTrackingDataAction) update action.
*
 */
type ParcelTrackingDataUpdatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// Unique identifier of the [Parcel](ctp:api:type:Parcel).
	ParcelId string `json:"parcelId"`
	// The [Tracking Data](ctp:api:type:TrackingData) that was added to the [Parcel](ctp:api:type:Parcel).
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Create Payment](/../api/projects/payments#create-a-payment) request.
*
 */
type PaymentCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Payment](ctp:api:type:Payment) that was created.
	Payment Payment `json:"payment"`
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

/**
*	Generated after a successful [Add InterfaceInteraction](ctp:api:type:PaymentAddInterfaceInteractionAction) update action.
*
 */
type PaymentInteractionAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The interface interaction that was added to the [Payment](ctp:api:type:Payment).
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

/**
*	Generated after a successful [Set StatusInterfaceCode](ctp:api:type:PaymentSetStatusInterfaceCodeAction) update action.
*
 */
type PaymentStatusInterfaceCodeSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier for the [Payment](ctp:api:type:Payment) for which the [Set StatusInterfaceCode](ctp:api:type:PaymentSetStatusInterfaceCodeAction) update action was applied.
	PaymentId string `json:"paymentId"`
	// The `interfaceCode` that was set during the [Set StatusInterfaceCode](ctp:api:type:PaymentSetStatusInterfaceCodeAction) update action.
	InterfaceCode *string `json:"interfaceCode,omitempty"`
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

/**
*	Generated after a successful [Transition State](ctp:api:type:PaymentTransitionStateAction) update action.
*
 */
type PaymentStatusStateTransitionMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [State](ctp:api:type:State) of the [Payment](ctp:api:type:Payment) after the [Transition State](ctp:api:type:PaymentTransitionStateAction) update action.
	State StateReference `json:"state"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Change Transaction State](ctp:api:type:PaymentChangeTransactionStateAction) update action.
	Force bool `json:"force"`
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

/**
*	Generated after a successful [Add Transaction](ctp:api:type:PaymentAddTransactionAction) update action.
*
 */
type PaymentTransactionAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Transaction](ctp:api:type:Transaction) that was added to the [Payment](ctp:api:type:Payment).
	Transaction Transaction `json:"transaction"`
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

/**
*	Generated after a successful [Change Transaction State](ctp:api:type:PaymentChangeTransactionStateAction) update action.
*
 */
type PaymentTransactionStateChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier for the [Transaction](ctp:api:type:Transaction) for which the [Transaction State](ctp:api:type:TransactionState) changed.
	TransactionId string `json:"transactionId"`
	// [Transaction State](ctp:api:type:TransactionState) after the [Change Transaction State](ctp:api:type:PaymentChangeTransactionStateAction) update action.
	State TransactionState `json:"state"`
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

/**
*	Generated after a successful [Add To Category](ctp:api:type:ProductAddToCategoryAction) update action.
*
 */
type ProductAddedToCategoryMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Category](ctp:api:type:Category) the [Product](ctp:api:type:Product) was added to.
	Category CategoryReference `json:"category"`
	// Whether the update was only applied to the staged [Product Projection](ctp:api:type:ProductProjection).
	Staged bool `json:"staged"`
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

/**
*	Generated after a successful [Create Product](/../api/projects/products#create-product) request.
*
 */
type ProductCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The staged [Product Projection](ctp:api:type:ProductProjection) of the [Product](ctp:api:type:Product) at the time of creation.
	ProductProjection ProductProjection `json:"productProjection"`
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

/**
*	Generated after a successful [Delete Product](/../api/projects/products#delete-product) request.
*
 */
type ProductDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// List of image URLs that were removed during the [Delete Product](ctp:api:type:Product) request.
	RemovedImageUrls []string `json:"removedImageUrls"`
	// Current [Product Projection](ctp:api:type:ProductProjection) of the deleted [Product](ctp:api:type:Product).
	CurrentProjection *ProductProjection `json:"currentProjection,omitempty"`
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

/**
*	Generated after a successful [Add External Image](ctp:api:type:ProductAddExternalImageAction) update action or after the successful [upload of an image](/../api/projects/products#upload-product-image).
*
 */
type ProductImageAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Product Variant](ctp:api:type:ProductVariant) to which the [Image](ctp:api:type:Image) was added.
	VariantId int `json:"variantId"`
	// [Image](ctp:api:type:Image) that was added.
	Image Image `json:"image"`
	// Whether the update was only applied to the staged [Product Projection](ctp:api:type:ProductProjection).
	Staged bool `json:"staged"`
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

/**
*	Generated after a Price is updated due to a [Product Discount](ctp:api:type:ProductDiscount).
*
 */
type ProductPriceDiscountsSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Array containing details about the [Embedded Prices](ctp:api:type:Price) that were updated.
	UpdatedPrices []ProductPriceDiscountsSetUpdatedPrice `json:"updatedPrices"`
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

/**
*	Details about a [Embedded Price](ctp:api:type:Price) that was updated due to a Discount. Specific to [ProductPriceDiscountsSet](ctp:api:type:ProductPriceDiscountsSetMessage) Message.
*
 */
type ProductPriceDiscountsSetUpdatedPrice struct {
	// Unique identifier of the [ProductVariant](ctp:api:type:ProductVariant) for which the Discount was set.
	VariantId int `json:"variantId"`
	// Key of the [ProductVariant](ctp:api:type:ProductVariant) for which Discount was set.
	VariantKey *string `json:"variantKey,omitempty"`
	// SKU of the [ProductVariant](ctp:api:type:ProductVariant) for which Discount was set.
	Sku *string `json:"sku,omitempty"`
	// Unique identifier of the [Embedded Price](ctp:api:type:Price).
	PriceId string `json:"priceId"`
	// Discounted Price for the [ProductVariant](ctp:api:type:ProductVariant) for which Discount was set.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Whether the update was only applied to the staged [ProductProjection](ctp:api:type:ProductProjection).
	Staged bool `json:"staged"`
}

/**
*	Generated after a successful [Set Discounted Embedded Price](ctp:api:type:ProductSetDiscountedPriceAction) update action.
*
 */
type ProductPriceExternalDiscountSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Product Variant](ctp:api:type:ProductVariant) for which the Discount was set.
	VariantId int `json:"variantId"`
	// Key of the [Product Variant](ctp:api:type:ProductVariant) for which the Discount was set.
	VariantKey *string `json:"variantKey,omitempty"`
	// SKU of the [Product Variant](ctp:api:type:ProductVariant) for which Discount was set.
	Sku *string `json:"sku,omitempty"`
	// Unique identifier of the [Embedded Price](ctp:api:type:Price).
	PriceId string `json:"priceId"`
	// Discounted Price for the [Product Variant](ctp:api:type:ProductVariant) for which Discount was set.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Whether the update was only applied to the staged [Product Projection](ctp:api:type:ProductProjection).
	Staged bool `json:"staged"`
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

/**
*	Generated after a successful [Publish](ctp:api:type:ProductPublishAction) update action.
*
 */
type ProductPublishedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// List of image URLs which were removed during the [Publish](ctp:api:type:ProductPublishAction) update action.
	RemovedImageUrls []string `json:"removedImageUrls"`
	// Current [Product Projection](ctp:api:type:ProductProjection) of the [Product](ctp:api:type:Product) at the time of creation.
	ProductProjection ProductProjection `json:"productProjection"`
	// [Publishing Scope](ctp:api:type:ProductPublishScope) that was used during the [Publish](ctp:api:type:ProductPublishAction) update action.
	Scope ProductPublishScope `json:"scope"`
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

/**
*	Generated after a successful [Remove From Category](ctp:api:type:ProductRemoveFromCategoryAction) update action.
*
 */
type ProductRemovedFromCategoryMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Category](ctp:api:type:Category) the [Product](ctp:api:type:Product) was removed from.
	Category CategoryReference `json:"category"`
	// Whether the update was only applied to the staged [Product Projection](ctp:api:type:ProductProjection).
	Staged bool `json:"staged"`
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

/**
*	Generated after a successful [Revert Staged Changes](ctp:api:type:ProductRevertStagedChangesAction) update action.
*
 */
type ProductRevertedStagedChangesMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// List of image URLs that were removed during the [Revert Staged Changes](ctp:api:type:ProductRevertStagedChangesAction) update action.
	RemovedImageUrls []string `json:"removedImageUrls"`
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

/**
*	Generated after a successful [Create Product Selection](/../api/projects/product-selections#create-product-selection) request.
*
 */
type ProductSelectionCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The `type` and `name` of the individual Product Selection.
	ProductSelection IndividualProductSelectionType `json:"productSelection"`
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

/**
*	Generated after a successful [Delete Product Selection](/../api/projects/product-selections#create-product-selection) request.
*
 */
type ProductSelectionDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
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

/**
*	Generated after a successful [Add Product](ctp:api:type:ProductSelectionAddProductAction) update action.
*
 */
type ProductSelectionProductAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Product](ctp:api:type:Product) that was added to the [Product Selection](ctp:api:type:ProductSelection).
	Product ProductReference `json:"product"`
	// Product Variant Selection after the [Add Product](ctp:api:type:ProductSelectionAddProductAction) update action.
	VariantSelection ProductVariantSelection `json:"variantSelection"`
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
	if obj.VariantSelection != nil {
		var err error
		obj.VariantSelection, err = mapDiscriminatorProductVariantSelection(obj.VariantSelection)
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

/**
*	Generated after a successful [Remove Product](ctp:api:type:ProductSelectionRemoveProductAction) update action.
*
 */
type ProductSelectionProductRemovedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Product](ctp:api:type:Product) that was removed from the Product Selection.
	Product ProductReference `json:"product"`
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

/**
*	Generated after a successful [Set Variant Selection](ctp:api:type:ProductSelectionSetVariantSelectionAction) update action.
*
 */
type ProductSelectionVariantSelectionChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Product](ctp:api:type:Product) for which the Product Variant Selection changed.
	Product ProductReference `json:"product"`
	// Product Variant Selection before the [Set Variant Selection](ctp:api:type:ProductSelectionSetVariantSelectionAction) update action.
	OldVariantSelection ProductVariantSelection `json:"oldVariantSelection"`
	// Product Variant Selection after the [Set Variant Selection](ctp:api:type:ProductSelectionSetVariantSelectionAction) update action.
	NewVariantSelection ProductVariantSelection `json:"newVariantSelection"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionVariantSelectionChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionVariantSelectionChangedMessage
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
	if obj.OldVariantSelection != nil {
		var err error
		obj.OldVariantSelection, err = mapDiscriminatorProductVariantSelection(obj.OldVariantSelection)
		if err != nil {
			return err
		}
	}
	if obj.NewVariantSelection != nil {
		var err error
		obj.NewVariantSelection, err = mapDiscriminatorProductVariantSelection(obj.NewVariantSelection)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionVariantSelectionChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionVariantSelectionChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSelectionVariantSelectionChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Slug](ctp:api:type:ProductChangeSlugAction) update action.
*
 */
type ProductSlugChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The slug of the [Product](ctp:api:type:Product) after the [Change Slug](ctp:api:type:ProductChangeSlugAction) update action.
	Slug LocalizedString `json:"slug"`
	// The slug of the [Product](ctp:api:type:Product) before the [Change Slug](ctp:api:type:ProductChangeSlugAction) update action.
	OldSlug *LocalizedString `json:"oldSlug,omitempty"`
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

/**
*	Generated after a successful [Transition State](ctp:api:type:ProductTransitionStateAction) update action.
*
 */
type ProductStateTransitionMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Product [State](ctp:api:type:State) after the [Transition State](ctp:api:type:ProductTransitionStateAction) update action.
	State StateReference `json:"state"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:ProductTransitionStateAction) update action.
	Force bool `json:"force"`
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

/**
*	Generated after a successful [Unpublish Product](ctp:api:type:ProductUnpublishAction) update action.
*
 */
type ProductUnpublishedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
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

/**
*	Generated after a successful [Add Product Variant](ctp:api:type:ProductAddVariantAction) update action.
*
 */
type ProductVariantAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Product Variant](ctp:api:type:ProductVariant) that was added.
	Variant ProductVariant `json:"variant"`
	// Whether the update was only applied to the staged [Product Projection](ctp:api:type:ProductProjection).
	Staged bool `json:"staged"`
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

/**
*	Generated after a successful [Remove Product Variant](ctp:api:type:ProductRemoveVariantAction) update action.
*
 */
type ProductVariantDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique identifier of the [Product Variant](ctp:api:type:ProductVariant) that was added.
	Variant *ProductVariant `json:"variant,omitempty"`
	// List of image URLs that were removed with the [Remove Product Variant](ctp:api:type:ProductRemoveVariantAction) update action.
	RemovedImageUrls []string `json:"removedImageUrls"`
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

/**
*	Generated after a successful [Create Quote](/../api/projects/quotes#create-quote) request.
*
 */
type QuoteCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Quote](/../api/projects/quotes) that was created.
	Quote Quote `json:"quote"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *QuoteCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias QuoteCreatedMessage
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
func (obj QuoteCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias QuoteCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteCreated", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Delete Quote](/../api/projects/quotes#delete-quote) request.
*
 */
type QuoteDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *QuoteDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias QuoteDeletedMessage
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
func (obj QuoteDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias QuoteDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteDeleted", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Create Quote Request](/../api/projects/quote-requests#create-quoterequest) request.
*
 */
type QuoteRequestCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Quote Request](/../api/projects/quote-requests) that was created.
	QuoteRequest QuoteRequest `json:"quoteRequest"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *QuoteRequestCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias QuoteRequestCreatedMessage
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
func (obj QuoteRequestCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias QuoteRequestCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteRequestCreated", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Delete Quote Request](/../api/projects/quote-requests#delete-quoterequest) request.
*
 */
type QuoteRequestDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *QuoteRequestDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias QuoteRequestDeletedMessage
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
func (obj QuoteRequestDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias QuoteRequestDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteRequestDeleted", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Quote Request State](ctp:api:type:QuoteRequestChangeQuoteRequestStateAction) update action.
*
 */
type QuoteRequestStateChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// State of the Quote Request after the [Change Quote Request State](ctp:api:type:QuoteRequestChangeQuoteRequestStateAction) update action.
	QuoteRequestState QuoteRequestState `json:"quoteRequestState"`
	// State of the Quote Request before the [Change Quote Request State](ctp:api:type:QuoteRequestChangeQuoteRequestStateAction) update action.
	OldQuoteRequestState QuoteRequestState `json:"oldQuoteRequestState"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *QuoteRequestStateChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias QuoteRequestStateChangedMessage
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
func (obj QuoteRequestStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias QuoteRequestStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteRequestStateChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Transition State](ctp:api:type:QuoteRequestTransitionStateAction) update action.
*
 */
type QuoteRequestStateTransitionMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) after the [Transition State](ctp:api:type:QuoteRequestTransitionStateAction) update action.
	State StateReference `json:"state"`
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) before the [Transition State](ctp:api:type:QuoteRequestTransitionStateAction) update action.
	OldState *StateReference `json:"oldState,omitempty"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:QuoteRequestTransitionStateAction) update action.
	Force bool `json:"force"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *QuoteRequestStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias QuoteRequestStateTransitionMessage
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
func (obj QuoteRequestStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias QuoteRequestStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteRequestStateTransition", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Quote State](ctp:api:type:QuoteChangeQuoteStateAction) update action.
*
 */
type QuoteStateChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// State of the Quote after the [Change Quote State](ctp:api:type:QuoteChangeQuoteStateAction) update action.
	QuoteState QuoteState `json:"quoteState"`
	// State of the Quote before the [Change Quote State](ctp:api:type:QuoteChangeQuoteStateAction) update action.
	OldQuoteState QuoteState `json:"oldQuoteState"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *QuoteStateChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias QuoteStateChangedMessage
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
func (obj QuoteStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias QuoteStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteStateChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Transition State](ctp:api:type:QuoteTransitionStateAction) update action.
*
 */
type QuoteStateTransitionMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) after the [Transition State](ctp:api:type:QuoteTransitionStateAction) update action.
	State StateReference `json:"state"`
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) before the [Transition State](ctp:api:type:QuoteTransitionStateAction) update action.
	OldState *StateReference `json:"oldState,omitempty"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:QuoteTransitionStateAction) update action.
	Force bool `json:"force"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *QuoteStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias QuoteStateTransitionMessage
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
func (obj QuoteStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias QuoteStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteStateTransition", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Add Return Info](ctp:api:type:OrderAddReturnInfoAction) update action.
*
 */
type ReturnInfoAddedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The [ReturnInfo](ctp:api:type:ReturnInfo) that was added to the [Order](ctp:api:type:Order).
	ReturnInfo ReturnInfo `json:"returnInfo"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReturnInfoAddedMessage) UnmarshalJSON(data []byte) error {
	type Alias ReturnInfoAddedMessage
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
func (obj ReturnInfoAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias ReturnInfoAddedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReturnInfoAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Return Info](ctp:api:type:OrderSetReturnInfoAction) update action on [Orders](ctp:api:type:Order) and [Order Edits](ctp:api:type:OrderEdit).
*
 */
type ReturnInfoSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The [ReturnInfo](ctp:api:type:ReturnInfo) that was set on the [Order](ctp:api:type:Order) or [Order Edit](ctp:api:type:OrderEdit).
	ReturnInfo []ReturnInfo `json:"returnInfo"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReturnInfoSetMessage) UnmarshalJSON(data []byte) error {
	type Alias ReturnInfoSetMessage
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
func (obj ReturnInfoSetMessage) MarshalJSON() ([]byte, error) {
	type Alias ReturnInfoSetMessage
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReturnInfoSet", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["returnInfo"] == nil {
		delete(raw, "returnInfo")
	}

	return json.Marshal(raw)

}

/**
*	Generated after a successful [Create Review](/../api/projects/reviews#create-a-review) request.
*
 */
type ReviewCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Review](ctp:api:type:Review) that was created.
	Review Review `json:"review"`
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

/**
*	Generated after a successful [Set Rating](ctp:api:type:ReviewSetRatingAction) update action.
*
 */
type ReviewRatingSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The `rating` of the [Review](ctp:api:type:Review) before the [Set Rating](ctp:api:type:ReviewSetRatingAction) update action.
	OldRating *float64 `json:"oldRating,omitempty"`
	// The `rating` of the [Review](ctp:api:type:Review) after the [Set Rating](ctp:api:type:ReviewSetRatingAction) update action.
	NewRating *float64 `json:"newRating,omitempty"`
	// Whether the [Review](ctp:api:type:Review) was taken into account in the ratings statistics of the target.
	IncludedInStatistics bool `json:"includedInStatistics"`
	// [Reference](ctp:api:type:Reference) to the resource that the [Review](ctp:api:type:Review) belongs to.
	Target Reference `json:"target,omitempty"`
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

/**
*	Generated after a successful [Transition State](ctp:api:type:ReviewTransitionStateAction) update action.
*
 */
type ReviewStateTransitionMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [State](ctp:api:type:State) of the [Review](ctp:api:type:Review) before the [Transition State](ctp:api:type:ReviewTransitionStateAction) update action.
	OldState *StateReference `json:"oldState,omitempty"`
	// [State](ctp:api:type:State) of the [Review](ctp:api:type:Review) after the [Transition State](ctp:api:type:ReviewTransitionStateAction) update action.
	NewState StateReference `json:"newState"`
	// Whether the old [Review](ctp:api:type:Review) was taken into account in the rating statistics of the target before the state transition.
	OldIncludedInStatistics bool `json:"oldIncludedInStatistics"`
	// Whether the new [Review](ctp:api:type:Review) was taken into account in the rating statistics of the target after the state transition.
	NewIncludedInStatistics bool `json:"newIncludedInStatistics"`
	// [Reference](ctp:api:type:Reference) to the resource that the [Review](ctp:api:type:Review) belongs to.
	Target Reference `json:"target,omitempty"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:ReviewTransitionStateAction) update action.
	Force bool `json:"force"`
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

/**
*	Generated after a successful [Create Staged Quote](/../api/projects/staged-quotes#create-stagedquote) request.
*
 */
type StagedQuoteCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Staged Quote](/../api/projects/staged-quotes) that was created.
	StagedQuote StagedQuote `json:"stagedQuote"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedQuoteCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias StagedQuoteCreatedMessage
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
func (obj StagedQuoteCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteCreated", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Delete Staged Quote](/../api/projects/staged-quotes#delete-stagedquote) request.
*
 */
type StagedQuoteDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedQuoteDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias StagedQuoteDeletedMessage
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
func (obj StagedQuoteDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteDeleted", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Seller Comment](ctp:api:type:StagedQuoteSetSellerCommentAction) update action.
*
 */
type StagedQuoteSellerCommentSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// `sellerComment` on the [StagedQuote](ctp:api:type:StagedQuote) after a successful [Set Seller Comment](ctp:api:type:StagedQuoteSetSellerCommentAction) update action.
	SellerComment string `json:"sellerComment"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedQuoteSellerCommentSetMessage) UnmarshalJSON(data []byte) error {
	type Alias StagedQuoteSellerCommentSetMessage
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
func (obj StagedQuoteSellerCommentSetMessage) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteSellerCommentSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteSellerCommentSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Staged Quote State](ctp:api:type:StagedQuoteChangeStagedQuoteStateAction) update action.
*
 */
type StagedQuoteStateChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// State of the Staged Quote after the [Change Staged Quote State](ctp:api:type:StagedQuoteChangeStagedQuoteStateAction) update action.
	StagedQuoteState StagedQuoteState `json:"stagedQuoteState"`
	// State of the Staged Quote before the [Change Staged Quote State](ctp:api:type:StagedQuoteChangeStagedQuoteStateAction) update action.
	OldStagedQuoteState StagedQuoteState `json:"oldStagedQuoteState"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedQuoteStateChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias StagedQuoteStateChangedMessage
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
func (obj StagedQuoteStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteStateChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteStateChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Transition State](ctp:api:type:StagedQuoteTransitionStateAction) update action.
*
 */
type StagedQuoteStateTransitionMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) after the [Transition State](ctp:api:type:StagedQuoteTransitionStateAction) update action.
	State StateReference `json:"state"`
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) before the [Transition State](ctp:api:type:StagedQuoteTransitionStateAction) update action.
	OldState *StateReference `json:"oldState,omitempty"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:StagedQuoteTransitionStateAction) update action.
	Force bool `json:"force"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedQuoteStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias StagedQuoteStateTransitionMessage
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
func (obj StagedQuoteStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteStateTransitionMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteStateTransition", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Valid To](ctp:api:type:StagedQuoteSetValidToAction) update action.
*
 */
type StagedQuoteValidToSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Expiration date for the Staged Quote after the [Set Valid To](ctp:api:type:StagedQuoteSetValidToAction) update action.
	ValidTo time.Time `json:"validTo"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedQuoteValidToSetMessage) UnmarshalJSON(data []byte) error {
	type Alias StagedQuoteValidToSetMessage
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
func (obj StagedQuoteValidToSetMessage) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteValidToSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteValidToSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Active](ctp:api:types:StandalonePriceChangeActiveAction) update action.
*
 */
type StandalonePriceActiveChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Value of the `active` field of the StandalonePrice after the [Change Active](ctp:api:types:StandalonePriceChangeActiveAction) update action.
	Active bool `json:"active"`
	// Value of the `active` field of the StandalonePrice before the [Change Active](ctp:api:types:StandalonePriceChangeActiveAction) update action.
	OldActive bool `json:"oldActive"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StandalonePriceActiveChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias StandalonePriceActiveChangedMessage
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
func (obj StandalonePriceActiveChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceActiveChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceActiveChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Create StandalonePrice](/../api/projects/standalone-prices#create-standaloneprice) request.
*
 */
type StandalonePriceCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Standalone Price](ctp:api:type:StandalonePrice) that was created.
	StandalonePrice StandalonePrice `json:"standalonePrice"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StandalonePriceCreatedMessage) UnmarshalJSON(data []byte) error {
	type Alias StandalonePriceCreatedMessage
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
func (obj StandalonePriceCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceCreatedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceCreated", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Delete StandalonePrice](/../api/projects/standalone-prices#delete-standaloneprice) request.
*
 */
type StandalonePriceDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StandalonePriceDeletedMessage) UnmarshalJSON(data []byte) error {
	type Alias StandalonePriceDeletedMessage
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
func (obj StandalonePriceDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceDeletedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceDeleted", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a [Product Discount](ctp:api:type:ProductDiscount) is successfully applied to a StandalonePrice.
*
 */
type StandalonePriceDiscountSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The new `discounted` value of the updated [StandalonePrice](ctp:api:type:StandalonePrice).
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StandalonePriceDiscountSetMessage) UnmarshalJSON(data []byte) error {
	type Alias StandalonePriceDiscountSetMessage
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
func (obj StandalonePriceDiscountSetMessage) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceDiscountSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceDiscountSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Discounted Price](ctp:api:type:StandalonePriceSetDiscountedPriceAction) update action.
*
 */
type StandalonePriceExternalDiscountSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The `discounted` value of the [StandalonePrice](ctp:api:type:StandalonePrice) after the [Set Discounted Price](ctp:api:type:StandalonePriceSetDiscountedPriceAction) update action.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StandalonePriceExternalDiscountSetMessage) UnmarshalJSON(data []byte) error {
	type Alias StandalonePriceExternalDiscountSetMessage
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
func (obj StandalonePriceExternalDiscountSetMessage) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceExternalDiscountSetMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceExternalDiscountSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Apply Staged Changes](ctp:api:types:StandalonePriceApplyStagedChangesAction) update action.
*
 */
type StandalonePriceStagedChangesAppliedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Applied changes of the [StandalonePrice](/../api/projects/standalone-prices) after the [Apply Staged Changes](ctp:api:types:StandalonePriceApplyStagedChangesAction) update action.
	StagedChanges StagedStandalonePrice `json:"stagedChanges"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StandalonePriceStagedChangesAppliedMessage) UnmarshalJSON(data []byte) error {
	type Alias StandalonePriceStagedChangesAppliedMessage
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
func (obj StandalonePriceStagedChangesAppliedMessage) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceStagedChangesAppliedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceStagedChangesApplied", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Value](ctp:api:type:StandalonePriceChangeValueAction) update action.
*
 */
type StandalonePriceValueChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The new value of the updated [StandalonePrice](ctp:api:type:StandalonePrice).
	Value Money `json:"value"`
	// Whether the new value was applied to the current or the staged representation of the StandalonePrice. Staged changes are stored on the [StagedStandalonePrice](ctp:api:type:StagedStandalonePrice).
	Staged bool `json:"staged"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StandalonePriceValueChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias StandalonePriceValueChangedMessage
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
func (obj StandalonePriceValueChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceValueChangedMessage
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceValueChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Create Store](/../api/projects/stores#create-store) request.
*
 */
type StoreCreatedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// The `name` of the [Store](ctp:api:type:Store) that was created.
	Name *LocalizedString `json:"name,omitempty"`
	// Languages of the [Store](ctp:api:type:Store) that was created. Languages are represented as [IETF language tags](https://en.wikipedia.org/wiki/IETF_language_tag).
	Languages []string `json:"languages"`
	// [Distribution Channels](ctp:api:type:ChannelRoleEnum) of the [Store](ctp:api:type:Store) that was created.
	DistributionChannels []ChannelReference `json:"distributionChannels"`
	// [Supply Channels](ctp:api:type:ChannelRoleEnum) of the [Store](ctp:api:type:Store) that was created.
	SupplyChannels []ChannelReference `json:"supplyChannels"`
	// [ProductSelectionSettings](ctp:api:type:ProductSelectionSetting) of the [Store](ctp:api:type:Store) that was created.
	ProductSelections []ProductSelectionSetting `json:"productSelections"`
	// [Custom Fields](ctp:api:type:CustomFields) on the [Store](ctp:api:type:Store) that was created.
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
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreCreated", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["languages"] == nil {
		delete(raw, "languages")
	}

	return json.Marshal(raw)

}

/**
*	Generated after a successful [Delete Store](/../api/projects/quote-requests#delete-quoterequest) request.
*
 */
type StoreDeletedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
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

/**
*	Generated after a successful [Add Distribution Channel](ctp:api:type:StoreAddDistributionChannelAction),
*	[Remove Distribution Channel](ctp:api:type:StoreRemoveDistributionChannelAction), or
*	[Set Distribution Channels](ctp:api:type:StoreSetDistributionChannelsAction) update action.
*
 */
type StoreDistributionChannelsChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Product distribution Channels that have been added to the [Store](ctp:api:type:Store).
	AddedDistributionChannels []ChannelReference `json:"addedDistributionChannels"`
	// Product distribution Channels that have been removed from the [Store](ctp:api:type:Store).
	RemovedDistributionChannels []ChannelReference `json:"removedDistributionChannels"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreDistributionChannelsChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias StoreDistributionChannelsChangedMessage
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
func (obj StoreDistributionChannelsChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias StoreDistributionChannelsChangedMessage
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreDistributionChannelsChanged", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["addedDistributionChannels"] == nil {
		delete(raw, "addedDistributionChannels")
	}

	if raw["removedDistributionChannels"] == nil {
		delete(raw, "removedDistributionChannels")
	}

	return json.Marshal(raw)

}

/**
*	Generated after a successful [Set Languages](ctp:api:type:StoreSetLanguagesAction) update action.
*
 */
type StoreLanguagesChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [Locales](ctp:api:type:Locale) added to the [Store](ctp:api:type:Store) after the [Set Languages](ctp:api:type:StoreSetLanguagesAction) update action.
	AddedLanguages []string `json:"addedLanguages"`
	// [Locales](ctp:api:type:Locale) removed from the [Store](ctp:api:type:Store) during the [Set Languages](ctp:api:type:StoreSetLanguagesAction) update action.
	RemovedLanguages []string `json:"removedLanguages"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreLanguagesChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias StoreLanguagesChangedMessage
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
func (obj StoreLanguagesChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias StoreLanguagesChangedMessage
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreLanguagesChanged", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["addedLanguages"] == nil {
		delete(raw, "addedLanguages")
	}

	if raw["removedLanguages"] == nil {
		delete(raw, "removedLanguages")
	}

	return json.Marshal(raw)

}

/**
*	Generated after a successful [Set Name](ctp:api:type:StoreSetNameAction) update action.
*
 */
type StoreNameSetMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Name of the [Store](ctp:api:type:Store) set during the [Set Name](ctp:api:type:StoreSetNameAction) update action.
	Name *LocalizedString `json:"name,omitempty"`
	// Names set for the [Store](ctp:api:type:Store) in different locales.
	NameAllLocales []LocalizedString `json:"nameAllLocales"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreNameSetMessage) UnmarshalJSON(data []byte) error {
	type Alias StoreNameSetMessage
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
func (obj StoreNameSetMessage) MarshalJSON() ([]byte, error) {
	type Alias StoreNameSetMessage
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreNameSet", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["nameAllLocales"] == nil {
		delete(raw, "nameAllLocales")
	}

	return json.Marshal(raw)

}

/**
*	Generated by a successful [Add Product Selection](ctp:api:type:StoreAddProductSelectionAction),
*	[Remove Product Selection](ctp:api:type:StoreRemoveProductSelectionAction),
*	[Set Product Selections](ctp:api:type:StoreSetProductSelectionsAction),
*	or [Change Product Selections Active](ctp:api:type:StoreChangeProductSelectionAction) update action.
*
 */
type StoreProductSelectionsChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// [ProductSelectionSettings](ctp:api:type:ProductSelectionSetting) that were added to the [Store](ctp:api:type:Store).
	AddedProductSelections []ProductSelectionSetting `json:"addedProductSelections"`
	// [ProductSelectionSettings](ctp:api:type:ProductSelectionSetting) that were removed from the [Store](ctp:api:type:Store).
	RemovedProductSelections []ProductSelectionSetting `json:"removedProductSelections"`
	// [ProductSelectionSettings](ctp:api:type:ProductSelectionSetting) that were updated in the [Store](ctp:api:type:Store).
	UpdatedProductSelections []ProductSelectionSetting `json:"updatedProductSelections"`
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["addedProductSelections"] == nil {
		delete(raw, "addedProductSelections")
	}

	if raw["removedProductSelections"] == nil {
		delete(raw, "removedProductSelections")
	}

	if raw["updatedProductSelections"] == nil {
		delete(raw, "updatedProductSelections")
	}

	return json.Marshal(raw)

}

/**
*	Generated after a successful [Add Supply Channel](ctp:api:type:StoreAddSupplyChannelAction),
*	[Remove Supply Channel](ctp:api:type:StoreRemoveSupplyChannelAction), or
*	[Set Supply Channels](ctp:api:type:StoreSetSupplyChannelsAction) update action.
*
 */
type StoreSupplyChannelsChangedMessage struct {
	// Unique identifier of the Message. Can be used to track which Messages have been processed.
	ID string `json:"id"`
	// Version of a resource. In case of Messages, this is always `1`.
	Version int `json:"version"`
	// Date and time (UTC) the Message was generated.
	CreatedAt time.Time `json:"createdAt"`
	// Value of `createdAt`.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Value of `createdBy`.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Message number in relation to other Messages for a given resource. The `sequenceNumber` of the next Message for the resource is the successor of the `sequenceNumber` of the current Message. Meaning, the `sequenceNumber` of the next Message equals the `sequenceNumber` of the current Message + 1.
	// `sequenceNumber` can be used to ensure that Messages are processed in the correct order for a particular resource.
	SequenceNumber int `json:"sequenceNumber"`
	// [Reference](ctp:api:type:Reference) to the resource on which the change or action was performed.
	Resource Reference `json:"resource"`
	// Version of the resource on which the change or action was performed.
	ResourceVersion int `json:"resourceVersion"`
	// User-provided identifiers of the resource, such as `key` or `externalId`. Only present if the resource has such identifiers.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Inventory supply Channels that have been added to the [Store](ctp:api:type:Store).
	AddedSupplyChannels []ChannelReference `json:"addedSupplyChannels"`
	// Inventory supply Channels that have been removed from the [Store](ctp:api:type:Store).
	RemovedSupplyChannels []ChannelReference `json:"removedSupplyChannels"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreSupplyChannelsChangedMessage) UnmarshalJSON(data []byte) error {
	type Alias StoreSupplyChannelsChangedMessage
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
func (obj StoreSupplyChannelsChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias StoreSupplyChannelsChangedMessage
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreSupplyChannelsChanged", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["addedSupplyChannels"] == nil {
		delete(raw, "addedSupplyChannels")
	}

	if raw["removedSupplyChannels"] == nil {
		delete(raw, "removedSupplyChannels")
	}

	return json.Marshal(raw)

}

/**
*	User-provided identifiers present on the resource for which the Message is created. The value of the identifier stored in the Message corresponds to the one that was set on the resource at the version shown in `resourceVersion`.
*
 */
type UserProvidedIdentifiers struct {
	// User-provided unique identifier of the resource.
	Key *string `json:"key,omitempty"`
	// User-provided unique identifier of the resource.
	ExternalId *string `json:"externalId,omitempty"`
	// User-provided unique identifier of an [Order](ctp:api:type:Order).
	OrderNumber *string `json:"orderNumber,omitempty"`
	// User-provided unique identifier of a [Customer](ctp:api:type:Customer).
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// Unique SKU of a [Product Variant](ctp:api:type:ProductVariant).
	Sku *string `json:"sku,omitempty"`
	// Unique identifier usually used in deep-link URLs for a [Product](ctp:api:type:Product). The value corresponds to the slug in the `current` [Product Projection](ctp:api:type:ProductProjection).
	Slug *LocalizedString `json:"slug,omitempty"`
	// Unique identifier of a [Custom Object](/../api/projects/custom-objects).
	ContainerAndKey *ContainerAndKey `json:"containerAndKey,omitempty"`
}

type MessagePayload interface{}

func mapDiscriminatorMessagePayload(input interface{}) (MessagePayload, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "BusinessUnitAddressAdded":
		obj := BusinessUnitAddressAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitAddressChanged":
		obj := BusinessUnitAddressChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitAddressRemoved":
		obj := BusinessUnitAddressRemovedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitAssociateAdded":
		obj := BusinessUnitAssociateAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitAssociateChanged":
		obj := BusinessUnitAssociateChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitAssociateRemoved":
		obj := BusinessUnitAssociateRemovedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitAssociatesSet":
		obj := BusinessUnitAssociatesSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitBillingAddressAdded":
		obj := BusinessUnitBillingAddressAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitBillingAddressRemoved":
		obj := BusinessUnitBillingAddressRemovedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitContactEmailSet":
		obj := BusinessUnitContactEmailSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitCreated":
		obj := BusinessUnitCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.BusinessUnit != nil {
			var err error
			obj.BusinessUnit, err = mapDiscriminatorBusinessUnit(obj.BusinessUnit)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "BusinessUnitDefaultBillingAddressSet":
		obj := BusinessUnitDefaultBillingAddressSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitDefaultShippingAddressSet":
		obj := BusinessUnitDefaultShippingAddressSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitDeleted":
		obj := BusinessUnitDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitNameChanged":
		obj := BusinessUnitNameChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitParentUnitChanged":
		obj := BusinessUnitParentUnitChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitShippingAddressAdded":
		obj := BusinessUnitShippingAddressAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitShippingAddressRemoved":
		obj := BusinessUnitShippingAddressRemovedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitStatusChanged":
		obj := BusinessUnitStatusChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitStoreAdded":
		obj := BusinessUnitStoreAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitStoreModeChanged":
		obj := BusinessUnitStoreModeChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitStoreRemoved":
		obj := BusinessUnitStoreRemovedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "BusinessUnitStoresSet":
		obj := BusinessUnitStoresSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
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
		if obj.VariantSelection != nil {
			var err error
			obj.VariantSelection, err = mapDiscriminatorProductVariantSelection(obj.VariantSelection)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ProductSelectionProductRemoved":
		obj := ProductSelectionProductRemovedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ProductSelectionVariantSelectionChanged":
		obj := ProductSelectionVariantSelectionChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.OldVariantSelection != nil {
			var err error
			obj.OldVariantSelection, err = mapDiscriminatorProductVariantSelection(obj.OldVariantSelection)
			if err != nil {
				return nil, err
			}
		}
		if obj.NewVariantSelection != nil {
			var err error
			obj.NewVariantSelection, err = mapDiscriminatorProductVariantSelection(obj.NewVariantSelection)
			if err != nil {
				return nil, err
			}
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
	case "QuoteCreated":
		obj := QuoteCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "QuoteDeleted":
		obj := QuoteDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "QuoteRequestCreated":
		obj := QuoteRequestCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "QuoteRequestDeleted":
		obj := QuoteRequestDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "QuoteRequestStateChanged":
		obj := QuoteRequestStateChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "QuoteRequestStateTransition":
		obj := QuoteRequestStateTransitionMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "QuoteStateChanged":
		obj := QuoteStateChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "QuoteStateTransition":
		obj := QuoteStateTransitionMessagePayload{}
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
	case "StagedQuoteCreated":
		obj := StagedQuoteCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StagedQuoteDeleted":
		obj := StagedQuoteDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StagedQuoteSellerCommentSet":
		obj := StagedQuoteSellerCommentSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StagedQuoteStateChanged":
		obj := StagedQuoteStateChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StagedQuoteStateTransition":
		obj := StagedQuoteStateTransitionMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StagedQuoteValidToSet":
		obj := StagedQuoteValidToSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StandalonePriceActiveChanged":
		obj := StandalonePriceActiveChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StandalonePriceCreated":
		obj := StandalonePriceCreatedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StandalonePriceDeleted":
		obj := StandalonePriceDeletedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StandalonePriceDiscountSet":
		obj := StandalonePriceDiscountSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StandalonePriceExternalDiscountSet":
		obj := StandalonePriceExternalDiscountSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StandalonePriceStagedChangesApplied":
		obj := StandalonePriceStagedChangesAppliedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StandalonePriceValueChanged":
		obj := StandalonePriceValueChangedMessagePayload{}
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
	case "StoreDistributionChannelsChanged":
		obj := StoreDistributionChannelsChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StoreLanguagesChanged":
		obj := StoreLanguagesChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StoreNameSet":
		obj := StoreNameSetMessagePayload{}
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
	case "StoreSupplyChannelsChanged":
		obj := StoreSupplyChannelsChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Generated after a successful [Add Address](ctp:api:type:BusinessUnitAddAddressAction) update action.
*
 */
type BusinessUnitAddressAddedMessagePayload struct {
	// The address that was added to the [Business Unit](ctp:api:type:BusinessUnit).
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAddressAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAddressAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAddressAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Address](ctp:api:type:BusinessUnitChangeAddressAction) update action.
*
 */
type BusinessUnitAddressChangedMessagePayload struct {
	// Updated address of the Business Unit.
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAddressChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAddressChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAddressChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Remove Address](ctp:api:type:BusinessUnitRemoveAddressAction) update action.
*
 */
type BusinessUnitAddressRemovedMessagePayload struct {
	// The address that was removed from the [Business Unit](ctp:api:type:BusinessUnit).
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAddressRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAddressRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAddressRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Add Associate](ctp:api:type:BusinessUnitAddAssociateAction) update action.
*
 */
type BusinessUnitAssociateAddedMessagePayload struct {
	// The [Associate](ctp:api:type:Associate) that was added to the [Business Unit](ctp:api:type:BusinessUnit).
	Associate Associate `json:"associate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAssociateAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAssociateAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAssociateAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Associate](ctp:api:type:BusinessUnitChangeAssociateAction) update action.
*
 */
type BusinessUnitAssociateChangedMessagePayload struct {
	// The [Associate](ctp:api:type:Associate) that was updated.
	Associate Associate `json:"associate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAssociateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAssociateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAssociateChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Remove Associate](ctp:api:type:BusinessUnitRemoveAssociateAction) update action.
*
 */
type BusinessUnitAssociateRemovedMessagePayload struct {
	// The [Associate](ctp:api:type:Associate) that was removed from the [Business Unit](ctp:api:type:BusinessUnit).
	Associate Associate `json:"associate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAssociateRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAssociateRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAssociateRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Associates](ctp:api:type:BusinessUnitSetAssociatesAction) update action.
*
 */
type BusinessUnitAssociatesSetMessagePayload struct {
	// The list of [Associates](ctp:api:type:Associate) that was updated on the [Business Unit](ctp:api:type:BusinessUnit).
	Associates []Associate `json:"associates"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitAssociatesSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitAssociatesSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitAssociatesSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Add Billing Address Identifier](ctp:api:type:BusinessUnitAddBillingAddressIdAction) update action.
*
 */
type BusinessUnitBillingAddressAddedMessagePayload struct {
	// The address that was added to the [Business Unit](ctp:api:type:BusinessUnit) as billing address.
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitBillingAddressAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitBillingAddressAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitBillingAddressAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Remove Billing Address Identifier](ctp:api:type:BusinessUnitRemoveBillingAddressIdAction) update action.
*
 */
type BusinessUnitBillingAddressRemovedMessagePayload struct {
	// The address that was removed from the billing addresses of the [Business Unit](ctp:api:type:BusinessUnit).
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitBillingAddressRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitBillingAddressRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitBillingAddressRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Contact Email](ctp:api:type:BusinessUnitSetContactEmailAction) update action.
*
 */
type BusinessUnitContactEmailSetMessagePayload struct {
	// The contact email that was updated on the [Business Unit](ctp:api:type:BusinessUnit).
	ContactEmail *string `json:"contactEmail,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitContactEmailSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitContactEmailSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitContactEmailSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Create Business Unit](/projects/business-units#create-businessunit) request.
*
 */
type BusinessUnitCreatedMessagePayload struct {
	// The [Business Unit](ctp:api:type:BusinessUnit) that was created.
	BusinessUnit BusinessUnit `json:"businessUnit"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *BusinessUnitCreatedMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias BusinessUnitCreatedMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.BusinessUnit != nil {
		var err error
		obj.BusinessUnit, err = mapDiscriminatorBusinessUnit(obj.BusinessUnit)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitCreated", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Default Billing Address](ctp:api:type:BusinessUnitSetDefaultBillingAddressAction) update action.
*
 */
type BusinessUnitDefaultBillingAddressSetMessagePayload struct {
	// The address that was set as the default billing address.
	Address *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitDefaultBillingAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitDefaultBillingAddressSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitDefaultBillingAddressSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Default Shipping Address](ctp:api:type:BusinessUnitSetDefaultShippingAddressAction) update action.
*
 */
type BusinessUnitDefaultShippingAddressSetMessagePayload struct {
	// The address that was set as the default shipping address.
	Address *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitDefaultShippingAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitDefaultShippingAddressSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitDefaultShippingAddressSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Delete Business Unit](/projects/business-units#delete-businessunit) request.
*
 */
type BusinessUnitDeletedMessagePayload struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitDeleted", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Name](ctp:api:type:BusinessUnitChangeNameAction) update action.
*
 */
type BusinessUnitNameChangedMessagePayload struct {
	// Updated name of the [Business Unit](ctp:api:type:BusinessUnit).
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitNameChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitNameChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitNameChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Parent Unit](ctp:api:type:BusinessUnitChangeParentUnitAction) update action.
*
 */
type BusinessUnitParentUnitChangedMessagePayload struct {
	// Parent unit of the [Business Unit](ctp:api:type:BusinessUnit) before the [Change Parent Unit](ctp:api:type:BusinessUnitChangeParentUnitAction) update action.
	OldParentUnit *BusinessUnitKeyReference `json:"oldParentUnit,omitempty"`
	// Parent unit of the [Business Unit](ctp:api:type:BusinessUnit) after the [Change Parent Unit](ctp:api:type:BusinessUnitChangeParentUnitAction) update action.
	NewParentUnit *BusinessUnitKeyReference `json:"newParentUnit,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitParentUnitChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitParentUnitChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitParentUnitChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Add Shipping Address Identifier](ctp:api:type:BusinessUnitAddShippingAddressIdAction) update action.
*
 */
type BusinessUnitShippingAddressAddedMessagePayload struct {
	// The address that was added to the [Business Unit](ctp:api:type:BusinessUnit) as shipping address.
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitShippingAddressAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitShippingAddressAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitShippingAddressAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Remove Shipping Address Identifier](ctp:api:type:BusinessUnitRemoveShippingAddressIdAction) update action.
*
 */
type BusinessUnitShippingAddressRemovedMessagePayload struct {
	// The address that was removed from shipping addresses of the [Business Unit](ctp:api:type:BusinessUnit).
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitShippingAddressRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitShippingAddressRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitShippingAddressRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Status](ctp:api:type:BusinessUnitChangeStatusAction) update action.
*
 */
type BusinessUnitStatusChangedMessagePayload struct {
	// Updated status of the [Business Unit](ctp:api:type:BusinessUnit).
	Active BusinessUnitStatus `json:"active"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitStatusChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitStatusChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitStatusChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Add Store](ctp:api:type:BusinessUnitAddStoreAction) update action.
*
 */
type BusinessUnitStoreAddedMessagePayload struct {
	// The [Store](ctp:api:type:Store) that was added to the [Business Unit](ctp:api:type:BusinessUnit).
	Store StoreKeyReference `json:"store"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitStoreAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitStoreAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitStoreAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Store Mode](ctp:api:type:BusinessUnitSetStoreModeAction) update action.
*
 */
type BusinessUnitStoreModeChangedMessagePayload struct {
	// [Stores](ctp:api:type:Store) of the [Business Unit](ctp:api:type:BusinessUnit) after the [Set Store Mode](ctp:api:type:BusinessUnitSetStoreModeAction) update action.
	Stores []StoreKeyReference `json:"stores"`
	// [BusinessUnitStoreMode](ctp:api:type:BusinessUnitStoreMode) of the Business Unit after the [Set Store Mode](ctp:api:type:BusinessUnitSetStoreModeAction) update action.
	StoreMode BusinessUnitStoreMode `json:"storeMode"`
	// [Stores](ctp:api:type:Store) of the [Business Unit](ctp:api:type:BusinessUnit) before the [Set Store Mode](ctp:api:type:BusinessUnitSetStoreModeAction) update action.
	OldStores []StoreKeyReference `json:"oldStores"`
	// [BusinessUnitStoreMode](ctp:api:type:BusinessUnitStoreMode) of the Business Unit before the [Set Store Mode](ctp:api:type:BusinessUnitSetStoreModeAction) update action.
	OldStoreMode BusinessUnitStoreMode `json:"oldStoreMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitStoreModeChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitStoreModeChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitStoreModeChanged", Alias: (*Alias)(&obj)})
}

type BusinessUnitStoreRemovedMessagePayload struct {
	// The [Store](ctp:api:type:Store) that was removed from the [Business Unit](ctp:api:type:BusinessUnit).
	Store StoreKeyReference `json:"store"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitStoreRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitStoreRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitStoreRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Stores](ctp:api:type:BusinessUnitSetStoresAction) update action.
*
 */
type BusinessUnitStoresSetMessagePayload struct {
	// [Stores](ctp:api:type:Store) of the [Business Unit](ctp:api:type:BusinessUnit) after the [Set Stores](ctp:api:type:BusinessUnitSetStoresAction) update action.
	Stores []StoreKeyReference `json:"stores"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BusinessUnitStoresSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias BusinessUnitStoresSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BusinessUnitStoresSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Create Category](/../api/projects/categories#create-category) request.
*
 */
type CategoryCreatedMessagePayload struct {
	// [Category](ctp:api:type:Category) that was created.
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

/**
*	Generated after a successful [Change Slug](ctp:api:type:CategoryChangeSlugAction) update action.
*
 */
type CategorySlugChangedMessagePayload struct {
	// The slug of the [Category](ctp:api:type:Category) after the [Change Slug](ctp:api:type:CategoryChangeSlugAction) update action.
	Slug LocalizedString `json:"slug"`
	// The slug of the [Category](ctp:api:type:Category) before the [Change Slug](ctp:api:type:CategoryChangeSlugAction) update action.
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

/**
*	Generated after a successful [Add Address](ctp:api:type:CustomerAddAddressAction) update action.
*
 */
type CustomerAddressAddedMessagePayload struct {
	// [Address](ctp:api:type:Address) that was added during the [Add Address](ctp:api:type:CustomerAddAddressAction) update action.
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

/**
*	Generated after a successful [Change Address](ctp:api:type:CustomerChangeAddressAction) update action.
*
 */
type CustomerAddressChangedMessagePayload struct {
	// [Address](ctp:api:type:Address) that was set during the [Change Address](ctp:api:type:CustomerChangeAddressAction) update action.
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

/**
*	Generated after a successful [Remove Address](ctp:api:type:CustomerRemoveAddressAction) update action.
*
 */
type CustomerAddressRemovedMessagePayload struct {
	// [Address](ctp:api:type:Address) that was removed during the [Remove Address](ctp:api:type:CustomerRemoveAddressAction) update action.
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

/**
*	Generated after a successful [Set Company Name](ctp:api:type:CustomerSetCompanyNameAction) update action.
*
 */
type CustomerCompanyNameSetMessagePayload struct {
	// The `companyName` that was set during the [Set Company Name](ctp:api:type:CustomerSetCompanyNameAction) update action.
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

/**
*	Generated after a successful [Create Customer](/../api/projects/customers#create-sign-up-customer) request.
*
 */
type CustomerCreatedMessagePayload struct {
	// [Customer](ctp:api:type:Customer) that was created.
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

/**
*	Generated after a successful [Set Date of Birth](ctp:api:type:CustomerSetDateOfBirthAction) update action.
*
 */
type CustomerDateOfBirthSetMessagePayload struct {
	// The `dateOfBirth` that was set during the [Set Date of Birth](ctp:api:type:CustomerSetDateOfBirthAction) update action.
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

/**
*	Generated after a successful [Delete Customer](/../api/projects/customers#delete-customer) request.
*
 */
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

/**
*	Generated after a successful [Change Email](ctp:api:type:CustomerChangeEmailAction) update action.
*
 */
type CustomerEmailChangedMessagePayload struct {
	// The `email` that was set during the [Change Email](ctp:api:type:CustomerChangeEmailAction) update action.
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

/**
*	Generated after a successful [Verify Customer's Email](/../api/projects/customers#verify-email-of-customer) request.
*
 */
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

/**
*	Generated after a successful [Set First Name](ctp:api:type:CustomerSetFirstNameAction) update action.
*
 */
type CustomerFirstNameSetMessagePayload struct {
	// The `firstName` that was set during the [Set First Name](ctp:api:type:CustomerSetFirstNameAction) update action.
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

/**
*	Generated after a successful [Set Customer Group](ctp:api:type:CustomerSetCustomerGroupAction) update action.
*
 */
type CustomerGroupSetMessagePayload struct {
	// [Customer Group](ctp:api:type:CustomerGroup) that was set during the [Set Customer Group](ctp:api:type:CustomerSetCustomerGroupAction) update action.
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

/**
*	Generated after a successful [Set Last Name](ctp:api:type:CustomerSetLastNameAction) update action.
*
 */
type CustomerLastNameSetMessagePayload struct {
	// The `lastName` that was set during the [Set Last Name](ctp:api:type:CustomerSetLastNameAction) update action.
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

/**
*	Generated after a successful [Reset Customer's Password](/../api/projects/customers#reset-password-of-customer), [Reset Customer's Password in a Store](/../api/projects/customers#reset-password-of-customer-in-store), [Change Customer's Password](/../api/projects/customers#change-password-of-customer), or [Change Customer's Password in a Store](/../api/projects/customers#change-password-of-customer-in-store) request. This Message is also produced during equivalent requests to the [My Customer Profile](/../api/projects/me-profile) endpoint.
*
 */
type CustomerPasswordUpdatedMessagePayload struct {
	// Whether the Customer's password was updated during the [Reset password](/../api/projects/customers#password-reset-of-customer) or [Change password](/../api/projects/customers#change-password-of-customer) flow.
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

/**
*	Generated after a successful [Set Title](ctp:api:type:CustomerSetTitleAction) update action.
*
 */
type CustomerTitleSetMessagePayload struct {
	// The `title` that was set during the [Set Title](ctp:api:type:CustomerSetTitleAction) update action.
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

/**
*	Generated after a successful [Create InventoryEntry](/../api/projects/inventory#create-inventoryentry) request.
*
 */
type InventoryEntryCreatedMessagePayload struct {
	// [InventoryEntry](ctp:api:type:InventoryEntry) that was created.
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

/**
*	Generated after a successful [Delete InventoryEntry](/../api/projects/inventory#delete-inventoryentry) request.
*
 */
type InventoryEntryDeletedMessagePayload struct {
	// The `sku` of the [InventoryEntry](ctp:api:type:InventoryEntry) that was deleted.
	Sku string `json:"sku"`
	// [Reference](ctp:api:type:Reference) to the [Channel](ctp:api:type:Channel) where the [InventoryEntry](ctp:api:type:InventoryEntry) was deleted.
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

/**
*	Generated after a successful [Add Quantity](ctp:api:type:InventoryEntryAddQuantityAction), [Remove Quantity](ctp:api:type:InventoryEntryRemoveQuantityAction) or [Change Quantity](ctp:api:type:InventoryEntryChangeQuantityAction) update action.
*	Inventory changes as a result of [Order creation](/../api/projects/orders#create-order) do not trigger this message.
*
 */
type InventoryEntryQuantitySetMessagePayload struct {
	// Quantity on stock for the [InventoryEntry](ctp:api:type:InventoryEntry) before the quantity was updated.
	OldQuantityOnStock int `json:"oldQuantityOnStock"`
	// Quantity on stock for the [InventoryEntry](ctp:api:type:InventoryEntry) after the quantity was updated.
	NewQuantityOnStock int `json:"newQuantityOnStock"`
	// Available quantity for the [InventoryEntry](ctp:api:type:InventoryEntry) before the quantity was updated.
	OldAvailableQuantity int `json:"oldAvailableQuantity"`
	// Available quantity for the [InventoryEntry](ctp:api:type:InventoryEntry) after the quantity was updated.
	NewAvailableQuantity int `json:"newAvailableQuantity"`
	// [Reference](ctp:api:type:Reference) to the [Channel](ctp:api:type:Channel) where the [InventoryEntry](ctp:api:type:InventoryEntry) quantity was set.
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
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
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
	case "OrderCustomLineItemAdded":
		obj := OrderCustomLineItemAddedMessagePayload{}
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
	case "OrderCustomLineItemQuantityChanged":
		obj := OrderCustomLineItemQuantityChangedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "OrderCustomLineItemRemoved":
		obj := OrderCustomLineItemRemovedMessagePayload{}
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
		return obj, nil
	case "OrderPaymentStateChanged":
		obj := OrderPaymentStateChangedMessagePayload{}
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
	case "ReturnInfoAdded":
		obj := ReturnInfoAddedMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ReturnInfoSet":
		obj := ReturnInfoSetMessagePayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Generated after a successful [Transition Custom Line Item State](ctp:api:type:OrderTransitionCustomLineItemStateAction) update action.
*
 */
type CustomLineItemStateTransitionMessagePayload struct {
	// Unique identifier of the [Custom Line Item](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
	// Date and time (UTC) when the transition of the [Custom Line Item](ctp:api:type:CustomLineItem) [State](ctp:api:type:State) was performed.
	TransitionDate time.Time `json:"transitionDate"`
	// Number of [Custom Line Items](ctp:api:type:CustomLineItem) for which the [State](ctp:api:type:State) was transitioned.
	Quantity int `json:"quantity"`
	// [State](ctp:api:type:State) the [Custom Line Item](ctp:api:type:CustomLineItem) was transitioned from.
	FromState StateReference `json:"fromState"`
	// [State](ctp:api:type:State) the [Custom Line Item](ctp:api:type:CustomLineItem) was transitioned to.
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

/**
*	Generated after a successful [Add Delivery](ctp:api:type:OrderAddDeliveryAction) update action.
*
 */
type DeliveryAddedMessagePayload struct {
	// [Delivery](ctp:api:type:Delivery) that was added to the [Order](ctp:api:type:Order). The [Delivery](ctp:api:type:Delivery) in the Message body does not contain [Parcels](ctp:api:type:Parcel) if those were part of the initial [Add Delivery](ctp:api:type:OrderAddDeliveryAction) update action. In that case, the update action produces an additional [ParcelAddedToDelivery](ctp:api:type:ParcelAddedToDeliveryMessage) Message containing information about the [Parcels](ctp:api:type:Parcel).
	Delivery Delivery `json:"delivery"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Set Delivery Address](ctp:api:type:OrderSetDeliveryAddressAction) update action.
*
 */
type DeliveryAddressSetMessagePayload struct {
	// Unique identifier of the [Parcel](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// [Address](ctp:api:type:Address) after the [Set Delivery Address](ctp:api:type:OrderSetDeliveryAddressAction) update action.
	Address *Address `json:"address,omitempty"`
	// [Address](ctp:api:type:Address) before the [Set Delivery Address](ctp:api:type:OrderSetDeliveryAddressAction) update action.
	OldAddress *Address `json:"oldAddress,omitempty"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Set Delivery Items](ctp:api:type:OrderSetDeliveryItemsAction) update action.
*
 */
type DeliveryItemsUpdatedMessagePayload struct {
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// [Delivery Items](ctp:api:type:DeliveryItem) after the [Set Delivery Items](ctp:api:type:OrderSetDeliveryItemsAction) update action.
	Items []DeliveryItem `json:"items"`
	// [Delivery Items](ctp:api:type:DeliveryItem) before the [Set Delivery Items](ctp:api:type:OrderSetDeliveryItemsAction) update action.
	OldItems []DeliveryItem `json:"oldItems"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Remove Delivery](ctp:api:type:OrderRemoveDeliveryAction) update action.
*
 */
type DeliveryRemovedMessagePayload struct {
	// The [Delivery](ctp:api:type:Delivery) that was removed from the [Order](ctp:api:type:Order).
	Delivery Delivery `json:"delivery"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Transition Line Item State](ctp:api:type:OrderTransitionLineItemStateAction) update action.
*
 */
type LineItemStateTransitionMessagePayload struct {
	// Unique identifier of the [Line Item](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
	// Date and time (UTC) when the transition of the [Line Item](ctp:api:type:LineItem) [State](ctp:api:type:State) was performed.
	TransitionDate time.Time `json:"transitionDate"`
	// Number of [Line Items](ctp:api:type:LineItem) for which the [State](ctp:api:type:State) was transitioned.
	Quantity int `json:"quantity"`
	// [State](ctp:api:type:State) the [Line Item](ctp:api:type:LineItem) was transitioned from.
	FromState StateReference `json:"fromState"`
	// [State](ctp:api:type:State) the [Line Item](ctp:api:type:LineItem) was transitioned to.
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

/**
*	Generated after a successful [Set Billing Address](ctp:api:type:OrderSetBillingAddressAction) update action.
*
 */
type OrderBillingAddressSetMessagePayload struct {
	// Billing address on the Order after the [Set Billing Address](ctp:api:type:OrderSetBillingAddressAction) update action.
	Address *Address `json:"address,omitempty"`
	// Billing address on the Order before the [Set Billing Address](ctp:api:type:OrderSetBillingAddressAction) update action.
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

/**
*	Generated after a successful [Create Order](/../api/projects/orders#create-order) request.
*
 */
type OrderCreatedMessagePayload struct {
	// [Order](ctp:api:type:Order) that was created.
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

/**
*	Generated after a successful [Add Custom Line Item](ctp:api:type:StagedOrderAddCustomLineItemAction) update action.
*
 */
type OrderCustomLineItemAddedMessagePayload struct {
	// [Custom Line Item](ctp:api:type:CustomLineItem) that was added to the [Order](ctp:api:type:Order).
	CustomLineItem CustomLineItem `json:"customLineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCustomLineItemAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomLineItemAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful recalculation of a Discount on a [Custom Line Item](ctp:api:type:CustomLineItem).
*
 */
type OrderCustomLineItemDiscountSetMessagePayload struct {
	// Unique identifier for the [Custom Line Item](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
	// Array of [DiscountedLineItemPriceForQuantity](ctp:api:type:DiscountedLineItemPriceForQuantity) after the Discount recalculation.
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	// [TaxedItemPrice](ctp:api:type:TaxedItemPrice) of the [Custom Line Item](ctp:api:type:CustomLineItem) after the Discount recalculation.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
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

/**
*	Generated after a successful [Change Custom Line Item Quantity](ctp:api:type:StagedOrderChangeCustomLineItemQuantityAction) update action.
*
 */
type OrderCustomLineItemQuantityChangedMessagePayload struct {
	// Unique identifier of the [Custom Line Item](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
	// [Custom Line Item](ctp:api:type:CustomLineItem) quantity after the [Change Custom Line Item Quantity](ctp:api:type:StagedOrderChangeCustomLineItemQuantityAction) update action.
	Quantity int `json:"quantity"`
	// [Custom Line Item](ctp:api:type:CustomLineItem) quantity before the [Change Custom Line Item Quantity](ctp:api:type:StagedOrderChangeCustomLineItemQuantityAction) update action.
	OldQuantity int `json:"oldQuantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCustomLineItemQuantityChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemQuantityChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomLineItemQuantityChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Remove Custom Line Item](ctp:api:type:StagedOrderRemoveCustomLineItemAction) update action.
*
 */
type OrderCustomLineItemRemovedMessagePayload struct {
	// Unique identifier of the [Custom Line Item](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
	// [Custom Line Item](ctp:api:type:CustomLineItem) that was removed from the [Order](ctp:api:type:Order).
	CustomLineItem CustomLineItem `json:"customLineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderCustomLineItemRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemRemovedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "OrderCustomLineItemRemoved", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Customer Email](ctp:api:type:OrderSetCustomerEmailAction) update action.
*
 */
type OrderCustomerEmailSetMessagePayload struct {
	// Email address on the [Order](ctp:api:type:Order) after the [Set Customer Email](ctp:api:type:OrderSetCustomerEmailAction) update action.
	Email *string `json:"email,omitempty"`
	// Email address on the [Order](ctp:api:type:Order) before the [Set Customer Email](ctp:api:type:OrderSetCustomerEmailAction) update action.
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

/**
*	Generated after a successful [Set Customer Group](ctp:api:type:StagedOrderSetCustomerGroupAction) update action.
*
 */
type OrderCustomerGroupSetMessagePayload struct {
	// [CustomerGroup](ctp:api:type:CustomerGroup) on the [Order](ctp:api:type:Order) after the [Set Customer Group](ctp:api:type:StagedOrderSetCustomerGroupAction) update action.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) on the [Order](ctp:api:type:Order) before the [Set Customer Group](ctp:api:type:StagedOrderSetCustomerGroupAction) update action.
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

/**
*	Generated after a successful [Set Customer Id](ctp:api:type:OrderSetCustomerIdAction) update action.
*
 */
type OrderCustomerSetMessagePayload struct {
	// [Customer](ctp:api:type:Customer) on the [Order](ctp:api:type:Order) after the [Set Customer Id](ctp:api:type:OrderSetCustomerIdAction) update action.
	Customer *CustomerReference `json:"customer,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) on the [Order](ctp:api:type:Order) after the [Set Customer Id](ctp:api:type:OrderSetCustomerIdAction) update action.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [Customer](ctp:api:type:Customer) on the [Order](ctp:api:type:Order) before the [Set Customer Id](ctp:api:type:OrderSetCustomerIdAction) update action.
	OldCustomer *CustomerReference `json:"oldCustomer,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) on the [Order](ctp:api:type:Order) before the [Set Customer Id](ctp:api:type:OrderSetCustomerIdAction) update action.
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

/**
*	Generated after a successful [Delete Order](/../api/projects/orders#delete-order) request.
*
 */
type OrderDeletedMessagePayload struct {
	// [Order](ctp:api:type:Order) that has been deleted.
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

/**
*	Generated after a successful [Add Discount Code](ctp:api:type:StagedOrderAddDiscountCodeAction) update action.
*
 */
type OrderDiscountCodeAddedMessagePayload struct {
	// [DiscountCode](ctp:api:type:DiscountCode) that was added.
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

/**
*	Generated after a successful [Remove Discount Code](ctp:api:type:StagedOrderRemoveDiscountCodeAction) update action.
*
 */
type OrderDiscountCodeRemovedMessagePayload struct {
	// [DiscountCode](ctp:api:type:DiscountCode) that was removed.
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

/**
*	Generated after the [DiscountCodeState](ctp:api:type:DiscountCodeState) changes due to a [recalculation](/../api/projects/carts#recalculate).
*
 */
type OrderDiscountCodeStateSetMessagePayload struct {
	// [DiscountCode](ctp:api:type:DiscountCode) that changed due to the recalculation.
	DiscountCode DiscountCodeReference `json:"discountCode"`
	// [DiscountCodeState](ctp:api:type:DiscountCodeState) after the recalculation.
	State DiscountCodeState `json:"state"`
	// [DiscountCodeState](ctp:api:type:DiscountCodeState) before the recalculation.
	OldState *DiscountCodeState `json:"oldState,omitempty"`
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

/**
*	Generated after a successfully applying an [OrderEdit](/../api/projects/order-edits#apply-an-orderedit).
*
 */
type OrderEditAppliedMessagePayload struct {
	// [OrderEdit](ctp:api:type:OrderEdit) that was applied.
	Edit OrderEdit `json:"edit"`
	// Information about a successfully applied [OrderEdit](ctp:api:type:OrderEdit).
	Result OrderEditApplied `json:"result"`
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

/**
*	Generated after a successful [Order Import](/../api/projects/orders-import#create-an-order-by-import).
*
 */
type OrderImportedMessagePayload struct {
	// [Order](ctp:api:type:Order) that was imported.
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

/**
*	Generated after a successful [Add Line Item](ctp:api:type:StagedOrderAddLineItemAction) update action.
*
 */
type OrderLineItemAddedMessagePayload struct {
	// [Line Item](ctp:api:type:LineItem) that was added to the [Order](ctp:api:type:Order).
	LineItem LineItem `json:"lineItem"`
	// Quantity of [Line Items](ctp:api:type:LineItem) that were added to the [Order](ctp:api:type:Order).
	AddedQuantity int `json:"addedQuantity"`
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

/**
*	Generated after a successful recalculation of a Discount on a [Line Item](ctp:api:type:LineItem).
*
 */
type OrderLineItemDiscountSetMessagePayload struct {
	// Unique identifier for the [Line Item](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
	// Array of [DiscountedLineItemPriceForQuantity](ctp:api:type:DiscountedLineItemPriceForQuantity) after the Discount recalculation.
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	// Total Price of the [Line Item](ctp:api:type:LineItem) after the Discount recalculation.
	TotalPrice Money `json:"totalPrice"`
	// [TaxedItemPrice](ctp:api:type:TaxedItemPrice) of the [Line Item](ctp:api:type:LineItem) after the Discount recalculation.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
	// Taxed price of the Shipping Methods in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode)..
	TaxedPricePortions []MethodTaxedPrice `json:"taxedPricePortions"`
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

/**
*	Generated after a successful [Set Line Item Distribution Channel](/../api/projects/order-edits#set-lineitem-distributionchannel) update action.
*
 */
type OrderLineItemDistributionChannelSetMessagePayload struct {
	// Unique identifier of the [Line Item](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
	// [Distribution Channel](ctp:api:type:Channel) that was set.
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

/**
*	Generated after a successful [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
*
 */
type OrderLineItemRemovedMessagePayload struct {
	// Unique identifier of the [Line Item](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
	// Quantity of [Line Items](ctp:api:type:LineItem) that were removed during the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	RemovedQuantity int `json:"removedQuantity"`
	// [Line Item](ctp:api:type:LineItem) quantity after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	NewQuantity int `json:"newQuantity"`
	// [ItemStates](ctp:api:type:ItemState) after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	NewState []ItemState `json:"newState"`
	// `totalPrice` of the [Order](ctp:api:type:Order) after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	NewTotalPrice CentPrecisionMoney `json:"newTotalPrice"`
	// [TaxedItemPrice](ctp:api:type:TaxedItemPrice) of the [Order](ctp:api:type:Order) after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	NewTaxedPrice *TaxedItemPrice `json:"newTaxedPrice,omitempty"`
	// [Price](ctp:api:type:Price) of the [Order](ctp:api:type:Order) after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	NewPrice *Price `json:"newPrice,omitempty"`
	// [Shipping Details](ctp:api:type:ItemShippingDetails) of the [Order](ctp:api:type:Order) after the [Remove Line Item](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
	NewShippingDetail *ItemShippingDetails `json:"newShippingDetail,omitempty"`
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

/**
*	Generated after a successful [Add Payment](ctp:api:type:OrderAddPaymentAction) update action or when a [Payment](ctp:api:type:Payment) is added via [Order Edits](ctp:api:type:StagedOrderAddPaymentAction).
*
 */
type OrderPaymentAddedMessagePayload struct {
	// [Payment](ctp:api:type:Payment) that was added to the [Order](ctp:api:type:Order).
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

/**
*	Generated after a successful [Change Payment State](ctp:api:type:OrderChangePaymentStateAction) update action.
*
 */
type OrderPaymentStateChangedMessagePayload struct {
	// [PaymentState](ctp:api:type:PaymentState) after the [Change Payment State](ctp:api:type:OrderChangePaymentStateAction) update action.
	PaymentState PaymentState `json:"paymentState"`
	// [PaymentState](ctp:api:type:PaymentState) before the [Change Payment State](ctp:api:type:OrderChangePaymentStateAction) update action.
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

/**
*	Generated after a successful [Set Return Shipment State](ctp:api:type:OrderSetReturnShipmentStateAction) update action.
*
 */
type OrderReturnShipmentStateChangedMessagePayload struct {
	// Unique identifier of the [ReturnItem](ctp:api:type:ReturnItem).
	ReturnItemId string `json:"returnItemId"`
	// State of the [ReturnItem](ctp:api:type:ReturnItem) after the [Set Return Shipment State](ctp:api:type:OrderSetReturnShipmentStateAction) update action.
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

/**
*	Generated after a successful [Change Shipment State](ctp:api:type:OrderChangeShipmentStateAction) update action.
*
 */
type OrderShipmentStateChangedMessagePayload struct {
	// [ShipmentState](ctp:api:type:ShipmentState) after the [Change Shipment State](ctp:api:type:OrderChangeShipmentStateAction) update action.
	ShipmentState ShipmentState `json:"shipmentState"`
	// [ShipmentState](ctp:api:type:ShipmentState) before the [Change Shipment State](ctp:api:type:OrderChangeShipmentStateAction) update action.
	OldShipmentState ShipmentState `json:"oldShipmentState"`
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

/**
*	Generated after a successful [Set Shipping Address](ctp:api:type:OrderSetShippingAddressAction) update action.
*
 */
type OrderShippingAddressSetMessagePayload struct {
	// Shipping address on the Order after the [Set Shipping Address](ctp:api:type:OrderSetShippingAddressAction) update action.
	Address *Address `json:"address,omitempty"`
	// Shipping address on the Order before the [Set Shipping Address](ctp:api:type:OrderSetShippingAddressAction) update action.
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

/**
*	Generated after a successful [Set Shipping Method](ctp:api:type:StagedOrderSetShippingMethodAction) and [Set Custom Shipping Method](ctp:api:type:StagedOrderSetCustomShippingMethodAction) update actions.
*
 */
type OrderShippingInfoSetMessagePayload struct {
	// [ShippingInfo](ctp:api:type:ShippingInfo) after the [Set Shipping Method](ctp:api:type:StagedOrderSetShippingMethodAction) or [Set Custom Shipping Method](ctp:api:type:StagedOrderSetCustomShippingMethodAction) update action.
	ShippingInfo *ShippingInfo `json:"shippingInfo,omitempty"`
	// [ShippingInfo](ctp:api:type:ShippingInfo) before the [Set Shipping Method](ctp:api:type:StagedOrderSetShippingMethodAction) or [Set Custom Shipping Method](ctp:api:type:StagedOrderSetCustomShippingMethodAction) update action.
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

/**
*	Generated after a successful [Set ShippingRateInput](ctp:api:type:StagedOrderSetShippingRateInputAction) update action.
*
 */
type OrderShippingRateInputSetMessagePayload struct {
	// [ShippingRateInput](ctp:api:type:ShippingRateInput) after the [Set ShippingRateInput](ctp:api:type:StagedOrderSetShippingRateInputAction) update action.
	ShippingRateInput ShippingRateInput `json:"shippingRateInput,omitempty"`
	// [ShippingRateInput](ctp:api:type:ShippingRateInput) before the [Set ShippingRateInput](ctp:api:type:StagedOrderSetShippingRateInputAction) update action.
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

/**
*	Generated after a successful [Change Order State](ctp:api:type:OrderChangeOrderStateAction) update action.
*
 */
type OrderStateChangedMessagePayload struct {
	// [OrderState](ctp:api:type:OrderState) after the [Change Order State](ctp:api:type:OrderChangeOrderStateAction) update action.
	OrderState OrderState `json:"orderState"`
	// [OrderState](ctp:api:type:OrderState) before the [Change Order State](ctp:api:type:OrderChangeOrderStateAction) update action.
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

/**
*	Generated after a successful [Transition State](ctp:api:type:OrderTransitionStateAction) update action.
*
 */
type OrderStateTransitionMessagePayload struct {
	// [OrderState](ctp:api:type:OrderState) after the [Transition State](ctp:api:type:OrderTransitionStateAction) update action.
	State StateReference `json:"state"`
	// [OrderState](ctp:api:type:OrderState) before the [Transition State](ctp:api:type:OrderTransitionStateAction) update action.
	OldState *StateReference `json:"oldState,omitempty"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:OrderTransitionStateAction) update action.
	Force bool `json:"force"`
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

/**
*	Generated after a successful [Set Store](ctp:api:type:OrderSetStoreAction) update action.
*
 */
type OrderStoreSetMessagePayload struct {
	// [Store](ctp:api:type:Store) that was set.
	Store *StoreKeyReference `json:"store,omitempty"`
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

/**
*	Generated after a successful [Add Parcel To Delivery](ctp:api:type:OrderAddParcelToDeliveryAction) update action.
*
 */
type ParcelAddedToDeliveryMessagePayload struct {
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	Delivery Delivery `json:"delivery"`
	// [Parcel](ctp:api:type:Parcel) that was added to the [Delivery](ctp:api:type:Delivery).
	Parcel Parcel `json:"parcel"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Set Parcel Items](ctp:api:type:OrderSetParcelItemsAction) update action.
*
 */
type ParcelItemsUpdatedMessagePayload struct {
	// Unique identifier of the [Parcel](ctp:api:type:Parcel).
	ParcelId string `json:"parcelId"`
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// [Delivery Items](ctp:api:type:DeliveryItem) after the [Set Parcel Items](ctp:api:type:OrderSetParcelItemsAction) update action.
	Items []DeliveryItem `json:"items"`
	// [Delivery Items](ctp:api:type:DeliveryItem) before the [Set Parcel Items](ctp:api:type:OrderSetParcelItemsAction) update action.
	OldItems []DeliveryItem `json:"oldItems"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Set Parcel Measurements](ctp:api:type:OrderSetParcelMeasurementsAction) update action.
*
 */
type ParcelMeasurementsUpdatedMessagePayload struct {
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// Unique identifier of the [Parcel](ctp:api:type:Parcel).
	ParcelId string `json:"parcelId"`
	// The [Parcel Measurements](ctp:api:type:ParcelMeasurements) that were set on the [Parcel](ctp:api:type:Parcel).
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Remove Parcel From Delivery](ctp:api:type:OrderRemoveParcelFromDeliveryAction) update action.
*
 */
type ParcelRemovedFromDeliveryMessagePayload struct {
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// [Parcel](ctp:api:type:Parcel) that was removed from the [Delivery](ctp:api:type:Delivery).
	Parcel Parcel `json:"parcel"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Set Parcel TrackingData](ctp:api:type:OrderSetParcelTrackingDataAction) update action.
*
 */
type ParcelTrackingDataUpdatedMessagePayload struct {
	// Unique identifier of the [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
	// Unique identifier of the [Parcel](ctp:api:type:Parcel).
	ParcelId string `json:"parcelId"`
	// The [Tracking Data](ctp:api:type:TrackingData) that was added to the [Parcel](ctp:api:type:Parcel).
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multi` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Generated after a successful [Create Payment](/../api/projects/payments#create-a-payment) request.
*
 */
type PaymentCreatedMessagePayload struct {
	// [Payment](ctp:api:type:Payment) that was created.
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

/**
*	Generated after a successful [Add InterfaceInteraction](ctp:api:type:PaymentAddInterfaceInteractionAction) update action.
*
 */
type PaymentInteractionAddedMessagePayload struct {
	// The interface interaction that was added to the [Payment](ctp:api:type:Payment).
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

/**
*	Generated after a successful [Set StatusInterfaceCode](ctp:api:type:PaymentSetStatusInterfaceCodeAction) update action.
*
 */
type PaymentStatusInterfaceCodeSetMessagePayload struct {
	// Unique identifier for the [Payment](ctp:api:type:Payment) for which the [Set StatusInterfaceCode](ctp:api:type:PaymentSetStatusInterfaceCodeAction) update action was applied.
	PaymentId string `json:"paymentId"`
	// The `interfaceCode` that was set during the [Set StatusInterfaceCode](ctp:api:type:PaymentSetStatusInterfaceCodeAction) update action.
	InterfaceCode *string `json:"interfaceCode,omitempty"`
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

/**
*	Generated after a successful [Transition State](ctp:api:type:PaymentTransitionStateAction) update action.
*
 */
type PaymentStatusStateTransitionMessagePayload struct {
	// [State](ctp:api:type:State) of the [Payment](ctp:api:type:Payment) after the [Transition State](ctp:api:type:PaymentTransitionStateAction) update action.
	State StateReference `json:"state"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Change Transaction State](ctp:api:type:PaymentChangeTransactionStateAction) update action.
	Force bool `json:"force"`
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

/**
*	Generated after a successful [Add Transaction](ctp:api:type:PaymentAddTransactionAction) update action.
*
 */
type PaymentTransactionAddedMessagePayload struct {
	// [Transaction](ctp:api:type:Transaction) that was added to the [Payment](ctp:api:type:Payment).
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

/**
*	Generated after a successful [Change Transaction State](ctp:api:type:PaymentChangeTransactionStateAction) update action.
*
 */
type PaymentTransactionStateChangedMessagePayload struct {
	// Unique identifier for the [Transaction](ctp:api:type:Transaction) for which the [Transaction State](ctp:api:type:TransactionState) changed.
	TransactionId string `json:"transactionId"`
	// [Transaction State](ctp:api:type:TransactionState) after the [Change Transaction State](ctp:api:type:PaymentChangeTransactionStateAction) update action.
	State TransactionState `json:"state"`
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

/**
*	Generated after a successful [Add To Category](ctp:api:type:ProductAddToCategoryAction) update action.
*
 */
type ProductAddedToCategoryMessagePayload struct {
	// [Category](ctp:api:type:Category) the [Product](ctp:api:type:Product) was added to.
	Category CategoryReference `json:"category"`
	// Whether the update was only applied to the staged [Product Projection](ctp:api:type:ProductProjection).
	Staged bool `json:"staged"`
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

/**
*	Generated after a successful [Create Product](/../api/projects/products#create-product) request.
*
 */
type ProductCreatedMessagePayload struct {
	// The staged [Product Projection](ctp:api:type:ProductProjection) of the [Product](ctp:api:type:Product) at the time of creation.
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

/**
*	Generated after a successful [Delete Product](/../api/projects/products#delete-product) request.
*
 */
type ProductDeletedMessagePayload struct {
	// List of image URLs that were removed during the [Delete Product](ctp:api:type:Product) request.
	RemovedImageUrls []string `json:"removedImageUrls"`
	// Current [Product Projection](ctp:api:type:ProductProjection) of the deleted [Product](ctp:api:type:Product).
	CurrentProjection *ProductProjection `json:"currentProjection,omitempty"`
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

/**
*	Generated after a successful [Add External Image](ctp:api:type:ProductAddExternalImageAction) update action or after the successful [upload of an image](/../api/projects/products#upload-product-image).
*
 */
type ProductImageAddedMessagePayload struct {
	// Unique identifier of the [Product Variant](ctp:api:type:ProductVariant) to which the [Image](ctp:api:type:Image) was added.
	VariantId int `json:"variantId"`
	// [Image](ctp:api:type:Image) that was added.
	Image Image `json:"image"`
	// Whether the update was only applied to the staged [Product Projection](ctp:api:type:ProductProjection).
	Staged bool `json:"staged"`
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

/**
*	Generated after a Price is updated due to a [Product Discount](ctp:api:type:ProductDiscount).
*
 */
type ProductPriceDiscountsSetMessagePayload struct {
	// Array containing details about the [Embedded Prices](ctp:api:type:Price) that were updated.
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

/**
*	Generated after a successful [Set Discounted Embedded Price](ctp:api:type:ProductSetDiscountedPriceAction) update action.
*
 */
type ProductPriceExternalDiscountSetMessagePayload struct {
	// Unique identifier of the [Product Variant](ctp:api:type:ProductVariant) for which the Discount was set.
	VariantId int `json:"variantId"`
	// Key of the [Product Variant](ctp:api:type:ProductVariant) for which the Discount was set.
	VariantKey *string `json:"variantKey,omitempty"`
	// SKU of the [Product Variant](ctp:api:type:ProductVariant) for which Discount was set.
	Sku *string `json:"sku,omitempty"`
	// Unique identifier of the [Embedded Price](ctp:api:type:Price).
	PriceId string `json:"priceId"`
	// Discounted Price for the [Product Variant](ctp:api:type:ProductVariant) for which Discount was set.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Whether the update was only applied to the staged [Product Projection](ctp:api:type:ProductProjection).
	Staged bool `json:"staged"`
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

/**
*	Generated after a successful [Publish](ctp:api:type:ProductPublishAction) update action.
*
 */
type ProductPublishedMessagePayload struct {
	// List of image URLs which were removed during the [Publish](ctp:api:type:ProductPublishAction) update action.
	RemovedImageUrls []string `json:"removedImageUrls"`
	// Current [Product Projection](ctp:api:type:ProductProjection) of the [Product](ctp:api:type:Product) at the time of creation.
	ProductProjection ProductProjection `json:"productProjection"`
	// [Publishing Scope](ctp:api:type:ProductPublishScope) that was used during the [Publish](ctp:api:type:ProductPublishAction) update action.
	Scope ProductPublishScope `json:"scope"`
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

/**
*	Generated after a successful [Remove From Category](ctp:api:type:ProductRemoveFromCategoryAction) update action.
*
 */
type ProductRemovedFromCategoryMessagePayload struct {
	// [Category](ctp:api:type:Category) the [Product](ctp:api:type:Product) was removed from.
	Category CategoryReference `json:"category"`
	// Whether the update was only applied to the staged [Product Projection](ctp:api:type:ProductProjection).
	Staged bool `json:"staged"`
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

/**
*	Generated after a successful [Revert Staged Changes](ctp:api:type:ProductRevertStagedChangesAction) update action.
*
 */
type ProductRevertedStagedChangesMessagePayload struct {
	// List of image URLs that were removed during the [Revert Staged Changes](ctp:api:type:ProductRevertStagedChangesAction) update action.
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

/**
*	Generated after a successful [Create Product Selection](/../api/projects/product-selections#create-product-selection) request.
*
 */
type ProductSelectionCreatedMessagePayload struct {
	// The `type` and `name` of the individual Product Selection.
	ProductSelection IndividualProductSelectionType `json:"productSelection"`
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

/**
*	Generated after a successful [Delete Product Selection](/../api/projects/product-selections#create-product-selection) request.
*
 */
type ProductSelectionDeletedMessagePayload struct {
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

/**
*	Generated after a successful [Add Product](ctp:api:type:ProductSelectionAddProductAction) update action.
*
 */
type ProductSelectionProductAddedMessagePayload struct {
	// [Product](ctp:api:type:Product) that was added to the [Product Selection](ctp:api:type:ProductSelection).
	Product ProductReference `json:"product"`
	// Product Variant Selection after the [Add Product](ctp:api:type:ProductSelectionAddProductAction) update action.
	VariantSelection ProductVariantSelection `json:"variantSelection"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionProductAddedMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionProductAddedMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.VariantSelection != nil {
		var err error
		obj.VariantSelection, err = mapDiscriminatorProductVariantSelection(obj.VariantSelection)
		if err != nil {
			return err
		}
	}

	return nil
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

/**
*	Generated after a successful [Remove Product](ctp:api:type:ProductSelectionRemoveProductAction) update action.
*
 */
type ProductSelectionProductRemovedMessagePayload struct {
	// [Product](ctp:api:type:Product) that was removed from the Product Selection.
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

/**
*	Generated after a successful [Set Variant Selection](ctp:api:type:ProductSelectionSetVariantSelectionAction) update action.
*
 */
type ProductSelectionVariantSelectionChangedMessagePayload struct {
	// [Product](ctp:api:type:Product) for which the Product Variant Selection changed.
	Product ProductReference `json:"product"`
	// Product Variant Selection before the [Set Variant Selection](ctp:api:type:ProductSelectionSetVariantSelectionAction) update action.
	OldVariantSelection ProductVariantSelection `json:"oldVariantSelection"`
	// Product Variant Selection after the [Set Variant Selection](ctp:api:type:ProductSelectionSetVariantSelectionAction) update action.
	NewVariantSelection ProductVariantSelection `json:"newVariantSelection"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSelectionVariantSelectionChangedMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias ProductSelectionVariantSelectionChangedMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.OldVariantSelection != nil {
		var err error
		obj.OldVariantSelection, err = mapDiscriminatorProductVariantSelection(obj.OldVariantSelection)
		if err != nil {
			return err
		}
	}
	if obj.NewVariantSelection != nil {
		var err error
		obj.NewVariantSelection, err = mapDiscriminatorProductVariantSelection(obj.NewVariantSelection)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionVariantSelectionChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionVariantSelectionChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ProductSelectionVariantSelectionChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Slug](ctp:api:type:ProductChangeSlugAction) update action.
*
 */
type ProductSlugChangedMessagePayload struct {
	// The slug of the [Product](ctp:api:type:Product) after the [Change Slug](ctp:api:type:ProductChangeSlugAction) update action.
	Slug LocalizedString `json:"slug"`
	// The slug of the [Product](ctp:api:type:Product) before the [Change Slug](ctp:api:type:ProductChangeSlugAction) update action.
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

/**
*	Generated after a successful [Transition State](ctp:api:type:ProductTransitionStateAction) update action.
*
 */
type ProductStateTransitionMessagePayload struct {
	// Product [State](ctp:api:type:State) after the [Transition State](ctp:api:type:ProductTransitionStateAction) update action.
	State StateReference `json:"state"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:ProductTransitionStateAction) update action.
	Force bool `json:"force"`
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

/**
*	Generated after a successful [Unpublish Product](ctp:api:type:ProductUnpublishAction) update action.
*
 */
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

/**
*	Generated after a successful [Add Product Variant](ctp:api:type:ProductAddVariantAction) update action.
*
 */
type ProductVariantAddedMessagePayload struct {
	// Unique identifier of the [Product Variant](ctp:api:type:ProductVariant) that was added.
	Variant ProductVariant `json:"variant"`
	// Whether the update was only applied to the staged [Product Projection](ctp:api:type:ProductProjection).
	Staged bool `json:"staged"`
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

/**
*	Generated after a successful [Remove Product Variant](ctp:api:type:ProductRemoveVariantAction) update action.
*
 */
type ProductVariantDeletedMessagePayload struct {
	// Unique identifier of the [Product Variant](ctp:api:type:ProductVariant) that was added.
	Variant *ProductVariant `json:"variant,omitempty"`
	// List of image URLs that were removed with the [Remove Product Variant](ctp:api:type:ProductRemoveVariantAction) update action.
	RemovedImageUrls []string `json:"removedImageUrls"`
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

/**
*	Generated after a successful [Create Quote](/../api/projects/quotes#create-quote) request.
*
 */
type QuoteCreatedMessagePayload struct {
	// [Quote](/../api/projects/quotes) that was created.
	Quote Quote `json:"quote"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias QuoteCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteCreated", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Delete Quote](/../api/projects/quotes#delete-quote) request.
*
 */
type QuoteDeletedMessagePayload struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias QuoteDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteDeleted", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Create Quote Request](/../api/projects/quote-requests#create-quoterequest) request.
*
 */
type QuoteRequestCreatedMessagePayload struct {
	// [Quote Request](/../api/projects/quote-requests) that was created.
	QuoteRequest QuoteRequest `json:"quoteRequest"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteRequestCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias QuoteRequestCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteRequestCreated", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Delete Quote Request](/../api/projects/quote-requests#delete-quoterequest) request.
*
 */
type QuoteRequestDeletedMessagePayload struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteRequestDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias QuoteRequestDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteRequestDeleted", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Quote Request State](ctp:api:type:QuoteRequestChangeQuoteRequestStateAction) update action.
*
 */
type QuoteRequestStateChangedMessagePayload struct {
	// State of the Quote Request after the [Change Quote Request State](ctp:api:type:QuoteRequestChangeQuoteRequestStateAction) update action.
	QuoteRequestState QuoteRequestState `json:"quoteRequestState"`
	// State of the Quote Request before the [Change Quote Request State](ctp:api:type:QuoteRequestChangeQuoteRequestStateAction) update action.
	OldQuoteRequestState QuoteRequestState `json:"oldQuoteRequestState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteRequestStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias QuoteRequestStateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteRequestStateChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Transition State](ctp:api:type:QuoteRequestTransitionStateAction) update action.
*
 */
type QuoteRequestStateTransitionMessagePayload struct {
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) after the [Transition State](ctp:api:type:QuoteRequestTransitionStateAction) update action.
	State StateReference `json:"state"`
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) before the [Transition State](ctp:api:type:QuoteRequestTransitionStateAction) update action.
	OldState *StateReference `json:"oldState,omitempty"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:QuoteRequestTransitionStateAction) update action.
	Force bool `json:"force"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteRequestStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias QuoteRequestStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteRequestStateTransition", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Quote State](ctp:api:type:QuoteChangeQuoteStateAction) update action.
*
 */
type QuoteStateChangedMessagePayload struct {
	// State of the Quote after the [Change Quote State](ctp:api:type:QuoteChangeQuoteStateAction) update action.
	QuoteState QuoteState `json:"quoteState"`
	// State of the Quote before the [Change Quote State](ctp:api:type:QuoteChangeQuoteStateAction) update action.
	OldQuoteState QuoteState `json:"oldQuoteState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias QuoteStateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteStateChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Transition State](ctp:api:type:QuoteTransitionStateAction) update action.
*
 */
type QuoteStateTransitionMessagePayload struct {
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) after the [Transition State](ctp:api:type:QuoteTransitionStateAction) update action.
	State StateReference `json:"state"`
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) before the [Transition State](ctp:api:type:QuoteTransitionStateAction) update action.
	OldState *StateReference `json:"oldState,omitempty"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:QuoteTransitionStateAction) update action.
	Force bool `json:"force"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias QuoteStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "QuoteStateTransition", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Add Return Info](ctp:api:type:OrderAddReturnInfoAction) update action.
*
 */
type ReturnInfoAddedMessagePayload struct {
	// The [ReturnInfo](ctp:api:type:ReturnInfo) that was added to the [Order](ctp:api:type:Order).
	ReturnInfo ReturnInfo `json:"returnInfo"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReturnInfoAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReturnInfoAddedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReturnInfoAdded", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Return Info](ctp:api:type:OrderSetReturnInfoAction) update action on [Orders](ctp:api:type:Order) and [Order Edits](ctp:api:type:OrderEdit).
*
 */
type ReturnInfoSetMessagePayload struct {
	// The [ReturnInfo](ctp:api:type:ReturnInfo) that was set on the [Order](ctp:api:type:Order) or [Order Edit](ctp:api:type:OrderEdit).
	ReturnInfo []ReturnInfo `json:"returnInfo"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReturnInfoSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReturnInfoSetMessagePayload
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReturnInfoSet", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["returnInfo"] == nil {
		delete(raw, "returnInfo")
	}

	return json.Marshal(raw)

}

/**
*	Generated after a successful [Create Review](/../api/projects/reviews#create-a-review) request.
*
 */
type ReviewCreatedMessagePayload struct {
	// [Review](ctp:api:type:Review) that was created.
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

/**
*	Generated after a successful [Set Rating](ctp:api:type:ReviewSetRatingAction) update action.
*
 */
type ReviewRatingSetMessagePayload struct {
	// The `rating` of the [Review](ctp:api:type:Review) before the [Set Rating](ctp:api:type:ReviewSetRatingAction) update action.
	OldRating *float64 `json:"oldRating,omitempty"`
	// The `rating` of the [Review](ctp:api:type:Review) after the [Set Rating](ctp:api:type:ReviewSetRatingAction) update action.
	NewRating *float64 `json:"newRating,omitempty"`
	// Whether the [Review](ctp:api:type:Review) was taken into account in the ratings statistics of the target.
	IncludedInStatistics bool `json:"includedInStatistics"`
	// [Reference](ctp:api:type:Reference) to the resource that the [Review](ctp:api:type:Review) belongs to.
	Target Reference `json:"target,omitempty"`
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

/**
*	Generated after a successful [Transition State](ctp:api:type:ReviewTransitionStateAction) update action.
*
 */
type ReviewStateTransitionMessagePayload struct {
	// [State](ctp:api:type:State) of the [Review](ctp:api:type:Review) before the [Transition State](ctp:api:type:ReviewTransitionStateAction) update action.
	OldState *StateReference `json:"oldState,omitempty"`
	// [State](ctp:api:type:State) of the [Review](ctp:api:type:Review) after the [Transition State](ctp:api:type:ReviewTransitionStateAction) update action.
	NewState StateReference `json:"newState"`
	// Whether the old [Review](ctp:api:type:Review) was taken into account in the rating statistics of the target before the state transition.
	OldIncludedInStatistics bool `json:"oldIncludedInStatistics"`
	// Whether the new [Review](ctp:api:type:Review) was taken into account in the rating statistics of the target after the state transition.
	NewIncludedInStatistics bool `json:"newIncludedInStatistics"`
	// [Reference](ctp:api:type:Reference) to the resource that the [Review](ctp:api:type:Review) belongs to.
	Target Reference `json:"target,omitempty"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:ReviewTransitionStateAction) update action.
	Force bool `json:"force"`
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
	// [Reference](/../api/types#reference) to a [Store](ctp:api:type:Store) by its key.
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

/**
*	Generated after a successful [Create Staged Quote](/../api/projects/staged-quotes#create-stagedquote) request.
*
 */
type StagedQuoteCreatedMessagePayload struct {
	// [Staged Quote](/../api/projects/staged-quotes) that was created.
	StagedQuote StagedQuote `json:"stagedQuote"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteCreated", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Delete Staged Quote](/../api/projects/staged-quotes#delete-stagedquote) request.
*
 */
type StagedQuoteDeletedMessagePayload struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteDeleted", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Seller Comment](ctp:api:type:StagedQuoteSetSellerCommentAction) update action.
*
 */
type StagedQuoteSellerCommentSetMessagePayload struct {
	// `sellerComment` on the [StagedQuote](ctp:api:type:StagedQuote) after a successful [Set Seller Comment](ctp:api:type:StagedQuoteSetSellerCommentAction) update action.
	SellerComment string `json:"sellerComment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteSellerCommentSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteSellerCommentSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteSellerCommentSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Staged Quote State](ctp:api:type:StagedQuoteChangeStagedQuoteStateAction) update action.
*
 */
type StagedQuoteStateChangedMessagePayload struct {
	// State of the Staged Quote after the [Change Staged Quote State](ctp:api:type:StagedQuoteChangeStagedQuoteStateAction) update action.
	StagedQuoteState StagedQuoteState `json:"stagedQuoteState"`
	// State of the Staged Quote before the [Change Staged Quote State](ctp:api:type:StagedQuoteChangeStagedQuoteStateAction) update action.
	OldStagedQuoteState StagedQuoteState `json:"oldStagedQuoteState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteStateChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteStateChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Transition State](ctp:api:type:StagedQuoteTransitionStateAction) update action.
*
 */
type StagedQuoteStateTransitionMessagePayload struct {
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) after the [Transition State](ctp:api:type:StagedQuoteTransitionStateAction) update action.
	State StateReference `json:"state"`
	// [State](ctp:api:type:State) of the [Quote](ctp:api:type:Quote) before the [Transition State](ctp:api:type:StagedQuoteTransitionStateAction) update action.
	OldState *StateReference `json:"oldState,omitempty"`
	// Whether [State](ctp:api:type:State) transition validations were turned off during the [Transition State](ctp:api:type:StagedQuoteTransitionStateAction) update action.
	Force bool `json:"force"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteStateTransitionMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteStateTransition", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Valid To](ctp:api:type:StagedQuoteSetValidToAction) update action.
*
 */
type StagedQuoteValidToSetMessagePayload struct {
	// Expiration date for the Staged Quote after the [Set Valid To](ctp:api:type:StagedQuoteSetValidToAction) update action.
	ValidTo time.Time `json:"validTo"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteValidToSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteValidToSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StagedQuoteValidToSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Active](ctp:api:types:StandalonePriceChangeActiveAction) update action.
*
 */
type StandalonePriceActiveChangedMessagePayload struct {
	// Value of the `active` field of the StandalonePrice after the [Change Active](ctp:api:types:StandalonePriceChangeActiveAction) update action.
	Active bool `json:"active"`
	// Value of the `active` field of the StandalonePrice before the [Change Active](ctp:api:types:StandalonePriceChangeActiveAction) update action.
	OldActive bool `json:"oldActive"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceActiveChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceActiveChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceActiveChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Create StandalonePrice](/../api/projects/standalone-prices#create-standaloneprice) request.
*
 */
type StandalonePriceCreatedMessagePayload struct {
	// [Standalone Price](ctp:api:type:StandalonePrice) that was created.
	StandalonePrice StandalonePrice `json:"standalonePrice"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceCreatedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceCreated", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Delete StandalonePrice](/../api/projects/standalone-prices#delete-standaloneprice) request.
*
 */
type StandalonePriceDeletedMessagePayload struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceDeletedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceDeleted", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a [Product Discount](ctp:api:type:ProductDiscount) is successfully applied to a StandalonePrice.
*
 */
type StandalonePriceDiscountSetMessagePayload struct {
	// The new `discounted` value of the updated [StandalonePrice](ctp:api:type:StandalonePrice).
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceDiscountSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceDiscountSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Set Discounted Price](ctp:api:type:StandalonePriceSetDiscountedPriceAction) update action.
*
 */
type StandalonePriceExternalDiscountSetMessagePayload struct {
	// The `discounted` value of the [StandalonePrice](ctp:api:type:StandalonePrice) after the [Set Discounted Price](ctp:api:type:StandalonePriceSetDiscountedPriceAction) update action.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceExternalDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceExternalDiscountSetMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceExternalDiscountSet", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Apply Staged Changes](ctp:api:types:StandalonePriceApplyStagedChangesAction) update action.
*
 */
type StandalonePriceStagedChangesAppliedMessagePayload struct {
	// Applied changes of the [StandalonePrice](/../api/projects/standalone-prices) after the [Apply Staged Changes](ctp:api:types:StandalonePriceApplyStagedChangesAction) update action.
	StagedChanges StagedStandalonePrice `json:"stagedChanges"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceStagedChangesAppliedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceStagedChangesAppliedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceStagedChangesApplied", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Change Value](ctp:api:type:StandalonePriceChangeValueAction) update action.
*
 */
type StandalonePriceValueChangedMessagePayload struct {
	// The new value of the updated [StandalonePrice](ctp:api:type:StandalonePrice).
	Value Money `json:"value"`
	// Whether the new value was applied to the current or the staged representation of the StandalonePrice. Staged changes are stored on the [StagedStandalonePrice](ctp:api:type:StagedStandalonePrice).
	Staged bool `json:"staged"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceValueChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceValueChangedMessagePayload
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StandalonePriceValueChanged", Alias: (*Alias)(&obj)})
}

/**
*	Generated after a successful [Create Store](/../api/projects/stores#create-store) request.
*
 */
type StoreCreatedMessagePayload struct {
	// The `name` of the [Store](ctp:api:type:Store) that was created.
	Name *LocalizedString `json:"name,omitempty"`
	// Languages of the [Store](ctp:api:type:Store) that was created. Languages are represented as [IETF language tags](https://en.wikipedia.org/wiki/IETF_language_tag).
	Languages []string `json:"languages"`
	// [Distribution Channels](ctp:api:type:ChannelRoleEnum) of the [Store](ctp:api:type:Store) that was created.
	DistributionChannels []ChannelReference `json:"distributionChannels"`
	// [Supply Channels](ctp:api:type:ChannelRoleEnum) of the [Store](ctp:api:type:Store) that was created.
	SupplyChannels []ChannelReference `json:"supplyChannels"`
	// [ProductSelectionSettings](ctp:api:type:ProductSelectionSetting) of the [Store](ctp:api:type:Store) that was created.
	ProductSelections []ProductSelectionSetting `json:"productSelections"`
	// [Custom Fields](ctp:api:type:CustomFields) on the [Store](ctp:api:type:Store) that was created.
	Custom *CustomFields `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StoreCreatedMessagePayload
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreCreated", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["languages"] == nil {
		delete(raw, "languages")
	}

	return json.Marshal(raw)

}

/**
*	Generated after a successful [Delete Store](/../api/projects/quote-requests#delete-quoterequest) request.
*
 */
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

/**
*	Generated after a successful [Add Distribution Channel](ctp:api:type:StoreAddDistributionChannelAction),
*	[Remove Distribution Channel](ctp:api:type:StoreRemoveDistributionChannelAction), or
*	[Set Distribution Channels](ctp:api:type:StoreSetDistributionChannelsAction) update action.
*
 */
type StoreDistributionChannelsChangedMessagePayload struct {
	// Product distribution Channels that have been added to the [Store](ctp:api:type:Store).
	AddedDistributionChannels []ChannelReference `json:"addedDistributionChannels"`
	// Product distribution Channels that have been removed from the [Store](ctp:api:type:Store).
	RemovedDistributionChannels []ChannelReference `json:"removedDistributionChannels"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreDistributionChannelsChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StoreDistributionChannelsChangedMessagePayload
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreDistributionChannelsChanged", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["addedDistributionChannels"] == nil {
		delete(raw, "addedDistributionChannels")
	}

	if raw["removedDistributionChannels"] == nil {
		delete(raw, "removedDistributionChannels")
	}

	return json.Marshal(raw)

}

/**
*	Generated after a successful [Set Languages](ctp:api:type:StoreSetLanguagesAction) update action.
*
 */
type StoreLanguagesChangedMessagePayload struct {
	// [Locales](ctp:api:type:Locale) added to the [Store](ctp:api:type:Store) after the [Set Languages](ctp:api:type:StoreSetLanguagesAction) update action.
	AddedLanguages []string `json:"addedLanguages"`
	// [Locales](ctp:api:type:Locale) removed from the [Store](ctp:api:type:Store) during the [Set Languages](ctp:api:type:StoreSetLanguagesAction) update action.
	RemovedLanguages []string `json:"removedLanguages"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreLanguagesChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StoreLanguagesChangedMessagePayload
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreLanguagesChanged", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["addedLanguages"] == nil {
		delete(raw, "addedLanguages")
	}

	if raw["removedLanguages"] == nil {
		delete(raw, "removedLanguages")
	}

	return json.Marshal(raw)

}

/**
*	Generated after a successful [Set Name](ctp:api:type:StoreSetNameAction) update action.
*
 */
type StoreNameSetMessagePayload struct {
	// Name of the [Store](ctp:api:type:Store) set during the [Set Name](ctp:api:type:StoreSetNameAction) update action.
	Name *LocalizedString `json:"name,omitempty"`
	// Names set for the [Store](ctp:api:type:Store) in different locales.
	NameAllLocales []LocalizedString `json:"nameAllLocales"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreNameSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StoreNameSetMessagePayload
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreNameSet", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["nameAllLocales"] == nil {
		delete(raw, "nameAllLocales")
	}

	return json.Marshal(raw)

}

/**
*	Generated by a successful [Add Product Selection](ctp:api:type:StoreAddProductSelectionAction),
*	[Remove Product Selection](ctp:api:type:StoreRemoveProductSelectionAction),
*	[Set Product Selections](ctp:api:type:StoreSetProductSelectionsAction),
*	or [Change Product Selections Active](ctp:api:type:StoreChangeProductSelectionAction) update action.
*
 */
type StoreProductSelectionsChangedMessagePayload struct {
	// [ProductSelectionSettings](ctp:api:type:ProductSelectionSetting) that were added to the [Store](ctp:api:type:Store).
	AddedProductSelections []ProductSelectionSetting `json:"addedProductSelections"`
	// [ProductSelectionSettings](ctp:api:type:ProductSelectionSetting) that were removed from the [Store](ctp:api:type:Store).
	RemovedProductSelections []ProductSelectionSetting `json:"removedProductSelections"`
	// [ProductSelectionSettings](ctp:api:type:ProductSelectionSetting) that were updated in the [Store](ctp:api:type:Store).
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["addedProductSelections"] == nil {
		delete(raw, "addedProductSelections")
	}

	if raw["removedProductSelections"] == nil {
		delete(raw, "removedProductSelections")
	}

	if raw["updatedProductSelections"] == nil {
		delete(raw, "updatedProductSelections")
	}

	return json.Marshal(raw)

}

/**
*	Generated after a successful [Add Supply Channel](ctp:api:type:StoreAddSupplyChannelAction),
*	[Remove Supply Channel](ctp:api:type:StoreRemoveSupplyChannelAction), or
*	[Set Supply Channels](ctp:api:type:StoreSetSupplyChannelsAction) update action.
*
 */
type StoreSupplyChannelsChangedMessagePayload struct {
	// Inventory supply Channels that have been added to the [Store](ctp:api:type:Store).
	AddedSupplyChannels []ChannelReference `json:"addedSupplyChannels"`
	// Inventory supply Channels that have been removed from the [Store](ctp:api:type:Store).
	RemovedSupplyChannels []ChannelReference `json:"removedSupplyChannels"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreSupplyChannelsChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias StoreSupplyChannelsChangedMessagePayload
	data, err := json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StoreSupplyChannelsChanged", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["addedSupplyChannels"] == nil {
		delete(raw, "addedSupplyChannels")
	}

	if raw["removedSupplyChannels"] == nil {
		delete(raw, "removedSupplyChannels")
	}

	return json.Marshal(raw)

}
