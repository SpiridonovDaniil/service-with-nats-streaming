package memory

import (
	"context"
	"encoding/json"
	"errors"
	"l0/internal/repository"
	"sync"
)

type Memory interface {
	Write(_ context.Context, data json.RawMessage, id string) error
	Read(_ context.Context, id string) (json.RawMessage, error)
}

type Cashe struct {
	data  map[string]json.RawMessage
	mutex sync.Mutex
}

func New(ctx context.Context, repo repository.Repository) (*Cashe, error) {
	var cashe Cashe
	var err error

	cashe.data, err = repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return &cashe, nil
}

func (c *Cashe) Write(_ context.Context, data json.RawMessage, id string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, ok := c.data[id]; !ok {
		c.data[id] = data
	} else {
		err := errors.New("the user already exists in the database")
		return err
	}

	return nil
}

func (c *Cashe) Read(_ context.Context, id string) (json.RawMessage, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if data, ok := c.data[id]; !ok {
		err := errors.New("user not found")
		return json.RawMessage{}, err
	} else {
		return data, nil
	}

}
