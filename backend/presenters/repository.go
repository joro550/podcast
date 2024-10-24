package presenters

import (
	"database/sql"
	"encoding/json"
	"strings"
)

type PresenterRepository struct {
	db *sql.DB
}

type Presenter struct {
	Name        string
	Description string
	ImageUrl    string
	AltText     string
	Socials     []Social
	Id          int
	Sha         string
}

type Social struct {
	Username string
	Url      string
}

type SlimPresenter struct {
	Id   int
	Name string
	Sha  string
}

func NewRepo(db *sql.DB) PresenterRepository {
	return PresenterRepository{db: db}
}

func (repo *PresenterRepository) GetNames() (map[string]SlimPresenter, error) {
	rows, err := repo.db.Query("select name, sha, id from presenters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	presenters := make(map[string]SlimPresenter)

	for rows.Next() {
		presenter := SlimPresenter{}

		err := rows.Scan(&presenter.Name, &presenter.Sha, &presenter.Id)
		if err != nil {
			return nil, err
		}

		presenters[strings.ToLower(presenter.Name)] = presenter
	}

	return presenters, nil
}

func (repo *PresenterRepository) Get() ([]Presenter, error) {
	rows, err := repo.db.Query("select id, name from presenters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	presenters := []Presenter{}

	for rows.Next() {
		var presenter Presenter

		err := rows.Scan(&presenter.Id, &presenter.Name)
		if err != nil {
			return nil, err
		}

		presenters = append(presenters, presenter)
	}

	return presenters, nil
}

func (repo *PresenterRepository) GetPresenter(name string) (Presenter, error) {
	rows := repo.db.QueryRow("select id, name from presenters where name = $1", name)

	var presenter Presenter
	err := rows.Scan(&presenter.Id, &presenter.Name)
	if err != nil {
		return Presenter{}, err
	}

	return presenter, nil
}

func (repo *PresenterRepository) Insert(model Presenter) error {
	socials, err := json.Marshal(model.Socials)
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(`insert into presenters (name, description, image_url, alt_text, socials, sha)
        values ($1, $2, $3, $4, $5, $6)`,
		model.Name,
		model.Description,
		model.ImageUrl,
		model.AltText,
		socials,
		model.Sha,
	)

	return err
}

func (repo *PresenterRepository) Update(id int, model Presenter) error {
	socials, err := json.Marshal(model.Socials)
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(`UPDATE presenters SET
        name=$1, description=$2, image_url=$3, alt_text=$4, socials=$5, sha=$6
        where id = $7`,
		model.Name,
		model.Description,
		model.ImageUrl,
		model.AltText,
		socials,
		model.Sha,
		id,
	)

	return err
}
