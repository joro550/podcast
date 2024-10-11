package main

import (
	"log"
	"net/http"
	"podcast-server/episodes"
	"podcast-server/presenters"
	"podcast-server/shared"
	"podcast-server/takes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	db, err := ConnectToDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	err = seedDatabase(db)
	if err != nil {
		log.Fatalln(err)
	}

	appServices := shared.NewAppService(db)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	presenters.AddPresenterEndpoints(r, &appServices)
	episodes.AddEpisodeEndpoints(r, &appServices)
	takes.AddTakeEndpoints(r, &appServices)

	r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.ListenAndServe(":3111", r)
}
