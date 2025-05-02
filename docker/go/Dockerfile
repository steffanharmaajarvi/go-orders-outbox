# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main ./cmd/app

# Final stage
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8081

CMD ["./main"]
