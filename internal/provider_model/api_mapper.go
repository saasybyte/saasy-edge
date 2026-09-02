package provider_model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/saasybyte/saasy-edge/db/sqlc"
	"github.com/saasybyte/saasy-edge/pkg/api"
)

func toAPIProviderModelsResponse(models *ProviderModelsResponse) api.ProviderModelsResponse {
	return api.ProviderModelsResponse{
		Llm: toAPILLMProviderModels(models.LLM),
		Tts: toAPITTSProviderModels(models.TTS),
		Stt: toAPISTTProviderModels(models.STT),
	}
}

func toAPILLMProviderModels(rows []sqlc.ListLLMProviderModelsRow) []api.LLMProviderModel {
	result := make([]api.LLMProviderModel, len(rows))
	for i, row := range rows {
		result[i] = api.LLMProviderModel{
			Id:          pgtypeToUUID(row.ID),
			Provider:    row.Provider,
			ModelId:     row.ModelID,
			DisplayName: row.DisplayName,
		}
	}
	return result
}

func toAPITTSProviderModels(rows []sqlc.ListTTSProviderModelsRow) []api.TTSProviderModel {
	result := make([]api.TTSProviderModel, len(rows))
	for i, row := range rows {
		result[i] = api.TTSProviderModel{
			Id:          pgtypeToUUID(row.ID),
			Provider:    row.Provider,
			ModelId:     row.ModelID,
			DisplayName: row.DisplayName,
		}
	}
	return result
}

func toAPISTTProviderModels(rows []sqlc.ListSTTProviderModelsRow) []api.STTProviderModel {
	result := make([]api.STTProviderModel, len(rows))
	for i, row := range rows {
		result[i] = api.STTProviderModel{
			Id:          pgtypeToUUID(row.ID),
			Provider:    row.Provider,
			ModelId:     row.ModelID,
			DisplayName: row.DisplayName,
		}
	}
	return result
}

func pgtypeToUUID(p pgtype.UUID) uuid.UUID {
	return uuid.UUID(p.Bytes)
}
