package takes

import (
	"database/sql"
	"encoding/json"
	"time"
)

type TakesRepository struct {
	db *sql.DB
}

type Take struct {
	CreatedDate   time.Time
	DueDate       time.Time
	Content       string
	PresenterName string
	Tags          []string
	PresenterId   int
	Id            int
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
	defer rows.Close()

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

func (db *TakesRepository) InsertTake(mode Take) error {
	result, err := json.Marshal(mode.Tags)
	if err != nil {
		return err
	}

	_, err = db.db.Exec(`insert into hot_take (content, presenter, tags, episode_id, result, created_date, due_date)
        VALUES (?, ?, ?, ?, ?, ?, ?)`,
		mode.Content,
		mode.PresenterId,
		string(result),
		mode.EpisodeId,
		mode.Result,
		mode.CreatedDate,
		mode.DueDate,
	)

	return err
}
