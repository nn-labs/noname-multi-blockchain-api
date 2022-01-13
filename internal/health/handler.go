package health

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"nn-blockchain-api/pkg/respond"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes(router chi.Router) {
	router.Get("/health", h.HealthCheckHandler)
}

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	respond.Respond(w, http.StatusOK, map[string]string{"status": "OK"})
}
