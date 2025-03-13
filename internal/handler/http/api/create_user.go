package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"test-api-task/internal/entity"
	"test-api-task/internal/service/userservice"

	"github.com/go-chi/chi/v5/middleware"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	const op = "handler.http.api.create_user"

	log := h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var newUser entity.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		log.Error("failed to decode request", slog.Any("error", err.Error()))
		http.Error(w, "invalid request", http.StatusNotFound)

		return
	}

	createdUser, err := h.us.CreateUser(&newUser)
	if err != nil {
		if errors.Is(err, userservice.ErrEmailAlreadyExists) {
			log.Info("couldn't create a user because the email is already busy", slog.Any("user_info", newUser))
			http.Error(w, "bad request", http.StatusBadRequest)

			return
		}

		log.Error("failed to create user", slog.Any("user_info", newUser), slog.Any("error", err.Error()))
		http.Error(w, "internal error", http.StatusInternalServerError)

		return
	}

	log.Info("user created", slog.Any("created_user", createdUser))

	resp, err := json.MarshalIndent(createdUser, "", " ")
	if err != nil {
		log.Error("failed to marshal response", slog.Any("error", err.Error()))
		http.Error(w, "internal error", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(resp)
	if err != nil {
		log.Error("failed to write response", slog.Any("error", err.Error()))
		http.Error(w, "internal error", http.StatusInternalServerError)

		return
	}
}
