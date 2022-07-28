package main

import (
	"net/http"

	"github.com/ggentile/bookings_go/pkg/config"
	"github.com/ggentile/bookings_go/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	mux := chi.NewRouter()

	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
