package provider_model

import (
	"context"

	edgev1 "github.com/saasybyte/saasy-edge/pkg/pb/edge/v1"
)

type GRPCHandler struct {
	edgev1.UnimplementedEdgeServiceServer // method promotion
	service *Service
}

func NewGRPCHandler(service *Service) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (h *GRPCHandler) ListProviderModels(ctx context.Context, req *edgev1.ListProviderModelsRequest) (*edgev1.ListProviderModelsResponse, error) {
	models, err := h.service.ListProviderModels(ctx)
	if err != nil {
		return nil, err
	}

	return toProtoListProviderModelsResponse(models), nil
}
