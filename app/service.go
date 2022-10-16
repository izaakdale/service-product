package app

import (
	"context"

	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	r := httprouter.New()
	r.HandlerFunc("GET", "/", ListProductsHandler)
	r.HandlerFunc("GET", "/:id", GetProductHandler)
	r.HandlerFunc("POST", "/", CreateProductHandler)
	return r
}

func getParam(ctx context.Context, key string) string {
	params := httprouter.ParamsFromContext(ctx)
	return params.ByName(key)
}
