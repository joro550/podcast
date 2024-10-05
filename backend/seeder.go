package main

import (
	"database/sql"
	_ "embed"
	"encoding/csv"
	"podcast-server/episodes"
	"podcast-server/presenters"
	"podcast-server/takes"
	"slices"
	"strconv"
	"strings"
)

// go:embed seed/takes.csv
var takesFile string

// go:embed seed/episodes.csv
var episodesFile string

func seedDatabase(db *sql.DB) error {
	presenterRepository := presenters.NewRepo(db)
	takesRespository := takes.NewRepo(db)
	episodeRepo := episodes.NewRepo(db)

	presenter, err := initPresenter(presenterRepository)
	if err != nil {
		return err
	}

	dbEpisodes, err := episodeRepo.GetNames()
	if err != nil {
		return err
	}

	stringReader := strings.NewReader(episodesFile)
	reader := csv.NewReader(stringReader)

	// read the header line
	reader.Read()
	rows, err := reader.ReadAll()

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

		episodeRepo.Insert(episode)

	}

	takesRespository.InsertTake(takes.Take{
		Content:     "",
		PresenterId: presenter.KassadId,
	})

	return nil
}

type presenter struct {
	ThorinId    int
	KassadId    int
	MauisnakeId int
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
