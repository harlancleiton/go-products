package handler

import (
	"encoding/json"
	"net/http"

	"github.com/harlancleiton/go-products/internal/dto"
	"github.com/harlancleiton/go-products/internal/entity"
	"github.com/harlancleiton/go-products/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(productDB database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: productDB}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(input.Name, input.Price)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
