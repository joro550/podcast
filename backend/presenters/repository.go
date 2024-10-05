package presenters

import "database/sql"

type PresenterRepository struct {
	db *sql.DB
}

type Presenter struct {
	Name string
	Id   int
}

func NewRepo(db *sql.DB) PresenterRepository {
	return PresenterRepository{db: db}
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
	rows := repo.db.QueryRow("select id, name from presenters where name = ?", name)

	var presenter Presenter
	err := rows.Scan(&presenter.Id, &presenter.Name)
	if err != nil {
		return Presenter{}, err
	}

	return presenter, nil
}
