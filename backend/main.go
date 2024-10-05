package main

import (
	"encoding/json"
	"log"
	"net/http"
	"podcast-server/presenters"
	"podcast-server/shared"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Weather struct {
	Location    string `json:"location"`
	Id          int    `json:"id"`
	Temperature int    `json:"temperature"`
}

func main() {
	db, err := ConnectToDatabase()
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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello mark"))
	})

	r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	r.Get("/weather", func(w http.ResponseWriter, r *http.Request) {
		weather := Weather{Id: 1, Location: "London"}

		weathers := []Weather{
			weather,
		}

		encoder := json.NewEncoder(w)
		err := encoder.Encode(weathers)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":3111", r)
}
