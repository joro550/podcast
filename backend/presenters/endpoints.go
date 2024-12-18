package presenters

import (
	"encoding/json"
	"log"
	"net/http"
	"podcast-server/shared"

	"github.com/go-chi/chi/v5"
)

func AddPresenterEndpoints(router *chi.Mux, services *shared.AppServices) {
	repo := NewRepo(services.Db)

	router.Route("/presenters", func(r chi.Router) {
		r.Get("/", getPresenters(repo))
	})
}

func getPresenters(repo PresenterRepository) http.HandlerFunc {
	log.Println("setting up getting presenters")

	return func(w http.ResponseWriter, r *http.Request) {
		presenters, err := repo.Get()
		if err != nil {
			log.Println("Could not get presenters", err)
			w.WriteHeader(500)
			return
		}

		responses := []PresenterResponse{}
		for _, p := range presenters {
			responses = append(responses, p.FromEntity())
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(responses)
	}
}

type PresenterResponse struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	ImageUrl    string           `json:"imageUrl"`
	AltText     string           `json:"altText"`
	Socials     []SocialResponse `json:"socials"`
	Id          int              `json:"id"`
}

type SocialResponse struct {
	Username string `json:"username"`
	Url      string `json:"url"`
	Icon     string `json:"icon"`
}

func (resp *Presenter) FromEntity() PresenterResponse {
	socials := []SocialResponse{}
	for _, social := range resp.Socials {
		socials = append(socials, social.FromEntity())
	}

	return PresenterResponse{
		Id:          resp.Id,
		Name:        resp.Name,
		Description: resp.Description,
		ImageUrl:    resp.ImageUrl,
		AltText:     resp.AltText,
		Socials:     socials,
	}
}

func (resp *Social) FromEntity() SocialResponse {
	return SocialResponse{
		Username: resp.Username,
		Url:      resp.Url,
	}
}
