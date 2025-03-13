package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"test-api-task/internal/entity"
	"test-api-task/internal/service/userservice"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	const op = "handler.http.api.update_user"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID := chi.URLParam(r, "id")

	var userUpdate entity.User
	if err := json.NewDecoder(r.Body).Decode(&userUpdate); err != nil {
		log.Error("failed to decode request", slog.Any("error", err.Error()))
		http.Error(w, "invalid request", http.StatusBadRequest)

		return
	}

	updatedUser, err := h.us.UpdateUser(userID, &userUpdate)
	fmt.Println(updatedUser, err)
	if err != nil {
		switch {
		case errors.Is(err, userservice.ErrUserNotFound):
			log.Info("user not found", slog.Any("user_id", userID))
			http.Error(w, "user not found", http.StatusNotFound)

		case errors.Is(err, userservice.ErrEmailAlreadyExists):
			log.Info(
				"user with this email already eixsts", slog.Any("user_id", userID),
				slog.Any("user_update", userUpdate),
			)
			http.Error(w, "bad request", http.StatusBadRequest)

		default:
			log.Error(
				"failed to update user", slog.Any("user_id", userID),
				slog.Any("user_update", userUpdate), slog.Any("error", err.Error()),
			)
			http.Error(w, "internal error", http.StatusInternalServerError)
		}

		return
	}

	log.Info("user updated", slog.Any("updated_user", updatedUser))

	resp, err := json.MarshalIndent(updatedUser, "", " ")
	if err != nil {
		log.Error("failed to marshal response", slog.Any("error", err.Error()))
		http.Error(w, "internal error", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		log.Error("failed to write response", slog.Any("error", err.Error()))
		http.Error(w, "internal error", http.StatusInternalServerError)

		return
	}
}
