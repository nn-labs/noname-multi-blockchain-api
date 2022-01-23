package bitcoin

import (
	"context"
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
	router.Get("/status", h.HealthCheckHandler)

	router.Post("/create-raw-tx", h.CreateRawTransaction)
	router.Post("/sign-raw-tx", h.SignRawTransaction)
	router.Post("/send-raw-tx", h.SendRawTransaction)
}

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	status, err := h.btcSvc.StatusNode(context.Background())
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	respond.Respond(w, http.StatusOK, status)
}

func (h *Handler) CreateRawTransaction(w http.ResponseWriter, r *http.Request) {
	var dto RawTransactionDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), errors.NewInternal(err.Error()))
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	transaction, err := h.btcSvc.CreateTransaction(context.Background(), &dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}
	respond.Respond(w, http.StatusOK, transaction)
}

func (h *Handler) SignRawTransaction(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) SendRawTransaction(w http.ResponseWriter, r *http.Request) {

}
