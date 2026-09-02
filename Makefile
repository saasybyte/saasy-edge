.PHONY: run build generate bruno proto api sqlc clean migrate-up migrate-down

# Run the API server
run:
	go run ./cmd/api

# Build binary
build:
	go build -o bin/api ./cmd/api

# Full codegen: protobuf, OpenAPI, sqlc, Bruno
generate: proto api sqlc bruno

# Regenerate protobuf/gRPC
proto:
	go generate ./pkg/pb/...

# Regenerate OpenAPI server
api:
	go generate ./pkg/api/...

# Regenerate sqlc queries
sqlc:
	sqlc generate

# Regenerate Bruno collection from OpenAPI
bruno:
	./scripts/bruno-gen.sh

# Clean build artifacts
clean:
	rm -rf bin/

# Run database migrations up
migrate-up:
	sh -c '. ./.env && migrate -path db/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" up'

# Roll back last migration
migrate-down:
	sh -c '. ./.env && migrate -path db/migrations -database "postgres://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" down 1'
