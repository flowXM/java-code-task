package routes

import (
	"java-code-task/internal/repository"
	"java-code-task/pkg/utils"
	"net/http"
)

func Init(mux *http.ServeMux) {
	walletRepository := repository.NewWalletRepository()
	mux.Handle("GET /api/v1/wallets/{id}", utils.RateLimitedHandler(HandleGetWallet(walletRepository)))
	mux.Handle("POST /api/v1/wallet", utils.RateLimitedHandler(HandlePostWallet(walletRepository)))
}
