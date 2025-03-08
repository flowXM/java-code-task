package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	WalletId uuid.UUID
	Amount   decimal.Decimal
}
