package handlers

import (
	"fmt"
	"github.com/mikusher/bookings/pkg/config"
	"github.com/mikusher/bookings/pkg/models"
	"github.com/mikusher/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by handler
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// StartRepo creates a new repository
func StartRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func StartHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})
}

// Generals is the handler for the generals page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors is the handler for the majors page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability is the handler for the Availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact is the handler for the Contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation is the handler for the Reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// PostAvailability is the handler for the PostAvailability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")

	w.Write([]byte(fmt.Sprintf("PostAvailability start: %s, end: %s", start, end)))
}
