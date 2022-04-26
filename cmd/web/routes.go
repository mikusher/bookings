package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mikusher/bookings/pkg/config"
	"github.com/mikusher/bookings/pkg/handlers"
	"net/http"
)

func routers(app *config.AppConfig) http.Handler {
	// use mux
	mux := chi.NewRouter()

	// middleware
	mux.Use(middleware.Recoverer)
	// personal middleware
	// use csrf token
	mux.Use(NoSurf)
	//use session
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
