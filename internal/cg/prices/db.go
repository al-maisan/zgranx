package prices

import (
	"github.com/jmoiron/sqlx"
)

func Persist(db *sqlx.DB, dsource, period string, data []Multi) error {
	log.Info(" #data = ", len(data))
	return nil
}
