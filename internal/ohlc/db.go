package ohlc

import "github.com/jmoiron/sqlx"

func Persist(db *sqlx.DB, dsource string, data []Data) error {
	return nil
}
