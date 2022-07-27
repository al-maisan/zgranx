package ohlc

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Persist(db *sqlx.DB, dsource, period string, data []Data) error {
	log.Info(" #data = ", len(data))
	for _, d := range data {
		q := `
		INSERT IGNORE INTO ohlc(
			ts, base, quote, data_source_id, open, high, low, close, period)
		SELECT
			FROM_UNIXTIME(:ts), ba.id, qa.id, ds.id, :open, :high, :low, :close,
			'%s'
		FROM asset ba, asset qa, data_source ds
		WHERE ba.name = '%s' AND qa.name = '%s' AND ds.name = '%s'
		`
		mq := fmt.Sprintf(q, period, d.Base, d.Quote, dsource)
		for _, od := range d.Data {
			_, err := db.NamedExec(mq, od)
			if err != nil {
				log.Errorf("failed to insert or update: %s/%s @ %d", d.Base, d.Quote, od.TS)
				log.Error(err)
			}
		}
	}
	return nil
}
