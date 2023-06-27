package repository

import (
	"context"
	"encoding/json"
)

type Repository interface {
	InsertData(ctx context.Context, data json.RawMessage, id string) error
	GetAll(ctx context.Context) (map[string]json.RawMessage, error)
}
