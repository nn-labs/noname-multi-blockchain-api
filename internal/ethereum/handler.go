package ethereum

import (
	"encoding/json"
	gErrors "errors"
	"net/http"
	"nn-blockchain-api/pkg/errors"
	"nn-blockchain-api/pkg/respond"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	ethSvc Service
}

func NewHandler(ethSvc Service) (*Handler, error) {
	if ethSvc == nil {
		return nil, gErrors.New("invalid bitcoin service")
	}

	return &Handler{
		ethSvc: ethSvc,
	}, nil
}

func (h *Handler) SetupRoutes(router chi.Router) {
	router.Post("/status", h.StatusNode)

	// Transaction
	router.Post("/create-raw-tx", h.CreateRawTransaction)
	router.Post("/sign-raw-tx", h.SignRawTransaction)
	router.Post("/send-raw-tx", h.SendRawTransaction)
}

func (h *Handler) StatusNode(w http.ResponseWriter, r *http.Request) {
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

	status, err := h.ethSvc.StatusNode(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	if status == nil {
		respond.Respond(w, http.StatusOK, &NodeInfoDTO{
			SyncMessage: "node has not synced yet",
		})
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

	transaction, err := h.ethSvc.CreateTransaction(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}
	respond.Respond(w, http.StatusOK, transaction)
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

	signedTx, err := h.ethSvc.SignTransaction(r.Context(), &dto)
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

	transactionId, err := h.ethSvc.SendTransaction(r.Context(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, transactionId)
}
