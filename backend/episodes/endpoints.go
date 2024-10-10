package episodes

import (
	"encoding/json"
	"log"
	"net/http"
	"podcast-server/shared"

	"github.com/go-chi/chi/v5"
)

func AddEpisodeEndpoints(router *chi.Mux, services *shared.AppServices) {
	repo := NewRepo(services.Db)
	router.Route("/episodes", func(r chi.Router) {
		r.Get("/", getAllEpisodes(repo))
	})
}

func getAllEpisodes(repo EpisodesRespository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		episodes, err := repo.Get()
		if err != nil {
			log.Println("Episodes could not be retrieved", err)
			w.WriteHeader(500)
			return
		}

		responses := []EpisodeResponse{}

		for _, m := range episodes {
			responses = append(responses, m.FromEntity())
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(responses)
	}
}

type EpisodeResponse struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func (model Episode) FromEntity() EpisodeResponse {
	return EpisodeResponse{
		Id:   model.Id,
		Name: model.Name,
	}
}
