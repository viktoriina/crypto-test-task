package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/viktoriina/crypto-test-task/services"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/api/rate", services.GetPrice)
	router.Post("/api/subscribe", services.Subscribe)

	http.ListenAndServe(":3000", router)
}
