package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func New(cfg *Config) (*sqlx.DB, error) {
	const op = "dbs.postgres.New"

	dataSource := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB, cfg.SSLMode,
	)

	conn, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return conn, nil
}
