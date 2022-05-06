package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mikusher/bookings/internal/config"
	"github.com/mikusher/bookings/internal/handlers"
	"github.com/mikusher/bookings/internal/helpers"
	"github.com/mikusher/bookings/internal/models"
	"github.com/mikusher/bookings/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var InfoLogger *log.Logger
var sanitizer *bluemonday.Policy

// main is the main function
func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	infoLogger := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	app.InfoLogger = infoLogger

	errorLogger := log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLogger = errorLogger

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	// security sanitizer
	sanitizer = bluemonday.StrictPolicy()
	sanitizer.AllowStandardURLs()

	app.SanitizerPolicy = sanitizer

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	helpers.NewHelpers(&app)
	return nil
}
