package products

import (
	"net/http"

	"github.com/tokzy/eccom-rest-api/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := []string{"hello", "world"}
	json.Write(w, http.StatusOK, products)
}
