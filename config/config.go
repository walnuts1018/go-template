package config

import (
	"log/slog"
	"reflect"
	"strings"
	"time"

	"github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	ServerPort string     `env:"SERVER_PORT" envDefault:"8080"`
	ServerURL  string     `env:"SERVER_URL" envDefault:"localhost"`
	LogLevel   slog.Level `env:"LOG_LEVEL"`

	// --------------------- PostgreSQL ---------------------
	PSQLDSN      string `env:"PSQL_DSN" envDefault:""` // If PSQL_DSN is set, other PSQL_* variables will be ignored
	PSQLHost     string `env:"PSQL_HOST" envDefault:"localhost"`
	PSQLPort     string `env:"PSQL_PORT" envDefault:"5432"`
	PSQLDatabase string `env:"PSQL_DATABASE" envDefault:"tobechanged"`
	PSQLUser     string `env:"PSQL_USER" envDefault:"postgres"`
	PSQLPassword string `env:"PSQL_PASSWORD" envDefault:"postgres"`
	PSQLSSLMode  string `env:"PSQL_SSL_MODE" envDefault:"disable"`
	PSQLTimeZone string `env:"PSQL_TIMEZONE" envDefault:"Asia/Tokyo"`
	// -------------------------------------------------------

}

func Load() (Config, error) {
	cfg := Config{}
	if err := env.ParseWithOptions(&cfg, env.Options{
		FuncMap: map[reflect.Type]env.ParserFunc{
			reflect.TypeOf(slog.Level(0)):    returnAny(ParseLogLevel),
			reflect.TypeOf(time.Duration(0)): returnAny(time.ParseDuration),
		},
	}); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func returnAny[T any](f func(v string) (t T, err error)) func(v string) (any, error) {
	return func(v string) (any, error) {
		t, err := f(v)
		return any(t), err
	}
}

func ParseLogLevel(v string) (slog.Level, error) {
	switch strings.ToLower(v) {
	case "":
		return slog.LevelInfo, nil
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		slog.Warn("Invalid log level, use default level: info")
		return slog.LevelInfo, nil
	}
}
