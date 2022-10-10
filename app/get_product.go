package app

import (
	"context"
	"net/http"

	db "github.com/izaakdale/service-product/db/sqlc"
	"github.com/izaakdale/utils-go/response"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id := getParam(r.Context(), "id")

	p, err := db.ClientQueries().GetProduct(context.Background(), id)
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	response.WriteJson(w, http.StatusOK, p)
}
