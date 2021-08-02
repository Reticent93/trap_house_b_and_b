package main

import (
	config2 "github.com/Reticent93/trap_house_b_and_b/internal/config"
	handlers2 "github.com/Reticent93/trap_house_b_and_b/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config2.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers2.Repo.Home)
	mux.Get("/about", handlers2.Repo.About)
	mux.Get("/dons-quarters", handlers2.Repo.Dons)
	mux.Get("/bastones-suite", handlers2.Repo.Bastones)

	mux.Get("/search-availability", handlers2.Repo.Availability)
	mux.Post("/search-availability", handlers2.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers2.Repo.AvailabilityJSON)

	mux.Get("/contact", handlers2.Repo.Contact)

	mux.Get("/make-reservation", handlers2.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
