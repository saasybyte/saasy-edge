package stt_provider_model

import (
	"context"

	"github.com/saasybyte/saasy-edge/db/sqlc"
)

type Repository interface {
	List(ctx context.Context) ([]sqlc.ListSTTProviderModelsRow, error)
}
