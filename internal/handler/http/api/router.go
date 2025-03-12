package api

import "github.com/go-chi/chi/v5"

func (h *Handler) AddRoutes(r chi.Router) {
	r.Get("/users/{id}", h.GetUser)
}
