-- name: ListLLMProviderModels :many
SELECT id, provider, model_id, display_name
FROM llm_provider_models
WHERE deleted_at IS NULL;
