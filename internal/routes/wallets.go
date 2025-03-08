package routes

import (
	"encoding/json"
	"github.com/google/uuid"
	"java-code-task/internal/models"
	"java-code-task/internal/repository"
	"java-code-task/pkg/validators"
	"net/http"
	"sync"
)

var walletMu sync.Mutex

func HandleGetWallet(db repository.WalletRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		amount, err := db.GetAmount(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		jsonBytes, err := json.Marshal(amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}

func HandlePostWallet(db repository.WalletRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction models.Transaction

		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		walletMu.Lock()
		defer walletMu.Unlock()

		amount, err := db.GetAmount(transaction.WalletId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = validators.ValidateDecimal(transaction.Amount, 2)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		switch transaction.OperationType {
		case models.Deposit:
			amount = amount.Add(transaction.Amount)
		case models.Withdraw:
			amount = amount.Sub(transaction.Amount)
		}

		err = db.UpdateAmount(transaction.WalletId, amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
