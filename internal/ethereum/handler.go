package ethereum

import (
	gErrors "errors"

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
}
