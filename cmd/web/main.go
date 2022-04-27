package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/mikusher/bookings/internal/config"
	"github.com/mikusher/bookings/internal/handlers"
	"github.com/mikusher/bookings/internal/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	// change to true if in production
	app.InProduction = false
	app.Static = "./static/"

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Sessions = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("Cannot create template cache: %v", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.StartRepo(&app)
	handlers.StartHandlers(repo)
	render.StartTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	// start server
	server := &http.Server{
		Addr:    portNumber,
		Handler: routers(&app),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
