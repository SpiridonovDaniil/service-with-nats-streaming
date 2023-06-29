package subscription

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"l0/internal/memory"
	"l0/internal/models"
	"l0/internal/repository"

	"github.com/nats-io/stan.go"
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

func (w *Worker) Start(ctx context.Context, natsStream stan.Conn) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		default:
			_, err := natsStream.Subscribe("message", func(m *stan.Msg) {
				var user models.User

				err := json.Unmarshal(m.Data, &user)
				if err != nil {
					log.Println(fmt.Errorf("[worker] %w", err))
					return
				}

				err = w.repo.InsertData(ctx, m.Data, user.OrderUid)
				if err != nil {
					log.Println(fmt.Errorf("[worker] failed to write to the database, err: %w", err))
					return
				}

				err = w.cashe.Write(ctx, m.Data, user.OrderUid)
				if err != nil {
					log.Println(fmt.Errorf("[worker] failed to write to the cashe, err: %w", err))
					return
				}
			}, stan.StartWithLastReceived())
			if err != nil {
				log.Println(fmt.Errorf("[worker] subscription error: %w", err))
				return
			}
		case <-ctx.Done():
			log.Println(ctx.Err())
			return
		}
	}()
	wg.Wait()
}
