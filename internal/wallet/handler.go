package wallet

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"nn-blockchain-api/pkg/errors"
	"nn-blockchain-api/pkg/respond"
)

type Handler struct {
	walletSvc Service
}

func NewHandler(walletSvc Service) *Handler {
	return &Handler{
		walletSvc: walletSvc,
	}
}

func (h *Handler) SetupRoutes(router chi.Router) {
	router.Get("/create-wallet", h.CreateWallet)
}

func (h *Handler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var dto CoinNameDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	wallet, err := h.walletSvc.CreateWallet(context.Background(), dto.Name)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewNotFound(err.Error()))
		return
	}

	respond.Respond(w, http.StatusOK, wallet)
}
