package service

import (
	"context"
	"encoding/json"

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

func (s *Service) Get(ctx context.Context, id string) (json.RawMessage, error) {
	resp, err := s.cashe.Read(ctx, id)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
