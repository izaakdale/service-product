package app

import (
	"context"
	"encoding/json"
	"net/http"

	db "github.com/izaakdale/service-product/db/sqlc"
	"github.com/izaakdale/utils-go/response"
)

type InsertProductRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Price       int32  `json:"price,omitempty"`
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {

	var req InsertProductRequest
	json.NewDecoder(r.Body).Decode(&req)

	product, err := db.ClientQueries().CreateProduct(context.Background(), db.CreateProductParams{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
	})
	if err != nil {
		panic(err)
	}

	response.WriteJson(w, http.StatusCreated, product)
}
