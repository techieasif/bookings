package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github/techieasif/bookings/pkg/config"
	"github/techieasif/bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//custom middleware.
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.HomePage)
	mux.Get("/about", handlers.Repo.AboutPage)
	return mux
}
