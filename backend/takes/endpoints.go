package takes

import (
	"encoding/json"
	"log"
	"net/http"
	"podcast-server/presenters"
	"podcast-server/shared"

	"github.com/go-chi/chi/v5"
)

func AddTakeEndpoints(router *chi.Mux, servives *shared.AppServices) {
	repo := NewRepo(servives.Db)
	presenterRepo := presenters.NewRepo(servives.Db)
	router.Route("/takes", func(r chi.Router) {
		r.Get("/", getAllTakes(repo, presenterRepo))
	})
}

func getAllTakes(repo TakesRepository, presenterRepo presenters.PresenterRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		takes, err := repo.GetTakes()
		if err != nil {
			log.Println("Takes could not be retrieved", err)
			w.WriteHeader(500)
			return
		}

		presenters, err := presenterRepo.Get()
		if err != nil {
			log.Println("Takes could not be retrieved", err)
			w.WriteHeader(500)
			return
		}
		presenterMap := presentersToMap(presenters)

		reponses := []TakeRsponse{}
		for _, t := range takes {
			reponses = append(reponses, t.FromEntity(presenterMap))
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(reponses)
	}
}

func presentersToMap(model []presenters.Presenter) map[int]presenters.Presenter {
	presenterMap := make(map[int]presenters.Presenter)
	for _, m := range model {
		presenterMap[m.Id] = m
	}
	return presenterMap
}

type TakeRsponse struct {
	Content       string   `json:"content"`
	PresenterName string   `json:"presenterName"`
	Tags          []string `json:"tags"`
	Id            int      `json:"id"`
	EpisodeId     int      `json:"episodeId"`
	Result        int      `json:"result"`
}

func (resp *Take) FromEntity(presenterMap map[int]presenters.Presenter) TakeRsponse {
	presenter, ok := presenterMap[resp.PresenterId]
	name := ""
	if ok {
		name = presenter.Name
	}

	return TakeRsponse{
		Content:       resp.Content,
		PresenterName: name,
		Tags:          resp.Tags,
		Id:            resp.Id,
		EpisodeId:     resp.EpisodeId,
	}
}
