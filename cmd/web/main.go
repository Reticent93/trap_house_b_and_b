package main

import (
	"fmt"
	"github.com/Reticent93/trap_house_b_and_b/pkg/config"
	"github.com/Reticent93/trap_house_b_and_b/pkg/handlers"
	"github.com/Reticent93/trap_house_b_and_b/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

var port = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change to true when in production
	app.InProduction = false
	
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	
	app.Session = session
	
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	
	render.NewTemplates(&app)
	
	fmt.Println(fmt.Sprintf("Starting on port %s", port))
	
	srv := &http.Server{Addr: port, Handler: routes(&app)}
	err = srv.ListenAndServe()
	log.Fatal(err)
	
}
