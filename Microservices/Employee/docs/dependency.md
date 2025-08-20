Aozora Employee Service — Dependency Documentation

This file documents every dependency declared in the go.mod file for the Employee microservice of the Aozora platform. Each dependency is grouped by category and includes a remark on its usage or purpose. This document supports strict compliance with Aozora’s 10/10 elite-tier, global-scale standards.

🧠 Core Runtime & System Utilities

Package

Usage

golang.org/x/crypto

Cryptographic primitives, used for password hashing, secure token generation.

golang.org/x/net

Network-related enhancements and extensions.

golang.org/x/sync

Lightweight primitives like errgroup, Once. Used in worker and handler concurrency.

golang.org/x/sys

Access to low-level OS primitives. Required by Vault and other system-level integrations.

golang.org/x/text

UTF handling, localization. Needed by validator's i18n.

golang.org/x/time

High-precision time utilities (e.g., retry backoff, timeout contexts).

📦 Database Layer

Package

Usage

github.com/jackc/pgx/v5

Primary PostgreSQL driver, high-performance and low-level control.

github.com/pashagolub/pgxmock

Mocking PGX queries for unit testing.

github.com/jackc/pgconn, pgproto3, pgtype, etc.

Internal PGX dependencies for parsing/querying.

📜 Configuration & Env Loading

Package

Usage

github.com/joho/godotenv

Load .env files during local dev and testing.

github.com/mitchellh/mapstructure

Convert map to struct (used in config loader, Vault decoding).

github.com/go-ini/ini

Legacy INI config parser (optional).

🔐 Security & Authorization

Package

Usage

github.com/golang-jwt/jwt/v5

Secure JWT creation, parsing, and claim validation.

github.com/hashicorp/vault/api

Secure secret management (tokens, DB creds).

github.com/open-policy-agent/opa

RBAC/ABAC policy enforcement engine.

🌐 HTTP & Routing

Package

Usage

github.com/gin-gonic/gin

Primary web framework for REST APIs.

github.com/gin-contrib/sse

SSE support for streaming endpoints.

github.com/fsnotify/fsnotify

Hot reload config watcher.

github.com/felixge/httpsnoop

Middleware HTTP metrics collection.

⚙️ Retry, Backoff, & Utility Logic

Package

Usage

github.com/cenkalti/backoff/v4

Exponential retry mechanism.

github.com/hashicorp/go-retryablehttp

Resilient HTTP client with retry support.

github.com/hashicorp/go-multierror

Aggregate multiple errors in one return.

📊 Observability & Logging

Package

Usage

go.opentelemetry.io/otel

Core OpenTelemetry library for tracing.

go.opentelemetry.io/otel/sdk

SDK for exporters and custom spans.

go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin

Middleware for trace propagation in Gin.

go.uber.org/zap

High-performance structured logging.

github.com/stretchr/testify

Used in test assertions.

📬 Messaging & Queue

Package

Usage

github.com/segmentio/kafka-go

Kafka messaging client, async event publishing.

🧱 Validation & i18n

Package

Usage

github.com/go-playground/validator/v10

Request and DTO validation.

github.com/go-playground/universal-translator

I18n for validation messages.

github.com/go-playground/locales

Localization language data.

🔄 Hashing, UUID, and Compression

Package

Usage

github.com/google/uuid

Generation of v4 UUIDs for identifiers.

github.com/cespare/xxhash/v2

Fast hashing algorithm, useful for cache or shard keys.

github.com/klauspost/compress

Optional compression helpers.

🔒 Secure Input Handling

Package

Usage

github.com/hashicorp/go-secure-stdlib/*

String utils and secure parsing (used in secrets, Vault config).

🧪 Fuzzing, WASM, Advanced

Package

Usage

github.com/AdaLogics/go-fuzz-headers

For fuzz testing edge cases.

github.com/bytecodealliance/wasmtime-go/v3

WASM runtime (optional/unused now).

✅ Status: All packages marked for Employee service are validated under Go 1.23.9.
📌 Update this file if packages are promoted to active use or deprecated.