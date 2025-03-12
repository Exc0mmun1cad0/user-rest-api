package http

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
)

type HandlerRouter interface {
	AddRoutes(r chi.Router)
}

type Router struct {
	router chi.Router
}

func NewRouter() *Router {
	return &Router{router: chi.NewRouter()}
}

func (r *Router) WithHandler(h HandlerRouter, log *slog.Logger) *Router {
	h.AddRoutes(r.router)

	// TODO: middlewares

	return r
}
