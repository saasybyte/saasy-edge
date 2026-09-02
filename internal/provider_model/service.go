package provider_model

import (
	"context"

	"github.com/saasybyte/saasy-edge/db/sqlc"
	"github.com/saasybyte/saasy-edge/internal/llm_provider_model"
	"github.com/saasybyte/saasy-edge/internal/stt_provider_model"
	"github.com/saasybyte/saasy-edge/internal/tts_provider_model"
)

type Service struct {
	llmProviderModelService *llm_provider_model.Service
	ttsProviderModelService *tts_provider_model.Service
	sttProviderModelService *stt_provider_model.Service
}

func NewService(
	llmProviderModelService *llm_provider_model.Service,
	ttsProviderModelService *tts_provider_model.Service,
	sttProviderModelService *stt_provider_model.Service,
) *Service {
	return &Service{
		llmProviderModelService: llmProviderModelService,
		ttsProviderModelService: ttsProviderModelService,
		sttProviderModelService: sttProviderModelService,
	}
}

func (s *Service) ListProviderModels(ctx context.Context) (*ProviderModelsResponse, error) {
	llmProviderModels, err := s.llmProviderModelService.List(ctx)
	if err != nil {
		return nil, err
	}

	ttsProviderModels, err := s.ttsProviderModelService.List(ctx)
	if err != nil {
		return nil, err
	}

	sttProviderModels, err := s.sttProviderModelService.List(ctx)
	if err != nil {
		return nil, err
	}

	if llmProviderModels == nil {
		llmProviderModels = []sqlc.ListLLMProviderModelsRow{}
	}
	if ttsProviderModels == nil {
		ttsProviderModels = []sqlc.ListTTSProviderModelsRow{}
	}
	if sttProviderModels == nil {
		sttProviderModels = []sqlc.ListSTTProviderModelsRow{}
	}

	return &ProviderModelsResponse{
		LLM: llmProviderModels,
		TTS: ttsProviderModels,
		STT: sttProviderModels,
	}, nil
}
