# services/regex/Dockerfile
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /regex-service ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /regex-service .

EXPOSE ${REGEX_PORT}
CMD ["./regex-service"]
