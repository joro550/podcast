package main

import (
	"encoding/json"
	"log"
	"net/http"
	"podcast-server/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/lib/pq"
)

type Weather struct {
	Location    string `json:"location"`
	Temperature int    `json:"temperature1"`
}

func main() {
	_, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello mark"))
	})

	r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	r.Get("/weather", func(w http.ResponseWriter, r *http.Request) {
		weather := Weather{Location: "London"}

		encoder := json.NewEncoder(w)
		err := encoder.Encode(weather)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":3111", r)
}
