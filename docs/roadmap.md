# 🛣️ Project Roadmap — Event-Driven Orders Service

This roadmap outlines step-by-step evolution of the project from MVP to production-grade, Kubernetes-deployable microservice.

---

## ✅ Level 1: MVP Core (DONE)
- [x] Orders Create usecase
- [x] Outbox pattern
- [x] Kafka integration (Redpanda)
- [x] SQLBoiler ORM
- [x] PostgreSQL + Migrations
- [x] Clean Architecture (domain/usecase/infra)
- [x] Docker Compose setup
- [x] Makefile workflow

---

## ⚙️ Level 2: Dev Experience & API
- [ ] Swagger/OpenAPI docs
- [ ] `.env` + config loader (cleanenv / Viper)
- [ ] Healthcheck route `/healthz`
- [ ] Logging (zap or slog) with trace IDs
- [ ] Graceful shutdown support

---

## 📊 Level 3: Observability & Resilience
- [ ] Prometheus metrics (`/metrics`)
- [ ] OpenTelemetry tracing (Jaeger compatible)
- [ ] Retry strategy + Kafka DLQ
- [ ] Circuit breaker middleware
- [ ] Rate limiting middleware

---

## 🧑‍🤝‍🧑 Level 4: Domain Expansion
- [ ] Update / delete orders
- [ ] User model & CRUD
- [ ] Webhooks for external services
- [ ] Email or Slack notifications (Kafka event)

---

## ☁️ Level 5: Infrastructure & CI/CD
- [ ] GitHub Actions CI pipeline
- [ ] DockerHub or GHCR image build
- [ ] Basic integration tests
- [ ] Static code analysis (golangci-lint)
- [ ] Secrets management (Vault / dotenv-secrets)

---

## ☸️ Level 6: Kubernetes Integration
- [ ] Helm Chart for the project
- [ ] K8s manifests: Deployments, Services, ConfigMap
- [ ] Postgres via Bitnami Helm Chart
- [ ] Kafka via Redpanda / Strimzi
- [ ] Ingress + TLS (cert-manager)
- [ ] HorizontalPodAutoscaler
- [ ] Readiness / Liveness probes

---

## Notes
Feel free to fork and adapt this roadmap to your needs.