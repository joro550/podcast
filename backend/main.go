package main

import (
	"database/sql"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/lib/pq"
)

type Weather struct {
	Location    string `json:"location"`
	Temperature int    `json:"temperature1"`
}

func main() {
	val, found := os.LookupEnv("Connection_string")
	if !found {
		log.Fatal("Could not find connecton string")
	}

	log.Println("Connecting string!!!", val)

	db, err := sql.Open("postgres", val)
	if err != nil {
		log.Fatal("Could not connect to the database")
	}

	err = runMigrations(db)
	if err != nil {
		log.Fatal("Migrations could not be run", err)
	}

	log.Println("Starting server")

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

func runMigrations(db *sql.DB) error {
	err := filepath.WalkDir("migrations", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
