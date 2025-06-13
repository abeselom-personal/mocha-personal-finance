# services/sms-parser/Dockerfile
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /sms-parser-service ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /sms-parser-service .

EXPOSE ${SMS_PARSER_PORT}
CMD ["./sms-parser-service"]
