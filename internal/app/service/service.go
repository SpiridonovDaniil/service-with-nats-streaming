package service

import (
	"context"
	"encoding/json"
	"fmt"
	"l0/internal/memory"
	"l0/internal/repository"
)

type Service struct {
	repo  repository.Repository
	cashe memory.Memory
}

func New(repo repository.Repository, cashe memory.Memory) *Service {
	return &Service{
		repo:  repo,
		cashe: cashe,
	}
}

func (s *Service) Get(id string) (json.RawMessage, error) {
	resp, err := s.cashe.Read(id)
	if err != nil {
		return resp, fmt.Errorf("[get] %w", err)
	}
	//TODO нужен ли контекст, если мы не обращаемся в базу данных?
	return resp, nil
}

func (s *Service) Create(ctx context.Context, data json.RawMessage, id string) error {
	err := s.repo.Create(ctx, data, id)
	if err != nil {
		return fmt.Errorf("[create] failed to write to the database, err: %w", err)
	}

	return nil
}
