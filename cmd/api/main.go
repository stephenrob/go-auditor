package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	api2 "github.lancs.ac.uk/library/auditor/pkg/api"
	"github.lancs.ac.uk/library/auditor/pkg/events"
	"os"
)

func main() {
	fmt.Printf("Starting auditor API\n")

	dbs := os.Getenv("DATABASE_CONNECTION_STRING")

	opt, err := pg.ParseURL(dbs)

	if err != nil {
		fmt.Printf("Failed to connect to Postgres DB: %s\n", err)
		return
	}

	db := pg.Connect(opt)

	repo := events.DBRepository{DB: db}

	api := api2.NewAPIService(&repo)

	api.Run(":8010")
}
