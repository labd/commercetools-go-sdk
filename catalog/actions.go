package catalog

import "github.com/labd/commercetools-go-sdk/common"

type ProductChangeName struct {
	common.UpdateAction
	Name *common.LocalizedString
}
type ProductAddPrice struct {
	common.UpdateAction
	VariantID int
	Price     int
	Staged    bool
}
