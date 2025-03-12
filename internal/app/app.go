package app

import (
	"fmt"
	"log/slog"
	"test-api-task/internal/config"
	pg "test-api-task/internal/dbs/postgres"

	"github.com/jmoiron/sqlx"
)

type App struct {
	cfg config.Config

	c *Container

	pgsqlx *sqlx.DB

	log *slog.Logger
}

func NewApp(cfg config.Config) (*App, error) {
	const op = "app.NewApp"

	app := &App{
		cfg: cfg,
	}

	app.initLogger()

	psqlConfig, err := pg.MustLoad()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	psqlConn, err := pg.New(psqlConfig)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	app.pgsqlx = psqlConn

	app.c = NewContainer(app.pgsqlx)

	return app, nil
}
