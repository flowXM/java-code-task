package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	WalletId      uuid.UUID       `json:"walletId"`
	OperationType OperationType   `json:"operationType"`
	Amount        decimal.Decimal `json:"amount"`
}
