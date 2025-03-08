package models

import (
	"encoding/json"
	"errors"
)

type OperationType string

const (
	Deposit  OperationType = "DEPOSIT"
	Withdraw OperationType = "WITHDRAW"
)

func (op *OperationType) IsValid() bool {
	switch *op {
	case Deposit, Withdraw:
		return true
	}
	return false
}

func (op *OperationType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	operation := OperationType(s)
	if !operation.IsValid() {
		return errors.New("invalid operation type")
	}

	*op = operation
	return nil
}
