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
	router.Get("/status", h.HealthCheckHandler)

	router.Post("/create-raw-tx", h.CreateRawTransaction)
}

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	status, err := h.btcSvc.StatusNode()
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
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	if err := Validate(dto); err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}

	transaction, err := h.btcSvc.CreateTransaction(&dto)
	if err != nil {
		respond.Respond(w, errors.HTTPCode(err), err)
		return
	}
	respond.Respond(w, http.StatusOK, transaction)
}
