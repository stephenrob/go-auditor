package main

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.lancs.ac.uk/library/auditor/pkg/events"
	"github.lancs.ac.uk/library/auditor/pkg/handler"
	"github.lancs.ac.uk/library/hutch"
	"log"
	"os"
	"os/signal"
)

func main() {
	fmt.Printf("Starting auditor handler\n")

	dbs := os.Getenv("DATABASE_CONNECTION_STRING")

	opt, err := pg.ParseURL(dbs)

	if err != nil {
		fmt.Printf("Failed to connect to Postgres DB: %s\n", err)
		return
	}

	db := pg.Connect(opt)

	repo := events.DBRepository{DB: db}

	done := make(chan os.Signal, 1)
	forever := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)
	signal.Notify(forever, os.Interrupt, os.Kill)
	rmq := os.Getenv("RABBITMQ_CONNECTION_STRING")
	c, err := hutch.NewClient(rmq, done)

	if err != nil {
		fmt.Printf("Error creating hutch client\n")
		return
	}

	h, err := handler.NewHutchHandler(c, "auditor", &repo)

	if err != nil {
		fmt.Printf("Error creating new hutch handler\n")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	err = h.Subscribe("auditor-persist", ctx)

	if err != nil {
		fmt.Printf("Error subscribing as auditor-persist: %s\n", err)
		return
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever

	cancel()

	err = h.Close()

	if err != nil {
		fmt.Printf("Failed to close client\n")
		return
	}

}
