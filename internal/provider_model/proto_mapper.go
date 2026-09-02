package provider_model

import (
	"github.com/saasybyte/saasy-edge/db/sqlc"
	edgev1 "github.com/saasybyte/saasy-edge/pkg/pb/edge/v1"
)

func toProtoListProviderModelsResponse(models *ProviderModelsResponse) *edgev1.ListProviderModelsResponse {
	return &edgev1.ListProviderModelsResponse{
		Llm: toProtoLLMProviderModel(models.LLM),
		Tts: toProtoTTSProviderModel(models.TTS),
		Stt: toProtoSTTProviderModel(models.STT),
	}
}

func toProtoLLMProviderModel(rows []sqlc.ListLLMProviderModelsRow) []*edgev1.LLMProviderModel {
	result := make([]*edgev1.LLMProviderModel, len(rows))
	for i, row := range rows {
		result[i] = &edgev1.LLMProviderModel{
			Id:          row.ID.String(),
			Provider:    row.Provider,
			ModelId:     row.ModelID,
			DisplayName: row.DisplayName,
		}
	}
	return result
}

func toProtoTTSProviderModel(rows []sqlc.ListTTSProviderModelsRow) []*edgev1.TTSProviderModel {
	result := make([]*edgev1.TTSProviderModel, len(rows))
	for i, row := range rows {
		result[i] = &edgev1.TTSProviderModel{
			Id:          row.ID.String(),
			Provider:    row.Provider,
			ModelId:     row.ModelID,
			DisplayName: row.DisplayName,
		}
	}
	return result
}

func toProtoSTTProviderModel(rows []sqlc.ListSTTProviderModelsRow) []*edgev1.STTProviderModel {
	result := make([]*edgev1.STTProviderModel, len(rows))
	for i, row := range rows {
		result[i] = &edgev1.STTProviderModel{
			Id:          row.ID.String(),
			Provider:    row.Provider,
			ModelId:     row.ModelID,
			DisplayName: row.DisplayName,
		}
	}
	return result
}
