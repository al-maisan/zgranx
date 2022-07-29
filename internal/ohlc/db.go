package ohlc

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Persist(db *sqlx.DB, dsource, period string, data []Data) error {
	log.Infof(" >> dsource = '%s' * period = '%s' * #data = %d", dsource, period, len(data))
	for _, d := range data {
		assetColumn := "name"
		if dsource != "coingecko" {
			assetColumn = "symbol"
		}
		q := `
		INSERT IGNORE INTO ohlc(
			ts, base, quote, data_source_id, open, high, low, close, period,
			count, q_volume)
		SELECT
			FROM_UNIXTIME(:ts), ba.id, qa.id, ds.id, :open, :high, :low, :close,
			'%s', :count, :q_volume
		FROM asset ba, asset qa, data_source ds
		WHERE ba.%s = '%s' AND qa.%s = '%s' AND ds.name = '%s'
		`
		mq := fmt.Sprintf(q, period, assetColumn, d.Base, assetColumn, d.Quote, dsource)
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
