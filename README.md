# 🌀 Aozora Platform

Aozora is architected for performance, resilience, and global scalability. Built from the ground up using modern backend best practices, Aozora enables enterprises to manage employees, authentication, messaging, and reporting with the robustness of cloud-native systems — without vendor lock-in.

This project exemplifies what an elite backend infrastructure should look like in 2025: poly-cloud compatible, container-first, observable, and DevOps-ready.

![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go)
![Terraform](https://img.shields.io/badge/Terraform-IaC-623ce4?logo=terraform)
![CI/CD](https://img.shields.io/badge/GitHub-Actions-blue?logo=github-actions)

---

## 🚀 Why This Stack?

### 💻 Language: **Go (Golang)**

Go is a battle-tested, statically typed systems language designed for simplicity, speed, and concurrency — ideal for scalable microservices. It compiles to a single binary, making it perfect for container-based deployment and low-overhead orchestration.

### 🧩 Architecture: **DDD + Hexagonal**

Each microservice is built using **Domain-Driven Design** with **Hexagonal (Ports & Adapters)** architecture, enabling:

- Clear separation of concerns
- Replaceable infrastructure (databases, queues, etc.)
- Testable and maintainable business logic
- Event-driven and command-based workflow support

### ☁️ Infrastructure-as-Code: **Terraform**

Terraform enables us to declaratively define infrastructure across:

- **AWS** (Amazon ECS, MSK, RDS, ALB)
- **GCP** (GKE, CloudSQL, Pub/Sub)
- **Azure** (AKS, PostgreSQL Flex, Event Grid)

It guarantees **reproducibility**, **auditability**, and **easy rollback** via version control.

### 📡 Messaging: **Apache Kafka**

Kafka is the backbone of our **event-driven architecture**, supporting:

- Loose coupling between services
- Scalable async processing
- Robust retry and dead-letter patterns
- Future support for stream analytics and auditing

### 🔐 Secrets: **AWS Secrets Manager / HashiCorp Vault**

Secrets are managed securely using Vault or cloud-native services to:

- Prevent leaking credentials in code
- Enable per-service scoped secrets
- Allow automatic rotation for long-lived systems

### 📊 Observability: **Prometheus + Grafana + Loki**

From day one, we’ve baked in deep observability:

- Real-time metrics via `/metrics` endpoint
- Centralized logging with Loki
- Dashboards and alerts with Grafana

### 🐳 Container-first: **Docker, Kubernetes, Swarm**

Every microservice is deployable to:

- **Local Docker** for dev
- **Docker Swarm** for lightweight production
- **Kubernetes** (or ECS/Fargate) for global-scale HA

### 🔁 CI/CD: **GitHub Actions**

Push-to-deploy infrastructure and microservices with:

- Per-environment deployment logic
- Lint, test, build, and push workflows
- Auto-tagged release pipelines and container registries

---

### 📁 Project Structure (Monorepo)

```bash
/Microservices/              # Microservices (Mailer, OAuth, Employee, etc.)
  ├── cmd/                   # Entrypoints (server, worker)
  ├── internal/              # DDD core: domain, usecase, event, handler
  ├── infra/                 # Ports: kafka, db, redis, logger, smtp, etc.
  ├── proto/                 # gRPC definitions
  ├── config/                # Environment/config loader
  └── test/                  # Unit, integration, contracts

/Infra/                      # Terraform IAC for all clouds
  ├── environments/          # AWS, GCP, Azure split
  ├── modules/               # Reusable infra components (vpc, kafka, etc.)
  ├── local/                 # Docker, Swarm, K8s, Caddy
  └── observability/         # Prometheus, Grafana, Loki

/Libs/                       # Shared Go libraries: logger, jwt, kafka client
/Docs/                       # Architecture diagrams, decision records
/DevOps/                     # GitHub Actions, CI/CD scripts
```

## 🧪 Development Ready

- `make run` / `make test` for every service
- `terraform apply` for infrastructure
- `.env.example` files for safe bootstrapping
- Docker and Kubernetes YAMLs ready to deploy in `/local/`

---

## 🌍 Built for Teams, Built for Scale

Whether you're:

- a **solo dev** shipping fast with Docker,
- a **team deploying to multi-cloud**,
- or a **global org** scaling across zones with Kubernetes,

**Aozora is designed to be your launchpad.**

---

> **Designed to scale. Engineered to last. Built like it's already running in production.**

---

## 🧑‍💻 Maintainer

Created by **Fauziyyan Rahman**  
📫 [Email](mailto:Fauziyyan.Rahman@gmail.com) • 🌐 [LinkedIn](https://linkedin.com/in/fauziyyan-rahman)  
Licensed under the [MIT License](./LICENSE)
