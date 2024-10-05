package main

import (
	"database/sql"
	_ "embed"
	"podcast-server/presenters"
	"podcast-server/takes"
)

// go:embed seed/takes.csv
var takesFile string

// go:embed seed/episodes.csv
var episodesFile string

func seedDatabase(db *sql.DB) error {
	presenterRepository := presenters.NewRepo(db)
	takesRespository := takes.NewRepo(db)

	presenter, err := initPresenter(presenterRepository)
	if err != nil {
		return err
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
