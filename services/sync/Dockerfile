# services/sync/Dockerfile
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /sync-service ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /sync-service .

EXPOSE ${SYNC_PORT}
CMD ["./sync-service"]
