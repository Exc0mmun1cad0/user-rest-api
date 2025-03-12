package app

import (
	"log/slog"
	"os"
)

func (a *App) initLogger() {
	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	a.log = log
}
