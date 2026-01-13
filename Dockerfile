FROM golang:1.25.4 AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-w -s -extldflags=-static" \
    -o server \
    .
RUN GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-w -s -extldflags=-static" \
    -o admin-cli \
    ./cmd/admin_cli/admin_cli.go
RUN GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-w -s -extldflags=-static" \
    -o migrate \
    ./cmd/migrate/migrate.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/admin-cli .
COPY --from=builder /app/migrate .
CMD ["./server"]
