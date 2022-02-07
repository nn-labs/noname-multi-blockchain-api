package bitcoin

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"nn-blockchain-api/pkg/errors"
	"nn-blockchain-api/pkg/respond"
)

type Handler struct {
	btcSvc Service
}

func NewHandler(btcSvc Service) *Handler {
	return &Handler{
		btcSvc: btcSvc,
	}
}

func (h *Handler) SetupRoutes(router chi.Router) {
	router.Post("/status", h.HealthCheckHandler)

	// Transaction
	router.Post("/create-raw-tx", h.CreateRawTransaction)
	router.Post("/decode-raw-tx", h.DecodeRawTransaction)
	router.Post("/fund-for-raw-tx", h.FundForRawTransaction)
	router.Post("/sign-raw-tx", h.SignRawTransaction)
	router.Post("/send-raw-tx", h.SendRawTransaction)

	// Wallet/Unspent transaction list
	router.Post("/wallet-info", h.WalletInfo)
	router.Post("/create-wallet", h.CreateWallet)
	router.Post("/load-wallet", h.LoadWallet)
	router.Post("/import-address", h.ImportAddress)
	router.Post("/rescan-wallet", h.RescanWallet)
	router.Post("/list-utx", h.ListUnspent)
}

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	var dto StatusNodeDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	status, err := h.btcSvc.StatusNode(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, status)
}

func (h *Handler) CreateRawTransaction(w http.ResponseWriter, r *http.Request) {
	var dto CreateRawTransactionDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	transaction, err := h.btcSvc.CreateTransaction(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}
	respond.Respond(w, http.StatusOK, transaction)
}

func (h *Handler) DecodeRawTransaction(w http.ResponseWriter, r *http.Request) {
	var dto DecodeRawTransactionDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	decodedTx, err := h.btcSvc.DecodeTransaction(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, decodedTx)
}

func (h *Handler) FundForRawTransaction(w http.ResponseWriter, r *http.Request) {
	var dto FundForRawTransactionDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	foundedTx, err := h.btcSvc.FoundForRawTransaction(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, foundedTx)
}

func (h *Handler) SignRawTransaction(w http.ResponseWriter, r *http.Request) {
	var dto SignRawTransactionDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	signedTx, err := h.btcSvc.SignTransaction(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, signedTx)
}

func (h *Handler) SendRawTransaction(w http.ResponseWriter, r *http.Request) {
	var dto SendRawTransactionDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	transactionId, err := h.btcSvc.SendTransaction(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, transactionId)
}

func (h *Handler) WalletInfo(w http.ResponseWriter, r *http.Request) {
	var dto WalletDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	info, err := h.btcSvc.WalletInfo(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, info)
}

func (h *Handler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var dto CreateWalletDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	walletId, err := h.btcSvc.CreateWallet(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, walletId)
}

func (h *Handler) LoadWallet(w http.ResponseWriter, r *http.Request) {
	var dto LoadWalletDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	info, err := h.btcSvc.LoadWaller(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, info)
}

func (h *Handler) ImportAddress(w http.ResponseWriter, r *http.Request) {
	var dto ImportAddressDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	info, err := h.btcSvc.ImportAddress(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, info)
}

func (h *Handler) RescanWallet(w http.ResponseWriter, r *http.Request) {
	var dto RescanWalletDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	info, err := h.btcSvc.RescanWallet(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, info)
}

func (h *Handler) ListUnspent(w http.ResponseWriter, r *http.Request) {
	var dto ListUnspentDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	list, err := h.btcSvc.ListUnspent(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, list)
}
