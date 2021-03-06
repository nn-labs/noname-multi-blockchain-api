package wallet

import (
	"encoding/json"
	gErrors "errors"
	"net/http"
	"nn-blockchain-api/pkg/errors"
	"nn-blockchain-api/pkg/respond"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	walletSvc Service
}

func NewHandler(walletSvc Service) (*Handler, error) {
	if walletSvc == nil {
		return nil, gErrors.New("invalid wallet service")
	}

	return &Handler{
		walletSvc: walletSvc,
	}, nil
}

func (h *Handler) SetupRoutes(router chi.Router) {
	router.Post("/create-wallet", h.CreateWallet)
	router.Post("/create-mnemonic", h.CreateMnemonic)
}

func (h *Handler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var dto CoinNameDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	wallet, err := h.walletSvc.CreateWallet(r.Context(), dto.Name, &dto.Mnemonic)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewNotFound(err.Error()))
		return
	}

	respond.Respond(w, http.StatusOK, wallet)
}

func (h *Handler) CreateMnemonic(w http.ResponseWriter, r *http.Request) {
	var dto MnemonicDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), ErrInvalidPayload)
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	mnemonic, err := h.walletSvc.CreateMnemonic(r.Context(), dto.Length, dto.Language)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	respond.Respond(w, http.StatusOK, mnemonic)
}
