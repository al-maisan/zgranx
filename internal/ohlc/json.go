package ohlc

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path"
	"path/filepath"
	"sort"
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
	TS    uint            `db:"ts"`
	O     decimal.Decimal `db:"open"`
	H     decimal.Decimal `db:"high"`
	L     decimal.Decimal `db:"low"`
	C     decimal.Decimal `db:"close"`
	Count uint            `db:"count"`
	QVol  decimal.Decimal `db:"q_volume"`
}

func Process(dsource, fpath string) ([]Data, error) {
	var data []Data
	files, err := find(fpath)
	if err != nil {
		log.Error("failed to find ohlc files ", fpath)
		return nil, err
	}

	log.Info(" #files = ", len(files))
	for _, file := range files {
		var (
			od  []OHLC
			err error
		)
		switch dsource {
		case "coingecko":
			od, err = coingeckoParse(file)
		case "huobi":
			od, err = huobiParse(file)
		case "binance":
			od, err = binanceParse(file)
		case "gateio":
			od, err = gateioParse(file)
		case "ftx":
			od, err = ftxParse(file)
		default:
			log.Errorf("unsupported data source: '%s'", dsource)
			continue
		}
		if err != nil {
			log.Error("failed to parse ohlc file ", file)
			continue
		}
		b, q := tradingPair(dsource, file)
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

func tradingPair(dsource, fpath string) (string, string) {
	if dsource == "huobi" {
		// file names all end on "usdt.ohlc"
		bn := strings.Split(path.Base(fpath), ".")[0]
		return strings.TrimSuffix(bn, "usdt"), "usdt"
	} else if dsource == "binance" {
		// file names all end on "USDT.ohlc"
		bn := strings.Split(path.Base(fpath), ".")[0]
		return strings.ToLower(strings.TrimSuffix(bn, "USDT")), "usdt"
	} else if dsource == "gateio" {
		pair := strings.Split(strings.Split(path.Base(fpath), ".")[0], "_")
		if len(pair) == 2 {
			return strings.ToLower(pair[0]), strings.ToLower(pair[1])
		}
	} else if dsource == "ftx" {
		pair := strings.Split(strings.Split(path.Base(fpath), ".")[0], "-")
		if len(pair) == 2 {
			return strings.ToLower(pair[0]), strings.ToLower(pair[1])
		}
	} else if dsource == "coingecko" {
		pair := strings.Split(strings.Split(path.Base(fpath), ".")[0], "_")
		if len(pair) == 2 {
			return pair[0], pair[1]
		}
	} else {
		log.Errorf("unsupported data source: '%s'", dsource)
		return "", ""
	}
	return "", ""
}

func coingeckoParse(fpath string) ([]OHLC, error) {
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
			// we want seconds
			TS:   uint(d[0].IntPart()) / 1e3,
			O:    d[1],
			H:    d[2],
			L:    d[3],
			C:    d[4],
			QVol: decimal.NewFromFloat(0.0),
		}
		res = append(res, r)
	}
	return res, nil
}

func huobiParse(fpath string) ([]OHLC, error) {
	type HK struct {
		TS    uint            `json:"id"`
		O     decimal.Decimal `json:"open"`
		H     decimal.Decimal `json:"high"`
		L     decimal.Decimal `json:"low"`
		C     decimal.Decimal `json:"close"`
		Count uint            `json:"count"`
		QVol  decimal.Decimal `json:"vol"`
	}
	type HKD struct {
		Status string `json:"status"`
		Data   []HK   `json:"data"`
	}
	var res []OHLC
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Error("failed to read ", fpath)
		return nil, err
	}
	var tl HKD
	err = json.Unmarshal(bs, &tl)
	if err != nil {
		log.Error("failed to parse ", fpath)
		return nil, err
	}
	if tl.Status == "ok" {
		for _, d := range tl.Data {
			r := OHLC(d)
			res = append(res, r)
		}
	} else {
		err = fmt.Errorf("status not `ok` for ohlc data file: '%s'", fpath)
		log.Error(err)
		return nil, err
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
	sort.Strings(files)
	return files, err
}

func binanceParse(fpath string) ([]OHLC, error) {
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
			// we want seconds
			TS:    uint(d[0].IntPart()) / 1e3,
			O:     d[1],
			H:     d[2],
			L:     d[3],
			C:     d[4],
			Count: uint(d[8].IntPart()),
			QVol:  d[7],
		}
		res = append(res, r)
	}
	return res, nil
}

func gateioParse(fpath string) ([]OHLC, error) {
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
			// we want seconds
			TS:   uint(d[0].IntPart()),
			O:    d[5],
			H:    d[3],
			L:    d[4],
			C:    d[2],
			QVol: d[1],
		}
		res = append(res, r)
	}
	return res, nil
}

func ftxParse(fpath string) ([]OHLC, error) {
	type HK struct {
		TS   decimal.Decimal `json:"time"`
		O    decimal.Decimal `json:"open"`
		H    decimal.Decimal `json:"high"`
		L    decimal.Decimal `json:"low"`
		C    decimal.Decimal `json:"close"`
		QVol decimal.Decimal `json:"volume"`
	}
	type HKD struct {
		Success bool `json:"success"`
		Data    []HK `json:"result"`
	}
	var res []OHLC
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Error("failed to read ", fpath)
		return nil, err
	}
	var tl HKD
	err = json.Unmarshal(bs, &tl)
	if err != nil {
		log.Errorf("failed to parse %s, %v", fpath, err)
		return nil, err
	}
	if tl.Success {
		for _, d := range tl.Data {
			r := OHLC{
				TS:   uint(d.TS.IntPart() / 1e3),
				O:    d.O,
				H:    d.H,
				L:    d.L,
				C:    d.C,
				QVol: d.QVol,
			}
			res = append(res, r)
		}
	} else {
		err = fmt.Errorf("success not `true` for ohlc data file: '%s'", fpath)
		log.Error(err)
		return nil, err
	}
	return res, nil
}
