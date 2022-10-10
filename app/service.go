package app

import (
	"context"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	r := httprouter.New()
	r.HandlerFunc("GET", "/product/:id", GetProductHandler)
	return r
}

func getParam(ctx context.Context, key string) string {
	params := httprouter.ParamsFromContext(ctx)
	return params.ByName(key)
}
