package health

import (
	"context"

	"github.com/rhzyyn/Aozora/Microservices/Employee/infra/postgres"
	"github.com/rhzyyn/Aozora/Microservices/Employee/infra/ratelimiter"
)

type HealthCheck struct {
	DB          *postgres.DB
	RateLimiter ratelimiter.Strategy
}

func New(db *postgres.DB, rl ratelimiter.Strategy) *HealthCheck {
	return &HealthCheck{
		DB:          db,
		RateLimiter: rl,
	}
}

func (h *HealthCheck) Livez() bool {
	// Liveness = process not dead
	return true
}

func (h *HealthCheck) Readyz(ctx context.Context) bool {
	// Readiness = infra is healthy
	if err := h.DB.Pool.Ping(ctx); err != nil {
		return false
	}
	if err := h.RateLimiter.HealthCheck(ctx); err != nil {
		return false
	}
	return true
}
