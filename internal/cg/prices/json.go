package prices

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type Multi struct {
	Base string `db:"base"`
	TS   uint   `db:"ts"`
	Data []Single
}

type Single struct {
	Quote      string          `db:"quote"`
	Price      decimal.Decimal `db:"price"`
	QVol       decimal.Decimal `db:"q_volume"`
	QVolChange decimal.Decimal `db:"q_volume_change"`
}

func Process(fpath string) ([]Multi, error) {
	var data []Multi
	files, err := find(fpath)
	if err != nil {
		log.Error("failed to find json files with price data", fpath)
		return nil, err
	}

	log.Info(" #files = ", len(files))
	for _, file := range files {
		pd, err := parse(file)
		if err != nil {
			log.Error("failed to parse json file with price data ", file)
			continue
		}
		data = append(data, pd...)
	}
	return data, nil
}

func parse(fpath string) ([]Multi, error) {
	fiat := []string{"usd", "eur", "jpy", "chf", "cad", "krw"}
	var res []Multi
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Error("failed to read ", fpath)
		return nil, err
	}
	var data map[string]map[string]decimal.Decimal
	err = json.Unmarshal(bs, &data)
	if err != nil {
		log.Error("failed to parse ", fpath)
		return nil, err
	}
	for base, qd := range data {
		m := Multi{
			Base: base,
			TS:   uint(qd["last_updated_at"].IntPart()),
		}
		for _, fc := range fiat {
			_, ok := qd[fc]
			if !ok {
				log.Warn("missing fiat: ", fc)
				continue
			}
			s := Single{
				Quote:      fc,
				Price:      qd[fc],
				QVol:       qd[fc+"_24h_vol"],
				QVolChange: qd[fc+"_24h_change"],
			}
			m.Data = append(m.Data, s)
		}
		res = append(res, m)
	}
	return res, nil
}

func find(fpath string) ([]string, error) {
	var files []string
	err := filepath.Walk(fpath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			// best effort -- ignore files/dirs we cannot read
			return nil
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		if info.Name() != "prices.json" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files, err
}
