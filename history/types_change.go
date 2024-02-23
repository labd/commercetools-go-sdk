package history

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	Difference between the previous and next version of a resource represented by `previousValue` (omitted, for example, on creations) and `nextValue` of the associated change. A Change can also contain extra fields that provide further information.
*
*	They are not identical to the actual update actions sent.
*
 */
type Change interface{}

func mapDiscriminatorChange(input interface{}) (Change, error) {
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
	case "AddAddressChange":
		obj := AddAddressChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddAssetChange":
		obj := AddAssetChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddAssociateChange":
		obj := AddAssociateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddAttributeDefinitionChange":
		obj := AddAttributeDefinitionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddBillingAddressIdChange":
		obj := AddBillingAddressIdChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddChannelRolesChange":
		obj := AddChannelRolesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddCustomLineItemChange":
		obj := AddCustomLineItemChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddDeliveryChange":
		obj := AddDeliveryChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddDiscountCodeChange":
		obj := AddDiscountCodeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddEnumValueChange":
		obj := AddEnumValueChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddExternalImageChange":
		obj := AddExternalImageChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddFieldDefinitionChange":
		obj := AddFieldDefinitionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddInheritedAssociateChange":
		obj := AddInheritedAssociateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddInterfaceInteractionChange":
		obj := AddInterfaceInteractionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddItemShippingAddressesChange":
		obj := AddItemShippingAddressesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddLocalizedEnumValueChange":
		obj := AddLocalizedEnumValueChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddLocationChange":
		obj := AddLocationChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddOrderLineItemChange":
		obj := AddOrderLineItemChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddParcelToDeliveryChange":
		obj := AddParcelToDeliveryChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddPaymentChange":
		obj := AddPaymentChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddPlainEnumValueChange":
		obj := AddPlainEnumValueChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddPriceChange":
		obj := AddPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddProductChange":
		obj := AddProductChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddProductSelectionChange":
		obj := AddProductSelectionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddPropertyChange":
		obj := AddPropertyChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddReturnInfoChange":
		obj := AddReturnInfoChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddShippingAddressIdChange":
		obj := AddShippingAddressIdChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddShoppingListLineItemChange":
		obj := AddShoppingListLineItemChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddStateRolesChange":
		obj := AddStateRolesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddTaxRateChange":
		obj := AddTaxRateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddTextLineItemChange":
		obj := AddTextLineItemChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddToCategoryChange":
		obj := AddToCategoryChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddTransactionChange":
		obj := AddTransactionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AddVariantChange":
		obj := AddVariantChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeAddressChange":
		obj := ChangeAddressChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeAmountAuthorizedChange":
		obj := ChangeAmountAuthorizedChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeAmountPlannedChange":
		obj := ChangeAmountPlannedChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeAssetNameChange":
		obj := ChangeAssetNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeAssetOrderChange":
		obj := ChangeAssetOrderChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeAssociateChange":
		obj := ChangeAssociateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeAssociateModeChange":
		obj := ChangeAssociateModeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeAttributeConstraintChange":
		obj := ChangeAttributeConstraintChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeAttributeOrderByNameChange":
		obj := ChangeAttributeOrderByNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeBuyerAssignableChange":
		obj := ChangeBuyerAssignableChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeCartDiscountsChange":
		obj := ChangeCartDiscountsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeCartPredicateChange":
		obj := ChangeCartPredicateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeCustomLineItemQuantityChange":
		obj := ChangeCustomLineItemQuantityChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeDescriptionChange":
		obj := ChangeDescriptionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeEmailChange":
		obj := ChangeEmailChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeEnumValueLabelChange":
		obj := ChangeEnumValueLabelChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeEnumValueOrderChange":
		obj := ChangeEnumValueOrderChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeFieldDefinitionOrderChange":
		obj := ChangeFieldDefinitionOrderChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeGroupsChange":
		obj := ChangeGroupsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeInheritedAssociateChange":
		obj := ChangeInheritedAssociateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeInitialChange":
		obj := ChangeInitialChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeInputHintChange":
		obj := ChangeInputHintChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeIsActiveChange":
		obj := ChangeIsActiveChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeIsSearchableChange":
		obj := ChangeIsSearchableChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeKeyChange":
		obj := ChangeKeyChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeLabelChange":
		obj := ChangeLabelChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeLineItemQuantityChange":
		obj := ChangeLineItemQuantityChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeLocalizedDescriptionChange":
		obj := ChangeLocalizedDescriptionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeLocalizedEnumValueLabelChange":
		obj := ChangeLocalizedEnumValueLabelChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeLocalizedEnumValueOrderChange":
		obj := ChangeLocalizedEnumValueOrderChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeLocalizedNameChange":
		obj := ChangeLocalizedNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeMasterVariantChange":
		obj := ChangeMasterVariantChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeNameChange":
		obj := ChangeNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeOrderHintChange":
		obj := ChangeOrderHintChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeOrderStateChange":
		obj := ChangeOrderStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeParentChange":
		obj := ChangeParentChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeParentUnitChange":
		obj := ChangeParentUnitChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangePaymentStateChange":
		obj := ChangePaymentStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangePlainEnumValueLabelChange":
		obj := ChangePlainEnumValueLabelChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangePlainEnumValueOrderChange":
		obj := ChangePlainEnumValueOrderChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangePredicateChange":
		obj := ChangePredicateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangePriceChange":
		obj := ChangePriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeProductSelectionActiveChange":
		obj := ChangeProductSelectionActiveChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeQuantityChange":
		obj := ChangeQuantityChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeQuoteRequestStateChange":
		obj := ChangeQuoteRequestStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeQuoteStateChange":
		obj := ChangeQuoteStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeRequiresDiscountCodeChange":
		obj := ChangeRequiresDiscountCodeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeReviewRatingStatisticsChange":
		obj := ChangeReviewRatingStatisticsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeShipmentStateChange":
		obj := ChangeShipmentStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeShoppingListLineItemQuantityChange":
		obj := ChangeShoppingListLineItemQuantityChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeShoppingListLineItemsOrderChange":
		obj := ChangeShoppingListLineItemsOrderChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeSlugChange":
		obj := ChangeSlugChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeSortOrderChange":
		obj := ChangeSortOrderChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeStackingModeChange":
		obj := ChangeStackingModeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeStagedQuoteStateChange":
		obj := ChangeStagedQuoteStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeStateTypeChange":
		obj := ChangeStateTypeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeStatusChange":
		obj := ChangeStatusChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeTargetChange":
		obj := ChangeTargetChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.PreviousValue != nil {
			var err error
			obj.PreviousValue, err = mapDiscriminatorChangeTargetChangeValue(obj.PreviousValue)
			if err != nil {
				return nil, err
			}
		}
		if obj.NextValue != nil {
			var err error
			obj.NextValue, err = mapDiscriminatorChangeTargetChangeValue(obj.NextValue)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ChangeTaxCalculationModeChange":
		obj := ChangeTaxCalculationModeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeTaxModeChange":
		obj := ChangeTaxModeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeTaxRoundingModeChange":
		obj := ChangeTaxRoundingModeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeTextLineItemNameChange":
		obj := ChangeTextLineItemNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeTextLineItemQuantityChange":
		obj := ChangeTextLineItemQuantityChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeTextLineItemsOrderChange":
		obj := ChangeTextLineItemsOrderChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeTransactionInteractionIdChange":
		obj := ChangeTransactionInteractionIdChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeTransactionStateChange":
		obj := ChangeTransactionStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeTransactionTimestampChange":
		obj := ChangeTransactionTimestampChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ChangeValueChange":
		obj := ChangeValueChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.PreviousValue != nil {
			var err error
			obj.PreviousValue, err = mapDiscriminatorChangeValueChangeValue(obj.PreviousValue)
			if err != nil {
				return nil, err
			}
		}
		if obj.NextValue != nil {
			var err error
			obj.NextValue, err = mapDiscriminatorChangeValueChangeValue(obj.NextValue)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "MoveImageToPositionChange":
		obj := MoveImageToPositionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PublishChange":
		obj := PublishChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveAddressChange":
		obj := RemoveAddressChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveAssetChange":
		obj := RemoveAssetChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveAssociateChange":
		obj := RemoveAssociateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveAttributeDefinitionChange":
		obj := RemoveAttributeDefinitionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveBillingAddressIdChange":
		obj := RemoveBillingAddressIdChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveChannelRolesChange":
		obj := RemoveChannelRolesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveCustomLineItemChange":
		obj := RemoveCustomLineItemChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveDeliveryItemsChange":
		obj := RemoveDeliveryItemsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveDiscountCodeChange":
		obj := RemoveDiscountCodeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveEnumValuesChange":
		obj := RemoveEnumValuesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveFieldDefinitionChange":
		obj := RemoveFieldDefinitionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveFromCategoryChange":
		obj := RemoveFromCategoryChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveImageChange":
		obj := RemoveImageChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveInheritedAssociateChange":
		obj := RemoveInheritedAssociateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveItemShippingAddressesChange":
		obj := RemoveItemShippingAddressesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveLocalizedEnumValuesChange":
		obj := RemoveLocalizedEnumValuesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveLocationChange":
		obj := RemoveLocationChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveOrderLineItemChange":
		obj := RemoveOrderLineItemChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveParcelFromDeliveryChange":
		obj := RemoveParcelFromDeliveryChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemovePaymentChange":
		obj := RemovePaymentChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemovePriceChange":
		obj := RemovePriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveProductChange":
		obj := RemoveProductChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveProductSelectionChange":
		obj := RemoveProductSelectionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemovePropertyChange":
		obj := RemovePropertyChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveShippingAddressIdChange":
		obj := RemoveShippingAddressIdChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveShoppingListLineItemChange":
		obj := RemoveShoppingListLineItemChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveStateRolesChange":
		obj := RemoveStateRolesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveTaxRateChange":
		obj := RemoveTaxRateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveTextLineItemChange":
		obj := RemoveTextLineItemChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RemoveVariantChange":
		obj := RemoveVariantChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "RequestQuoteRenegotiationChange":
		obj := RequestQuoteRenegotiationChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAddressChange":
		obj := SetAddressChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAddressCustomFieldChange":
		obj := SetAddressCustomFieldChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAddressCustomTypeChange":
		obj := SetAddressCustomTypeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAnonymousIdChange":
		obj := SetAnonymousIdChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetApplicationVersionChange":
		obj := SetApplicationVersionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAssetCustomFieldChange":
		obj := SetAssetCustomFieldChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAssetCustomTypeChange":
		obj := SetAssetCustomTypeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAssetDescriptionChange":
		obj := SetAssetDescriptionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAssetKeyChange":
		obj := SetAssetKeyChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAssetSourcesChange":
		obj := SetAssetSourcesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAssetTagsChange":
		obj := SetAssetTagsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAttributeChange":
		obj := SetAttributeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAuthenticationModeChange":
		obj := SetAuthenticationModeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetAuthorNameChange":
		obj := SetAuthorNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetBillingAddressChange":
		obj := SetBillingAddressChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCartPredicateChange":
		obj := SetCartPredicateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCategoryOrderHintChange":
		obj := SetCategoryOrderHintChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetChannelRolesChange":
		obj := SetChannelRolesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCompanyNameChange":
		obj := SetCompanyNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetContactEmailChange":
		obj := SetContactEmailChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCountriesChange":
		obj := SetCountriesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCountryChange":
		obj := SetCountryChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomFieldChange":
		obj := SetCustomFieldChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomLineItemCustomFieldChange":
		obj := SetCustomLineItemCustomFieldChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomLineItemCustomTypeChange":
		obj := SetCustomLineItemCustomTypeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomLineItemMoneyChange":
		obj := SetCustomLineItemMoneyChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomLineItemShippingDetailsChange":
		obj := SetCustomLineItemShippingDetailsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomLineItemTaxAmountChange":
		obj := SetCustomLineItemTaxAmountChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomLineItemTaxCategoryChange":
		obj := SetCustomLineItemTaxCategoryChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomLineItemTaxRateChange":
		obj := SetCustomLineItemTaxRateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomLineItemTaxedPriceChange":
		obj := SetCustomLineItemTaxedPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomLineItemTotalPriceChange":
		obj := SetCustomLineItemTotalPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomShippingMethodChange":
		obj := SetCustomShippingMethodChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomTypeChange":
		obj := SetCustomTypeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomerChange":
		obj := SetCustomerChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomerEmailChange":
		obj := SetCustomerEmailChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomerGroupChange":
		obj := SetCustomerGroupChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomerIdChange":
		obj := SetCustomerIdChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetCustomerNumberChange":
		obj := SetCustomerNumberChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetDateOfBirthChange":
		obj := SetDateOfBirthChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetDefaultBillingAddressChange":
		obj := SetDefaultBillingAddressChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetDefaultShippingAddressChange":
		obj := SetDefaultShippingAddressChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetDeleteDaysAfterLastModificationChange":
		obj := SetDeleteDaysAfterLastModificationChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetDeliveryAddressChange":
		obj := SetDeliveryAddressChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetDeliveryItemsChange":
		obj := SetDeliveryItemsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetDescriptionChange":
		obj := SetDescriptionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetDiscountedPriceChange":
		obj := SetDiscountedPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetDistributionChannelsChange":
		obj := SetDistributionChannelsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetExpectedDeliveryChange":
		obj := SetExpectedDeliveryChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetExternalIdChange":
		obj := SetExternalIdChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetFirstNameChange":
		obj := SetFirstNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetGeoLocationChange":
		obj := SetGeoLocationChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetImageLabelChange":
		obj := SetImageLabelChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetInputTipChange":
		obj := SetInputTipChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetInterfaceIdChange":
		obj := SetInterfaceIdChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetIsValidChange":
		obj := SetIsValidChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetKeyChange":
		obj := SetKeyChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLanguagesChange":
		obj := SetLanguagesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLastNameChange":
		obj := SetLastNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemDeactivatedAtChange":
		obj := SetLineItemDeactivatedAtChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemDiscountedPriceChange":
		obj := SetLineItemDiscountedPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemDiscountedPricePerQuantityChange":
		obj := SetLineItemDiscountedPricePerQuantityChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemDistributionChannelChange":
		obj := SetLineItemDistributionChannelChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemPriceChange":
		obj := SetLineItemPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemProductKeyChange":
		obj := SetLineItemProductKeyChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemProductSlugChange":
		obj := SetLineItemProductSlugChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemShippingDetailsChange":
		obj := SetLineItemShippingDetailsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemTaxAmountChange":
		obj := SetLineItemTaxAmountChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemTaxRateChange":
		obj := SetLineItemTaxRateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemTaxedPriceChange":
		obj := SetLineItemTaxedPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLineItemTotalPriceChange":
		obj := SetLineItemTotalPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLocaleChange":
		obj := SetLocaleChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLocalizedDescriptionChange":
		obj := SetLocalizedDescriptionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetMaxApplicationsChange":
		obj := SetMaxApplicationsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetMaxApplicationsPerCustomerChange":
		obj := SetMaxApplicationsPerCustomerChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetMetaDescriptionChange":
		obj := SetMetaDescriptionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetMetaKeywordsChange":
		obj := SetMetaKeywordsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetMetaTitleChange":
		obj := SetMetaTitleChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetMethodInfoInterfaceChange":
		obj := SetMethodInfoInterfaceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetMethodInfoMethodChange":
		obj := SetMethodInfoMethodChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetMethodInfoNameChange":
		obj := SetMethodInfoNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetMiddleNameChange":
		obj := SetMiddleNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetNameChange":
		obj := SetNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetLocalizedNameChange":
		obj := SetLocalizedNameChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetOrderLineItemCustomFieldChange":
		obj := SetOrderLineItemCustomFieldChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetOrderLineItemCustomTypeChange":
		obj := SetOrderLineItemCustomTypeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetOrderNumberChange":
		obj := SetOrderNumberChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetOrderTaxedPriceChange":
		obj := SetOrderTaxedPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetOrderTotalPriceChange":
		obj := SetOrderTotalPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetOrderTotalTaxChange":
		obj := SetOrderTotalTaxChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetParcelItemsChange":
		obj := SetParcelItemsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetParcelMeasurementsChange":
		obj := SetParcelMeasurementsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetParcelTrackingDataChange":
		obj := SetParcelTrackingDataChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetPermissionsChange":
		obj := SetPermissionsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetPricesChange":
		obj := SetPricesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetProductCountChange":
		obj := SetProductCountChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetProductPriceCustomFieldChange":
		obj := SetProductPriceCustomFieldChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetProductPriceCustomTypeChange":
		obj := SetProductPriceCustomTypeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetProductSelectionsChange":
		obj := SetProductSelectionsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetProductVariantKeyChange":
		obj := SetProductVariantKeyChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetPropertyChange":
		obj := SetPropertyChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetPurchaseOrderNumberChange":
		obj := SetPurchaseOrderNumberChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetRatingChange":
		obj := SetRatingChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetReservationsChange":
		obj := SetReservationsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetRestockableInDaysChange":
		obj := SetRestockableInDaysChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetReturnPaymentStateChange":
		obj := SetReturnPaymentStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetReturnShipmentStateChange":
		obj := SetReturnShipmentStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetSalutationChange":
		obj := SetSalutationChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetSearchKeywordsChange":
		obj := SetSearchKeywordsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetSellerCommentChange":
		obj := SetSellerCommentChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetShippingAddressChange":
		obj := SetShippingAddressChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetShippingInfoPriceChange":
		obj := SetShippingInfoPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetShippingInfoTaxedPriceChange":
		obj := SetShippingInfoTaxedPriceChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetShippingMethodChange":
		obj := SetShippingMethodChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetShippingMethodTaxAmountChange":
		obj := SetShippingMethodTaxAmountChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetShippingMethodTaxRateChange":
		obj := SetShippingMethodTaxRateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetShippingRateChange":
		obj := SetShippingRateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetShippingRateInputChange":
		obj := SetShippingRateInputChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetShoppingListLineItemCustomFieldChange":
		obj := SetShoppingListLineItemCustomFieldChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetShoppingListLineItemCustomTypeChange":
		obj := SetShoppingListLineItemCustomTypeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetSkuChange":
		obj := SetSkuChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetSlugChange":
		obj := SetSlugChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetStateRolesChange":
		obj := SetStateRolesChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetStatusInterfaceCodeChange":
		obj := SetStatusInterfaceCodeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetStatusInterfaceTextChange":
		obj := SetStatusInterfaceTextChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetStoreChange":
		obj := SetStoreChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetStoreModeChange":
		obj := SetStoreModeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetStoresChange":
		obj := SetStoresChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetSupplyChannelChange":
		obj := SetSupplyChannelChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetSupplyChannelsChange":
		obj := SetSupplyChannelsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetTargetChange":
		obj := SetTargetChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetTaxCategoryChange":
		obj := SetTaxCategoryChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetTextChange":
		obj := SetTextChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetTextLineItemCustomFieldChange":
		obj := SetTextLineItemCustomFieldChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetTextLineItemCustomTypeChange":
		obj := SetTextLineItemCustomTypeChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetTextLineItemDescriptionChange":
		obj := SetTextLineItemDescriptionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetTitleChange":
		obj := SetTitleChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetTransitionsChange":
		obj := SetTransitionsChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetValidFromAndUntilChange":
		obj := SetValidFromAndUntilChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetValidFromChange":
		obj := SetValidFromChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetValidToChange":
		obj := SetValidToChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetValidUntilChange":
		obj := SetValidUntilChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetValueChange":
		obj := SetValueChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetVariantAvailabilityChange":
		obj := SetVariantAvailabilityChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetVariantSelectionChange":
		obj := SetVariantSelectionChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SetVatIdChange":
		obj := SetVatIdChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "TransitionCustomLineItemStateChange":
		obj := TransitionCustomLineItemStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "TransitionLineItemStateChange":
		obj := TransitionLineItemStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "TransitionStateChange":
		obj := TransitionStateChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "UnknownChange":
		obj := UnknownChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "UnpublishChange":
		obj := UnpublishChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "UpdateSyncInfoChange":
		obj := UpdateSyncInfoChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "VerifyEmailChange":
		obj := VerifyEmailChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Change triggered by the [Add Address](ctp:api:type:CustomerAddAddressAction) update action.
 */
type AddAddressChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue Address `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddAddressChange) MarshalJSON() ([]byte, error) {
	type Alias AddAddressChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddAddressChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Add Asset](ctp:api:type:CategoryAddAssetAction) on Categories.
*	- [Add Asset](ctp:api:type:ProductAddAssetAction) on Products.
*
 */
type AddAssetChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Asset `json:"previousValue"`
	// Value after the change.
	NextValue Asset `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddAssetChange) MarshalJSON() ([]byte, error) {
	type Alias AddAssetChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddAssetChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Associate](ctp:api:type:BusinessUnitAddAssociateAction) update action.
 */
type AddAssociateChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue Associate `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddAssociateChange) MarshalJSON() ([]byte, error) {
	type Alias AddAssociateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddAssociateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Attribute Definition](ctp:api:type:ProductTypeAddAttributeDefinitionAction) update action.
 */
type AddAttributeDefinitionChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue AttributeDefinition `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddAttributeDefinitionChange) MarshalJSON() ([]byte, error) {
	type Alias AddAttributeDefinitionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddAttributeDefinitionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Billing Address ID](ctp:api:type:CustomerAddBillingAddressIdAction) update action.
 */
type AddBillingAddressIdChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []string `json:"previousValue"`
	// Value after the change.
	NextValue []string `json:"nextValue"`
	// Address added to `billingAddressIds`.
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddBillingAddressIdChange) MarshalJSON() ([]byte, error) {
	type Alias AddBillingAddressIdChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddBillingAddressIdChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Roles](ctp:api:type:ChannelAddRolesAction) update action.
 */
type AddChannelRolesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []ChannelRoleEnum `json:"previousValue"`
	// Value after the change.
	NextValue []ChannelRoleEnum `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddChannelRolesChange) MarshalJSON() ([]byte, error) {
	type Alias AddChannelRolesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddChannelRolesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add CustomLineItem](ctp:api:type:StagedOrderAddCustomLineItemAction) update action.
 */
type AddCustomLineItemChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CustomLineItem `json:"previousValue"`
	// Value after the change.
	NextValue CustomLineItem `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddCustomLineItemChange) MarshalJSON() ([]byte, error) {
	type Alias AddCustomLineItemChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddCustomLineItemChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Delivery](ctp:api:type:OrderAddDeliveryAction) update action.
 */
type AddDeliveryChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue DeliveryChangeValue `json:"previousValue"`
	// Value after the change.
	NextValue DeliveryChangeValue `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddDeliveryChange) MarshalJSON() ([]byte, error) {
	type Alias AddDeliveryChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddDeliveryChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add DiscountCode](ctp:api:type:StagedOrderAddDiscountCodeAction) update action.
 */
type AddDiscountCodeChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue DiscountCodeInfo `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddDiscountCodeChange) MarshalJSON() ([]byte, error) {
	type Alias AddDiscountCodeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddDiscountCodeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add EnumValue to FieldDefinition](ctp:api:type:TypeAddEnumValueAction) update action.
 */
type AddEnumValueChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue EnumValue `json:"nextValue"`
	// Name of the updated [FieldDefinition](ctp:api:type:FieldDefinition).
	FieldName string `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddEnumValueChange) MarshalJSON() ([]byte, error) {
	type Alias AddEnumValueChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddEnumValueChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add External Image](ctp:api:type:ProductAddExternalImageAction) update action.
 */
type AddExternalImageChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Image `json:"previousValue"`
	// Value after the change.
	NextValue []Image `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddExternalImageChange) MarshalJSON() ([]byte, error) {
	type Alias AddExternalImageChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddExternalImageChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add FieldDefinition](ctp:api:type:TypeAddFieldDefinitionAction) update action.
 */
type AddFieldDefinitionChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue FieldDefinition `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddFieldDefinitionChange) MarshalJSON() ([]byte, error) {
	type Alias AddFieldDefinitionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddFieldDefinitionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Associate](ctp:api:type:BusinessUnitAddAssociateAction) update action on a parent of a Business Unit in cases where [inheritance applies](/../api/associates-overview#conditions-for-inheritance).
 */
type AddInheritedAssociateChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue InheritedAssociate `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddInheritedAssociateChange) MarshalJSON() ([]byte, error) {
	type Alias AddInheritedAssociateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddInheritedAssociateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add InterfaceInteraction](ctp:api:type:PaymentAddInterfaceInteractionAction) update action.
 */
type AddInterfaceInteractionChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue CustomFieldExpandedValue `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddInterfaceInteractionChange) MarshalJSON() ([]byte, error) {
	type Alias AddInterfaceInteractionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddInterfaceInteractionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Add ItemShippingAddress](ctp:api:type:OrderAddItemShippingAddressAction) on Orders.
*	- [Add ItemShippingAddress](ctp:api:type:StagedOrderAddItemShippingAddressAction) on Staged Orders.
*
 */
type AddItemShippingAddressesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Address `json:"previousValue"`
	// Value after the change.
	NextValue Address `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddItemShippingAddressesChange) MarshalJSON() ([]byte, error) {
	type Alias AddItemShippingAddressesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddItemShippingAddressesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Add LocalizableEnumValue to AttributeDefinition](ctp:api:type:ProductTypeAddLocalizedEnumValueAction) on Product Types.
*	- [Add LocalizedEnumValue to FieldDefinition](ctp:api:type:TypeAddLocalizedEnumValueAction) on Types.
*
 */
type AddLocalizedEnumValueChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue AttributeLocalizedEnumValue `json:"nextValue"`
	// Name of the updated [FieldDefinition](ctp:api:type:FieldDefinition); only present on changes to Types.
	FieldName string `json:"fieldName"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition); only present on changes to Product Types.
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddLocalizedEnumValueChange) MarshalJSON() ([]byte, error) {
	type Alias AddLocalizedEnumValueChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddLocalizedEnumValueChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Location](ctp:api:type:ZoneAddLocationAction) update action.
 */
type AddLocationChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue Location `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddLocationChange) MarshalJSON() ([]byte, error) {
	type Alias AddLocationChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddLocationChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add LineItem](ctp:api:type:StagedOrderAddLineItemAction) update action.
 */
type AddOrderLineItemChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LineItem `json:"previousValue"`
	// Value after the change.
	NextValue LineItem `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddOrderLineItemChange) MarshalJSON() ([]byte, error) {
	type Alias AddOrderLineItemChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddOrderLineItemChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Add Parcel](ctp:api:type:OrderAddParcelToDeliveryAction) on Orders.
*	- [Add Parcel](ctp:api:type:StagedOrderAddParcelToDeliveryAction) on Staged Orders.
*
 */
type AddParcelToDeliveryChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue Parcel `json:"nextValue"`
	// `id` of the [Delivery](ctp:api:type:Delivery) to which the Parcel was added.
	DeliveryId string `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddParcelToDeliveryChange) MarshalJSON() ([]byte, error) {
	type Alias AddParcelToDeliveryChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddParcelToDeliveryChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Add Payment](ctp:api:type:OrderAddPaymentAction) on Orders.
*	- [Add Payment](ctp:api:type:StagedOrderAddPaymentAction) on Staged Orders.
*
 */
type AddPaymentChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue PaymentInfo `json:"previousValue"`
	// Value after the change.
	NextValue PaymentInfo `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddPaymentChange) MarshalJSON() ([]byte, error) {
	type Alias AddPaymentChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddPaymentChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add PlainEnumValue to AttributeDefinition](ctp:api:type:ProductTypeAddPlainEnumValueAction) update action.
 */
type AddPlainEnumValueChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue AttributePlainEnumValue `json:"nextValue"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition).
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddPlainEnumValueChange) MarshalJSON() ([]byte, error) {
	type Alias AddPlainEnumValueChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddPlainEnumValueChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Price](ctp:api:type:ProductAddPriceAction) update action.
 */
type AddPriceChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue Price `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
	// `id` of the Embedded [Price](ctp:api:type:Price).
	PriceId string `json:"priceId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddPriceChange) MarshalJSON() ([]byte, error) {
	type Alias AddPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddPriceChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Product](ctp:api:type:ProductSelectionAddProductAction) update action.
 */
type AddProductChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue Reference `json:"nextValue"`
	// The [Product Variants](ctp:api:type:ProductVariant) included in the [Product Selection](ctp:api:type:ProductSelection).
	VariantSelection ProductVariantSelection `json:"variantSelection"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddProductChange) MarshalJSON() ([]byte, error) {
	type Alias AddProductChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddProductChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Product Selection](ctp:api:type:StoreAddProductSelectionAction) update action.
 */
type AddProductSelectionChange struct {
	Change        string                  `json:"change"`
	PreviousValue ProductSelectionSetting `json:"previousValue"`
	NextValue     ProductSelectionSetting `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddProductSelectionChange) MarshalJSON() ([]byte, error) {
	type Alias AddProductSelectionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddProductSelectionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Update CustomObject](ctp:api:endpoint:/{projectKey}/custom-objects:POST) request when a new property is added.
 */
type AddPropertyChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
	// Path to the new property that was added.
	Path string `json:"path"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddPropertyChange) MarshalJSON() ([]byte, error) {
	type Alias AddPropertyChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddPropertyChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Add ReturnInfo](ctp:api:type:OrderAddReturnInfoAction) on Orders.
*	- [Add ReturnInfo](ctp:api:type:StagedOrderAddReturnInfoAction) on Staged Orders.
*
 */
type AddReturnInfoChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue ReturnInfo `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddReturnInfoChange) MarshalJSON() ([]byte, error) {
	type Alias AddReturnInfoChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddReturnInfoChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Shipping Address ID](ctp:api:type:CustomerAddShippingAddressIdAction) update action.
 */
type AddShippingAddressIdChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []string `json:"previousValue"`
	// Value after the change.
	NextValue []string `json:"nextValue"`
	// Address added to `shippingAddressIds`.
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddShippingAddressIdChange) MarshalJSON() ([]byte, error) {
	type Alias AddShippingAddressIdChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddShippingAddressIdChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add ShoppingListLineItem](ctp:api:type:ShoppingListAddLineItemAction) update action.
 */
type AddShoppingListLineItemChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LineItem `json:"previousValue"`
	// Value after the change.
	NextValue LineItem `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddShoppingListLineItemChange) MarshalJSON() ([]byte, error) {
	type Alias AddShoppingListLineItemChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddShoppingListLineItemChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add State roles](ctp:api:type:StateAddRolesAction) update action.
 */
type AddStateRolesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []StateRoleEnum `json:"previousValue"`
	// Value after the change.
	NextValue []StateRoleEnum `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddStateRolesChange) MarshalJSON() ([]byte, error) {
	type Alias AddStateRolesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddStateRolesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add TaxRate](ctp:api:type:TaxCategoryAddTaxRateAction) update action.
 */
type AddTaxRateChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue TaxRate `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddTaxRateChange) MarshalJSON() ([]byte, error) {
	type Alias AddTaxRateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddTaxRateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add TextLineItem](ctp:api:type:ShoppingListAddTextLineItemAction) update action.
 */
type AddTextLineItemChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue TextLineItem `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddTextLineItemChange) MarshalJSON() ([]byte, error) {
	type Alias AddTextLineItemChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddTextLineItemChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add to Category](ctp:api:type:ProductAddToCategoryAction) update action.
 */
type AddToCategoryChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Reference `json:"previousValue"`
	// Value after the change.
	NextValue []Reference `json:"nextValue"`
	// Category to which the Product was added.
	Category Reference `json:"category"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddToCategoryChange) MarshalJSON() ([]byte, error) {
	type Alias AddToCategoryChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddToCategoryChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add Transaction](ctp:api:type:PaymentAddTransactionAction) update action.
 */
type AddTransactionChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue Transaction `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddTransactionChange) MarshalJSON() ([]byte, error) {
	type Alias AddTransactionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddTransactionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Add ProductVariant](ctp:api:type:ProductAddVariantAction) update action.
 */
type AddVariantChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Variant `json:"previousValue"`
	// Value after the change.
	NextValue Variant `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AddVariantChange) MarshalJSON() ([]byte, error) {
	type Alias AddVariantChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AddVariantChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Address](ctp:api:type:CustomerChangeAddressAction) update action.
 */
type ChangeAddressChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Address `json:"previousValue"`
	// Value after the change.
	NextValue Address `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeAddressChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeAddressChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeAddressChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered automatically due to a user-initiated change.
 */
type ChangeAmountAuthorizedChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Money `json:"previousValue"`
	// Value after the change.
	NextValue Money `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeAmountAuthorizedChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeAmountAuthorizedChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeAmountAuthorizedChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change AmountPlanned](ctp:api:type:PaymentChangeAmountPlannedAction) update action.
 */
type ChangeAmountPlannedChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Money `json:"previousValue"`
	// Value after the change.
	NextValue Money `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeAmountPlannedChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeAmountPlannedChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeAmountPlannedChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change Asset Name](ctp:api:type:CategoryChangeAssetNameAction) on Categories.
*	- [Change Asset Name](ctp:api:type:ProductChangeAssetNameAction) on Products.
*
 */
type ChangeAssetNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
	// Information about the updated Asset.
	Asset AssetChangeValue `json:"asset"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeAssetNameChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeAssetNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeAssetNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change Asset Order](ctp:api:type:CategoryChangeAssetOrderAction) on Categories.
*	- [Change Asset Order](ctp:api:type:ProductChangeAssetOrderAction) on Products.
*
 */
type ChangeAssetOrderChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue []LocalizedString `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeAssetOrderChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeAssetOrderChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeAssetOrderChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Associate](ctp:api:type:BusinessUnitChangeAssociateAction) update action.
 */
type ChangeAssociateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Associate `json:"previousValue"`
	// Value after the change.
	NextValue Associate `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeAssociateChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeAssociateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeAssociateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Associate Mode](ctp:api:type:BusinessUnitChangeAssociateModeAction) update action.
 */
type ChangeAssociateModeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue BusinessUnitAssociateMode `json:"previousValue"`
	// Value after the change.
	NextValue BusinessUnitAssociateMode `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeAssociateModeChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeAssociateModeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeAssociateModeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change AttributeDefinition AttributeConstraint](ctp:api:type:ProductTypeChangeAttributeConstraintAction) update action.
 */
type ChangeAttributeConstraintChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue AttributeConstraintEnum `json:"previousValue"`
	// Value after the change.
	NextValue AttributeConstraintEnum `json:"nextValue"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition).
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeAttributeConstraintChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeAttributeConstraintChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeAttributeConstraintChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change the order of AttributeDefinitions](ctp:api:type:ProductTypeChangeAttributeOrderByNameAction) update action.
 */
type ChangeAttributeOrderByNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []string `json:"previousValue"`
	// Value after the change.
	NextValue []string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeAttributeOrderByNameChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeAttributeOrderByNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeAttributeOrderByNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change BuyerAssignable](ctp:api:type:AssociateRoleChangeBuyerAssignableAction) update action.
*
 */
type ChangeBuyerAssignableChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue bool `json:"previousValue"`
	// Value after the change.
	NextValue bool `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeBuyerAssignableChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeBuyerAssignableChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeBuyerAssignableChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change CartDiscounts](ctp:api:type:DiscountCodeChangeCartDiscountsAction) update action.
 */
type ChangeCartDiscountsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Reference `json:"previousValue"`
	// Value after the change.
	NextValue []Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeCartDiscountsChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeCartDiscountsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeCartDiscountsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Cart Predicate](ctp:api:type:CartDiscountChangeCartPredicateAction) update action.
 */
type ChangeCartPredicateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeCartPredicateChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeCartPredicateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeCartPredicateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change CustomLineItem Quantity](ctp:api:type:StagedOrderChangeCustomLineItemQuantityAction) update action.
 */
type ChangeCustomLineItemQuantityChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue int `json:"previousValue"`
	// Value after the change.
	NextValue int `json:"nextValue"`
	// Name of the CustomLineItem.
	CustomLineItem LocalizedString `json:"customLineItem"`
	// `id` of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeCustomLineItemQuantityChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeCustomLineItemQuantityChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeCustomLineItemQuantityChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Description](ctp:api:type:ProductTypeChangeDescriptionAction) update action.
*
 */
type ChangeDescriptionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeDescriptionChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeDescriptionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeDescriptionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Email](ctp:api:type:CustomerChangeEmailAction) update action.
 */
type ChangeEmailChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeEmailChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeEmailChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeEmailChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change EnumValue Label](ctp:api:type:TypeChangeEnumValueLabelAction) update action.
 */
type ChangeEnumValueLabelChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
	// Name of the updated [FieldDefinition](ctp:api:type:FieldDefinition).
	FieldName string `json:"fieldName"`
	// Key of the updated values.
	ValueKey string `json:"valueKey"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeEnumValueLabelChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeEnumValueLabelChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeEnumValueLabelChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change the order of EnumValues](ctp:api:type:TypeChangeEnumValueOrderAction) update action.
 */
type ChangeEnumValueOrderChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []EnumValue `json:"previousValue"`
	// Value after the change.
	NextValue []EnumValue `json:"nextValue"`
	// Name of the updated [FieldDefinition](ctp:api:type:FieldDefinition).
	FieldName string `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeEnumValueOrderChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeEnumValueOrderChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeEnumValueOrderChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change the order of FieldDefinitions](ctp:api:type:TypeChangeFieldDefinitionOrderAction) update action.
 */
type ChangeFieldDefinitionOrderChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []FieldDefinitionOrderValue `json:"previousValue"`
	// Value after the change.
	NextValue []FieldDefinitionOrderValue `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeFieldDefinitionOrderChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeFieldDefinitionOrderChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeFieldDefinitionOrderChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Groups](ctp:api:type:DiscountCodeChangeGroupsAction) update action.
 */
type ChangeGroupsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []string `json:"previousValue"`
	// Value after the change.
	NextValue []string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeGroupsChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeGroupsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeGroupsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Associate](ctp:api:type:BusinessUnitChangeAssociateAction) update action on a parent of a Business Unit in cases where [inheritance applies](/../api/associates-overview#conditions-for-inheritance).
 */
type ChangeInheritedAssociateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue InheritedAssociate `json:"previousValue"`
	// Value after the change.
	NextValue InheritedAssociate `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeInheritedAssociateChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeInheritedAssociateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeInheritedAssociateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change initial State](ctp:api:type:StateChangeInitialAction) update action.
 */
type ChangeInitialChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue bool `json:"previousValue"`
	// Value after the change.
	NextValue bool `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeInitialChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeInitialChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeInitialChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change AttributeDefinition InputHint](ctp:api:type:ProductTypeChangeInputHintAction) on Product Types.
*	- [Change InputHint](ctp:api:type:TypeChangeInputHintAction) on Types.
*
 */
type ChangeInputHintChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TextInputHint `json:"previousValue"`
	// Value after the change.
	NextValue TextInputHint `json:"nextValue"`
	// Name of the updated [FieldDefinition](ctp:api:type:FieldDefinition); only present on changes to Types.
	FieldName string `json:"fieldName"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition); only present on changes to Product Types.
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeInputHintChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeInputHintChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeInputHintChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change IsActive](ctp:api:type:CartDiscountChangeIsActiveAction) on Cart Discounts.
*	- [Change IsActive](ctp:api:type:DiscountCodeChangeIsActiveAction) on Discount Codes.
*	- [Change IsActive](ctp:api:type:ProductDiscountChangeIsActiveAction) on Product Discounts.
*
 */
type ChangeIsActiveChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue bool `json:"previousValue"`
	// Value after the change.
	NextValue bool `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeIsActiveChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeIsActiveChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeIsActiveChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change AttributeDefinition IsSearchable](ctp:api:type:ProductTypeChangeIsSearchableAction) update action.
 */
type ChangeIsSearchableChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue bool `json:"previousValue"`
	// Value after the change.
	NextValue bool `json:"nextValue"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition).
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeIsSearchableChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeIsSearchableChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeIsSearchableChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change Key](ctp:api:type:ChannelChangeKeyAction) on Channels.
*	- [Change State key](ctp:api:type:StateChangeKeyAction) on States.
*	- [Change Key](ctp:api:type:TypeChangeKeyAction) on Types.
*
 */
type ChangeKeyChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeKeyChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeKeyChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeKeyChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change AttributeDefinition Label](ctp:api:type:ProductTypeChangeLabelAction) on Product Types.
*	- [Change FieldDefinition Label](ctp:api:type:TypeChangeLabelAction) on Types.
*
 */
type ChangeLabelChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
	// Name of the updated [FieldDefinition](ctp:api:type:FieldDefinition); only present on changes to Types).
	FieldName string `json:"fieldName"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition); only present on changes to Product Types.
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeLabelChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeLabelChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeLabelChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change LineItem Quantity](ctp:api:type:StagedOrderChangeLineItemQuantityAction) update action.
 */
type ChangeLineItemQuantityChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue int `json:"previousValue"`
	// Value after the change
	NextValue int `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the updated Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
	// `id` of the updated [LineItem](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeLineItemQuantityChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeLineItemQuantityChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeLineItemQuantityChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Description](ctp:api:type:ChannelChangeDescriptionAction) update action.
*
 */
type ChangeLocalizedDescriptionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeLocalizedDescriptionChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeLocalizedDescriptionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeLocalizedDescriptionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change the label of a LocalizedEnumValue](ctp:api:type:ProductTypeChangeLocalizedEnumValueLabelAction) on Product Types.
*	- [Change LocalizedEnumValue Label](ctp:api:type:TypeChangeLocalizedEnumValueLabelAction) on Types.
*
 */
type ChangeLocalizedEnumValueLabelChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
	// Name of the updated [FieldDefinition](ctp:api:type:FieldDefinition); only present on changes to Types.
	FieldName string `json:"fieldName"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition); only present on changes to Product Types.
	AttributeName string `json:"attributeName"`
	// Key of the updated values.
	ValueKey string `json:"valueKey"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeLocalizedEnumValueLabelChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeLocalizedEnumValueLabelChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeLocalizedEnumValueLabelChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change the order of LocalizedEnumValues](ctp:api:type:ProductTypeChangeLocalizedEnumValueOrderAction) on Product Types.
*	- [Change the order of LocalizedEnumValues](ctp:api:type:TypeChangeLocalizedEnumValueOrderAction) on Types.
*
 */
type ChangeLocalizedEnumValueOrderChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []LocalizedEnumValue `json:"previousValue"`
	// Value after the change.
	NextValue []LocalizedEnumValue `json:"nextValue"`
	// Name of the updated [FieldDefinition](ctp:api:type:FieldDefinition); only present on changes to Types.
	FieldName string `json:"fieldName"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition); only present on changes to Product Types.
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeLocalizedEnumValueOrderChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeLocalizedEnumValueOrderChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeLocalizedEnumValueOrderChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change Name](ctp:api:type:CartDiscountChangeNameAction) on Cart Discounts.
*	- [Change Name](ctp:api:type:CategoryChangeNameAction) on Categories.
*	- [Change Name](ctp:api:type:ChannelChangeNameAction) on Channels.
*	- [Change Name](ctp:api:type:ProductChangeNameAction) on Products.
*	- [Change Name](ctp:api:type:ProductDiscountChangeNameAction) on Product Discounts.
*	- [Change Name](ctp:api:type:ProductSelectionChangeNameAction) on Product Selections.
*	- [Change Name](ctp:api:type:ShoppingListChangeNameAction) on Shopping Lists.
*	- [Change Name](ctp:api:type:ZoneChangeNameAction) on Zones.
*
 */
type ChangeLocalizedNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeLocalizedNameChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeLocalizedNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeLocalizedNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Master Variant](ctp:api:type:ProductChangeMasterVariantAction) update action.
 */
type ChangeMasterVariantChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Variant `json:"previousValue"`
	// Value after the change.
	NextValue Variant `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeMasterVariantChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeMasterVariantChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeMasterVariantChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change Name](ctp:api:type:CustomerGroupChangeNameAction) on Customer Groups.
*	- [Change Name](ctp:api:type:ProductTypeChangeNameAction) on Product Types.
*	- [Change Name](ctp:api:type:TaxCategoryChangeNameAction) on Tax Categories.
*	- [Change Name](ctp:api:type:ZoneChangeNameAction) on Zones.
*
 */
type ChangeNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeNameChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change OrderHint](ctp:api:type:CategoryChangeOrderHintAction) update action.
 */
type ChangeOrderHintChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeOrderHintChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeOrderHintChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeOrderHintChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change OrderState](ctp:api:type:OrderChangeOrderStateAction) on Orders.
*	- [Change OrderState](ctp:api:type:StagedOrderChangeOrderStateAction) on Staged Orders.
*
 */
type ChangeOrderStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue OrderState `json:"previousValue"`
	// Value after the change.
	NextValue OrderState `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeOrderStateChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeOrderStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeOrderStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Parent](ctp:api:type:CategoryChangeParentAction) update action.
 */
type ChangeParentChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Reference `json:"previousValue"`
	// Value after the change.
	NextValue Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeParentChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeParentChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeParentChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Parent Unit](ctp:api:type:BusinessUnitChangeParentUnitAction) update action.
 */
type ChangeParentUnitChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue KeyReference `json:"previousValue"`
	// Value after the change.
	NextValue KeyReference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeParentUnitChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeParentUnitChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeParentUnitChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change PaymentState](ctp:api:type:OrderChangePaymentStateAction) on Orders.
*	- [Change PaymentState](ctp:api:type:StagedOrderChangePaymentStateAction) on Staged Orders.
*
 */
type ChangePaymentStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue PaymentState `json:"previousValue"`
	// Value after the change.
	NextValue PaymentState `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangePaymentStateChange) MarshalJSON() ([]byte, error) {
	type Alias ChangePaymentStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangePaymentStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change the label of an EnumValue](ctp:api:type:ProductTypeChangePlainEnumValueLabelAction) update action.
 */
type ChangePlainEnumValueLabelChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition).
	AttributeName string `json:"attributeName"`
	// Key of the updated values.
	ValueKey string `json:"valueKey"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangePlainEnumValueLabelChange) MarshalJSON() ([]byte, error) {
	type Alias ChangePlainEnumValueLabelChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangePlainEnumValueLabelChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change the order of EnumValues](ctp:api:type:ProductTypeChangePlainEnumValueOrderAction) update action.
 */
type ChangePlainEnumValueOrderChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []EnumValue `json:"previousValue"`
	// Value after the change.
	NextValue []EnumValue `json:"nextValue"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition).
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangePlainEnumValueOrderChange) MarshalJSON() ([]byte, error) {
	type Alias ChangePlainEnumValueOrderChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangePlainEnumValueOrderChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Predicate](ctp:api:type:ProductDiscountChangePredicateAction) update action.
 */
type ChangePredicateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangePredicateChange) MarshalJSON() ([]byte, error) {
	type Alias ChangePredicateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangePredicateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Price](ctp:api:type:ProductChangePriceAction) update action.
 */
type ChangePriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Price `json:"previousValue"`
	// Value after the change.
	NextValue Price `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
	// `id` of the Embedded [Price](ctp:api:type:Price).
	PriceId string `json:"priceId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangePriceChange) MarshalJSON() ([]byte, error) {
	type Alias ChangePriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangePriceChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Product Selection Active](ctp:api:type:StoreChangeProductSelectionAction) update action.
 */
type ChangeProductSelectionActiveChange struct {
	Change string `json:"change"`
	// Reference to the Product Selection which was changed.
	ProductSelection Reference `json:"productSelection"`
	// Value before the change.
	PreviousValue bool `json:"previousValue"`
	// Value after the change.
	NextValue bool `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeProductSelectionActiveChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeProductSelectionActiveChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeProductSelectionActiveChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Quantity](ctp:api:type:InventoryEntryChangeQuantityAction) update action.
 */
type ChangeQuantityChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue InventoryQuantityValue `json:"previousValue"`
	// Value after the change.
	NextValue InventoryQuantityValue `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeQuantityChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeQuantityChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeQuantityChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Quote Request State](ctp:api:type:QuoteRequestChangeQuoteRequestStateAction) update action.
 */
type ChangeQuoteRequestStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue QuoteRequestState `json:"previousValue"`
	// Value after the change.
	NextValue QuoteRequestState `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeQuoteRequestStateChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeQuoteRequestStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeQuoteRequestStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Quote State](ctp:api:type:QuoteChangeQuoteStateAction) update action.
 */
type ChangeQuoteStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue QuoteState `json:"previousValue"`
	// Value after the change.
	NextValue QuoteState `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeQuoteStateChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeQuoteStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeQuoteStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Requires DiscountCode](ctp:api:type:CartDiscountChangeRequiresDiscountCodeAction) update action.
 */
type ChangeRequiresDiscountCodeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue bool `json:"previousValue"`
	// Value after the change.
	NextValue bool `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeRequiresDiscountCodeChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeRequiresDiscountCodeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeRequiresDiscountCodeChange", Alias: (*Alias)(&obj)})
}

type ChangeReviewRatingStatisticsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ReviewRatingStatistics `json:"previousValue"`
	// Value after the change.
	NextValue ReviewRatingStatistics `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeReviewRatingStatisticsChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeReviewRatingStatisticsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeReviewRatingStatisticsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change ShipmentState](ctp:api:type:OrderChangeShipmentStateAction) on Orders.
*	- [Change ShipmentState](ctp:api:type:StagedOrderChangeShipmentStateAction) on Staged Orders.
*
 */
type ChangeShipmentStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ShipmentState `json:"previousValue"`
	// Value after the change.
	NextValue ShipmentState `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeShipmentStateChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeShipmentStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeShipmentStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change ShoppingListLineItem Quantity](ctp:api:type:ShoppingListChangeLineItemQuantityAction) update action.
 */
type ChangeShoppingListLineItemQuantityChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue int `json:"previousValue"`
	// Value after the change.
	NextValue int `json:"nextValue"`
	// Holds information about the updated Shopping List Line Item.
	LineItem ShoppingListLineItemValue `json:"lineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeShoppingListLineItemQuantityChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeShoppingListLineItemQuantityChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeShoppingListLineItemQuantityChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change ShoppingListLineItems Order](ctp:api:type:ShoppingListChangeLineItemsOrderAction) update action.
 */
type ChangeShoppingListLineItemsOrderChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []ShoppingListLineItemValue `json:"previousValue"`
	// Value after the change.
	NextValue []ShoppingListLineItemValue `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeShoppingListLineItemsOrderChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeShoppingListLineItemsOrderChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeShoppingListLineItemsOrderChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change Slug](ctp:api:type:CategoryChangeSlugAction) on Categories.
*	- [Change Slug](ctp:api:type:ProductChangeSlugAction) on Products.
*
 */
type ChangeSlugChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeSlugChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeSlugChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeSlugChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change Sort Order](ctp:api:type:CartDiscountChangeSortOrderAction) on Cart Discounts.
*	- [Change Sort Order](ctp:api:type:ProductDiscountChangeSortOrderAction) on Product Discounts.
*
 */
type ChangeSortOrderChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeSortOrderChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeSortOrderChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeSortOrderChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Stacking Mode](ctp:api:type:CartDiscountChangeStackingModeAction) update action.
 */
type ChangeStackingModeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue StackingMode `json:"previousValue"`
	// Value after the change.
	NextValue StackingMode `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeStackingModeChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeStackingModeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeStackingModeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [ChangeStagedQuoteState](ctp:api:type:StagedQuoteChangeStagedQuoteStateAction) update action.
 */
type ChangeStagedQuoteStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue StagedQuoteState `json:"previousValue"`
	// Value after the change.
	NextValue StagedQuoteState `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeStagedQuoteStateChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeStagedQuoteStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeStagedQuoteStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change State Type](ctp:api:type:StateChangeTypeAction) update action.
 */
type ChangeStateTypeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue StateTypeEnum `json:"previousValue"`
	// Value after the change.
	NextValue StateTypeEnum `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeStateTypeChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeStateTypeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeStateTypeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Status](ctp:api:type:BusinessUnitChangeStatusAction) update action.
 */
type ChangeStatusChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue BusinessUnitStatus `json:"previousValue"`
	// Value after the change.
	NextValue BusinessUnitStatus `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeStatusChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeStatusChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeStatusChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Target](ctp:api:type:CartDiscountChangeTargetAction) update action.
 */
type ChangeTargetChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ChangeTargetChangeValue `json:"previousValue"`
	// Value after the change.
	NextValue ChangeTargetChangeValue `json:"nextValue"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ChangeTargetChange) UnmarshalJSON(data []byte) error {
	type Alias ChangeTargetChange
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.PreviousValue != nil {
		var err error
		obj.PreviousValue, err = mapDiscriminatorChangeTargetChangeValue(obj.PreviousValue)
		if err != nil {
			return err
		}
	}
	if obj.NextValue != nil {
		var err error
		obj.NextValue, err = mapDiscriminatorChangeTargetChangeValue(obj.NextValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTargetChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeTargetChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeTargetChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change TaxCalculationMode](ctp:api:type:StagedOrderChangeTaxCalculationModeAction) update action.
 */
type ChangeTaxCalculationModeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TaxCalculationMode `json:"previousValue"`
	// Value after the change.
	NextValue TaxCalculationMode `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTaxCalculationModeChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeTaxCalculationModeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeTaxCalculationModeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change TaxMode](ctp:api:type:StagedOrderChangeTaxModeAction) update action.
 */
type ChangeTaxModeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TaxMode `json:"previousValue"`
	// Value after the change.
	NextValue TaxMode `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTaxModeChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeTaxModeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeTaxModeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change Tax RoundingMode](ctp:api:type:StagedOrderChangeTaxRoundingModeAction) update action.
 */
type ChangeTaxRoundingModeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue RoundingMode `json:"previousValue"`
	// Value after the change.
	NextValue RoundingMode `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTaxRoundingModeChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeTaxRoundingModeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeTaxRoundingModeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change TextLineItem Name](ctp:api:type:ShoppingListChangeTextLineItemNameAction) update action.
 */
type ChangeTextLineItemNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
	// Holds information about the updated Text Line Item.
	TextLineItem TextLineItemValue `json:"textLineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTextLineItemNameChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeTextLineItemNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeTextLineItemNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change TextLineItem Quantity](ctp:api:type:ShoppingListChangeTextLineItemQuantityAction) update action.
 */
type ChangeTextLineItemQuantityChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue int `json:"previousValue"`
	// Value after the change.
	NextValue int `json:"nextValue"`
	// Holds information about the updated Text Line Item.
	TextLineItem TextLineItemValue `json:"textLineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTextLineItemQuantityChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeTextLineItemQuantityChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeTextLineItemQuantityChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change TextLineItems Order](ctp:api:type:ShoppingListChangeTextLineItemsOrderAction) update action.
 */
type ChangeTextLineItemsOrderChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []TextLineItemValue `json:"previousValue"`
	// Value after the change.
	NextValue []TextLineItemValue `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTextLineItemsOrderChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeTextLineItemsOrderChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeTextLineItemsOrderChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change TransactionInteractionId](ctp:api:type:PaymentChangeTransactionInteractionIdAction) update action.
 */
type ChangeTransactionInteractionIdChange struct {
	Change string `json:"change"`
	// Value after the change.
	PreviousValue string `json:"previousValue"`
	// Value before the change.
	NextValue string `json:"nextValue"`
	// Holds information about the updated Transaction.
	Transaction TransactionChangeValue `json:"transaction"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTransactionInteractionIdChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeTransactionInteractionIdChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeTransactionInteractionIdChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change TransactionState](ctp:api:type:PaymentChangeTransactionStateAction) update action.
 */
type ChangeTransactionStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TransactionState `json:"previousValue"`
	// Value after the change.
	NextValue TransactionState `json:"nextValue"`
	// Holds information about the updated Transaction.
	Transaction TransactionChangeValue `json:"transaction"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTransactionStateChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeTransactionStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeTransactionStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Change TransactionTimestamp](ctp:api:type:PaymentChangeTransactionTimestampAction) update action.
 */
type ChangeTransactionTimestampChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
	// Holds information about the updated Transaction.
	Transaction TransactionChangeValue `json:"transaction"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTransactionTimestampChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeTransactionTimestampChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeTransactionTimestampChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change Value](ctp:api:type:CartDiscountChangeValueAction) on Cart Discounts.
*	- [Change Value](ctp:api:type:ProductDiscountChangeValueAction) on Product Discounts.
*
 */
type ChangeValueChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ChangeValueChangeValue `json:"previousValue"`
	// Value after the change.
	NextValue ChangeValueChangeValue `json:"nextValue"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ChangeValueChange) UnmarshalJSON(data []byte) error {
	type Alias ChangeValueChange
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.PreviousValue != nil {
		var err error
		obj.PreviousValue, err = mapDiscriminatorChangeValueChangeValue(obj.PreviousValue)
		if err != nil {
			return err
		}
	}
	if obj.NextValue != nil {
		var err error
		obj.NextValue, err = mapDiscriminatorChangeValueChangeValue(obj.NextValue)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeValueChange) MarshalJSON() ([]byte, error) {
	type Alias ChangeValueChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ChangeValueChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Move image to position](ctp:api:type:ProductMoveImageToPositionAction) update action.
 */
type MoveImageToPositionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Image `json:"previousValue"`
	// Value after the change.
	NextValue []Image `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MoveImageToPositionChange) MarshalJSON() ([]byte, error) {
	type Alias MoveImageToPositionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "MoveImageToPositionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Publish](ctp:api:type:ProductPublishAction) update action.
 */
type PublishChange struct {
	Change string `json:"change"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PublishChange) MarshalJSON() ([]byte, error) {
	type Alias PublishChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PublishChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Address](ctp:api:type:CustomerRemoveAddressAction) update action.
 */
type RemoveAddressChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Address `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveAddressChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveAddressChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveAddressChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Remove Asset](ctp:api:type:CategoryRemoveAssetAction) on Categories.
*	- [Remove Asset](ctp:api:type:ProductRemoveAssetAction) on Products.
*
 */
type RemoveAssetChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Asset `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveAssetChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveAssetChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveAssetChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Associate](ctp:api:type:BusinessUnitRemoveAssociateAction) update action.
 */
type RemoveAssociateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Associate `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveAssociateChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveAssociateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveAssociateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove AttributeDefinition](ctp:api:type:ProductTypeRemoveAttributeDefinitionAction) update action.
 */
type RemoveAttributeDefinitionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue AttributeDefinition `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveAttributeDefinitionChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveAttributeDefinitionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveAttributeDefinitionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Billing Address ID](ctp:api:type:CustomerRemoveBillingAddressIdAction) update action.
 */
type RemoveBillingAddressIdChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []string `json:"previousValue"`
	// Value after the change.
	NextValue []string `json:"nextValue"`
	// Address removed from `billingAddressesIds`.
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveBillingAddressIdChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveBillingAddressIdChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveBillingAddressIdChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Roles](ctp:api:type:ChannelRemoveRolesAction) update action.
 */
type RemoveChannelRolesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []ChannelRoleEnum `json:"previousValue"`
	// Value after the change.
	NextValue []ChannelRoleEnum `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveChannelRolesChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveChannelRolesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveChannelRolesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove CustomLineItem](ctp:api:type:StagedOrderRemoveCustomLineItemAction) update action.
 */
type RemoveCustomLineItemChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CustomLineItem `json:"previousValue"`
	// Value after the change.
	NextValue CustomLineItem `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveCustomLineItemChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveCustomLineItemChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveCustomLineItemChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Delivery](ctp:api:type:OrderRemoveDeliveryAction) update action.
 */
type RemoveDeliveryItemsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Delivery `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveDeliveryItemsChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveDeliveryItemsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveDeliveryItemsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove DiscountCode](ctp:api:type:StagedOrderRemoveDiscountCodeAction) update action.
 */
type RemoveDiscountCodeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue DiscountCodeInfo `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveDiscountCodeChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveDiscountCodeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveDiscountCodeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove EnumValues from AttributeDefinition](ctp:api:type:ProductTypeRemoveEnumValuesAction) update action.
 */
type RemoveEnumValuesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue EnumValue `json:"previousValue"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition).
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveEnumValuesChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveEnumValuesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveEnumValuesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove FieldDefinition](ctp:api:type:TypeRemoveFieldDefinitionAction) update action.
 */
type RemoveFieldDefinitionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue FieldDefinition `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveFieldDefinitionChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveFieldDefinitionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveFieldDefinitionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove from Category](ctp:api:type:ProductRemoveFromCategoryAction) update action.
 */
type RemoveFromCategoryChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Reference `json:"previousValue"`
	// Value after the change.
	NextValue []Reference `json:"nextValue"`
	// Category from which the Product was removed.
	Category Reference `json:"category"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveFromCategoryChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveFromCategoryChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveFromCategoryChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Image](ctp:api:type:ProductRemoveImageAction) update action.
 */
type RemoveImageChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Image `json:"previousValue"`
	// Value after the change.
	NextValue []Image `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveImageChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveImageChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveImageChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Associate](ctp:api:type:BusinessUnitRemoveAssociateAction) update action on a parent of a Business Unit in cases where [inheritance applies](/../api/associates-overview#conditions-for-inheritance).
 */
type RemoveInheritedAssociateChange struct {
	Change string `json:"change"`
	// The value before the change.
	PreviousValue InheritedAssociate `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveInheritedAssociateChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveInheritedAssociateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveInheritedAssociateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Item Shipping Address](ctp:api:type:OrderRemoveItemShippingAddressAction) update action.
 */
type RemoveItemShippingAddressesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Address `json:"previousValue"`
	// Value after the change.
	NextValue Address `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveItemShippingAddressesChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveItemShippingAddressesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveItemShippingAddressesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove EnumValues from AttributeDefinition](ctp:api:type:ProductTypeRemoveEnumValuesAction) update action.
 */
type RemoveLocalizedEnumValuesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedEnumValue `json:"previousValue"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition).
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveLocalizedEnumValuesChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveLocalizedEnumValuesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveLocalizedEnumValuesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Location](ctp:api:type:ZoneRemoveLocationAction) update action.
 */
type RemoveLocationChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Location `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveLocationChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveLocationChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveLocationChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Parcel From Delivery](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
 */
type RemoveOrderLineItemChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LineItem `json:"previousValue"`
	// Value after the change.
	NextValue LineItem `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveOrderLineItemChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveOrderLineItemChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveOrderLineItemChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Remove Parcel From Delivery](ctp:api:type:OrderRemoveParcelFromDeliveryAction) on Orders.
*	- [Remove Parcel From Delivery](ctp:api:type:StagedOrderRemoveParcelFromDeliveryAction) on Staged Orders.
*
 */
type RemoveParcelFromDeliveryChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Parcel `json:"previousValue"`
	// `id` of the [Delivery](ctp:api:type:Delivery) from which the Parcel was removed.
	DeliveryId string `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveParcelFromDeliveryChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveParcelFromDeliveryChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveParcelFromDeliveryChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Remove Payment](ctp:api:type:OrderRemovePaymentAction) on Orders.
*	- [Remove Payment](ctp:api:type:StagedOrderRemovePaymentAction) on Staged Orders.
*
 */
type RemovePaymentChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue PaymentInfo `json:"previousValue"`
	// Value after the change.
	NextValue PaymentInfo `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemovePaymentChange) MarshalJSON() ([]byte, error) {
	type Alias RemovePaymentChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemovePaymentChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Embedded Price](ctp:api:type:ProductRemovePriceAction) update action.
 */
type RemovePriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Price `json:"previousValue"`
	// Value after the change.
	NextValue Price `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
	// `id` of the Embedded [Price](ctp:api:type:Price).
	PriceId string `json:"priceId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemovePriceChange) MarshalJSON() ([]byte, error) {
	type Alias RemovePriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemovePriceChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Product](ctp:api:type:ProductSelectionRemoveProductAction) update action.
 */
type RemoveProductChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Reference `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveProductChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveProductChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveProductChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Product Selection](ctp:api:type:StoreRemoveProductSelectionAction) update action.
 */
type RemoveProductSelectionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ProductSelectionSetting `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveProductSelectionChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveProductSelectionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveProductSelectionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Update CustomObject](ctp:api:endpoint:/{projectKey}/custom-objects:POST) request when an existing property is removed.
 */
type RemovePropertyChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Path to the property that was removed.
	Path string `json:"path"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemovePropertyChange) MarshalJSON() ([]byte, error) {
	type Alias RemovePropertyChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemovePropertyChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove Shipping Address ID](ctp:api:type:CustomerRemoveShippingAddressIdAction) update action.
 */
type RemoveShippingAddressIdChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []string `json:"previousValue"`
	// Value after the change.
	NextValue []string `json:"nextValue"`
	// Address removed from `shippingAddressesIds`.
	Address Address `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveShippingAddressIdChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveShippingAddressIdChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveShippingAddressIdChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove ShoppingListLineItem](ctp:api:type:ShoppingListRemoveLineItemAction) update action.
 */
type RemoveShoppingListLineItemChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LineItem `json:"previousValue"`
	// Value after the change.
	NextValue LineItem `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveShoppingListLineItemChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveShoppingListLineItemChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveShoppingListLineItemChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove State roles](ctp:api:type:StateRemoveRolesAction) update action.
 */
type RemoveStateRolesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []StateRoleEnum `json:"previousValue"`
	// Value after the change.
	NextValue []StateRoleEnum `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveStateRolesChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveStateRolesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveStateRolesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove TaxRate](ctp:api:type:TaxCategoryRemoveTaxRateAction) update action.
 */
type RemoveTaxRateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TaxRate `json:"previousValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveTaxRateChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveTaxRateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveTaxRateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove TextLineItem](ctp:api:type:ShoppingListRemoveTextLineItemAction) update action.
 */
type RemoveTextLineItemChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TextLineItem `json:"previousValue"`
	// Value after the change.
	NextValue TextLineItem `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveTextLineItemChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveTextLineItemChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveTextLineItemChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Remove ProductVariant](ctp:api:type:ProductRemoveVariantAction) update action.
 */
type RemoveVariantChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Variant `json:"previousValue"`
	// Value after the change.
	NextValue Variant `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RemoveVariantChange) MarshalJSON() ([]byte, error) {
	type Alias RemoveVariantChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RemoveVariantChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Request Quote Renegotiation](ctp:api:type:QuoteRequestQuoteRenegotiationAction) update action.
*
 */
type RequestQuoteRenegotiationChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue QuoteState `json:"previousValue"`
	// Value after the change.
	NextValue QuoteState `json:"nextValue"`
	// Message from the [Buyer](/../api/quotes-overview#buyer) regarding the [Quote](ctp:api:type:Quote) renegotiation request.
	BuyerComment string `json:"buyerComment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RequestQuoteRenegotiationChange) MarshalJSON() ([]byte, error) {
	type Alias RequestQuoteRenegotiationChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "RequestQuoteRenegotiationChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Address](ctp:api:type:ChannelSetAddressAction) update action.
 */
type SetAddressChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Address `json:"previousValue"`
	// Value after the change.
	NextValue Address `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAddressChange) MarshalJSON() ([]byte, error) {
	type Alias SetAddressChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAddressChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Address Custom Field](ctp:api:type:BusinessUnitSetAddressCustomFieldAction) update action.
*
 */
type SetAddressCustomFieldChange struct {
	Change string `json:"change"`
	// [Address](ctp:api:type:Address) which was extended.
	Address Address `json:"address"`
	// Name of the [Custom Field](ctp:api:type:CustomFields).
	Name string `json:"name"`
	// `id` of the referenced [Type](ctp:api:type:Type).
	CustomTypeId string `json:"customTypeId"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAddressCustomFieldChange) MarshalJSON() ([]byte, error) {
	type Alias SetAddressCustomFieldChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAddressCustomFieldChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Address Custom Type](ctp:api:type:BusinessUnitSetAddressCustomTypeAction) update action.
 */
type SetAddressCustomTypeChange struct {
	Change string `json:"change"`
	// [Address](ctp:api:type:Address) which was extended.
	Address Address `json:"address"`
	// Value before the change.
	PreviousValue CustomFields `json:"previousValue"`
	// Value after the change.
	NextValue CustomFields `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAddressCustomTypeChange) MarshalJSON() ([]byte, error) {
	type Alias SetAddressCustomTypeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAddressCustomTypeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set AnonymousId](ctp:api:type:PaymentSetAnonymousIdAction) on Payments.
*	- [Set AnonymousId](ctp:api:type:ShoppingListSetAnonymousIdAction) on Shopping Lists.
*
 */
type SetAnonymousIdChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAnonymousIdChange) MarshalJSON() ([]byte, error) {
	type Alias SetAnonymousIdChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAnonymousIdChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered automatically due to a user-initiated change.
 */
type SetApplicationVersionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue int `json:"previousValue"`
	// Value after the change.
	NextValue int `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetApplicationVersionChange) MarshalJSON() ([]byte, error) {
	type Alias SetApplicationVersionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetApplicationVersionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Asset CustomField](ctp:api:type:CategorySetAssetCustomFieldAction) on Categories.
*	- [Set Asset CustomField](ctp:api:type:ProductSetAssetCustomFieldAction) on Products.
*
 */
type SetAssetCustomFieldChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// `id` of the referenced [Type](ctp:api:type:Type).
	CustomTypeId string `json:"customTypeId"`
	// Information about the updated Asset.
	Asset AssetChangeValue `json:"asset"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAssetCustomFieldChange) MarshalJSON() ([]byte, error) {
	type Alias SetAssetCustomFieldChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAssetCustomFieldChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Asset Custom Type](ctp:api:type:CategorySetAssetCustomTypeAction) on Categories.
*	- [Set Asset Custom Type](ctp:api:type:ProductSetAssetCustomTypeAction) on Products.
*
 */
type SetAssetCustomTypeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CustomFields `json:"previousValue"`
	// Value after the change.
	NextValue CustomFields `json:"nextValue"`
	// Information about the updated Asset.
	Asset AssetChangeValue `json:"asset"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAssetCustomTypeChange) MarshalJSON() ([]byte, error) {
	type Alias SetAssetCustomTypeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAssetCustomTypeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Asset Description](ctp:api:type:CategorySetAssetDescriptionAction) on Categories.
*	- [Set Asset Description](ctp:api:type:ProductSetAssetDescriptionAction) on Products.
*
 */
type SetAssetDescriptionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
	// Information about the updated Asset.
	Asset AssetChangeValue `json:"asset"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAssetDescriptionChange) MarshalJSON() ([]byte, error) {
	type Alias SetAssetDescriptionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAssetDescriptionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Asset Key](ctp:api:type:CategorySetAssetKeyAction) on Categories.
*	- [Set Asset Key](ctp:api:type:ProductSetAssetKeyAction) on Products.
*
 */
type SetAssetKeyChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
	// Information about the updated Asset.
	Asset AssetChangeValue `json:"asset"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAssetKeyChange) MarshalJSON() ([]byte, error) {
	type Alias SetAssetKeyChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAssetKeyChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Asset Sources](ctp:api:type:CategorySetAssetSourcesAction) on Categories.
*	- [Set Asset Sources](ctp:api:type:ProductSetAssetSourcesAction) on Products.
*
 */
type SetAssetSourcesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []AssetSource `json:"previousValue"`
	// Value after the change.
	NextValue []AssetSource `json:"nextValue"`
	// Information about the updated Asset.
	Asset AssetChangeValue `json:"asset"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAssetSourcesChange) MarshalJSON() ([]byte, error) {
	type Alias SetAssetSourcesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAssetSourcesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change Asset Tags](ctp:api:type:CategorySetAssetTagsAction) on Categories.
*	- [Change Asset Tags](ctp:api:type:ProductSetAssetTagsAction) on Products.
*
 */
type SetAssetTagsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []string `json:"previousValue"`
	// Value after the change.
	NextValue []string `json:"nextValue"`
	// Information about the updated Asset.
	Asset AssetChangeValue `json:"asset"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAssetTagsChange) MarshalJSON() ([]byte, error) {
	type Alias SetAssetTagsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAssetTagsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Attribute](ctp:api:type:ProductSetAttributeAction) update action.
 */
type SetAttributeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue AttributeValue `json:"previousValue"`
	// Value after the change.
	NextValue AttributeValue `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAttributeChange) MarshalJSON() ([]byte, error) {
	type Alias SetAttributeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAttributeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set AuthenticationMode](ctp:api:type:CustomerSetAuthenticationModeAction) update action.
 */
type SetAuthenticationModeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue AuthenticationMode `json:"previousValue"`
	// Value after the change.
	NextValue AuthenticationMode `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAuthenticationModeChange) MarshalJSON() ([]byte, error) {
	type Alias SetAuthenticationModeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAuthenticationModeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Author Name](ctp:api:type:ReviewSetAuthorNameAction) update action.
 */
type SetAuthorNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAuthorNameChange) MarshalJSON() ([]byte, error) {
	type Alias SetAuthorNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetAuthorNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Billing Address](ctp:api:type:OrderSetBillingAddressAction) on Orders.
*	- [Set Billing Address](ctp:api:type:StagedOrderSetBillingAddressAction) on Staged Orders.
*
 */
type SetBillingAddressChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Address `json:"previousValue"`
	// Value after the change.
	NextValue Address `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetBillingAddressChange) MarshalJSON() ([]byte, error) {
	type Alias SetBillingAddressChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetBillingAddressChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Cart Predicate](ctp:api:type:DiscountCodeSetCartPredicateAction) update action.
 */
type SetCartPredicateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCartPredicateChange) MarshalJSON() ([]byte, error) {
	type Alias SetCartPredicateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCartPredicateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Category Order Hint](ctp:api:type:ProductSetCategoryOrderHintAction) update action.
 */
type SetCategoryOrderHintChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CategoryOrderHints `json:"previousValue"`
	// Value after the change.
	NextValue CategoryOrderHints `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
	// `id` of the updated [Category](ctp:api:type:Category).
	CategoryId string `json:"categoryId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCategoryOrderHintChange) MarshalJSON() ([]byte, error) {
	type Alias SetCategoryOrderHintChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCategoryOrderHintChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Roles](ctp:api:type:ChannelSetRolesAction) update action.
 */
type SetChannelRolesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []ChannelRoleEnum `json:"previousValue"`
	// Value after the change.
	NextValue []ChannelRoleEnum `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetChannelRolesChange) MarshalJSON() ([]byte, error) {
	type Alias SetChannelRolesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetChannelRolesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Company Name](ctp:api:type:CustomerSetCompanyNameAction) update action.
 */
type SetCompanyNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCompanyNameChange) MarshalJSON() ([]byte, error) {
	type Alias SetCompanyNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCompanyNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Contact Email](ctp:api:type:BusinessUnitSetContactEmailAction) update action.
 */
type SetContactEmailChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetContactEmailChange) MarshalJSON() ([]byte, error) {
	type Alias SetContactEmailChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetContactEmailChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Countries](ctp:api:type:StoreSetCountriesAction) update action.
 */
type SetCountriesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []StoreCountry `json:"previousValue"`
	// Value after the change.
	NextValue []StoreCountry `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCountriesChange) MarshalJSON() ([]byte, error) {
	type Alias SetCountriesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCountriesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Country](ctp:api:type:StagedOrderSetCountryAction) update action.
 */
type SetCountryChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCountryChange) MarshalJSON() ([]byte, error) {
	type Alias SetCountryChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCountryChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set CustomField](ctp:api:type:CartDiscountSetCustomFieldAction) on Cart Discounts.
*	- [Set CustomField](ctp:api:type:CategorySetCustomFieldAction) on Categories.
*	- [Set CustomField](ctp:api:type:ChannelSetCustomFieldAction) on Channels.
*	- [Set CustomField](ctp:api:type:CustomerSetCustomFieldAction) on Customers.
*	- [Set CustomField](ctp:api:type:CustomerGroupSetCustomFieldAction) on Customer Groups.
*	- [Set CustomField](ctp:api:type:DiscountCodeSetCustomFieldAction) on Discount Codes.
*	- [Set CustomField](ctp:api:type:InventoryEntrySetCustomFieldAction) on Inventories.
*	- [Set CustomField](ctp:api:type:OrderSetCustomFieldAction) on Orders.
*	- [Set CustomField](ctp:api:type:OrderEditSetCustomFieldAction) on Order Edits.
*	- [Set CustomField](ctp:api:type:PaymentSetCustomFieldAction) on Payments.
*	- [Set CustomField](ctp:api:type:ProductSelectionSetCustomFieldAction) on Product Selections.
*	- [Set CustomField](ctp:api:type:QuoteSetCustomFieldAction) on Quotes.
*	- [Set CustomField](ctp:api:type:QuoteRequestSetCustomFieldAction) on Quote Requests.
*	- [Set CustomField](ctp:api:type:ReviewSetCustomFieldAction) on Reviews.
*	- [Set CustomField](ctp:api:type:ShoppingListSetCustomFieldAction) on Shopping Lists.
*	- [Set CustomField](ctp:api:type:StagedOrderSetCustomFieldAction) on Staged Orders.
*	- [Set CustomField](ctp:api:type:StagedQuoteSetCustomFieldAction) on Staged Quotes.
*	- [Set CustomField](ctp:api:type:StoreSetCustomFieldAction) on Stores.
*
 */
type SetCustomFieldChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// `id` of the referenced [Type](ctp:api:type:Type).
	CustomTypeId string `json:"customTypeId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomFieldChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomFieldChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomFieldChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set CustomLineItem Custom Type](ctp:api:type:OrderSetCustomLineItemCustomFieldAction) on Orders.
*	- [Set CustomLineItem Custom Type](ctp:api:type:StagedOrderSetCustomLineItemCustomFieldAction) on Staged Orders.
*
 */
type SetCustomLineItemCustomFieldChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// Name of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItem LocalizedString `json:"customLineItem"`
	// `id` of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomLineItemCustomFieldChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomLineItemCustomFieldChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomLineItemCustomFieldChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set CustomLineItem CustomField](ctp:api:type:OrderSetCustomLineItemCustomTypeAction) on Orders.
*	- [Set CustomLineItem CustomField](ctp:api:type:StagedOrderSetCustomLineItemCustomTypeAction) on Staged Orders.
*
 */
type SetCustomLineItemCustomTypeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CustomFields `json:"previousValue"`
	// Value after the change.
	NextValue CustomFields `json:"nextValue"`
	// Name of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItem LocalizedString `json:"customLineItem"`
	// `id` of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomLineItemCustomTypeChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomLineItemCustomTypeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomLineItemCustomTypeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set CustomLineItem Money](ctp:api:type:StagedOrderChangeCustomLineItemMoneyAction) update action.
 */
type SetCustomLineItemMoneyChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Money `json:"previousValue"`
	// Value after the change.
	NextValue Money `json:"nextValue"`
	// Name of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItem LocalizedString `json:"customLineItem"`
	// `id` of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomLineItemMoneyChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomLineItemMoneyChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomLineItemMoneyChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set CustomLineItem ShippingDetails](ctp:api:type:OrderSetCustomLineItemShippingDetailsAction) on Orders.
*	- [Set CustomLineItem ShippingDetails](ctp:api:type:StagedOrderSetCustomLineItemShippingDetailsAction) on Staged Orders.
*
 */
type SetCustomLineItemShippingDetailsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ItemShippingDetails `json:"previousValue"`
	// Value after the change.
	NextValue ItemShippingDetails `json:"nextValue"`
	// `id` of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomLineItemShippingDetailsChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomLineItemShippingDetailsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomLineItemShippingDetailsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set CustomLineItem TaxAmount](ctp:api:type:StagedOrderSetCustomLineItemTaxAmountAction) update action.
 */
type SetCustomLineItemTaxAmountChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TaxRate `json:"previousValue"`
	// Value after the change.
	NextValue TaxRate `json:"nextValue"`
	// Name of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItem LocalizedString `json:"customLineItem"`
	// `id` of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
	// `"ExternalAmount"`
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomLineItemTaxAmountChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomLineItemTaxAmountChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomLineItemTaxAmountChange", Alias: (*Alias)(&obj)})
}

type SetCustomLineItemTaxCategoryChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Reference `json:"previousValue"`
	// Value after the change.
	NextValue Reference `json:"nextValue"`
	// Name of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItem LocalizedString `json:"customLineItem"`
	// `id` of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomLineItemTaxCategoryChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomLineItemTaxCategoryChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomLineItemTaxCategoryChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set CustomLineItem TaxRate](ctp:api:type:StagedOrderSetCustomLineItemTaxRateAction) update action.
 */
type SetCustomLineItemTaxRateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TaxRate `json:"previousValue"`
	// Value after the change.
	NextValue TaxRate `json:"nextValue"`
	// Name of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItem LocalizedString `json:"customLineItem"`
	// `id` of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
	// `"External"`
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomLineItemTaxRateChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomLineItemTaxRateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomLineItemTaxRateChange", Alias: (*Alias)(&obj)})
}

type SetCustomLineItemTaxedPriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Money `json:"previousValue"`
	// Value after the change.
	NextValue Money `json:"nextValue"`
	// Name of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItem LocalizedString `json:"customLineItem"`
	// `id` of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomLineItemTaxedPriceChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomLineItemTaxedPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomLineItemTaxedPriceChange", Alias: (*Alias)(&obj)})
}

type SetCustomLineItemTotalPriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Money `json:"previousValue"`
	// Value after the change.
	NextValue Money `json:"nextValue"`
	// Name of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItem LocalizedString `json:"customLineItem"`
	// `id` of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomLineItemTotalPriceChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomLineItemTotalPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomLineItemTotalPriceChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Custom ShippingMethod](ctp:api:type:StagedOrderSetCustomShippingMethodAction) update action.
 */
type SetCustomShippingMethodChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CustomShippingMethodChangeValue `json:"previousValue"`
	// Value after the change.
	NextValue CustomShippingMethodChangeValue `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomShippingMethodChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomShippingMethodChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomShippingMethodChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Custom Type](ctp:api:type:CartDiscountSetCustomTypeAction) on Cart Discounts.
*	- [Set Custom Type](ctp:api:type:CategorySetCustomTypeAction) on Categories.
*	- [Set Custom Type](ctp:api:type:ChannelSetCustomTypeAction) on Channels.
*	- [Set Custom Type](ctp:api:type:CustomerSetCustomTypeAction) on Customers.
*	- [Set Custom Type](ctp:api:type:CustomerGroupSetCustomTypeAction) on Customer Groups.
*	- [Set Custom Type](ctp:api:type:DiscountCodeSetCustomTypeAction) on Discount Codes.
*	- [Set Custom Type](ctp:api:type:InventoryEntrySetCustomTypeAction) on Inventories.
*	- [Set Custom Type](ctp:api:type:OrderSetCustomTypeAction) on Orders.
*	- [Set Custom Type](ctp:api:type:OrderEditSetCustomTypeAction) on Order Edits.
*	- [Set Custom Type](ctp:api:type:StagedOrderSetCustomTypeAction) on Staged Orders.
*	- [Set Custom Type](ctp:api:type:PaymentSetCustomTypeAction) on Payments.
*	- [Set Custom Type](ctp:api:type:ProductSelectionSetCustomTypeAction) on Product Selections.
*	- [Set Custom Type](ctp:api:type:QuoteSetCustomTypeAction) on Quotes.
*	- [Set Custom Type](ctp:api:type:StagedQuoteSetCustomTypeAction) on Staged Quotes.
*	- [Set Custom Type](ctp:api:type:QuoteRequestSetCustomTypeAction) on Quote Requests.
*	- [Set Custom Type](ctp:api:type:ReviewSetCustomTypeAction) on Reviews.
*	- [Set Custom Type](ctp:api:type:ShoppingListSetCustomTypeAction) on Shopping Lists.
*	- [Set Custom Type](ctp:api:type:StoreSetCustomTypeAction) on Stores.
*
 */
type SetCustomTypeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CustomFields `json:"previousValue"`
	// Value after the change.
	NextValue CustomFields `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomTypeChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomTypeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomTypeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Customer](ctp:api:type:PaymentSetCustomerAction) on Payments.
*	- [Set Customer](ctp:api:type:ReviewSetCustomerAction) on Reviews.
*	- [Set Customer](ctp:api:type:ShoppingListSetCustomerAction) on Shopping Lists.
*
 */
type SetCustomerChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Reference `json:"previousValue"`
	// Value after the change.
	NextValue Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomerChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomerChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomerChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Customer Email](ctp:api:type:OrderSetCustomerEmailAction) on Orders.
*	- [Set Customer Email](ctp:api:type:StagedOrderSetCustomerEmailAction) on Staged Orders.
*
 */
type SetCustomerEmailChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomerEmailChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomerEmailChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomerEmailChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set CustomerGroup](ctp:api:type:CustomerSetCustomerGroupAction) update action.
 */
type SetCustomerGroupChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Reference `json:"previousValue"`
	// Value after the change.
	NextValue Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomerGroupChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomerGroupChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomerGroupChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Customer ID](ctp:api:type:OrderSetCustomerIdAction) on Orders.
*	- [Set Customer ID](ctp:api:type:StagedOrderSetCustomerIdAction) on Staged Orders.
*
 */
type SetCustomerIdChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomerIdChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomerIdChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomerIdChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Customer Number](ctp:api:type:CustomerSetCustomerNumberAction) update action.
 */
type SetCustomerNumberChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetCustomerNumberChange) MarshalJSON() ([]byte, error) {
	type Alias SetCustomerNumberChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetCustomerNumberChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Date of Birth](ctp:api:type:CustomerSetDateOfBirthAction) update action.
 */
type SetDateOfBirthChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDateOfBirthChange) MarshalJSON() ([]byte, error) {
	type Alias SetDateOfBirthChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetDateOfBirthChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Default Billing Address](ctp:api:type:CustomerSetDefaultBillingAddressAction) update action.
 */
type SetDefaultBillingAddressChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Address `json:"previousValue"`
	// Value after the change.
	NextValue Address `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDefaultBillingAddressChange) MarshalJSON() ([]byte, error) {
	type Alias SetDefaultBillingAddressChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetDefaultBillingAddressChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Default Shipping Address](ctp:api:type:CustomerSetDefaultShippingAddressAction) update action.
 */
type SetDefaultShippingAddressChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Address `json:"previousValue"`
	// Value after the change.
	NextValue Address `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDefaultShippingAddressChange) MarshalJSON() ([]byte, error) {
	type Alias SetDefaultShippingAddressChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetDefaultShippingAddressChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set DeleteDaysAfterLastModification](ctp:api:type:ShoppingListSetDeleteDaysAfterLastModificationAction) update action.
 */
type SetDeleteDaysAfterLastModificationChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue int `json:"previousValue"`
	// Value after the change.
	NextValue int `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDeleteDaysAfterLastModificationChange) MarshalJSON() ([]byte, error) {
	type Alias SetDeleteDaysAfterLastModificationChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetDeleteDaysAfterLastModificationChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set DeliveryAddress](ctp:api:type:OrderSetDeliveryAddressAction) on Orders.
*	- [Set DeliveryAddress](ctp:api:type:StagedOrderSetDeliveryAddressAction) on Staged Orders.
*
 */
type SetDeliveryAddressChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Address `json:"previousValue"`
	// Value after the change.
	NextValue Address `json:"nextValue"`
	// `id` of the updated [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDeliveryAddressChange) MarshalJSON() ([]byte, error) {
	type Alias SetDeliveryAddressChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetDeliveryAddressChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Delivery Items](ctp:api:type:OrderSetDeliveryItemsAction) on Orders.
*	- [Set Delivery Items](ctp:api:type:StagedOrderSetDeliveryItemsAction) on Staged Orders.
*
 */
type SetDeliveryItemsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []DeliveryItem `json:"previousValue"`
	// Value after the change.
	NextValue []DeliveryItem `json:"nextValue"`
	// `id` of the updated [Delivery](ctp:api:type:Delivery).
	DeliveryId string `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDeliveryItemsChange) MarshalJSON() ([]byte, error) {
	type Alias SetDeliveryItemsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetDeliveryItemsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Description](ctp:api:type:TaxCategorySetDescriptionAction) on Tax Categories.
*	- [Set Description](ctp:api:type:ZoneSetDescriptionAction) on Zones.
*
 */
type SetDescriptionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDescriptionChange) MarshalJSON() ([]byte, error) {
	type Alias SetDescriptionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetDescriptionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Discounted Embedded Price](ctp:api:type:ProductSetDiscountedPriceAction) update action.
 */
type SetDiscountedPriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Price `json:"previousValue"`
	// Value after the change.
	NextValue Price `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
	// `sku` or `key` of the updated [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
	// `id` of the Embedded [Price](ctp:api:type:Price).
	PriceId string `json:"priceId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDiscountedPriceChange) MarshalJSON() ([]byte, error) {
	type Alias SetDiscountedPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetDiscountedPriceChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Distribution Channels](ctp:api:type:StoreSetDistributionChannelsAction) update action.
 */
type SetDistributionChannelsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Reference `json:"previousValue"`
	// Value after the change.
	NextValue []Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDistributionChannelsChange) MarshalJSON() ([]byte, error) {
	type Alias SetDistributionChannelsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetDistributionChannelsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set ExpectedDelivery](ctp:api:type:InventoryEntrySetExpectedDeliveryAction) update action.
 */
type SetExpectedDeliveryChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetExpectedDeliveryChange) MarshalJSON() ([]byte, error) {
	type Alias SetExpectedDeliveryChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetExpectedDeliveryChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set External ID](ctp:api:type:CategorySetExternalIdAction) on Categories.
*	- [Set External ID](ctp:api:type:CustomerSetExternalIdAction) on Customers.
*
 */
type SetExternalIdChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetExternalIdChange) MarshalJSON() ([]byte, error) {
	type Alias SetExternalIdChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetExternalIdChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set First Name](ctp:api:type:CustomerSetFirstNameAction) update action.
 */
type SetFirstNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetFirstNameChange) MarshalJSON() ([]byte, error) {
	type Alias SetFirstNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetFirstNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set GeoLocation](ctp:api:type:ChannelSetGeoLocationAction) update action.
 */
type SetGeoLocationChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue GeoLocation `json:"previousValue"`
	// Value after the change.
	NextValue GeoLocation `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetGeoLocationChange) MarshalJSON() ([]byte, error) {
	type Alias SetGeoLocationChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetGeoLocationChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Image Label](ctp:api:type:ProductSetImageLabelAction) update action.
 */
type SetImageLabelChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Image `json:"previousValue"`
	// Value after the change.
	NextValue Image `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetImageLabelChange) MarshalJSON() ([]byte, error) {
	type Alias SetImageLabelChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetImageLabelChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set AttributeDefinition InputTip](ctp:api:type:ProductTypeSetInputTipAction) update action.
 */
type SetInputTipChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
	// Name of the updated [AttributeDefinition](ctp:api:type:AttributeDefinition).
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetInputTipChange) MarshalJSON() ([]byte, error) {
	type Alias SetInputTipChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetInputTipChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set InterfaceId](ctp:api:type:PaymentSetInterfaceIdAction) update action.
 */
type SetInterfaceIdChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetInterfaceIdChange) MarshalJSON() ([]byte, error) {
	type Alias SetInterfaceIdChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetInterfaceIdChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered automatically due to a user-initiated change.
 */
type SetIsValidChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue bool `json:"previousValue"`
	// Value after the change.
	NextValue bool `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetIsValidChange) MarshalJSON() ([]byte, error) {
	type Alias SetIsValidChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetIsValidChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Key](ctp:api:type:CartDiscountSetKeyAction) on Cart Discounts.
*	- [Set Key](ctp:api:type:CategorySetKeyAction) on Categories.
*	- [Set Key](ctp:api:type:CustomerSetKeyAction) on Customers.
*	- [Set Key](ctp:api:type:CustomerGroupSetKeyAction) on Customer Groups.
*	- [Set Key](ctp:api:type:PaymentSetKeyAction) on Payments.
*	- [Set Key](ctp:api:type:ProductSetKeyAction) on Products.
*	- [Set Key](ctp:api:type:ProductDiscountSetKeyAction) on Product Discounts.
*	- [Set Key](ctp:api:type:ProductSelectionSetKeyAction) on Product Selections.
*	- [Set Key](ctp:api:type:ProductTypeSetKeyAction) on Product Types.
*	- [Set Key](ctp:api:type:ReviewSetKeyAction) on Reviews.
*	- [Set Key](ctp:api:type:ShoppingListSetKeyAction) on Shopping Lists.
*	- [Set Key](ctp:api:type:ZoneSetKeyAction) on Zones.
*
 */
type SetKeyChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetKeyChange) MarshalJSON() ([]byte, error) {
	type Alias SetKeyChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetKeyChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by [Set Languages](ctp:api:type:StoreSetLanguagesAction) update action.
 */
type SetLanguagesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []string `json:"previousValue"`
	// Value after the change.
	NextValue []string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLanguagesChange) MarshalJSON() ([]byte, error) {
	type Alias SetLanguagesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLanguagesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by [Set Last Name](ctp:api:type:CustomerSetLastNameAction) update action.
 */
type SetLastNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLastNameChange) MarshalJSON() ([]byte, error) {
	type Alias SetLastNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLastNameChange", Alias: (*Alias)(&obj)})
}

type SetLineItemDeactivatedAtChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
	// Holds information about the updated Shopping List Line Item.
	LineItem ShoppingListLineItemValue `json:"lineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemDeactivatedAtChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemDeactivatedAtChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemDeactivatedAtChange", Alias: (*Alias)(&obj)})
}

type SetLineItemDiscountedPriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue DiscountedLineItemPrice `json:"previousValue"`
	// Value after the change.
	NextValue DiscountedLineItemPrice `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
	// `sku` or `key` of the updated [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemDiscountedPriceChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemDiscountedPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemDiscountedPriceChange", Alias: (*Alias)(&obj)})
}

type SetLineItemDiscountedPricePerQuantityChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue DiscountedLineItemPriceForQuantity `json:"previousValue"`
	// Value after the change.
	NextValue DiscountedLineItemPriceForQuantity `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
	// `sku` or `key` of the updated [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemDiscountedPricePerQuantityChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemDiscountedPricePerQuantityChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemDiscountedPricePerQuantityChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set LineItem DistributionChannel](ctp:api:type:StagedOrderSetLineItemDistributionChannelAction) update action.
 */
type SetLineItemDistributionChannelChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Reference `json:"previousValue"`
	// Value after the change.
	NextValue Reference `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
	// `sku` or `key` of the updated [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemDistributionChannelChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemDistributionChannelChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemDistributionChannelChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set LineItem Price](ctp:api:type:StagedOrderSetLineItemPriceAction) update action.
 */
type SetLineItemPriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Price `json:"previousValue"`
	// Value after the change.
	NextValue Price `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the updated Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemPriceChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemPriceChange", Alias: (*Alias)(&obj)})
}

type SetLineItemProductKeyChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
	// `id` of the updated [LineItem](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
	// `sku` or `key` of the updated [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemProductKeyChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemProductKeyChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemProductKeyChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered automatically due to a user-initiated change.
 */
type SetLineItemProductSlugChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the updated Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
	// `sku` or `key` of the updated [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemProductSlugChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemProductSlugChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemProductSlugChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set LineItem ShippingDetails](ctp:api:type:OrderSetLineItemShippingDetailsAction) on Orders.
*	- [Set LineItem ShippingDetails](ctp:api:type:StagedOrderSetLineItemShippingDetailsAction) on Staged Orders.
*
 */
type SetLineItemShippingDetailsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ItemShippingDetails `json:"previousValue"`
	// Value after the change.
	NextValue ItemShippingDetails `json:"nextValue"`
	// `id` of the updated [LineItem](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemShippingDetailsChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemShippingDetailsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemShippingDetailsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set LineItem TaxAmount](ctp:api:type:StagedOrderSetLineItemTaxAmountAction) update action.
 */
type SetLineItemTaxAmountChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TaxRate `json:"previousValue"`
	// Value after the change.
	NextValue TaxRate `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
	// `sku` or `key` of the [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
	// `"ExternalAmount"`
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemTaxAmountChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemTaxAmountChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemTaxAmountChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set LineItemTaxRate](ctp:api:type:StagedOrderSetLineItemTaxRateAction) update action.
 */
type SetLineItemTaxRateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TaxRate `json:"previousValue"`
	// Value after the change.
	NextValue TaxRate `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
	// `sku` or `key` of the [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
	// `"External"`
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemTaxRateChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemTaxRateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemTaxRateChange", Alias: (*Alias)(&obj)})
}

type SetLineItemTaxedPriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TaxedItemPrice `json:"previousValue"`
	// Value after the change.
	NextValue TaxedItemPrice `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
	// `id` of the updated [LineItem](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemTaxedPriceChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemTaxedPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemTaxedPriceChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set LineItemTotalPrice](ctp:api:type:StagedOrderSetLineItemTotalPriceAction) update action.
 */
type SetLineItemTotalPriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Money `json:"previousValue"`
	// Value after the change.
	NextValue Money `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the updated Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLineItemTotalPriceChange) MarshalJSON() ([]byte, error) {
	type Alias SetLineItemTotalPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLineItemTotalPriceChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Locale](ctp:api:type:CustomerSetLocaleAction) on Customers.
*	- [Set Locale](ctp:api:type:OrderSetLocaleAction) on Orders.
*	- [Set Locale](ctp:api:type:StagedOrderSetLocaleAction) on Staged Orders.
*	- [Set Locale](ctp:api:type:ReviewSetLocaleAction) on Reviews.
*
 */
type SetLocaleChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLocaleChange) MarshalJSON() ([]byte, error) {
	type Alias SetLocaleChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLocaleChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Description](ctp:api:type:CartDiscountSetDescriptionAction) on Cart Discounts.
*	- [Set Description](ctp:api:type:CategorySetDescriptionAction) on Categories.
*	- [Set Description](ctp:api:type:DiscountCodeSetDescriptionAction) on Discount Codes.
*	- [Set Description](ctp:api:type:ProductSetDescriptionAction) on Products.
*	- [Set Description](ctp:api:type:ProductDiscountSetDescriptionAction) on Product Discounts.
*	- [Set Description](ctp:api:type:ShoppingListSetDescriptionAction) on Shopping Lists.
*	- [Set Description](ctp:api:type:StateSetDescriptionAction) on States.
*	- [Set Description](ctp:api:type:TypeSetDescriptionAction) on Types.
*
 */
type SetLocalizedDescriptionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLocalizedDescriptionChange) MarshalJSON() ([]byte, error) {
	type Alias SetLocalizedDescriptionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLocalizedDescriptionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Max Applications](ctp:api:type:DiscountCodeSetMaxApplicationsAction) update action.
 */
type SetMaxApplicationsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue int `json:"previousValue"`
	// Value after the change.
	NextValue int `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetMaxApplicationsChange) MarshalJSON() ([]byte, error) {
	type Alias SetMaxApplicationsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetMaxApplicationsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Max Applications Per Customer](ctp:api:type:DiscountCodeSetMaxApplicationsPerCustomerAction) update action.
 */
type SetMaxApplicationsPerCustomerChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue int `json:"previousValue"`
	// Value after the change.
	NextValue int `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetMaxApplicationsPerCustomerChange) MarshalJSON() ([]byte, error) {
	type Alias SetMaxApplicationsPerCustomerChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetMaxApplicationsPerCustomerChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Meta Description](ctp:api:type:CategorySetMetaDescriptionAction) on Categories.
*	- [Set Meta Description](ctp:api:type:ProductSetMetaDescriptionAction) on Products.
*
 */
type SetMetaDescriptionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetMetaDescriptionChange) MarshalJSON() ([]byte, error) {
	type Alias SetMetaDescriptionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetMetaDescriptionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Meta Keywords](ctp:api:type:CategorySetMetaKeywordsAction) on Categories.
*	- [Set Meta Keywords](ctp:api:type:ProductSetMetaKeywordsAction) on Products.
*
 */
type SetMetaKeywordsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetMetaKeywordsChange) MarshalJSON() ([]byte, error) {
	type Alias SetMetaKeywordsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetMetaKeywordsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Meta Title](ctp:api:type:CategorySetMetaTitleAction) on Categories.
*	- [Set Meta Title](ctp:api:type:ProductSetMetaTitleAction) on Products.
*
 */
type SetMetaTitleChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetMetaTitleChange) MarshalJSON() ([]byte, error) {
	type Alias SetMetaTitleChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetMetaTitleChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set MethodInfoInterface](ctp:api:type:PaymentSetMethodInfoInterfaceAction) update action.
 */
type SetMethodInfoInterfaceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetMethodInfoInterfaceChange) MarshalJSON() ([]byte, error) {
	type Alias SetMethodInfoInterfaceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetMethodInfoInterfaceChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set MethodInfoInterface](ctp:api:type:PaymentSetMethodInfoMethodAction) update action.
 */
type SetMethodInfoMethodChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetMethodInfoMethodChange) MarshalJSON() ([]byte, error) {
	type Alias SetMethodInfoMethodChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetMethodInfoMethodChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set MethodInfoName](ctp:api:type:PaymentSetMethodInfoNameAction) update action.
 */
type SetMethodInfoNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetMethodInfoNameChange) MarshalJSON() ([]byte, error) {
	type Alias SetMethodInfoNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetMethodInfoNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Middle Name](ctp:api:type:CustomerSetMiddleNameAction) update action.
 */
type SetMiddleNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetMiddleNameChange) MarshalJSON() ([]byte, error) {
	type Alias SetMiddleNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetMiddleNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Name](ctp:api:type:AssociateRoleSetNameAction) update action.
 */
type SetNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetNameChange) MarshalJSON() ([]byte, error) {
	type Alias SetNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Name](ctp:api:type:DiscountCodeSetNameAction) on Discount Codes.
*	- [Set State Name](ctp:api:type:StateSetNameAction) on States.
*	- [Set Name](ctp:api:type:StoreSetNameAction) on Stores.
*
 */
type SetLocalizedNameChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetLocalizedNameChange) MarshalJSON() ([]byte, error) {
	type Alias SetLocalizedNameChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetLocalizedNameChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set LineItem CustomField](ctp:api:type:OrderSetLineItemCustomFieldAction) on Orders.
*	- [Set LineItem CustomField](ctp:api:type:StagedOrderSetLineItemCustomFieldAction) on Staged Orders.
*
 */
type SetOrderLineItemCustomFieldChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
	// `id` of the referenced [Type](ctp:api:type:Type).
	CustomTypeId string `json:"customTypeId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// Name of the [Product](ctp:api:type:Product) the Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
	// `sku` or `key` of the [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetOrderLineItemCustomFieldChange) MarshalJSON() ([]byte, error) {
	type Alias SetOrderLineItemCustomFieldChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetOrderLineItemCustomFieldChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set LineItem Custom Type](ctp:api:type:OrderSetLineItemCustomTypeAction) on Orders.
*	- [Set LineItem Custom Type](ctp:api:type:StagedOrderSetLineItemCustomTypeAction) on Staged Orders.
*
 */
type SetOrderLineItemCustomTypeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CustomFields `json:"previousValue"`
	// Value after the change.
	NextValue CustomFields `json:"nextValue"`
	// Name of the [Product](ctp:api:type:Product) the updated Line Item is based on.
	LineItem LocalizedString `json:"lineItem"`
	// `sku` or `key` of the [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetOrderLineItemCustomTypeChange) MarshalJSON() ([]byte, error) {
	type Alias SetOrderLineItemCustomTypeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetOrderLineItemCustomTypeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Order Number](ctp:api:type:OrderSetOrderNumberAction) on Orders.
*	- [Set Order Number](ctp:api:type:StagedOrderSetOrderNumberAction) on Staged Order.
*
 */
type SetOrderNumberChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetOrderNumberChange) MarshalJSON() ([]byte, error) {
	type Alias SetOrderNumberChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetOrderNumberChange", Alias: (*Alias)(&obj)})
}

type SetOrderTaxedPriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TaxedItemPrice `json:"previousValue"`
	// Value after the change.
	NextValue TaxedItemPrice `json:"nextValue"`
	TaxMode   TaxMode        `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetOrderTaxedPriceChange) MarshalJSON() ([]byte, error) {
	type Alias SetOrderTaxedPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetOrderTaxedPriceChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered automatically due to a user-initiated change.
 */
type SetOrderTotalPriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Money `json:"previousValue"`
	// Value after the change.
	NextValue Money `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetOrderTotalPriceChange) MarshalJSON() ([]byte, error) {
	type Alias SetOrderTotalPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetOrderTotalPriceChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set OrderTotalTax](ctp:api:type:StagedOrderSetOrderTotalTaxAction) update action.
 */
type SetOrderTotalTaxChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Money `json:"previousValue"`
	// Value after the change.
	NextValue Money `json:"nextValue"`
	// `"ExternalAmount"`
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetOrderTotalTaxChange) MarshalJSON() ([]byte, error) {
	type Alias SetOrderTotalTaxChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetOrderTotalTaxChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Parcel Items](ctp:api:type:OrderSetParcelItemsAction) on Orders.
*	- [Set Parcel Items](ctp:api:type:StagedOrderSetParcelItemsAction) on Staged Orders.
*
 */
type SetParcelItemsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []DeliveryItem `json:"previousValue"`
	// Value after the change.
	NextValue []DeliveryItem `json:"nextValue"`
	// Information about the updated Parcel.
	Parcel ParcelChangeValue `json:"parcel"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetParcelItemsChange) MarshalJSON() ([]byte, error) {
	type Alias SetParcelItemsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetParcelItemsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [SetParcelMeasurements](ctp:api:type:OrderSetParcelMeasurementsAction) on Orders.
*	- [SetParcelMeasurements](ctp:api:type:StagedOrderSetParcelMeasurementsAction) on Staged Orders.
*
 */
type SetParcelMeasurementsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ParcelMeasurements `json:"previousValue"`
	// Value after the change.
	NextValue ParcelMeasurements `json:"nextValue"`
	// Information about the updated Parcel.
	Parcel ParcelChangeValue `json:"parcel"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetParcelMeasurementsChange) MarshalJSON() ([]byte, error) {
	type Alias SetParcelMeasurementsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetParcelMeasurementsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Parcel Tracking Data](ctp:api:type:OrderSetParcelTrackingDataAction) on Orders.
*	- [Set Parcel Tracking Data](ctp:api:type:StagedOrderSetParcelTrackingDataAction) on Staged Orders.
*
 */
type SetParcelTrackingDataChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TrackingData `json:"previousValue"`
	// Value after the change.
	NextValue TrackingData `json:"nextValue"`
	// Information about the updated Parcel.
	Parcel ParcelChangeValue `json:"parcel"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetParcelTrackingDataChange) MarshalJSON() ([]byte, error) {
	type Alias SetParcelTrackingDataChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetParcelTrackingDataChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Permissions](ctp:api:type:AssociateRoleSetPermissionsAction), [Add Permission](ctp:api:type:AssociateRoleAddPermissionAction), and [Remove Permission](ctp:api:type:AssociateRoleRemovePermissionAction) update actions.
*
 */
type SetPermissionsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Permission `json:"previousValue"`
	// Value after the change.
	NextValue []Permission `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetPermissionsChange) MarshalJSON() ([]byte, error) {
	type Alias SetPermissionsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetPermissionsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Prices](ctp:api:type:ProductSetPricesAction) update action.
 */
type SetPricesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Price `json:"previousValue"`
	// Value after the change.
	NextValue []Price `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
	// `sku` or `key` of the [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetPricesChange) MarshalJSON() ([]byte, error) {
	type Alias SetPricesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetPricesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered automatically by the [Add Product](ctp:api:type:ProductSelectionAddProductAction) or [Remove Product](ctp:api:type:ProductSelectionRemoveProductAction) update action.
 */
type SetProductCountChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue int `json:"previousValue"`
	// Value after the change.
	NextValue int `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetProductCountChange) MarshalJSON() ([]byte, error) {
	type Alias SetProductCountChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetProductCountChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Price CustomField](ctp:api:type:ProductSetProductPriceCustomFieldAction) update action.
 */
type SetProductPriceCustomFieldChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CustomFields `json:"previousValue"`
	// Value after the change.
	NextValue CustomFields `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetProductPriceCustomFieldChange) MarshalJSON() ([]byte, error) {
	type Alias SetProductPriceCustomFieldChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetProductPriceCustomFieldChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Price Custom Type](ctp:api:type:ProductSetProductPriceCustomTypeAction) update action.
 */
type SetProductPriceCustomTypeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CustomFields `json:"previousValue"`
	// Value after the change.
	NextValue CustomFields `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetProductPriceCustomTypeChange) MarshalJSON() ([]byte, error) {
	type Alias SetProductPriceCustomTypeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetProductPriceCustomTypeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Product Selections](ctp:api:type:StoreSetProductSelectionsAction) update action.
 */
type SetProductSelectionsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []ProductSelectionSetting `json:"previousValue"`
	// Value after the change.
	NextValue []ProductSelectionSetting `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetProductSelectionsChange) MarshalJSON() ([]byte, error) {
	type Alias SetProductSelectionsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetProductSelectionsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set ProductVariant Key](ctp:api:type:ProductSetProductVariantKeyAction) update action.
 */
type SetProductVariantKeyChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetProductVariantKeyChange) MarshalJSON() ([]byte, error) {
	type Alias SetProductVariantKeyChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetProductVariantKeyChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Update CustomObject](ctp:api:endpoint:/{projectKey}/custom-objects:POST) request when an existing property is updated.
 */
type SetPropertyChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
	// Path to the property that was updated.
	Path string `json:"path"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetPropertyChange) MarshalJSON() ([]byte, error) {
	type Alias SetPropertyChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetPropertyChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Purchase Order Number](ctp:api:type:OrderSetPurchaseOrderNumberAction) on Orders.
*	- [Set Purchase Order Number](ctp:api:type:StagedOrderSetPurchaseOrderNumberAction) on Staged Orders.
*
 */
type SetPurchaseOrderNumberChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetPurchaseOrderNumberChange) MarshalJSON() ([]byte, error) {
	type Alias SetPurchaseOrderNumberChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetPurchaseOrderNumberChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Rating](ctp:api:type:ReviewSetRatingAction) update action.
 */
type SetRatingChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue int `json:"previousValue"`
	// Value after the change.
	NextValue int `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetRatingChange) MarshalJSON() ([]byte, error) {
	type Alias SetRatingChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetRatingChange", Alias: (*Alias)(&obj)})
}

type SetReservationsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Reservation `json:"previousValue"`
	// Value after the change.
	NextValue []Reservation `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetReservationsChange) MarshalJSON() ([]byte, error) {
	type Alias SetReservationsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetReservationsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set RestockableInDays](ctp:api:type:InventoryEntrySetRestockableInDaysAction) update action.
 */
type SetRestockableInDaysChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue int `json:"previousValue"`
	// Value after the change.
	NextValue int `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetRestockableInDaysChange) MarshalJSON() ([]byte, error) {
	type Alias SetRestockableInDaysChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetRestockableInDaysChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set PaymentShipmentState](ctp:api:type:OrderSetReturnPaymentStateAction) on Orders.
*	- [Set PaymentShipmentState](ctp:api:type:StagedOrderSetReturnPaymentStateAction) on Staged Orders.
*
 */
type SetReturnPaymentStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ReturnPaymentState `json:"previousValue"`
	// Value after the change.
	NextValue ReturnPaymentState `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetReturnPaymentStateChange) MarshalJSON() ([]byte, error) {
	type Alias SetReturnPaymentStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetReturnPaymentStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set ReturnShipmentState](ctp:api:type:OrderSetReturnShipmentStateAction) on Orders.
*	- [Set ReturnShipmentState](ctp:api:type:StagedOrderSetReturnShipmentStateAction) on Staged Orders.
*
 */
type SetReturnShipmentStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ReturnShipmentState `json:"previousValue"`
	// Value after the change.
	NextValue ReturnShipmentState `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetReturnShipmentStateChange) MarshalJSON() ([]byte, error) {
	type Alias SetReturnShipmentStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetReturnShipmentStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Salutation](ctp:api:type:CustomerSetSalutationAction) update action.
 */
type SetSalutationChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetSalutationChange) MarshalJSON() ([]byte, error) {
	type Alias SetSalutationChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetSalutationChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set SearchKeywords](ctp:api:type:ProductSetSearchKeywordsAction) update action.
 */
type SetSearchKeywordsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue SearchKeywords `json:"previousValue"`
	// Value after the change.
	NextValue SearchKeywords `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetSearchKeywordsChange) MarshalJSON() ([]byte, error) {
	type Alias SetSearchKeywordsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetSearchKeywordsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Seller Comment](ctp:api:type:StagedQuoteSetSellerCommentAction) update action.
 */
type SetSellerCommentChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetSellerCommentChange) MarshalJSON() ([]byte, error) {
	type Alias SetSellerCommentChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetSellerCommentChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Shipping Address](ctp:api:type:OrderSetShippingAddressAction) on Orders.
*	- [Set Shipping Address](ctp:api:type:StagedOrderSetShippingAddressAction) on Staged Orders.
*
 */
type SetShippingAddressChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Address `json:"previousValue"`
	// Value after the change.
	NextValue Address `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetShippingAddressChange) MarshalJSON() ([]byte, error) {
	type Alias SetShippingAddressChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetShippingAddressChange", Alias: (*Alias)(&obj)})
}

type SetShippingInfoPriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Money `json:"previousValue"`
	// Value after the change.
	NextValue Money `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetShippingInfoPriceChange) MarshalJSON() ([]byte, error) {
	type Alias SetShippingInfoPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetShippingInfoPriceChange", Alias: (*Alias)(&obj)})
}

type SetShippingInfoTaxedPriceChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TaxedPrice `json:"previousValue"`
	// Value after the change.
	NextValue TaxedPrice `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetShippingInfoTaxedPriceChange) MarshalJSON() ([]byte, error) {
	type Alias SetShippingInfoTaxedPriceChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetShippingInfoTaxedPriceChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set ShippingMethod](ctp:api:type:StagedOrderSetShippingMethodAction) update action.
 */
type SetShippingMethodChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ShippingMethodChangeValue `json:"previousValue"`
	// Value after the change.
	NextValue ShippingMethodChangeValue `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetShippingMethodChange) MarshalJSON() ([]byte, error) {
	type Alias SetShippingMethodChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetShippingMethodChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set ShippingMethod TaxAmount](ctp:api:type:StagedOrderSetShippingMethodTaxAmountAction) update action.
 */
type SetShippingMethodTaxAmountChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ShippingMethodTaxAmountChangeValue `json:"previousValue"`
	// Value after the change.
	NextValue ShippingMethodTaxAmountChangeValue `json:"nextValue"`
	// `"ExternalAmount"`
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetShippingMethodTaxAmountChange) MarshalJSON() ([]byte, error) {
	type Alias SetShippingMethodTaxAmountChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetShippingMethodTaxAmountChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set ShippingMethod TaxRate](ctp:api:type:StagedOrderSetShippingMethodTaxRateAction) update action.
 */
type SetShippingMethodTaxRateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue TaxRate `json:"previousValue"`
	// Value after the change.
	NextValue TaxRate `json:"nextValue"`
	// `"External"`
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetShippingMethodTaxRateChange) MarshalJSON() ([]byte, error) {
	type Alias SetShippingMethodTaxRateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetShippingMethodTaxRateChange", Alias: (*Alias)(&obj)})
}

type SetShippingRateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Money `json:"previousValue"`
	// Value after the change.
	NextValue Money `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetShippingRateChange) MarshalJSON() ([]byte, error) {
	type Alias SetShippingRateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetShippingRateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Shipping Rate Input](ctp:api:type:StagedOrderSetShippingRateInputAction) update action.
 */
type SetShippingRateInputChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetShippingRateInputChange) MarshalJSON() ([]byte, error) {
	type Alias SetShippingRateInputChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetShippingRateInputChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set ShoppingListLineItem Custom Field](ctp:api:type:ShoppingListSetLineItemCustomFieldAction) update action.
 */
type SetShoppingListLineItemCustomFieldChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// `id` of the referenced [Type](ctp:api:type:Type).
	CustomTypeId string `json:"customTypeId"`
	// Holds information about the updated Shopping List Line Item.
	LineItem ShoppingListLineItemValue `json:"lineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetShoppingListLineItemCustomFieldChange) MarshalJSON() ([]byte, error) {
	type Alias SetShoppingListLineItemCustomFieldChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetShoppingListLineItemCustomFieldChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set ShoppingListLineItem Custom Type](ctp:api:type:ShoppingListSetLineItemCustomTypeAction) update action.
 */
type SetShoppingListLineItemCustomTypeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CustomFields `json:"previousValue"`
	// Value after the change.
	NextValue CustomFields `json:"nextValue"`
	// Holds information about the updated Shopping List Line Item.
	LineItem ShoppingListLineItemValue `json:"lineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetShoppingListLineItemCustomTypeChange) MarshalJSON() ([]byte, error) {
	type Alias SetShoppingListLineItemCustomTypeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetShoppingListLineItemCustomTypeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set SKU](ctp:api:type:ProductSetSkuAction) update action.
 */
type SetSkuChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetSkuChange) MarshalJSON() ([]byte, error) {
	type Alias SetSkuChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetSkuChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Slug](ctp:api:type:ShoppingListSetSlugAction) update action.
 */
type SetSlugChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetSlugChange) MarshalJSON() ([]byte, error) {
	type Alias SetSlugChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetSlugChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set State roles](ctp:api:type:StateSetRolesAction) update action.
 */
type SetStateRolesChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []StateRoleEnum `json:"previousValue"`
	// Value after the change.
	NextValue []StateRoleEnum `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetStateRolesChange) MarshalJSON() ([]byte, error) {
	type Alias SetStateRolesChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetStateRolesChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set StatusInterfaceCode](ctp:api:type:PaymentSetStatusInterfaceCodeAction) update action.
 */
type SetStatusInterfaceCodeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetStatusInterfaceCodeChange) MarshalJSON() ([]byte, error) {
	type Alias SetStatusInterfaceCodeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetStatusInterfaceCodeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set StatusInterfaceText](ctp:api:type:PaymentSetStatusInterfaceTextAction) update action.
 */
type SetStatusInterfaceTextChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetStatusInterfaceTextChange) MarshalJSON() ([]byte, error) {
	type Alias SetStatusInterfaceTextChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetStatusInterfaceTextChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Store](ctp:api:type:OrderSetStoreAction) on Orders.
*	- [Set Store](ctp:api:type:ShoppingListSetStoreAction) on Shopping Lists.
*
 */
type SetStoreChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Reference `json:"previousValue"`
	// Value after the change.
	NextValue Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetStoreChange) MarshalJSON() ([]byte, error) {
	type Alias SetStoreChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetStoreChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Store Mode](ctp:api:type:BusinessUnitSetStoreModeAction) update action.
 */
type SetStoreModeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue BusinessUnitStoreMode `json:"previousValue"`
	// Value after the change.
	NextValue BusinessUnitStoreMode `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetStoreModeChange) MarshalJSON() ([]byte, error) {
	type Alias SetStoreModeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetStoreModeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Stores](ctp:api:type:CustomerSetStoresAction) update action.
 */
type SetStoresChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Reference `json:"previousValue"`
	// Value after the change.
	NextValue []Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetStoresChange) MarshalJSON() ([]byte, error) {
	type Alias SetStoresChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetStoresChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set SupplyChannel](ctp:api:type:InventoryEntrySetSupplyChannelAction) update action.
 */
type SetSupplyChannelChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Reference `json:"previousValue"`
	// Value after the change.
	NextValue Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetSupplyChannelChange) MarshalJSON() ([]byte, error) {
	type Alias SetSupplyChannelChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetSupplyChannelChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Supply Channels](ctp:api:type:StoreSetSupplyChannelsAction) update action.
 */
type SetSupplyChannelsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Reference `json:"previousValue"`
	// Value after the change.
	NextValue []Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetSupplyChannelsChange) MarshalJSON() ([]byte, error) {
	type Alias SetSupplyChannelsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetSupplyChannelsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Target](ctp:api:type:ReviewSetTargetAction) update action.
 */
type SetTargetChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Reference `json:"previousValue"`
	// Value after the change.
	NextValue Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetTargetChange) MarshalJSON() ([]byte, error) {
	type Alias SetTargetChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetTargetChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set TaxCategory](ctp:api:type:ProductSetTaxCategoryAction) update action.
 */
type SetTaxCategoryChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Reference `json:"previousValue"`
	// Value after the change.
	NextValue Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetTaxCategoryChange) MarshalJSON() ([]byte, error) {
	type Alias SetTaxCategoryChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetTaxCategoryChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Text](ctp:api:type:ReviewSetTextAction) update action.
 */
type SetTextChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetTextChange) MarshalJSON() ([]byte, error) {
	type Alias SetTextChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetTextChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set TextLineItem CustomField](ctp:api:type:ShoppingListSetTextLineItemCustomFieldAction) update action.
 */
type SetTextLineItemCustomFieldChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// `id` of the referenced [Type](ctp:api:type:Type).
	CustomTypeId string `json:"customTypeId"`
	// Holds information about the updated Text Line Item.
	TextLineItem TextLineItemValue `json:"textLineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetTextLineItemCustomFieldChange) MarshalJSON() ([]byte, error) {
	type Alias SetTextLineItemCustomFieldChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetTextLineItemCustomFieldChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set TextLineItem Custom Type](ctp:api:type:ShoppingListSetTextLineItemCustomTypeAction) update action.
 */
type SetTextLineItemCustomTypeChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue CustomFields `json:"previousValue"`
	// Value after the change.
	NextValue CustomFields `json:"nextValue"`
	// Holds information about the updated Text Line Item.
	TextLineItem TextLineItemValue `json:"textLineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetTextLineItemCustomTypeChange) MarshalJSON() ([]byte, error) {
	type Alias SetTextLineItemCustomTypeChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetTextLineItemCustomTypeChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set TextLineItem Description](ctp:api:type:ShoppingListSetTextLineItemDescriptionAction) update action.
 */
type SetTextLineItemDescriptionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue LocalizedString `json:"previousValue"`
	// Value after the change.
	NextValue LocalizedString `json:"nextValue"`
	// Holds information about the updated Text Line Item.
	TextLineItem TextLineItemValue `json:"textLineItem"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetTextLineItemDescriptionChange) MarshalJSON() ([]byte, error) {
	type Alias SetTextLineItemDescriptionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetTextLineItemDescriptionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Title](ctp:api:type:CustomerSetTitleAction) on Customers.
*	- [Set Title](ctp:api:type:ReviewSetTitleAction) on Reviews.
*
 */
type SetTitleChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetTitleChange) MarshalJSON() ([]byte, error) {
	type Alias SetTitleChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetTitleChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Transitions](ctp:api:type:StateSetTransitionsAction) update action.
 */
type SetTransitionsChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []Reference `json:"previousValue"`
	// Value after the change.
	NextValue []Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetTransitionsChange) MarshalJSON() ([]byte, error) {
	type Alias SetTransitionsChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetTransitionsChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Valid From and Until](ctp:api:type:CartDiscountSetValidFromAndUntilAction) on Cart Discounts.
*	- [Set Valid From and Until](ctp:api:type:DiscountCodeSetValidFromAndUntilAction) on Discount Codes.
*	- [Set Valid From and Until](ctp:api:type:ProductDiscountSetValidFromAndUntilAction) on Product Discounts.
*
 */
type SetValidFromAndUntilChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ValidFromAndUntilValue `json:"previousValue"`
	// Value after the change.
	NextValue ValidFromAndUntilValue `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetValidFromAndUntilChange) MarshalJSON() ([]byte, error) {
	type Alias SetValidFromAndUntilChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetValidFromAndUntilChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Valid From](ctp:api:type:CartDiscountSetValidFromAction) on Cart Discounts.
*	- [Set Valid From](ctp:api:type:DiscountCodeSetValidFromAction) on Discount Codes.
*	- [Set Valid From](ctp:api:type:ProductDiscountSetValidFromAction) on Product Discounts.
*
 */
type SetValidFromChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetValidFromChange) MarshalJSON() ([]byte, error) {
	type Alias SetValidFromChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetValidFromChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Valid To](ctp:api:type:StagedQuoteSetValidToAction) update action.
 */
type SetValidToChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetValidToChange) MarshalJSON() ([]byte, error) {
	type Alias SetValidToChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetValidToChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Set Valid Until](ctp:api:type:CartDiscountSetValidUntilAction) on Cart Discounts.
*	- [Set Valid Until](ctp:api:type:DiscountCodeSetValidUntilAction) on Discount Codes.
*	- [Set Valid Until](ctp:api:type:ProductDiscountSetValidUntilAction) on Product Discounts.
*
 */
type SetValidUntilChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetValidUntilChange) MarshalJSON() ([]byte, error) {
	type Alias SetValidUntilChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetValidUntilChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Update CustomObject](ctp:api:endpoint:/{projectKey}/custom-objects:POST) request when a value of a property is updated.
 */
type SetValueChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetValueChange) MarshalJSON() ([]byte, error) {
	type Alias SetValueChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetValueChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered automatically when an [InventoryEntry](ctp:api:type:InventoryEntry) associated with a Product changes.
 */
type SetVariantAvailabilityChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ProductVariantAvailability `json:"previousValue"`
	// Value after the change.
	NextValue ProductVariantAvailability `json:"nextValue"`
	// - `staged`, if the staged [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	// - `current`, if the current [ProductCatalogData](ctp:api:type:ProductCatalogData) was updated.
	CatalogData string `json:"catalogData"`
	// `sku` or `key` of the [ProductVariant](ctp:api:type:ProductVariant).
	Variant string `json:"variant"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetVariantAvailabilityChange) MarshalJSON() ([]byte, error) {
	type Alias SetVariantAvailabilityChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetVariantAvailabilityChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Variant Selection](ctp:api:type:ProductSelectionSetVariantSelectionAction) update action.
 */
type SetVariantSelectionChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue ProductVariantSelection `json:"previousValue"`
	// Value after the change.
	NextValue ProductVariantSelection `json:"nextValue"`
	// Reference to the updated [Product](ctp:api:type:Product).
	Product Reference `json:"product"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetVariantSelectionChange) MarshalJSON() ([]byte, error) {
	type Alias SetVariantSelectionChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetVariantSelectionChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Set Vat ID](ctp:api:type:CustomerSetVatIdAction) update action.
 */
type SetVatIdChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue string `json:"previousValue"`
	// Value after the change.
	NextValue string `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetVatIdChange) MarshalJSON() ([]byte, error) {
	type Alias SetVatIdChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SetVatIdChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change the state of CustomLineItem according to allowed transitions](ctp:api:type:OrderTransitionCustomLineItemStateAction) on Orders.
*	- [Change the state of CustomLineItem according to allowed transitions](ctp:api:type:StagedOrderTransitionCustomLineItemStateAction) on Staged Orders.
*
 */
type TransitionCustomLineItemStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []ItemState `json:"previousValue"`
	// Value after the change.
	NextValue []ItemState `json:"nextValue"`
	// `id` of the updated [CustomLineItem](ctp:api:type:CustomLineItem).
	LineItemId string `json:"lineItemId"`
	// `id` of the [State](ctp:api:type:State) involved in the transition.
	StateId string `json:"stateId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TransitionCustomLineItemStateChange) MarshalJSON() ([]byte, error) {
	type Alias TransitionCustomLineItemStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "TransitionCustomLineItemStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Change the state of LineItem according to allowed transitions](ctp:api:type:OrderTransitionLineItemStateAction) on Orders.
*	- [Change the state of LineItem according to allowed transitions](ctp:api:type:OrderTransitionLineItemStateAction) on Staged Orders.
*
 */
type TransitionLineItemStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue []ItemState `json:"previousValue"`
	// Value after the change.
	NextValue []ItemState `json:"nextValue"`
	// `id` of the updated [LineItem](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
	// `id` of the [State](ctp:api:type:State) involved in the transition.
	StateId string `json:"stateId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TransitionLineItemStateChange) MarshalJSON() ([]byte, error) {
	type Alias TransitionLineItemStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "TransitionLineItemStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Transition State](ctp:api:type:OrderTransitionStateAction) on Orders.
*	- [Transition State](ctp:api:type:StagedOrderTransitionStateAction) on Staged Orders.
*	- [Transition State](ctp:api:type:PaymentTransitionStateAction) on Payments.
*	- [Transition State](ctp:api:type:ProductTransitionStateAction) on Products.
*	- [Transition State](ctp:api:type:QuoteTransitionStateAction) on Quotes.
*	- [Transition State](ctp:api:type:StagedQuoteTransitionStateAction) on Staged Quotes.
*	- [Transition State](ctp:api:type:QuoteRequestTransitionStateAction) on Quote Requests.
*	- [Transition State](ctp:api:type:ReviewTransitionStateAction) on Reviews.
*	- [Transition State](ctp:api:type:StateSetTransitionsAction) on States.
*
 */
type TransitionStateChange struct {
	Change string `json:"change"`
	// Value before the change.
	PreviousValue Reference `json:"previousValue"`
	// Value after the change.
	NextValue Reference `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TransitionStateChange) MarshalJSON() ([]byte, error) {
	type Alias TransitionStateChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "TransitionStateChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered when the format of changes on an entity is not identified by Audit Log.
 */
type UnknownChange struct {
	// Identifier for the type of modification.
	Change string `json:"change"`
	// Value before the change.
	PreviousValue interface{} `json:"previousValue"`
	// Value after the change.
	NextValue interface{} `json:"nextValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj UnknownChange) MarshalJSON() ([]byte, error) {
	type Alias UnknownChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "UnknownChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the [Unpublish](ctp:api:type:ProductUnpublishAction) update action.
 */
type UnpublishChange struct {
	Change string `json:"change"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj UnpublishChange) MarshalJSON() ([]byte, error) {
	type Alias UnpublishChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "UnpublishChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by the following update actions:
*
*	- [Update SyncInfo](ctp:api:type:OrderUpdateSyncInfoAction) on Orders.
*	- [Update SyncInfo](ctp:api:type:StagedOrderUpdateSyncInfoAction) on Staged Orders.
*
 */
type UpdateSyncInfoChange struct {
	Change string `json:"change"`
	// Value after the change.
	NextValue SyncInfo `json:"nextValue"`
	// `id` of the updated [Channel](ctp:api:type:Channel).
	ChannelId string `json:"channelId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj UpdateSyncInfoChange) MarshalJSON() ([]byte, error) {
	type Alias UpdateSyncInfoChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "UpdateSyncInfoChange", Alias: (*Alias)(&obj)})
}

/**
*	Change triggered by a Customer email verification.
 */
type VerifyEmailChange struct {
	Change string `json:"change"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VerifyEmailChange) MarshalJSON() ([]byte, error) {
	type Alias VerifyEmailChange
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "VerifyEmailChange", Alias: (*Alias)(&obj)})
}
