package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"test-api-task/internal/service/userservice"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	const op = "handler.http.api.get_user"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID := chi.URLParam(r, "id")

	user, err := h.us.GetUser(userID)
	if errors.Is(err, userservice.ErrUserNotFound) {
		log.Info("user info not found", slog.String("user_id", userID))
		http.Error(w, "invalid request", http.StatusNotFound)

		return
	}
	if err != nil {
		log.Info("failed to get user information")
		http.Error(w, "user info not found", http.StatusInternalServerError)

		return
	}

	log.Info("user information found", slog.Any("id", userID))

	resp, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Error("failed to marhal response")
		http.Error(w, "internal error", http.StatusInternalServerError)

		return
	}

	_, err = w.Write(resp)
	if err != nil {
		log.Error("failed to write response")
		http.Error(w, "internal error", http.StatusInternalServerError)

		return
	}
}
