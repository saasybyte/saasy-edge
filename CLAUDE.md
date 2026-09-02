# saasy-edge

## Commands
- `make run` — run dev server
- `make build` — build binary to `bin/api`
- `make generate` — full codegen (proto, api, sqlc, bruno)
- `make proto` — regenerate protobuf/gRPC
- `make api` — regenerate OpenAPI server
- `make sqlc` — regenerate sqlc queries
- `make bruno` — generate Bruno collection from OpenAPI
- `make migrate-up` — run all pending migrations (DB creds from `.env`)
- `make migrate-down` — roll back last migration (DB creds from `.env`)
- `make clean` — remove build artifacts

## Conventions
- **Domain modules**: each module in `internal/` follows model → repository → service → handler layering.
- **Aggregate module**: `provider_model` composes sub-domain services and exposes both REST and gRPC handlers.
- **DI wiring**: `cmd/api/wire.go` — repos → services → handlers. No DI framework.
- **REST handlers**: implement generated `api.ServerInterface` (oapi-codegen).
- **gRPC handlers**: embed `Unimplemented*Server` for forward compatibility.
- **Mapper files**: separate API and proto conversion (`api_mapper.go`, `proto_mapper.go`) — do not inline mapping logic in handlers.
- **Repositories**: wrap `sqlc.Queries` — do not call sqlc directly from services.
- **OpenAPI summaries**: filename-friendly names (lowercase, hyphens, no spaces). Bruno generates `.bru` filenames from summaries.
- **Logging**: stdlib `log` package. Not `logrus`, `zap`, or `slog`.
- **Codegen outputs**: `pkg/api/`, `pkg/pb/`, `db/sqlc/` are committed. Regenerate via `make generate`.

## Service Boundaries
- **Serves saasy-orchestrator** (gRPC): provider model catalog.
- **Serves web** (REST): provider model selection.
- **Proto types from saasy-proto** (git submodule): do not define proto types locally.
- **Does not own**: auth (saasy-core), media (saasy-sfu), AI inference (saasy-orchestrator), proto schema (saasy-proto).
