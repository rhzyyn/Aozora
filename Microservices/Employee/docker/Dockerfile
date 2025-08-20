FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go build -o app-employee ./cmd/server

FROM gcr.io/distroless/cc-debian12
WORKDIR /app
COPY --from=builder /app/app-employee .
COPY --from=builder /app/.env.local .env.local
EXPOSE 8081
EXPOSE 5051
ENTRYPOINT ["/app/app-employee"]
