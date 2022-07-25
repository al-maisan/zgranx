package ohlc

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type Data struct {
	Base  string
	Quote string
	Data  []OHLC
}

type OHLC struct {
	TS uint
	O  decimal.Decimal
	H  decimal.Decimal
	L  decimal.Decimal
	C  decimal.Decimal
}

func Process(fpath string) ([]Data, error) {
	var data []Data
	files, err := find(fpath)
	if err != nil {
		log.Error("failed to find ohlc files ", fpath)
		return nil, err
	}

	for _, file := range files {
		od, err := parse(file)
		if err != nil {
			log.Error("failed to parse ohlc file ", file)
			continue
		}
		b, q := tradingPair(file)
		if b != "" && q != "" {
			d := Data{
				Base:  b,
				Quote: q,
				Data:  od,
			}
			data = append(data, d)
		}
	}
	return data, nil
}

func tradingPair(fpath string) (string, string) {
	pair := strings.Split(strings.Split(path.Base(fpath), ".")[0], "_")
	if len(pair) == 2 {
		return pair[0], pair[1]
	}
	return "", ""
}

func parse(fpath string) ([]OHLC, error) {
	var res []OHLC
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
		r := OHLC{
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
