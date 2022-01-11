package health

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_HealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	handler := NewHandler()
	handler.HealthCheckHandler(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
