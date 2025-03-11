package postgres

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Port     int    `env:"POSTGRES_PORT"`
	Host     string `env:"POSTGRES_HOST"`
	DB       string `env:"POSTGRES_DB"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	SSLMode  string `env:"POSTGRES_SSLMODE"`
}

func MustLoad() (*Config, error) {
	const op = "dbs.postgres.MustLoad"

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &cfg, nil
}
