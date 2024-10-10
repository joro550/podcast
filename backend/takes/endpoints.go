package takes

import (
	"encoding/json"
	"log"
	"net/http"
	"podcast-server/shared"

	"github.com/go-chi/chi/v5"
)

func AddTakeEndpoints(router *chi.Mux, servives *shared.AppServices) {
	repo := NewRepo(servives.Db)
	router.Route("/takes", func(r chi.Router) {
		r.Get("/", getAllTakes(repo))
	})
}

func getAllTakes(repo TakesRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		takes, err := repo.GetTakes()
		if err != nil {
			log.Println("Takes could not be retrieved", err)
			w.WriteHeader(500)
			return
		}

		reponses := []TakeRsponse{}
		for _, t := range takes {
			reponses = append(reponses, t.FromEntity())
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(reponses)
	}
}

type TakeRsponse struct {
	Content       string
	PresenterName string
	Tags          []string
	Id            int
	EpisodeId     int
	Result        int
}

func (resp *Take) FromEntity() TakeRsponse {
	return TakeRsponse{
		Content:       resp.Content,
		PresenterName: resp.PresenterName,
		Tags:          resp.Tags,
		Id:            resp.Id,
		EpisodeId:     resp.EpisodeId,
	}
}
