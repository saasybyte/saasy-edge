package llm_provider_model

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

func (r *SqlcRepository) List(ctx context.Context) ([]sqlc.ListLLMProviderModelsRow, error) {
	return r.queries.ListLLMProviderModels(ctx)
}
