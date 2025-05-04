# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main ./cmd/consumer

# Final stage
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8082

CMD ["./main"]
