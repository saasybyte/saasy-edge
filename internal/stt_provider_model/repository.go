package stt_provider_model

import (
	"context"

	"github.com/saasybyte/saasy-edge/db/sqlc"
)

type SqlcRepository struct {
	queries *sqlc.Queries
}

func NewSqlcRepository(queries *sqlc.Queries) *SqlcRepository {
	return &SqlcRepository{queries: queries}
}

func (r *SqlcRepository) List(ctx context.Context) ([]sqlc.ListSTTProviderModelsRow, error) {
	return r.queries.ListSTTProviderModels(ctx)
}
