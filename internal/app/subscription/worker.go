package subscription

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"l0/internal/memory"
	"l0/internal/models"
	"l0/internal/repository"
	"log"
)

type Worker struct {
	cashe memory.Memory
	repo  repository.Repository
}

func New(cashe memory.Memory, repo repository.Repository) *Worker {
	return &Worker{
		cashe: cashe,
		repo:  repo,
	}
}

type worker interface {
	Write(data json.RawMessage, id string) error
	Read(id string) (json.RawMessage, error)
	InsertData(ctx context.Context, data json.RawMessage, id string) error
}

func Start(ctx context.Context, natsStream stan.Conn, worker worker) error {
	_, err := natsStream.Subscribe("message", func(m *stan.Msg) {
		var user models.User

		err := json.Unmarshal(m.Data, &user)
		if err != nil {
			log.Println(err)
		}

		err = worker.Write(m.Data, user.OrderUid)
		if err != nil {
			log.Println(err)
		}

		err = worker.InsertData(ctx, m.Data, user.OrderUid)
		if err != nil {
			log.Println(fmt.Errorf("[create] failed to write to the database, err: %w", err))
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
