package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/microcosm-cc/bluemonday"
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache        bool
	TemplateCache   map[string]*template.Template
	InfoLogger      *log.Logger
	ErrorLogger     *log.Logger
	InProduction    bool
	Session         *scs.SessionManager
	SanitizerPolicy *bluemonday.Policy
}
