package stt_provider_model

import (
	"context"

	"github.com/saasybyte/saasy-edge/db/sqlc"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context) ([]sqlc.ListSTTProviderModelsRow, error) {
	return s.repo.List(ctx)
}
