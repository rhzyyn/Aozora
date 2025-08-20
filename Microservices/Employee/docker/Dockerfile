FROM golang:1.22.4-slim AS builder
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/* && go build -o app-employee ./cmd/server

FROM gcr.io/distroless/cc-debian12
WORKDIR /app
COPY --from=builder /app/app-employee .
COPY --from=builder /app/.env.local .env.local
EXPOSE 8081
EXPOSE 5051
ENTRYPOINT ["/app/app-employee"]
