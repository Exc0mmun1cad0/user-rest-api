package app

import (
	"log/slog"
	"test-api-task/internal/handler/http"
	httpapi "test-api-task/internal/handler/http/api"
)

func (a *App) StartHTTPServer() {
	// TODO: add graceful shutdown here
	a.startHttpServer()
}

func (a *App) startHttpServer() {
	handler := httpapi.NewHandler(a.c.GetUserService(), a.log)

	router := http.NewRouter()
	router.WithHandler(handler, a.log)

	srv := http.NewServer(a.cfg.HTTP)
	srv.RegisterRoutes(router)

	a.log.Info("start HTTP server", slog.String("host", a.cfg.HTTP.Host), slog.Int("port", a.cfg.HTTP.Port))
	err := srv.Start()
	if err != nil {
		a.log.Error("failed to start http server", slog.Any("error", err))
	}
}
