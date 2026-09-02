# saasy-edge

User-facing API service for [SaasyByte](https://github.com/saasybyte/saasybyte), an open-source real-time AI voice platform.

Edge owns the AI provider/model registry: which LLM, STT, and TTS providers and models are available for a session. It serves the catalog to the web client over REST and to the Orchestrator over gRPC, backed by PostgreSQL.

## How It Fits

- **Serves saasy-web** (REST): provider model selection for the session UI.
- **Serves saasy-orchestrator** (gRPC): provider model catalog validation.
- **Proto types** from the [saasy-proto](https://github.com/saasybyte/saasy-proto) git submodule; REST contracts generated from `api/openapi.yaml` (oapi-codegen); queries via sqlc.

See the [platform overview](https://github.com/saasybyte/saasybyte) for the full architecture.

## Build & Run

Requirements: Go 1.25+, PostgreSQL, [golang-migrate](https://github.com/golang-migrate/migrate) for migrations.

```bash
git submodule update --init   # saasy-proto
make run          # run the API server
make build        # build binary to bin/api
make generate     # regenerate protobuf, OpenAPI server, sqlc, Bruno collection
make migrate-up   # apply migrations (DB credentials from .env)
```

Database and server settings come from environment variables; see `.env.example`. Ports: 8080 (REST), 9090 (gRPC).

Generated code (`pkg/api/`, `pkg/pb/`, `db/sqlc/`) is committed, so building needs none of the codegen tools. `make generate` (only needed when changing protos, the OpenAPI spec, or queries) additionally requires: `protoc` with `protoc-gen-go` and `protoc-gen-go-grpc`, `sqlc`, `oapi-codegen`, and the Bruno CLI (`bru`).

A `Dockerfile` is included; `docker build .` needs no credentials. The Dockerfile is the reference build environment.

## License

Apache-2.0, see [LICENSE](LICENSE).
