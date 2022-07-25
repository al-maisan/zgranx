package main

import (
	"fmt"
	"os"

	"github.com/alphabot-fi/T-801/internal/cg/ohlc"
	"github.com/alphabot-fi/T-801/internal/cg/prices"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	bts, rev, version string
	log               = logrus.New()
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}
	version = fmt.Sprintf("%s::%s", bts, rev)
	log.Info("version = ", version)

	dsn := getDSN()
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var dsource, fpath, period string
	app := &cli.App{
		Name:  "dit",
		Usage: "data import tool",
		Commands: []*cli.Command{
			{
				Name:    "process-ohlc-data",
				Aliases: []string{"pod"},
				Usage:   "process a collection of ohlc files",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "dsource",
						Usage:       "data source (e.g. \"coingecko\")",
						Required:    true,
						Destination: &dsource,
					},
					&cli.StringFlag{
						Name:        "period",
						Usage:       "data collection period (e.g. \"30M\")",
						Required:    true,
						Destination: &period,
					},
					&cli.StringFlag{
						Name:        "fpath",
						Usage:       "root directory that contains the ohlc files",
						Required:    true,
						Destination: &fpath,
					},
				},
				Action: func(c *cli.Context) error {
					log.Info("fpath = ", fpath)
					data, err := ohlc.Process(fpath)
					if err != nil {
						return err
					}
					err = ohlc.Persist(db, dsource, period, data)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "process-price-data",
				Aliases: []string{"ppd"},
				Usage:   "process a collection of prices.json files",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "dsource",
						Usage:       "data source (e.g. \"coingecko\")",
						Required:    true,
						Destination: &dsource,
					},
					&cli.StringFlag{
						Name:        "period",
						Usage:       "data collection period (e.g. \"5M\")",
						Required:    true,
						Destination: &period,
					},
					&cli.StringFlag{
						Name:        "fpath",
						Usage:       "root directory that contains the ohlc files",
						Required:    true,
						Destination: &fpath,
					},
				},
				Action: func(c *cli.Context) error {
					log.Info("fpath = ", fpath)
					data, err := prices.Process(fpath)
					if err != nil {
						return err
					}
					err = prices.Persist(db, dsource, period, data)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getDSN() string {
	var (
		host, port, user, passwd, database string
		present                            bool
	)

	host, present = os.LookupEnv("T_801_PDB_HOST")
	if !present {
		log.Fatal("T_801_PDB_HOST variable not set")
	}
	port, present = os.LookupEnv("T_801_PDB_PORT")
	if !present {
		log.Fatal("T_801_PDB_PORT variable not set")
	}
	user, present = os.LookupEnv("T_801_PDB_USER")
	if !present {
		log.Fatal("T_801_PDB_USER variable not set")
	}
	passwd, present = os.LookupEnv("T_801_PDB_PASSWORD")
	if !present {
		log.Fatal("T_801_PDB_PASSWORD variable not set")
	}
	database, present = os.LookupEnv("T_801_PDB_DATABASE")
	if !present {
		log.Fatal("T_801_PDB_DATABASE variable not set")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true&time_zone=UTC", user, passwd, host, port, database)
	return dsn
}
