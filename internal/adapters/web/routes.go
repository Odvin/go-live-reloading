package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (web *WebService) routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/v1/healthcheck", web.healthcheck)

	router.Post("/v1/customers", web.createCustomer)
	router.Get("/v1/customers/{id}", web.getCustomer)

	return router
}
