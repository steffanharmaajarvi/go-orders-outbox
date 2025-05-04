# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main ./cmd/app

# Final stage
FROM alpine:latest

RUN apk add --no-cache curl

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate /usr/local/bin/

WORKDIR /root/
COPY --from=builder /app/main .
COPY ./migrations/ ./migrations/
EXPOSE 8081

CMD ["./cmd/app"]
