package main

import (
	"crypto/sha256"
	"database/sql"
	_ "embed"
	"encoding/base64"
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
	log.Println("seeding takes")
	dbPresenters, err := initPresenter(presenterRepository)
	if err != nil {
		return err
	}
	log.Println("Got presenters from database")

	dbTakes, err := takesRespository.GetTakeSha()
	if err != nil {
		return err
	}
	log.Println("Got existing takes from database")

	stringReader := strings.NewReader(takesFile)
	reader := csv.NewReader(stringReader)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	_, err = reader.Read()
	if err != nil {
		log.Println("Read failed")
		return err
	}

	rows, err := reader.ReadAll()
	if err != nil {
		log.Println("Read failed")
		return err
	}

	for _, row := range rows {

		presenter := row[1]
		content := row[2]
		tags := row[3]
		createdDateString := row[4]
		dueDateString := row[5]
		wasCorrect := row[6]
		episode, err := strconv.Atoi(row[7])
		if err != nil {
			return err
		}

		layout := "2006-01-02"
		createDate, err := time.Parse(layout, createdDateString)
		if err != nil {
			return err
		}

		if len(dueDateString) == 0 {
			dueDateString = "2022-01-01"
		}

		dueDate, err := time.Parse(layout, dueDateString)
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
			EpisodeId:   episode,
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

		newTake.Sha = takeSha
		newTake.CompleteSha = completeSha

		if !exists {
			log.Println("Inserting new take", newTake)
			err = takesRespository.InsertTake(newTake)
			if err != nil {
				return err
			}

		} else if dbTake.CompleteSha != completeSha {

			log.Println("Updating take", newTake)
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
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

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

	return base64.URLEncoding.EncodeToString(h.Sum(nil)), nil
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
