package memory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"l0/internal/repository"
	"sync"
)

type Memory interface {
	Write(data json.RawMessage, id string) error
	Read(id string) (json.RawMessage, error)
}

type CashRecovery struct {
	repo repository.Repository
}

func Recovery(repo repository.Repository) *CashRecovery {
	return &CashRecovery{
		repo: repo,
	}
}

//type C struct {
//	Memory Memory
//}
//
//func NewMemory(cashe Memory) *C {
//	return &C{
//		Memory: cashe,
//	}
//}

//TODO как то некрасиво получилось в воркер загнать интерфейс

type Cashe struct {
	Data  map[string]json.RawMessage
	Mutex sync.Mutex
}

func New(ctx context.Context, recovery *CashRecovery) (*Cashe, error) {
	var cashe Cashe
	err := fmt.Errorf("")
	// TODO как убрать это страшное? и давай проверим указатели еще раз)))
	cashe.Data, err = recovery.repo.GetAll(ctx)
	if err != nil {
		return &cashe, err
	}

	return &cashe, nil
}

func (c *Cashe) Write(data json.RawMessage, id string) error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	if _, ok := c.Data[id]; !ok {
		c.Data[id] = data
	} else {
		err := errors.New("the user already exists in the database")
		return err
	}

	return nil
}

func (c *Cashe) Read(id string) (json.RawMessage, error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	if data, ok := c.Data[id]; !ok {
		err := errors.New("user not found")
		return json.RawMessage{}, err
	} else {
		return data, nil
	}

}
