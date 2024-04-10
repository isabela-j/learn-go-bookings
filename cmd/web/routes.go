package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/isabela-j/bookings/pkg/config"
	"github.com/isabela-j/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}

// func routes(app *config.AppConfig) http.Handler {
// 	multiplexer := pat.New()

// 	multiplexer.Get("/", http.HandlerFunc(handlers.Repo.Home))
// 	multiplexer.Get("/about", http.HandlerFunc(handlers.Repo.About))

// 	return multiplexer
// }
