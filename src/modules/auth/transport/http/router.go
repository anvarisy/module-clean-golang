package http

import (
	"anvarisy.tech/cleangolang/src/modules/auth/endpoint"
	"anvarisy.tech/cleangolang/src/modules/auth/transport/http/handler"
	"github.com/go-chi/chi"
)

func Init(endpoint endpoint.Endpoint) *chi.Mux {
	var (
		handler = handler.Init(endpoint)
		router  = chi.NewRouter()
	)
	router.Post("/login", handler.SignIn)
	return router
}
