package history

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

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
	case "ChangeTargetChange":
		obj := ChangeTargetChange{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.NextValue != nil {
			var err error
			obj.NextValue, err = mapDiscriminatorChangeTargetChangeValue(obj.NextValue)
			if err != nil {
				return nil, err
			}
		}
		if obj.PreviousValue != nil {
			var err error
			obj.PreviousValue, err = mapDiscriminatorChangeTargetChangeValue(obj.PreviousValue)
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
		if obj.NextValue != nil {
			var err error
			obj.NextValue, err = mapDiscriminatorChangeValueChangeValue(obj.NextValue)
			if err != nil {
				return nil, err
			}
		}
		if obj.PreviousValue != nil {
			var err error
			obj.PreviousValue, err = mapDiscriminatorChangeValueChangeValue(obj.PreviousValue)
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
	case "SetAddressChange":
		obj := SetAddressChange{}
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

type AddAddressChange struct {
	// Update action for `setAddress` action.
	Change        string  `json:"change"`
	NextValue     Address `json:"nextValue"`
	PreviousValue Address `json:"previousValue"`
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

type AddAssetChange struct {
	// Update action for `addAsset`
	Change        string `json:"change"`
	NextValue     Asset  `json:"nextValue"`
	PreviousValue Asset  `json:"previousValue"`
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

type AddAttributeDefinitionChange struct {
	// Update action for `addAttributeDefinition` on product types
	Change    string              `json:"change"`
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

type AddBillingAddressIdChange struct {
	// Update action for `addBillingAddressId` action on customers.
	Change        string   `json:"change"`
	NextValue     []string `json:"nextValue"`
	PreviousValue []string `json:"previousValue"`
	Address       Address  `json:"address"`
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

type AddChannelRolesChange struct {
	Change        string        `json:"change"`
	PreviousValue []ChannelRole `json:"previousValue"`
	NextValue     []ChannelRole `json:"nextValue"`
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

type AddCustomLineItemChange struct {
	// Update action for adding and removing custom line items
	Change        string         `json:"change"`
	NextValue     CustomLineItem `json:"nextValue"`
	PreviousValue CustomLineItem `json:"previousValue"`
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

type AddDeliveryChange struct {
	// Update action for `addDelivery`
	Change        string              `json:"change"`
	NextValue     DeliveryChangeValue `json:"nextValue"`
	PreviousValue DeliveryChangeValue `json:"previousValue"`
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

type AddDiscountCodeChange struct {
	// Update action for `addDiscountCode`
	Change    string           `json:"change"`
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

type AddEnumValueChange struct {
	// Update action for `addEnumValue` on types
	Change string `json:"change"`
	// The name of the field/attribute definition updated.
	FieldName string    `json:"fieldName"`
	NextValue EnumValue `json:"nextValue"`
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

type AddExternalImageChange struct {
	// Update actions for adding an external image
	Change        string  `json:"change"`
	CatalogData   string  `json:"catalogData"`
	PreviousValue []Image `json:"previousValue"`
	NextValue     []Image `json:"nextValue"`
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

type AddFieldDefinitionChange struct {
	// Update action for `addFieldDefinition` on payments
	Change    string          `json:"change"`
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

type AddInterfaceInteractionChange struct {
	// Update action for `addInterfaceInteraction` on payments
	Change string `json:"change"`
	// Only available if `expand` is set to true
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

type AddItemShippingAddressesChange struct {
	// Update action for `addItemShippingAddress`
	Change        string  `json:"change"`
	NextValue     Address `json:"nextValue"`
	PreviousValue Address `json:"previousValue"`
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

type AddLocalizedEnumValueChange struct {
	// Update action for `addLocalizedEnumValue` on types
	Change string `json:"change"`
	// The name of the field definition updated.
	FieldName string `json:"fieldName"`
	// The name of the attribute updated.
	AttributeName string             `json:"attributeName"`
	NextValue     LocalizedEnumValue `json:"nextValue"`
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

type AddLocationChange struct {
	// Update action for `addLocation` on zones
	Change string `json:"change"`
	// Shape of the value for `addLocation` and `removeLocation` actions
	PreviousValue Location `json:"previousValue"`
	// Shape of the value for `addLocation` and `removeLocation` actions
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

type AddOrderLineItemChange struct {
	Change        string   `json:"change"`
	PreviousValue LineItem `json:"previousValue"`
	NextValue     LineItem `json:"nextValue"`
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

type AddParcelToDeliveryChange struct {
	// Update action for `addParcelToDelivery`
	Change     string `json:"change"`
	DeliveryId string `json:"deliveryId"`
	NextValue  Parcel `json:"nextValue"`
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

type AddPaymentChange struct {
	// Update action for `addPayment` & `removePayment`
	Change        string      `json:"change"`
	NextValue     PaymentInfo `json:"nextValue"`
	PreviousValue PaymentInfo `json:"previousValue"`
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

type AddPlainEnumValueChange struct {
	// Update action for `addPlainEnumValue` on product types
	Change string `json:"change"`
	// The name of the attribute updated.
	AttributeName string    `json:"attributeName"`
	NextValue     EnumValue `json:"nextValue"`
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

type AddPriceChange struct {
	// Update action for adding prices
	Change      string `json:"change"`
	CatalogData string `json:"catalogData"`
	PriceId     string `json:"priceId"`
	NextValue   Price  `json:"nextValue"`
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

type AddProductChange struct {
	// Update action for when a product is assigned to a product selection
	Change    string    `json:"change"`
	NextValue Reference `json:"nextValue"`
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

type AddPropertyChange struct {
	// Update action for `addProperty` on custom objects
	Change string `json:"change"`
	// Value path to the property that was added
	Path      string      `json:"path"`
	NextValue interface{} `json:"nextValue"`
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

type AddReturnInfoChange struct {
	// Update action for `addReturnInfo`
	Change    string     `json:"change"`
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

type AddShippingAddressIdChange struct {
	// Update action for `addShippingAddressId` action on customers.
	Change        string   `json:"change"`
	NextValue     []string `json:"nextValue"`
	PreviousValue []string `json:"previousValue"`
	Address       Address  `json:"address"`
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

type AddShoppingListLineItemChange struct {
	Change        string   `json:"change"`
	PreviousValue LineItem `json:"previousValue"`
	NextValue     LineItem `json:"nextValue"`
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

type AddStateRolesChange struct {
	Change        string      `json:"change"`
	PreviousValue []StateRole `json:"previousValue"`
	NextValue     []StateRole `json:"nextValue"`
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

type AddTaxRateChange struct {
	// Update action for `addTaxRate` on tax categories
	Change string `json:"change"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
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

type AddTextLineItemChange struct {
	Change    string       `json:"change"`
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

type AddToCategoryChange struct {
	// Update action for `addToCategory`
	Change        string      `json:"change"`
	Category      Reference   `json:"category"`
	PreviousValue []Reference `json:"previousValue"`
	NextValue     []Reference `json:"nextValue"`
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

type AddTransactionChange struct {
	// Update action for `addTransaction` on payments
	Change    string      `json:"change"`
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

type AddVariantChange struct {
	// Update action for `addVariant`
	Change        string  `json:"change"`
	CatalogData   string  `json:"catalogData"`
	PreviousValue Variant `json:"previousValue"`
	NextValue     Variant `json:"nextValue"`
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

type ChangeAddressChange struct {
	// Update action `changeAddress` action.
	Change        string  `json:"change"`
	NextValue     Address `json:"nextValue"`
	PreviousValue Address `json:"previousValue"`
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

type ChangeAmountAuthorizedChange struct {
	// Internal Update action for `changeAmountAuthorized`
	Change        string `json:"change"`
	PreviousValue Money  `json:"previousValue"`
	NextValue     Money  `json:"nextValue"`
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

type ChangeAmountPlannedChange struct {
	Change        string `json:"change"`
	PreviousValue Money  `json:"previousValue"`
	NextValue     Money  `json:"nextValue"`
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

type ChangeAssetNameChange struct {
	// Update action for `changeAssetName`
	Change        string           `json:"change"`
	Asset         AssetChangeValue `json:"asset"`
	NextValue     LocalizedString  `json:"nextValue"`
	PreviousValue LocalizedString  `json:"previousValue"`
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

type ChangeAssetOrderChange struct {
	Change        string            `json:"change"`
	PreviousValue []LocalizedString `json:"previousValue"`
	NextValue     []LocalizedString `json:"nextValue"`
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

type ChangeAttributeConstraintChange struct {
	Change string `json:"change"`
	// name of the updated attribute
	AttributeName string                  `json:"attributeName"`
	PreviousValue AttributeConstraintEnum `json:"previousValue"`
	NextValue     AttributeConstraintEnum `json:"nextValue"`
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

type ChangeAttributeOrderByNameChange struct {
	// Update action for `changeAttributeOrderByName` on product types
	Change        string   `json:"change"`
	PreviousValue []string `json:"previousValue"`
	NextValue     []string `json:"nextValue"`
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

type ChangeCartDiscountsChange struct {
	// Shape of the action for `changeCartDiscounts`
	Change        string      `json:"change"`
	PreviousValue []Reference `json:"previousValue"`
	NextValue     []Reference `json:"nextValue"`
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

type ChangeCartPredicateChange struct {
	// Shape of the action for `changeCartPredicate`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type ChangeCustomLineItemQuantityChange struct {
	// Update action for `changeCustomLineItemQuantity`
	Change           string          `json:"change"`
	CustomLineItem   LocalizedString `json:"customLineItem"`
	CustomLineItemId string          `json:"customLineItemId"`
	NextValue        int             `json:"nextValue"`
	PreviousValue    int             `json:"previousValue"`
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

type ChangeDescriptionChange struct {
	// Shape of the action for `changeDescription`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type ChangeEmailChange struct {
	// Shape of the action for `changeEmail`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type ChangeEnumValueLabelChange struct {
	// Update action for `changeEnumValueLabel` on types
	Change string `json:"change"`
	// The name of the field definition updated.
	FieldName string `json:"fieldName"`
	// Key of the values that was updated
	ValueKey      string `json:"valueKey"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type ChangeEnumValueOrderChange struct {
	// Update action for `changeEnumValueOrder` on types
	Change string `json:"change"`
	// The name of the field/attribute definition updated.
	FieldName     string      `json:"fieldName"`
	NextValue     []EnumValue `json:"nextValue"`
	PreviousValue []EnumValue `json:"previousValue"`
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

type ChangeFieldDefinitionOrderChange struct {
	// Update action for `changeFieldDefinitionOrder` on types
	Change        string                      `json:"change"`
	PreviousValue []FieldDefinitionOrderValue `json:"previousValue"`
	NextValue     []FieldDefinitionOrderValue `json:"nextValue"`
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

type ChangeGroupsChange struct {
	// Update action for `changeGroups` on stores
	Change        string   `json:"change"`
	PreviousValue []string `json:"previousValue"`
	NextValue     []string `json:"nextValue"`
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

type ChangeInitialChange struct {
	// Shape of the action for `changeInitial`
	Change        string `json:"change"`
	PreviousValue bool   `json:"previousValue"`
	NextValue     bool   `json:"nextValue"`
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

type ChangeInputHintChange struct {
	// Update action for `changeInputHint` on product types and types
	Change string `json:"change"`
	// The name of the field definition updated.
	FieldName string `json:"fieldName"`
	// The name of the attribute updated.
	AttributeName string        `json:"attributeName"`
	NextValue     TextInputHint `json:"nextValue"`
	PreviousValue TextInputHint `json:"previousValue"`
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

type ChangeIsActiveChange struct {
	// Shape of the action for `changeIsActive`
	Change        string `json:"change"`
	PreviousValue bool   `json:"previousValue"`
	NextValue     bool   `json:"nextValue"`
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

type ChangeIsSearchableChange struct {
	// Update action for `changeIsSearchable` on product types
	Change string `json:"change"`
	// The name of the updated attribute.
	AttributeName string `json:"attributeName"`
	NextValue     bool   `json:"nextValue"`
	PreviousValue bool   `json:"previousValue"`
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

type ChangeKeyChange struct {
	// Shape of the action for `changeKey`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type ChangeLabelChange struct {
	// Update action for `changeLabel` on product types and types
	Change string `json:"change"`
	// The name of the field definition to update (types).
	FieldName string `json:"fieldName"`
	// The name of the attribute definition to update (product-type).
	AttributeName string          `json:"attributeName"`
	NextValue     LocalizedString `json:"nextValue"`
	PreviousValue LocalizedString `json:"previousValue"`
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

type ChangeLineItemQuantityChange struct {
	// Update action for `changeLineItemQuantity`
	Change     string          `json:"change"`
	LineItem   LocalizedString `json:"lineItem"`
	LineItemId string          `json:"lineItemId"`
	// The amount of a LineItem in the cart. Must be a positive integer.
	NextValue int `json:"nextValue"`
	// The amount of a LineItem in the cart. Must be a positive integer.
	PreviousValue int `json:"previousValue"`
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

type ChangeLocalizedDescriptionChange struct {
	// Shape of the action for `changeDescription`
	Change        string          `json:"change"`
	PreviousValue LocalizedString `json:"previousValue"`
	NextValue     LocalizedString `json:"nextValue"`
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

type ChangeLocalizedEnumValueLabelChange struct {
	// Update action for `changeLocalizedEnumValueLabel` on types
	Change string `json:"change"`
	// The name of the field definition updated.
	FieldName string `json:"fieldName"`
	// The name of the attribute updated.
	AttributeName string `json:"attributeName"`
	// Key of the values that was updated
	ValueKey      string          `json:"valueKey"`
	PreviousValue LocalizedString `json:"previousValue"`
	NextValue     LocalizedString `json:"nextValue"`
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

type ChangeLocalizedEnumValueOrderChange struct {
	// Update action for `changeLocalizedEnumValueOrder` on types and product types
	Change string `json:"change"`
	// The name of the field definition updated.
	FieldName string `json:"fieldName"`
	// The name of the attribute updated.
	AttributeName string               `json:"attributeName"`
	NextValue     []LocalizedEnumValue `json:"nextValue"`
	PreviousValue []LocalizedEnumValue `json:"previousValue"`
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

type ChangeLocalizedNameChange struct {
	// Shape of the action for `changeName`
	Change        string          `json:"change"`
	PreviousValue LocalizedString `json:"previousValue"`
	NextValue     LocalizedString `json:"nextValue"`
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

type ChangeMasterVariantChange struct {
	// Update action for `changeMasterVariant`
	Change        string  `json:"change"`
	CatalogData   string  `json:"catalogData"`
	PreviousValue Variant `json:"previousValue"`
	NextValue     Variant `json:"nextValue"`
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

type ChangeNameChange struct {
	// Shape of the action for `changeName`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type ChangeOrderHintChange struct {
	// Shape of the action for `changeOrderHint`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type ChangeOrderStateChange struct {
	// Update action for `changeOrderState`
	Change        string     `json:"change"`
	NextValue     OrderState `json:"nextValue"`
	PreviousValue OrderState `json:"previousValue"`
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

type ChangeParentChange struct {
	// Shape of the action for `changeParent`
	Change        string    `json:"change"`
	PreviousValue Reference `json:"previousValue"`
	NextValue     Reference `json:"nextValue"`
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

type ChangePaymentStateChange struct {
	// Update action for `changePaymentState`
	Change        string       `json:"change"`
	NextValue     PaymentState `json:"nextValue"`
	PreviousValue PaymentState `json:"previousValue"`
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

type ChangePlainEnumValueLabelChange struct {
	// Update action for `changePlainEnumValueLabel` on types
	Change string `json:"change"`
	// The name of the attribute updated.
	AttributeName string `json:"attributeName"`
	// Key of the values that was updated
	ValueKey      string `json:"valueKey"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type ChangePlainEnumValueOrderChange struct {
	// Update action for `changePlainEnumValueOrder` on product types
	Change string `json:"change"`
	// The name of the attribute updated.
	AttributeName string      `json:"attributeName"`
	NextValue     []EnumValue `json:"nextValue"`
	PreviousValue []EnumValue `json:"previousValue"`
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

type ChangePredicateChange struct {
	// Shape of the action for `changePredicate`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type ChangePriceChange struct {
	// Update action for changing prices
	Change        string `json:"change"`
	CatalogData   string `json:"catalogData"`
	PriceId       string `json:"priceId"`
	PreviousValue Price  `json:"previousValue"`
	NextValue     Price  `json:"nextValue"`
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
*	Update action for `changeQuantity` on inventories
 */
type ChangeQuantityChange struct {
	Change        string                 `json:"change"`
	NextValue     InventoryQuantityValue `json:"nextValue"`
	PreviousValue InventoryQuantityValue `json:"previousValue"`
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
*	Update action for `changeQuoteRequestState` on `quote-request`
 */
type ChangeQuoteRequestStateChange struct {
	Change        string            `json:"change"`
	NextValue     QuoteRequestState `json:"nextValue"`
	PreviousValue QuoteRequestState `json:"previousValue"`
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
*	Update action for `changeQuoteState` on `quote`
 */
type ChangeQuoteStateChange struct {
	Change        string     `json:"change"`
	NextValue     QuoteState `json:"nextValue"`
	PreviousValue QuoteState `json:"previousValue"`
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

type ChangeRequiresDiscountCodeChange struct {
	// Shape of the action for `changeRequiresDiscountCode`
	Change        string `json:"change"`
	PreviousValue bool   `json:"previousValue"`
	NextValue     bool   `json:"nextValue"`
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
	// Update action for `changeReviewRatingStatistics`
	Change        string                 `json:"change"`
	NextValue     ReviewRatingStatistics `json:"nextValue"`
	PreviousValue ReviewRatingStatistics `json:"previousValue"`
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

type ChangeShipmentStateChange struct {
	// Update action for `changeShipmentState`
	Change        string        `json:"change"`
	NextValue     ShipmentState `json:"nextValue"`
	PreviousValue ShipmentState `json:"previousValue"`
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

type ChangeShoppingListLineItemQuantityChange struct {
	Change        string                    `json:"change"`
	LineItem      ShoppingListLineItemValue `json:"lineItem"`
	PreviousValue int                       `json:"previousValue"`
	NextValue     int                       `json:"nextValue"`
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

type ChangeShoppingListLineItemsOrderChange struct {
	Change        string                      `json:"change"`
	PreviousValue []ShoppingListLineItemValue `json:"previousValue"`
	NextValue     []ShoppingListLineItemValue `json:"nextValue"`
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

type ChangeSlugChange struct {
	// Shape of the action for `changeSlug`
	Change        string          `json:"change"`
	PreviousValue LocalizedString `json:"previousValue"`
	NextValue     LocalizedString `json:"nextValue"`
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

type ChangeSortOrderChange struct {
	// Shape of the action for `changeSortOrder`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type ChangeStackingModeChange struct {
	// Update action for `changeStackingMode` on cart discounts
	Change        string       `json:"change"`
	NextValue     StackingMode `json:"nextValue"`
	PreviousValue StackingMode `json:"previousValue"`
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
*	Update action for `changeStagedQuoteState` on `staged-quote`
 */
type ChangeStagedQuoteStateChange struct {
	Change        string           `json:"change"`
	NextValue     StagedQuoteState `json:"nextValue"`
	PreviousValue StagedQuoteState `json:"previousValue"`
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

type ChangeStateTypeChange struct {
	// Update action for `changeType` on state
	Change        string    `json:"change"`
	PreviousValue StateType `json:"previousValue"`
	NextValue     StateType `json:"nextValue"`
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

type ChangeTargetChange struct {
	// Update action for `changeTarget` on cart discounts
	Change        string                  `json:"change"`
	NextValue     ChangeTargetChangeValue `json:"nextValue"`
	PreviousValue ChangeTargetChangeValue `json:"previousValue"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ChangeTargetChange) UnmarshalJSON(data []byte) error {
	type Alias ChangeTargetChange
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.NextValue != nil {
		var err error
		obj.NextValue, err = mapDiscriminatorChangeTargetChangeValue(obj.NextValue)
		if err != nil {
			return err
		}
	}
	if obj.PreviousValue != nil {
		var err error
		obj.PreviousValue, err = mapDiscriminatorChangeTargetChangeValue(obj.PreviousValue)
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

type ChangeTaxCalculationModeChange struct {
	// Shape of the action for `changeTaxCalculationMode`
	Change        string             `json:"change"`
	PreviousValue TaxCalculationMode `json:"previousValue"`
	NextValue     TaxCalculationMode `json:"nextValue"`
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

type ChangeTaxModeChange struct {
	// Shape of the action for `changeTaxMode`
	Change        string  `json:"change"`
	PreviousValue TaxMode `json:"previousValue"`
	NextValue     TaxMode `json:"nextValue"`
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

type ChangeTaxRoundingModeChange struct {
	// Shape of the action for `changeTaxRoundingMode`
	Change        string       `json:"change"`
	PreviousValue RoundingMode `json:"previousValue"`
	NextValue     RoundingMode `json:"nextValue"`
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

type ChangeTextLineItemNameChange struct {
	// Update action for `changeTextLineItemName`
	Change        string            `json:"change"`
	TextLineItem  TextLineItemValue `json:"textLineItem"`
	NextValue     LocalizedString   `json:"nextValue"`
	PreviousValue LocalizedString   `json:"previousValue"`
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

type ChangeTextLineItemQuantityChange struct {
	Change        string            `json:"change"`
	TextLineItem  TextLineItemValue `json:"textLineItem"`
	PreviousValue int               `json:"previousValue"`
	NextValue     int               `json:"nextValue"`
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

type ChangeTextLineItemsOrderChange struct {
	Change        string              `json:"change"`
	PreviousValue []TextLineItemValue `json:"previousValue"`
	NextValue     []TextLineItemValue `json:"nextValue"`
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

type ChangeTransactionInteractionIdChange struct {
	// Update action for `changeTransactionInteractionId` on payments
	Change        string                 `json:"change"`
	Transaction   TransactionChangeValue `json:"transaction"`
	NextValue     string                 `json:"nextValue"`
	PreviousValue string                 `json:"previousValue"`
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

type ChangeTransactionStateChange struct {
	// Update action for `changeTransactionState` on payments
	Change        string                 `json:"change"`
	Transaction   TransactionChangeValue `json:"transaction"`
	NextValue     TransactionState       `json:"nextValue"`
	PreviousValue TransactionState       `json:"previousValue"`
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

type ChangeTransactionTimestampChange struct {
	// Update action for `changeTransactionTimestamp` on payments
	Change        string                 `json:"change"`
	Transaction   TransactionChangeValue `json:"transaction"`
	NextValue     string                 `json:"nextValue"`
	PreviousValue string                 `json:"previousValue"`
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

type ChangeValueChange struct {
	// Update action for `changeValue` on cart discounts and product discounts
	Change        string                 `json:"change"`
	NextValue     ChangeValueChangeValue `json:"nextValue"`
	PreviousValue ChangeValueChangeValue `json:"previousValue"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ChangeValueChange) UnmarshalJSON(data []byte) error {
	type Alias ChangeValueChange
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.NextValue != nil {
		var err error
		obj.NextValue, err = mapDiscriminatorChangeValueChangeValue(obj.NextValue)
		if err != nil {
			return err
		}
	}
	if obj.PreviousValue != nil {
		var err error
		obj.PreviousValue, err = mapDiscriminatorChangeValueChangeValue(obj.PreviousValue)
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

type MoveImageToPositionChange struct {
	// Update actions for moving images
	Change        string  `json:"change"`
	CatalogData   string  `json:"catalogData"`
	PreviousValue []Image `json:"previousValue"`
	NextValue     []Image `json:"nextValue"`
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

type RemoveAddressChange struct {
	// Update action for `removeAddress` action.
	Change        string  `json:"change"`
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

type RemoveAssetChange struct {
	// Update action for `removeAsset`
	Change        string `json:"change"`
	PreviousValue Asset  `json:"previousValue"`
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

type RemoveAttributeDefinitionChange struct {
	// Update action for `removeAttributeDefinition` on product types
	Change        string              `json:"change"`
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

type RemoveBillingAddressIdChange struct {
	// Update action for `removeBillingAddressId` action on customers.
	Change        string   `json:"change"`
	NextValue     []string `json:"nextValue"`
	PreviousValue []string `json:"previousValue"`
	Address       Address  `json:"address"`
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

type RemoveChannelRolesChange struct {
	Change        string        `json:"change"`
	PreviousValue []ChannelRole `json:"previousValue"`
	NextValue     []ChannelRole `json:"nextValue"`
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

type RemoveCustomLineItemChange struct {
	// Update action for adding and removing custom line items
	Change        string         `json:"change"`
	NextValue     CustomLineItem `json:"nextValue"`
	PreviousValue CustomLineItem `json:"previousValue"`
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

type RemoveDeliveryItemsChange struct {
	// Update action for `removeDelivery`
	Change        string   `json:"change"`
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

type RemoveDiscountCodeChange struct {
	// Update action for `removeDiscountCode`
	Change        string           `json:"change"`
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

type RemoveEnumValuesChange struct {
	// Update action for `removeEnumValues` on product types
	Change string `json:"change"`
	// The name of the attribute updated.
	AttributeName string    `json:"attributeName"`
	PreviousValue EnumValue `json:"previousValue"`
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

type RemoveFieldDefinitionChange struct {
	// Update action for `removeFieldDefinition` on payments
	Change        string          `json:"change"`
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

type RemoveFromCategoryChange struct {
	// Update action for `addToCategory`
	Change        string      `json:"change"`
	Category      Reference   `json:"category"`
	PreviousValue []Reference `json:"previousValue"`
	NextValue     []Reference `json:"nextValue"`
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

type RemoveImageChange struct {
	// Update actions for removing images
	Change        string  `json:"change"`
	CatalogData   string  `json:"catalogData"`
	PreviousValue []Image `json:"previousValue"`
	NextValue     []Image `json:"nextValue"`
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

type RemoveItemShippingAddressesChange struct {
	// Update action for `removeItemShippingAddress`
	Change        string  `json:"change"`
	NextValue     Address `json:"nextValue"`
	PreviousValue Address `json:"previousValue"`
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

type RemoveLocalizedEnumValuesChange struct {
	// Update action for `removeEnumValues` on product types
	Change string `json:"change"`
	// The name of the attribute updated.
	AttributeName string             `json:"attributeName"`
	PreviousValue LocalizedEnumValue `json:"previousValue"`
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

type RemoveLocationChange struct {
	// Update action for `removeLocation` on zones
	Change string `json:"change"`
	// Shape of the value for `addLocation` and `removeLocation` actions
	PreviousValue Location `json:"previousValue"`
	// Shape of the value for `addLocation` and `removeLocation` actions
	NextValue Location `json:"nextValue"`
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

type RemoveOrderLineItemChange struct {
	Change        string   `json:"change"`
	PreviousValue LineItem `json:"previousValue"`
	NextValue     LineItem `json:"nextValue"`
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

type RemoveParcelFromDeliveryChange struct {
	// Update action for `removeParcelFromDelivery`
	Change        string `json:"change"`
	DeliveryId    string `json:"deliveryId"`
	PreviousValue Parcel `json:"previousValue"`
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

type RemovePaymentChange struct {
	// Update action for `addPayment` & `removePayment`
	Change        string      `json:"change"`
	NextValue     PaymentInfo `json:"nextValue"`
	PreviousValue PaymentInfo `json:"previousValue"`
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

type RemovePriceChange struct {
	// Update action for removing prices
	Change        string `json:"change"`
	CatalogData   string `json:"catalogData"`
	PriceId       string `json:"priceId"`
	PreviousValue Price  `json:"previousValue"`
	NextValue     Price  `json:"nextValue"`
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

type RemoveProductChange struct {
	// Update action for when a product is unassigned from a product selection
	Change        string    `json:"change"`
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

type RemovePropertyChange struct {
	// Update action for `removeProperty` on custom objects
	Change string `json:"change"`
	// Value path to the property that was removed
	Path          string      `json:"path"`
	PreviousValue interface{} `json:"previousValue"`
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

type RemoveShippingAddressIdChange struct {
	// Update action for `removeShippingAddressId` action on customers.
	Change        string   `json:"change"`
	NextValue     []string `json:"nextValue"`
	PreviousValue []string `json:"previousValue"`
	Address       Address  `json:"address"`
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

type RemoveShoppingListLineItemChange struct {
	Change        string   `json:"change"`
	PreviousValue LineItem `json:"previousValue"`
	NextValue     LineItem `json:"nextValue"`
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

type RemoveStateRolesChange struct {
	Change        string      `json:"change"`
	PreviousValue []StateRole `json:"previousValue"`
	NextValue     []StateRole `json:"nextValue"`
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

type RemoveTaxRateChange struct {
	// Update action for `removeTaxRate` on tax categories
	Change string `json:"change"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	PreviousValue TaxRate `json:"previousValue"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	NextValue TaxRate `json:"nextValue"`
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

type RemoveTextLineItemChange struct {
	Change        string       `json:"change"`
	PreviousValue TextLineItem `json:"previousValue"`
	NextValue     TextLineItem `json:"nextValue"`
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

type RemoveVariantChange struct {
	// Update action for `removeVariant`
	Change        string  `json:"change"`
	CatalogData   string  `json:"catalogData"`
	PreviousValue Variant `json:"previousValue"`
	NextValue     Variant `json:"nextValue"`
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

type SetAddressChange struct {
	// Update action for `setAddress` action.
	Change        string  `json:"change"`
	NextValue     Address `json:"nextValue"`
	PreviousValue Address `json:"previousValue"`
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

type SetAnonymousIdChange struct {
	// Shape of the action for `setAnonymousId`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetApplicationVersionChange struct {
	// Internal Update action for `setApplicationVersion`
	Change        string `json:"change"`
	PreviousValue int    `json:"previousValue"`
	NextValue     int    `json:"nextValue"`
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

type SetAssetCustomFieldChange struct {
	// Update action for `setAssetCustomField`
	Change        string           `json:"change"`
	Name          string           `json:"name"`
	CustomTypeId  string           `json:"customTypeId"`
	Asset         AssetChangeValue `json:"asset"`
	NextValue     interface{}      `json:"nextValue"`
	PreviousValue interface{}      `json:"previousValue"`
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

type SetAssetCustomTypeChange struct {
	// Update action for `setAssetCustomType`
	Change        string           `json:"change"`
	Asset         AssetChangeValue `json:"asset"`
	NextValue     CustomFields     `json:"nextValue"`
	PreviousValue CustomFields     `json:"previousValue"`
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

type SetAssetDescriptionChange struct {
	// Update action for `setAssetDescription`
	Change        string           `json:"change"`
	Asset         AssetChangeValue `json:"asset"`
	NextValue     LocalizedString  `json:"nextValue"`
	PreviousValue LocalizedString  `json:"previousValue"`
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

type SetAssetKeyChange struct {
	// Update action for `setAssetKey`
	Change        string           `json:"change"`
	Asset         AssetChangeValue `json:"asset"`
	NextValue     string           `json:"nextValue"`
	PreviousValue string           `json:"previousValue"`
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

type SetAssetSourcesChange struct {
	// Update action for `setAssetSources`
	Change        string           `json:"change"`
	Asset         AssetChangeValue `json:"asset"`
	NextValue     []AssetSource    `json:"nextValue"`
	PreviousValue []AssetSource    `json:"previousValue"`
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

type SetAssetTagsChange struct {
	// Update action for `setAssetTags`
	Change        string           `json:"change"`
	Asset         AssetChangeValue `json:"asset"`
	NextValue     []string         `json:"nextValue"`
	PreviousValue []string         `json:"previousValue"`
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

type SetAttributeChange struct {
	// Update action for `setAttribute`
	Change        string         `json:"change"`
	CatalogData   string         `json:"catalogData"`
	PreviousValue AttributeValue `json:"previousValue"`
	NextValue     AttributeValue `json:"nextValue"`
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

type SetAuthenticationModeChange struct {
	// Update action for `setAuthenticationMode`
	Change        string             `json:"change"`
	PreviousValue AuthenticationMode `json:"previousValue"`
	NextValue     AuthenticationMode `json:"nextValue"`
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

type SetAuthorNameChange struct {
	// Shape of the action for `setAuthorName`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetBillingAddressChange struct {
	// Update action for `setBillingAddress`
	Change        string  `json:"change"`
	NextValue     Address `json:"nextValue"`
	PreviousValue Address `json:"previousValue"`
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

type SetCartPredicateChange struct {
	// Shape of the action for `setCartPredicate`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetCategoryOrderHintChange struct {
	// Update action for `setCategoryOrderHint`
	Change        string             `json:"change"`
	CatalogData   string             `json:"catalogData"`
	CategoryId    string             `json:"categoryId"`
	PreviousValue CategoryOrderHints `json:"previousValue"`
	NextValue     CategoryOrderHints `json:"nextValue"`
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

type SetChannelRolesChange struct {
	Change        string        `json:"change"`
	PreviousValue []ChannelRole `json:"previousValue"`
	NextValue     []ChannelRole `json:"nextValue"`
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

type SetCompanyNameChange struct {
	// Shape of the action for `setCompanyName`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetCountryChange struct {
	// Update action for `setCountry`
	Change string `json:"change"`
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	PreviousValue string `json:"previousValue"`
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
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

type SetCustomFieldChange struct {
	// Update action for setting a custom field
	Change string `json:"change"`
	// Custom field name
	Name          string      `json:"name"`
	CustomTypeId  string      `json:"customTypeId"`
	NextValue     interface{} `json:"nextValue"`
	PreviousValue interface{} `json:"previousValue"`
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

type SetCustomLineItemCustomFieldChange struct {
	// Update action for `setCustomLineItemCustomField`
	Change           string          `json:"change"`
	Name             string          `json:"name"`
	CustomLineItem   LocalizedString `json:"customLineItem"`
	CustomLineItemId string          `json:"customLineItemId"`
	NextValue        interface{}     `json:"nextValue"`
	PreviousValue    interface{}     `json:"previousValue"`
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

type SetCustomLineItemCustomTypeChange struct {
	// Update action for `setCustomLineItemCustomType`
	Change           string          `json:"change"`
	CustomLineItem   LocalizedString `json:"customLineItem"`
	CustomLineItemId string          `json:"customLineItemId"`
	NextValue        CustomFields    `json:"nextValue"`
	PreviousValue    CustomFields    `json:"previousValue"`
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

type SetCustomLineItemMoneyChange struct {
	// Update action for `setCustomLineItemMoney`
	Change           string          `json:"change"`
	CustomLineItem   LocalizedString `json:"customLineItem"`
	CustomLineItemId string          `json:"customLineItemId"`
	NextValue        Money           `json:"nextValue"`
	PreviousValue    Money           `json:"previousValue"`
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

type SetCustomLineItemShippingDetailsChange struct {
	// Update action for `setCustomLineItemShippingDetails`
	Change           string              `json:"change"`
	CustomLineItemId string              `json:"customLineItemId"`
	NextValue        ItemShippingDetails `json:"nextValue"`
	PreviousValue    ItemShippingDetails `json:"previousValue"`
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

type SetCustomLineItemTaxAmountChange struct {
	// Update action for `setCustomLineItemTaxAmount`
	Change           string          `json:"change"`
	CustomLineItem   LocalizedString `json:"customLineItem"`
	CustomLineItemId string          `json:"customLineItemId"`
	TaxMode          TaxMode         `json:"taxMode"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	NextValue TaxRate `json:"nextValue"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	PreviousValue TaxRate `json:"previousValue"`
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
	// Update action for `setCustomLineItemTaxCategory`
	Change           string          `json:"change"`
	CustomLineItem   LocalizedString `json:"customLineItem"`
	CustomLineItemId string          `json:"customLineItemId"`
	NextValue        Reference       `json:"nextValue"`
	PreviousValue    Reference       `json:"previousValue"`
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

type SetCustomLineItemTaxRateChange struct {
	// Update action for `setCustomLineItemTaxRate`
	Change           string          `json:"change"`
	CustomLineItem   LocalizedString `json:"customLineItem"`
	CustomLineItemId string          `json:"customLineItemId"`
	TaxMode          TaxMode         `json:"taxMode"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	NextValue TaxRate `json:"nextValue"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	PreviousValue TaxRate `json:"previousValue"`
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
	// Update action for `setCustomLineItemTaxedPrice`
	Change           string          `json:"change"`
	CustomLineItem   LocalizedString `json:"customLineItem"`
	CustomLineItemId string          `json:"customLineItemId"`
	NextValue        Money           `json:"nextValue"`
	PreviousValue    Money           `json:"previousValue"`
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
	// Update action for `setCustomLineItemTotalPrice`
	Change           string          `json:"change"`
	CustomLineItem   LocalizedString `json:"customLineItem"`
	CustomLineItemId string          `json:"customLineItemId"`
	NextValue        Money           `json:"nextValue"`
	PreviousValue    Money           `json:"previousValue"`
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

type SetCustomShippingMethodChange struct {
	// Update action for `setCustomShippingMethod`
	Change        string                          `json:"change"`
	NextValue     CustomShippingMethodChangeValue `json:"nextValue"`
	PreviousValue CustomShippingMethodChangeValue `json:"previousValue"`
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

type SetCustomTypeChange struct {
	// Update action for setting a custom type
	Change        string       `json:"change"`
	NextValue     CustomFields `json:"nextValue"`
	PreviousValue CustomFields `json:"previousValue"`
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

type SetCustomerChange struct {
	// Shape of the action for `setCustomer`
	Change        string    `json:"change"`
	PreviousValue Reference `json:"previousValue"`
	NextValue     Reference `json:"nextValue"`
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

type SetCustomerEmailChange struct {
	// Shape of the action for `setCustomerEmail`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetCustomerGroupChange struct {
	// Shape of the action for `setCustomerGroup`
	Change        string    `json:"change"`
	PreviousValue Reference `json:"previousValue"`
	NextValue     Reference `json:"nextValue"`
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

type SetCustomerIdChange struct {
	// Shape of the action for `setCustomerId`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetCustomerNumberChange struct {
	// Shape of the action for `setCustomerNumber`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetDateOfBirthChange struct {
	// Shape of the action for `setDateOfBirth`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetDefaultBillingAddressChange struct {
	// Update action for `setDefaultBillingAddress` action.
	Change        string  `json:"change"`
	NextValue     Address `json:"nextValue"`
	PreviousValue Address `json:"previousValue"`
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

type SetDefaultShippingAddressChange struct {
	// Update action for `setDefaultShippingAddress` action.
	Change        string  `json:"change"`
	NextValue     Address `json:"nextValue"`
	PreviousValue Address `json:"previousValue"`
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

type SetDeleteDaysAfterLastModificationChange struct {
	// Shape of the action for `setDeleteDaysAfterLastModification`
	Change        string `json:"change"`
	PreviousValue int    `json:"previousValue"`
	NextValue     int    `json:"nextValue"`
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

type SetDeliveryAddressChange struct {
	// Update action for `setDeliveryAddress`
	Change        string  `json:"change"`
	DeliveryId    string  `json:"deliveryId"`
	NextValue     Address `json:"nextValue"`
	PreviousValue Address `json:"previousValue"`
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

type SetDeliveryItemsChange struct {
	// Update action for `setDeliveryItems`
	Change        string         `json:"change"`
	DeliveryId    string         `json:"deliveryId"`
	NextValue     []DeliveryItem `json:"nextValue"`
	PreviousValue []DeliveryItem `json:"previousValue"`
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

type SetDescriptionChange struct {
	// Shape of the action for `setDescription`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetDiscountedPriceChange struct {
	// Update action for `setDiscountedPrice`
	Change        string `json:"change"`
	CatalogData   string `json:"catalogData"`
	Variant       string `json:"variant"`
	PriceId       string `json:"priceId"`
	PreviousValue Price  `json:"previousValue"`
	NextValue     Price  `json:"nextValue"`
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

type SetDistributionChannelsChange struct {
	// Shape of the action for `setDistributionChannels`
	Change        string      `json:"change"`
	PreviousValue []Reference `json:"previousValue"`
	NextValue     []Reference `json:"nextValue"`
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

type SetExpectedDeliveryChange struct {
	// Shape of the action for `setExpectedDelivery`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetExternalIdChange struct {
	// Shape of the action for `setExternalId`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetFirstNameChange struct {
	// Shape of the action for `setFirstName`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetGeoLocationChange struct {
	// Update action for `setGeoLocation`
	Change        string      `json:"change"`
	NextValue     GeoLocation `json:"nextValue"`
	PreviousValue GeoLocation `json:"previousValue"`
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

type SetImageLabelChange struct {
	// Update action for `setImageLabel`
	Change        string `json:"change"`
	CatalogData   string `json:"catalogData"`
	PreviousValue Image  `json:"previousValue"`
	NextValue     Image  `json:"nextValue"`
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

type SetInputTipChange struct {
	// Update action for `setInputTip` on product types
	Change string `json:"change"`
	// The name of the updated attribute.
	AttributeName string          `json:"attributeName"`
	NextValue     LocalizedString `json:"nextValue"`
	PreviousValue LocalizedString `json:"previousValue"`
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

type SetInterfaceIdChange struct {
	// Shape of the action for `setInterfaceId`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetIsValidChange struct {
	// Shape of the action for `setIsValid`
	Change        string `json:"change"`
	PreviousValue bool   `json:"previousValue"`
	NextValue     bool   `json:"nextValue"`
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

type SetKeyChange struct {
	// Shape of the action for `setKey`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetLanguagesChange struct {
	// Update action for `setLanguages` on stores
	Change        string   `json:"change"`
	PreviousValue []string `json:"previousValue"`
	NextValue     []string `json:"nextValue"`
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

type SetLastNameChange struct {
	// Shape of the action for `setLastName`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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
	// Update action for `setLineItemDeactivatedAt`
	Change        string                    `json:"change"`
	LineItem      ShoppingListLineItemValue `json:"lineItem"`
	PreviousValue string                    `json:"previousValue"`
	NextValue     string                    `json:"nextValue"`
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
	// Update action for `setLineItemDiscountedPrice`
	Change        string                  `json:"change"`
	LineItem      LocalizedString         `json:"lineItem"`
	Variant       string                  `json:"variant"`
	NextValue     DiscountedLineItemPrice `json:"nextValue"`
	PreviousValue DiscountedLineItemPrice `json:"previousValue"`
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
	// Update action for `setLineItemDiscountedPricePerQuantity`
	Change        string                             `json:"change"`
	LineItem      LocalizedString                    `json:"lineItem"`
	Variant       string                             `json:"variant"`
	NextValue     DiscountedLineItemPriceForQuantity `json:"nextValue"`
	PreviousValue DiscountedLineItemPriceForQuantity `json:"previousValue"`
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

type SetLineItemDistributionChannelChange struct {
	// Update action for `setLineItemDistributionChannel`
	Change        string          `json:"change"`
	LineItem      LocalizedString `json:"lineItem"`
	Variant       string          `json:"variant"`
	NextValue     Reference       `json:"nextValue"`
	PreviousValue Reference       `json:"previousValue"`
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

type SetLineItemPriceChange struct {
	// Update action for `setLineItemPrice`
	Change        string          `json:"change"`
	LineItem      LocalizedString `json:"lineItem"`
	NextValue     Price           `json:"nextValue"`
	PreviousValue Price           `json:"previousValue"`
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
	// Update action for `setLineItemProductKey`
	Change        string          `json:"change"`
	LineItem      LocalizedString `json:"lineItem"`
	LineItemId    string          `json:"lineItemId"`
	Variant       string          `json:"variant"`
	PreviousValue string          `json:"previousValue"`
	NextValue     string          `json:"nextValue"`
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

type SetLineItemProductSlugChange struct {
	// Update action for `setLineItemProductSlug`
	Change        string          `json:"change"`
	LineItem      LocalizedString `json:"lineItem"`
	Variant       string          `json:"variant"`
	NextValue     LocalizedString `json:"nextValue"`
	PreviousValue LocalizedString `json:"previousValue"`
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

type SetLineItemShippingDetailsChange struct {
	// Update action for `setLineItemShippingDetails`
	Change        string              `json:"change"`
	LineItemId    string              `json:"lineItemId"`
	NextValue     ItemShippingDetails `json:"nextValue"`
	PreviousValue ItemShippingDetails `json:"previousValue"`
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

type SetLineItemTaxAmountChange struct {
	// Update action for `setLineItemTaxAmount`
	Change   string          `json:"change"`
	LineItem LocalizedString `json:"lineItem"`
	Variant  string          `json:"variant"`
	TaxMode  TaxMode         `json:"taxMode"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	NextValue TaxRate `json:"nextValue"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	PreviousValue TaxRate `json:"previousValue"`
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

type SetLineItemTaxRateChange struct {
	// Update action for `setLineItemTaxRate`
	Change   string          `json:"change"`
	LineItem LocalizedString `json:"lineItem"`
	Variant  string          `json:"variant"`
	TaxMode  TaxMode         `json:"taxMode"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	NextValue TaxRate `json:"nextValue"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	PreviousValue TaxRate `json:"previousValue"`
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
	// Update action for `setLineItemTaxedPrice`
	Change        string          `json:"change"`
	LineItem      LocalizedString `json:"lineItem"`
	LineItemId    string          `json:"lineItemId"`
	NextValue     TaxedItemPrice  `json:"nextValue"`
	PreviousValue TaxedItemPrice  `json:"previousValue"`
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

type SetLineItemTotalPriceChange struct {
	// Update action for `setLineItemTotalPrice`
	Change        string          `json:"change"`
	LineItem      LocalizedString `json:"lineItem"`
	NextValue     Money           `json:"nextValue"`
	PreviousValue Money           `json:"previousValue"`
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

type SetLocaleChange struct {
	// Update action for `setLocale` on reviews
	Change string `json:"change"`
	// A locale of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag).
	PreviousValue string `json:"previousValue"`
	// A locale of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag).
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

type SetLocalizedDescriptionChange struct {
	// Shape of the action for `setDescription`
	Change        string          `json:"change"`
	PreviousValue LocalizedString `json:"previousValue"`
	NextValue     LocalizedString `json:"nextValue"`
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

type SetMaxApplicationsChange struct {
	// Shape of the action for `setMaxApplications`
	Change        string `json:"change"`
	PreviousValue int    `json:"previousValue"`
	NextValue     int    `json:"nextValue"`
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

type SetMaxApplicationsPerCustomerChange struct {
	// Shape of the action for `setMaxApplicationsPerCustomer`
	Change        string `json:"change"`
	PreviousValue int    `json:"previousValue"`
	NextValue     int    `json:"nextValue"`
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

type SetMetaDescriptionChange struct {
	// Shape of the action for `setMetaDescription`
	Change        string          `json:"change"`
	PreviousValue LocalizedString `json:"previousValue"`
	NextValue     LocalizedString `json:"nextValue"`
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

type SetMetaKeywordsChange struct {
	// Shape of the action for `setMetaKeywords`
	Change        string          `json:"change"`
	PreviousValue LocalizedString `json:"previousValue"`
	NextValue     LocalizedString `json:"nextValue"`
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

type SetMetaTitleChange struct {
	// Shape of the action for `setMetaTitle`
	Change        string          `json:"change"`
	PreviousValue LocalizedString `json:"previousValue"`
	NextValue     LocalizedString `json:"nextValue"`
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

type SetMethodInfoInterfaceChange struct {
	// Shape of the action for `setMethodInfoInterface`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetMethodInfoMethodChange struct {
	// Shape of the action for `setMethodInfoMethod`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetMethodInfoNameChange struct {
	// Shape of the action for `setMethodInfoName`
	Change        string          `json:"change"`
	PreviousValue LocalizedString `json:"previousValue"`
	NextValue     LocalizedString `json:"nextValue"`
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

type SetMiddleNameChange struct {
	// Shape of the action for `setMiddleName`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetNameChange struct {
	// Shape of the action for `setName`
	Change        string          `json:"change"`
	PreviousValue LocalizedString `json:"previousValue"`
	NextValue     LocalizedString `json:"nextValue"`
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

type SetOrderLineItemCustomFieldChange struct {
	// Update action for `setLineItemCustomField`
	Change        string          `json:"change"`
	CustomTypeId  string          `json:"customTypeId"`
	Name          string          `json:"name"`
	Variant       string          `json:"variant"`
	LineItem      LocalizedString `json:"lineItem"`
	NextValue     interface{}     `json:"nextValue"`
	PreviousValue interface{}     `json:"previousValue"`
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

type SetOrderLineItemCustomTypeChange struct {
	// Update action for `setLineItemCustomType`
	Change        string          `json:"change"`
	LineItem      LocalizedString `json:"lineItem"`
	Variant       string          `json:"variant"`
	NextValue     CustomFields    `json:"nextValue"`
	PreviousValue CustomFields    `json:"previousValue"`
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

type SetOrderNumberChange struct {
	// Shape of the action for `setOrderNumber`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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
	// Update action for `setOrderTaxedPrice`
	Change        string         `json:"change"`
	TaxMode       TaxMode        `json:"taxMode"`
	NextValue     TaxedItemPrice `json:"nextValue"`
	PreviousValue TaxedItemPrice `json:"previousValue"`
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

type SetOrderTotalPriceChange struct {
	// Update action for `setOrderTotalPrice`
	Change        string `json:"change"`
	NextValue     Money  `json:"nextValue"`
	PreviousValue Money  `json:"previousValue"`
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

type SetOrderTotalTaxChange struct {
	// Update action for `setOrderTotalTax`
	Change        string  `json:"change"`
	TaxMode       TaxMode `json:"taxMode"`
	NextValue     Money   `json:"nextValue"`
	PreviousValue Money   `json:"previousValue"`
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

type SetParcelItemsChange struct {
	// Update action for `setParcelItems`
	Change        string            `json:"change"`
	Parcel        ParcelChangeValue `json:"parcel"`
	NextValue     []DeliveryItem    `json:"nextValue"`
	PreviousValue []DeliveryItem    `json:"previousValue"`
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

type SetParcelMeasurementsChange struct {
	// Update action for `setParcelMeasurements`
	Change        string             `json:"change"`
	Parcel        ParcelChangeValue  `json:"parcel"`
	NextValue     ParcelMeasurements `json:"nextValue"`
	PreviousValue ParcelMeasurements `json:"previousValue"`
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

type SetParcelTrackingDataChange struct {
	// Update action for `setParcelTrackingData`
	Change        string            `json:"change"`
	Parcel        ParcelChangeValue `json:"parcel"`
	NextValue     TrackingData      `json:"nextValue"`
	PreviousValue TrackingData      `json:"previousValue"`
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

type SetPricesChange struct {
	// Update action for `setPrices`
	Change        string  `json:"change"`
	CatalogData   string  `json:"catalogData"`
	Variant       string  `json:"variant"`
	PreviousValue []Price `json:"previousValue"`
	NextValue     []Price `json:"nextValue"`
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

type SetProductCountChange struct {
	// Update action for `setProductCount`
	Change        string `json:"change"`
	PreviousValue int    `json:"previousValue"`
	NextValue     int    `json:"nextValue"`
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

type SetProductPriceCustomFieldChange struct {
	// Update action for `setProductPriceCustomField`
	Change        string       `json:"change"`
	CatalogData   string       `json:"catalogData"`
	PreviousValue CustomFields `json:"previousValue"`
	NextValue     CustomFields `json:"nextValue"`
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

type SetProductPriceCustomTypeChange struct {
	// Update action for `setProductPriceCustomType`
	Change        string       `json:"change"`
	CatalogData   string       `json:"catalogData"`
	PreviousValue CustomFields `json:"previousValue"`
	NextValue     CustomFields `json:"nextValue"`
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

type SetProductSelectionsChange struct {
	// Update action for `setProductSelections`
	Change        string                    `json:"change"`
	PreviousValue []ProductSelectionSetting `json:"previousValue"`
	NextValue     []ProductSelectionSetting `json:"nextValue"`
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

type SetProductVariantKeyChange struct {
	// Update action for `setProductVariantKey`
	Change        string `json:"change"`
	CatalogData   string `json:"catalogData"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetPropertyChange struct {
	// Update action for `setProperty` on custom objects
	Change string `json:"change"`
	// Value path to the property that was changed
	Path          string      `json:"path"`
	NextValue     interface{} `json:"nextValue"`
	PreviousValue interface{} `json:"previousValue"`
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

type SetRatingChange struct {
	// Shape of the action for `setRating`
	Change        string `json:"change"`
	PreviousValue int    `json:"previousValue"`
	NextValue     int    `json:"nextValue"`
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
	// Update action for `setReservations` on inventories
	Change        string        `json:"change"`
	NextValue     []Reservation `json:"nextValue"`
	PreviousValue []Reservation `json:"previousValue"`
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

type SetRestockableInDaysChange struct {
	// Shape of the action for `setRestockableInDays`
	Change        string `json:"change"`
	PreviousValue int    `json:"previousValue"`
	NextValue     int    `json:"nextValue"`
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

type SetReturnPaymentStateChange struct {
	// Update action for `setReturnPaymentState`
	Change        string             `json:"change"`
	NextValue     ReturnPaymentState `json:"nextValue"`
	PreviousValue ReturnPaymentState `json:"previousValue"`
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

type SetReturnShipmentStateChange struct {
	// Update action for `setReturnShipmentState`
	Change        string              `json:"change"`
	NextValue     ReturnShipmentState `json:"nextValue"`
	PreviousValue ReturnShipmentState `json:"previousValue"`
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

type SetSalutationChange struct {
	// Shape of the action for `setSalutation`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetSearchKeywordsChange struct {
	// Update action for `setSearchKeywords`
	Change        string         `json:"change"`
	CatalogData   string         `json:"catalogData"`
	PreviousValue SearchKeywords `json:"previousValue"`
	NextValue     SearchKeywords `json:"nextValue"`
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

type SetSellerCommentChange struct {
	// Shape of the action for `setSellerComment`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetShippingAddressChange struct {
	// Update action for `setShippingAddress`
	Change        string  `json:"change"`
	NextValue     Address `json:"nextValue"`
	PreviousValue Address `json:"previousValue"`
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
	// Update action for `setShippingInfoPrice`
	Change        string `json:"change"`
	NextValue     Money  `json:"nextValue"`
	PreviousValue Money  `json:"previousValue"`
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
	// Update action for `setShippingInfoTaxedPrice`
	Change        string     `json:"change"`
	NextValue     TaxedPrice `json:"nextValue"`
	PreviousValue TaxedPrice `json:"previousValue"`
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

type SetShippingMethodChange struct {
	// Update action for `setShippingMethod`
	Change        string                    `json:"change"`
	NextValue     ShippingMethodChangeValue `json:"nextValue"`
	PreviousValue ShippingMethodChangeValue `json:"previousValue"`
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

type SetShippingMethodTaxAmountChange struct {
	// Update action for `setShippingMethodTaxAmount`
	Change        string                             `json:"change"`
	TaxMode       TaxMode                            `json:"taxMode"`
	NextValue     ShippingMethodTaxAmountChangeValue `json:"nextValue"`
	PreviousValue ShippingMethodTaxAmountChangeValue `json:"previousValue"`
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

type SetShippingMethodTaxRateChange struct {
	// Update action for `setShippingMethodTaxRate`
	Change  string  `json:"change"`
	TaxMode TaxMode `json:"taxMode"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	NextValue TaxRate `json:"nextValue"`
	// Shape of the value for `addTaxRate` and `removeTaxRate` actions
	PreviousValue TaxRate `json:"previousValue"`
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
	// Update action for `setShippingRate`
	Change        string `json:"change"`
	NextValue     Money  `json:"nextValue"`
	PreviousValue Money  `json:"previousValue"`
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

type SetShippingRateInputChange struct {
	// Update action for `setShippingRateInput`
	Change        string      `json:"change"`
	NextValue     interface{} `json:"nextValue"`
	PreviousValue interface{} `json:"previousValue"`
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

type SetShoppingListLineItemCustomFieldChange struct {
	// Update action for `setLineItemCustomField`
	Change        string                    `json:"change"`
	Name          string                    `json:"name"`
	CustomTypeId  string                    `json:"customTypeId"`
	LineItem      ShoppingListLineItemValue `json:"lineItem"`
	NextValue     interface{}               `json:"nextValue"`
	PreviousValue interface{}               `json:"previousValue"`
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

type SetShoppingListLineItemCustomTypeChange struct {
	// Update action for `setLineItemCustomType`
	Change        string                    `json:"change"`
	LineItem      ShoppingListLineItemValue `json:"lineItem"`
	NextValue     CustomFields              `json:"nextValue"`
	PreviousValue CustomFields              `json:"previousValue"`
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

type SetSkuChange struct {
	// Update action for `setSku`
	Change        string `json:"change"`
	CatalogData   string `json:"catalogData"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetSlugChange struct {
	// Shape of the action for `setSlug`
	Change        string          `json:"change"`
	PreviousValue LocalizedString `json:"previousValue"`
	NextValue     LocalizedString `json:"nextValue"`
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

type SetStateRolesChange struct {
	Change        string      `json:"change"`
	PreviousValue []StateRole `json:"previousValue"`
	NextValue     []StateRole `json:"nextValue"`
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

type SetStatusInterfaceCodeChange struct {
	// Shape of the action for `setStatusInterfaceCode`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetStatusInterfaceTextChange struct {
	// Shape of the action for `setStatusInterfaceText`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetStoreChange struct {
	// Shape of the action for `setStore`
	Change        string    `json:"change"`
	PreviousValue Reference `json:"previousValue"`
	NextValue     Reference `json:"nextValue"`
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

type SetStoresChange struct {
	// Shape of the action for `setStores`
	Change        string      `json:"change"`
	PreviousValue []Reference `json:"previousValue"`
	NextValue     []Reference `json:"nextValue"`
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

type SetSupplyChannelChange struct {
	// Shape of the action for `setSupplyChannel`
	Change        string    `json:"change"`
	PreviousValue Reference `json:"previousValue"`
	NextValue     Reference `json:"nextValue"`
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

type SetSupplyChannelsChange struct {
	// Shape of the action for `setSupplyChannels`
	Change        string      `json:"change"`
	PreviousValue []Reference `json:"previousValue"`
	NextValue     []Reference `json:"nextValue"`
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

type SetTargetChange struct {
	// Shape of the action for `setTarget`
	Change        string    `json:"change"`
	PreviousValue Reference `json:"previousValue"`
	NextValue     Reference `json:"nextValue"`
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

type SetTaxCategoryChange struct {
	// Shape of the action for `setTaxCategory`
	Change        string    `json:"change"`
	PreviousValue Reference `json:"previousValue"`
	NextValue     Reference `json:"nextValue"`
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

type SetTextChange struct {
	// Shape of the action for `setText`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetTextLineItemCustomFieldChange struct {
	// Update action for `setTextLineItemCustomField`
	Change        string            `json:"change"`
	Name          string            `json:"name"`
	CustomTypeId  string            `json:"customTypeId"`
	TextLineItem  TextLineItemValue `json:"textLineItem"`
	NextValue     interface{}       `json:"nextValue"`
	PreviousValue interface{}       `json:"previousValue"`
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

type SetTextLineItemCustomTypeChange struct {
	// Update action for `setTextLineItemCustomType`
	Change        string            `json:"change"`
	TextLineItem  TextLineItemValue `json:"textLineItem"`
	NextValue     CustomFields      `json:"nextValue"`
	PreviousValue CustomFields      `json:"previousValue"`
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

type SetTextLineItemDescriptionChange struct {
	Change        string            `json:"change"`
	TextLineItem  TextLineItemValue `json:"textLineItem"`
	PreviousValue LocalizedString   `json:"previousValue"`
	NextValue     LocalizedString   `json:"nextValue"`
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

type SetTitleChange struct {
	// Shape of the action for `setTitle`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetTransitionsChange struct {
	// Shape of the action for `setTransitions`
	Change        string      `json:"change"`
	PreviousValue []Reference `json:"previousValue"`
	NextValue     []Reference `json:"nextValue"`
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

type SetValidFromAndUntilChange struct {
	Change string `json:"change"`
	// Shape of the value for `setValidFromAndUntil` action
	PreviousValue ValidFromAndUntilValue `json:"previousValue"`
	// Shape of the value for `setValidFromAndUntil` action
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

type SetValidFromChange struct {
	// Shape of the action for `setValidFrom`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetValidToChange struct {
	// Shape of the action for `setValidTo`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetValidUntilChange struct {
	// Shape of the action for `setValidUntil`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type SetValueChange struct {
	// Update action for `setValue` on custom objects
	Change        string      `json:"change"`
	NextValue     interface{} `json:"nextValue"`
	PreviousValue interface{} `json:"previousValue"`
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

type SetVariantAvailabilityChange struct {
	// Update action for `setVariantAvailability`
	Change        string                     `json:"change"`
	CatalogData   string                     `json:"catalogData"`
	Variant       string                     `json:"variant"`
	PreviousValue ProductVariantAvailability `json:"previousValue"`
	NextValue     ProductVariantAvailability `json:"nextValue"`
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

type SetVatIdChange struct {
	// Shape of the action for `setVatId`
	Change        string `json:"change"`
	PreviousValue string `json:"previousValue"`
	NextValue     string `json:"nextValue"`
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

type TransitionCustomLineItemStateChange struct {
	// Update action for `transitionCustomLineItemState`
	Change        string      `json:"change"`
	LineItemId    string      `json:"lineItemId"`
	StateId       string      `json:"stateId"`
	NextValue     []ItemState `json:"nextValue"`
	PreviousValue []ItemState `json:"previousValue"`
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

type TransitionLineItemStateChange struct {
	// Update action for `transitionLineItemState`
	Change        string      `json:"change"`
	LineItemId    string      `json:"lineItemId"`
	StateId       string      `json:"stateId"`
	NextValue     []ItemState `json:"nextValue"`
	PreviousValue []ItemState `json:"previousValue"`
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

type TransitionStateChange struct {
	// Shape of the action for `transitionState`
	Change        string    `json:"change"`
	PreviousValue Reference `json:"previousValue"`
	NextValue     Reference `json:"nextValue"`
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

type UnknownChange struct {
	Change        string      `json:"change"`
	PreviousValue interface{} `json:"previousValue"`
	NextValue     interface{} `json:"nextValue"`
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

type UpdateSyncInfoChange struct {
	// Update action for `updateSyncInfo`
	Change    string   `json:"change"`
	ChannelId string   `json:"channelId"`
	NextValue SyncInfo `json:"nextValue"`
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
