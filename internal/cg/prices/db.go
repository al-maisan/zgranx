package prices

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Persist(db *sqlx.DB, dsource, period string, data []Multi) error {
	log.Info(" #data = ", len(data))
	for _, d := range data {
		for _, s := range d.Data {
			q := `
				INSERT IGNORE INTO price(
					ts, base, quote, data_source_id, price, q_volume,
					q_volume_change, period)
				SELECT
					FROM_UNIXTIME(%d), ba.id, qa.id, ds.id, :price, :q_volume,
					:q_volume_change, '%s'
				FROM asset ba, asset qa, data_source ds
				WHERE ba.name = '%s' AND qa.name = '%s' AND ds.name = '%s'
				`
			mq := fmt.Sprintf(q, d.TS, period, d.Base, s.Quote, dsource)
			_, err := db.NamedExec(mq, s)
			if err != nil {
				log.Errorf("failed to insert or update: %s/%s @ %d", d.Base, s.Quote, d.TS)
				log.Error(err)
			}
		}
	}
	return nil
}
