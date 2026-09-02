package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/saasybyte/saasy-edge/db/sqlc"
	"github.com/saasybyte/saasy-edge/internal/llm_provider_model"
	"github.com/saasybyte/saasy-edge/internal/provider_model"
	"github.com/saasybyte/saasy-edge/internal/stt_provider_model"
	"github.com/saasybyte/saasy-edge/internal/tts_provider_model"
	"github.com/saasybyte/saasy-edge/pkg/api"
	edgev1 "github.com/saasybyte/saasy-edge/pkg/pb/edge/v1"
)

type Handlers struct {
	ProviderModel     *provider_model.Handler
	ProviderModelGRPC *provider_model.GRPCHandler
}

func InitializeHandlers(queries *sqlc.Queries) *Handlers {
	llmRepo := llm_provider_model.NewSqlcRepository(queries)
	ttsRepo := tts_provider_model.NewSqlcRepository(queries)
	sttRepo := stt_provider_model.NewSqlcRepository(queries)

	llmService := llm_provider_model.NewService(llmRepo)
	ttsService := tts_provider_model.NewService(ttsRepo)
	sttService := stt_provider_model.NewService(sttRepo)
	providerModelService := provider_model.NewService(llmService, ttsService, sttService)

	providerModelHandler := provider_model.NewHandler(providerModelService)

	providerModelGRPCHandler := provider_model.NewGRPCHandler(providerModelService)

	return &Handlers{
		ProviderModel:     providerModelHandler,
		ProviderModelGRPC: providerModelGRPCHandler,
	}
}

func RegisterRoutes(routerGroup *gin.RouterGroup, h *Handlers) {
	api.RegisterHandlers(routerGroup, h.ProviderModel)
}

func RegisterGRPCServices(grpcServer *grpc.Server, h *Handlers) {
	edgev1.RegisterEdgeServiceServer(grpcServer, h.ProviderModelGRPC)
}
