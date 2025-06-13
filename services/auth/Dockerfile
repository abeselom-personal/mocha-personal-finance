# services/auth/Dockerfile
FROM golang:1.24.4-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download
RUN apk add --no-cache make
CMD ["air", "-c", ".air.toml"]
