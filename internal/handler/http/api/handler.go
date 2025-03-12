package api

import (
	"log/slog"
	"test-api-task/internal/entity"
)

type UserService interface {
	GetUser(userID string) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(userID string, newUserInfo *entity.User) (*entity.User, error)
	DeleteUser(userID string) error
}

type Handler struct {
	us  UserService
	log *slog.Logger
}

func NewHandler(us UserService, log *slog.Logger) *Handler {
	return &Handler{us: us, log: log}
}
