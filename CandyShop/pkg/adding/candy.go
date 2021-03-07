package adding

import (
	"github.com/shopspring/decimal"
)

type Candy struct {
	Id       string          `json:"id",omitempty`
	Category string          `json:"category"`
	Name     string          `json:"name"`
	Price    decimal.Decimal `json:"price"`
}
