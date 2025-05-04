# 📦 Event-Driven Orders Service (Go + Postgres + Kafka)
![Go Version](https://img.shields.io/badge/go-1.24-blue)
![Docker](https://img.shields.io/badge/containerized-docker-blue)

A minimal event-driven microservice built in Go using the **Transactional Outbox Pattern**, PostgreSQL, and Kafka. Designed for learning and practicing **strong consistency** in distributed systems through **clean architecture** and **Docker-based environments**.

---

## 🧱 Features

- Create order with transactional outbox
- Kafka publisher for asynchronous event delivery
- PostgreSQL with SQL migrations (via golang-migrate)
- Clean, layered architecture (domain → usecase → infra)
- Live reloading in development via [`air`](https://github.com/cosmtrek/air)
- Separate containers for app, consumer, migrator
- Docker Compose setup for local dev

--

## 🚀 Stack

- **Go 1.24+**
- **PostgreSQL**
- **Kafka (Redpanda)**
- **golang-migrate**
- **Air (live reload)**
- **Docker + Docker Compose**

---

## 📂 Project Structure

```
.
├── cmd/                # Application entrypoints (app / consumer)
├── internal/
│   ├── domain/         # Business entities and logic
│   ├── usecase/        # Use cases
│   ├── infrastructure/ # Postgres, Kafka, Outbox, etc.
│   └── setup/          # Dependency injection, wiring
├── migrations/         # SQL migration files
├── configs/            # YAML config files
├── docker/
│   ├── app/            # Dockerfile for main app
│   ├── consumer/       # Dockerfile for Kafka consumer
│   └── migrator/       # Dockerfile for DB migrator
├── .air.toml           # Air config for hot reload
├── docker-compose.yml
└── Makefile
```

---

## 🐳 Docker Compose Setup

Containers:

- `postgres`: the main database
- `redpanda`: Kafka broker
- `migrator`: runs once to apply DB migrations via `wait-and-migrate.sh`
- `app`: main API server
- `consumer`: Kafka consumer reading from outbox topic

---

## 🔁 Air (Live Reload)

Used for development with live code reloading.

1. Install `air`:

   ```bash
   go install github.com/cosmtrek/air@latest
   ```

2. Config is placed in `.air.toml`

3. Dockerfile uses:

   ```Dockerfile
   CMD ["air"]
   ```

Make sure the `.air.toml` is **not overwritten** by volumes in `docker-compose.yml`.

---

## 📦 Makefile Commands

```bash
make start   # init + build
make build   # docker-compose up -d --build
make up      # docker-compose up -d
```

---

## 🧪 TODO

- Add update/delete order
- Retry policy for Kafka delivery
- Observability (logging, metrics)
- Integration tests