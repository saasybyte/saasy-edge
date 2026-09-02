# ------------------------------
# Build stage
# ------------------------------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install migrate CLI
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy dependency files first (layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN go build -o bin/api ./cmd/api

# ------------------------------
# Runtime stage
# ------------------------------
FROM alpine

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bin/api ./api

# Copy migrate CLI from builder
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Copy migrations
COPY --from=builder /app/db/migrations ./db/migrations

# Expose ports
EXPOSE 8080 9090

# Run the API
CMD ["./api"]
