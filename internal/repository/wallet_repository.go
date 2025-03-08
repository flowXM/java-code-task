package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"java-code-task/pkg/database"
	"java-code-task/pkg/models"
	"java-code-task/pkg/validators"
)

type WalletRepository interface {
	UpdateAmount(walletId uuid.UUID, amount decimal.Decimal) error
	GetAmount(walletId uuid.UUID) (decimal.Decimal, error)
}

type walletRepository struct{}

func NewWalletRepository() WalletRepository {
	return &walletRepository{}
}

func (w *walletRepository) UpdateAmount(walletId uuid.UUID, amount decimal.Decimal) error {
	if err := validators.ValidateDecimal(amount, 2); err != nil {
		return err
	}

	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	var wallet models.Wallet

	result := db.QueryRow("SELECT * FROM wallets WHERE wallet_id = $1", walletId)
	err = result.Scan(&wallet.WalletId, &wallet.Amount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("wallet %s: not found", walletId)
		}
		return fmt.Errorf("wallet %s: %v", walletId, err)
	}

	_, err = db.Exec("UPDATE wallets SET amount = $1 WHERE wallet_id = $2", amount, walletId)
	if err != nil {
		return err
	}

	return nil
}

func (w *walletRepository) GetAmount(walletId uuid.UUID) (decimal.Decimal, error) {
	db, err := database.Connect()
	if err != nil {
		return decimal.Decimal{}, err
	}
	defer db.Close()

	var wallet models.Wallet

	result := db.QueryRow("SELECT * FROM wallets WHERE wallet_id = $1", walletId)
	err = result.Scan(&wallet.WalletId, &wallet.Amount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return decimal.Decimal{}, fmt.Errorf("wallet %s: not found", walletId)
		}
		return decimal.Decimal{}, fmt.Errorf("wallet %s: %v", walletId, err)
	}

	return wallet.Amount, nil
}
