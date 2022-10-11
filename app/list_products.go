package app

import (
	"context"
	"net/http"

	db "github.com/izaakdale/service-product/db/sqlc"
	"github.com/izaakdale/utils-go/response"
)

func ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	lp, err := db.ClientQueries().ListProducts(context.Background(), db.ListProductsParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	response.WriteJson(w, http.StatusOK, lp)
}
