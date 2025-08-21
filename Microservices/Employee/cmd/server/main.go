package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/rhzyyn/Aozora/Microservices/Employee/config"
	"github.com/rhzyyn/Aozora/Microservices/Employee/infra/logger"
	"github.com/rhzyyn/Aozora/Microservices/Employee/infra/postgres"
)

func main() {
	// Panic recovery
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("💥 Panic occurred: %v", r)
		}
	}()

	// Signal-aware context
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Config load failed: %v", err)
	}

	// Init logger
	logg := logger.New(cfg)
	logg.Info("🟢 Logger initialized")

	// Print loaded config (debug)
	logg.Info("🧾 DB USER: %s | HOST: %s | NAME: %s | ENABLED_VAULT=%v\n", cfg.DB.User, cfg.DB.Host, cfg.DB.Name, cfg.Vault.Enabled)

	// Connect to DB
	logg.Info("🔌 Attempting PostgreSQL connection...")
	db, err := postgres.Connect(ctx, cfg, logg)
	if err != nil {
		logg.Fatal("❌ PostgreSQL connection failed", "error", err.Error())
	}
	defer db.Close()
	logg.Info("🗄 PostgreSQL connected")

	// Test DB query
	row := db.Pool.QueryRow(ctx, `SELECT 1`)
	var result int
	if err := row.Scan(&result); err != nil {
		logg.Error("❌ DB test query failed", "error", err.Error())
	} else {
		logg.Info("✅ DB test query success", "result", result)
	}

	// Fiber REST setup
	app := fiber.New(fiber.Config{
		AppName:       cfg.App.Name,
		CaseSensitive: true,
		Prefork:       false,
		StrictRouting: true,
		ServerHeader:  "Aozora",
	})

	// Liveness and Readiness
	app.Get("/livez", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Register Prometheus middleware if enabled
	if cfg.Metrics.Enabled {
		prom := fiberprometheus.New(cfg.App.Name)
		prom.RegisterAt(app, cfg.Metrics.Path)
		app.Use(prom.Middleware)
		logg.Info("📈 Prometheus metrics enabled", "path", cfg.Metrics.Path)
	}

	// Start REST (non-TLS for now)
	go func() {
		logg.Info("🚀 Fiber server running", "port", cfg.App.Port)
		if err := app.Listen(":" + cfg.App.Port); err != nil {
			logg.Fatal("❌ Fiber failed", "error", err.Error())
		}
	}()

	// Wait for shutdown
	<-ctx.Done()
	logg.Info("📴 Shutdown initiated")

	if err := app.Shutdown(); err != nil {
		logg.Error("❌ Fiber shutdown failed", "error", err.Error())
	} else {
		logg.Info("✅ Fiber shutdown complete")
	}

	logg.Info("🛑 Application stopped gracefully")
}
