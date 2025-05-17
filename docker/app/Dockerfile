# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN apk add --no-cache git build-base

COPY . ./
COPY ./configs/.air.toml ./

RUN go install github.com/cosmtrek/air@v1.45.0

EXPOSE 8081

CMD ["air"]
