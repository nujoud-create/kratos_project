package data

import "database/sql"

type Data struct {
	db *sql.DB
}

func NewData(db *sql.DB) (*Data, func(), error) {
	cleanup := func() {
		db.Close()
	}

	return &Data{
		db: db,
	}, cleanup, nil
}
