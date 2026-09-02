-- name: ListSTTProviderModels :many
SELECT id, provider, model_id, display_name
FROM stt_provider_models
WHERE deleted_at IS NULL;
