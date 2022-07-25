package ohlc

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type Data struct {
	TS uint
	O  decimal.Decimal
	H  decimal.Decimal
	L  decimal.Decimal
	C  decimal.Decimal
}

func Process(dsource, fpath string) ([]Data, error) {
	return nil, nil
}

func parse(fpath string) ([]Data, error) {
	var res []Data
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Error("failed to read ", fpath)
		return nil, err
	}
	var data [][]decimal.Decimal
	err = json.Unmarshal(bs, &data)
	if err != nil {
		log.Error("failed to parse ", fpath)
		return nil, err
	}
	for _, d := range data {
		r := Data{
			TS: uint(d[0].IntPart()),
			O:  d[1],
			H:  d[2],
			L:  d[3],
			C:  d[4],
		}
		res = append(res, r)
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
		if !strings.HasSuffix(info.Name(), ".ohlc") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files, err
}
