# services/notification/Dockerfile
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /notification-service ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /notification-service .

EXPOSE ${NOTIFICATION_PORT}
CMD ["./notification-service"]
