package handlers

import (
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
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
