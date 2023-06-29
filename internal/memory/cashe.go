package memory

import (
	"context"
	"encoding/json"
	"sync"

	"l0/internal/models"
	"l0/internal/repository"
)

type Memory interface {
	Write(_ context.Context, data json.RawMessage, id string) error
	Read(_ context.Context, id string) (json.RawMessage, error)
	Recover(ctx context.Context, repo repository.Repository, cashe *Cashe) (*Cashe, error)
}

type Cashe struct {
	data  map[string]json.RawMessage
	mutex sync.Mutex
}

func New() *Cashe {
	m := make(map[string]json.RawMessage)
	var cashe Cashe
	cashe.data = m
	return &cashe
}

func (c *Cashe) Write(_ context.Context, data json.RawMessage, id string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, ok := c.data[id]; !ok {
		c.data[id] = data
	} else {
		return models.ErrAlreadyInTheDB
	}

	return nil
}

func (c *Cashe) Read(_ context.Context, id string) (json.RawMessage, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if data, ok := c.data[id]; !ok {
		err := models.ErrNotFound
		return json.RawMessage{}, err
	} else {
		return data, nil
	}

}

func (c *Cashe) Recover(ctx context.Context, repo repository.Repository, cashe *Cashe) (*Cashe, error) {
	var err error

	cashe.data, err = repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return cashe, nil
}
