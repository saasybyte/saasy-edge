-- name: ListTTSProviderModels :many
SELECT id, provider, model_id, display_name
FROM tts_provider_models
WHERE deleted_at IS NULL;
