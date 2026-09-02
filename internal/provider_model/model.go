package provider_model

import "github.com/saasybyte/saasy-edge/db/sqlc"

type ProviderModelsResponse struct {
	LLM []sqlc.ListLLMProviderModelsRow `json:"llm"`
	TTS []sqlc.ListTTSProviderModelsRow `json:"tts"`
	STT []sqlc.ListSTTProviderModelsRow `json:"stt"`
}
