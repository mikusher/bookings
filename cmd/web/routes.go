package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mikusher/bookings/internal/config"
	"github.com/mikusher/bookings/internal/handlers"
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
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)

	// load file Server
	fileServer := http.FileServer(http.Dir(app.Static))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
