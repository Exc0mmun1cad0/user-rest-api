package api

import "github.com/go-chi/chi/v5"

func (h *Handler) AddRoutes(r chi.Router) {
	r.Get("/users/{id}", h.GetUser)
	r.Post("/users", h.CreateUser)
	r.Patch("/users/{id}", h.UpdateUser)
	r.Delete("/users/{id}", h.DeleteUser)
}
