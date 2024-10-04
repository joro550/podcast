package shared

import "database/sql"

type AppServices struct {
	Db *sql.DB
}

func NewAppService(db *sql.DB) AppServices {
	return AppServices{Db: db}
}
