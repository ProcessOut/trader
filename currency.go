package trader

import (
	"github.com/shopspring/decimal"
)

// Currency represents a currency and its value relative to the dollar
type Currency struct {
	Code  string           `json:"code"`  // ISO 4217 code of the currency
	Value *decimal.Decimal `json:"value"` // Value of the currency, relative to USD
}
