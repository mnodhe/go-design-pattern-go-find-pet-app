package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second)) // for nasty people that trying to add
	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("../../static"))))
	// display test page
	mux.Get("/test-patterns", app.TestPatterns)
	mux.Get("/api/dog-from-factory", app.CreateDogFromFactory)
	mux.Get("/api/cat-from-factory", app.CreateCatFromFactory)

	mux.Get("/api/cat-from-abstract-factory", app.CreateCatFromAbstractFactory)
	mux.Get("/api/dog-from-abstract-factory", app.CreateDogFromAbstractFactory)

	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)
	mux.Get("/api/dog-breeds", app.GetAllDogBreedsJSON)
	return mux
}
