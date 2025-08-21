package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name     string
	Port     string
	Env      string
	LogLevel string
	Version  string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type TLSConfig struct {
	CertFile string
	KeyFile  string
}

type VaultConfig struct {
	Address string
	Token   string
	Enabled bool
}

type OTelConfig struct {
	Endpoint string
}

type GRPCConfig struct {
	Port string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type RateLimitConfig struct {
	MaxRequestsPerMinute int
}

type TenantConfig struct {
	HeaderName string // default: "X-Tenant-ID"
}

type MetricsConfig struct {
	Enabled bool
	Path    string // default: "/metrics"
}

type JWTConfig struct {
	PublicKey string
	Issuer    string
}

type Config struct {
	App       AppConfig
	DB        DBConfig
	TLS       TLSConfig
	Vault     VaultConfig
	OTel      OTelConfig
	GRPC      GRPCConfig
	Redis     RedisConfig
	RateLimit RateLimitConfig
	Tenant    TenantConfig
	Metrics   MetricsConfig
	JWT       JWTConfig
}

// Load initializes configuration from .env.local, config.yaml, and optionally Vault
func Load() (*Config, error) {
	// Try to load .env.local from common locations
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load("../.env.local")
	_ = godotenv.Load("../../.env.local")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Look in current dir and parents (handles running from cmd/server)
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	_ = viper.ReadInConfig() // optional

	viper.SetEnvPrefix("EMPLOYEE")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	cfg := &Config{
		App: AppConfig{
			Name:     viper.GetString("APP.NAME"),
			Port:     Fallback(viper.GetString("APP.PORT"), "8081"),
			Env:      Fallback(viper.GetString("APP.ENV"), "local"),
			LogLevel: Fallback(viper.GetString("APP.LOGLEVEL"), "debug"),
			Version:  Fallback(viper.GetString("APP.VERSION"), "v1.0.0"),
		},
		DB: DBConfig{
			Host:     Fallback(viper.GetString("DB.HOST"), "localhost"),
			Port:     Fallback(viper.GetString("DB.PORT"), "5432"),
			User:     viper.GetString("DB.USER"),
			Password: viper.GetString("DB.PASSWORD"),
			Name:     viper.GetString("DB.NAME"),
			SSLMode:  Fallback(viper.GetString("DB.SSLMODE"), "disable"),
		},
		TLS: TLSConfig{
			CertFile: Fallback(viper.GetString("TLS.CERTFILE"), "./certs/server.crt"),
			KeyFile:  Fallback(viper.GetString("TLS.KEYFILE"), "./certs/server.key"),
		},
		Vault: VaultConfig{
			Address: Fallback(viper.GetString("VAULT.ADDRESS"), "http://localhost:8200"),
			Token:   viper.GetString("VAULT.TOKEN"),
			Enabled: viper.GetBool("VAULT_ENABLED"), // ✅ CORRECT: mapped from EMPLOYEE_VAULT_ENABLED
		},
		OTel: OTelConfig{
			Endpoint: Fallback(viper.GetString("OTEL.EXPORTER.OTLP.ENDPOINT"), "localhost:4317"),
		},
		GRPC: GRPCConfig{
			Port: Fallback(viper.GetString("GRPC.PORT"), "5051"),
		},
		Redis: RedisConfig{
			Addr:     Fallback(viper.GetString("REDIS.ADDR"), "localhost:6379"),
			Password: viper.GetString("REDIS.PASSWORD"),
			DB:       viper.GetInt("REDIS.DB"),
		},
		RateLimit: RateLimitConfig{
			MaxRequestsPerMinute: viper.GetInt("RATELIMIT.MAX_REQUESTS_PER_MINUTE"),
		},
		Tenant: TenantConfig{
			HeaderName: Fallback(viper.GetString("TENANT.HEADER_NAME"), "X-Tenant-ID"),
		},
		Metrics: MetricsConfig{
			Enabled: viper.GetBool("METRICS.ENABLED"),
			Path:    Fallback(viper.GetString("METRICS.PATH"), "/metrics"),
		},
		JWT: JWTConfig{
			PublicKey: viper.GetString("JWT.PUBLIC_KEY"),
			Issuer:    viper.GetString("JWT.ISSUER"),
		},
	}

	return cfg, nil
}

func Fallback(val, defaultVal string) string {
	if val == "" {
		return defaultVal
	}
	return val
}
