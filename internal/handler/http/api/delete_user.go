package api

import (
	"errors"
	"log/slog"
	"net/http"
	"test-api-task/internal/service/userservice"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	const op = "handler.http.api.delete_user"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID := chi.URLParam(r, "id")
	err := h.us.DeleteUser(userID)
	if errors.Is(err, userservice.ErrUserNotFound) {
		log.Info("user not found", slog.String("user_id", userID))
		http.Error(w, "invalid request", http.StatusNotFound)

		return
	}
	if err != nil {
		log.Info("failed to delete user")
		http.Error(w, "user info not found", http.StatusInternalServerError)

		return
	}

	log.Info("user deleted", slog.Any("id", userID))

	w.WriteHeader(http.StatusNoContent)
}
