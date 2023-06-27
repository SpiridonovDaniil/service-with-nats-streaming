package subscription

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"l0/internal/memory"
	"l0/internal/models"
	"log"
)

//go:generate mockgen -source=worker.go -destination=mocks/mock.go

type service interface {
	Create(ctx context.Context, data json.RawMessage, id string) error
}

func Worker(ctx context.Context, natsStream stan.Conn, cashe *memory.Cashe, service service) error {
	_, err := natsStream.Subscribe("message", func(m *stan.Msg) {
		var user models.User

		err := json.Unmarshal(m.Data, &user)
		if err != nil {
			log.Println(err)
		}

		err = cashe.Write(m.Data, user.OrderUid)
		if err != nil {
			log.Println(err)
		}

		err = service.Create(ctx, m.Data, user.OrderUid)
		if err != nil {
			log.Println(err)
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

//TODO вернуть наверх ошибки
