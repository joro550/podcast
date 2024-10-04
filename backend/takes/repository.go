package takes

import (
	"database/sql"
	"encoding/json"
)

type TakesRepository struct {
	db *sql.DB
}

type Take struct {
	Content       string
	Tags          []string
	Id            int
	PresenterName string
	EpisodeId     int
	Result        int
}

func NewRepo(db *sql.DB) TakesRepository {
	return TakesRepository{db: db}
}

func (db *TakesRepository) GetTakes() ([]Take, error) {
	rows, err := db.db.Query(`select id, p.name, episode, content, tags, result, created_date, due_date
        from hot_take ht
        inner join presenter p on ht.presenter = p.id`)
	if err != nil {
		return nil, err
	}

	takes := []Take{}

	for rows.Next() {
		var take Take
		var tags string
		rows.Scan(
			&take.Id,
			&take.PresenterName,
			&take.EpisodeId,
			&take.Content,
			&tags,
			&take.Result,
		)
		err := json.Unmarshal([]byte(tags), &take.Tags)
		if err != nil {
			return takes, err
		}
		takes = append(takes, take)
	}

	return takes, nil
}
