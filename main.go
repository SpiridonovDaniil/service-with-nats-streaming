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

	natsStream, err := stan.Connect("spiridonov", "daniil", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		panic(err)
	}
	defer natsStream.Close()

	db := postgres.New(cfg.Postgres)
	recovery := memory.Recovery(db)
	cashe, err := memory.New(ctx, recovery)
	if err != nil {
		log.Println("Data could not be recovered, error:", err)
	}
	// TODO проверить правильно ли я решил головоломку с ссылками на cashe
	service := service.New(db, cashe)

	//c := memory.NewMemory(cashe)
	err = subscription.Worker(ctx, natsStream, cashe, service)
	if err != nil {
		log.Println("[worker]", err)
	}

	r := router.NewServer(service)
	err = r.Listen(":" + cfg.Service.Port)
	if err != nil {
		panic(err)
	}
}
