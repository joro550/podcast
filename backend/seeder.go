package main

import (
	"crypto/sha256"
	"database/sql"
	_ "embed"
	"encoding/csv"
	"encoding/json"
	"log"
	"podcast-server/episodes"
	"podcast-server/presenters"
	"podcast-server/takes"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed seed/takes.csv
var takesFile string

//go:embed seed/episodes.csv
var episodesFile string

func seedDatabase(db *sql.DB) error {
	presenterRepository := presenters.NewRepo(db)
	takesRespository := takes.NewRepo(db)
	episodeRepo := episodes.NewRepo(db)

	err := createEpisodes(episodeRepo)
	if err != nil {
		return err
	}

	err = seedTakes(presenterRepository, takesRespository)
	if err != nil {
		return err
	}

	return nil
}

func seedTakes(presenterRepository presenters.PresenterRepository, takesRespository takes.TakesRepository) error {
	dbPresenters, err := initPresenter(presenterRepository)
	if err != nil {
		return err
	}

	dbTakes, err := takesRespository.GetTakeSha()
	if err != nil {
		return err
	}

	stringReader := strings.NewReader(takesFile)
	reader := csv.NewReader(stringReader)

	reader.Read()
	rows, err := reader.ReadAll()
	for _, row := range rows {
		presenter := row[1]
		content := row[2]
		tags := row[3]
		createdDateString := row[4]
		dueDateString := row[5]
		wasCorrect := row[6]

		createDate, err := time.Parse("2023-01-01", createdDateString)
		if err != nil {
			return err
		}

		dueDate, err := time.Parse("2023-01-01", dueDateString)
		if err != nil {
			return err
		}

		tagsSplit := strings.Split(",", tags)

		var presenterId int
		switch strings.ToLower(presenter) {
		case "thorin":
			presenterId = dbPresenters.ThorinId

		case "kassad":
			presenterId = dbPresenters.KassadId

		case "mauisnake":
			presenterId = dbPresenters.MauisnakeId

		}

		newTake := takes.Take{
			Content:     content,
			Tags:        tagsSplit,
			PresenterId: presenterId,
		}

		takeSha, err := shaTake(newTake)
		if err != nil {
			return err
		}

		dbTake, exists := dbTakes[takeSha]

		newTake.CreatedDate = createDate
		newTake.DueDate = dueDate
		newTake.WasCorrect = wasCorrect

		completeSha, err := shaTake(newTake)
		if err != nil {
			return err
		}

		if !exists {
			takesRespository.InsertTake(newTake)
		} else if dbTake.CompleteSha != completeSha {
			takesRespository.UpdateTake(newTake)
		}

	}
	return nil
}

func createEpisodes(episodeRepo episodes.EpisodesRespository) error {
	dbEpisodes, err := episodeRepo.GetNames()
	if err != nil {
		return err
	}

	log.Println("Got episodes from database")

	stringReader := strings.NewReader(episodesFile)
	reader := csv.NewReader(stringReader)

	reader.Read()
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, row := range rows {
		episodeId, err := strconv.Atoi(row[1])
		if err != nil {
			return err
		}
		episode := episodes.Episode{
			Name:      row[0],
			EpisodeId: episodeId,
		}

		if slices.Contains(dbEpisodes, episode.Name) {
			continue
		}

		err = episodeRepo.Insert(episode)
		if err != nil {
			return err
		}
		log.Println("Inserted episode ", episode.Name)
	}
	return nil
}

type presenter struct {
	ThorinId    int
	KassadId    int
	MauisnakeId int
}

func shaTake(take takes.Take) (string, error) {
	takeJson, err := json.Marshal(take)
	if err != nil {
		return "", err
	}

	h := sha256.New()
	h.Write(takeJson)
	shaCode := h.Sum(nil)
	return string(shaCode), nil
}

func initPresenter(repo presenters.PresenterRepository) (presenter, error) {
	thorin, err := repo.GetPresenter("Thorin")
	if err != nil {
		return presenter{}, err
	}

	kassad, err := repo.GetPresenter("Kassad")
	if err != nil {
		return presenter{}, err
	}

	maui, err := repo.GetPresenter("mauisnake")
	if err != nil {
		return presenter{}, err
	}
	return presenter{
		ThorinId:    thorin.Id,
		KassadId:    kassad.Id,
		MauisnakeId: maui.Id,
	}, nil
}
