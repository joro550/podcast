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
	Sha           string
	CompleteSha   string
	WasCorrect    string
	Tags          []string
	PresenterId   int
	Id            int
	EpisodeId     int
	Result        int
}

type TakeSha struct {
	Sha         string
	CompleteSha string
	Id          int
}

func NewRepo(db *sql.DB) TakesRepository {
	return TakesRepository{db: db}
}

func (db *TakesRepository) GetTakes() ([]Take, error) {
	rows, err := db.db.Query(`select id, presenter, episode, content, tags, result, created_date, due_date, sha, complete_sha
        from hot_take ht`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	takes := []Take{}

	for rows.Next() {
		var take Take
		var tags string

		err := rows.Scan(
			&take.Id,
			&take.PresenterId,
			&take.EpisodeId,
			&take.Content,
			&tags,
			&take.Result,
			&take.CreatedDate,
			&take.DueDate,
			&take.Sha,
			&take.CompleteSha,
		)
		if err != nil {
			return takes, err
		}

		err = json.Unmarshal([]byte(tags), &take.Tags)
		if err != nil {
			return takes, err
		}
		takes = append(takes, take)
	}

	return takes, nil
}

func (db *TakesRepository) GetTakeSha() (map[string]TakeSha, error) {
	rows, err := db.db.Query(`select sha, complete_sha, id from hot_take ht`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	takes := map[string]TakeSha{}

	for rows.Next() {
		var takeSha TakeSha
		err := rows.Scan(
			&takeSha.Sha,
			&takeSha.CompleteSha,
			&takeSha.Id,
		)
		if err != nil {
			return nil, err
		}

		takes[takeSha.Sha] = takeSha
	}

	return takes, nil
}

func (db *TakesRepository) InsertTake(mode Take) error {
	result, err := json.Marshal(mode.Tags)
	if err != nil {
		return err
	}

	_, err = db.db.Exec(`insert into hot_take (content, presenter, tags, episode, result, created_date, due_date, sha, complete_sha, wascorrect)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		mode.Content,
		mode.PresenterId,
		string(result),
		mode.EpisodeId,
		mode.Result,
		mode.CreatedDate,
		mode.DueDate,
		mode.Sha,
		mode.CompleteSha,
		mode.WasCorrect,
	)

	return err
}

func (db *TakesRepository) UpdateTake(mode Take) error {
	_, err := db.db.Exec(`UPDATE hot_take SET created_date = $1, due_date = $2, was_correct = $3 where id = $4`,
		mode.CreatedDate,
		mode.DueDate,
		mode.WasCorrect,
		mode.Id,
	)

	return err
}
