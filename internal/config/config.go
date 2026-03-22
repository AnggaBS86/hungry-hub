package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppPort       string
	DB            DBConfig
	RunMigrations bool
	MigrationsDir string
	RunSeed       bool
}

type DBConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	TLS             bool
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func Load() (Config, error) {
	var cfg Config

	cfg.AppPort = getEnv("APP_PORT", "8080")
	cfg.DB.Host = getEnv("DB_HOST", "127.0.0.1")
	cfg.DB.Port = getEnvInt("DB_PORT", 3306)
	cfg.DB.User = getEnv("DB_USER", "root")
	cfg.DB.Password = getEnv("DB_PASSWORD", "")
	cfg.DB.Name = getEnv("DB_NAME", "hungryhub")
	cfg.DB.TLS = getEnvBool("DB_TLS", false)
	cfg.DB.MaxOpenConns = getEnvInt("DB_MAX_OPEN_CONNS", 10)
	cfg.DB.MaxIdleConns = getEnvInt("DB_MAX_IDLE_CONNS", 5)
	cfg.DB.ConnMaxLifetime = getEnvDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute)
	cfg.RunMigrations = getEnvBool("RUN_MIGRATIONS", true)
	cfg.MigrationsDir = getEnv("MIGRATIONS_DIR", "migrations")
	cfg.RunSeed = getEnvBool("RUN_SEED", false)

	if cfg.DB.Host == "" || cfg.DB.User == "" || cfg.DB.Name == "" {
		return Config{}, fmt.Errorf("missing required db config")
	}

	return cfg, nil
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}

func getEnvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		n, err := strconv.Atoi(v)
		if err == nil {
			return n
		}
	}

	return def
}

func getEnvBool(key string, def bool) bool {
	if v := os.Getenv(key); v != "" {
		b, err := strconv.ParseBool(v)
		if err == nil {
			return b
		}
	}

	return def
}

func getEnvDuration(key string, def time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		d, err := time.ParseDuration(v)
		if err == nil {
			return d
		}
	}

	return def
}
