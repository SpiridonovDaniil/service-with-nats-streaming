package service

import (
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

	return resp, nil
}
