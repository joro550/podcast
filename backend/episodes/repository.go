package episodes

import "database/sql"

type EpisodesRespository struct {
	db *sql.DB
}

type Episode struct {
	Name      string
	Id        int
	EpisodeId int
}

func NewRepo(db *sql.DB) EpisodesRespository {
	return EpisodesRespository{db: db}
}

func (repo *EpisodesRespository) Get() ([]Episode, error) {
	rows, err := repo.db.Query("select id, name, episode_id from episode")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	episodes := []Episode{}

	for rows.Next() {
		var episode Episode

		err := rows.Scan(&episode.Id, &episode.Name, &episode.EpisodeId)
		if err != nil {
			return nil, err
		}

		episodes = append(episodes, episode)
	}

	return episodes, nil
}

func (repo *EpisodesRespository) GetNames() ([]string, error) {
	rows, err := repo.db.Query("select id, name, episode_id from episode")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	episodes := []string{}

	for rows.Next() {
		var episode Episode

		err := rows.Scan(&episode.Id, &episode.Name, &episode.EpisodeId)
		if err != nil {
			return nil, err
		}

		episodes = append(episodes, episode.Name)
	}

	return episodes, nil
}

func (repo *EpisodesRespository) Insert(model Episode) error {
	_, err := repo.db.Exec(`insert into episode (name, episode_id)  Values ($1, $2)`,
		model.Name,
		model.EpisodeId,
	)

	return err
}
