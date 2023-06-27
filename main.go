package main

import (
	"context"
	"github.com/nats-io/stan.go"
	router "l0/internal/app/http"
	"l0/internal/app/service"
	"l0/internal/app/subscription"
	"l0/internal/config"
	"l0/internal/memory"
	"l0/internal/repository/postgres"
	"log"
)

func main() {
	cfg := config.Read()

	ctx := context.Background()

	natsStream, err := stan.Connect(cfg.Nats.Cluster, cfg.Nats.Client, stan.NatsURL(cfg.Nats.Url))
	if err != nil {
		panic(err)
	}
	defer natsStream.Close()

	db := postgres.New(cfg.Postgres)

	cashe, err := memory.New(ctx, db)
	if err != nil {
		log.Println("Data could not be recovered, error:", err)
	}

	service := service.New(db, cashe)
	worker := subscription.New(cashe, db)

	worker.Start(ctx, natsStream)

	r := router.NewServer(service)
	err = r.Listen(":" + cfg.Service.Port)
	if err != nil {
		panic(err)
	}
}
